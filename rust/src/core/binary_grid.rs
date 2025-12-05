//! Ultra-compact 2D binary grid.
//!
//! Stores one bit per cell and provides fast access and iteration.
//! This implementation uses [`BitVec`] from the [`bitvec`] crate.

use bitvec::prelude::*;

/// A compact 2D binary grid based on [`BitVec`].
#[derive(Debug, Clone)]
pub struct BinaryGrid {
    width: usize,
    height: usize,
    bits: BitVec,
}

impl BinaryGrid {
    /// Creates a new grid with all cells set to `value` (true/false).
    pub fn new(width: usize, height: usize, value: bool) -> Self {
        let len = width.checked_mul(height).expect("overflow");
        let mut bits = BitVec::with_capacity(len);
        bits.fill(value);
        Self {
            width,
            height,
            bits,
        }
    }

    /// Parses a string into a grid.
    ///
    /// - `false_char` is the character used to represent `false`.
    /// - `true_char` is the character used to represent `true`.
    pub fn from_utf8_str(input: &str, false_char: char, true_char: char) -> Self {
        let width = input.lines().next().unwrap().len();
        let height = input.lines().count();
        let mut grid = Self::new(width, height, false);
        for (y, line) in input.lines().enumerate() {
            for (x, c) in line.chars().enumerate() {
                if c == true_char {
                    grid.set(x, y, true);
                } else {
                    assert!(c == false_char);
                }
            }
        }
        grid
    }

    /// Parses a string into a grid.
    ///
    /// - `false_char` is the byte used to represent `false`.
    /// - `true_char` is the byte used to represent `true`.
    pub fn from_ascii_str(input: &str, false_char: u8, true_char: u8) -> Self {
        let width = input.lines().next().unwrap().len();
        let height = input.lines().count();
        let mut grid = Self::new(width, height, false);
        for (y, line) in input.lines().enumerate() {
            for (x, c) in line.bytes().enumerate() {
                if c == true_char {
                    grid.set(x, y, true);
                } else {
                    assert!(c == false_char);
                }
            }
        }
        grid
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

    /// Converts `(x, y)` to bit index.
    #[inline(always)]
    fn index(&self, x: usize, y: usize) -> usize {
        y * self.width + x
    }

    /// Returns the value at `(x, y)`.
    #[inline(always)]
    pub fn get(&self, x: usize, y: usize) -> bool {
        self.bits[self.index(x, y)]
    }

    /// Sets the value at `(x, y)`.
    #[inline(always)]
    pub fn set(&mut self, x: usize, y: usize, value: bool) {
        let idx = self.index(x, y);
        self.bits.set(idx, value);
    }

    /// Iterates over all positions returning `(x, y)`.
    pub fn iter_coords<'a>(&'a self) -> impl Iterator<Item = (usize, usize)> + 'a {
        (0..self.height).zip(0..self.width)
    }

    /// Iterates over all cells returning `(x, y, value)`.
    pub fn iter<'a>(&'a self) -> impl Iterator<Item = (usize, usize, bool)> + 'a {
        self.iter_coords().map(|(x, y)| (x, y, self.get(x, y)))
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

    /// Counts the number of cells that are `true`.
    pub fn count_ones(&self) -> usize {
        self.bits.count_ones()
    }

    /// Counts the number of cells that are `false`.
    pub fn count_zeros(&self) -> usize {
        self.bits.len() - self.bits.count_ones()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn basic() {
        let mut g = BinaryGrid::new(4, 4, false);
        g.set(1, 1, true);
        g.set(2, 3, true);
        assert!(g.get(1, 1));
        assert!(!g.get(0, 0));
        let ones: Vec<_> = g.iter().filter(|(_, _, v)| *v).collect();
        assert_eq!(ones.len(), 2);
        let n4: Vec<_> = g
            .neighbors4(1, 1)
            .filter(|(nx, ny)| g.get(*nx, *ny))
            .collect();
        assert_eq!(n4.len(), 0);
    }

    #[test]
    fn parsing() {
        let g = BinaryGrid::from_ascii_str("##\n.#\n##", b'.', b'#');
        assert!(g.get(0, 0));
        assert!(g.get(1, 0));
        assert!(!g.get(0, 1));
        assert!(g.get(1, 1));
        assert!(g.get(0, 2));
        assert!(g.get(1, 2));
    }
}
