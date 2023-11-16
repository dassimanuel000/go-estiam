package main

import "fmt"

func main() {
    fmt.Println("Hello, world.s			d")

	maMap := make(map[string]string)
    maMap["article1"] = "VOIZL "
    maMap["article2"] = " ERR"

	for titre, description := range maMap {
        fmt.Println("Titre :", titre)
        fmt.Println("Description :", description)
        fmt.Println()
    }

    fmt.Println("Carte vide :", maMap)
}