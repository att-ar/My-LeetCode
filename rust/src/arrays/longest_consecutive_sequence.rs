use std::collections::HashSet;

use crate::utils::DisjointSet;

/// I checked my python solution, because I didn't feel like figuring out O(n) version
/// The fold is a lot slower than a naive loop interestingly. probably compiler optimizations im missing out on.
///
/// I also thought .filter would be faster than branching with continue but apparently not.
pub fn longest_consecutive(nums: Vec<i32>) -> i32 {
    let set = nums.into_iter().collect::<HashSet<i32>>();
    let mut res = 0;
    for &n in set.iter().filter(|&&n| !set.contains(&(n - 1))) {
        let mut interim = 1;
        let mut m = n + 1;
        while set.contains(&m) {
            interim += 1;
            m += 1;
        }
        res = res.max(interim);
    }
    res
}

/// 5% time, 6% memory
/// Significantly slower than the HashSet method.
///
/// Redoing the problem as intended with Union Find for practice since
/// I will have to do Graph problems later anyways.
pub fn longest_consecutive_union_find(nums: Vec<i32>) -> i32 {
    let mut set = DisjointSet::with_capacity(nums.len());

    // Go through nums once and Union-Find : O(n)
    for n in nums {
        // Insert if new node
        if !set.contains(n) {
            set.insert(n);
        }

        // Union with predecessor
        if set.contains(n - 1) {
            set.union(n, n - 1);
        }

        // Union with successor
        if set.contains(n + 1) {
            set.union(n, n + 1);
        }
    }

    // Return largest set size found
    set.get_largest_size()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_longest_consecutive() {
        println!("{:?}", longest_consecutive(vec![100, 4, 200, 1, 3, 2]));
        println!(
            "{:?}",
            longest_consecutive(vec![0, 3, 7, 2, 5, 8, 4, 6, 0, 1])
        );
        println!("{:?}", longest_consecutive(vec![1, 0, 1, 2]));
        println!(
            "{:?}",
            longest_consecutive(vec![
                121, 1, 0, 1, 2, 120, 112, 113, 119, 118, 115, 116, 111, 117, 114
            ])
        )
    }

    #[test]
    fn test_longest_consecutive_union_find() {
        println!(
            "{:?}",
            longest_consecutive_union_find(vec![100, 4, 200, 1, 3, 2])
        );
        println!(
            "{:?}",
            longest_consecutive_union_find(vec![0, 3, 7, 2, 5, 8, 4, 6, 0, 1])
        );
        println!("{:?}", longest_consecutive_union_find(vec![1, 0, 1, 2]));
        println!(
            "{:?}",
            longest_consecutive_union_find(vec![
                121, 1, 0, 1, 2, 120, 112, 113, 119, 118, 115, 116, 111, 117, 114
            ])
        );
    }
}
