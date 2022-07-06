package map_practice

func RemoveDuplicates(input []string) []string {
	d := make(map[string]bool, len(input))
	r := make([]string, len(input))

	outputIdx := 0

	for _, v := range input {
		if _, ok := d[v]; !ok {
			d[v] = true
			r[outputIdx] = v
			outputIdx++
		}
	}

	return r[:outputIdx]
}
