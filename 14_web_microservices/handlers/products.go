package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/MFCaballero/go_crash_course/14_web_microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URL more than one id")
			http.Error(w, "Invalid URL", http.StatusBadRequest)
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URL more than one capture group")
			http.Error(w, "Invalid URL", http.StatusBadRequest)
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			p.l.Println("Invalid URL unable to convert to number", idString)
			http.Error(w, "Invalid URL", http.StatusBadRequest)
		}

		p.updateProducts(id, w, r)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJson(w)
	if err != nil {
		http.Error(w, "Unable to parse json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p *Products) updateProducts(id int, w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
	}
}
