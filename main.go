package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type JsonAnalyzer struct{}

func (e *JsonAnalyzer) getTagData(tag reflect.StructTag) (string, string) {
	var tagName, tagValue string

	tagName = "json"
	tagValue = tag.Get(tagName)
	if tagValue != "" {
		return tagName, tagValue
	}

	return "", ""
}

func (e *JsonAnalyzer) DataAnalysisInCaseOfError(expectedStruct any, toAnalize map[string]any) (err error) {
	for k, v := range toAnalize {
		err = e.typeAnalysis(expectedStruct, k, reflect.TypeOf(v), reflect.ValueOf(v))
		if err != nil {
			return
		}
	}

	return
}

func (e *JsonAnalyzer) KeyAnalysis(expectedStruct any, jsonData []byte) (err error) {
	structElements := reflect.ValueOf(expectedStruct).Elem()

	var tagValue string
	var jsonParserd map[string]any

	err = json.Unmarshal(jsonData, &jsonParserd)
	if err != nil {
		return
	}

	for key := range jsonParserd {
		var pass = false

		for i := 0; i < structElements.NumField(); i += 1 {
			typeField := structElements.Type().Field(i)
			tag := typeField.Tag

			_, tagValue = e.getTagData(tag)

			if tagValue == "-" || tagValue == "" {
				continue
			}

			if tagValue == key {
				pass = true
			}
		}

		if pass == false {
			err = fmt.Errorf("there is a new key in the received json that was not predicted in the json handling code. key: %v", key)
			return
		}
	}

	return
}

func (e *JsonAnalyzer) typeAnalysis(expectedStruct any, key string, typeOfTheArrivedData reflect.Type, valueOfTheArrivedData reflect.Value) (err error) {
	structElements := reflect.ValueOf(expectedStruct).Elem()

	var tagValue string

	for i := 0; i < structElements.NumField(); i += 1 {
		typeField := structElements.Type().Field(i)
		tag := typeField.Tag

		_, tagValue = e.getTagData(tag)

		if tagValue == "-" || tagValue == "" {
			continue
		}

		if tagValue == key {
			if typeOfTheArrivedData.Kind() != typeField.Type.Kind() {
				err = fmt.Errorf("there is an inconsistency between the received type (%v) and the expected type (%v). value: `%v`", typeOfTheArrivedData.Kind().String(), typeField.Type.Kind().String(), valueOfTheArrivedData.Interface())
				return
			}
		}
	}

	return
}

// Test Este é o struct do dado esperado. Ele é feito para ter todos os campos recebidos pelo json, mesmo os não usados.
// Se houverem dois structs, um com apenas as chaves usadas, e outro completo para analize, o desempenho pode ser melhorado.
type Test struct {

	// O struct recebe a injeção de objeto para herdar as propriedades do mesmo
	JsonAnalyzer

	// Todos os dados esperados daqui para baixo
	Id int `json:"id"`
}

func main() {
	log.Printf("O dado json está conforme esperado")
	example([]byte(`{"id":2}`))
	log.Print("")

	log.Printf("O dado json deveria ter o campo id como inteiro, mas, enviou float")
	example([]byte(`{"id":2.2}`))
	log.Print("")

	log.Printf("O dado json deveria ter o campo id como inteiro, mas, enviou string")
	example([]byte(`{"id":"2"}`))
	log.Print("")

	log.Printf("O dado json chegou correto, mas, com um campo a mais, não previsto na documentação")
	example([]byte(`{"id":2, "name":"Dino Da Silva Sauro"}`))
	log.Print("")
}

func example(data []byte) {
	t := &Test{}

	err := json.Unmarshal(data, t)
	if err != nil {
		analize := make(map[string]any)
		err = json.Unmarshal(data, &analize)
		if err != nil {
			panic(err)
		}

		err = t.DataAnalysisInCaseOfError(t, analize)
		if err != nil {
			log.Printf("pismo transaction parser problem: %v", err)
		}
	}

	// Aqui eu colocaria um valor aleatório para que sejam processados 1% dos dados recebidos
	// assim, não se gasta muito poder de processamento do servidor processando tudo o que entra e ao mesmo tempo, permite
	// uma verificação em caso de mudança dos dados recebidos da pismo
	err = t.KeyAnalysis(t, data)
	if err != nil {
		log.Printf("pismo transaction parser problem: %v", err)
	}
}

//
//
//
//
//
//
//
//
//
//
//
//
//
