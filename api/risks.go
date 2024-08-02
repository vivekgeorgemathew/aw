package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vivekgeorgemathew/aw/db/models"
	"net/http"
)

type CreateRiskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	State       string `json:"state" binding:"required,oneof=open closed accepted investigating"`
}

// Get a single risk
func (server *Server) getRisk(ctx *gin.Context) {
	riskId := ctx.Param("id")
	risk, err := server.store.Get(riskId)
	if err != nil {
		if risk == nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, risk)
}

// Get all risks
func (server *Server) getAllRisks(ctx *gin.Context) {
	risks, err := server.store.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, risks)
}

// Create Risk
func (server *Server) createRisk(ctx *gin.Context) {
	var createAccReq CreateRiskRequest
	if err := ctx.ShouldBindJSON(&createAccReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	riskId := uuid.New().String()
	riskRecord := models.Risk{}
	riskRecord.RiskID = riskId
	riskRecord.State = createAccReq.State
	riskRecord.Title = createAccReq.Title
	riskRecord.Description = createAccReq.Description

	if err := server.store.Save(riskId, riskRecord); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, riskRecord)
}
