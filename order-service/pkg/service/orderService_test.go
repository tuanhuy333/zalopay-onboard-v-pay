package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"order-service/models"
)

func mockDB() *gorm.DB {
	// create DB in mem for testing
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&models.Order{})

	return db
}

func TestStorage_New(t *testing.T) {

	db := mockDB()

	o := New(db)

	require.NotNil(t, o)

}

func TestStorage_CreateOrder(t *testing.T) {

	db := mockDB()
	o := Storage{db: db}

	err := o.CreateOrder(&models.Order{})

	require.Nil(t, err)

}

func TestStorage_GetAllOrder(t *testing.T) {
	t.Run("get all success", func(t *testing.T) {
		db := mockDB()
		o := Storage{db: db}

		o1, _ := o.GetAllOrder(&[]models.Order{})

		require.NotNil(t, o1)
	})

}

func TestStorage_GetOrderById(t *testing.T) {
	t.Run("get order by id not found", func(t *testing.T) {
		db := mockDB()
		o := Storage{db: db}

		_, err := o.GetOrderById(10)

		require.Error(t, err, "record not found")
	})
	t.Run("get order by id success", func(t *testing.T) {
		db := mockDB()
		db.Create(&models.Order{
			OrderNo:     22,
			MerchantID:  "abc",
			AppID:       0,
			Status:      0,
			Amount:      0,
			ProductCode: "",
			Description: "",
			CreateTime:  time.Time{},
		})
		o := Storage{db: db}

		o1, _ := o.GetOrderById(22)

		require.NotNil(t, o1)
	})
}

func TestStorage_UpdateOrderById(t *testing.T) {
	t.Run("update order by id not found", func(t *testing.T) {
		db := mockDB()
		o := Storage{db: db}

		_, err := o.UpdateOrderById(10, &models.Order{})

		require.Error(t, err, "record not found")
	})
	t.Run("update order by id success", func(t *testing.T) {
		db := mockDB()
		db.Create(&models.Order{
			OrderNo:     33,
			MerchantID:  "abc",
			AppID:       0,
			Status:      0,
			Amount:      0,
			ProductCode: "",
			Description: "",
			CreateTime:  time.Time{},
		})
		o := Storage{db: db}

		o1, err := o.UpdateOrderById(33, &models.Order{})

		require.Nil(t, err, o1)
	})

}
