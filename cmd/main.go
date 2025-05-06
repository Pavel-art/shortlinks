package main

import (
	"fmt"
	"net/http"
	"shortlinks/configs"
	"shortlinks/internal/auth"
	"shortlinks/internal/link"
	"shortlinks/internal/user"
	"shortlinks/pkg/db"
	"shortlinks/pkg/middleware"
)

func main() {

	// create dependencies
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositories init
	linkRepository := link.NewLinkRepository(database)
	userRepository := user.NewUserRepository(database)

	// Services
	authService := auth.NewAuthService(userRepository)

	//Handlers init
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
	})

	//Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}
	fmt.Println("Starting server on port 8081")
	server.ListenAndServe()

}
