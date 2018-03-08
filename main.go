package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	lp := flag.Int("l", 1000, "segment length")
	ap := flag.Int("a", 2, "suffix length")
	bp := flag.String("b", "", "split by bytes")
	in := os.Stdin

	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		var err error
		in, err = os.Open(args[0])
		if err != nil {
			log.Fatal(err)
		}
	}
	prefix := "x"
	if len(args) > 1 {
		prefix = args[1]
	}

	if *bp == "" {
		splitLine(in, *lp, *ap, prefix)
	} else {
		splitBytes(in, *bp, *ap, prefix)
	}

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

func splitBytes(in io.Reader, b string, a int, prefix string) {
	seq := 0
	factor := 1
	if b[len(b)-1] == 'k' {
		b = b[:len(b)-1]
		factor = 1000
	}
	if b[len(b)-1] == 'm' {
		b = b[:len(b)-1]
		factor = 1000000
	}
	l, err := strconv.Atoi(b)
	if err != nil {
		log.Fatal(err)
	}
	l *= factor
	buf := make([]byte, l)

	for {
		n, err := in.Read(buf)
		if n > 0 {
			var out *os.File
			out, seq = NewOutput(seq, a, prefix)
			out.Write(buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
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
		os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}
	return out, seq + 1
}
