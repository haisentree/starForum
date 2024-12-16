package main

import "starForum/internal/server"

func main() {
	s := server.NewStarForumServer()
	s.Init()
	s.Start()
}
