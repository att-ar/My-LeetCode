use std::collections::HashMap;

/// Disjoint Set used in Union-Find algos
/// I am choosing to do union-by-size since it tends to make more sense
/// for leetcode problems.
pub struct DisjointSet {
    parents: HashMap<i32, i32>,
    sizes: HashMap<i32, i32>,
}

impl DisjointSet {
    pub fn new() -> Self {
        DisjointSet {
            parents: HashMap::new(),
            sizes: HashMap::new(),
        }
    }
    pub fn with_capacity(capacity: usize) -> Self {
        DisjointSet {
            parents: HashMap::with_capacity(capacity),
            sizes: HashMap::with_capacity(capacity),
        }
    }

    /// check if a key has been inserted into the set
    /// We only check one since self.insert() auto mutates both
    pub fn contains(&self, n: i32) -> bool {
        self.parents.contains_key(&n)
    }

    /// Add a new element to the data structure as its own set.
    pub fn insert(&mut self, n: i32) {
        self.parents.insert(n, n);
        self.sizes.insert(n, 1); // Initial size is 1 (itself)
    }

    /// Union-by-size the sets containing n1 and n2
    pub fn union(&mut self, n1: i32, n2: i32) {
        // Get roots
        let root1 = self.find(n1);
        let root2 = self.find(n2);

        // No-op if n1 and n2 are in the same set
        if root1 == root2 {
            return;
        }

        // Get sizes of roots
        let size1 = self.sizes[&root1];
        let size2 = self.sizes[&root2];

        // Union roots
        if size1 > size2 {
            self.parents.insert(root2, root1);
            self.sizes.insert(root1, size1 + size2);
        } else {
            self.parents.insert(root1, root2);
            self.sizes.insert(root2, size1 + size2);
        }
    }

    /// Find the root of the set containing `n` with path compression.
    fn find(&mut self, n: i32) -> i32 {
        let mut root = self.parents[&n];
        if root != n {
            root = self.find(root); // Safe unwrap
            self.parents.insert(n, root); // Path compression
        }
        root
    }

    pub fn get_largest_size(&self) -> i32 {
        self.sizes.values().max().cloned().unwrap_or_default()
    }
}

impl Default for DisjointSet {
    fn default() -> Self {
        Self::new()
    }
}
