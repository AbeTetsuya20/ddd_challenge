package handler

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Service interface {
	Server(ctx context.Context)
}

type ServiceController struct {
	// now 現在時刻を取得するための関数
	Now    func() time.Time
	Driver *ServiceDriver
}

func NewServiceController(service *ServiceDriver) *ServiceController {
	return &ServiceController{
		Now:    time.Now,
		Driver: service,
	}
}

func (s *ServiceController) Server(ctx context.Context) {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/debug", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"message": "hello",
		}
		render.JSON(w, r, data)
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/message", func(r chi.Router) {
			r.Route("/get", func(r chi.Router) {
				r.Get("/send", func(w http.ResponseWriter, r *http.Request) {
					s.Driver.MessageGetSend(ctx, w, r)
				})

				r.Get("/notsend", func(w http.ResponseWriter, r *http.Request) {
					data := map[string]string{
						"message": "notsend",
					}
					render.JSON(w, r, data)
				})
			})
		})

		r.Route("/channel", func(r chi.Router) {
			r.Get("/get", func(w http.ResponseWriter, r *http.Request) {
				data := map[string]string{
					"message": "get",
				}
				render.JSON(w, r, data)
			})

			r.Get("/create", func(w http.ResponseWriter, r *http.Request) {
				data := map[string]string{
					"message": "create",
				}
				render.JSON(w, r, data)
			})
		})

		r.Route("/user", func(r chi.Router) {
			r.Get("/get", func(w http.ResponseWriter, r *http.Request) {
				data := map[string]string{
					"message": "/user/get",
				}
				render.JSON(w, r, data)
			})
		})

		r.Route("/join", func(r chi.Router) {
			r.Get("/delete", func(w http.ResponseWriter, r *http.Request) {
				data := map[string]string{
					"message": "delete",
				}
				render.JSON(w, r, data)
			})

			r.Get("/create", func(w http.ResponseWriter, r *http.Request) {
				data := map[string]string{
					"message": "create",
				}
				render.JSON(w, r, data)
			})

			r.Route("/get", func(r chi.Router) {
				r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
					data := map[string]string{
						"message": "/join/get/user",
					}
					render.JSON(w, r, data)
				})

				r.Get("/channel", func(w http.ResponseWriter, r *http.Request) {
					data := map[string]string{
						"message": "channel",
					}
					render.JSON(w, r, data)
				})
			})
		})
	})

	addr := os.Getenv("Addr")
	if addr == "" {
		addr = ":8080"
	}

	log.Printf("listen: %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("!! %+v", err)
	}
}
