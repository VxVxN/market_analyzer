package server

import (
	"github.com/VxVxN/log"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func Init() (*Server, error) {
	gin.SetMode(gin.ReleaseMode)
	server := Server{router: gin.Default()}

	if err := log.Init("market_analyzer.log", log.CommonLog, false); err != nil {
		return nil, err
	}

	return &server, nil
}

func (server *Server) ListenAndServe(listen string) error {
	log.Info.Printf("Listening %s", listen)
	if err := server.router.Run(listen); err != nil {
		return err
	}
	return nil
}

func (server *Server) SetRoutes() {
	server.router.POST("/emitters/list", server.emittersListHandler)
	server.router.POST("/emitters/report", server.emittersReportHandler)
	server.router.POST("/emitters/load", server.emittersLoadHandler)
}
