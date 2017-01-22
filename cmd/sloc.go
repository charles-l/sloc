package main

import (
	"fmt"
	"github.com/charles-l/sloc"
	"os"
	"path"
	"path/filepath"
)

func countLinesRecursive(p string) {
	slocs := make(map[*sloc.Language]int)
	filepath.Walk(p, func(pp string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		f, err := os.Open(pp)
		if err != nil {
			return err
		}

		e := path.Ext(pp)
		if e != "" {
			e = e[1:]
		}

		l := sloc.GetLanguage(e)
		slocs[l] += sloc.CountLines(f, l)
		f.Close()
		return nil
	})
	for k, v := range slocs {
		if k != nil {
			fmt.Printf("%s: %d SLOC\n", k.Name, v)
		} else {
			fmt.Printf("Other: %d L\n", v)
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [DIR]\n", path.Base(os.Args[0]))
	os.Exit(0)
}

func main() {
	if len(os.Args[1:]) != 1 {
		usage()
	}

	countLinesRecursive(os.Args[1])
}
