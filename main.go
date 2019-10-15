package main

import (
	"github.com/gophersbd/gouberfmt/analyzers/interfacepointer"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		interfacepointer.Analyzer,
	)
}
