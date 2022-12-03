package repository

import (
	"database/sql"
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

type UserTokenRepositorySuite struct {
	suite.Suite
	conn *sql.DB
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	repo *UserTokenRepository

	userToken *model.UserToken
}

func (rs *UserTokenRepositorySuite) AfterTest(_, _ string) {
	assert.NoError(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *UserTokenRepositorySuite) SetupSuite() {
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

	rs.repo = NewUserTokenRepository(rs.DB)

	assert.IsType(rs.T(), &UserTokenRepository{}, rs.repo)

	rs.userToken = buildUserToken()
}

func (rs *UserTokenRepositorySuite) TestGenerate() {
	rs.mock.ExpectBegin()

	rs.mock.ExpectExec(
		regexp.QuoteMeta(`INSERT INTO "user_tokens" ("id","token","user_id","created_at","updated_at") VALUES ($1,$2,$3,$4,$5)`),
	).WithArgs(
		rs.userToken.ID,
		rs.userToken.Token,
		rs.userToken.UserID,
		rs.userToken.CreatedAt,
		rs.userToken.UpdatedAt,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	rs.mock.ExpectCommit()

	userToken := rs.repo.Generate(rs.userToken)

	assert.Equal(rs.T(), rs.userToken, userToken)
}

func (rs *UserTokenRepositorySuite) TestFindByToken() {
	rows := sqlmock.NewRows([]string{"id", "token", "user_id", "created_at", "updated_at"}).
		AddRow(
			rs.userToken.ID,
			rs.userToken.Token,
			rs.userToken.UserID,
			rs.userToken.CreatedAt,
			rs.userToken.UpdatedAt,
		)

	rs.mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "user_tokens" WHERE token = $1`),
	).WithArgs(
		rs.userToken.Token,
	).WillReturnRows(rows)

	userToken := rs.repo.FindByToken(rs.userToken.Token.String())

	assert.Equal(rs.T(), rs.userToken, userToken)
}

func TestSuiteUserTokenRepository(t *testing.T) {
	suite.Run(t, new(UserTokenRepositorySuite))
}

func buildUserToken() *model.UserToken {
	return &model.UserToken{
		ID:        uuid.MustParse("c02bc48d-d72b-4e6c-82df-b7f55d4ae491"),
		Token:     uuid.MustParse("38bca023-4a02-4fc7-90e6-12ea70d8a2cf"),
		UserID:    "5d8b9bb7-da43-41ff-9078-fee0451fb532",
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}
}
