package aplikasi

import (
	"log"
	"rsudlampung/helper"

	"gorm.io/gorm"
)

type AplikasiService interface {
	Create(Aplikasi) (Aplikasi, error)
	Update(Aplikasi) error
	Delete(Aplikasi) error
	FindAll() []Aplikasi
	FindByKd(kd string) (Aplikasi, error)
	FindByLimit(limit int) ([]Aplikasi, error)
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
		db.AutoMigrate(&Aplikasi{})
	}

	return &aplikasiService{
		conn: db,
	}
}

// Create creates a new aplikasi record
func (service *aplikasiService) Create(aplikasi Aplikasi) (Aplikasi, error) {
	err := service.conn.Create(&aplikasi).Error
	if err != nil {
		return Aplikasi{}, err
	}
	return aplikasi, nil
}

// Update updates an existing aplikasi record
func (service *aplikasiService) Update(aplikasi Aplikasi) error {
	err := service.conn.Save(&aplikasi).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes an aplikasi record
func (service *aplikasiService) Delete(aplikasi Aplikasi) error {
	err := service.conn.Delete(&aplikasi).Error
	if err != nil {
		return err
	}
	return nil
}

// FindAll returns all aplikasi records
func (service *aplikasiService) FindAll() []Aplikasi {
	var aplikasis []Aplikasi
	service.conn.Find(&aplikasis)
	return aplikasis
}

// FindByKd returns an aplikasi by kd
func (service *aplikasiService) FindByKd(kd string) (Aplikasi, error) {
	var aplikasi Aplikasi
	err := service.conn.Where("kd = ?", kd).First(&aplikasi).Error
	if err != nil {
		return Aplikasi{}, err
	}
	return aplikasi, nil
}

// FindByLimit returns a limited number of aplikasi records
func (service *aplikasiService) FindByLimit(limit int) ([]Aplikasi, error) {
	var aplikasis []Aplikasi
	err := service.conn.Limit(limit).Find(&aplikasis).Error
	if err != nil {
		return nil, err
	}
	return aplikasis, nil
}
