package dto

type LinksRequest struct {
	UserID      int32   `json:"-"`
	OriginalURL string  `json:"original_url" validate:"required,url" example:"https://example.com/long-url-path"`
	ShortURL    *string `json:"short_url,omitempty" validate:"omitempty,min=3,max=20,alphanum" example:"my-custom-slug"`
	IsCustom    *bool   `json:"is_custom,omitempty" example:"false"`
}

type LinksResponse struct {
	OriginalURL string  `json:"original_url" example:"https://example.com/long-url-path"`
	ShortURL    *string `json:"short_url,omitempty" example:"x7K9aL"`
}
