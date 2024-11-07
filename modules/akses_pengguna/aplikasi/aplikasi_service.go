package aplikasi

import (
	"errors"
	"log"
	"rsudlampung/helper"
	"time"

	"gorm.io/gorm"
)

// AplikasiService interface with updated method signature
type AplikasiService interface {
	FindAll() ([]Aplikasi, error)
	FindByKd(kd int16) (Aplikasi, error)
	FindByLimit(limit, offset int) ([]Aplikasi, error)
	Create(aplikasi Aplikasi) (Aplikasi, error)
	Update(aplikasi Aplikasi) error
	Delete(kd int16) error
}

type aplikasiService struct {
	conn *gorm.DB
}

func NewAplikasiService(db *gorm.DB) AplikasiService {
	configEnv, errEnv := helper.LoadConfig("../.")
	if errEnv != nil {
		log.Fatal("cannot load config:", errEnv)
	}
	am := configEnv.AutoMigrate

	if am == "on" {
		if err := db.AutoMigrate(&Aplikasi{}); err != nil {
			log.Fatal("failed to auto-migrate Aplikasi model:", err)
		}
	}

	return &aplikasiService{
		conn: db,
	}
}

// Create creates a new aplikasi record
func (service *aplikasiService) Create(aplikasi Aplikasi) (Aplikasi, error) {
	if aplikasi.Nama == "" || aplikasi.Label == "" || aplikasi.Logo == "" || aplikasi.UrlFE == "" || aplikasi.UrlAPI == "" {
		return Aplikasi{}, errors.New("semua field wajib diisi")
	}
	aplikasi.CreatedAt = time.Now()
	aplikasi.UpdatedAt = time.Now()

	if err := service.conn.Create(&aplikasi).Error; err != nil {
		return Aplikasi{}, err
	}
	return aplikasi, nil
}

// Update updates an existing aplikasi record
func (service *aplikasiService) Update(aplikasi Aplikasi) error {
	if aplikasi.Nama == "" || aplikasi.Label == "" || aplikasi.Logo == "" || aplikasi.UrlFE == "" || aplikasi.UrlAPI == "" {
		return errors.New("semua field wajib diisi")
	}

	return service.conn.Transaction(func(tx *gorm.DB) error {
		var existingAplikasi Aplikasi
		if err := tx.First(&existingAplikasi, "kd = ?", aplikasi.Kd).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("Aplikasi tidak ditemukan untuk diupdate")
			}
			return err
		}

		aplikasi.UpdatedAt = time.Now()
		if err := tx.Model(&existingAplikasi).Omit("created_at").Updates(aplikasi).Error; err != nil {
			return err
		}

		return nil
	})
}

// Delete deletes an aplikasi record
func (service *aplikasiService) Delete(kd int16) error {
	return service.conn.Transaction(func(tx *gorm.DB) error {
		var aplikasi Aplikasi
		if err := tx.First(&aplikasi, "kd = ?", kd).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("Aplikasi tidak ditemukan untuk dihapus")
			}
			return err
		}

		if err := tx.Delete(&aplikasi).Error; err != nil {
			return err
		}
		return nil
	})
}

// FindAll returns all aplikasi records
func (service *aplikasiService) FindAll() ([]Aplikasi, error) {
	var aplikasis []Aplikasi
	if err := service.conn.Find(&aplikasis).Error; err != nil {
		return nil, err
	}
	return aplikasis, nil
}

// FindByKd returns an aplikasi by kd
func (service *aplikasiService) FindByKd(kd int16) (Aplikasi, error) {
	var aplikasi Aplikasi
	if err := service.conn.First(&aplikasi, "kd = ?", kd).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Aplikasi{}, errors.New("Aplikasi tidak ditemukan")
		}
		return Aplikasi{}, err
	}
	return aplikasi, nil
}

// FindByLimit retrieves aplikasi records with offset and limit
func (service *aplikasiService) FindByLimit(limit, offset int) ([]Aplikasi, error) {
	var aplikasiList []Aplikasi
	if err := service.conn.Offset(offset).Limit(limit).Find(&aplikasiList).Error; err != nil {
		return nil, err
	}
	return aplikasiList, nil
}

