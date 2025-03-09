package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/insaneadinesia/go-boilerplate/internal/app/entity"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/helper"
	"github.com/insaneadinesia/gobang/gotel"
	"gorm.io/gorm"
)

type User interface {
	GetByUUID(ctx context.Context, uuid uuid.UUID) (entity entity.User, err error)
	GetByUsernameOrEmail(ctx context.Context, username, email string) (entity entity.User, err error)
	GetAll(ctx context.Context, filter UserFilter) (entities []entity.User, err error)
	CountTotal(ctx context.Context, filter UserFilter) (total int64, err error)
	Create(ctx context.Context, entity *entity.User) (err error)
	Update(ctx context.Context, entity *entity.User) (err error)
	Delete(ctx context.Context, entity *entity.User) (err error)
}

type UserFilter struct {
	Name     string
	Username string
	Email    string
	Limit    int
	Offset   int
}

type user struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) User {
	if db == nil {
		panic("database is nil")
	}

	return &user{db}
}

func (u *user) GetByUUID(ctx context.Context, uuid uuid.UUID) (entity entity.User, err error) {
	ctx, span := gotel.DefaultTracer().Start(ctx, helper.GetFuncName())
	defer span.End()

	err = u.db.WithContext(ctx).Last(&entity, uuid).Error
	return
}

func (u *user) GetByUsernameOrEmail(ctx context.Context, username, email string) (entity entity.User, err error) {
	ctx, span := gotel.DefaultTracer().Start(ctx, helper.GetFuncName())
	defer span.End()

	err = u.db.WithContext(ctx).Where("username = ? OR email = ?", username, email).Take(&entity).Error
	return
}

func (u *user) GetAll(ctx context.Context, filter UserFilter) (entities []entity.User, err error) {
	ctx, span := gotel.DefaultTracer().Start(ctx, helper.GetFuncName())
	defer span.End()

	err = u.buildQuery(ctx, filter).Limit(filter.Limit).Offset(filter.Offset).Find(&entities).Error
	return
}

func (u *user) CountTotal(ctx context.Context, filter UserFilter) (total int64, err error) {
	ctx, span := gotel.DefaultTracer().Start(ctx, helper.GetFuncName())
	defer span.End()

	err = u.buildQuery(ctx, filter).Model(&entity.User{}).Count(&total).Error
	return
}

func (u *user) Create(ctx context.Context, entity *entity.User) (err error) {
	ctx, span := gotel.DefaultTracer().Start(ctx, helper.GetFuncName())
	defer span.End()

	err = u.db.WithContext(ctx).Create(entity).Error
	return
}

func (u *user) Update(ctx context.Context, entity *entity.User) (err error) {
	ctx, span := gotel.DefaultTracer().Start(ctx, helper.GetFuncName())
	defer span.End()

	err = u.db.WithContext(ctx).Save(entity).Error
	return
}

func (u *user) Delete(ctx context.Context, entity *entity.User) (err error) {
	ctx, span := gotel.DefaultTracer().Start(ctx, helper.GetFuncName())
	defer span.End()

	err = u.db.WithContext(ctx).Delete(entity).Error
	return
}

func (u *user) buildQuery(ctx context.Context, filter UserFilter) (db *gorm.DB) {
	db = u.db.WithContext(ctx)

	if filter.Name != "" {
		db = db.Where("name ILIKE ?", "%"+filter.Name+"%")
	}

	if filter.Email != "" {
		db = db.Where("email = ?", filter.Email)
	}

	if filter.Username != "" {
		db = db.Where("username = ?", filter.Username)
	}

	return
}
