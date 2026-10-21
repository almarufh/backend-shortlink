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

type Links struct {
	LinksID     int64  `json:"id" db:"id"`
	OriginalURL string `json:"original_url" db:"original_url"`
	ShortURL    string `json:"short_url" db:"short_url"`
	Slug        string `json:"slug" db:"slug"`
}

type MyLinksResponse struct {
	Data []Links `json:"data"`
}

type MyLinksRequest struct {
	UserID   int64  `form:"-"`
	ShortURL string `form:"short_url"`
}
