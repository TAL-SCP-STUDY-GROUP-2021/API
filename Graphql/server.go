package main

import (
	"Graphql/app/mysql"
	"Graphql/graph"
	"Graphql/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	ginMiddlewares "github.com/Laisky/gin-middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
)

const defaultPort = "8080"

func main() {
	r := gin.Default()
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	h.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.Options{})
	h.AddTransport(transport.MultipartForm{})
	r.Any("/", ginMiddlewares.FromStd(playground.Handler("GraphQL playground", "/query")))
	r.Any("/query", ginMiddlewares.FromStd(h.ServeHTTP))
	r.GET("/article/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		c.JSON(http.StatusOK, gin.H{
			"article": mysql.CreateArticleDao().Article(id),
		})
	})
	r.GET("/user/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		c.JSON(http.StatusOK, gin.H{
			"user": mysql.CreateUserDao().User(id),
		})
	})
	r.GET("/comments/:id", func(c *gin.Context) {
		comments := make([]*mysql.Comment, 0)
		for _, s := range strings.Split(c.Param("id"), ",") {
			id, _ := strconv.Atoi(s)
			comments = append(comments, mysql.CreateCommentDao().Comment(id))
		}
		c.JSON(http.StatusOK, gin.H{
			"comments": comments,
		})
	})
	r.Run()

}
