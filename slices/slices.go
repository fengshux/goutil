package slices

func ChunkStringSlice( ss []string, size int ) ( [][]string, int ) {

	var (
		sursor int = 0
		end = size
		length := len( ss )
		pice := length / size + 1
		chunk [][]string
	)

	for i:=0; i < pice ; i++ {
		
		end = min( end, length )
		chunk = append( chunk, ss[sursor:end] )
		cursor = end
		end = end + size
	}

	return chunk
}
