package main

import (
	"flag"

	"github.com/Rajeevnita1993/cut-tool/fileio"
)

func main() {

	// Parse command line arguments
	fieldListPtr := flag.String("f", "1", "field number")
	delimiterPtr := flag.String("d", "\t", "delimiter")
	flag.Parse()

	// Check if a filename is provided
	if flag.NArg() == 0 || flag.Arg(0) == "-" {
		// Read from standard input
		fileio.CutfFromStdin(*fieldListPtr, *delimiterPtr)
	} else {
		// Read from the provided file
		filename := flag.Arg(0)
		fileio.Cutf(filename, *fieldListPtr, *delimiterPtr)
	}

}
