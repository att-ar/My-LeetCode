use crate::utils::ListNode;
use std::cmp::Ordering;
use std::collections::BinaryHeap;

impl PartialOrd for ListNode {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Ord for ListNode {
    /// Using negatives to get min heap
    fn cmp(&self, other: &Self) -> Ordering {
        (-self.val).cmp(&(-other.val))
    }
}

/// 100% time, 60% memory
///
/// Basically just put all active heads in a heap and then insert into list
///
/// Since ListNode doesn't derive PartialOrd, and Ord I will impl myself
pub fn merge_k_lists(lists: Vec<Option<Box<ListNode>>>) -> Option<Box<ListNode>> {
    let mut head = Box::new(ListNode::new(0));
    let mut cur = head.as_mut();

    // naming it min_heap because I am using negatives
    let mut min_heap: BinaryHeap<Box<ListNode>> = lists.into_iter().flatten().collect();

    while !min_heap.is_empty() {
        let mut node = min_heap.pop().unwrap();

        // move node.next out of node and into heap
        if let Some(next) = node.next.take() {
            // take needed to avoid partial move
            min_heap.push(next);
        };

        // Can freely move node now
        cur.next = Some(node);
        // need deref so I don't move the Option from cur.next but I get a &mut Option<T>
        cur = cur.next.as_deref_mut().unwrap();
    }

    head.next
}

#[cfg(test)]
pub mod tests {
    use super::*;

    #[test]
    fn test_merge_k_lists() {
        let lists: Vec<Vec<i32>> = vec![vec![1, 4, 5], vec![1, 3, 4], vec![2, 6]];
        let lists: Vec<Option<Box<ListNode>>> = lists
            .into_iter()
            .map(|v| Some(Box::<ListNode>::from(v)))
            .collect();
        println!("{}", merge_k_lists(lists).unwrap());

        let lists: Vec<Vec<i32>> = vec![vec![1, 4, 5], vec![1, 3, 4], vec![], vec![]];
        let lists: Vec<Option<Box<ListNode>>> = lists
            .into_iter()
            .map(|v| {
                if v.is_empty() {
                    None
                } else {
                    Some(Box::<ListNode>::from(v))
                }
            })
            .collect();
        println!("{lists:?}");
        println!("{}", merge_k_lists(lists).unwrap());
    }
}
