package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Clientes struct{
	Clientes []Cliente `json:"clientes"`
}

type Cliente struct {
	Nome string  `json:"nome"` 
	Idade int  `json:"idade"`
    Telefone string `json:"telefone"`
}

//como não uso banco de dados, guardo os clientes no array
var clientesArray []Cliente

func main() {

	rotas := mux.NewRouter().StrictSlash(true)

	rotas.HandleFunc("/cliente", getAll).Methods("GET")
	rotas.HandleFunc("/cliente", create).Methods("POST")
	var port = ":8080"
	fmt.Println("Servidor rodando na porta:", port)
	log.Fatal(http.ListenAndServe(port, rotas))
	

}


//o Get faz a busca no array
func getAll(w http.ResponseWriter, r *http.Request) {
	
	json.NewEncoder(w).Encode(clientesArray)
	
}

func create(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var clientesBody Cliente
	
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}
	
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	
	if err := json.Unmarshal(body, &clientesBody); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	
	json.Unmarshal(body, &clientesBody)

	//for i := 0; i < len(clientesBody.Cliente); i++ {
	//	clientesArray = append(clientesArray, clientesBody.Clientes[i])
	//}

	clientesArray = append(clientesArray, clientesBody)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(clientesBody); err != nil {
		panic(err)
	}
}