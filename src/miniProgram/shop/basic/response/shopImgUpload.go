package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseShopImgUpload struct {
	*response.ResponseMiniProgram

	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	ImgInfo *struct {
		ErrCode    string `json:"media_id,omitempty"`     // 	开放平台media_id
		TempImgUrl string `json:"temp_img_url"`           // 图片链接，有访问频率限制，不可用作图床，后果自负
		PayMediaId string `json:"pay_media_id,omitempty"` //	微信支付media_id
	} `json:"img_info"` // 图片信息
}
