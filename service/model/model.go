package model

import (
	"dwarf/database"
)

type Dwarf struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"not null"`
	Dwarf    string `json:"dwarf" gorm:"uniquel; not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

func GetAllDwarves() ([]Dwarf, error) {
	var dwarves []Dwarf
	db := database.DB
	tx := db.Find(&dwarves)

	if tx.Error != nil {
		return []Dwarf{}, tx.Error
	}

	return dwarves, nil
}

func GetDwarf(id uint64) (Dwarf, error) {
	var dwarf Dwarf
	db := database.DB
	tx := db.Where("id = ?", id).First(&dwarf)

	if tx.Error != nil {
		return Dwarf{}, tx.Error
	}

	return dwarf, nil
}

func CreateDwarf(dwarf Dwarf) (Dwarf, error) {
	db := database.DB
	tx := db.Create(&dwarf)
	return dwarf, tx.Error
}

func UpdateDwarf(dwarf Dwarf) error {
	db := database.DB
	tx := db.Save(&dwarf)
	return tx.Error
}

func DeleteDwarf(id uint64) error {
	db := database.DB
	tx := db.Unscoped().Delete(&Dwarf{}, id)
	return tx.Error
}

func FindByDwarfUrl(url string) (Dwarf, error) {
	var dwarf Dwarf
	db := database.DB
	tx := db.Where("dwarf = ?", url).First(&dwarf)

	return dwarf, tx.Error
}
