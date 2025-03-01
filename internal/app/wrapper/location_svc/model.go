package location_svc

type LocationDetail struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type GetSubDistrictDetailResponse struct {
	LocationDetail
	District GetDistrictDetailResponse `json:"district"`
}

type GetDistrictDetailResponse struct {
	LocationDetail
	City GetCityDetailResponse `json:"city"`
}

type GetCityDetailResponse struct {
	LocationDetail
	Province GetProvinceDetailResponse `json:"province"`
}

type GetProvinceDetailResponse struct {
	LocationDetail
}
