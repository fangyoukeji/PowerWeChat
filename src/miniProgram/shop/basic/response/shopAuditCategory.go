package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type ResponseShopAuditCategory struct {
	*response.ResponseMiniProgram

	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	AuditId string `json:"audit_id"`
}
