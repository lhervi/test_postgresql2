package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Repo *Postgresrepo
}

func NewHandlers(repo *Postgresrepo) *Handlers {
	return &Handlers{
		Repo: repo,
	}
}

func (h *Handlers) InitRoutes(router *gin.Engine) {
	router.GET("/all", h.getAll)
}

func (h *Handlers) getAll(c *gin.Context) {
	users, err := h.Repo.GetAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": users,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": users,
	})
}
