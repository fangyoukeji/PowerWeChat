package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

/**
 * @description: 商品类目返回结构体
 * @Author: fangyoukeji
 */
type ResponseShopCatGet struct {
	*response.ResponseMiniProgram

	ErrCode      int                 `json:"errcode"`
	ErrMsg       string              `json:"errmsg"`
	ThirdCatList []*ThirdCatListItem `json:"third_cat_list"`
}

/**
 * @description: 类目列表
 * @Author: fangyoukeji
 */
type ThirdCatListItem struct {
	ThirdCatId               int    `json:"third_cat_id"`
	ThirdCatName             string `json:"third_cat_name"`
	Qualification            string `json:"qualification"`
	QualificationType        int    `json:"qualification_type"`
	ProductQualification     string `json:"product_qualification"`
	ProductQualificationType int    `json:"product_qualification_type"`
	SecondCatId              int    `json:"second_cat_id"`
	SecondCatName            string `json:"second_cat_name"`
	FirstCatId               int    `json:"first_cat_id"`
	FirstCatName             string `json:"first_cat_name"`
}
