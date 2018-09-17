package randgen

import (
	"testing"
)

func TestGenArray(t *testing.T) {
	n := 4
	arr1 := GenArray(n)
	arr2 := GenArray(n)
	if len(arr1) != n {
		t.Errorf("Unexpected arr1 length, got: %d, want: %d.", len(arr1), n)
	}
	if len(arr2) != n {
		t.Errorf("Unexpected arr2 length, got: %d, want: %d.", len(arr2), n)
	}
	same := true
	for i := 0; i < n; i++ {
		if arr1[i] != arr2[i] {
			same = false
			break
		}
	}
	if same {
		t.Errorf("GenArray() should be random, but got %v == %v\n", arr1, arr2)
	}
}

func TestGenVec(t *testing.T) {
	n := 4
	docVec1 := GenVec(n)
	docVec2 := GenVec(n)
	v1 := docVec1.Vec
	v2 := docVec2.Vec

	if len(v1) != n {
		t.Errorf("Unexpected v1 length, got: %d, want: %d.", len(v1), n)
	}
	if len(v2) != n {
		t.Errorf("Unexpected v2 length, got: %d, want: %d.", len(v2), n)
	}
	same := true
	for i := 0; i < n; i++ {
		if v1[i] != v2[i] {
			same = false
			break
		}
	}
	if same {
		t.Errorf("GenVec() should be random, but got %v == %v\n", v1, v2)
	}
}
