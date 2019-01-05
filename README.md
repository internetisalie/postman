# Golang Postman Test Generator
This golang package will automatically generate Golang unit tests directly from the Postman JSON export.


## Features
- [x] Create Go HTTP Tests from a Postman JSON file
- [ ] Include Authentication Headers from Postman to HTTP tests
- [ ] Include Content Type Headers from Postman to HTTP tests
- [ ] Convert Postman Javascript tests to Golang `assert` tests
- [ ] Dynamic variables based on Postman Javascript tests

## Installation
```go
go get github.com/hunterlong/postman
```

## Configuration
```go
configs := postman.Config{
    Package: "handlers",         // Package name for the test
    PostmanFile: "postman.json", // Location of the Postman JSON file
    TestFile: "api_test.go",     // Test filename to generate
    RouterFunc: "Router()",      // Function or variable name for the HTTP Router
    Variables: variables,        // Optional mapping of variables
    PreTest: pretest,            // Optional test to run first
}
```

## Example
This package requires a `go generate` file to automatically generator your golang tests. After you run the generator, the [api_test.go](https://github.com/hunterlong/postman/blob/master/examples/server_test.go) test file will be rendered. 
###### `generator.go` will create `api_test.go` with `go run generator.go`
```go
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

    // OPTIONAL - include a test to run first 
    pretest := `
    func TestResetHandlerDatabase(t *testing.T) {
        Clean()
        loadDatabase()
        resetDatabase()
    }`

    // create the postman config
    configs := postman.Config{
        Package: "handlers",
        PostmanFile: "../source/tmpl/postman.json",
        TestFile: "api_test.go",
        RouterFunc: "Router()",
        Variables: variables,
        PreTest: pretest,
    }
	
    // generate the Golang Postman Tests
    err := configs.Generate()
    if err != nil {
        panic(err)
    }
}
```