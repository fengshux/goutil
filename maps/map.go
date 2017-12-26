package maps

func Keys( m map[string]string ) []string {
	var keys []string

	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}


