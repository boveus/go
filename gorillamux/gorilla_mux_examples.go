package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HandlerFunc is similar to http.HandleFunc
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)

func variablesinpath() {
	// You can also add variables to a path like so
	r := mux.NewRouter()
	r.HandleFunc("/products/{key}", ProductHandler)
	r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
}
func subroutes()   {
  //Subroutes to remove repitition for routes sharing a subroute:
  r := mux.NewRouter()
  s := r.Host("www.example.com").Subrouter()
  // Then register routes in the subrouter:

  s.HandleFunc("/products/", ProductsHandler)
  s.HandleFunc("/products/{key}", ProductHandler)
  s.HandleFunc("/articles/{category}/{id:[0-9]+}"), ArticleHandler)
}

func statichandlers() {
    flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
    flag.Parse()
    r := mux.NewRouter()

    // This will serve files under http://localhost:8000/static/<filename>
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

    srv := &http.Server{
        Handler:      r,
        Addr:         "127.0.0.1:8000",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())
}


}
