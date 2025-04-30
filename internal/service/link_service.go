package service

import (
	"errors"
	"math/big"
	"time"

	"github.com/Mutter0815/short-url-go/internal/models"
	"github.com/Mutter0815/short-url-go/internal/repository"
	"github.com/google/uuid"
)

type LinkService struct {
	repo repository.LinkRepository
}

func NewLinkService(repo repository.LinkRepository) *LinkService {
	return &LinkService{
		repo: repo,
	}
}

func (s *LinkService) GetShortURL() string {
	id := uuid.New()
	num := new(big.Int).SetBytes(id[:5])
	return num.Text(62)

}

func (s *LinkService) CreateShortLink(originalURL string) (string, error) {
	shortUrl := s.GetShortURL()
	link := models.Links{
		Link:       originalURL,
		Short_URL:  shortUrl,
		Created_at: time.Now(),
	}
	err := s.repo.SaveLink(link)
	if err != nil {
		return "", err
	}
	return shortUrl, nil
}
func (s *LinkService) GetOriginalURLByShortLink(shortLink string) (string, error) {
	originalURL, err := s.repo.GetOriginalURL(shortLink)
	if err != nil {
		return "", errors.New("не найдена соответсвующая основная ссылка ссылка")
	}
	return originalURL, nil
}

func (s *LinkService) GetShortLinkByOriginalURL(originalURL string) (string, error) {
	shortUrl, err := s.repo.GetShortLink(originalURL)
	if err != nil {
		return "", errors.New("не найдена соответсвующая основная ссылка ссылка")
	}
	return shortUrl, nil
}
