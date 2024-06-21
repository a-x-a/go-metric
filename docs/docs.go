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
        "/list": {
            "get": {
                "description": "Возвращает список полученных от сервиса метрик ввиде обычного текста.",
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "list"
                ],
                "summary": "List",
                "operationId": "list",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "Performs a health check by pinging the service.",
                "tags": [
                    "ping"
                ],
                "summary": "Ping",
                "operationId": "ping",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/update": {
            "post": {
                "description": "Обновляет текущее значение метрики с указанным имененм и типом.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "update"
                ],
                "summary": "UpdateJSON",
                "operationId": "updateJSON",
                "parameters": [
                    {
                        "description": "Параметры метрики: имя, тип, значение",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/adapter.RequestMetric"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/update/{kind}/{name}/{value}": {
            "post": {
                "description": "Обновляет текущее значение метрики с указанным имененм и типом.",
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "update"
                ],
                "summary": "Update",
                "operationId": "update",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Тип метрики",
                        "name": "kind",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Имяметрики",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Значение метрики",
                        "name": "value",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/updates": {
            "post": {
                "description": "Обновляет текущие значения метрик из набора.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "updatebatch"
                ],
                "summary": "UpdateBatch",
                "operationId": "updatebatch",
                "parameters": [
                    {
                        "description": "Набор метрик",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/adapter.RequestMetric"
                            }
                        }
                    }
                ],
                "responses": {
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/value": {
            "get": {
                "description": "Возвращает текущее значение метрики в формате JSON с указанным имененм и типом.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "value"
                ],
                "summary": "GetJSON",
                "operationId": "getJSON",
                "parameters": [
                    {
                        "description": "Параметры метрики: имя, тип",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/adapter.RequestMetric"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/value/{kind}/{name}": {
            "get": {
                "description": "Возвращает текущее значение метрики с указанным имененм и типом.",
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "value"
                ],
                "summary": "Get",
                "operationId": "get",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Тип метрики",
                        "name": "kind",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Имя метрики",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        }
    },
    "definitions": {
        "adapter.RequestMetric": {
            "type": "object",
            "properties": {
                "delta": {
                    "description": "значение метрики в случае передачи counter",
                    "type": "integer"
                },
                "id": {
                    "description": "имя метрики",
                    "type": "string"
                },
                "type": {
                    "description": "параметр, принимающий значение gauge или counter",
                    "type": "string"
                },
                "value": {
                    "description": "значение метрики в случае передачи gauge",
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Сервис сбора метрик и алертинга.",
	Description:      "Сервис для сбора рантайм-метрик.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}