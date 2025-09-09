package main

// UNION-FIND algo functions

/*
@parents should be an slice of 2-digit arrays where
  - the index is the node
  - first digit is the parent of the node
  - second digit is the degree of the parent (this isn't necessary accurate as only the connected group's parent gets its degree updated on `union` calls)
*/
func find(parents [][2]int, a int) [2]int {
	var parent [2]int
	parent = parents[a]

	for parents[parent[0]][0] != parent[0] {
		// this reduces the number of searches needed later
		parents[a] = parents[parent[0]]
		parent = parents[a]
	}
	// when you reach value whose parent is itself then you have finished the `find`
	return parent
}

/*
@parents should be an slice of 2-digit arrays where
  - the index is the node
  - first digit is the parent of the node
  - second digit is the degree of the parent (this isn't necessary accurate as only the connected group's parent gets its degree updated on `union` calls)
*/
func union(parents [][2]int, a, b int) bool {
	var parentA, parentB [2]int
	parentA = find(parents, a)
	parentB = find(parents, b)
	// NOTE: parents[parentA[0]][0] is equal to parentA[0] since the `find` function returns a node whose parent is itself

	// check if the parents are the same, in which case return false because the union is unsuccessful (already happened)
	if parentA[0] == parentB[0] {
		return false
	}

	// union the two parents by adding the samller degree one to the larger degree one
	// and set the parent of the smaller one to the larger one
	if parentA[1] >= parentB[1] {
		parents[parentA[0]][1] += parents[parentB[0]][1]
		parents[parentB[0]] = parentA
	} else {
		parents[parentB[0]][1] += parents[parentA[0]][1]
		parents[parentA[0]] = parentB
	}
	return true
}
