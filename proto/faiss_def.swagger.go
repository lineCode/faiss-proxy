package faiss_server

const Swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "faiss_def.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/faiss/1.0/db/del": {
      "post": {
        "operationId": "DbDel",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/faiss_serverEmptyResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/faiss_serverDbDelRequest"
            }
          }
        ],
        "tags": [
          "FaissService"
        ]
      }
    },
    "/faiss/1.0/db/list": {
      "post": {
        "operationId": "DbList",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/faiss_serverDbListResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/faiss_serverDbListRequest"
            }
          }
        ],
        "tags": [
          "FaissService"
        ]
      }
    },
    "/faiss/1.0/db/new": {
      "post": {
        "operationId": "DbNew",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/faiss_serverEmptyResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/faiss_serverDbNewRequest"
            }
          }
        ],
        "tags": [
          "FaissService"
        ]
      }
    },
    "/faiss/1.0/hget": {
      "post": {
        "operationId": "HGet",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/faiss_serverHGetResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/faiss_serverHGetDelRequest"
            }
          }
        ],
        "tags": [
          "FaissService"
        ]
      }
    },
    "/faiss/1.0/hsearch": {
      "post": {
        "operationId": "HSearch",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/faiss_serverHSearchResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/faiss_serverHSearchRequest"
            }
          }
        ],
        "tags": [
          "FaissService"
        ]
      }
    },
    "/faiss/1.0/hset": {
      "post": {
        "operationId": "HSet",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/faiss_serverHSetResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/faiss_serverHSetRequest"
            }
          }
        ],
        "tags": [
          "FaissService"
        ]
      }
    },
    "/faiss/1.0/ping": {
      "post": {
        "operationId": "Ping",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/faiss_serverPingResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/faiss_serverPingRequest"
            }
          }
        ],
        "tags": [
          "FaissService"
        ]
      }
    }
  },
  "definitions": {
    "DbListResponseDbStatus": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "ntotal": {
          "type": "string",
          "format": "uint64"
        },
        "max_size": {
          "type": "string",
          "format": "uint64"
        },
        "curr_max_id": {
          "type": "string",
          "format": "uint64"
        },
        "curr_persist_max_id": {
          "type": "string",
          "format": "uint64"
        },
        "persist_path": {
          "type": "string"
        },
        "raw_data_path": {
          "type": "string"
        },
        "dimension": {
          "type": "string",
          "format": "uint64"
        },
        "model": {
          "type": "string"
        },
        "black_list_len": {
          "type": "string",
          "format": "uint64"
        }
      }
    },
    "HSearchRequestDistanceType": {
      "type": "string",
      "enum": [
        "Euclid",
        "Cosine"
      ],
      "default": "Euclid"
    },
    "HSearchResponseResult": {
      "type": "object",
      "properties": {
        "score": {
          "type": "number",
          "format": "float"
        },
        "id": {
          "type": "string",
          "format": "uint64"
        }
      }
    },
    "faiss_serverDbDelRequest": {
      "type": "object",
      "properties": {
        "db_name": {
          "type": "string"
        },
        "request_id": {
          "type": "string"
        }
      },
      "title": "删除db请求"
    },
    "faiss_serverDbListRequest": {
      "type": "object",
      "properties": {
        "request_id": {
          "type": "string"
        }
      },
      "title": "查询db列表请求"
    },
    "faiss_serverDbListResponse": {
      "type": "object",
      "properties": {
        "db_status": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/DbListResponseDbStatus"
          }
        },
        "error_code": {
          "type": "string",
          "format": "int64"
        },
        "error_msg": {
          "type": "string"
        },
        "request_id": {
          "type": "string"
        }
      },
      "title": "查询db列表返回结果"
    },
    "faiss_serverDbNewRequest": {
      "type": "object",
      "properties": {
        "db_name": {
          "type": "string"
        },
        "max_size": {
          "type": "string",
          "format": "uint64"
        },
        "model": {
          "type": "string"
        },
        "request_id": {
          "type": "string"
        }
      },
      "title": "创建db请求"
    },
    "faiss_serverEmptyResponse": {
      "type": "object",
      "properties": {
        "error_code": {
          "type": "string",
          "format": "int64"
        },
        "error_msg": {
          "type": "string"
        },
        "request_id": {
          "type": "string"
        }
      },
      "title": "默认返回"
    },
    "faiss_serverHGetDelRequest": {
      "type": "object",
      "properties": {
        "db_name": {
          "type": "string"
        },
        "id": {
          "type": "string",
          "format": "uint64"
        },
        "request_id": {
          "type": "string"
        }
      },
      "title": "获取或者删除一条特征请求"
    },
    "faiss_serverHGetResponse": {
      "type": "object",
      "properties": {
        "feature": {
          "type": "string",
          "format": "byte"
        },
        "dimension": {
          "type": "string",
          "format": "uint64"
        },
        "request_id": {
          "type": "string"
        },
        "error_code": {
          "type": "string",
          "format": "int64"
        },
        "error_msg": {
          "type": "string"
        }
      },
      "title": "获取一个特征的返回"
    },
    "faiss_serverHSearchRequest": {
      "type": "object",
      "properties": {
        "db_name": {
          "type": "string"
        },
        "feature": {
          "type": "string",
          "format": "byte"
        },
        "top_k": {
          "type": "string",
          "format": "uint64"
        },
        "distance_type": {
          "$ref": "#/definitions/HSearchRequestDistanceType"
        },
        "request_id": {
          "type": "string"
        }
      },
      "title": "ANN检索请求"
    },
    "faiss_serverHSearchResponse": {
      "type": "object",
      "properties": {
        "results": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/HSearchResponseResult"
          }
        },
        "request_id": {
          "type": "string"
        },
        "error_code": {
          "type": "string",
          "format": "int64"
        },
        "error_msg": {
          "type": "string"
        }
      },
      "title": "ANN 检索返回"
    },
    "faiss_serverHSetRequest": {
      "type": "object",
      "properties": {
        "db_name": {
          "type": "string"
        },
        "feature": {
          "type": "string",
          "format": "byte"
        },
        "request_id": {
          "type": "string"
        }
      },
      "title": "添加一条特征的请求"
    },
    "faiss_serverHSetResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "uint64"
        },
        "request_id": {
          "type": "string"
        },
        "error_code": {
          "type": "string",
          "format": "int64"
        },
        "error_msg": {
          "type": "string"
        }
      },
      "title": "添加一条特征的返回"
    },
    "faiss_serverPingRequest": {
      "type": "object",
      "properties": {
        "payload": {
          "type": "string"
        }
      },
      "title": "ping请求接口"
    },
    "faiss_serverPingResponse": {
      "type": "object",
      "properties": {
        "payload": {
          "type": "string"
        }
      },
      "title": "ping返回报文"
    }
  }
}
`
