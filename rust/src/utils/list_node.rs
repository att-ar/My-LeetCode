#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

impl std::fmt::Display for Box<ListNode> {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let mut cur = Some(self);
        while let Some(node) = cur {
            write!(f, "ListNode({}) -> ", node.val)?;
            cur = node.next.as_ref()
        }
        writeln!(f, "None")?;

        Ok(())
    }
}

impl From<Vec<i32>> for Box<ListNode> {
    /// Return the first entry as the head of the list
    /// Very cool foldr
    fn from(list: Vec<i32>) -> Self {
        let opt = list
            .into_iter()
            .rev()
            .fold(None, |acc, val| Some(Box::new(ListNode { val, next: acc })));
        opt.unwrap()
    }
}
