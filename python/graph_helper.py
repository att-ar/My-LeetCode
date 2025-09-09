from typing import Optional, List
import matplotlib.pyplot as plt
import networkx as nx


class Node:
    def __init__(self, val=0, neighbors: Optional[List["Node"]] = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []
        # neighbors should be a list of Node

    def __repr__(self):
        return f"Node: {self.val}, Neighbors: {[node.val for node in self.neighbors]}"


def display_graph(node: Node):
    seen = set()
    stack = [node]
    while stack:
        for _ in range(len(stack)):
            cur = stack.pop(0)
            if cur not in seen:
                seen.add(cur)
                print(f"{cur}")
                stack.extend(cur.neighbors)


def create_graph(adjacency_list: list[list[int]]):
    # Create a dictionary to store nodes based on their values
    nodes_dict = {}

    # Helper function to get or create a node with a given value
    def get_or_create_node(value):
        if value not in nodes_dict:
            nodes_dict[value] = Node(value)
        return nodes_dict[value]

    # Iterate through the adjacency list and create nodes and edges
    for vertex, neighbors in enumerate(adjacency_list):
        current_node = get_or_create_node(vertex + 1)
        current_node.neighbors = [
            get_or_create_node(neighbor) for neighbor in neighbors
        ]

    # Return the first node in the dictionary (assuming the graph is connected)
    return nodes_dict[next(iter(nodes_dict))]


def draw_graph(graph):
    UG = nx.Graph()
    if isinstance(graph, list):
        for node, neighbors in enumerate(graph):
            for neighbor in neighbors:
                UG.add_edge(node, neighbor)
    else:
        for node, neighbors in graph.items():
            for neighbor in neighbors:
                UG.add_edge(node, neighbor)
    plt.figure(figsize=(8, 6))
    nx.draw(
        UG,
        with_labels=True,
        node_size=800,
        font_size=13,
        font_color="black",
        font_weight="bold",
        arrows=True,
    )
    plt.title("Graph Visualization for Undirected Graph")

    plt.show()
