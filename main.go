package main

import (
    "fmt"
    "cours1/dictionary"
)

func main() {

    //*TEST POUR COMPRENDRE **/
    fmt.Println("Hello, world.s			d")

	maMap := make(map[string]string)
    maMap["article1"] = "VOIZL "
    maMap["article2"] = " ERR"

	for titre, description := range maMap {
        fmt.Println("Titre :", titre)
        fmt.Println("Description :", description)
        fmt.Println()
    }

    fmt.Println("Carte   :", maMap)

    //*FIN DU TEST POUR COMPRENDRE **/

    dict := dictionary.New() 
    dict.Add("mot1", "definition1")
    dict.Add("mot2", "definition2")

    err := dict.SavuverEnFichier("./dictionnaire.txt")

    if err != nil {
        fmt.Println("Erreurrrrrrrr:", err)
    } else {
        fmt.Println("Succèssssssssssssss.")
    }

	if dict == nil {
		fmt.Println("Dictionary  nul")
	} else {
		fmt.Println("Dictionary plein")
	}
}