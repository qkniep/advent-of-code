use arrayvec::ArrayVec;

/// Represents a 2D grid of items.
#[derive(Clone, Debug)]
pub struct Grid<T: Clone> {
    pub width: usize,
    pub height: usize,
    /// Cells stored in row-major order.
    pub cells: Vec<T>,
}

impl<T: Copy> Grid<T> {
    /// Create a grid from a 2D vector.
    pub fn from_vec(vec: Vec<Vec<T>>) -> Self {
        let height = vec.len();
        let width = vec[0].len();
        let mut cells = Vec::with_capacity(width * height);
        for row in vec {
            assert_eq!(row.len(), width);
            cells.extend_from_slice(&row);
        }
        Grid {
            width,
            height,
            cells,
        }
    }

    /// Get cell at (x, y).
    pub fn get(&self, x: usize, y: usize) -> Option<T> {
        if x < self.width && y < self.height {
            Some(self.cells[y * self.width + x])
        } else {
            None
        }
    }

    /// Mutable access to cell at (x, y).
    pub fn get_mut(&mut self, x: usize, y: usize) -> Option<&mut T> {
        if x < self.width && y < self.height {
            Some(&mut self.cells[y * self.width + x])
        } else {
            None
        }
    }

    /// Iterator over all positions and values.
    pub fn iter(&self) -> impl Iterator<Item = ((usize, usize), &T)> {
        self.cells.iter().enumerate().map(move |(i, v)| {
            let y = i / self.width;
            let x = i % self.width;
            ((x, y), v)
        })
    }

    /// Return neighbors (4-directional).
    pub fn neighbors4(&self, x: usize, y: usize) -> ArrayVec<(usize, usize), 4> {
        let mut n = ArrayVec::new();
        if x > 0 {
            n.push((x - 1, y));
        }
        if x + 1 < self.width {
            n.push((x + 1, y));
        }
        if y > 0 {
            n.push((x, y - 1));
        }
        if y + 1 < self.height {
            n.push((x, y + 1));
        }
        n
    }

    /// Returns neighbors (8-directional).
    pub fn neighbors8(&self, x: usize, y: usize) -> ArrayVec<(usize, usize), 8> {
        let mut n = ArrayVec::new();
        for dy in [-1isize, 0, 1] {
            for dx in [-1isize, 0, 1] {
                if dx == 0 && dy == 0 {
                    continue;
                }
                let nx = x as isize + dx;
                let ny = y as isize + dy;
                if nx >= 0 && nx < self.width as isize && ny >= 0 && ny < self.height as isize {
                    n.push((nx as usize, ny as usize));
                }
            }
        }
        n
    }

    /// ASCII visualization.
    pub fn to_ascii_with<F>(&self, mut f: F) -> String
    where
        F: FnMut(&T) -> char,
    {
        let mut s = String::with_capacity(self.width * self.height + self.height);
        for y in 0..self.height {
            for x in 0..self.width {
                s.push(f(&self.cells[y * self.width + x]));
            }
            s.push('\n');
        }
        s
    }
}

impl<T: TryFrom<char> + Clone> TryFrom<&str> for Grid<T> {
    type Error = ();

    fn try_from(s: &str) -> Result<Self, Self::Error> {
        let height = s.lines().count();
        let width = s.lines().next().unwrap().len();
        let mut cells = Vec::with_capacity(width * height);
        for line in s.lines() {
            assert_eq!(line.len(), width);
            for c in line.chars() {
                cells.push(T::try_from(c).map_err(|_| ())?);
            }
        }
        Ok(Grid {
            width,
            height,
            cells,
        })
    }
}

impl<T: Into<char> + Copy> Grid<T> {
    /// ASCII visualization.
    pub fn to_ascii(&self) -> String {
        let mut s = String::with_capacity(self.width * self.height + self.height);
        for y in 0..self.height {
            for x in 0..self.width {
                s.push(self.cells[y * self.width + x].into());
            }
            s.push('\n');
        }
        s
    }
}
