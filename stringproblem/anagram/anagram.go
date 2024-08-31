package stringproblem

/*
["eat","tea","bat","tab","cat"] , [["ate","tea"],["bat","tab"],["cat"]]

*/

type Key [26]byte

func strKey(str string) (key Key) {
	for i := range str {
		key[str[i]-'a']++
	}
	return
}

func findAnagramInSlice(strs []string) [][]string {
	groups := make(map[Key][]string)

	for _, v := range strs {
		key := strKey(v)
		groups[key] = append(groups[key], v)
	}

	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		result = append(result, v)
	}
	return result
}

// func findAnagramInSlice(input []string) [][]string {

// 	mapWordAnagram := make(map[string][]string)
// 	var result [][]string

// 	for _, str := range input {

// 		chArr := strings.Split(str, "")

// 		sort.Strings(chArr)

// 		sortedString := strings.Join(chArr, "")

// 		mapWordAnagram[sortedString] = append(mapWordAnagram[sortedString], str)

// 	}

// 	for _, anagrams := range mapWordAnagram {

// 		if len(anagrams) > 1 {

// 			result = append(result, anagrams)
// 		}

// 	}

// 	return result

// }

func IsAnagramMap(str1 string, str2 string) bool {

	if len(str1) != len(str2) {
		return false
	}

	str1Map := make(map[rune]int, len(str1))

	for _, s := range str1 {

		str1Map[s-'a']++

	}

	for _, s := range str2 {

		str1Map[s-'a']--

	}

	for _, r := range str1Map {

		if r != 0 {
			return false
		}

	}

	return true
}
