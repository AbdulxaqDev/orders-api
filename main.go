package main

import (
    "fmt"
    "net/http"
)



func main(){
    server := &http.Server{
        Addr: ":9000",
        Handler: http.HandlerFunc(basicHandler),
    }

    if err := server.ListenAndServe(); err!=nil {
        fmt.Println("Failed to listen server:", err)
    }

}

func basicHandler(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Hell, world")) 
}
