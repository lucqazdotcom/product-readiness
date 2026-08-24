package ingestion

import (
	"path"
	"testing"
)

func BenchmarkIngestion(b *testing.B) {
	datadir := path.Join("..", "..", "testdata")
	for b.Loop() {
		if err := Ingestion(datadir); err != nil {
			b.Fatal(err)
		}
	}
}
