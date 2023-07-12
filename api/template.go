package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kierquebs/capcut-scraper/pkg/scrape"
	"github.com/kierquebs/capcut-scraper/util"
)

type getRenderDataRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type getRenderDataResponse struct {
	RenderData interface{}
}

func (server *Server) getRenderData(ctx *gin.Context) {
	var req getRenderDataRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	data, err := scrape.GetRenderData(server.config,req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}

	json, err := util.DataDecoder(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, json)
}

func (server *Server) getRenderDetail(ctx *gin.Context) {
	var req getRenderDataRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	data, err := scrape.GetRenderDetail(server.config,req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}

	json, err := util.DataDecoder(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, json)
}