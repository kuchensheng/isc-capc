package category

import "strings"

const SEP = "/"

type SearchVO struct {
	Name     string       `json:"name"`
	Type     CategoryType `json:"type"`
	Ids      []int        `json:"ids"`
	Codes    []string     `json:"codes"`
	ParentId int          `json:"parentId"`
}

type CategoryType int

var categoryTypeNames = []string{"DEFAULT", "NATIVE", "POLYMERIC", "LIGHT", "OS", "UDMP", "TDDM"}

const (
	//DEFAULT 查询所有
	DEFAULT CategoryType = iota
	//NATIVE 云原生接口，表示三方应用
	NATIVE
	//POLYMERIC 逻辑编排接口
	POLYMERIC
	//LIGHT 轻应用接口
	LIGHT
	//OS 系统内置接口
	OS
	//UDMP 通用数据管理生成的接口
	UDMP
	//TDDM 表对象建模生成的接口
	TDDM
)

func (t CategoryType) GetIdx() int {
	return int(t)
}

func (t CategoryType) GetName() string {
	return categoryTypeNames[t]
}

func String2CategoryType(name string) CategoryType {
	for i, typeName := range categoryTypeNames {
		if typeName == name {
			return CategoryType(i)
		}
	}
	return NATIVE
}

var statusNames = []string{"PUBLISHED", "DESIGN", "DEVELOP", "DEBUG", "TESTING", "OBSOLETED", "UNPUBLISHED"}

type Status int

const (
	//PUBLISHED 已发布，默认值
	PUBLISHED Status = iota
	//DESIGN 设计中
	DESIGN
	//DEVELOP 开发中
	DEVELOP
	//DEBUG 联调中
	DEBUG
	//TESTING 测试中
	TESTING
	//OBSOLETED 已废弃
	OBSOLETED
	//UNPUBLISHED 未发布
	UNPUBLISHED
)

func String2Status(name string) Status {
	for i, typeName := range statusNames {
		if typeName == name {
			return Status(i)
		}
	}
	return PUBLISHED
}

func (t Status) GetName() string {
	return statusNames[t]
}

var authTypeNames = []string{"NO_AUTH", "BASE_AUTH"}

type AuthType int

const (
	NO_AUTH = iota
	BASE_AUTH
)

func (t AuthType) GetName() string {
	return authTypeNames[t]
}

func String2AuthType(name string) AuthType {
	for i, typeName := range authTypeNames {
		if typeName == name {
			return AuthType(i)
		}
	}
	return NO_AUTH
}

var protocolNames = []string{"HTTP", "HTTPS", "WS", "WSS"}

type Protocol int

const (
	HTTP Protocol = iota
	HTTPS
	WS
	WSS
)

func (t Protocol) GetName() string {
	return protocolNames[t]
}

func String2Protocol(name string) Protocol {
	for i, typeName := range protocolNames {
		if typeName == strings.ToUpper(name) {
			return Protocol(i)
		}
	}
	return HTTP
}
