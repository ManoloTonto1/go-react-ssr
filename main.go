package main

import (
	"github.com/ManoloTonto1/go-react-ssr/models"
	"github.com/gin-gonic/gin"
	go_ssr "github.com/natewong1313/go-react-ssr"
	"github.com/natewong1313/go-react-ssr/config"
	"github.com/natewong1313/go-react-ssr/react_renderer"
)

var APP_ENV string

func main() {
	g := gin.Default()
	g.StaticFile("favicon.ico", "./frontend/public/favicon.ico")
	g.Static("/assets", "./frontend/public")
	go_ssr.Init(config.Config{
		AppEnv:             APP_ENV,
		AssetRoute:         "/assets",
		FrontendDir:        "./frontend/src",
		GeneratedTypesPath: "./frontend/src/generated.d.ts",
		TailwindConfigPath: "./frontend/tailwind.config.js",
		GlobalCSSFilePath:  "./frontend/src/Main.css",
		PropsStructsPath:   "./models/props.go",
	})

	g.GET("/", func(c *gin.Context) {
		response := react_renderer.RenderRoute(react_renderer.Config{
			File:  "Home.tsx",
			Title: "Manuel's SSR",
			MetaTags: map[string]string{
				"og:title":    "Gin example app",
				"description": "Hello world!",
			},
			Props: &models.IndexRouteProps{
				InitialCount: 0,
				Name:         "Manuel",
			},
		})
		c.Writer.Write(response)
	})
	g.Run()
}
