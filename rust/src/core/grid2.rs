//! Optimized 2D Grid with support for 4-way and 8-way neighbors
//! and optional space-filling curve layouts (Morton / Hilbert).
//!
//! - `Grid<T>`: generic storage over `Vec<T>`
//! - Layouts: RowMajor (standard), Morton (Z-order) for cache-friendly locality
//! - Neighbor iterators: `neighbors4`, `neighbors8` returning (x,y, &T) or indices
//! - Works for rectangular grids; Morton layout internally pads to next power of two

use std::ops::{Index, IndexMut};

///
pub trait Layout {
    fn new(width: usize, height: usize) -> Self;

    fn len(&self) -> usize;

    fn coord_to_index(&self, x: usize, y: usize) -> usize;

    fn index_to_coord(&self, idx: usize) -> (usize, usize);

    fn iter_coords<'a>(&'a self) -> impl Iterator<Item = (usize, usize)> {
        (0..self.len()).map(move |idx| self.index_to_coord(idx))
    }
}

/// Standard row-major layout.
pub struct RowMajor(usize, usize);

impl Layout for RowMajor {
    fn new(width: usize, height: usize) -> Self {
        Self(width, height)
    }

    fn len(&self) -> usize {
        self.0 * self.1
    }

    fn coord_to_index(&self, x: usize, y: usize) -> usize {
        y * self.0 + x
    }

    fn index_to_coord(&self, idx: usize) -> (usize, usize) {
        (idx % self.0, idx / self.0)
    }

    fn iter_coords<'a>(&'a self) -> impl Iterator<Item = (usize, usize)> {
        (0..self.1).zip(0..self.0)
    }
}

/// Z-order space-filling curve.
pub struct Morton {
    padded_size: usize,
}

impl Layout for Morton {
    fn new(width: usize, height: usize) -> Self {
        let longer_side = width.max(height);
        let padded_size = next_power_of_two(longer_side);
        Self { padded_size }
    }

    fn len(&self) -> usize {
        self.padded_size * self.padded_size
    }

    fn coord_to_index(&self, x: usize, y: usize) -> usize {
        morton_encode(x, y, self.padded_size)
    }

    fn index_to_coord(&self, idx: usize) -> (usize, usize) {
        morton_decode(idx, self.padded_size)
    }
}

/// Hilbert space-filling curve.
pub struct Hilbert {
    padded_size: usize,
}

impl Layout for Hilbert {
    fn new(width: usize, height: usize) -> Self {
        let longer_side = width.max(height);
        let padded_size = next_power_of_two(longer_side);
        Self { padded_size }
    }

    fn len(&self) -> usize {
        self.padded_size * self.padded_size
    }

    fn coord_to_index(&self, x: usize, y: usize) -> usize {
        hilbert_encode(x, y, self.padded_size)
    }

    fn index_to_coord(&self, idx: usize) -> (usize, usize) {
        hilbert_decode(idx, self.padded_size)
    }
}

/// Generic 2D grid.
#[derive(Debug, Clone)]
pub struct VecGrid<T, L: Layout = Hilbert> {
    width: usize,
    height: usize,
    layout: L,
    data: Vec<T>,
}

impl<T, L: Layout> VecGrid<T, L> {
    ///
    pub fn new(width: usize, height: usize, value: T) -> Self
    where
        T: Clone,
    {
        let layout = L::new(width, height);
        let data = vec![value; layout.len()];

        Self {
            width,
            height,
            layout,
            data,
        }
    }

    ///
    pub fn new_default(width: usize, height: usize) -> Self
    where
        T: Clone + Default,
    {
        let layout = L::new(width, height);
        let data = vec![T::default(); layout.len()];

        Self {
            width,
            height,
            layout,
            data,
        }
    }

    /// Returns the width of the grid.
    #[inline(always)]
    pub fn width(&self) -> usize {
        self.width
    }

    /// Returns the height of the grid.
    #[inline(always)]
    pub fn height(&self) -> usize {
        self.height
    }

    /// Returns the number of cells in the grid.
    #[inline(always)]
    pub fn len(&self) -> usize {
        self.width * self.height
    }

    /// Returns the value at `(x, y)`, or `None` if out of bounds.
    pub fn get(&self, x: usize, y: usize) -> Option<&T> {
        if x >= self.width || y >= self.height {
            return None;
        }
        Some(&self.data[self.layout.coord_to_index(x, y)])
    }

