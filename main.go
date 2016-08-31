package main

import (
	"flag"
	"fmt"
	"image/color"
	"image/gif"
	"log"
	"math"
	"os"

	"github.com/soh335/truecolor"
)

var (
	file   = flag.String("file", "", "input file")
	number = flag.Int("number", 1, "number of image")
)

func main() {
	flag.Parse()
	if err := _main(); err != nil {
		log.Fatal(err)
	}
}

func _main() error {
	f, err := os.Open(*file)
	if err != nil {
		return err
	}
	defer f.Close()
	g, err := gif.DecodeAll(f)
	if err != nil {
		return err
	}
	if l := len(g.Image); l < *number {
		return fmt.Errorf("number of image is %d, but number is %d", l, number)
	}
	p := g.Image[*number-1].ColorModel().(color.Palette)
	size := int(math.Ceil(math.Sqrt(float64(len(p)))))
	for i, c := range p {
		tc := truecolor.New()
		tc.Add(truecolor.NewBackgrond(c))
		tc.Fprint(os.Stdout, "  ")
		if i%size == size-1 {
			fmt.Print("\n")
		}
	}
	return nil
}
