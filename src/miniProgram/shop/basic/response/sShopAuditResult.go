package response

type ResponseShopAuditResult struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Data    *struct {
		Status       int    `json:"status"`
		BrandID      int    `json:"brand_id"`
		RejectReason string `json:"reject_reason"`
	} `json:"data"`
}
