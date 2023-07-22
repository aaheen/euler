package main

import (
	"eaheen/euler/ans"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var usage string = `Usage: euler
		<p> 	- Run the solution for problem <p>. Directly outputs answer
	-i	<p> 	- Run the interactive solution session for problem <p>
	-s 	<p> 	- Open the writeup for problem <p>
	-g 		- Open this repo on GitHub`

	fInter := flag.Bool("i", false, "Run the interactive solution session for problem <p>")
	fWriteup := flag.Bool("s", false, "Open the writeup for problem <p>")
	fGithub := flag.Bool("g", false, "Open this repo on GitHub")
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
		case *fWriteup:
			ans.Writeup(p)
		case *fGithub:
			fmt.Println(usage)
		default:
			ans.Sol(p)
		}
	default:
		fmt.Println(usage)
	}

}
