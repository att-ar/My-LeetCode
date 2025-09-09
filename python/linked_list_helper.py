from typing import Optional


# Definition for singly-linked list.
class ListNode:
    def __init__(self, val, next=None):
        self.val = val
        self.next = next

    def __repr__(self) -> str:
        return f"{self.val}" + (
            f" -> {self.next.__repr__()}" if self.next is not None else ""
        )


def makeLinkedList(lst: list[int]) -> Optional[ListNode]:
    if lst:
        root = cur = ListNode(lst[0])
        for val in lst[1:]:
            cur.next = ListNode(val=val)
            cur = cur.next
        return root
    return None


def makeListOfLinkedLists(lst: list[list[int]]) -> list[Optional[ListNode]]:
    return [makeLinkedList(l) for l in lst]  # noqa: E741
