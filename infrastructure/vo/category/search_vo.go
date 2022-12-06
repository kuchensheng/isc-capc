package category

type SearchVO struct {
	Name     string       `json:"name"`
	Type     CategoryType `json:"type"`
	Ids      []int        `json:"ids"`
	Codes    []string     `json:"codes"`
	ParentId int          `json:"parentId"`
}

type CategoryType int

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
