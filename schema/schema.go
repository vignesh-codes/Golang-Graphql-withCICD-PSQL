package schema

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"fmt"
)

var (
	rawSchema []byte
	Schema    string
)

func init() {
	fmt.Println(rawSchema)
	path, _ := filepath.Abs("./schema/schema.graphql")
	var rawSchema, err = ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error getting schema", err)
	}
	Schema = string(rawSchema)
	
}
