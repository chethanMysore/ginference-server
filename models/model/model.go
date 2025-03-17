package model

import (
	"example/ginference-server/utils"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type AIModel struct {
	ModelID   uuid.UUID `json:"id"`
	ModelName string    `json:"name"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

func (m AIModel) ErrEmptyList() error {
	return fmt.Errorf("no models subscribed")
}

func (m AIModel) ErrNotFound(params ...any) error {
	if len(params) == 0 {
		return fmt.Errorf("no such model found")
	} else {
		return fmt.Errorf("no such model found with %v", params[0])
	}
}

type AIModels []AIModel

func (m AIModels) ErrNotFound(params ...any) error {
	if len(params) == 0 {
		return fmt.Errorf("no such model found")
	} else {
		return fmt.Errorf("no models found with the given filter criteria - %v", params[0])
	}
}

func (m AIModels) FindByName(modelName string) (AIModel, error) {
	filteredModels, err := utils.Filter(m, func(model AIModel) bool {
		return strings.Contains(strings.ToLower(model.ModelName), strings.ToLower(modelName))
	})
	if err != nil {
		var model AIModel
		return model, model.ErrNotFound(modelName)
	}
	return utils.First(filteredModels)
}

func (m AIModels) FindByUUID(uuid string) (AIModel, error) {
	filteredModels, err := utils.Filter(m, func(model AIModel) bool {
		return strings.Contains(strings.ToLower(model.ModelID.String()), strings.ToLower(uuid))
	})
	if err != nil {
		var model AIModel
		return model, model.ErrNotFound(uuid)
	}
	return utils.First(filteredModels)
}

func (m AIModels) FindByUser(userID string) ([]AIModel, error) {
	filteredModels, err := utils.Filter(m, func(model AIModel) bool {
		return strings.Contains(strings.ToLower(model.CreatedBy), strings.ToLower(userID))
	})
	if err != nil {
		var m AIModels
		return m, m.ErrNotFound(userID)
	}
	return filteredModels, err
}
