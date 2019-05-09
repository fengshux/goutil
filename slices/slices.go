package slices

import(
	"github.com/radrupt/goutil/integer"
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


func StringSliceIndexOf( s[]string, item string ) int {

	for i, v := range s {
		if ( v == item ) {
			return i
		}
	}
	return -1
}

func StringSliceRemove( s1, s2[]string ) []string {
	var result = []string{}

	for _, v := range s1 {
		if StringSliceIndexOf(s2, v) < 0 {
			result = append(result, v)
		}
	}
	return result
}
