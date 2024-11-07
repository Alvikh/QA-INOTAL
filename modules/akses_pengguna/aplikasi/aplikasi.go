package aplikasi

import "time"

type Aplikasi struct {
	Kd        int16     `json:"kd" gorm:"primaryKey; auto_increment"`
	Nama      string    `json:"nama" gorm:"type:varchar(50);not null"`
	Label     string    `json:"label" gorm:"type:varchar(50);not null"`
	Logo      string    `json:"logo" gorm:"type:varchar(30);not null"`
	UrlFE     string    `json:"url_fe" gorm:"type:varchar(50);not null"`
	UrlAPI    string    `json:"url_api" gorm:"type:varchar(50);not null"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
