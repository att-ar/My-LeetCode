class TreeNode(object):
    def __init__(
        self, val=0, left: "TreeNode | None" = None, right: "TreeNode | None" = None
    ):
        self.val = val
        self.left = left
        self.right = right

    def __repr__(self) -> str:
        return f"TreeNode: {self.val}"


def makeTree(nodes: list[int | None]) -> TreeNode | None:
    if not nodes:
        return None

    root = TreeNode(nodes[0])  # type:ignore
    queue = [root]
    i = 1

    while i < len(nodes):
        current_node = queue.pop(0)

        # Left child
        if i < len(nodes) and nodes[i] is not None:
            current_node.left = TreeNode(nodes[i])  # type:ignore
            queue.append(current_node.left)
        i += 1

        # Right child
        if i < len(nodes) and nodes[i] is not None:
            current_node.right = TreeNode(nodes[i])  # type:ignore
            queue.append(current_node.right)
        i += 1

    return root


def print_tree(root: TreeNode, level=0, prefix="Root: "):
    if root is not None:
        print(" " * (level * 4) + prefix + str(root.val))
        if root.left is not None or root.right is not None:
            print_tree(root.left, level + 1, "L--- ")  # type:ignore
            print_tree(root.right, level + 1, "R--- ")  # type:ignore
