package data

import (
	"example/ginference-server/models/model"
	"example/ginference-server/models/user"
	"time"

	"github.com/google/uuid"
)

var RegisteredUsers = user.Users{
	{UserID: uuid.New(), UserName: "Tom", CreatedAt: time.Now()},
	{UserID: uuid.New(), UserName: "Joe", CreatedAt: time.Now()},
	{UserID: uuid.New(), UserName: "Harry", CreatedAt: time.Now()},
}

var user1, _ = RegisteredUsers.FindByName("Tom")
var user2, _ = RegisteredUsers.FindByName("Joe")

//var user3, _ = RegisteredUsers.FindByName("Harry")

var SubscribedModels = model.AIModels{
	{ModelID: uuid.New(), ModelName: "pickachu_1", CreatedBy: user1.UserID.String(), CreatedAt: time.Now()},
	{ModelID: uuid.New(), ModelName: "bulbasaur_1", CreatedBy: user2.UserID.String(), CreatedAt: time.Now()},
	{ModelID: uuid.New(), ModelName: "bulbasaur_2", CreatedBy: user2.UserID.String(), CreatedAt: time.Now()},
}
