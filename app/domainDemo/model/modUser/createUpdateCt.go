package modUser

// CreateUpdateCt 增加/更新
type CreateUpdateCt struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Name        string `json:"name" binding:"required" label:"名称"`
	NameFl      string `json:"name_fl"`
	NameFull    string `json:"name_full"`
	Sort        int64  `json:"sort"`
	State       int8   `json:"state"`
}
