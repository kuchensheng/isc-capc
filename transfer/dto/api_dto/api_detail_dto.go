package api_dto

type IscApiDetailDTO struct {
	IscApiInfoDTO

	//Parameters 入参信息模型
	Parameters string `json:"parameters"`
	//Resposes 出参信息模型
	Responses string `json:"responses"`
	//Type 出参类型，JSON/XML
	RespType string `json:"respType"`

	//Import 是否是导入类型
	Import bool `json:"import"`
}
