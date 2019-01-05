// +build ignore

package main

import "github.com/hunterlong/postman"

func main() {
	// OPTIONAL - create replacement values for postman variables
	variables := make(map[string]string)
	variables["endpoint"] = "http://localhost:8080"
	variables["service"] = "1"
	variables["message"] = "1"
	variables["group"] = "1"

	// create the postman config
	configs := postman.Config{
		Package:     "main",
		PostmanFile: "./postman.json",
		TestFile:    "server_test.go",
		RouterFunc:  "router()",
		Variables:   variables,
	}

	// generate the Golang Postman Tests
	err := configs.Generate()
	if err != nil {
		panic(err)
	}
}
