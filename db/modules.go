package db

import "github.com/datvvan/sample/models"

func ModuleFindAll() ([]models.Module, error) {
	instance := GetInstance()
	data := []models.Module{}
	if err := instance.DB.Model(&data).Select(); err != nil {
		return nil, err
	}
	return data, nil
}
