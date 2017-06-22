package helper

func ChunkSlice(data []interface{}, size int) []interface{} {
	var chunks []interface{}

	for i := 0; i < len(data); i += size {
		end := i + size

		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, data[i:end])
	}
	return chunks
}
