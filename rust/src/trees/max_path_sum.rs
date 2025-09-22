use crate::utils::TreeNode;
use std::{cell::RefCell, rc::Rc};

/// 17% time, 65% memory
///
/// Gonna have to use recursion no cap
/// I don't remember my python solution but I am pretty sure it is straight-forward
///
/// Values at node:
/// - Max path from left child
///     - Excluding node
///     - Including node (iff includes left child)
/// - Max path from right child
///     - Excluding node
///     - Including node (iff includes right child)
/// - Sum of max path from children and node (iff children max include children)
/// - Value of node itself
///
/// Since I should always have a path including the current node to allow for growth up the tree,
/// I'll have two values returned.
/// - Max path at node
/// - Path at node including node with the larger of the two inclusionary paths from the child
///
/// Or, I can reset both accumulators to just be the current node
///     this is useful if I all of a sudden get a positive node while carrying negative sums
pub fn max_path_sum(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    /// Returns (max path at node, max path at node allowing connection to parent)
    fn dfs(node: Option<Rc<RefCell<TreeNode>>>) -> Option<(i32, i32)> {
        match node {
            None => {
                // println!("None");
                None
            }
            Some(node) => {
                let borrowed = node.borrow();
                // Consider node alone:
                let val = borrowed.val;
                let mut path = val;
                let mut path_with = val;

                // Get left child
                let left_with = match dfs(borrowed.left.clone()) {
                    Some((left, left_with)) => {
                        path = path.max(left); // larger node val
                        // println!("Left|{left}, {left_with}| ");
                        left_with
                    }
                    None => 0, // 0 okay even tho we have neg nums since it is in a sum
                };

                // Get right child
                let right_with = match dfs(borrowed.right.clone()) {
                    Some((right, right_with)) => {
                        path = path.max(right); // larger node val
                        // println!("Right|{right}, {right_with}| ");
                        right_with
                    }
                    None => 0, // 0 okay even tho we have neg nums since it is in a sum
                };

                // `path` is either the largest node or the largest sum in a subtree's path
                // Basically, `path` holds the max disconnected from parent node.
                path = path.max(val + left_with + right_with);
                // `path_with` holds the max going through node and allows connection to parent node.
                // Thus, cannot include both children.
                path_with += left_with.max(right_with).max(0);

                // println!("Node|{path}, {path_with}, {val}|");
                Some((path.max(path_with), path_with))
            }
        }
    }

    // return path since it is >= path_with
    dfs(root).unwrap().0
}

/// 100% time, 65% memory
///
/// Trying to make a simple and more efficient version without Rc cloning
/// and without Option because that adds a lot of overhead and kills runtime.
pub fn max_path_sum_clean(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    /// Returns (max path at node, max path at node allowing connection to parent)
    ///
    /// Note: .as_deref() is safe here because we only hit each node once so Rc isn't needed
    fn dfs(node: Option<&RefCell<TreeNode>>) -> (i32, i32) {
        match node {
            None => {
                (i32::MIN, 0) // Need 0 because left_with and right_with are in sums
            }
            Some(node) => {
                let borrowed = node.borrow();
                let val = borrowed.val;
                let (left, left_with) = dfs(borrowed.left.as_deref());
                let (right, right_with) = dfs(borrowed.right.as_deref());

                // `path` holds the largest sum disconnected from parent node.
                // i.e. largest sum in subtree (can be individual node)
                let path = val.max(left).max(right).max(val + left_with + right_with);
                // `path_with` holds the max going through node and allows connection to parent node.
                // Thus, cannot include both children.
                let path_with = val + left_with.max(right_with).max(0);

                // path should always be >= path_with because path_with will not persist
                // up the tree, so path will need to store it
                (path.max(path_with), path_with)
            }
        }
    }

    dfs(root.as_deref()).0
}

#[cfg(test)]
pub mod tests {
    use super::*;

    #[test]
    fn test_max_path_sum() {
        // root max
        let tree = TreeNode::from_vec(vec![
            Some(5),
            Some(4),
            Some(8),
            Some(11),
            None,
            Some(13),
            Some(4),
            Some(7),
            Some(2),
            None,
            None,
            None,
            Some(1),
        ]);
        println!("{}", tree.as_ref().unwrap().borrow());
        println!("{}", max_path_sum(tree));

        // subtree max
        let tree = TreeNode::from_vec(vec![
            Some(-10),
            Some(9),
            Some(20),
            None,
            None,
            Some(15),
            Some(7),
        ]);
        println!("\n{}", tree.as_ref().unwrap().borrow());
        println!("{}", max_path_sum(tree));
    }

    #[test]
    fn test_max_path_sum_clean() {
        // root max
        let tree = TreeNode::from_vec(vec![
            Some(5),
            Some(4),
            Some(8),
            Some(11),
            None,
            Some(13),
            Some(4),
            Some(7),
            Some(2),
            None,
            None,
            None,
            Some(1),
        ]);
        println!("{}", tree.as_ref().unwrap().borrow());
        println!("{}", max_path_sum_clean(tree));

        // subtree max
        let tree = TreeNode::from_vec(vec![
            Some(-10),
            Some(9),
            Some(20),
            None,
            None,
            Some(15),
            Some(7),
        ]);
        println!("\n{}", tree.as_ref().unwrap().borrow());
        println!("{}", max_path_sum_clean(tree));
    }
}
