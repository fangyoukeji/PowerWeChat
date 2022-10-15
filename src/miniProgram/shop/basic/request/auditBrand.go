package request

type RequestAuditBrand struct {
	AuditBrandReq AuditBrandReq `json:"audit_req"`
}

type AuditBrandReq struct {
	License        []string  `json:"license"`
	BrandInfo      BrandInfo `json:"brand_info"`
	SceneGroupList []int     `json:"scene_group_list"`
}

type BrandInfo struct {
	BrandAuditType                   int      `json:"brand_audit_type"`
	TrademarkType                    string   `json:"trademark_type"`
	BrandManagementType              int      `json:"brand_management_type"`
	CommodityOriginType              int      `json:"commodity_origin_type"`
	BrandWording                     string   `json:"brand_wording"`
	SaleAuthorization                []string `json:"sale_authorization"`
	TrademarkRegistrationCertificate []string `json:"trademark_registration_certificate"`
	TrademarkChangeCertificate       []string `json:"trademark_change_certificate"`
	TrademarkRegistrant              string   `json:"trademark_registrant"`
	TrademarkRegistrantNu            string   `json:"trademark_registrant_nu"`
	TrademarkAuthorizationPeriod     string   `json:"trademark_authorization_period"`
	TrademarkRegistrationApplication []string `json:"trademark_registration_application"`
	TrademarkApplicant               string   `json:"trademark_applicant"`
	TrademarkApplicationTime         string   `json:"trademark_application_time"`
	ImportedGoodsForm                []string `json:"imported_goods_form"`
}
