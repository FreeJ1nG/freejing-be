// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Andrew Jeremy",
            "email": "Andrewjeremy12345@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/auth": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create new user",
                "parameters": [
                    {
                        "description": "Create new user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dbquery.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpm.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dbquery.User"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/auth/{username}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update user with a certain username",
                "parameters": [
                    {
                        "description": "Update user with a certain username",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dbquery.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpm.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dbquery.User"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "user"
                ],
                "summary": "Delete user with a certain username",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/v1/blogs": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blog"
                ],
                "summary": "Get blog posts",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "type": "array",
                            "items": {
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/httpm.Response"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "data": {
                                                "type": "array",
                                                "items": {
                                                    "$ref": "#/definitions/dbquery.Blog"
                                                }
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blog"
                ],
                "summary": "Create new blog post",
                "parameters": [
                    {
                        "description": "Create Blog Post",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dbquery.Blog"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpm.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dbquery.Blog"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/blogs/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blog"
                ],
                "summary": "Get blog post with a certain id",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpm.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dbquery.Blog"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blog"
                ],
                "summary": "Update blog post with a certain id",
                "parameters": [
                    {
                        "description": "Update Blog Post",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dbquery.Blog"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpm.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dbquery.Blog"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "blog"
                ],
                "summary": "Delete blog post with a certain id",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/v1/user/{username}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user with a certain username",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpm.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dbquery.User"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dbquery.Blog": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "create_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dbquery.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "passwordHash": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "httpm.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "api.freejing.com",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "FreeJinG API",
	Description:      "This API is a self-made project made with golang, this repository can be accessed on: https://github.com/FreeJ1nG/freejing-be",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
