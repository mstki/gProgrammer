package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]*fileObj)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, obj := range counts {
		if obj.count > 1 {
			fmt.Printf("%d\t%s\t%s\n", obj.count, line, obj.name)
		}
	}
}

type fileObj struct {
	count int
	name  []string
}

func countLines(f *os.File, obj map[string]*fileObj) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if len(line) == 0 {
			break
		}
		if obj[line] == nil {
			obj[line] = &fileObj{}
		}
		obj[line].count++
		obj[line].name = append(obj[line].name, f.Name())
	}
	//NOTES:ignoring potential errors from input.Err()
}
