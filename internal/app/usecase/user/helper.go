package user

import (
	"github.com/insaneadinesia/go-boilerplate/internal/app/entity"
	"github.com/insaneadinesia/go-boilerplate/internal/app/wrapper/location_svc"
)

func (u *usecase) mappingUserResponse(user entity.User, location *location_svc.GetSubDistrictDetailResponse) (result UserDetailResponse) {
	result = UserDetailResponse{
		UserResponse: UserResponse{
			UUID:      user.UUID.String(),
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z0700"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05Z0700"),
		},
	}

	if location != nil {
		result.Location = &UserLocation{
			SubDistrictID:   location.ID,
			SubDistrictName: location.Name,
			DistrictID:      location.District.ID,
			DistrictName:    location.District.Name,
			CityID:          location.District.City.ID,
			CityName:        location.District.City.Name,
			ProvinceID:      location.District.City.Province.ID,
			ProvinceName:    location.District.City.Province.Name,
		}
	}

	return
}
