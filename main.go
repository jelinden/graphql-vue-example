package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
)

var schema graphql.Schema

const jsonDataFile = "data.json"

func filterPerson(data []map[string]interface{}, args map[string]interface{}) map[string]interface{} {
	for _, person := range data {
		for k, v := range args {
			if person[k] != v {
				goto nextperson
			}
			return person
		}

	nextperson:
	}
	return nil
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v\n", result.Errors)
	}
	return result
}

func importJSONDataFromFile(fileName string) error {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	var data []map[string]interface{}

	err = json.Unmarshal(content, &data)
	if err != nil {
		return err
	}

	fields := make(graphql.Fields)
	args := make(graphql.FieldConfigArgument)
	for _, item := range data {
		for k := range item {
			fields[k] = &graphql.Field{
				Type: graphql.String,
			}
			args[k] = &graphql.ArgumentConfig{
				Type: graphql.String,
			}
		}
	}

	var personType = graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "person",
			Fields: fields,
		},
	)

	var queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"person": &graphql.Field{
				Type: personType,
				Args: args,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return filterPerson(data, p.Args), nil
				},
			},
			"personList": &graphql.Field{
				Type: graphql.NewList(personType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return data, nil
				},
			},
		},
	})

	schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)

	return nil
}

func main() {
	err := importJSONDataFromFile(jsonDataFile)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query()["query"][0], schema)
		w.Header().Add("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(result)
	})
	fs := http.FileServer(http.Dir("frontend/dist"))
	http.Handle("/", fs)

	fmt.Println("Now server is running on port 8000")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8000/graphql?query={person(name:\"Luke\"){id,name,surname}}'")
	http.ListenAndServe(":8000", nil)
}
