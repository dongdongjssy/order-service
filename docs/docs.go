// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths" : {
        "/orders/transform" : {
            "post" : {
                "tags" : [ "Orders" ],
                "summary" : "Transform a list of orders",
                "description" : "It transforms a list of orders to the format of a list of customer items which consists of item details, the count of purchased items, and total cost etc.",
                "parameters" : [ {
                    "in" : "body",
                    "name" : "orders",
                    "description" : "a list of orders",
                    "required" : true,
                    "schema" : {
                        "type" : "array",
                        "items" : {
                        "$ref" : "#/definitions/Order"
                        }
                    }
                } ],
                "responses" : {
                    "200" : {
                        "description" : "success",
                        "schema" : {
                        "$ref" : "#/definitions/SuccessResponse"
                        }
                    },
                    "400" : {
                        "description" : "bad request",
                        "schema" : {
                        "$ref" : "#/definitions/ErrorResponse"
                        }
                    },
                    "500" : {
                        "description" : "internal server error",
                        "schema" : {
                        "$ref" : "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions" : {
        "ItemRequest" : {
            "required" : [ "costEur", "itemId" ],
            "properties" : {
                "itemId" : {
                    "type" : "string",
                    "example" : "20201"
                },
                "costEur" : {
                    "type" : "number",
                    "example" : 2.5
                }
            }
        },
        "ItemResponse" : {
            "required" : [ "costEur", "customerId", "itemId" ],
            "properties" : {
                "customerId" : {
                    "type" : "string",
                    "example" : "01"
                },
                "itemId" : {
                    "type" : "string",
                    "example" : "20201"
                },
                "costEur" : {
                    "type" : "number",
                    "example" : 2.5
                }
            }
        },
        "Order" : {
            "properties" : {
                "customerId" : {
                    "type" : "string",
                    "example" : "01"
                },
                "orderId" : {
                    "type" : "string",
                    "example" : "50"
                },
                "timestamp" : {
                    "type" : "string",
                    "example" : "1637245070513"
                },
                "items" : {
                    "type" : "array",
                    "items" : {
                        "$ref" : "#/definitions/ItemRequest"
                    }
                }
            }
        },
        "Summary" : {
            "properties" : {
                "customerId" : {
                    "type" : "string",
                    "example" : "01"
                },
                "nbrOfPurchasedItems" : {
                    "type" : "number",
                    "example" : 200.0
                },
                "totalAmountEur" : {
                    "type" : "number",
                    "example" : 15000.0
                },
                "items" : {
                    "type" : "array",
                    "items" : {
                        "$ref" : "#/definitions/ItemResponse"
                    }
                }
            }
        },
        "SuccessResponse" : {
            "required" : [ "code", "data", "message" ],
            "properties" : {
                "code" : {
                    "type" : "integer",
                    "example" : 200
                },
                "message" : {
                    "type" : "string",
                    "example" : "success"
                },
                "data" : {
                    "type" : "array",
                    "items" : {
                        "$ref" : "#/definitions/Summary"
                    }
                }
            }
        },
        "ErrorResponse" : {
            "required" : [ "code", "errors", "message" ],
            "properties" : {
                "code" : {
                    "type" : "integer"
                },
                "message" : {
                    "type" : "string"
                },
                "errors" : {
                    "type" : "array",
                    "items" : {
                        "type" : "string"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Orders Transform Service",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
