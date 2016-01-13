package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {

/**
  Sets up the web server
  base url: "/" will have Hello World printed on the page
**/
  http.HandleFunc("/",func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello World")
  })

/**
  making a simple get request to google
  Make sure "net/http" is included
**/
  resp, err := http.Get("https://google.com")
  if err != nil {
    //handle error
    fmt.Println("Error")
  }else {
    //print the response
    fmt.Println(resp)
  }

  //running on 8080
  log.Fatal(http.ListenAndServe(":8080",nil))
}
