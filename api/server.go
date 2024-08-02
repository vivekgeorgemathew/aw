package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vivekgeorgemathew/aw/db/store"
)

type Server struct {
	store  store.Store
	router *gin.Engine
}

func NewServer(store store.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Risk Routes
	router.GET("/api/v1/risks", server.getAllRisks)
	router.GET("/api/v1/risks/:id", server.getRisk)
	router.POST("/api/v1/risks", server.createRisk)

	server.router = router
	return server
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}
