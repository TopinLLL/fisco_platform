// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/add/role": {
            "post": {
                "tags": [
                    "管理员操作"
                ],
                "summary": "添加用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "角色id",
                        "name": "role_id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "角色名称",
                        "name": "role_name",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/amend/calculate": {
            "post": {
                "tags": [
                    "电子合同通信与修正模块"
                ],
                "summary": "计算信誉值",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据ID",
                        "name": "data_id",
                        "in": "formData"
                    },
                    {
                        "type": "boolean",
                        "description": "该数据是否合法",
                        "name": "agree",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/amend/contract/edittx": {
            "post": {
                "tags": [
                    "电子合同通信与修正模块"
                ],
                "summary": "请求修正电子合同",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户私钥",
                        "name": "privateKey",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "交易地址",
                        "name": "txHash",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "编辑后的数据",
                        "name": "data",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/amend/contract/singletx": {
            "post": {
                "tags": [
                    "电子合同通信与修正模块"
                ],
                "summary": "上传电子合同",
                "parameters": [
                    {
                        "type": "string",
                        "description": "交易发送方",
                        "name": "from",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "交易接收方",
                        "name": "to",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "交易金额",
                        "name": "money",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/amend/contract/tracetx": {
            "post": {
                "tags": [
                    "电子合同通信与修正模块"
                ],
                "summary": "追溯修正操作",
                "parameters": [
                    {
                        "type": "string",
                        "description": "监管者私钥",
                        "name": "privateKey",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "交易地址",
                        "name": "txHash",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/amend/eliminate": {
            "post": {
                "tags": [
                    "电子合同通信与修正模块"
                ],
                "summary": "淘汰低信誉用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据ID",
                        "name": "data_id",
                        "in": "formData"
                    },
                    {
                        "type": "boolean",
                        "description": "该数据是否合法",
                        "name": "agree",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/blockchain/contract/deploy": {
            "post": {
                "tags": [
                    "用户操作"
                ],
                "summary": "部署智能合约",
                "responses": {}
            }
        },
        "/blockchain/height": {
            "get": {
                "tags": [
                    "管理员操作"
                ],
                "summary": "获取区块链高度",
                "responses": {}
            }
        },
        "/login": {
            "post": {
                "tags": [
                    "用户操作"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "账号",
                        "name": "username",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/privacy/aggregate": {
            "post": {
                "tags": [
                    "隐私保护模块"
                ],
                "summary": "聚合密钥",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "是否开启聚合",
                        "name": "start",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/privacy/groupcreate": {
            "post": {
                "tags": [
                    "隐私保护模块"
                ],
                "summary": "群创建",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户签名",
                        "name": "signature",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/privacy/groupjoin": {
            "post": {
                "tags": [
                    "隐私保护模块"
                ],
                "summary": "群加入",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户签名",
                        "name": "signature",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/privacy/open": {
            "post": {
                "tags": [
                    "隐私保护模块"
                ],
                "summary": "群签名打开",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户签名",
                        "name": "signature",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/privacy/sign": {
            "post": {
                "tags": [
                    "隐私保护模块"
                ],
                "summary": "电子合同签名",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户签名",
                        "name": "signature",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/privacy/verify": {
            "post": {
                "tags": [
                    "隐私保护模块"
                ],
                "summary": "群签名验证",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户签名",
                        "name": "signature",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/register": {
            "post": {
                "tags": [
                    "用户操作"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "账号",
                        "name": "username",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "角色id",
                        "name": "role_id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "mail",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/test/amend": {
            "post": {
                "tags": [
                    "实验"
                ],
                "summary": "可修正",
                "responses": {}
            }
        },
        "/test/amendattack": {
            "post": {
                "tags": [
                    "实验"
                ],
                "summary": "可修正作恶节点",
                "responses": {}
            }
        },
        "/test/privacy": {
            "post": {
                "tags": [
                    "实验"
                ],
                "summary": "群签名",
                "responses": {}
            }
        },
        "/test/privacydoor": {
            "post": {
                "tags": [
                    "实验"
                ],
                "summary": "不同实验群签名",
                "responses": {}
            }
        },
        "/verify/smartcontract": {
            "post": {
                "tags": [
                    "管理员操作"
                ],
                "summary": "确认智能合约",
                "parameters": [
                    {
                        "type": "string",
                        "description": "验证邮件",
                        "name": "verify_email",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/verify/user": {
            "post": {
                "tags": [
                    "管理员操作"
                ],
                "summary": "处理用户账号申请",
                "parameters": [
                    {
                        "type": "string",
                        "description": "申请id",
                        "name": "offer_id",
                        "in": "formData"
                    },
                    {
                        "type": "boolean",
                        "description": "是否同意创建",
                        "name": "ok",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:9091",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "面向电子合同的可修正区块链系统",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
