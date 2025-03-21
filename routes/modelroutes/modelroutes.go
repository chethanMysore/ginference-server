package modelroutes

import (
	"net/http"
	"strings"
	"time"

	"example/ginference-server/config/devconfig"
	"example/ginference-server/data"
	"example/ginference-server/models/model"
	"example/ginference-server/models/user"
	"example/ginference-server/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// GET (/models)
func GetAllModels(c *gin.Context) {
	filter := bson.D{{}}
	findOptions := options.Find()
	var subscribedModels model.AIModels
	subscribedModels, err := data.Find(subscribedModels, devconfig.DBName, devconfig.ModelCollection, filter, findOptions)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(subscribedModels) == 0 {
		c.IndentedJSON(http.StatusNotFound, subscribedModels.ErrEmptyList().Error())
		return
	}
	c.IndentedJSON(http.StatusOK, subscribedModels)
}

// GET (/models/id/:id)
func GetModelByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, "Please specify the id or use /models for a generic search")
		return
	}
	modelID, parseErr := uuid.Parse(id)
	if parseErr != nil {
		c.IndentedJSON(http.StatusBadRequest, "invalid modelid")
		return
	}
	filter := bson.D{{Key: "modelid", Value: modelID}}
	findOptions := options.Find()
	var subscribedModels model.AIModels
	subscribedModels, err := data.Find(subscribedModels, devconfig.DBName, devconfig.ModelCollection, filter, findOptions)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(subscribedModels) == 0 {
		c.IndentedJSON(http.StatusNotFound, subscribedModels.ErrNotFound(id).Error())
		return
	}
	mod, err := utils.First(subscribedModels)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, mod)
}

// GET (/models/name/:name)
func GetModelByName(c *gin.Context) {
	modelName := c.Param("name")
	if modelName == "" {
		c.IndentedJSON(http.StatusBadRequest, "Please specify the name or use /models for a generic search")
		return
	}
	modelName = strings.Join([]string{".*", modelName, ".*"}, "")
	filter := bson.D{{Key: "modelname", Value: bson.Regex{Pattern: modelName, Options: "i"}}}
	findOptions := options.Find()
	var subscribedModels model.AIModels
	subscribedModels, err := data.Find(subscribedModels, devconfig.DBName, devconfig.ModelCollection, filter, findOptions)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(subscribedModels) == 0 {
		c.IndentedJSON(http.StatusNotFound, subscribedModels.ErrNotFound(modelName).Error())
		return
	}
	c.IndentedJSON(http.StatusOK, subscribedModels)
}

// GET (/models/username/:username)
func GetModelsByUsername(c *gin.Context) {
	userName := c.Param("username")
	if userName == "" {
		c.IndentedJSON(http.StatusBadRequest, "Please specify the name or use /models for a generic search")
		return
	}
	filter := bson.D{{Key: "username", Value: userName}}
	findOptions := options.Find()
	var registeredUsers user.Users
	registeredUsers, err := data.Find(registeredUsers, devconfig.DBName, devconfig.UserCollection, filter, findOptions)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(registeredUsers) == 0 {
		c.IndentedJSON(http.StatusNotFound, registeredUsers.ErrNotFound(userName).Error())
		return
	}
	usr, err := utils.First(registeredUsers)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	filter = bson.D{{Key: "createdby", Value: usr.UserID}}
	var subscribedModels model.AIModels
	subscribedModels, err = data.Find(subscribedModels, devconfig.DBName, devconfig.ModelCollection, filter, findOptions)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(subscribedModels) == 0 {
		c.IndentedJSON(http.StatusNotFound, subscribedModels.ErrNotFound(userName).Error())
		return
	}
	c.IndentedJSON(http.StatusOK, subscribedModels)
}

// POST (/models/create)
func CreateNewModel(c *gin.Context) {
	var newModel model.AIModel
	if err := c.BindJSON(&newModel); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	filter := bson.D{{Key: "userid", Value: newModel.CreatedBy}}
	findOptions := options.Find()
	var registeredUsers user.Users
	registeredUsers, err := data.Find(registeredUsers, devconfig.DBName, devconfig.UserCollection, filter, findOptions)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(registeredUsers) == 0 {
		c.IndentedJSON(http.StatusNotFound, registeredUsers.ErrNotFound(newModel.CreatedBy).Error())
		return
	}
	_, err = utils.First(registeredUsers)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	newModel.ModelID = uuid.New()
	newModel.CreatedAt = time.Now()
	newModel.ModifiedAt = time.Now()
	if err := data.Create(newModel, devconfig.DBName, devconfig.ModelCollection); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, newModel)
}

// PUT (/models/edit)
func EditModel(c *gin.Context) {
	var mod model.AIModelUpdate
	if err := c.BindJSON(&mod); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	filter := bson.D{{Key: "modelid", Value: mod.ModelID}}
	updateOptions := options.UpdateOne().SetUpsert(false)
	mod.ModifiedAt = time.Now()
	if err := data.EditOne(mod, devconfig.DBName, devconfig.ModelCollection, filter, updateOptions); err != nil {
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, mod)
}
