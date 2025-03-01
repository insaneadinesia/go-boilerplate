package location_svc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/insaneadinesia/gobang/logger"
)

func (w *wrapper) GetSubDistrictDetail(ctx context.Context, id int64) (resp GetSubDistrictDetailResponse, err error) {
	path := fmt.Sprintf("/sub-district/%d?name=1212", id)

	headers := http.Header{}
	headers.Add("Content-Type", "application/json")
	body, status, err := w.client.Get(ctx, path, headers)
	if err != nil {
		err = fmt.Errorf("wrapper.GetSubDistrictDetail error: %s", err.Error())
		return
	}

	if status >= http.StatusBadRequest {
		err = fmt.Errorf("wrapper.GetSubDistrictDetail return status %d", status)
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		logger.Log.Error(ctx, err.Error())
		err = fmt.Errorf("wrapper.GetSubDistrictDetail unmarshal error: %s", err.Error())
		return
	}

	return
}
