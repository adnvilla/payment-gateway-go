package postgresql

import (
	"context"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/vo"
	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateOrderRepository(t *testing.T) {

	sqldb, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer sqldb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       sqldb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	id := uuid.NewV4()

	input := vo.CreateOrderDetail{
		OrderId:      id.String(),
		ProviderType: shared_domain.ProviderType_Stripe,
		Amount:       "100",
		Currency:     "mxn",
	}
	sqlRegex := `INSERT INTO "create_orders"`
	sqlRegex2 := `INSERT INTO "create_order_providers"`
	dbMock.ExpectBegin()
	dbMock.ExpectExec(sqlRegex).WithArgs(AnyString{}, AnyTime{}, input.Amount, input.Currency, input.ProviderType).WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectExec(sqlRegex2).WithArgs(AnyString{}, AnyTime{}, AnyString{}, AnyString{}, input.ProviderType, AnyString{}).WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()
	repo := NewOrderRepository(db)

	_, err = repo.CreateOrder(context.TODO(), input)

	assert.NoError(t, err)

	if err := dbMock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestCaptureOrderRepository(t *testing.T) {

	sqldb, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer sqldb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       sqldb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	id := uuid.NewV4()

	input := vo.CaptureOrderDetail{
		CaptureOrderId: id.String(),
		ProviderType:   shared_domain.ProviderType_Stripe,
	}
	sqlRegex := `INSERT INTO "capture_orders"`
	sqlRegex2 := `INSERT INTO "capture_order_providers"`
	dbMock.ExpectBegin()
	dbMock.ExpectExec(sqlRegex).WithArgs(AnyString{}, AnyTime{}, AnyString{}, AnyString{}, input.ProviderType).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectExec(sqlRegex2).WithArgs(AnyString{}, AnyTime{}, AnyString{}, AnyString{}, input.ProviderType, AnyString{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()
	repo := NewOrderRepository(db)

	_, err = repo.CaptureOrder(context.TODO(), input)

	assert.NoError(t, err)

	if err := dbMock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestGetOrderProviderRepository(t *testing.T) {

	sqldb, dbMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer sqldb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       sqldb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	id := uuid.NewV4()
	id2 := uuid.NewV4()
	id3 := uuid.NewV4()

	sqlRegex := `SELECT * FROM "create_order_providers" WHERE create_order_id = $1 ORDER BY "create_order_providers"."id" LIMIT $2`
	columns := []string{"id", "create_order_id", "provider_order_id", "provider_type", "payload", "created_at"}
	dbMock.ExpectQuery(sqlRegex).
		WithArgs(id.String(), 1).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(id.String(), id2.String(), id3.String(), 1, "rd", 123))
	repo := NewOrderRepository(db)

	_, err = repo.GetOrderProvider(context.TODO(), id)

	assert.NoError(t, err)

	if err := dbMock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(int64)
	return ok
}

type AnyString struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyString) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}
