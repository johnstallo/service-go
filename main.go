package main

import(
"fmt"
"net/http"
"os"
)

func indexHandler( w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello JOHN from service B!!")
}

func main(){
    http.HandleFunc("/", indexHandler)

    var myport = ":" + os.Getenv("PORT")
    fmt.Println("PORT: ", myport)
    http.ListenAndServe(myport, nil)
}

