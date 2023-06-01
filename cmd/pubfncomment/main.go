package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/sivchari/pubfncomment"
)

func main() { unitchecker.Main(pubfncomment.Analyzer) }
