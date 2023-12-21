package app

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/GoLangWebSDK/rest"
	"github.com/GoLangWebSDK/rest/example/controllers"
)

type AppRoutes struct {
}

var _ rest.Routes = (*AppRoutes)(nil)

func NewRoutes() *AppRoutes {
	return &AppRoutes{}
}
func (routes *AppRoutes) LoadRoutes(router *rest.Rest) {
	router.Route("/test_ctrl").Controller(&controllers.TestController{})

	api := router.API("v2")
	api.Route("/users").Controller(&controllers.UsersController{})

	a := router.API("v1")
	a.Route("/articles").Controller(&controllers.ArticlesController{})
}

func (routes *AppRoutes) LoadMiddlewares(router *rest.Rest) {
	// LOGGING MIDDLEWARE
	router.Mux.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

			if r.Method == "POST" || r.Method == "PUT" {
				bodyBytes, err := io.ReadAll(r.Body)
				if err != nil {
					log.Printf("Error reading body: %v", err)
					http.Error(w, "can't read body", http.StatusBadRequest)
					return
				}

				// After reading the body, you need to replace it for further handlers
				r.Body = io.NopCloser(strings.NewReader(string(bodyBytes)))

				log.Printf("Body: %s\n", string(bodyBytes))
			}

			next.ServeHTTP(w, r)
		})
	})
}
