package loop_over

func LoopOver[T any](sl []T) func() T {
	maxLen := len(sl)
	idx := -1
	return func() T {
		idx++
		if idx == maxLen {
			idx = 0
		}
		return sl[idx]
	}
}
