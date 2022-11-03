package model

func GetShorts() ([]Short, error) {
	var shorts []Short

	tx := db.Find(&shorts)

	if tx.Error != nil {
		return []Short{}, tx.Error
	}

	return shorts, nil
}

func GetShort(id uint64) (Short, error) {
	var short Short

	tx := db.Where("id = ?", id).First(short)

	if tx.Error != nil {
		return Short{}, tx.Error
	}

	return short, nil
}

func GetShortByURL(url string) (Short, error) {
	var short Short

	tx := db.Where("short = ?", url).First(&short)

	return short, tx.Error
}

func CreateShort(short Short) error {
	tx := db.Create(&short)

	return tx.Error
}

func UpdateShort(short Short) error {
	tx := db.Save(&short)

	return tx.Error
}

func DeleteShort(id uint64) error {
	tx := db.Unscoped().Delete(&Short{}, id)

	return tx.Error
}
