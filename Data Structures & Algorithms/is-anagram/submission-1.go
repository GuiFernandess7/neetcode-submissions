import "slices"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var a []int32
	var b []int32 

	for _, char := range s {
		a = append(a, rune(char))
	}

	for _, char := range t {
		b = append(b, rune(char))
	}

	slices.Sort(a)
	slices.Sort(b)
	return slices.Equal(a, b)
}
