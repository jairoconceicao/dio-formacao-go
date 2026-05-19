package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jairoconceicao/dio-formacao-go/api/responses"
)

const POKEAPI_URL = "http://pokeapi.co/api/v2/pokedex/kanto/"

func main() {
	response, err := http.Get(POKEAPI_URL)

	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		os.Exit(1)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
		log.Fatal(err)
	}

	var responseData responses.PokedexResponse

	errjson := json.Unmarshal(body, &responseData)

	if errjson != nil {
		fmt.Println("Erro ao fazer o parse do JSON:", errjson)
		log.Fatal(errjson)
	}

	fmt.Println(responseData)
}
