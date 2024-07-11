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
            "name": "API Support",
            "url": "https://github.com/barganakukuhraditya",
            "email": "714220013@std.ulbi.ac.id"
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
        "/parfume": {
            "get": {
                "description": "Mengambil semua data parfume.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Parfume"
                ],
                "summary": "Get All Data Parfume.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Parfume"
                        }
                    }
                }
            }
        },
        "/parfume/{id}": {
            "get": {
                "description": "Ambil per ID data parfume.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Parfume"
                ],
                "summary": "Get By ID Data Parfume.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Parfume"
                        }
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
        }
    },
    "definitions": {
        "controller.Parfume": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string",
                    "example": "123456789"
                },
                "deskripsi": {
                    "type": "string",
                    "example": "Parfum yang sangat wangi"
                },
                "harga": {
                    "type": "integer",
                    "example": 1000000
                },
                "jenis_parfume": {
                    "type": "string",
                    "example": "Eau de Parfum"
                },
                "merk": {
                    "type": "string",
                    "example": "Dior"
                },
                "nama_parfume": {
                    "type": "string",
                    "example": "Chirstian Dior"
                },
                "stok": {
                    "type": "integer",
                    "example": 100
                },
                "tahun_peluncuran": {
                    "type": "integer",
                    "example": 2021
                },
                "ukuran": {
                    "type": "string",
                    "example": "100ml"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "tb-parfume2024-34a7b650de40.herokuapp.com",
	BasePath:         "/",
	Schemes:          []string{"https", "http"},
	Title:            "SWAGGER TUGAS BESAR",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
