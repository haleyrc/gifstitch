package main

import (
	"flag"
	"log"
	"strconv"
	"strings"

	"github.com/haleyrc/gifstitch/stitch"
	"github.com/pkg/errors"
)

func main() {
	var (
		fs     = flag.String("files", "", "A comma-separated list of files to parse")
		ls     = flag.String("loops", "", "A comma-separated list of loop counts for each file (count should equal number of files)")
		ofname = flag.String("o", "merged.gif", "The name of the output file; defaults to merged.gif")
	)
	flag.Parse()

	files, err := parseFiles(*fs)
	if err != nil {
		log.Fatalln(err)
	}

	loops, err := parseLoops(*ls)
	if err != nil {
		log.Fatalln(err)
	}

	gif, err := stitch.Create(files, loops)
	if err != nil {
		log.Fatalln(err)
	}

	if err := gif.Save(*ofname); err != nil {
		log.Fatalln(err)
	}
}

func parseFiles(s string) ([]string, error) {
	if s == "" {
		return nil, errors.Errorf("no input files specified")
	}

	files := strings.Split(s, ",")
	return files, nil
}

func parseLoops(s string) ([]int, error) {
	if s == "" {
		return []int{}, nil
	}

	var ls []int
	svals := strings.Split(s, ",")
	for _, sval := range svals {
		val, err := strconv.Atoi(sval)
		if err != nil {
			return nil, errors.Errorf("could not parse loop count: %s: %v", sval, err)
		}
		ls = append(ls, val)
	}

	return ls, nil
}
