package main

import (
	"fmt"
	"os"

	"github.com/schildwech/sycli/src/internal/spotify"
)

func main() {

	if err := spotify.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
