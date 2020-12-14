package toolbox

import (
	"fmt"
	"log"
	"net/http"
)

//GetServer opens an http server at the address string to the function
func GetServer(fulladdress string, port string, fn func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(fulladdress, fn)
	// http.HandleFunc("", api.Messaging)
	fmt.Printf("Starting server at port " + port + "\n")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}