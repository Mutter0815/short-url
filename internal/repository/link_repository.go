package repository

import (
	"database/sql"
	"errors"

	"github.com/Mutter0815/short-url-go/internal/models"
)

type LinkRepository struct {
	db *sql.DB
}

func NewLinkRepository(db *sql.DB) *LinkRepository {
	return &LinkRepository{db: db}
}

func (r *LinkRepository) SaveLink(link models.Links) error {
	_, err := r.db.Exec("INSERT INTO links (link,short_url,created_at) VALUES ($1,$2,$3)", link.Link, link.Short_URL, link.Created_at)
	return err
}

func (r *LinkRepository) GetOriginalURL(shortLink string) (string, error) {
	var originalURL string
	err := r.db.QueryRow("SELECT link FROM links WHERE short_url=$1", shortLink).Scan(&originalURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("ссылка не найдена")
		}
		return "", err
	}

	return originalURL, err
}

func (r *LinkRepository) GetShortLink(originalURL string) (string, error) {
	var shortUrl string
	err := r.db.QueryRow("SELECT short_url FROM links WHERE link=$1", originalURL).Scan(&shortUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("ссылка не найдена")
		}
		return "", err
	}

	return shortUrl, err
}
