package main

import (
	"cours1/dictionary"
	"encoding/json"
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

	err := dict.SavuverEnFichier("./dictionary.txt")

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
	fmt.Fprintf(w, "Word : %s", word)

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

func ajouterEntree(dict *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	var entree dictionary.Entry
	err := json.NewDecoder(r.Body).Decode(&entree)
	if err != nil {
		http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
		log.Println("Erreur de décodage JSON :", err)
		return
	}

	if len(entree.Definition) < 3 || len(entree.Definition) < 5 {
		http.Error(w, "Les données ne respectent pas les règles de validation", http.StatusBadRequest)
		log.Println("Erreur de validation des données")
		return
	}

	err = dict.Add(entree.Word, entree.Definition)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Erreur lors de l'ajout de l'entrée :", err)
		return
	}

	reponse := map[string]string{"message": "Entrée ajoutée avec succès"}
	json.NewEncoder(w).Encode(reponse)
}

func definirEntree(dict *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	if len(word) < 3 {
		http.Error(w, "Paramètre de word invalide", http.StatusBadRequest)
		log.Println("Paramètre de word invalide")
		return
	}

	entree, err := dict.Get(word)
	if err != nil {
		http.Error(w, "Erreur", http.StatusNotFound)
		log.Println("Erreur", err)
		return
	}

	if entree.Word == "" {
		http.Error(w, "Word non trouvé", http.StatusNotFound)
		log.Println("Word non trouvé")
		return
	}

	json.NewEncoder(w).Encode(entree)
}

func supprimerEntree(dict *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	dict.Remove(word)
	reponse := map[string]string{"message": "Suppression réussie"}
	json.NewEncoder(w).Encode(reponse)
}

func listerEntrees(dict *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {

	entrees := dict.List()
	fmt.Println("Liste des entrées du dictionary :", entrees)

	json.NewEncoder(w).Encode(entrees)
}
