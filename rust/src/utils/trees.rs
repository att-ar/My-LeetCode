use std::{
    cell::RefCell,
    collections::VecDeque,
    fmt::{self, Display, Formatter},
    rc::Rc,
};

#[derive(Debug)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    /// Build a binary tree from a level-order vector where `None` represents missing nodes.
    /// Returns `None` if the vector is empty or the first element is `None`.
    pub fn from_vec(values: Vec<Option<i32>>) -> Option<Rc<RefCell<TreeNode>>> {
        if values.is_empty() {
            return None;
        }

        let mut iter = values.into_iter();
        let root_val = iter.next().unwrap()?;

        let root = Rc::new(RefCell::new(TreeNode {
            val: root_val,
            left: None,
            right: None,
        }));
        let mut queue: VecDeque<Rc<RefCell<TreeNode>>> = VecDeque::new();
        queue.push_back(Rc::clone(&root));

        while let Some(parent) = queue.pop_front() {
            if let Some(next) = iter.next() {
                if let Some(v) = next {
                    let left = Rc::new(RefCell::new(TreeNode {
                        val: v,
                        left: None,
                        right: None,
                    }));
                    parent.borrow_mut().left = Some(Rc::clone(&left));
                    queue.push_back(left);
                }
            } else {
                break;
            }

            if let Some(next) = iter.next() {
                if let Some(v) = next {
                    let right = Rc::new(RefCell::new(TreeNode {
                        val: v,
                        left: None,
                        right: None,
                    }));
                    parent.borrow_mut().right = Some(Rc::clone(&right));
                    queue.push_back(right);
                }
            } else {
                break;
            }
        }

        Some(root)
    }
}

/// A helper function for the Display implementation to format the tree recursively.
fn fmt_recursive(node: &TreeNode, f: &mut Formatter<'_>, indent: usize) -> fmt::Result {
    // Print the current node's value with indentation.
    writeln!(f, "{}{}", "  ".repeat(indent), node.val)?;

    // Recursively print the left child with increased indentation.
    if let Some(left_child) = &node.left {
        fmt_recursive(&left_child.borrow(), f, indent + 1)?;
    }

    // Recursively print the right child with increased indentation.
    if let Some(right_child) = &node.right {
        fmt_recursive(&right_child.borrow(), f, indent + 1)?;
    }

    Ok(())
}

/// Implements a multi-line, indented pre-order traversal string representation for the tree.
impl Display for TreeNode {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        // Start the recursive formatting from the root with an indent of 0.
        fmt_recursive(self, f, 0)
    }
}
