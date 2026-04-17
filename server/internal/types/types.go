package types

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

type HTTPError struct {
	Message   string
	ErrorCode int
}

type RequestStatus struct {
	Status int `json:"status_code"`
}
