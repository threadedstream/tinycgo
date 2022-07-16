package aid

func Contains(val any, ins ...any) bool {
	for _, in := range ins {
		if val == in {
			return true
		}
	}
	return false
}
