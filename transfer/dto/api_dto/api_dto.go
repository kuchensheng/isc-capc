package api_dto

import (
	"encoding/json"
	"github.com/kuchensheng/capc/infrastructure/model/api"
	"github.com/kuchensheng/capc/infrastructure/vo/category"
	"strconv"
	"strings"
)

type IscApiInfoDTO struct {
	ID int `json:"id"`
	//Summary api名称
	Summary string `json:"summary"`
	//Path api路径
	Path string `json:"path"`
	//Method 请求方法,GET、POST、DELETE、PUT、PATCH等
	Method string `json:"method"`
	//Code api唯一标识
	Code string `json:"code"`
	//Type api类型，OS、NATIVE、POLYMERIC、LIGHT、UDMP、TDDM等
	Type string `json:"type"`
	//CategoryId 分组ID
	CategoryId int `json:"groupId"`
	//CategoryPath 分组全路径ID
	CategoryPath []int `json:"groupPath"`
	//CategoryFullName 分组全路径名称
	CategoryFullName []string `json:"groupNamePath"`
	//Status 生命周期状态
	Status string `json:"status"`
	//Tags 标签信息
	Tags string `json:"tags"`
	//Version 版本信息，默认1.0.0
	Version string `json:"version"`
	//Description 描述信息
	Description string `json:"description"`
	//Protocol 协议信息，HTTP、HTTPS、WS、WSS、WQ等，默认HTTP
	Protocol string `json:"protocol"`
	//Principal 责任人信息
	Principal string `json:"principal"`
	//AuthType 鉴权方式
	AuthType string `json:"authType"`
	//AuthConfig 鉴权配置
	AuthConfig string `json:"authConfig"`
	//Cosumes 解析方式，是个数组
	Consumes string `json:"consumes"`
}

func (dto *IscApiDetailDTO) Dto2DO() (*api.IscCapcApiInfo, *api.IscCapcApiReqResp) {
	model := &api.IscCapcApiInfo{}
	data, _ := json.Marshal(dto)
	_ = json.Unmarshal(data, model)
	model.Type = category.String2CategoryType(dto.Type).GetIdx()
	model.CategoryPath = strings.Join(func() []string {
		var strIds []string
		for _, id := range dto.CategoryPath {
			strIds = append(strIds, strconv.Itoa(id))
		}
		return strIds
	}(), category.SEP)
	model.CategoryId = dto.CategoryId
	model.CategoryFullName = strings.Join(dto.CategoryFullName, category.SEP)
	model.Status = int(category.String2Status(dto.Status))
	model.AuthType = int(category.String2AuthType(dto.AuthType))
	model.Protocol = category.String2Protocol(dto.Protocol).GetName()
	var resp *api.IscCapcApiReqResp
	if dto.Parameters != "" || dto.Responses != "" {
		resp = &api.IscCapcApiReqResp{
			Code:       dto.Code,
			ApiId:      dto.ID,
			Parameters: dto.Parameters,
			Responses:  dto.Responses,
			Type:       int(category.String2RespType(dto.RespType)),
		}
	}
	return model, resp
}
