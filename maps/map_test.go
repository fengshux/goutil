package maps


import (
	"testing"
)


func TestKeys( t *testing.T) {

	m := map[string]string{
		"a":"1",
		"b":"2",
	}

	s := Keys(m)
	t.Log(s)
	if len(m) != len(s) {
		t.Fatal("s and m length is not equal")
	}

}
