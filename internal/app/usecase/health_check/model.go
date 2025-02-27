package health_check

type StatusCheck struct {
	DBStatus string `json:"db_status" example:"OK/ERROR"`
}
