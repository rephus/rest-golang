package main

import (
    "encoding/json"
    "fmt"
    "html"
    "log"
    "net/http"
    "time"
    "io/ioutil"
)

type Response struct {
    Name      string `json:"name"`
    Params    string `jsong:"params"`
    Completed bool `json:"completed"`
    Time      time.Time `json:"time"`
}

type PostBody struct {
    Foo   string  `json:"foo"`
}

func main() {
    log.Println("Server started in http://localhost:8080")
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    //curl -X POST  http://localhost:8080/testPOST -d "{\"foo\": \"bar\"}" 
    http.HandleFunc("/testPOST", func(w http.ResponseWriter, r *http.Request) {
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {  panic(err) }
        var t PostBody
        err = json.Unmarshal(body, &t)
        if err != nil { panic(err)}

        params := t.Foo
        response := Response{Name: "Post", Params: params, Completed: true, Time: time.Now()}
        json.NewEncoder(w).Encode(response)
    })

    //http://localhost:8080/testGET?foo=bar
    http.HandleFunc("/testGET", func(w http.ResponseWriter, r *http.Request) {
        params := r.URL.Query().Get("foo")
        response := Response{Name: "Get", Params: params, Completed: true, Time: time.Now()}
        json.NewEncoder(w).Encode(response)
    })
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}
