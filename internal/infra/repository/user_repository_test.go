package repository

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepositorySuite struct {
	suite.Suite
	conn *sql.DB
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	repo *UserRepository

	user *model.User
}

func (rs *UserRepositorySuite) AfterTest(_, _ string) {
	assert.NoError(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *UserRepositorySuite) SetupSuite() {
	var err error

	rs.conn, rs.mock, err = sqlmock.New()
	assert.NoError(rs.T(), err)

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 rs.conn,
		PreferSimpleProtocol: true,
	})

	rs.DB, err = gorm.Open(dialector, &gorm.Config{})
	assert.NoError(rs.T(), err)

	rs.DB.Callback().Create().Remove("gorm:before_create")

	rs.repo = NewUserRepository(rs.DB)

	assert.IsType(rs.T(), &UserRepository{}, rs.repo)

	rs.user = buildUser()
}

func (rs *UserRepositorySuite) TestSave() {
	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(
		regexp.QuoteMeta(`INSERT INTO "users" ("id","name","email","password","avatar","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6,$7)`),
	).WithArgs(
		rs.user.ID,
		rs.user.Name,
		rs.user.Email,
		rs.user.Password,
		"",
		rs.user.CreatedAt,
		rs.user.UpdatedAt,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	rs.mock.ExpectCommit()

	rs.repo.Save(rs.user)
}

func (rs *UserRepositorySuite) TestUpdate() {
	rs.user.Name = "LUCAS UPDATE"

	sql := fmt.Sprint(`UPDATE "users" SET "name"=$1,"email"=$2,"password"=$3,"avatar"=$4,"created_at"=$5,"updated_at"=$6 WHERE "id" = $7`)

	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(
		regexp.QuoteMeta(sql),
	).WithArgs(
		rs.user.Name,
		rs.user.Email,
		rs.user.Password,
		"",
		AnyTime{},
		AnyTime{},
		rs.user.ID,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	rs.mock.ExpectCommit()

	rs.repo.Update(rs.user)
}

func (rs *UserRepositorySuite) TestFindByEmail() {
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "avatar", "created_at", "updated_at"}).
		AddRow(
			rs.user.ID,
			rs.user.Name,
			rs.user.Email,
			rs.user.Password,
			rs.user.Avatar,
			rs.user.CreatedAt,
			rs.user.UpdatedAt,
		)

	rs.mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1`),
	).WithArgs(
		rs.user.Email,
	).WillReturnRows(rows)

	user := rs.repo.FindByEmail(rs.user.Email)

	assert.Equal(rs.T(), rs.user.ID, user.ID)
	assert.Equal(rs.T(), rs.user.Name, user.Name)
	assert.Equal(rs.T(), rs.user.Email, user.Email)
}

func (rs *UserRepositorySuite) TestFindById() {
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "avatar", "created_at", "updated_at"}).
		AddRow(
			rs.user.ID,
			rs.user.Name,
			rs.user.Email,
			rs.user.Password,
			rs.user.Avatar,
			rs.user.CreatedAt,
			rs.user.UpdatedAt,
		)

	rs.mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "users" WHERE id = $1`),
	).WithArgs(
		rs.user.ID,
	).WillReturnRows(rows)

	user := rs.repo.FindById(rs.user.ID.String())

	assert.Equal(rs.T(), rs.user.ID, user.ID)
	assert.Equal(rs.T(), rs.user.Name, user.Name)
	assert.Equal(rs.T(), rs.user.Email, user.Email)
}

func TestSuiteUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}

func buildUser() *model.User {
	return &model.User{
		ID:        uuid.MustParse("13a1d230-f016-4450-a714-3c4475d32a7f"),
		Name:      "Lucas",
		Email:     "lucas@gmail.com",
		Password:  "123456",
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
