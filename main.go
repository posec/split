package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	lp := flag.Int("l", 1000, "segment length")
	ap := flag.Int("a", 2, "suffix length")

	flag.Parse()

	splitLine(os.Stdin, *lp, *ap, "x")

	os.Exit(0)
}

func splitLine(inio io.Reader, n, a int, prefix string) {
	i := 0
	seq := 0
	var out *os.File

	in := bufio.NewReader(inio)

	for {
		bs, err := in.ReadBytes('\n')
		if len(bs) > 0 {
			i += 1
			if out == nil {
				out, seq = NewOutput(seq, a, prefix)
			}
			out.Write(bs)
			if i >= n {
				out.Close()
				out = nil
			}
		}
		if err != nil {
			break
		}
	}

	return

}

const letters = "abcdefghijklmnopqrstuvwxyz"

func NewOutput(seq, a int, prefix string) (*os.File, int) {
	suffix := make([]byte, 0)
	s := seq
	for i := 0; i < a; i++ {
		d := s % 26
		s /= 26
		suffix = append([]byte{letters[d]},
			suffix...)
	}
	if s > 0 {
		log.Fatal("ran out of possible suffixes\n")
	}
	name := fmt.Sprintf("%s%s", prefix, suffix)
	out, err := os.OpenFile(name,
		os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	return out, seq + 1
}
