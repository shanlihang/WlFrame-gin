package arrayProcessing

// []uint去重
func RemoveDuplicates(ID2s []uint) []uint {
	seen := make(map[uint]bool)
	result := []uint{}

	for _, id := range ID2s {
		if _, exists := seen[id]; !exists {
			seen[id] = true
			result = append(result, id)
		}
	}

	return result
}
