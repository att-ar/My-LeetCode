use crate::utils::ListNode;

/// 100% time, 83% memory
/// doubled memory % by swapping `is_some()` for let `Some()`
pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut cur = head;
    let mut prev: Option<Box<ListNode>> = None;

    while let Some(mut node) = cur {
        // move next into tmp
        let tmp = node.next;
        // set next to prev
        node.next = prev;
        // set prev to current
        prev = Some(node);
        // set current to tmp (move to next node)
        cur = tmp;
    }

    prev
}

#[cfg(test)]
pub mod tests {
    use super::*;

    #[test]
    fn test_reverse_list() {
        let nums: Vec<i32> = vec![1, 2, 3, 4, 5];
        let head = Box::<ListNode>::from(nums);
        println!(
            "Before: {}\n\tAfter {}",
            head.clone(),
            reverse_list(Some(head)).expect("Expected a ListNode")
        );

        let nums: Vec<i32> = vec![1, 2, 3, 4, 5, 10, 11, 7, 8, 9, 100, 101, 102, 2003, 200];
        let head = Box::<ListNode>::from(nums);
        println!(
            "Before: {}\n\tAfter {}",
            head.clone(),
            reverse_list(Some(head)).expect("Expected a ListNode")
        );
    }
}
