package main

import (
    "fmt"
    "net/http"
    
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"

)



func main(){

    router := chi.NewRouter()
    router.Use(middleware.Logger)

    router.Get("/hello", basicHandler)

    server := &http.Server{
        Addr: ":9000",
        Handler: router,
    }

    if err := server.ListenAndServe(); err!=nil {
        fmt.Println("Failed to listen server:", err)
    }

}

func basicHandler(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Hell, world")) 
}
