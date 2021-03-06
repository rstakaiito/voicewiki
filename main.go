package main

import (
	"fmt"
	"github.com/tamnd/voicewiki/api"
	"github.com/tamnd/voicewiki/handler"
	"github.com/tamnd/voicewiki/middleware"
	"github.com/tamnd/voicewiki/model"
	"os"
)

func main() {
	var configFile string = "app.config"
	if len(os.Args) >= 2 {
		configFile = os.Args[1]
	}
	err := middleware.LoadConfig(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	model.Init()

	middleware.Route("/article/:id", handler.Article)
	middleware.Route("/list", handler.List)
	middleware.Route("/", handler.Home)

	middleware.Route("/api/article/:id", api.GetArticle)
	middleware.Route("/api/search", api.SearchArticle)
	middleware.Route("/api/list", api.ListArticle)
	middleware.Route("/api/read", api.ReadSection)

	middleware.Run()
}
