use std::collections::{BTreeMap, HashMap};

/// 48% time, 82% memory (btreemap is slower)
///
/// Keep track of sorted order anagram in a map
/// while storing result index of the anagrams
///
/// Faster Solution:
/// ```ignore
/// let mut freq = vec![0; 26];
/// for &c in s.as_bytes() {
///     freq[(c - b'a') as usize] += 1;
/// }
/// ```
///
pub fn group_anagrams_btreemap(strs: Vec<String>) -> Vec<Vec<String>> {
    let mut bmap = BTreeMap::<Vec<char>, Vec<String>>::new();
    strs.into_iter().for_each(|s| {
        let mut anagram: Vec<char> = s.chars().collect();
        anagram.sort();

        // If the entry doesn't exist, it will be inserted as a new empty Vec.
        // `Entry.or_insert()` returns a mutable reference to the value in the entry
        let anagrams_vec: &mut Vec<String> = bmap.entry(anagram).or_insert(Vec::with_capacity(1));
        anagrams_vec.push(s); // Can chain .push(s) after .or_insert()
    });

    bmap.into_values().collect() // avoid clone via into_values()
}

/// 65% time, 82% memory (btreemap is slower)
pub fn group_anagrams_hashmap(strs: Vec<String>) -> Vec<Vec<String>> {
    let mut bmap = HashMap::<Vec<char>, Vec<String>>::new();
    strs.into_iter().for_each(|s| {
        let mut anagram: Vec<char> = s.chars().collect();
        anagram.sort();

        // If the entry doesn't exist, it will be inserted as a new empty Vec.
        // `Entry.or_insert()` returns a mutable reference to the value in the entry
        let anagrams_vec: &mut Vec<String> = bmap.entry(anagram).or_insert(Vec::with_capacity(1));
        anagrams_vec.push(s); // Can chain .push(s) after .or_insert()
    });

    bmap.into_values().collect() // avoid clone via into_values()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_group_anagrams_btreemap() {
        println!(
            "{:?}",
            group_anagrams_btreemap(
                ["eat", "tea", "tan", "ate", "nat", "bat"]
                    .iter()
                    .map(|s| s.to_string())
                    .collect()
            )
        );
        println!(
            "{:?}",
            group_anagrams_btreemap([""].iter().map(|s| s.to_string()).collect())
        );
        println!(
            "{:?}",
            group_anagrams_btreemap(["a"].iter().map(|s| s.to_string()).collect())
        );
    }

    #[test]
    fn test_group_anagrams_hashmap() {
        println!(
            "{:?}",
            group_anagrams_hashmap(
                ["eat", "tea", "tan", "ate", "nat", "bat"]
                    .iter()
                    .map(|s| s.to_string())
                    .collect()
            )
        );
        println!(
            "{:?}",
            group_anagrams_hashmap([""].iter().map(|s| s.to_string()).collect())
        );
        println!(
            "{:?}",
            group_anagrams_hashmap(["a"].iter().map(|s| s.to_string()).collect())
        );
    }
}
