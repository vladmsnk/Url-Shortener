package v1

import (
	"net/http"
	"vladmsnk/taskrec/internal/dto"

	"github.com/gin-gonic/gin"

	"vladmsnk/taskrec/internal/usecase"
	"vladmsnk/taskrec/pkg/logger"
)

type selectionRoutes struct {
	t usecase.Selection
	l logger.Interface
}

func newTranslationRoutes(handler *gin.RouterGroup, t usecase.Selection, l logger.Interface) {
	r := &selectionRoutes{t, l}

	h := handler.Group("/")
	{
		h.POST("/activity", r.postActivity)
		h.GET("selection")
	}
}

func (r *selectionRoutes) postActivity(c *gin.Context) {
	var activity dto.PostActivityRequest

	if err := c.ShouldBindJSON(&activity); err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	err := r.t.PostActivity(c.Request.Context(), activity)
	if err != nil {
		r.l.Error(err, "http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "database problems")
		return
	}

	c.JSON(http.StatusOK, activity)
}

func (r *selectionRoutes) getSelection(c *gin.Context) {
	var selection dto.GetSelectionResponse

	c.JSON(http.StatusOK, selection)
}