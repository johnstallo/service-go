package main

import(
"fmt"
"net/http"
"os"
)

func indexHandler( w http.ResponseWriter, r *http.Request){
    hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hello from service B running on ", hostname)
}

func main(){
    http.HandleFunc("/", indexHandler)

    var myport = ":" + os.Getenv("PORT")
    fmt.Println("PORT: ", myport)
    http.ListenAndServe(myport, nil)
}

