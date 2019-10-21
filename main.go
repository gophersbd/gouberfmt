package main

import (
	"github.com/gophersbd/gouberfmt/analyzers/copyboundary"
	"github.com/gophersbd/gouberfmt/analyzers/interfacepointer"
	"github.com/gophersbd/gouberfmt/analyzers/mutexpointer"

	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		interfacepointer.Analyzer,
		mutexpointer.Analyzer,
		copyboundary.Analyzer,
	)
}
