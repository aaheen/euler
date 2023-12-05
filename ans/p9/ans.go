package p9

import "fmt"

// Returns solution for Problem 9
func Sol() {
	tri := brutePythagTriN(1000)
	sum := sumTriple(tri)
	prod := prodTriple(tri)
	fmt.Println(tri, "Sum:", sum, "Product:", prod)
}

// Interactive solution session for Problem 9
func Ask() {
	fmt.Print("Goal \nInput n: ")
	
	var n int
	fmt.Scanln(&n)
	tri := brutePythagTriN(n)
	sum := sumTriple(tri)
	prod := prodTriple(tri)
	
	if tri == [3]int{-1, -1, -1} {
		fmt.Println("No triple with sum", n)
	} else {
		fmt.Println(tri, "Sum:", sum, "Product:", prod)
	}
}
