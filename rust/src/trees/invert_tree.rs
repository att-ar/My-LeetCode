use std::cell::RefCell;
use std::rc::Rc;

use crate::utils::TreeNode;

/// 100% time, 14% memory
pub fn invert_tree(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
    match root {
        Some(node) => {
            // use Rc increment to avoid moving
            let left = invert_tree(node.borrow_mut().left.clone());
            let right = invert_tree(node.borrow_mut().right.clone());

            // swap children
            node.borrow_mut().left = right;
            node.borrow_mut().right = left;

            Some(node)
        }
        None => None,
    }
}

/// Claude's solution
pub fn invert_tree_claude(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
    if let Some(node) = &root {
        let mut borrowed = node.borrow_mut();
        let left = borrowed.left.take(); // moves node out of .left
        let right = borrowed.right.take(); // moves node out of .right

        borrowed.left = invert_tree(right);
        borrowed.right = invert_tree(left);
    }
    root
}

#[cfg(test)]
pub mod tests {
    use super::*;

    #[test]
    fn test_invert_tree() {
        let root = TreeNode::from_vec(vec![
            Some(4),
            Some(2),
            Some(7),
            Some(1),
            Some(3),
            Some(6),
            Some(9),
        ])
        .unwrap();

        println!("Original:\n{}", root.borrow());

        let inverted = super::invert_tree(Some(root.clone()));

        if let Some(inv) = inverted {
            println!("Inverted:\n{}", inv.borrow());
        } else {
            println!("Inverted: None");
        }
    }
}
