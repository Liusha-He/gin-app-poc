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
        "/api/v1/accounts": {
            "get": {
                "description": "Responds with the list of accounts",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Get account list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dao.Account"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Takes an account json and store in DB, Returned saved json.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "create account",
                "parameters": [
                    {
                        "description": "account json",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.createAccountRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dao.Account"
                        }
                    }
                }
            }
        },
        "/api/v1/accounts/{id}": {
            "get": {
                "description": "Takes an account id with path, Returned account info json.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "get account by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dao.Account"
                        }
                    }
                }
            }
        },
        "/api/v1/transfers": {
            "post": {
                "description": "Takes an transfer json and store in DB, Returned saved json.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transfers"
                ],
                "summary": "create transfer",
                "parameters": [
                    {
                        "description": "transfer json",
                        "name": "transfer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.transferRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dao.Transfer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.createAccountRequest": {
            "type": "object",
            "required": [
                "currency",
                "owner"
            ],
            "properties": {
                "currency": {
                    "type": "string",
                    "enum": [
                        "USD",
                        "EUR",
                        "GBP",
                        "CNY"
                    ]
                },
                "owner": {
                    "type": "string"
                }
            }
        },
        "api.transferRequest": {
            "type": "object",
            "required": [
                "amount",
                "currency",
                "from_account_id",
                "to_account_id"
            ],
            "properties": {
                "amount": {
                    "type": "number",
                    "minimum": 1
                },
                "currency": {
                    "type": "string"
                },
                "from_account_id": {
                    "type": "integer",
                    "minimum": 1
                },
                "to_account_id": {
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "dao.Account": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "owner": {
                    "type": "string"
                }
            }
        },
        "dao.Transfer": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "from_account_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "to_account_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "simple bank API",
	Description:      "A Golang and gin API template",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
