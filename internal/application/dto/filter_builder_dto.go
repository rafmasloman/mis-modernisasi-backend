package dto

type FilterBuilderDTO struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Query    string `json:"query"`
	ReportId int    `json:"report_id"`
	Type     string `json:"type"`
	Value    string `json:"value"`
	Label    string `json:"label"`
	Api      string `json:"api"`
	ReqBody  string `json:"req_body"`
	Required bool   `json:"required"`
	OrderNum int    `json:"order_num"`
	OnLoad   bool   `json:"on_load"`
	OnChange string `json:"on_change"`
}

type FilterBuilderResponseDTO struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Title    string `json:"title"`
	Query    string `json:"query"`
	ReportId int    `json:"report_id"`
	Type     string `json:"type"`
	Value    string `json:"value"`
	Label    string `json:"label"`
	Api      string `json:"api"`
	ReqBody  string `json:"req_body"`
	Required bool   `json:"required"`
	OrderNum int    `json:"order_num"`
	OnLoad   bool   `json:"on_load"`
	OnChange string `json:"on_change"`
}
