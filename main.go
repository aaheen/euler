package main

import (
	"flag"
	"fmt"
	"m9ple/euler/ans"
	"os"
	"strconv"
)

func main() {

	var usage string = `Usage: euler
		<p> 	- Run the solution for problem <p>. Directly outputs answer
	-i	<p> 	- Run the interactive solution session for problem <p>
	-g 		- Open this repo on github.com/m9ple/euler
	-g 	<p> 	- Open the subdirectory for problem <p> in this repo on github`

	fInter := flag.Bool("i", false, "Run the interactive solution session for problem <p>")
	fGithub := flag.Bool("g", false, "Open the subdirectory for problem <p> in this repo on github")
	flag.Parse()
	fArgs := flag.Args()

	switch len(fArgs) {
	case 0: // Either open Github or invalid usage
		if *fGithub {
			ans.Repo()
		} else {
			fmt.Println(usage)
		}
	case 1: // Asking about a specific problem number
		p, err := strconv.Atoi(fArgs[0])
		if err != nil {
			fmt.Println("Error parsing problem ID", fArgs[0])
			os.Exit(2)
		}
		// Recover if Sol, Ask, or solution code panics
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Error trying to run problem", p, ":\n", err)
				os.Exit(1)
			}
		}()

		switch {
		case *fInter:
			ans.Ask(p)
		case *fGithub:
			ans.Page(p)
		default:
			ans.Sol(p)
		}
	default:
		fmt.Println(usage)
	}

}
