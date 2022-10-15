package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type ResponseShopGetMiniAppCertificate struct {
	*response.ResponseMiniProgram
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	BrandInfo *struct {
		BrandWording                     string   `json:"brand_wording"`
		SaleAuthorization                []string `json:"sale_authorization"`
		TrademarkRegistrationCertificate []string `json:"trademark_registration_certificate"`
	} `json:"brand_info"`
	CategoryInfoList *struct {
		FirstCategoryID    int      `json:"first_category_id"`
		FirstCategoryName  string   `json:"first_category_name"`
		SecondCategoryID   int      `json:"second_category_id"`
		SecondCategoryName string   `json:"second_category_name"`
		CertificateURL     []string `json:"certificate_url"`
	} `json:"category_info_list"`
}
