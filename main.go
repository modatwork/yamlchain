package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	help := flag.Bool("h", false, "show usage")
	flag.Parse()
	if *help {
		fmt.Println("yamlchain merges multiple YAML files into a single file, inserting a '---' separator line between each file's content. The combined output is then written to stdout.")
		fmt.Println("Usage: yamlchain file1 file2 ...")
		return
	}

	chained := 0
	for _, file := range flag.Args() {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		chained++
		if chained < len(flag.Args()) {
			fmt.Println("---")
		}
	}
}
