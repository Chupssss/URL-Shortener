package handler

import (
	"URL-Shortener/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UrlHandler struct {
	service *service.UrlServ
	baseURL string
}

type createUrlRequest struct {
	OriginalUrl string `json:"original_url"`
}

func NewUrlHandler(service *service.UrlServ, baseURL string) *UrlHandler {
	return &UrlHandler{service: service,
		baseURL: baseURL}
}

func (h *UrlHandler) Create(c *gin.Context) {
	var request createUrlRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	shortCode := h.service.Create(request.OriginalUrl)

	c.JSON(http.StatusCreated, gin.H{
		"short_url": h.baseURL + "/" + shortCode,
	})
}

func (h *UrlHandler) RedirectURL(c *gin.Context) {
	shortCode := c.Param("shortCode")
	originalUrl, ok := h.service.GetOriginalURL(shortCode)
	if ok {
		c.Redirect(http.StatusFound, originalUrl)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "invalid request"})
	return
}
