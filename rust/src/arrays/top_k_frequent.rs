use std::collections::{BTreeMap, BinaryHeap, HashMap};

/// 100% time, 32% memory
/// O(nlogn)
/// 1 <= k <= unique values <= n
/// Counter into Heap over the counts
pub fn top_k_frequent_one_liner(nums: Vec<i32>, k: i32) -> Vec<i32> {
    // O(nlogn); logn because of BTreeMap insert
    let counter: BTreeMap<i32, i32> = nums.into_iter().fold(BTreeMap::new(), |mut acc, n| {
        // acc.entry(n).and_modify(|v| *v += 1).or_insert(1);
        // idiomatic:
        *acc.entry(n).or_default() += 1;
        acc
    });

    // Build max heap using counter values
    // O(n + klogn); Heap build is O(n)
    let mut max_heap =
        BinaryHeap::<(i32, i32)>::from_iter(counter.into_iter().map(|(k, v)| (v, k)));
    let mut result = Vec::<i32>::with_capacity(k as usize);
    for _ in 0..k {
        result.push(max_heap.pop().unwrap().1);
    }
    result
}

/// 17% time, 47% memory
/// O(nlogk)
pub fn top_k_frequent_efficient(nums: Vec<i32>, k: i32) -> Vec<i32> {
    // O(n)
    let counter: HashMap<i32, i32> = nums.into_iter().fold(HashMap::new(), |mut acc, n| {
        *acc.entry(n).or_default() -= 1; // decrement because I need a min heap
        acc
    });

    // I think: O(nlogk)
    // O(k): Populate with k (0,0)s to avoid any branching (ignore Option branch and ignore length k branch)
    let mut min_heap = BinaryHeap::<(i32, i32)>::from(vec![(0, 0); k as usize]);
    // O(nlogk)
    for (num, freq) in counter.into_iter() {
        if freq < min_heap.peek().unwrap().0 {
            min_heap.pop();
            min_heap.push((freq, num));
        }
    }

    // O(k)
    min_heap.into_iter().map(|tup| tup.1).collect()
}

pub fn top_k_frequent_functional(nums: Vec<i32>, k: i32) -> Vec<i32> {
    // O(nlogn) logn because of BTreeMap
    let counter: BTreeMap<i32, i32> = nums.into_iter().fold(BTreeMap::new(), |mut acc, n| {
        // acc.entry(n).and_modify(|v| *v += 1).or_insert(1);
        // idiomatic:
        *acc.entry(n).or_default() += 1;
        acc
    });

    // Functional programming: slower since it is nlogn (the sorted vec)
    BinaryHeap::<(i32, i32)>::from_iter(counter.into_iter().map(|(k, v)| (v, k)))
        .into_sorted_vec()
        .into_iter()
        .take(k as usize)
        .map(|e| e.0)
        .collect()
}

#[cfg(test)]
pub mod tests {
    use super::*;

    #[test]
    fn test_top_k_frequent_one_liner() {
        println!("{:?}", top_k_frequent_one_liner(vec![1, 1, 1, 2, 2, 3], 2));
    }

    #[test]
    fn test_top_k_frequent_efficient() {
        println!("{:?}", top_k_frequent_efficient(vec![1, 1, 1, 2, 2, 3], 2));
    }

    #[test]
    fn test_top_k_frequent_functional() {
        println!("{:?}", top_k_frequent_functional(vec![1, 1, 1, 2, 2, 3], 2));
    }
}
