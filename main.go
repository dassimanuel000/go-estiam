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
	router.Use(gestionnaireErreur)

	router.HandleFunc("/accueil", GetDefinition).Methods("GET")

	router.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		listerEntrees(dict, w, r)
	}).Methods("GET")

	router.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		ajouterEntree(dict, w, r)
	}).Methods("POST")

	router.HandleFunc("/definir/{word}", func(w http.ResponseWriter, r *http.Request) {
		definirEntree(dict, w, r)
	}).Methods("GET")

	router.HandleFunc("/retirer/{word}", func(w http.ResponseWriter, r *http.Request) {
		supprimerEntree(dict, w, r)
	}).Methods("DELETE")

	http.Handle("/", router)

	fmt.Printf("Serveur écoutant sur le port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, router))

	//router.HandleFunc("/api/{word}", GetDefinition).Methods("GET")

}

func GetDefinition(w http.ResponseWriter, r *http.Request) {

	// Écrivez une réponse simple
	w.WriteHeader(http.StatusOK)                // Code de statut HTTP 200 (OK)
	w.Header().Set("Content-Type", "text/html") // Type de contenu HTML

	message := "<html><body><h3>Bienvenue sur ma page web</h3></body></html>"
	fmt.Fprintf(w, message)
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
	if err := json.NewDecoder(r.Body).Decode(&entree); err != nil {
		http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
		log.Println("Erreur de décodage JSON :", err)
		return
	}

	if len(entree.Word) < 3 || len(entree.Definition) < 5 {
		http.Error(w, "Les données ne respectent pas les règles de validation", http.StatusBadRequest)
		log.Println("Erreur de validation des données")
		return
	}

	if err := dict.Add(entree.Word, entree.Definition); err != nil {
		http.Error(w, "Erreur lors de l'ajout de l'entrée", http.StatusInternalServerError)
		log.Println("Erreur lors de l'ajout de l'entrée :", err)
		return
	}

	reponse := map[string]string{"message": "Entrée ajoutée avec succès"}
	w.WriteHeader(http.StatusCreated) // Code de statut HTTP 201 (Created)
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
	//json.NewEncoder(w).Encode(entrees)

	var message string

	message += "<html><body><h3>listerEntrees</h3></body></html>"
	for mot, entree := range entrees {
		message += fmt.Sprintf("Mot : %s, Définition : %s\n", mot, entree.Definition)
	}

	w.WriteHeader(http.StatusOK)                // Code de statut HTTP 200 (OK)
	w.Header().Set("Content-Type", "text/html") // Type de contenu HTML

	fmt.Fprintf(w, message)
}
