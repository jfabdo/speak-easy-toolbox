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

//GetFileServer returns a file server to serve static files
func GetFileServer(fulladdress string, port string, dir string) {
	fs := http.FileServer(http.Dir(dir))
	http.Handle(fulladdress, fs)

	log.Println("Listening on :" + port + "...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
