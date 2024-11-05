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
	FindByKd(kd int) Aplikasi
	FindByLimit(offset int, limit int) []Aplikasi
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

func (service *aplikasiService) Create(aplikasi Aplikasi) (Aplikasi, error) {
	err := service.conn.Create(&aplikasi).Error
	if err != nil {
		return Aplikasi{}, err
	}
	return aplikasi, nil
}

func (service *aplikasiService) Update(aplikasi Aplikasi) error {
	err := service.conn.Save(&aplikasi).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *aplikasiService) Delete(aplikasi Aplikasi) error {
	err := service.conn.Delete(&aplikasi).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *aplikasiService) FindAll() []Aplikasi {
	var aplikasis []Aplikasi
	service.conn.Find(&aplikasis)
	return aplikasis
}

func (service *aplikasiService) FindByKd(kd int) Aplikasi {
	var aplikasi Aplikasi
	service.conn.Where("kd=?", kd).Find(&aplikasi)
	return aplikasi
}

func (service *aplikasiService) FindByLimit(offset int, limit int) []Aplikasi {
	var aplikasis []Aplikasi
	service.conn.Offset(offset).Limit(limit).Find(&aplikasis)
	return aplikasis
}
