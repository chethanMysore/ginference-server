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

// @BasePath /api/v1

// PingExample godoc
// @Summary Get all models
// @Schemes
// @Description Find all AI Models subscribed to the ginference-server
// @Tags AIModels
// @Accept json
// @Produce json
// @Success 200 {array} model.AIModel
// @Router /models [get]
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

// @BasePath /api/v1

// PingExample godoc
// @Summary Search model by modelID
// @Schemes
// @Description Find an AI Model using the given modelID
// @Tags AIModels
// @Param id path string true "Model ID" minlength(36) maxlength(36)
// @Accept json
// @Produce json
// @Success 200 {object} model.AIModel
// @Router /models/id/{id} [get]
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

// @BasePath /api/v1

// PingExample godoc
// @Summary Search model by model name
// @Schemes
// @Description Find AI Models matching the given model name
// @Tags AIModels
// @Param name path string true "Model Name" minlength(2) maxlength(18)
// @Accept json
// @Produce json
// @Success 200 {array} model.AIModel
// @Router /models/name/{name} [get]
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

// @BasePath /api/v1

// PingExample godoc
// @Summary Search model by username
// @Schemes
// @Description Find AI Models created by the user with the given username
// @Tags AIModels
// @Param username path string true "Username" minlength(5) maxlength(18)
// @Accept json
// @Produce json
// @Success 200 {array} model.AIModel
// @Router /models/username/{username} [get]
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

// @BasePath /api/v1

// PingExample godoc
// @Summary Create new model
// @Schemes
// @Description Subscribe new AIModel for inference
// @Tags AIModels
// @Param AIModel body model.AIModelCreate true "Create AIModel"
// @Accept json
// @Produce json
// @Success 201 {object} model.AIModel
// @Router /models/create [post]
func CreateNewModel(c *gin.Context) {
	var newModel model.AIModelCreate
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
	var mod model.AIModel
	mod.ModelID = uuid.New()
	mod.CreatedAt = time.Now()
	mod.ModifiedAt = time.Now()
	mod.ModelName = newModel.ModelName
	mod.CreatedBy = newModel.CreatedBy
	if err := data.Create(mod, devconfig.DBName, devconfig.ModelCollection); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, mod)
}

// @BasePath /api/v1

// PingExample godoc
// @Summary Edit a model
// @Schemes
// @Description Update a subscribed AIModel's details
// @Tags AIModels
// @Param AIModel body model.AIModelUpdate true "Update AIModel"
// @Accept json
// @Produce json
// @Success 200 {object} model.AIModel
// @Router /models/edit [put]
func EditModel(c *gin.Context) {
	var modUpdate model.AIModelUpdate
	if err := c.BindJSON(&modUpdate); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	filter := bson.D{{Key: "modelid", Value: modUpdate.ModelID}}
	findOptions := options.Find()
	var subscribedModels model.AIModels
	subscribedModels, err := data.Find(subscribedModels, devconfig.DBName, devconfig.ModelCollection, filter, findOptions)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(subscribedModels) == 0 {
		c.IndentedJSON(http.StatusNotFound, subscribedModels.ErrNotFound(modUpdate.ModelID).Error())
		return
	}
	mod, err := utils.First(subscribedModels)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	updateOptions := options.UpdateOne().SetUpsert(false)
	mod.ModelName = modUpdate.ModelName
	mod.ModifiedAt = time.Now()
	if err := data.EditOne(mod, devconfig.DBName, devconfig.ModelCollection, filter, updateOptions); err != nil {
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, mod)
}
