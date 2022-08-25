package common_function

type success struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func SuccessResponse(data, paging, filter interface{}) *success {
	return &success{
		Data:   data,
		Paging: paging,
		Filter: filter,
	}
}

func DataResponseOnly(data interface{}) *success {
	return &success{
		Data:   data,
		Paging: nil,
		Filter: nil,
	}
}
