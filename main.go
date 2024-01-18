package main

import (
	"cours1/dictionary"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	dict := dictionary.NewDictionary()
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

	router := mux.NewRouter()
	port := ":8080"

	fmt.Printf("Serveur écoutant sur le port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, router))

	router.Use(gestionnaireErreur)

	router.HandleFunc("/api/{word}", GetDefinition).Methods("GET")

}

func GetDefinition(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, World!")
	params := mux.Vars(r)
	word := params["word"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Mot : %s", word)

	/*entry, err := dictionary.NewDictionary().Get(word)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonResponse, _ := json.Marshal(entry)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)*/
}

func gestionnaireErreur(suivant http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Erreur non capturée :", err)
				http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			}
		}()

		suivant.ServeHTTP(w, r)
	})
}
