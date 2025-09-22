use std::cell::RefCell;
use std::rc::Rc;

use crate::utils::TreeNode;

/// 100% time, 14% memory
///
/// Easy ahhh question
/// just keep going down on the correct side while you are not within [p,q]
/// then return the first node that satisfies the bounds
///
/// Actually wasn't easy because of Rust borrow checker
pub fn lowest_common_ancestor(
    root: Option<Rc<RefCell<TreeNode>>>,
    p: Option<Rc<RefCell<TreeNode>>>,
    q: Option<Rc<RefCell<TreeNode>>>,
) -> Option<Rc<RefCell<TreeNode>>> {
    let (low, high) = {
        let mut bounds = [p?.borrow().val, q?.borrow().val];
        bounds.sort();
        (bounds[0], bounds[1])
    };

    // Guaranteed to have an answer since p and q exist in the BST
    let mut node = root?;
    loop {
        let next: Option<Rc<RefCell<TreeNode>>> = {
            let borrowed = node.borrow();
            let val = borrowed.val;
            if val < low {
                borrowed.right.clone() // clone because I can't move out of a borrow
            } else if val > high {
                borrowed.left.clone() // clone because I can't move out of a borrow
            } else {
                return Some(node.clone()); // clone because node can't be moved while borrowed
            }
            // drops borrow here so can reassign node outside
        };
        node = next?;
    }
}

#[cfg(test)]
pub mod tests {
    use super::*;

    #[test]
    fn test_lowest_common_ancestor() {
        let tree = TreeNode::from_vec(vec![
            Some(6),
            Some(2),
            Some(8),
            Some(0),
            Some(4),
            Some(7),
            Some(9),
            None,
            None,
            Some(3),
            Some(5),
        ]);
        let p = Some(Rc::new(RefCell::new(TreeNode::new(2))));
        let q = Some(Rc::new(RefCell::new(TreeNode::new(4))));
        println!("Tree:\n{}", tree.as_ref().unwrap().borrow());
        println!(
            "LCA:\n{}",
            lowest_common_ancestor(tree, p, q).unwrap().borrow()
        );
    }
}
