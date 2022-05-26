package main

import (
    "fmt"
    "io/ioutil"
    "log"
	"strconv"
    "net/http"
)

type NodoOperacion struct {
	id			string
	tipo 		string
	input1		NodoDato
	input2		NodoDato
}

type NodoDato struct{
	id			string
	valor 		string
}

func calcular(a NodoOperacion) string {
	if a.tipo == "+"
		return Itoa(Atoi(a.input1.valor) + Atoi(a.input2.valor))
	if a.tipo == "*"
		return Itoa(Atoi(a.input1.valor) * Atoi(a.input2.valor))
	if a.tipo == "-"
		if Atoi(a.input1.valor) >=  Atoi(a.input2.valor)
			return Itoa(Atoi(a.input1.valor) - Atoi(a.input2.valor))
		else
			return Itoa(Atoi(a.input2.valor) - Atoi(a.input1.valor))
	if a.tipo == "/"
		if (Atoi(a.input1.valor) && Atoi(a.input2.valor)) > 0
			return Itoa(Atoi(a.input1.valor) / Atoi(a.input2.valor))
	if a.tipo == "="
		if a.input1.valor == a.input2.valor
			return "Verdadero"
		else
			return "Falso"
	if a.tipo == ">="
		if a.input1.valor >= a.input2.valor
			return "Verdadero"
		else
			return "Falso"
	if a.tipo == "=<"
		if a.input1.valor =< a.input2.valor
			return "Verdadero"
		else
			return "Falso"
	if a.tipo == "!="
		if a.input1.valor != a.input2.valor
			return "Verdadero"
		else
			return "Falso"	
	
}

func main(NodoOperacion) {

    resp, err := http.Get("http://127.0.0.1:5503/")

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(body))
}

