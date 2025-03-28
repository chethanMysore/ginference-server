// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "chethanMysore",
            "url": "https://chethanmysore.github.io/portfolio/#contact",
            "email": "willishardrock94@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "performs user authentication and returns JWT Auth token on success",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User Login",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register new user for inference",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Create new user",
                "parameters": [
                    {
                        "description": "Create User",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UserCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                }
            }
        },
        "/models": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Find all AI Models subscribed to the ginference-server",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AIModels"
                ],
                "summary": "Get all models",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.AIModel"
                            }
                        }
                    }
                }
            }
        },
        "/models/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Subscribe new AIModel for inference",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AIModels"
                ],
                "summary": "Create new model",
                "parameters": [
                    {
                        "description": "Create AIModel",
                        "name": "AIModel",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AIModelCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.AIModel"
                        }
                    }
                }
            }
        },
        "/models/edit": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update a subscribed AIModel's details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AIModels"
                ],
                "summary": "Edit a model",
                "parameters": [
                    {
                        "description": "Update AIModel",
                        "name": "AIModel",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AIModelUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AIModel"
                        }
                    }
                }
            }
        },
        "/models/id/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Find an AI Model using the given modelID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AIModels"
                ],
                "summary": "Search model by modelID",
                "parameters": [
                    {
                        "maxLength": 36,
                        "minLength": 36,
                        "type": "string",
                        "description": "Model ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AIModel"
                        }
                    }
                }
            }
        },
        "/models/name/{name}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Find AI Models matching the given model name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AIModels"
                ],
                "summary": "Search model by model name",
                "parameters": [
                    {
                        "maxLength": 18,
                        "minLength": 2,
                        "type": "string",
                        "description": "Model Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.AIModel"
                            }
                        }
                    }
                }
            }
        },
        "/models/username/{username}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Find AI Models created by the user with the given username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AIModels"
                ],
                "summary": "Search model by username",
                "parameters": [
                    {
                        "maxLength": 18,
                        "minLength": 5,
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.AIModel"
                            }
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Find all users registered with the ginference-server",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/user.User"
                            }
                        }
                    }
                }
            }
        },
        "/users/auth/id/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Find the user role created with the given username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Search user role by userID",
                "parameters": [
                    {
                        "maxLength": 36,
                        "minLength": 36,
                        "type": "string",
                        "description": "UserID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/edit": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update a registered User's details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Edit a user",
                "parameters": [
                    {
                        "description": "Update User",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UserUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                }
            }
        },
        "/users/id/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Find the user created with the given ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Search user by ID",
                "parameters": [
                    {
                        "maxLength": 36,
                        "minLength": 36,
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                }
            }
        },
        "/users/name/{name}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Find the users created with the given name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Search users by name",
                "parameters": [
                    {
                        "maxLength": 18,
                        "minLength": 2,
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/user.User"
                            }
                        }
                    }
                }
            }
        },
        "/users/username/{username}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Find the user created with the given username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Search user by username",
                "parameters": [
                    {
                        "maxLength": 18,
                        "minLength": 5,
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AIModel": {
            "type": "object",
            "required": [
                "createdBy",
                "modelName"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "createdBy": {
                    "type": "string"
                },
                "modelID": {
                    "type": "string"
                },
                "modelName": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 2
                },
                "modifiedAt": {
                    "type": "string"
                }
            }
        },
        "model.AIModelCreate": {
            "type": "object",
            "required": [
                "createdBy",
                "modelName"
            ],
            "properties": {
                "createdBy": {
                    "type": "string"
                },
                "modelName": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 2
                }
            }
        },
        "model.AIModelUpdate": {
            "type": "object",
            "required": [
                "modelID",
                "modelName"
            ],
            "properties": {
                "modelID": {
                    "type": "string"
                },
                "modelName": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 2
                }
            }
        },
        "user.User": {
            "type": "object",
            "required": [
                "countryCode",
                "emailID",
                "firstName",
                "lastName",
                "phone",
                "userName"
            ],
            "properties": {
                "countryCode": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "emailID": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 2
                },
                "fullName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 2
                },
                "modifiedAt": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                },
                "userName": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 5
                }
            }
        },
        "user.UserCreate": {
            "type": "object",
            "required": [
                "countryCode",
                "emailID",
                "firstName",
                "lastName",
                "password",
                "phone",
                "userName"
            ],
            "properties": {
                "countryCode": {
                    "type": "string"
                },
                "emailID": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 2
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 2
                },
                "password": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 8
                },
                "phone": {
                    "type": "string"
                },
                "userName": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 5
                }
            }
        },
        "user.UserUpdate": {
            "type": "object",
            "required": [
                "countryCode",
                "emailID",
                "firstName",
                "lastName",
                "phone",
                "userID"
            ],
            "properties": {
                "countryCode": {
                    "type": "string"
                },
                "emailID": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 2
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 2
                },
                "phone": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "\"Type 'Bearer TOKEN' to correctly set the API Key\"",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger ginference-server API",
	Description:      "This is a GO REST server for sentix inference.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
