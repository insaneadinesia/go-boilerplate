package rest

import "time"

type Option struct {
	Address string
	Timeout time.Duration
	SkipTLS bool
}
