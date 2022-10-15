/*
 * @Author: fangyoukeji
 * @Date: 2022-10-14 13:48:25
 * @FilePath: /PowerWeChat/src/miniProgram/shop/register/client.go
 * @Description:自定义版交易组件及开放接口 - 申请接入接口
 * Copyright (c) 2022 by ArtisanCloud.Ltd., All Rights Reserved.
 */
package register

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/shop/register/response"
)

type Client struct {
	*kernel.BaseClient
}

/**
 * @description: 自定义组件 - 接入申请
 * @return {*}
 * @Author: fangyoukeji
 * @url: https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/ministore/minishopopencomponent2/API/enter/enter_apply.html
 */
func (comp *Client) ShopRegisterApply() (*response.ResponseShopRegisterApply, error) {
	result := &response.ResponseShopRegisterApply{}
	_, err := comp.HttpPostJson("shop/register/apply", nil, nil, nil, result)
	return result, err
}

/**
 * @description: 自定义组件 - 获取接入状态
 * @return {*}
 * @Author: fangyoukeji
 * @url: https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/ministore/minishopopencomponent2/API/enter/enter_check.html
 */
func (comp *Client) ShopRegisterCheck() (*response.ResponseShopRegisterCheck, error) {
	result := &response.ResponseShopRegisterCheck{}
	_, err := comp.HttpPostJson("shop/register/check", nil, nil, nil, result)
	return result, err
}

/**
 * @description: 自定义组件 - 完成接入任务
 * @param {int} accessInfoItem 6:完成 spu 接口，7:完成订单接口 / 19:完成二级商户号订单，8:完成物流接口，9:完成售后接口 / 20:完成二级商户号售后，10:测试完成，11:发版完成
 * @return {*}
 * @Author: fangyoukeji
 * @url: https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/ministore/minishopopencomponent2/API/enter/finish_access_info.html
 */
func (comp *Client) ShopRegisterFinishAccessInfo(accessInfoItem int) (*response.ResponseShopRegisterFinishAccessInfo, error) {
	result := &response.ResponseShopRegisterFinishAccessInfo{}

	data := &object.HashMap{
		"access_info_item": accessInfoItem,
	}

	_, err := comp.HttpPostJson("shop/register/finish_access_info", data, nil, nil, result)
	return result, err
}

/**
 * @description: 自定义组件 - 场景接入申请
 * @param {int} sceneGroupId 1:视频号
 * @return {*}
 * @Author: fangyoukeji
 * @url: https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/ministore/minishopopencomponent2/API/enter/scene_apply.html
 */
func (comp *Client) ShopRegisterApplyScene(sceneGroupId int) (*response.ResponseShopRegisterApplyScene, error) {
	result := &response.ResponseShopRegisterApplyScene{}

	data := &object.HashMap{
		"scene_group_id": sceneGroupId,
	}

	_, err := comp.HttpPostJson("shop/register/apply_scene", data, nil, nil, result)
	return result, err
}
