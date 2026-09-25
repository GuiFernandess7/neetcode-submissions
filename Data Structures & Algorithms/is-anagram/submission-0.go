func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashS := make(map[string]int)
	for _, char := range s {
		if _, ok := hashS[string(char)]; ok {
			hashS[string(char)] ++
		} else {
			hashS[string(char)] = 1
		}
	}

	hashT := make(map[string]int)
	for _, char := range t {
		if _, ok := hashT[string(char)]; ok {
			hashT[string(char)] ++
		} else {
			hashT[string(char)] = 1
		}
	}

	for _, char := range t {
		if hashS[string(char)] != hashT[string(char)] {
			return false
		}
	}
	return true
}
