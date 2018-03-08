package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	lp := flag.Int("l", 1000, "segment length")
	ap := flag.Int("a", 2, "suffix length")

	flag.Parse()
	fmt.Println("a", *ap)
	fmt.Println("l", *lp)
	os.Exit(0)
}
