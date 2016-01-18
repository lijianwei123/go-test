package service

import (
	"testing"
)

func Test_Batch(t *testing.T) {
	bc := BatchConfig{
		Start:   0,
		Limit:   10,
		Size:    100,
		ReqTmpl: "http://localhost/test10.php?offset=%s&limit=%s",
	}
	batch := NewBatch(bc)
	batch.Run()
}
