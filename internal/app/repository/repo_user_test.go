package repository

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/insaneadinesia/go-boilerplate/internal/app/entity"
)

type UserRepoTestSuite struct {
	suite.Suite
	db       *gorm.DB
	mock     sqlmock.Sqlmock
	repo     User
	testUser entity.User
}

func (s *UserRepoTestSuite) SetupTest() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	assert.NoError(s.T(), err)

	dialector := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})
	s.db, err = gorm.Open(dialector)
	assert.NoError(s.T(), err)

	s.repo = NewUserRepository(s.db)

	// Setup test user
	s.testUser = entity.User{
		UUID:          uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
		Name:          "Test User",
		Username:      "testuser",
		Email:         "test@example.com",
		SubDistrictID: 10001,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

func (s *UserRepoTestSuite) TestGetByUUID_Success() {
	rows := sqlmock.NewRows([]string{"uuid", "name", "username", "email", "created_at", "updated_at", "deleted_at"}).
		AddRow(s.testUser.UUID, s.testUser.Name, s.testUser.Username, s.testUser.Email, s.testUser.CreatedAt, s.testUser.UpdatedAt, nil)

	s.mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."uuid" = \$1 AND "users"."deleted_at" IS NULL ORDER BY "users"."uuid" DESC LIMIT \$2`).
		WithArgs(s.testUser.UUID, 1).
		WillReturnRows(rows)

	user, err := s.repo.GetByUUID(context.Background(), s.testUser.UUID)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.testUser.UUID, user.UUID)
	assert.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *UserRepoTestSuite) TestGetByUUID_NotFound() {
	s.mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."uuid" = \$1 AND "users"."deleted_at" IS NULL ORDER BY "users"."uuid" DESC LIMIT \$2`).
		WithArgs(s.testUser.UUID, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	_, err := s.repo.GetByUUID(context.Background(), s.testUser.UUID)

	assert.Error(s.T(), err)
	assert.True(s.T(), errors.Is(err, gorm.ErrRecordNotFound))
	assert.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *UserRepoTestSuite) TestGetByUsernameOrEmail_Success() {
	rows := sqlmock.NewRows([]string{"uuid", "name", "username", "email", "created_at", "updated_at", "deleted_at"}).
		AddRow(s.testUser.UUID, s.testUser.Name, s.testUser.Username, s.testUser.Email, s.testUser.CreatedAt, s.testUser.UpdatedAt, nil)

	s.mock.ExpectQuery(`SELECT \* FROM "users" WHERE \(username = \$1 OR email = \$2\) AND "users"."deleted_at" IS NULL LIMIT \$3`).
		WithArgs("testuser", "test@example.com", 1).
		WillReturnRows(rows)

	user, err := s.repo.GetByUsernameOrEmail(context.Background(), "testuser", "test@example.com")

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.testUser.UUID, user.UUID)
	assert.Equal(s.T(), s.testUser.Name, user.Name)
	assert.Equal(s.T(), s.testUser.Username, user.Username)
	assert.Equal(s.T(), s.testUser.Email, user.Email)
	assert.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *UserRepoTestSuite) TestGetByUsernameOrEmail_NotFound() {
	s.mock.ExpectQuery(`SELECT \* FROM "users" WHERE \(username = \$1 OR email = \$2\) AND "users"."deleted_at" IS NULL LIMIT \$3`).
		WithArgs("nonexistent", "nonexistent@example.com", 1).
		WillReturnError(gorm.ErrRecordNotFound)

	_, err := s.repo.GetByUsernameOrEmail(context.Background(), "nonexistent", "nonexistent@example.com")

	assert.Error(s.T(), err)
	assert.True(s.T(), errors.Is(err, gorm.ErrRecordNotFound))
	assert.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *UserRepoTestSuite) TestGetAll_Success() {
	rows := sqlmock.NewRows([]string{"uuid", "name", "username", "email", "created_at", "updated_at", "deleted_at"}).
		AddRow(s.testUser.UUID, s.testUser.Name, s.testUser.Username, s.testUser.Email, s.testUser.CreatedAt, s.testUser.UpdatedAt, nil)

	s.mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."deleted_at" IS NULL LIMIT \$1 OFFSET \$2`).
		WithArgs(10, 1).
		WillReturnRows(rows)

	users, err := s.repo.GetAll(context.Background(), UserFilter{
		Limit:  10,
		Offset: 1,
	})

	assert.NoError(s.T(), err)
	assert.Len(s.T(), users, 1)
	assert.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *UserRepoTestSuite) TestGetAll_Success_WithFilterName() {
	rows := sqlmock.NewRows([]string{"uuid", "name", "username", "email", "created_at", "updated_at", "deleted_at"}).
		AddRow(s.testUser.UUID, s.testUser.Name, s.testUser.Username, s.testUser.Email, s.testUser.CreatedAt, s.testUser.UpdatedAt, nil)

	s.mock.ExpectQuery(`SELECT \* FROM "users" WHERE name ILIKE \$1\ AND "users"."deleted_at" IS NULL LIMIT \$2 OFFSET \$3`).
		WithArgs("%test%", 10, 1).
		WillReturnRows(rows)

	users, err := s.repo.GetAll(context.Background(), UserFilter{
		Name:   "test",
		Limit:  10,
		Offset: 1,
	})

	assert.NoError(s.T(), err)
	assert.Len(s.T(), users, 1)
	assert.Equal(s.T(), users[0].Name, "Test User")
	assert.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *UserRepoTestSuite) TestGetAll_Success_WithFilterEmail() {
	rows := sqlmock.NewRows([]string{"uuid", "name", "username", "email", "created_at", "updated_at", "deleted_at"}).
		AddRow(s.testUser.UUID, s.testUser.Name, s.testUser.Username, s.testUser.Email, s.testUser.CreatedAt, s.testUser.UpdatedAt, nil)

	s.mock.ExpectQuery(`SELECT \* FROM "users" WHERE email = \$1 AND "users"."deleted_at" IS NULL LIMIT \$2 OFFSET \$3`).
		WithArgs("test@example.com", 10, 1).
		WillReturnRows(rows)

	users, err := s.repo.GetAll(context.Background(), UserFilter{
		Email:  "test@example.com",
		Limit:  10,
		Offset: 1,
	})

	assert.NoError(s.T(), err)
	assert.Len(s.T(), users, 1)
	assert.Equal(s.T(), users[0].Email, "test@example.com")
	assert.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *UserRepoTestSuite) TestGetAll_Success_WithFilterUsername() {
	rows := sqlmock.NewRows([]string{"uuid", "name", "username", "email", "created_at", "updated_at", "deleted_at"}).
		AddRow(s.testUser.UUID, s.testUser.Name, s.testUser.Username, s.testUser.Email, s.testUser.CreatedAt, s.testUser.UpdatedAt, nil)

	s.mock.ExpectQuery(`SELECT \* FROM "users" WHERE username = \$1 AND "users"."deleted_at" IS NULL LIMIT \$2 OFFSET \$3`).
		WithArgs("testuser", 10, 1).
		WillReturnRows(rows)

	users, err := s.repo.GetAll(context.Background(), UserFilter{
		Username: "testuser",
		Limit:    10,
		Offset:   1,
	})

	assert.NoError(s.T(), err)
	assert.Len(s.T(), users, 1)
	assert.Equal(s.T(), users[0].Username, "testuser")
	assert.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *UserRepoTestSuite) TestGetAll_NoResults() {
	s.mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."deleted_at" IS NULL LIMIT \$1`).
		WithArgs(10).
		WillReturnRows(sqlmock.NewRows([]string{}))

	users, err := s.repo.GetAll(context.Background(), UserFilter{
		Limit:  10,
		Offset: 0,
	})

	assert.NoError(s.T(), err)
	assert.Empty(s.T(), users)
	assert.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *UserRepoTestSuite) TestCountTotal_Success() {
	s.mock.ExpectQuery(`SELECT count\(\*\) FROM "users" WHERE "users"."deleted_at" IS NULL`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(10))

	total, err := s.repo.CountTotal(context.Background(), UserFilter{})

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), int64(10), total)
	assert.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *UserRepoTestSuite) TestCreate_Success() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(`INSERT INTO "users"`).
		WithArgs(
			sqlmock.AnyArg(), // UUID is generated in BeforeCreate hook
			s.testUser.Name,
			s.testUser.Username,
			s.testUser.Email,
			s.testUser.SubDistrictID,
			sqlmock.AnyArg(), // created_at
			sqlmock.AnyArg(), // updated_at
			nil,              // deleted_at
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.repo.Create(context.Background(), &s.testUser)

	assert.NoError(s.T(), err)
	assert.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *UserRepoTestSuite) TestUpdate_Success() {
	s.testUser.Name = "Updated Name"

	s.mock.ExpectBegin()
	s.mock.ExpectExec(`UPDATE "users" SET "name"=\$1,"username"=\$2,"email"=\$3,"sub_district_id"=\$4,"updated_at"=\$5,"deleted_at"=\$6 WHERE "users"."deleted_at" IS NULL AND "uuid" = \$7`).
		WithArgs(
			s.testUser.Name,
			s.testUser.Username,
			s.testUser.Email,
			s.testUser.SubDistrictID,
			sqlmock.AnyArg(), // updated_at
			nil,              // deleted_at
			s.testUser.UUID,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.repo.Update(context.Background(), &s.testUser)

	assert.NoError(s.T(), err)
	assert.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *UserRepoTestSuite) TestDelete_Success() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(`UPDATE "users" SET "deleted_at"=\$1 WHERE "users"."uuid" = \$2 AND "users"."deleted_at" IS NULL`).
		WithArgs(sqlmock.AnyArg(), s.testUser.UUID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.repo.Delete(context.Background(), &s.testUser)

	assert.NoError(s.T(), err)
	assert.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}

// Additional Test For Panic Test
func TestPanic_NewUserRepository(t *testing.T) {
	assert.Panics(t, func() {
		NewUserRepository(nil)
	})
}
