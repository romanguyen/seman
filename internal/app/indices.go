package app

func (m *Model) visibleIndex(items []int, value int) int {
	for i, idx := range items {
		if idx == value {
			return i
		}
	}
	return -1
}
