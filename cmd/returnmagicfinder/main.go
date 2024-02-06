package main

import (
	"github.com/mshr0969/returnmagicfinder"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(returnmagicfinder.Analyzer) }
