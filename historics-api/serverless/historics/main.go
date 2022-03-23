package main

import (
	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	"historics-api/internal"
	"historics-api/internal/repo"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := repo.NewMongo()
	if err != nil {
		log.Fatalln(err)
	}

	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	app, err := internal.Inject(db)
	if err != nil {
		log.Fatalln(err)
	}

	port := ":" + os.Getenv("PORT")

	if gin.Mode() == gin.ReleaseMode {
		log.Fatal(gateway.ListenAndServe(port, app))
	} else {
		log.Fatal(http.ListenAndServe(port, app))
	}

}
