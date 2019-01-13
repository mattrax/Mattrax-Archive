package main

import (
	"context"
	"net/http"

	gqlhandler "github.com/99designs/gqlgen/handler"
	"github.com/kataras/muxie"
	"github.com/mattrax/Mattrax/cmd/mattrax/assets"
	"github.com/mattrax/Mattrax/internal/graphql"
	"github.com/oscartbeaumont/go-utils/handler"
	"github.com/oscartbeaumont/go-utils/middleware"
	"github.com/rs/zerolog/log"
	"github.com/shurcooL/httpgzip"
)

// TODO: Request Logging Middleware
// TODO: Do HTTP/2 With The React Asssets
// TODO: Look. Might have good idea's -> https://github.com/strukturag/httputils
// TODO: Only Serve the the stuff via GET minus GraphQL
// TODO: Go Doc Funcs

func routes(mux *muxie.Mux) {
	mux.Use(middleware.Headers("Mattrax"))

	/*		Frontend		*/
	mux.HandleFunc("/", handler.ServeFile("/index.html"))
	mux.HandleFunc("/login", handler.ServeFile("/index.html"))
	mux.HandleFunc("/favicon.ico", handler.ServeFile("/favicon.ico"))

	/*		Frontend Static Files (Js, Css, etc)		*/
	static := mux.Of("/static")
	static.Use(middleware.CacheHeader("max-age=31536000")) // Cache for 1 Year
	static.Handle("/*", httpgzip.FileServer(assets.Assets, httpgzip.FileServerOptions{IndexHTML: true}))

	/*		Custom Handlers		*/

	/*		GraphQL And 404 Handler		*/
	if *flgDebug {
		mux.Handle("/playground", gqlhandler.Playground("Mattrax GraphQL playground", "/query"))
	}
	mux.Handle("/query", gqlhandler.GraphQL(
		graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{}}),
		gqlhandler.IntrospectionEnabled(*flgDebug),
		gqlhandler.ComplexityLimit(5),
		gqlhandler.RecoverFunc(func(ctx context.Context, err interface{}) error {
			log.Error().Interface("err", err).Msg("GraphQL resolver error")
			return nil
		}),
	))
	mux.HandleFunc("/*path", notFoundHandler())
}

func notFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound) // TODO: Gives 'multiple response.WriteHeader calls' Make style one here + in React that look the same
		handler.ServeFile("/404/index.html")(w, r)
	}
}
