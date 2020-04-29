/*Ejercicio: web crawler
En este ejercicio usarás la concurrencia de Go 
con el fin de paralelizar un web crawler 
(que obtiene recursivamente páginas web a partir 
de una URL).

Modifica la función Crawl para cargar las URLs en
 paralelo evitando repeticiones.
*/
package main

import (
	"fmt"
)

type Fetcher interface {
   // Fetch devuelve el cuerpo de una URL y 
   // un slice de las URLs encontradas en la página
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl usa el 'fetcher' para obtener recursivamente las páginas
// empezando en cierta 'url', con cierta profundidad máxima 'depth'.
func Crawl(url string, depth int, fetcher Fetcher) {
   // TODO: Obtener URLs en paralelo.
	// TODO: Evitar repeticiones.
	// Esta implementación no hace ninguna de las dos cosas.
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}


// fakeFetcher es un Fetcher the devuelve resultados ficticios.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls     []string
}

func (f *fakeFetcher) Fetch(url string) (str---ing, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher contiene los datos ficticios.
var fetcher = &fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}