package server

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"starForum/internal/controllers/api"
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
	mvc.Configure(s.server.Party("/api"), func(m *mvc.Application) {
		m.Party("/user").Handle(new(api.UserController))
		m.Party("/captcha").Handle(new(api.CaptchaController))
	})
}
