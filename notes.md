# Notes

### No need to set up a type only to attacha method to due to http.handleFunc
*unnecessary*
```go
type myServerHandler struct{}

func (server *myServerHandler) ServeHTTP(myWriter http.ResponseWriter, myRequest *http.Request) {
	myWriter.Header().Set("Content-Type", "application/json")
 	switch myRequest.Method {
 	case "GET":
		myWriter.WriteHeader(http.StatusOK)
		myWriter.Write([]byte(`{"message": "hello world"}`))
	case "POST":
        myWriter.WriteHeader(http.StatusCreated)
        myWriter.Write([]byte(`{"message": "post called"}`))
    // Etc.
}

func main() {
	s := &myServerHandler{}
	http.Handle("/", s)
```
