package embalm

import (
	"encoding/json"
	"github.com/isyscore/isc-gobase/config"
	"github.com/isyscore/isc-gobase/http"
	"github.com/rs/zerolog/log"
	http2 "net/http"
	"strings"
)

type UserStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    data   `json:"data"`
}

type data struct {
	Token      string   `json:"token"`
	UserId     string   `json:"userId"`
	LoginName  string   `json:"loginName"`
	NickName   string   `json:"nickName"`
	Role       []string `json:"role"`
	RoleId     []string `json:"roleId"`
	TenantId   string   `json:"tenantId"`
	SuperAdmin bool     `json:"superAdmin"`
}

var statusUri = "/api/permission/auth/status"
var defaultHost = "http://isc-permission-service:32100"

func GetUserStatus(token string) UserStatus {
	status := UserStatus{
		Data: data{},
	}
	defer func() UserStatus {
		if x := recover(); x != nil {
			log.Error().Msgf("获取用户数据异常，%v", x)
		}
		return status
	}()

	header := http2.Header{
		"token": []string{token},
		//"content-type": []string{"application/json"},
	}
	host := config.GetValueStringDefault("feign.permission", defaultHost)
	url := func(h string) string {
		h = h + statusUri
		h = strings.ReplaceAll(h, "//api", "/api")
		return h
	}(host)
	code, _, result, err := http.Get(url, header, make(map[string]string))
	if code != http2.StatusOK {
		log.Error().Msgf("不能正确地获取到用户状态,code = %d", code)
		return status
	}
	if err != nil {
		log.Error().Msgf("获取用户状态时发生了异常,%v", err)
		return status
	}

	if err = json.Unmarshal(result.([]byte), &status); err != nil {
		log.Error().Msgf("用户状态信息解析异常,%v", err)
		return status
	}
	return status
}
