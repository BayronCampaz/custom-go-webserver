# How to use
This is an example of how to use this module usin a main.go file.
package main

```
func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("POST", "/api", server.AddMidleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}
```
