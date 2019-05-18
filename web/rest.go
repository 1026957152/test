package web

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"sync"
)

// Product 商品信息
type Product struct {
	Name      string
	ProductID int64
	Number    int
	Price     float64
	IsOnSale  bool
}

// counter echoes the number of calls so far.
func aboutHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

	data, _ := json.Marshal(status)
	fmt.Fprintf(w, "Form[%q] = %q\n", "Status", string(data))

	//fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	//...
}

var status map[string]string

func Webserver(Wg *sync.WaitGroup, status_ map[string]string) {
	status = status_

	Wg.Add(1)
	var thingName string = "Led"
	var region string = "us-west-2"
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello, ", html.EscapeString(request.URL.Path))
	})
	http.HandleFunc("/about", aboutHandler)
	log.Printf("开启web 服务 $v")

	log.Fatal(http.ListenAndServe(":8080", nil))

	log.Printf("web 服务已经起来")
	p := &Product{}
	p.Name = thingName
	p.Name = region

	p.IsOnSale = true
	p.Number = 10000
	p.Price = 2499.00
	p.ProductID = 1
	data, _ := json.Marshal(p)
	fmt.Println(string(data))
}
func main() {
	//Webserver(nil)
}
