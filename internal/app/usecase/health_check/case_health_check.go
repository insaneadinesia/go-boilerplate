package health_check

import "context"

func (u *usecase) HealthCheck(ctx context.Context) (resp StatusCheck, err error) {
	const (
		STATUS_OK    = "OK"
		STATUS_ERROR = "ERROR"
	)

	resp.DBStatus = STATUS_OK

	err = u.healthCheckRepository.PingDB(ctx)
	if err != nil {
		resp.DBStatus = STATUS_ERROR
		return
	}

	return
}
