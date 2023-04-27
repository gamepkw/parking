package domain

type Fee struct {
	Feeid       int64  `json:"id"`
	Config_name string `json:"config_name"`
	Value       string `json:"value"`
}
