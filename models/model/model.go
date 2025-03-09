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

func (m AIModel) ErrEmptyList() string {
	return "No models subscribed!"
}

func (m AIModel) ErrNotFound(params ...any) string {
	if len(params) == 0 {
		return "No such model found!"
	} else {
		return fmt.Sprintf("No such model found with %v", params[0])
	}
}

type AIModels []AIModel

func (m AIModels) ErrNotFound(params ...any) string {
	if len(params) == 0 {
		return "No such model found!"
	} else {
		return fmt.Sprintf("No models found with the given filter criteria - %v", params[0])
	}
}

func (m AIModels) FindByName(modelName string) (AIModel, string) {
	filteredModels, err := utils.Filter(m, func(model AIModel) bool {
		return strings.Contains(strings.ToLower(model.ModelName), strings.ToLower(modelName))
	})
	if err == "" {
		return utils.First(filteredModels)
	} else {
		var model AIModel
		return model, model.ErrNotFound(modelName)
	}
}

func (m AIModels) FindByUUID(uuid string) (AIModel, string) {
	filteredModels, err := utils.Filter(m, func(model AIModel) bool {
		return strings.Contains(strings.ToLower(model.ModelID.String()), strings.ToLower(uuid))
	})
	if err == "" {
		return utils.First(filteredModels)
	} else {
		var model AIModel
		return model, model.ErrNotFound(uuid)
	}
}

func (m AIModels) FindByUser(userID string) ([]AIModel, string) {
	filteredModels, err := utils.Filter(m, func(model AIModel) bool {
		return strings.Contains(strings.ToLower(model.CreatedBy), strings.ToLower(userID))
	})
	if err == "" {
		return filteredModels, err
	} else {
		var m AIModels
		return m, m.ErrNotFound(userID)
	}
}
