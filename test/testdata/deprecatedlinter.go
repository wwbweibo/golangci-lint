package testdata

import (
	"compress/gzip"
	"log"

	"golang.org/x/tools/go/analysis"
)

func SpewDebugInfo() {
	log.Println(gzip.BestCompression)
	_ = analysis.Analyzer{}
}
