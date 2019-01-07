package stitch

import (
	"image/gif"
	"io"
	"os"

	"github.com/pkg/errors"
)

// File represents a single GIF to be added to the final product. The full GIF
// will be added at least once, but possibly more depending on the arguments
// supplied to the Create function.
type File struct {
	name  string
	loops int
}

// FileList is a list of Files to process.
type FileList []File

// Create takes a list of files and a list of loop counts. If the list of loop
// counts is empty, each file will be added to the final product exactly once.
// Otherwise, the number of loop counts must equal the number of files, and each
// file will be added a number of times equal to its loop count.
func Create(files []string, loops []int) (GIF, error) {
	fl, err := parseArgs(files, loops)
	if err != nil {
		return GIF{}, err
	}

	var g GIF
	for _, f := range fl {
		if err := g.Add(f); err != nil {
			return GIF{}, err
		}
	}

	return g, nil
}

// GIF is a wrapper around a gif.GIF from the standard library.
type GIF struct {
	g gif.GIF
}

// Add adds the provided file to the GIF a number of times equal to its loop
// count.
func (g *GIF) Add(f File) error {
	in, err := os.Open(f.name)
	if err != nil {
		return errors.Errorf("could not open gif: %s: %v", f.name, err)
	}
	defer in.Close()

	img, err := gif.DecodeAll(in)
	if err != nil {
		return errors.Errorf("could not decode gif: %s: %v\n", f.name, err)
	}

	for i := 0; i < f.loops; i++ {
		g.g.Image = append(g.g.Image, img.Image...)
		g.g.Delay = append(g.g.Delay, img.Delay...)
	}

	return nil
}

// Save outputs the GIF to the file with the name specified.
func (g *GIF) Save(fn string) error {
	of, err := os.Create("merged.gif")
	if err != nil {
		return errors.Errorf("could not create merged gif: %v\n", err)
	}
	defer of.Close()

	if err := g.Encode(of); err != nil {
		return errors.Errorf("could not encode merged gif: %v\n", err)
	}
	return nil
}

// Encode outputs the GIF to the supplied io.Writer.
func (g *GIF) Encode(w io.Writer) error {
	if err := gif.EncodeAll(w, &g.g); err != nil {
		return errors.Errorf("could not encode merged gif: %v\n", err)
	}
	return nil
}

// parseArgs is a helper to make sure that the supplied arguments are of the
// correct size and that each file gets added at least once.
func parseArgs(files []string, loops []int) (FileList, error) {
	if len(loops) != 0 && len(loops) != len(files) {
		return nil, errors.Errorf("number of files (%d) and loops (%d) don't match", len(files), len(loops))
	}

	if len(loops) == 0 {
		for range files {
			loops = append(loops, 1)
		}
	}

	var fl FileList
	for i := range files {
		f := File{
			name:  files[i],
			loops: loops[i],
		}
		fl = append(fl, f)
	}

	return fl, nil
}
