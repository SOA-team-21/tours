package repo

import (
	"gorm.io/gorm"
	"tours.xws.com/model"
)

type PreferenceRepo struct {
	DatabaseConnection *gorm.DB
}

func (repo *PreferenceRepo) Create(preference *model.Preference) error {
	dbResult := repo.DatabaseConnection.Create(preference)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *PreferenceRepo) GetAllByUser(userId string) ([]model.Preference, error) {
	var preferences []model.Preference

	dbResult := repo.DatabaseConnection.Find(&preferences, "user_id = ?", userId)
	if dbResult.Error != nil {
		return preferences, dbResult.Error
	}

	return preferences, nil
}

func (repo *PreferenceRepo) Update(preference *model.Preference) error {
	dbResult := repo.DatabaseConnection.Save(&preference)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("updated successfully")
	return nil
}

func (repo *PreferenceRepo) Delete(preferenceId string) error {
	dbResult := repo.DatabaseConnection.Exec("DELETE FROM preferences WHERE id = ?", preferenceId)

	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}
