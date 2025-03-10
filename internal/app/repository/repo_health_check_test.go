package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserHealthCheckTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo HealthCheck
}

func (s *UserHealthCheckTestSuite) SetupTest() {
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

	s.repo = NewHealthCheckRepository(s.db)
}

func (s *UserHealthCheckTestSuite) TestPingDB_Success() {
	err := s.repo.PingDB(context.Background())

	assert.NoError(s.T(), err)
	assert.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestHealthCheckRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserHealthCheckTestSuite))
}

// Additional Test For Panic Test
func TestPanic_NewHealthCheckRepository(t *testing.T) {
	assert.Panics(t, func() {
		NewHealthCheckRepository(nil)
	})
}
