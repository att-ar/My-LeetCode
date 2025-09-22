use std::collections::HashMap;

pub struct FolderTrie {
    pub children: HashMap<String, FolderTrie>,
    pub is_end: bool,
}

impl FolderTrie {
    pub fn new() -> Self {
        Self {
            children: HashMap::new(),
            is_end: false,
        }
    }
    /// Short circuits if a parent folder exists
    /// Returns bool on success
    ///     Time optimization to avoid doing a DFS to get the parents
    ///     NOTE: Only works if you insert in order of increasing length
    pub fn insert(&mut self, key: &str) -> bool {
        // can .map(|s| s.to_string()) but that would call it more times than having it in the for loop
        let parts: Vec<&str> = key.split('/').filter(|&s| !s.is_empty()).collect();

        let mut cur = self;
        for part in parts.into_iter() {
            cur = cur
                .children
                .entry(part.to_string())
                .or_insert_with(Self::new);

            if cur.is_end {
                // short-circuit
                return false;
            }
        }
        cur.is_end = true;
        true
    }
}

/// 40% time, 40% memory
pub fn remove_subfolders(folder: Vec<String>) -> Vec<String> {
    // Sort by length since insert short circuits
    let mut folder = folder;
    folder.sort_by_key(|folder_path| folder_path.len());

    let mut trie = FolderTrie::new();
    folder
        .into_iter()
        .filter(|folder_path| trie.insert(folder_path))
        .collect()
}

#[cfg(test)]
pub mod tests {
    use super::*;

    #[test]
    fn test_remove_subfolders() {
        let folder: Vec<String> = ["/a", "/a/b", "/c/d", "/c/d/e", "/c/f"]
            .into_iter()
            .map(|s| s.to_string())
            .collect();
        println!("{:?}", remove_subfolders(folder));

        let folder: Vec<String> = ["/a", "/a/b/c", "/a/b/d"]
            .into_iter()
            .map(|s| s.to_string())
            .collect();
        println!("{:?}", remove_subfolders(folder));

        let folder: Vec<String> = ["/a/b/c", "/a/b/ca", "/a/b/d"]
            .into_iter()
            .map(|s| s.to_string())
            .collect();
        println!("{:?}", remove_subfolders(folder));
    }
}
