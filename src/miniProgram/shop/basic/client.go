/*
 * @Author: fangyoukeji
 * @Date: 2022-10-14 17:02:00
 * @FilePath: /PowerWeChat/src/miniProgram/shop/basic/client.go
 * @Description:自定义版交易组件及开放接口 - 接入商品前必需接口
 * Copyright (c) 2022 by ArtisanCloud.Ltd., All Rights Reserved.
 */
package basic

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/shop/basic/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/shop/basic/response"
)

type Client struct {
	*kernel.BaseClient
}

/**
 * @Author: fangyoukeji
 * @description: 自定义组件 - 获取商品类目
 * @return {*}
 * @url: https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/ministore/minishopopencomponent2/API/cat/get_children_cateogry.html
 */
func (comp *Client) ShopGetCat() (*response.ResponseShopCatGet, error) {
	result := &response.ResponseShopCatGet{}
	_, err := comp.HttpPostJson("shop/cat/get", nil, nil, nil, result)
	return result, err
}

/**
 * @Author: fangyoukeji
 * @description:自定义组件 - 上传图片
 * @param {string} imageFilePath uploadType=0为图片文件路径 uploadType=1为网络图片url
 * @param {int} respType 响应格式 0:media_id(目前只用于品牌申请品牌和类目)，1：返回链接, 2：返回微信支付media_id
 * @param {int} uploadType 上传类型，本地图片或者网络图片
 * @return {*}
 * @url: https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/ministore/minishopopencomponent2/API/public/upload_img.html
 */
func (comp *Client) ShopImgUpload(imageFilePath string, respType int, uploadType int) (*response.ResponseShopImgUpload, error) {
	result := &response.ResponseShopImgUpload{}
	if uploadType == 0 {
		// 上传图片为二进制流
		var files *object.HashMap
		if imageFilePath != "" {
			files = &object.HashMap{
				"media": imageFilePath,
			}
		}
		_, err := comp.HttpUpload("shop/img/upload", files, nil, &object.HashMap{
			"resp_type":   respType,
			"upload_type": uploadType,
		}, nil, result)
		return result, err
	} else {
		// 上传图片为图片url
		data := &object.HashMap{
			"img_url":     imageFilePath,
			"resp_type":   respType,
			"upload_type": uploadType,
		}
		_, err := comp.HttpPostJson("shop/img/upload", data, nil, nil, result)
		return result, err
	}
}

/**
 * @Author: fangyoukeji
 * @description: 自定义组件 - 上传品牌审核
 * @param {*request.RequestAuditBrand} options
 * @return {*}
 * @url: https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/ministore/minishopopencomponent2/API/audit/audit_brand.html
 */
func (comp *Client) ShopAuditBrand(options *request.RequestAuditBrand) (*response.ResponseShopAuditBrand, error) {
	result := &response.ResponseShopAuditBrand{}

	_, err := comp.HttpPostJson("shop/audit/audit_brand", options, nil, nil, result)
	return result, err
}

/**
 * @Author: fangyoukeji
 * @description: 自定义组件 - 上传类目资质审核
 * @param {*request.RequestAuditCategory} options
 * @return {*}
 * @url: https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/ministore/minishopopencomponent2/API/audit/audit_category.html
 */
func (comp *Client) ShopAuditCategory(options *request.RequestAuditCategory) (*response.ResponseShopAuditCategory, error) {
	result := &response.ResponseShopAuditCategory{}

	_, err := comp.HttpPostJson("shop/audit/audit_category", options, nil, nil, result)
	return result, err
}

/**
 * @Author: fangyoukeji
 * @description: 自定义组件 - 查询获取审核结果
 * @param {string} auditId
 * @return {*}
 * @url: https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/ministore/minishopopencomponent2/API/audit/audit_result.html
 */
func (comp *Client) ShopAuditResult(auditId string) (*response.ResponseShopAuditResult, error) {
	result := &response.ResponseShopAuditResult{}
	data := &object.HashMap{
		"audit_id": auditId,
	}
	_, err := comp.HttpPostJson("shop/audit/result", data, nil, nil, result)
	return result, err
}

/**
 * @Author: fangyoukeji
 * @description: 自定义组件 - 获取小程序以往提交资质
 * @param {string} reqType 1-类目 2-品牌
 * @return {*}
 * @url: https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/business-capabilities/ministore/minishopopencomponent2/API/audit/audit_result.html
 */
func (comp *Client) ShopGetMiniAppCertificate(reqType int) (*response.ResponseShopGetMiniAppCertificate, error) {
	result := &response.ResponseShopGetMiniAppCertificate{}
	data := &object.HashMap{
		"req_type": reqType,
	}
	_, err := comp.HttpPostJson("shop/audit/get_miniapp_certificate", data, nil, nil, result)
	return result, err
}
