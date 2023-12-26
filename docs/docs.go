// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/backoffice/laptops": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve a list of laptops",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "backoffice"
                ],
                "summary": "List laptops",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of laptops",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        },
        "/backoffice/laptops/{laptop_id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete a laptop by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "backoffice"
                ],
                "summary": "Delete laptop",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "ID of the laptop to be deleted",
                        "name": "laptop_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Laptop not found",
                        "schema": {}
                    }
                }
            }
        },
        "/users/": {
            "put": {
                "description": "Update user information with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "User update information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/param.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully updated user",
                        "schema": {
                            "$ref": "#/definitions/param.UpdateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {}
                    }
                }
            }
        },
        "/users/laptops/favorites": {
            "get": {
                "description": "Retrieve the list of laptops marked as favorites by the authenticated user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user's favorite laptops",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "user_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of favorite laptops",
                        "schema": {}
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            },
            "post": {
                "description": "Add a laptop to the user's favorites list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Add a laptop to favorites",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Laptop information to add to favorites",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/param.AddLaptopRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully added laptop to favorites",
                        "schema": {
                            "$ref": "#/definitions/param.AddLaptopResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        },
        "/users/laptops/favorites/{laptop_id}": {
            "delete": {
                "description": "Remove a laptop from the user's favorites list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Remove laptop from favorites",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "ID of the laptop to be removed from favorites",
                        "name": "laptop_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request parameters for removing a laptop from favorites",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/param.RemoveFavoriteLaptopRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully removed laptop from favorites",
                        "schema": {
                            "$ref": "#/definitions/param.RemoveFavoriteLaptopResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        },
        "/users/laptops/search": {
            "get": {
                "description": "Perform a search for laptops based on the provided criteria",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Search for laptops",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Search criteria for laptops",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/param.SearchRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Search results for laptops",
                        "schema": {
                            "$ref": "#/definitions/param.SearchResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Login with the provided credentials",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "User login information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/param.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully logged in",
                        "schema": {
                            "$ref": "#/definitions/param.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    }
                }
            }
        },
        "/users/profile": {
            "get": {
                "description": "Retrieve the profile information of the authenticated user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User profile information",
                        "schema": {
                            "$ref": "#/definitions/param.ProfileResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Register a new user with the provided information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/param.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully registered",
                        "schema": {
                            "$ref": "#/definitions/param.RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "param.AddLaptopRequest": {
            "type": "object",
            "properties": {
                "laptop_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "param.AddLaptopResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "param.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "param.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "$ref": "#/definitions/param.Tokens"
                },
                "user": {
                    "$ref": "#/definitions/param.UserInfo"
                }
            }
        },
        "param.ProfileResponse": {
            "type": "object",
            "properties": {
                "Info": {
                    "$ref": "#/definitions/param.UserInfo"
                }
            }
        },
        "param.RegisterRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "param.RegisterResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/param.UserInfo"
                }
            }
        },
        "param.RemoveFavoriteLaptopRequest": {
            "type": "object",
            "properties": {
                "laptop_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "param.RemoveFavoriteLaptopResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "param.SearchInfo": {
            "type": "object",
            "properties": {
                "company": {
                    "type": "string"
                },
                "cpu": {
                    "type": "string"
                },
                "graphic": {
                    "type": "integer"
                },
                "hdd": {
                    "type": "integer"
                },
                "ram": {
                    "type": "integer"
                },
                "screen_size": {
                    "type": "string"
                },
                "ssd": {
                    "type": "integer"
                }
            }
        },
        "param.SearchRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "laptop": {
                    "$ref": "#/definitions/param.SearchInfo"
                }
            }
        },
        "param.SearchResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "param.Tokens": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "param.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "param.UpdateUserResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "param.UserInfo": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
