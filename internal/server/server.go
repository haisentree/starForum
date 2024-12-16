package server

import (
	"github.com/kataras/iris/v12/mvc"
	"starForum/internal/controllers"
	"starForum/internal/controllers/api"

	"github.com/kataras/iris/v12"
)

type StarForumServer struct {
	server *iris.Application
}

func NewStarForumServer() *StarForumServer {
	app := iris.New()
	return &StarForumServer{
		server: app,
	}
}

func (s *StarForumServer) Init() {
	s.createRouters()
}

func (s *StarForumServer) Start() {
	s.server.Listen(":8080")
}

func (s *StarForumServer) createRouters() {
	booksAPI := s.server.Party("/books")
	{
		booksAPI.Use(iris.Compression)

		// GET: http://localhost:8080/books
		booksAPI.Get("/", controllers.List)
		// POST: http://localhost:8080/books
		booksAPI.Post("/", controllers.Create)
	}

	mvc.Configure(s.server.Party("/api"), func(m *mvc.Application) {
		m.Party("/user").Handle(new(api.UserController))
	})

}
