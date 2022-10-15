package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

/**
 * @description: 接入申请
 * @return {*}
 */
type ResponseShopRegisterApply struct {
	*response.ResponseMiniProgram

	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

/**
 * @description:获取接入状态
 * @return {*}
 */
type ResponseShopRegisterCheck struct {
	*response.ResponseMiniProgram

	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Data    *struct {
		AccessInfo *struct {
			SpuAuditSuccess         int `json:"spu_audit_success"`          // 上传商品并审核成功
			EcOrderSuccess          int `json:"ec_order_success"`           // 发起一笔订单并支付成功
			SendDeliverySuccess     int `json:"send_delivery_success"`      // 物流接口调用成功
			EcAfterSaleSuccess      int `json:"ec_after_sale_success"`      // 售后接口调用成功
			SpuAuditFinished        int `json:"spu_audit_finished"`         // 商品接口调试完成
			EcOrderFinished         int `json:"ec_order_finished"`          // 订单接口调试完成
			SendDeliveryFinished    int `json:"send_delivery_finished"`     // 物流接口调试完成
			EcAfterSaleFinished     int `json:"ec_after_sale_finished"`     // 售后接口调试完成
			TestApiFinished         int `json:"test_api_finished"`          // 测试完成
			DeployWxaFinished       int `json:"deploy_wxa_finished"`        // 发版完成
			OpenProductTaskFinished int `json:"open_product_task_finished"` // 完成自定义组件全部任务
		} `json:"access_info"`
		SceneGroupList []*power.HashMap `json:"scene_group_list"` // 场景接入相关
	} `json:"data,omitempty"`
}

/**
 * @description: 完成接入任务
 * @return {*}
 */
type ResponseShopRegisterFinishAccessInfo struct {
	*response.ResponseMiniProgram

	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

/**
 * @description: 场景接入申请
 * @return {*}
 */
type ResponseShopRegisterApplyScene struct {
	*response.ResponseMiniProgram

	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
