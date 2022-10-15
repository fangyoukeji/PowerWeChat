package request

type RequestAuditCategory struct {
	AuditCategoryReq AuditCategoryReq `json:"audit_req"`
}

type AuditCategoryReq struct {
	License        []string     `json:"license"`
	CategoryInfo   CategoryInfo `json:"category_info"`
	SceneGroupList []int        `json:"scene_group_list"`
}

type CategoryInfo struct {
	Level1      int      `json:"level1"`
	Level2      int      `json:"level2"`
	Level3      int      `json:"level3"`
	Certificate []string `json:"certificate"`
}
