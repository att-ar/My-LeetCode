package main

import "fmt"

func runTries() {
	fmt.Printf("\nTries\n------\n")

	fmt.Println("Sum of Prefix Scores")
	words := []string{"abc", "ab", "bc", "b"}
	fmt.Println(words)
	fmt.Println(sumPrefixScores(words))

	fmt.Println("Word Break II")
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
	fmt.Println(wordBreak("cbca", []string{"bc", "ca"}))
}

/*
336. Palindrome Pairs

# I can store the reverses of the words in the Tries too

Before inserting, iterate through the characters of the word in reverse and
When inserting the forward, I can iterate through the characters in reverse and
*/
func palindromePairs(words []string) [][]int {

}

/*
140. Word Break II
100% time, 8% memory

This is just backtracking Trie problem where the validity of the DFS is based on reaching the end of string `s`
Easy dubs tbh...

Use the TrieNodeWord struct in data_structs.go

Don't need to make a search feature because the search will happen in the wordBreak itself
*/
func wordBreak(s string, wordDict []string) []string {
	// build the Trie
	trie := NewTrieNodeWord()
	for _, word := range wordDict {
		trie.Insert(word)
	}

	// backtrack backed by the trie
	result := []string{}
	/* DFS each path possible */
	var backtrack func(i int, curS string)
	backtrack = func(i int, curS string) {
		if i == len(s) {
			// success base case
			result = append(result, curS[1:])
			return
		}

		/*
			use the trie starting from `i`
			this loop will exhaust all words that can be made starting from index `j`
		*/
		node := trie
		for j := i; j < len(s); j++ {
			if _, exists := node.children[s[j]]; exists {
				// can add `j`th letter
				node = node.children[s[j]]
				if node.word != "" {
					// DFS with the found word added to current path
					backtrack(j+1, curS+" "+node.word)
				}
			} else {
				// broke the word if no char is found so early exit
				return
			}
		}
	}

	backtrack(0, "")
	return result
}

/*
2416. Sum of Prefix Scores of Strings
54% time, 53% memory
You are given an array words of size n consisting of non-empty strings.

We define the score of a string term as the number of strings words[i] such that term is a prefix of words[i].

For example, if words = ["a", "ab", "abc", "cab"], then the score of "ab" is 2, since "ab" is a prefix of both "ab" and "abc".
Return an array answer of size n where answer[i] is the sum of scores of every non-empty prefix of words[i].

Note that a string is considered as a prefix of itself.
*/
func sumPrefixScores(words []string) []int {
	// will populate the Trie with every word O(n * k), k = max length of word
	// will increment the number stored at nodes where a word ends.
	trie := NewTrieNodeNum()
	// O(n * k)
	for _, word := range words {
		trie.Insert(word)
	}

	// O(n * k)
	res := make([]int, len(words))
	for i := range words {
		res[i] = trie.SearchPrefixes(words[i])
	}
	return res
}
