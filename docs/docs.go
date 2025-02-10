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
        "/ping": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ping"
                ],
                "summary": "GET ping used for healthcheck",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/get_ping_response.Response"
                        }
                    }
                }
            }
        },
        "/user/auth/sign_up": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Auth-SignUp"
                ],
                "summary": "User sign up",
                "parameters": [
                    {
                        "description": "Sign up request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/sign_up_request.Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sign_up_response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/validation_error_response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_server_error_response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "field_error_response.Response": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string",
                    "example": "email"
                },
                "tag": {
                    "type": "string",
                    "example": "required"
                }
            }
        },
        "get_ping_response.Data": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "pong"
                }
            }
        },
        "get_ping_response.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/get_ping_response.Data"
                }
            }
        },
        "internal_server_error_response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "INTERNAL_SERVER_ERROR"
                },
                "message": {
                    "type": "string",
                    "example": "Internal Server Error"
                }
            }
        },
        "sign_up_request.Request": {
            "type": "object",
            "required": [
                "confirm_password",
                "email",
                "password"
            ],
            "properties": {
                "confirm_password": {
                    "type": "string",
                    "example": "Aa@123456"
                },
                "email": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "user@example.com"
                },
                "password": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "Aa@123456"
                }
            }
        },
        "sign_up_response.Data": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string",
                    "example": "xxx"
                },
                "expires_at": {
                    "type": "string",
                    "example": "2025-12-31T12:30:00.0000+07:00"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "refresh_token": {
                    "type": "string",
                    "example": "xxx"
                }
            }
        },
        "sign_up_response.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/sign_up_response.Data"
                }
            }
        },
        "validation_error_response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "VALIDATION_ERROR"
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/field_error_response.Response"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "Validation Error"
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
