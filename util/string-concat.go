package util

func SConcat(strs ...string) string {
	var str string
	for _, s := range strs {
		str += s
	}

	return str
}
