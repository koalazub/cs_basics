package stringsthings

func CompareStrings(strs []string) string {

	if strs == nil {
		return ""
	}

	prefix := strs[0]
	for _, str := range strs[1:] {
		i := 0
		for i < len(prefix) && i < len(str) && str[i] == prefix[i] {
			i++
		}

		prefix = prefix[:i]
		if prefix == "" {
			return ""
		}
	}
	return prefix
}
