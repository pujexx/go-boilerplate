package gen

import (
	"fmt"
	"github.com/gorilla/mux"
	"strings"
)

type postManGenerator struct {
	R *mux.Router
}

func NewPostmanGenerator(r *mux.Router) *postManGenerator {
	return &postManGenerator{R: r}
}

func (p *postManGenerator) Generate() {

	p.R.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {

		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		name := route.GetName()
		fmt.Println("name : ", name)

		return nil
	})
}
