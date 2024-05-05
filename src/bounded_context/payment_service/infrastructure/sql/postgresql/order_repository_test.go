package postgresql

import (
	"context"
	"database/sql/driver"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/vo"
	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateOrderRepository(t *testing.T) {
	r := require.New(t)
	db, dbMock, err := getDbMock(sqlmock.QueryMatcherRegexp)
	r.NoError(err)

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

	r.NoError(err)
	r.NoError(dbMock.ExpectationsWereMet())
}

func TestCreateOrderRepositoryFail(t *testing.T) {
	r := require.New(t)
	db, dbMock, err := getDbMock(sqlmock.QueryMatcherRegexp)
	r.NoError(err)

	id := uuid.NewV4()
	input := vo.CreateOrderDetail{
		OrderId:      id.String(),
		ProviderType: shared_domain.ProviderType_Stripe,
		Amount:       "100",
		Currency:     "mxn",
	}

	sqlRegex := `INSERT INTO "create_orders"`
	dbMock.ExpectBegin()
	dbMock.ExpectExec(sqlRegex).WithArgs(AnyString{}, AnyTime{}, input.Amount, input.Currency, input.ProviderType).WillReturnError(fmt.Errorf("error"))
	dbMock.ExpectRollback()

	repo := NewOrderRepository(db)
	_, err = repo.CreateOrder(context.TODO(), input)

	r.Error(err)
	r.NoError(dbMock.ExpectationsWereMet())
}

func TestCaptureOrderRepository(t *testing.T) {
	r := require.New(t)
	db, dbMock, err := getDbMock(sqlmock.QueryMatcherRegexp)
	r.NoError(err)

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

	r.NoError(err)
	r.NoError(dbMock.ExpectationsWereMet())
}

func TestCaptureOrderRepositoryFail(t *testing.T) {
	r := require.New(t)
	db, dbMock, err := getDbMock(sqlmock.QueryMatcherRegexp)
	r.NoError(err)

	id := uuid.NewV4()
	input := vo.CaptureOrderDetail{
		CaptureOrderId: id.String(),
		ProviderType:   shared_domain.ProviderType_Stripe,
	}

	sqlRegex := `INSERT INTO "capture_orders"`
	dbMock.ExpectBegin()
	dbMock.ExpectExec(sqlRegex).WithArgs(AnyString{}, AnyTime{}, AnyString{}, AnyString{}, input.ProviderType).
		WillReturnError(fmt.Errorf("error"))
	dbMock.ExpectRollback()

	repo := NewOrderRepository(db)
	_, err = repo.CaptureOrder(context.TODO(), input)

	r.Error(err)
	r.NoError(dbMock.ExpectationsWereMet())
}

func TestGetCaptureOrderProviderRepository(t *testing.T) {
	r := require.New(t)
	db, dbMock, err := getDbMock(sqlmock.QueryMatcherEqual)
	r.NoError(err)

	id := uuid.NewV4()
	id2 := uuid.NewV4()
	id3 := uuid.NewV4()

	sqlRegex := `SELECT * FROM "capture_order_providers" WHERE capture_order_id = $1 ORDER BY "capture_order_providers"."id" LIMIT $2`
	columns := []string{"id", "capture_order_id", "provider_order_id", "provider_type", "payload", "created_at"}
	dbMock.ExpectQuery(sqlRegex).
		WithArgs(id.String(), 1).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(id.String(), id2.String(), id3.String(), 1, "rd", 123))

	repo := NewOrderRepository(db)
	_, err = repo.GetCaptureOrderProvider(context.TODO(), id)

	r.NoError(err)
	r.NoError(dbMock.ExpectationsWereMet())
}

func TestGetCaptureOrderProviderRepositoryFail(t *testing.T) {
	r := require.New(t)
	db, dbMock, err := getDbMock(sqlmock.QueryMatcherEqual)
	r.NoError(err)

	id := uuid.NewV4()

	sqlRegex := `SELECT * FROM "capture_order_providers" WHERE capture_order_id = $1 ORDER BY "capture_order_providers"."id" LIMIT $2`
	dbMock.ExpectQuery(sqlRegex).
		WithArgs(id.String(), 1).
		WillReturnError(fmt.Errorf("error"))

	repo := NewOrderRepository(db)
	_, err = repo.GetCaptureOrderProvider(context.TODO(), id)

	r.Error(err)
	r.NoError(dbMock.ExpectationsWereMet())
}

func TestGetOrderProviderRepository(t *testing.T) {
	r := require.New(t)
	db, dbMock, err := getDbMock(sqlmock.QueryMatcherEqual)
	r.NoError(err)

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

	r.NoError(err)
	r.NoError(dbMock.ExpectationsWereMet())
}

func TestGetOrderProviderRepositoryFail(t *testing.T) {
	r := require.New(t)
	db, dbMock, err := getDbMock(sqlmock.QueryMatcherEqual)
	r.NoError(err)

	id := uuid.NewV4()

	sqlRegex := `SELECT * FROM "create_order_providers" WHERE create_order_id = $1 ORDER BY "create_order_providers"."id" LIMIT $2`
	dbMock.ExpectQuery(sqlRegex).
		WithArgs(id.String(), 1).
		WillReturnError(fmt.Errorf("error"))

	repo := NewOrderRepository(db)
	_, err = repo.GetOrderProvider(context.TODO(), id)

	r.Error(err)
	r.NoError(dbMock.ExpectationsWereMet())

}