    /// Returns a mutable reference to the value at `(x, y)`, or `None` if out of bounds.
    pub fn get_mut<'a>(&'a mut self, x: usize, y: usize) -> Option<&'a mut T> {
        if x >= self.width || y >= self.height {
            return None;
        }
        let idx = self.layout.coord_to_index(x, y);
        self.data.get_mut(idx)
    }

    /// Iterates over 4-way neighbors.
    pub fn neighbors4(&self, x: usize, y: usize) -> impl Iterator<Item = (usize, usize)> + '_ {
        let w = self.width;
        let h = self.height;

        // inline array of candidates; invalid ones filtered out
        [
            (x.wrapping_sub(1), y, x > 0),
            (x + 1, y, x + 1 < w),
            (x, y.wrapping_sub(1), y > 0),
            (x, y + 1, y + 1 < h),
        ]
        .into_iter()
        .filter(|(_, _, is_valid)| *is_valid)
        .map(|(nx, ny, _)| (nx, ny))
    }

    /// Iterates over 4-way neighbors.
    pub fn neighbors4_wrap(&self, x: usize, y: usize) -> impl Iterator<Item = (usize, usize)> + '_ {
        let w = self.width;
        let h = self.height;

        // inline array of candidates
        [
            (x.wrapping_sub(1), y),
            ((x + 1) % w, y),
            (x, y.wrapping_sub(1)),
            (x, (y + 1) % h),
        ]
        .into_iter()
        .map(|(nx, ny)| (nx, ny))
    }

    /// Iterates over 8-way neighbors.
    pub fn neighbors8(&self, x: usize, y: usize) -> impl Iterator<Item = (usize, usize)> + '_ {
        let w = self.width;
        let h = self.height;

        // inline array of candidates
        [
            (x.wrapping_sub(1), y, x > 0),
            (x + 1, y, x + 1 < w),
            (x, y.wrapping_sub(1), y > 0),
            (x, y + 1, y + 1 < h),
            (x.wrapping_sub(1), y.wrapping_sub(1), x > 0 && y > 0),
            (x + 1, y + 1, x + 1 < w && y + 1 < h),
            (x + 1, y.wrapping_sub(1), x + 1 < w && y > 0),
            (x.saturating_sub(1), (y + 1) % h, x > 0 && y + 1 < h),
        ]
        .into_iter()
        .filter(|(_, _, is_valid)| *is_valid)
        .map(|(nx, ny, _)| (nx, ny))
    }

    /// Iterates over 8-way neighbors.
    pub fn neighbors8_wrap(&self, x: usize, y: usize) -> impl Iterator<Item = (usize, usize)> + '_ {
        let w = self.width;
        let h = self.height;

        // inline array of candidates
        [
            (x.wrapping_sub(1), y),
            ((x + 1) % w, y),
            (x, y.wrapping_sub(1)),
            (x, (y + 1) % h),
            (x.wrapping_sub(1), y.wrapping_sub(1)),
            ((x + 1) % w, (y + 1) % h),
            ((x + 1) % w, y.wrapping_sub(1)),
            (x.saturating_sub(1), (y + 1) % h),
        ]
        .into_iter()
        .map(|(nx, ny)| (nx, ny))
    }

    /// Iterates over all cells as `(x, y, &value)`.
    pub fn iter<'a>(&'a self) -> impl Iterator<Item = (usize, usize, &'a T)> {
        self.layout
            .iter_coords()
            .map(move |(x, y)| (x, y, &self.data[self.layout.coord_to_index(x, y)]))
    }
}

impl<T> Index<(usize, usize)> for VecGrid<T> {
    type Output = T;
    fn index(&self, (x, y): (usize, usize)) -> &T {
        self.get(x, y).expect("out of bounds")
    }
}

impl<T> IndexMut<(usize, usize)> for VecGrid<T> {
    fn index_mut(&mut self, (x, y): (usize, usize)) -> &mut T {
        self.get_mut(x, y).expect("out of bounds")
    }
}

fn next_power_of_two(mut x: usize) -> usize {
    if x == 0 {
        return 1;
    }
    x -= 1;
    x |= x >> 1;
    x |= x >> 2;
    x |= x >> 4;
    x |= x >> 8;
    x |= x >> 16;
    if std::mem::size_of::<usize>() >= 8 {
        x |= x >> 32;
    }
    x + 1
}

fn morton_encode(x: usize, y: usize, size: usize) -> usize {
    let bits = (size as f64).log2() as usize;
    let mut idx = 0;
    for i in 0..bits {
        let bx = (x >> i) & 1;
        let by = (y >> i) & 1;
        idx |= bx << (2 * i);
        idx |= by << (2 * i + 1);
    }
    idx
}

fn morton_decode(idx: usize, size: usize) -> (usize, usize) {
    let bits = (size as f64).log2() as usize;
    let mut x = 0;
    let mut y = 0;
    for i in 0..bits {
        x |= ((idx >> (2 * i)) & 1) << i;
        y |= ((idx >> (2 * i + 1)) & 1) << i;
    }
    (x, y)
}

fn hilbert_encode(mut x: usize, mut y: usize, n: usize) -> usize {
    let mut idx = 0;
    let mut s = n / 2;
    while s > 0 {
        let rx = ((x & s) > 0) as usize;
        let ry = ((y & s) > 0) as usize;
        idx += s * s * ((3 * rx) ^ ry);
        rotate_hilbert(&mut x, &mut y, s, rx, ry);
        s /= 2;
    }
    idx
}

fn hilbert_decode(mut idx: usize, n: usize) -> (usize, usize) {
    let mut x = 0;
    let mut y = 0;
    let mut s = 1;
    while s < n {
        let rx = 1 & (idx / 2);
        let ry = 1 & (idx ^ rx);
        rotate_hilbert(&mut x, &mut y, s, rx, ry);
        x += s * rx;
        y += s * ry;
        idx /= 4;
        s *= 2;
    }
    (x, y)
}

fn rotate_hilbert(x: &mut usize, y: &mut usize, n: usize, rx: usize, ry: usize) {
    if ry == 0 {
        if rx == 1 {
            *x = n - 1 - *x;
            *y = n - 1 - *y;
        }
        std::mem::swap(x, y);
    }
}
