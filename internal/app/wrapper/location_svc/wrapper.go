package location_svc

import (
	"context"

	"github.com/insaneadinesia/go-boilerplate/config"
	"github.com/insaneadinesia/gobang/rest"
)

type LocationScvWrapper interface {
	GetSubDistrictDetail(ctx context.Context, id int64) (resp GetSubDistrictDetailResponse, err error)
}

type wrapper struct {
	config config.Config
	client rest.RestClient
}

func NewWrapper() *wrapper {
	return &wrapper{}
}

func (w *wrapper) SetConfig(config config.Config) *wrapper {
	w.config = config
	return w
}

func (w *wrapper) Setup() *wrapper {
	restOption := rest.Option{
		Address: w.config.LocationSvcUrl,
		Timeout: w.config.LocationSvcTimeout,
		SkipTLS: w.config.LocationSvcSkipTLS,
	}

	w.client = rest.New(restOption)
	return w
}

func (w *wrapper) Validate() LocationScvWrapper {
	if w.client == nil {
		panic("client is nil")
	}

	return w
}