func TestGetOrderRepository(t *testing.T) {
	r := require.New(t)
	db, dbMock, err := getDbMock(sqlmock.QueryMatcherEqual)
	r.NoError(err)

	id := uuid.NewV4()
	id2 := uuid.NewV4()
	id3 := uuid.NewV4()

	sqlRegex := `SELECT * FROM "create_orders" WHERE "create_orders"."id" = $1 ORDER BY "create_orders"."id" LIMIT $2`
	columns := []string{"id", "amount", "currency", "provider_type", "created_at"}
	dbMock.ExpectQuery(sqlRegex).
		WithArgs(id.String(), 1).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(id.String(), id2.String(), id3.String(), 1, 123))
	repo := NewOrderRepository(db)

	_, err = repo.GetOrder(context.TODO(), id)

	r.NoError(err)
	r.NoError(dbMock.ExpectationsWereMet())
}

func TestGetOrderRepositoryFail(t *testing.T) {
	r := require.New(t)
	db, dbMock, err := getDbMock(sqlmock.QueryMatcherEqual)
	r.NoError(err)

	id := uuid.NewV4()

	sqlRegex := `SELECT * FROM "create_orders" WHERE "create_orders"."id" = $1 ORDER BY "create_orders"."id" LIMIT $2`
	dbMock.ExpectQuery(sqlRegex).
		WithArgs(id.String(), 1).
		WillReturnError(fmt.Errorf("error"))
	repo := NewOrderRepository(db)

	_, err = repo.GetOrder(context.TODO(), id)

	r.Error(err)
	r.NoError(dbMock.ExpectationsWereMet())

}

func TestCreateRefundRepository(t *testing.T) {
	r := require.New(t)
	db, dbMock, err := getDbMock(sqlmock.QueryMatcherRegexp)
	r.NoError(err)

	id := uuid.NewV4()

	input := vo.CreateRefundDetail{
		RefundOrderId: id.String(),
		ProviderType:  shared_domain.ProviderType_Stripe,
	}
	sqlRegex := `INSERT INTO "refunds"`
	sqlRegex2 := `INSERT INTO "refund_providers"`
	dbMock.ExpectBegin()
	dbMock.ExpectExec(sqlRegex).WithArgs(AnyString{}, AnyTime{}, AnyString{}, AnyString{}, input.ProviderType).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectExec(sqlRegex2).WithArgs(AnyString{}, AnyTime{}, AnyString{}, AnyString{}, input.ProviderType, AnyString{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()
	repo := NewOrderRepository(db)

	_, err = repo.CreateRefund(context.TODO(), input)

	r.NoError(err)
	r.NoError(dbMock.ExpectationsWereMet())
}

func TestCreateRefundRepositoryFail(t *testing.T) {
	r := require.New(t)
	db, dbMock, err := getDbMock(sqlmock.QueryMatcherRegexp)
	r.NoError(err)

	id := uuid.NewV4()

	input := vo.CreateRefundDetail{
		RefundOrderId: id.String(),
		ProviderType:  shared_domain.ProviderType_Stripe,
	}
	sqlRegex := `INSERT INTO "refunds"`
	dbMock.ExpectBegin()
	dbMock.ExpectExec(sqlRegex).WithArgs(AnyString{}, AnyTime{}, AnyString{}, AnyString{}, input.ProviderType).
		WillReturnError(fmt.Errorf("error"))
	dbMock.ExpectRollback()
	repo := NewOrderRepository(db)

	_, err = repo.CreateRefund(context.TODO(), input)

	r.Error(err)
	r.NoError(dbMock.ExpectationsWereMet())

}

func TestGetRefundProviderRepository(t *testing.T) {
	r := require.New(t)
	db, dbMock, err := getDbMock(sqlmock.QueryMatcherEqual)
	r.NoError(err)

	id := uuid.NewV4()
	id2 := uuid.NewV4()
	id3 := uuid.NewV4()

	sqlRegex := `SELECT * FROM "refund_providers" WHERE refund_id = $1 ORDER BY "refund_providers"."id" LIMIT $2`
	columns := []string{"id", "refund_id", "provider_order_id", "provider_type", "payload", "created_at"}
	dbMock.ExpectQuery(sqlRegex).
		WithArgs(id.String(), 1).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(id.String(), id2.String(), id3.String(), 1, "rd", 123))
	repo := NewOrderRepository(db)

	_, err = repo.GetRefundProvider(context.TODO(), id)

	r.NoError(err)
	r.NoError(dbMock.ExpectationsWereMet())
}

func TestGetRefundProviderRepositoryFail(t *testing.T) {
	r := require.New(t)
	db, dbMock, err := getDbMock(sqlmock.QueryMatcherEqual)
	r.NoError(err)

	id := uuid.NewV4()

	sqlRegex := `SELECT * FROM "refund_providers" WHERE refund_id = $1 ORDER BY "refund_providers"."id" LIMIT $2`
	dbMock.ExpectQuery(sqlRegex).
		WithArgs(id.String(), 1).
		WillReturnError(fmt.Errorf("error"))
	repo := NewOrderRepository(db)

	_, err = repo.GetRefundProvider(context.TODO(), id)

	r.Error(err)
	r.NoError(dbMock.ExpectationsWereMet())
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

func getDbMock(queryMatcher sqlmock.QueryMatcher) (*gorm.DB, sqlmock.Sqlmock, error) {
	sqldb, dbMock, err := sqlmock.New(sqlmock.QueryMatcherOption(queryMatcher))
	if err != nil {
		return nil, nil, err
	}
	dialector := postgres.New(postgres.Config{
		Conn:       sqldb,
		DriverName: "postgres",
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	return db, dbMock, nil
}
