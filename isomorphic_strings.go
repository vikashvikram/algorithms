package algorithms

func IsIsomorphic(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	one_to_two_dict := make(map[byte]byte)
	two_to_one_dict := make(map[byte]byte)
	for index := range str1 {
		char1 := str1[index]
		char2 := str2[index]
		if _, ok1 := one_to_two_dict[char1]; !ok1 {
			if _, ok2 := two_to_one_dict[char2]; ok2 {
				return false
			}
			one_to_two_dict[char1] = char2
			two_to_one_dict[char2] = char1
		} else {
			if _, ok3 := two_to_one_dict[char2]; !ok3 {
				return false
			}
			if one_to_two_dict[char1] != char2 || two_to_one_dict[char2] != char1 {
				return false
			}
		}
	}
	return true
}
