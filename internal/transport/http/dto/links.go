package dto

type CreateLinkRequest struct {
	Link string `json:"link" validate:"required,url"`
}

type CreateLinkResponse struct {
	ID       uint   `json:"id"`
	Link     string `json:"link"`
	ShortURL string `json:"short_url"`
}
