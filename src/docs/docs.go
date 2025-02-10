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
        "/v1/expense/": {
            "get": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "List and filter expenses",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Expenses"
                ],
                "summary": "List Expenses",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filter (past_week, past_month, last_3_months)",
                        "name": "filter",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Start Date (YYYY-MM-DD)",
                        "name": "startDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "End Date (YYYY-MM-DD)",
                        "name": "endDate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Expense response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_dto.ExpenseResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "Create a new Expense",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Expenses"
                ],
                "summary": "Create a Expense",
                "parameters": [
                    {
                        "description": "Create a Expense",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_dto.CreateExpenseRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Expense response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_dto.ExpenseResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v1/expense/{id}": {
            "get": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "Get a Expense by Id",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Expenses"
                ],
                "summary": "Get a Expense",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Expense response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_dto.ExpenseResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "Update a Expense by Id",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Expenses"
                ],
                "summary": "Update a Expense",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update a Expense",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_dto.UpdateExpenseRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Expense response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_dto.ExpenseResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "Delete a Expense by Id",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Expenses"
                ],
                "summary": "Delete a Expense",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No content",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v1/tokens/refresh-token": {
            "post": {
                "description": "RefreshToken",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "RefreshToken",
                "responses": {
                    "201": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse"
                        }
                    },
                    "409": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v1/users/login-by-username": {
            "post": {
                "description": "LoginByUsername",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "LoginByUsername",
                "parameters": [
                    {
                        "description": "LoginByUsernameRequest",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_dto.LoginByUsernameRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse"
                        }
                    },
                    "409": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v1/users/register-by-username": {
            "post": {
                "description": "RegisterByUsername",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "RegisterByUsername",
                "parameters": [
                    {
                        "description": "RegisterUserByUsernameRequest",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_dto.RegisterUserByUsernameRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse"
                        }
                    },
                    "409": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_alielmi98_golang-expense-tracker-api_api_dto.CreateExpenseRequest": {
            "type": "object",
            "required": [
                "amount",
                "title"
            ],
            "properties": {
                "amount": {
                    "type": "number"
                },
                "category": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "github_com_alielmi98_golang-expense-tracker-api_api_dto.ExpenseResponse": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "category": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "github_com_alielmi98_golang-expense-tracker-api_api_dto.LoginByUsernameRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string",
                    "minLength": 5
                }
            }
        },
        "github_com_alielmi98_golang-expense-tracker-api_api_dto.RegisterUserByUsernameRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "minLength": 6
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string",
                    "minLength": 5
                }
            }
        },
        "github_com_alielmi98_golang-expense-tracker-api_api_dto.UpdateExpenseRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "category": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "github_com_alielmi98_golang-expense-tracker-api_api_helper.BaseHttpResponse": {
            "type": "object",
            "properties": {
                "error": {},
                "result": {},
                "resultCode": {
                    "$ref": "#/definitions/github_com_alielmi98_golang-expense-tracker-api_api_helper.ResultCode"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "github_com_alielmi98_golang-expense-tracker-api_api_helper.ResultCode": {
            "type": "integer",
            "enum": [
                0,
                40001,
                40101,
                40301,
                40401,
                42901,
                42902,
                50001,
                50002,
                50003
            ],
            "x-enum-varnames": [
                "Success",
                "ValidationError",
                "AuthError",
                "ForbiddenError",
                "NotFoundError",
                "LimiterError",
                "OtpLimiterError",
                "CustomRecovery",
                "InternalError",
                "InvalidInputError"
            ]
        }
    },
    "securityDefinitions": {
        "AuthBearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
