package internal

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thedevsaddam/gojsonq"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Foo struct{Name string}
func generate(structName string, amount int) *[]Foo {
	return &[]Foo{{"foo"}, {"bar"}}
}

type generateFn = func(amount int) []interface{}

func getHandler(structName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("STRUCTNAME", structName)

		amount := 10 // TODO: use quantity from path var
		response, err := json.Marshal(generate(structName, amount))

		if err != nil {
			fmt.Println("MARSHAL ERROR", err)
		}

		_, writeErr := w.Write(response)

		if writeErr != nil {
			fmt.Println("WRITE ERROR", writeErr)
		}
	}
}

func populateRouter(router *mux.Router, structNames []string) {
	for _, structName := range structNames {
		lowercased := strings.ToLower(structName)
		fmt.Println("Populating /" + lowercased)
		router.HandleFunc("/" + lowercased + "/{quantity}", getHandler(structName)).Methods("GET")
	}
}

func getSchemaEntityName() string {
	jq := gojsonq.New().File("../schema.avsc")
	res := jq.From("name").Get()
	return fmt.Sprint(res)
}

func StartServer(host string, port int) {
	r := mux.NewRouter()

	populateRouter(r, []string{getSchemaEntityName()})

	addr := host + ":" + strconv.Itoa(port)

	fmt.Println("Starting server on " + addr)
	log.Fatal(http.ListenAndServe(addr, r))
}