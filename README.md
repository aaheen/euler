# [Project Euler](https://projecteuler.net)
**Solutions by Erik Heen**, *written in Go*.

I have tried to generalize these solutions somewhat, so they can be repurposed later. Each subdirectory of [`ans`](ans) is numbered according to the problem it solves on [Project Euler](https://projecteuler.net). Presently, I have only a few problems solved. If in the future this number grows, I will most likely fork this repository and make it private, leaving only the first 100 open for outside examination.

## Structure:

* This is all under the Go module `m9ple/euler`.
* `main` accesses answers through the [`ans`](ans) package, through the `Sol(p int)` and `Ask(p int)` functions, where `p` is the problem number.
* Each problem is contained in its own sub-package of [`ans`](ans), named `p<number>`.
* ***Example:*** The full name for the package that contains the solution for Problem 3 is [`m9ple/euler/ans/p3`](ans/p3). To get the exact answer for Problem 3, call `ans.Sol(3)`. To run the interactive session for Problem 3, call `ans.Ask(3)`.
* **Note:** All of the $$ you see in problem READMEs is just unrendered LaTeX. GitHub does not currently support any easy method of rendering LaTeX. This is such a hotly demanded feature that I have to believe they will do so sometime relatively soon. As such, I'm not using any workarounds to display it anyway, because it's too much work.
    * [![xhub](https://img.shields.io/badge/Chrome%20extension-xhub-f2eecb?style=flat-square)](https://chrome.google.com/webstore/detail/xhub/anidddebgkllnnnnjfkmjcaallemhjee) is a great workaround that I considered using, but decided against it.
    * Codecogs is also a great resource, I just didn't feel like it.

## The `euler` CLI

A utility used to run any solution, without having to travel to each solution's directory and `go run .`

**I am *very* new to Go**, so I'm still trying to figure out how to implement this. I'm literally just grabbing from two big `map[int]func()` (declared in [probmap.go](ans/probmap.go)), where the index corresponds to the problem number. I have struggled trying to come up with a more elegant implementation, but right now I'm not sure what that would be.

## Usage:
```
euler  <p>     - Run the solution for problem <p>. Directly outputs answer
    -i <p>     - Run the interactive solution session for problem <p>
    -g         - Open this repo at github.com/m9ple/euler
    -g <p>     - Open the subdirectory for problem <p> on github.
```

![Portrait of Leonhard Euler](euler.png)
