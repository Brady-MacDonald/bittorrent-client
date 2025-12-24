package main

import "os"

func main() {
	file := os.Args[1]
	output := os.Args[2]

	if file == "" || output == "" {
		panic("Bad")
	}

}
