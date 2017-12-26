package slices

import(
	"goutil/integer"
)

func ChunkStringSlice( ss []string, size int ) ( [][]string, int ) {

	var (
		cursor int = 0
		end = size
		length = len( ss )
		pice = length / size + 1
		chunk [][]string
	)

	for i:=0; i < pice ; i++ {
		
		end = integer.Min( end, length )
		chunk = append( chunk, ss[cursor:end] )
		cursor = end
		end = end + size
	}

	return chunk, pice
}


func StringSliceEqual(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }

    if (a == nil) != (b == nil) {
        return false
    }

    for i, v := range a {
        if v != b[i] {
            return false
        }
    }

    return true
}
