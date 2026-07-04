package sync

import "github.com/gin-gonic/gin"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) SyncStudents(c *gin.Context) {

}

func (h *Handler) SyncProformas(c *gin.Context) {

}
