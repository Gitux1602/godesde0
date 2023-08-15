package main

import (
	"fmt"
	"godesde0/variables"
	"runtime"
)

func main() {
	variables.RestoVariables()
	variables.ConviertoaTexto(100)

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es Windows")
	} else {
		fmt.Println("Estas en Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linex")
	case "darwin":
		fmt.Println("Darwin")
	default:
		fmt.Printf("Operative System: %s", os)
	}
}
