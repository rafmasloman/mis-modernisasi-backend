package dto

type FilterBuilderDTO struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Query    string `json:"query"`
	ReportId int    `json:"report_id"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	OrderNum int    `json:"order_num"`
}

type FilterBuilderResponseDTO struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Title    string `json:"title"`
	Query    string `json:"query"`
	ReportId int    `json:"report_id"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	OrderNum int    `json:"order_num"`
}
