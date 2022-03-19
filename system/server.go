package system

import (
	"gozealous/environment"
	"gozealous/route"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func (s *Server) Run() {
	s.engine = gin.Default()

	route.Configure(s.engine)

	port, err := environment.ServerPort()

	if err == nil {
		s.engine.Run(":" + port)
	} else {
		panic(err)
	}
}
