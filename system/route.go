package system

import (
	"gozealous/nexus"

	"github.com/gin-gonic/gin"
)

func ConfigureRoute(engine *gin.Engine, nexus *nexus.Store) {
	engine.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Welcome to gozealous")
	})

	health := engine.Group("/health")
	health.GET("/", nexus.SystemHealth())
	if gin.IsDebugging() {
		health.GET("/database", nexus.DatabaseHealth())
	}

	railway := engine.Group("/railway")
	railway.GET("/stations", nexus.RailwayStations())
	railway.GET("/lines", nexus.RailwayLines())
	railway.GET("/journey", nexus.RailwayJourney())
}
