package hw_1

func LongestCommonPrefix(strs []string) string {

	if len(strs) == 1 {
		return strs[0]
	}

	var minRange int
	firstWord := strs[0]
	res := ""

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(firstWord) {
			minRange = len(strs[i])
		} else {
			minRange = len(firstWord)
		}

		if res != "" && len(res) < minRange {
			minRange = len(res)
		}

		curRes := ""
		for j := 0; j < minRange; j++ {
			if firstWord[j] == strs[i][j] {
				curRes = string(append([]byte(curRes), strs[i][j]))
			} else if curRes == "" {
				return ""
			} else {
				break
			}
		}
		res = curRes
	}
	return res
}
