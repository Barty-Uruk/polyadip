package models

import (
	"fmt"
	"time"
)

type (
	//Ad объявления
	Ad struct {
		ID        int
		Comment   string `form:"comment"`
		Title     string `form:"title"`
		Time      int    `form:"time"`
		UserID    int    `sql:"user_id,notnull,default:0"`
		Category  string `form:"category" sql:"category"`
		Price     int    `form:"price" sql:",notnull,default:0"`
		CreatedAt time.Time
	}
	FilterData struct {
		Title    string `form:"title"`
		Category string `form:"category"`
		MinPrice int    `form:"minprice"`
		MaxPrice int    `form:"maxprice"`
	}
)

//Create создает объявление
func (ad Ad) Create() (Ad, error) {
	ad.CreatedAt = time.Now()
	_, err := db.Model(&ad).Insert()
	if err != nil {
		return ad, fmt.Errorf("Ошибка сохранения объявления в базу,%v", err)
	}
	return ad, nil
}

//Update обновляет объявление
func (ad Ad) Update() (Ad, error) {

	_, err := db.Model(&ad).WherePK().UpdateNotNull()
	if err != nil {
		return ad, fmt.Errorf("Ошибка обновления,%v", err)
	}
	return ad, nil
}

//DeleteAdByID удаляет объявление
func DeleteAdByID(id int) (Ad, error) {
	var (
		ad Ad
	)
	ad.ID = id
	_, err := db.Model(&ad).WherePK().Delete()
	if err != nil {
		return ad, fmt.Errorf("Ошибка удаления,%v", err)
	}
	return ad, nil
}

// GetAdByID возвращает объявление по id
func GetAdByID(id int) (Ad, error) {
	var (
		ad Ad
	)
	err := db.Model(&ad).Where("id = ?", id).Select()
	return ad, err
}

// GetAllAds возвращает все объявления
func GetAllAds() ([]Ad, error) {
	var (
		ads []Ad
	)
	err := db.Model(&ads).Select()
	return ads, err
}

// GetFilterAds возвращает отфильтрованные объявления
func (fd FilterData) GetFilterAds() ([]Ad, error) {
	var (
		ads []Ad
	)
	fd.Title = "%" + fd.Title + "%"
	fd.Category = "%" + fd.Category + "%"
	fmt.Println(fd.MinPrice)
	if fd.MaxPrice == 0 {
		fd.MaxPrice = 999999999
	}

	err := db.Model(&ads).
		Where("title ILIKE ?", fd.Title).
		Where("price <= ?", fd.MaxPrice).
		Where("price >= ?", fd.MinPrice).
		Where("category ILIKE ?", fd.Category).
		Select()
	return ads, err
}

func GetAdsByUserID(userid int) ([]Ad, error) {
	var (
		Ads []Ad
	)
	err := db.Model(&Ads).Where("user_id = ?", userid).Select()
	return Ads, err
}
