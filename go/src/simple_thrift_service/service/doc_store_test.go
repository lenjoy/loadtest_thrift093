package service

import (
	"testing"
)

func TestGetVector(t *testing.T) {
	ds := NewDocStore()
	dv := ds.GetVector(123)
	if len(dv.Vec) != int(ds.Dim) {
		t.Errorf("Unexpected Vec length, got: %d, want: %d.", len(dv.Vec), ds.Dim)
	}
}

func TestGetRelatedDocs(t *testing.T) {
	seedDocID := int32(9527)
	n := 5
	ds := NewDocStore()
	helloDocArr := ds.GetRelatedDocs(seedDocID, n)
	if len(helloDocArr) != n {
		t.Errorf("Unexpected helloDocArr length, got: %d, want: %d.", len(helloDocArr), n)
	}
	for i := 1; i < len(helloDocArr); i++ {
		if *helloDocArr[i].Score > *helloDocArr[i-1].Score {
			t.Errorf("Expect helloDocArr[%d].Score %f > helloDocArr[%d].Score %f",
				i-1, *helloDocArr[i-1].Score, i, *helloDocArr[i].Score)
		}
		break
	}
}
