package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/Mutter0815/short-url-go/internal/service"
	"github.com/Mutter0815/short-url-go/internal/transport/http/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LinkHandler struct {
	linkService *service.LinkService
}

func NewLinkHandler(linkService *service.LinkService) *LinkHandler {
	return &LinkHandler{linkService: linkService}
}

func (h *LinkHandler) CreateLink(c *gin.Context) {
	request := dto.CreateLinkRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Ошибка привязки данных:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный запрос"})
		return
	}
	if !strings.HasPrefix(request.Link, "http://") && !strings.HasPrefix(request.Link, "https://") {
		request.Link = "https://" + request.Link
	}
	validate := validator.New()
	if err := validate.Struct(&request); err != nil {
		log.Printf("Ошибка валидации:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}
	shortLink, err := h.linkService.CreateShortLink(request.Link)
	if err != nil {
		log.Println("Ошибка при создании короткой ссылки", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Неккоректный запрос"})
		return
	}
	shortLink = c.Request.Host + c.Request.URL.Port() + "/" + shortLink
	c.JSON(http.StatusOK, gin.H{"shortLink": shortLink})
}

func (h *LinkHandler) RedirectLink(c *gin.Context) {
	shortLink := c.Param("short_code")
	log.Printf("Попытка найти оригинальную ссылку:%s", shortLink)
	originalURL, err := h.linkService.GetOriginalURLByShortLink(shortLink)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ссылка не найдена"})
		return
	}
	if !strings.HasPrefix(originalURL, "http") {
		originalURL = "https://" + originalURL
	}
	c.Redirect(http.StatusMovedPermanently, originalURL)

}

func (h *LinkHandler) GetOriginalURLByShortLink(c *gin.Context) {
	shortLink := c.Query("shortlink")
	originalURL, err := h.linkService.GetOriginalURLByShortLink(shortLink)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "не найдена ссылка оригинала"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"originalURL": originalURL})

}

func (h *LinkHandler) GetShortLinkByOriginalLink(c *gin.Context) {
	originalURL := c.Query("originalurl")
	log.Printf("Запрос на короткую ссылку для: %s", originalURL) // Добавь логирование для проверки
	shortLink, err := h.linkService.GetShortLinkByOriginalURL(originalURL)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "не найдена короткая ссылка"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"shortLink": shortLink})

}
