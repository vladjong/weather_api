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
        "/cities": {
            "get": {
                "description": "get all cities",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "weather"
                ],
                "summary": "Get Cities",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.AllCities"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/cities/{name}": {
            "get": {
                "description": "get string by NANE",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "weather"
                ],
                "summary": "Get prediction weather in City",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.WeatherPredict"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/detail_weather/{name}/{date}": {
            "get": {
                "description": "get STRING by NAME and DATE (YYYY-MM-DDTHH:MM:SSZ)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "weather"
                ],
                "summary": "Get all information about the weather in the city on the exact day",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.WeatherDetails"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.AllCities": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "entities.Info": {
            "type": "object",
            "properties": {
                "clouds": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "number"
                    }
                },
                "infoWeather": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.InfoWeather"
                    }
                },
                "main": {
                    "$ref": "#/definitions/entities.Main"
                },
                "visibility": {
                    "type": "number"
                },
                "wind": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "number"
                    }
                }
            }
        },
        "entities.InfoWeather": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "main": {
                    "type": "string"
                }
            }
        },
        "entities.Main": {
            "type": "object",
            "properties": {
                "feels_like": {
                    "type": "number"
                },
                "grind_level": {
                    "type": "number"
                },
                "humidity": {
                    "type": "number"
                },
                "pressure": {
                    "type": "number"
                },
                "sea_level": {
                    "type": "number"
                },
                "temp": {
                    "type": "number"
                },
                "temp_kf": {
                    "type": "number"
                },
                "temp_map": {
                    "type": "number"
                },
                "temp_min": {
                    "type": "number"
                }
            }
        },
        "entities.WeatherDetails": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "info": {
                    "$ref": "#/definitions/entities.Info"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entities.WeatherPredict": {
            "type": "object",
            "properties": {
                "av_temp": {
                    "type": "number"
                },
                "country": {
                    "type": "string"
                },
                "dates": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "v1.errorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Weather API",
	Description:      "This is a service that predicts the weather",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}