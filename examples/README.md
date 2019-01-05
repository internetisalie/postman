### `generator.go`
The Postman Test generator file.

### `main.go`
A basic HTTP server with a couple routes. By including `//go:generate go run generator.go` under the package name, you can generate the test file with `go generate`.

### `postman.json`
The Postman JSON export file that includes all of your routes, method, and body.

### `server_test.go`
Generated Golang test file that includes your Postman routes and tests.