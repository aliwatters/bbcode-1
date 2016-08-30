package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/aliwatters/go-bbcode"
)

// TODO: all the custom bbcode for travelblog.org

func main() {
	compiler := bbcode.NewCompiler(true, true, false)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(compiler.Compile(scanner.Text()))
	}
}
