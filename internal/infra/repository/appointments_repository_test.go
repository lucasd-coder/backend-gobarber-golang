package repository

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AppointmentsRepositorySuite struct {
	suite.Suite
	conn *sql.DB
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	repo *AppointmentsRepository

	appointment *model.Appointment

	dtoFindAllInMonthFromProviderDTO *dtos.FindAllInMonthFromProviderDTO

	dtoFindAllInDayFromProviderDTO *dtos.FindAllInDayFromProviderDTO
}

func (rs *AppointmentsRepositorySuite) AfterTest(_, _ string) {
	assert.NoError(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *AppointmentsRepositorySuite) SetupSuite() {
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

	rs.repo = NewAppointmentsRepository((rs.DB))

	assert.IsType(rs.T(), &AppointmentsRepository{}, rs.repo)

	rs.appointment = buildAppointments()

	rs.dtoFindAllInMonthFromProviderDTO = buildFindAllInMonthFromProviderDTO()

	rs.dtoFindAllInDayFromProviderDTO = buildFindAllInDayFromProviderDTO()
}

func (rs *AppointmentsRepositorySuite) TestSave() {
	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(
		regexp.QuoteMeta(`INSERT INTO "appointments" ("id","user_id","provider_id","date","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6)`),
	).WithArgs(
		rs.appointment.ID,
		rs.appointment.UserID,
		rs.appointment.ProviderID,
		rs.appointment.Date,
		rs.appointment.CreatedAt,
		rs.appointment.UpdatedAt,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	rs.mock.ExpectCommit()

	rs.repo.Save(rs.appointment)
}

func (rs *AppointmentsRepositorySuite) TestFindByDate() {
	rows := sqlmock.NewRows([]string{"id", "date", "provider_id", "user_id", "created_at", "updated_at"}).
		AddRow(
			rs.appointment.ID,
			rs.appointment.Date,
			rs.appointment.ProviderID,
			rs.appointment.UserID,
			rs.appointment.CreatedAt,
			rs.appointment.UpdatedAt,
		)

	rs.mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "appointments" WHERE date = $1 AND provider_id = $2`),
	).WithArgs(
		rs.appointment.Date,
		rs.appointment.ProviderID,
	).WillReturnRows(rows)

	appointment := rs.repo.FindByDate(&rs.appointment.Date, rs.appointment.ProviderID)

	assert.Equal(rs.T(), rs.appointment.Date, appointment.Date)
	assert.Equal(rs.T(), rs.appointment.ProviderID, appointment.ProviderID)
}

func (rs *AppointmentsRepositorySuite) TestFindAllInMonthFromProvider() {
	rows := sqlmock.NewRows([]string{"id", "date", "provider_id", "user_id", "created_at", "updated_at"}).
		AddRow(
			rs.appointment.ID,
			rs.appointment.Date,
			rs.appointment.ProviderID,
			rs.appointment.UserID,
			rs.appointment.CreatedAt,
			rs.appointment.UpdatedAt,
		)

	parsedMonth, _ := fmt.Printf("'%02d'", rs.dtoFindAllInMonthFromProviderDTO.Month)
	dateFieldName := time.Date(rs.dtoFindAllInMonthFromProviderDTO.Year, time.Month(parsedMonth), 0, 0, 0, 0, 0, time.UTC)
	date := dateFieldName.Format("2006-01-02")

	rs.mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "appointments" WHERE provider_id = $1 AND date LIKE $2`),
	).WithArgs(
		rs.appointment.ProviderID,
		"%"+date+"%",
	).WillReturnRows(rows)

	appointments := rs.repo.FindAllInMonthFromProvider(rs.dtoFindAllInMonthFromProviderDTO)

	assert.Equal(rs.T(), rs.appointment.ID, appointments[0].ID)
	assert.Equal(rs.T(), rs.appointment.ProviderID, appointments[0].ProviderID)
	assert.Equal(rs.T(), rs.appointment.Date, appointments[0].Date)
}

func (rs *AppointmentsRepositorySuite) TestFindAllInDayFromProviderDTO() {
	rows := sqlmock.NewRows([]string{"id", "date", "provider_id", "user_id", "created_at", "updated_at"}).
		AddRow(
			rs.appointment.ID,
			rs.appointment.Date,
			rs.appointment.ProviderID,
			rs.appointment.UserID,
			rs.appointment.CreatedAt,
			rs.appointment.UpdatedAt,
		)

	parsedMonth, _ := fmt.Printf("'%02d'", rs.dtoFindAllInDayFromProviderDTO.Month)
	parsedDay, _ := fmt.Printf("'%02d'", rs.dtoFindAllInDayFromProviderDTO.Day)
	dateFieldName := time.Date(rs.dtoFindAllInDayFromProviderDTO.Year, time.Month(parsedMonth), parsedDay, 0, 0, 0, 0, time.UTC)
	date := dateFieldName.Format("2006-01-02")

	rs.mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "appointments" WHERE provider_id = $1 AND date LIKE $2`),
	).WithArgs(
		rs.appointment.ProviderID,
		"%"+date+"%",
	).WillReturnRows(rows)

	appointments := rs.repo.FindAllInDayFromProvider(rs.dtoFindAllInDayFromProviderDTO)

	assert.Equal(rs.T(), rs.appointment.ID, appointments[0].ID)
	assert.Equal(rs.T(), rs.appointment.ProviderID, appointments[0].ProviderID)
	assert.Equal(rs.T(), rs.appointment.Date, appointments[0].Date)
}

func TestSuiteAppointmentsRepository(t *testing.T) {
	suite.Run(t, new(AppointmentsRepositorySuite))
}

func buildAppointments() *model.Appointment {
	return &model.Appointment{
		ID:         uuid.MustParse("61d4e027-9e40-42e9-84e2-0f68d3aeb9f1"),
		UserID:     "522e3950-5e4e-4cf6-8c95-5ba00880bdbd",
		ProviderID: "8820216f-e77a-47a9-bf18-d5f28a7b519d",
		Date:       time.Date(2022, time.December, 30, 10, 0, 0, 0, time.UTC),
		CreatedAt:  time.Date(2022, time.December, 30, 10, 0, 0, 0, time.UTC),
		UpdatedAt:  time.Date(2022, time.December, 30, 10, 0, 0, 0, time.UTC),
	}
}

func buildFindAllInMonthFromProviderDTO() *dtos.FindAllInMonthFromProviderDTO {
	return &dtos.FindAllInMonthFromProviderDTO{
		ProviderID: "8820216f-e77a-47a9-bf18-d5f28a7b519d",
		Month:      12,
		Year:       2022,
	}
}

func buildFindAllInDayFromProviderDTO() *dtos.FindAllInDayFromProviderDTO {
	return &dtos.FindAllInDayFromProviderDTO{
		ProviderID: "8820216f-e77a-47a9-bf18-d5f28a7b519d",
		Day:        30,
		Month:      12,
		Year:       2022,
	}
}
