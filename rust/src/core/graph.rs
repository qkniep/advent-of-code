//! Simple adjacency list graph implementation.
//!
//!

use std::collections::VecDeque;

/// Generic adjacency list graph
pub struct Graph {
    pub edges: Vec<Vec<usize>>, // Node index -> neighbors
}

impl Graph {
    /// Creates a new graph with `n` nodes and no edges.
    pub fn new(n: usize) -> Self {
        Graph {
            edges: vec![Vec::new(); n],
        }
    }

    /// Adds an edge from `u` to `v`.
    pub fn add_edge(&mut self, u: usize, v: usize) {
        self.edges[u].push(v);
    }

    /// Adds an undirected edge between `u` and `v`.
    pub fn add_undirected_edge(&mut self, u: usize, v: usize) {
        self.edges[u].push(v);
        self.edges[v].push(u);
    }

    /// BFS returning distances from start node.
    pub fn bfs(&self, start: usize) -> Vec<Option<usize>> {
        let mut dist = vec![None; self.edges.len()];
        let mut queue = VecDeque::new();
        dist[start] = Some(0);
        queue.push_back(start);

        while let Some(u) = queue.pop_front() {
            let d = dist[u].unwrap();
            for &v in &self.edges[u] {
                if dist[v].is_none() {
                    dist[v] = Some(d + 1);
                    queue.push_back(v);
                }
            }
        }

        dist
    }

    /// DFS traversal (recursive).
    pub fn dfs<F>(&self, start: usize, mut visit: F)
    where
        F: FnMut(usize),
    {
        let mut visited = vec![false; self.edges.len()];
        fn dfs_inner<F>(graph: &Graph, u: usize, visited: &mut [bool], visit: &mut F)
        where
            F: FnMut(usize),
        {
            visited[u] = true;
            visit(u);
            for &v in &graph.edges[u] {
                if !visited[v] {
                    dfs_inner(graph, v, visited, visit);
                }
            }
        }
        dfs_inner(self, start, &mut visited, &mut visit);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn basic() {}
}
