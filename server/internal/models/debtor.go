package models

type Debtor struct {
	Id       int    `json:"debtor_id"`
	Inn      string `json:"debtor_inn"`
	Ogrnip   string `json:"debtor_ogrnip"`
	Name     string `json:"debtor_name"`
	Category string `json:"debtor_category"`
	Snils    string `json:"debtor_snils"`
	Region   string `json:"debtor_region"`
	Address  string `json:"debtor_address"`
}
