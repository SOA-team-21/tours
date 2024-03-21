package service

import (
	"tours.xws.com/model"
	"tours.xws.com/repo"
)

type PreferenceService struct {
	PreferenceRepo *repo.PreferenceRepo
}

func (service *PreferenceService) Create(preference *model.Preference) error {
	err := service.PreferenceRepo.Create(preference)

	if err != nil {
		return err
	}

	return nil
}

func (service *PreferenceService) Update(preference *model.Preference) error {
	err := service.PreferenceRepo.Update(preference)

	if err != nil {
		return err
	}

	return nil
}

func (service *PreferenceService) Delete(preferenceId string) error {
	err := service.PreferenceRepo.Delete(preferenceId)

	if err != nil {
		return err
	}

	return nil
}

func (service *PreferenceService) GetAllByUser(userId string) ([]model.Preference, error) {
	var preferences []model.Preference

	preferences, err := service.PreferenceRepo.GetAllByUser(userId)
	if err != nil {
		return nil, err
	}

	return preferences, nil
}
