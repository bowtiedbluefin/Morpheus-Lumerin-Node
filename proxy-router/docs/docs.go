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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/blockchain/allowance": {
            "get": {
                "description": "Get MOR allowance for spender",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get Allowance for MOR",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Spender address",
                        "name": "spender",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "description": "Approve MOR allowance for spender",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Approve MOR allowance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Spender address",
                        "name": "spender",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Amount",
                        "name": "amount",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/blockchain/balance": {
            "get": {
                "description": "Get ETH and MOR balance of the user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get ETH and MOR balance",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/blockchain/bids/{id}/session": {
            "post": {
                "description": "Full flow to open a session by bidId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sessions"
                ],
                "summary": "Open Session by bidId in blockchain",
                "parameters": [
                    {
                        "description": "Open session",
                        "name": "opensession",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.OpenSessionWithDurationRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bid ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/blockchain/latestBlock": {
            "get": {
                "description": "Get latest block number from blockchain",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get Latest Block",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/blockchain/models": {
            "get": {
                "description": "Get models list from blokchain",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get models list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "object"
                            }
                        }
                    }
                }
            }
        },
        "/blockchain/models/{id}/bids": {
            "get": {
                "description": "Get bids from blockchain by model agent",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get Active Bids by Model Agent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ModelAgent ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "object"
                            }
                        }
                    }
                }
            }
        },
        "/blockchain/models/{id}/session": {
            "post": {
                "description": "Full flow to open a session by modelId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sessions"
                ],
                "summary": "Open Session by ModelID in blockchain",
                "parameters": [
                    {
                        "description": "Open session",
                        "name": "opensession",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.OpenSessionWithDurationRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Model ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/blockchain/providers": {
            "get": {
                "description": "Get providers list from blokchain",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get providers list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "object"
                            }
                        }
                    }
                }
            }
        },
        "/blockchain/providers/{id}/bids": {
            "get": {
                "description": "Get bids from blockchain by provider",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get Bids by Provider",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Provider ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "object"
                            }
                        }
                    }
                }
            }
        },
        "/blockchain/providers/{id}/bids/active": {
            "get": {
                "description": "Get bids from blockchain by provider",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get Bids by Provider",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Provider ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "object"
                            }
                        }
                    }
                }
            }
        },
        "/blockchain/send/eth": {
            "post": {
                "description": "Send Eth to address",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Send Eth",
                "parameters": [
                    {
                        "description": "Send Eth",
                        "name": "sendeth",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.SendRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/blockchain/send/mor": {
            "post": {
                "description": "Send Mor to address",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Send Mor",
                "parameters": [
                    {
                        "description": "Send Mor",
                        "name": "sendmor",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.SendRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/blockchain/sessions": {
            "get": {
                "description": "Get sessions from blockchain by user or provider",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sessions"
                ],
                "summary": "Get Sessions",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Provider address",
                        "name": "provider",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "User address",
                        "name": "user",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "object"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Sends transaction in blockchain to open a session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sessions"
                ],
                "summary": "Open Session with Provider in blockchain",
                "parameters": [
                    {
                        "description": "Open session",
                        "name": "opensession",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.OpenSessionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/blockchain/sessions/budget": {
            "get": {
                "description": "Get todays budget from blockchain",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get Todays Budget",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/blockchain/sessions/{id}/close": {
            "post": {
                "description": "Sends transaction in blockchain to close a session",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sessions"
                ],
                "summary": "Close Session with Provider",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Session ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/blockchain/token/supply": {
            "get": {
                "description": "Get MOR token supply from blockchain",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get Token Supply",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/blockchain/transactions": {
            "get": {
                "description": "Get MOR and ETH transactions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get Transactions",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "object"
                            }
                        }
                    }
                }
            }
        },
        "/config": {
            "get": {
                "description": "Return the current config of proxy router",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthcheck"
                ],
                "summary": "Get Config",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/system.ConfigResponse"
                        }
                    }
                }
            }
        },
        "/files": {
            "get": {
                "description": "Returns opened files",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthcheck"
                ],
                "summary": "Get files",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/system.FD"
                            }
                        }
                    }
                }
            }
        },
        "/healthcheck": {
            "get": {
                "description": "do ping",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthcheck"
                ],
                "summary": "Healthcheck example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/system.HealthCheckResponse"
                        }
                    }
                }
            }
        },
        "/proxy/sessions/initiate": {
            "post": {
                "description": "sends a handshake to the provider",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sessions"
                ],
                "summary": "Initiate Session with Provider",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/proxy/sessions/{id}/providerClaim": {
            "post": {
                "description": "Claim provider balance from session",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sessions"
                ],
                "summary": "Claim Provider Balance",
                "parameters": [
                    {
                        "description": "Claim",
                        "name": "claim",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.SendRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Session ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/proxy/sessions/{id}/providerClaimableBalance": {
            "get": {
                "description": "Get provider claimable balance from session",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sessions"
                ],
                "summary": "Get Provider Claimable Balance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Session ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/v1/chat/completions": {
            "post": {
                "description": "Send prompt to a local or remote model based on session id in header",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Send Local Or Remote Prompt",
                "parameters": [
                    {
                        "description": "Prompt",
                        "name": "prompt",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/proxyapi.OpenAiCompletitionRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Session ID",
                        "name": "session_id",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/v1/models": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get local models",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/aiengine.LocalModel"
                            }
                        }
                    }
                }
            }
        },
        "/wallet": {
            "get": {
                "description": "Get wallet address",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get Wallet",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "description": "Set wallet private key",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Set Wallet",
                "parameters": [
                    {
                        "description": "Private key",
                        "name": "privatekey",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/walletapi.SetupWalletReqBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "aiengine.LocalModel": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "lib.BigInt": {
            "type": "object"
        },
        "proxyapi.ChatCompletionMessage": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "multiContent": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/proxyapi.ChatMessagePart"
                    }
                },
                "name": {
                    "description": "This property isn't in the official documentation, but it's in\nthe documentation for the official library for python:\n- https://github.com/openai/openai-python/blob/main/chatml.md\n- https://github.com/openai/openai-cookbook/blob/main/examples/How_to_count_tokens_with_tiktoken.ipynb",
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "tool_call_id": {
                    "description": "For Role=tool prompts this should be set to the ID given in the assistant's prior request to call a tool.",
                    "type": "string"
                }
            }
        },
        "proxyapi.ChatCompletionResponseFormat": {
            "type": "object",
            "properties": {
                "type": {
                    "type": "string"
                }
            }
        },
        "proxyapi.ChatMessageImageURL": {
            "type": "object",
            "properties": {
                "detail": {
                    "$ref": "#/definitions/proxyapi.ImageURLDetail"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "proxyapi.ChatMessagePart": {
            "type": "object",
            "properties": {
                "image_url": {
                    "$ref": "#/definitions/proxyapi.ChatMessageImageURL"
                },
                "text": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/proxyapi.ChatMessagePartType"
                }
            }
        },
        "proxyapi.ChatMessagePartType": {
            "type": "string",
            "enum": [
                "text",
                "image_url"
            ],
            "x-enum-varnames": [
                "ChatMessagePartTypeText",
                "ChatMessagePartTypeImageURL"
            ]
        },
        "proxyapi.ImageURLDetail": {
            "type": "string",
            "enum": [
                "high",
                "low",
                "auto"
            ],
            "x-enum-varnames": [
                "ImageURLDetailHigh",
                "ImageURLDetailLow",
                "ImageURLDetailAuto"
            ]
        },
        "proxyapi.OpenAiCompletitionRequest": {
            "type": "object",
            "properties": {
                "frequency_penalty": {
                    "type": "number"
                },
                "function_call": {
                    "description": "Deprecated: use ToolChoice instead."
                },
                "logit_bias": {
                    "description": "LogitBias is must be a token id string (specified by their token ID in the tokenizer), not a word string.\nincorrect: ` + "`" + `\"logit_bias\":{\"You\": 6}` + "`" + `, correct: ` + "`" + `\"logit_bias\":{\"1639\": 6}` + "`" + `\nrefs: https://platform.openai.com/docs/api-reference/chat/create#chat/create-logit_bias",
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "logprobs": {
                    "description": "LogProbs indicates whether to return log probabilities of the output tokens or not.\nIf true, returns the log probabilities of each output token returned in the content of message.\nThis option is currently not available on the gpt-4-vision-preview model.",
                    "type": "boolean"
                },
                "max_tokens": {
                    "type": "integer"
                },
                "messages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/proxyapi.ChatCompletionMessage"
                    }
                },
                "model": {
                    "type": "string"
                },
                "n": {
                    "type": "integer"
                },
                "presence_penalty": {
                    "type": "number"
                },
                "response_format": {
                    "$ref": "#/definitions/proxyapi.ChatCompletionResponseFormat"
                },
                "seed": {
                    "type": "integer"
                },
                "stop": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "stream": {
                    "type": "boolean"
                },
                "temperature": {
                    "type": "number"
                },
                "tool_choice": {
                    "description": "This can be either a string or an ToolChoice object."
                },
                "top_logprobs": {
                    "description": "TopLogProbs is an integer between 0 and 5 specifying the number of most likely tokens to return at each\ntoken position, each with an associated log probability.\nlogprobs must be set to true if this parameter is used.",
                    "type": "integer"
                },
                "top_p": {
                    "type": "number"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "structs.OpenSessionRequest": {
            "type": "object"
        },
        "structs.OpenSessionWithDurationRequest": {
            "type": "object",
            "properties": {
                "sessionDuration": {
                    "$ref": "#/definitions/lib.BigInt"
                }
            }
        },
        "structs.SendRequest": {
            "type": "object"
        },
        "system.ConfigResponse": {
            "type": "object",
            "properties": {
                "commit": {
                    "type": "string"
                },
                "config": {},
                "derivedConfig": {},
                "version": {
                    "type": "string"
                }
            }
        },
        "system.FD": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "system.HealthCheckResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                },
                "uptime": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "walletapi.SetupWalletReqBody": {
            "type": "object",
            "required": [
                "privateKey"
            ],
            "properties": {
                "privateKey": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "ApiBus Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
