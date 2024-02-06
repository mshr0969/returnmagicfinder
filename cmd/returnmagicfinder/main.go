package main

import (
	"returnmagicfinder"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(returnmagicfinder.Analyzer) }
