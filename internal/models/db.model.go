package models

import "time"

type Users struct {
	ID         int32      `db:"id"`
	Name       string     `db:"name"`
	ImageURL   *string    `db:"image_url"`
	Email      string     `db:"email"`
	Password   string     `db:"password"`
	VerifirdAt *time.Time `db:"verified_at"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
}

type Links struct {
	ID          int64      `db:"id"`
	UserID      int32      `db:"user_id"`
	OriginalURL string     `db:"original_url"`
	ShortURL    string     `db:"short_url"`
	IsCustom    bool       `db:"is_custom"`
	Password    string     `db:"password"`
	ClickCount  int64      `db:"click_count"`
	DeletedAt   *time.Time `db:"deleted_at"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
}
