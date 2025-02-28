package user

import "github.com/insaneadinesia/go-boilerplate/internal/pkg/pagination"

type CreateUpdateUserRequest struct {
	Name          string `json:"name" validate:"required" example:"Rachmat Adi Prakoso"`
	Username      string `json:"username" validate:"required" example:"mamatosai"`
	Email         string `json:"email" validate:"required,email" example:"rachmat.adi.p@gmail.com"`
	SubDistrictID int64  `json:"sub_district_id" validate:"required,gt=10" example:"10001"`
}

type GetAllUserResponse struct {
	Users      []UserDetailResponse          `json:"users"`
	Pagination pagination.PaginationResponse `json:"pagination"`
}

type UserDetailResponse struct {
	Name      string       `json:"name" example:"Rachmat Adi Prakoso"`
	Username  string       `json:"username" example:"mamatosai"`
	Email     string       `json:"email" example:"rachmat.adi.p@gmail.com"`
	CreatedAt string       `json:"created_at" example:"2025-02-28T12:00:00+0700"`
	UpdatedAt string       `json:"updated_at" example:"2025-02-28T12:00:00+0700"`
	Location  UserLocation `json:"location,omitempty"`
}

type UserLocation struct {
	SubDistrictID   int64  `json:"sub_district_id" example:"10001"`
	SubDistrictName string `json:"sub_district_name" example:"CIMONE JAYA"`
	DistrictID      int64  `json:"district_id" example:"1001"`
	DistrictName    string `json:"district_name" example:"KARAWACI"`
	CityID          int64  `json:"city_id" example:"101"`
	CityName        string `json:"city_name" example:"KOTA TANGERANG"`
	ProvinceID      int64  `json:"province_id" example:"11"`
	ProvinceName    string `json:"province_name" example:"BANTEN"`
}
