package aplikasi

type Aplikasi struct {
	Kd     int16  `json:"kd" gorm:"primaryKey;autoIncrement"`
	Nama   string `json:"nama" gorm:"type:varchar(50);not null"`
	Label  string `json:"label" gorm:"type:varchar(50);not null"`
	Logo   string `json:"logo" gorm:"type:varchar(30);not null"`
	UrlFE  string `json:"url_fe" gorm:"type:varchar(50);not null"`
	UrlAPI string `json:"url_api" gorm:"type:varchar(50);not null"`
}
