package gorm

import (
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGeConnection(t *testing.T) {

	os.Setenv("PAYMENT_GATEWAY_POSTGRES_HOST", "1")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_USER", "q")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_PASSWORD", "s")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_DBNAME", "t")

	UseDefaultConnection()

	con := GetConnection()

	assert.NotNil(t, con)
}

func TestGormOpenFail(t *testing.T) {
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_PORT", "loc=alhos&$#/$)t")

	UseDefaultConnection()

	conn := GetConnection()

	assert.Error(t, conn.Error)
}

type User struct {
	Model
}

func TestVerifyBeforeCreateIsExecute(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sql mock: %v", err)
	}
	dialector := postgres.New(postgres.Config{
		Conn:       sqlDB,
		DriverName: "postgres",
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	user := User{}
	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	assert.NotEmpty(t, user.ID)
}

func TestResetConnection(t *testing.T) {
	r := require.New(t)
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_HOST", "1")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_PORT", "1")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_USER", "q")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_PASSWORD", "s")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_DBNAME", "t")

	ResetConnection()
	con := GetConnection().Instance

	ResetConnection()
	con2 := GetConnection().Instance

	r.NotNil(con)
	r.NotSame(con, con2)
}
