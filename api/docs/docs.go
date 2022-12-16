// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/film": {
            "get": {
                "description": "Get List Film",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Film"
                ],
                "summary": "Get List Film",
                "operationId": "get_list_film",
                "responses": {
                    "200": {
                        "description": "GetFilmBody",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Film"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Film",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Film"
                ],
                "summary": "Create Film",
                "operationId": "create_film",
                "parameters": [
                    {
                        "description": "CreateFilmRequestBody",
                        "name": "film",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Film"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "GetFilmBody",
                        "schema": {
                            "$ref": "#/definitions/models.Film"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/film/": {
            "put": {
                "description": "Update Film",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Film"
                ],
                "summary": "Update Film",
                "operationId": "update_film",
                "parameters": [
                    {
                        "description": "CreateFilmRequestBody",
                        "name": "film",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Film"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetFilmsBody",
                        "schema": {
                            "$ref": "#/definitions/models.Film"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/film/{id}": {
            "get": {
                "description": "Get By Id Film",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Film"
                ],
                "summary": "Get By Id Film",
                "operationId": "get_by_id_film",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetFilmBody",
                        "schema": {
                            "$ref": "#/definitions/models.Film"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete By Id Film",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Film"
                ],
                "summary": "Delete By Id Film",
                "operationId": "delete_by_id_film",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetFilmBody",
                        "schema": {
                            "$ref": "#/definitions/models.Film"
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Film": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "film_id": {
                    "type": "integer"
                },
                "fulltext": {
                    "type": "string"
                },
                "language_id": {
                    "type": "integer"
                },
                "last_update": {
                    "type": "string"
                },
                "length": {
                    "type": "integer"
                },
                "rating": {},
                "release_year": {
                    "type": "integer"
                },
                "rentalRate": {
                    "type": "number"
                },
                "rental_duration": {
                    "type": "integer"
                },
                "replacement_cost": {
                    "type": "number"
                },
                "special_features": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}