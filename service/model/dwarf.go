package model

func GetAllDwarves() ([]Dwarf, error) {
	var dwarves []Dwarf

	tx := db.Find(&dwarves)

	if tx.Error != nil {
		return []Dwarf{}, tx.Error
	}

	return dwarves, nil
}

func GetDwarf(id uint64) (Dwarf, error) {
	var dwarf Dwarf

	tx := db.Where("id = ?", id).First(&dwarf)

	if tx.Error != nil {
		return Dwarf{}, tx.Error
	}

	return dwarf, nil
}

func CreateDwarf(dwarf Dwarf) (Dwarf, error) {

	tx := db.Create(&dwarf)
	return dwarf, tx.Error
}

func UpdateDwarf(dwarf Dwarf) error {

	tx := db.Save(&dwarf)
	return tx.Error
}

func DeleteDwarf(id uint64) error {

	tx := db.Unscoped().Delete(&Dwarf{}, id)
	return tx.Error
}

func FindByDwarfUrl(url string) (Dwarf, error) {
	var dwarf Dwarf

	tx := db.Where("dwarf = ?", url).First(&dwarf)

	return dwarf, tx.Error
}
