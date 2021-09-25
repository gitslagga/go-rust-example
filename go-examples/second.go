package main

import (
	"fmt"
)

func main() {
	// You are given a string s and an array of strings words of the same length. Return all starting indices of substring(s) in s that is a concatenation of each word in words exactly once, in any order, and without any intervening characters.
	// You can return the answer in any order.

	// Example 1:

	// Input: s = "barfoothefoobarman", words = ["foo","bar"]
	// Output: [0,9]
	// Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar" respectively.
	// The output order does not matter, returning [9,0] is fine too.
	// Example 2:

	// Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
	// Output: []
	// Example 3:

	// Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
	// Output: [6,9,12]

	// Constraints:

	// 1 <= s.length <= 104
	// s consists of lower-case English letters.
	// 1 <= words.length <= 5000
	// 1 <= words[i].length <= 30
	// words[i] consists of lower-case English letters.

	StringArrayConvert("barfoothefoobarman", []string{"foo", "bar"})
	StringArrayConvert("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"})
	StringArrayConvert("barfoofoobarthefoobarman", []string{"bar", "foo", "the"})
}

func StringArrayConvert(str string, worlds []string) (indexs []int) {
	lenStr := len(str)
	lenWords := len(worlds)
	length := len(worlds[0])

	if 1 <= lenStr && lenStr <= 104 && 1 <= lenWords && lenWords <= 5000 {
		for _, v := range worlds {
			if 1 > len(v) && len(v) > 30 {
				return
			}
		}

		tempWords := map[string]int{}
		for _, v := range worlds {
			tempWords[v]++
		}

		for index := 0; index+lenWords*length < lenStr; index++ {
			tempStr := map[string]int{}

			var j int
			for j = 0; j < lenWords; j++ {
				temp := str[index+j*length : index+(j+1)*length]
				tempStr[temp]++

				if tempStr[temp] > tempWords[temp] {
					break
				}
			}

			if j == lenWords {
				indexs = append(indexs, index)
			}
		}
	}

	fmt.Println(indexs)
	return
}
