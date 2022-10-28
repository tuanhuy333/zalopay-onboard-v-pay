package handlers

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	"order-service/mock"
)

func TestHandler_GetOrders(t *testing.T) {
	t.Run("GetOrders success", func(t *testing.T) {
		// mock service
		o := mock.NewMockOrderService(gomock.NewController(t))
		o.EXPECT().GetAllOrder(gomock.Any()).Return(nil, nil)
		p := mock.NewMockPublisherService(gomock.NewController(t))
		h := Handler{
			Service:          o,
			PublisherService: p,
		}

		res := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(res)

		// call
		h.GetOrders(context)

		// assert
		if res.Code != http.StatusOK {
			t.Error("error")
		}
	})
	t.Run("GetOrders error", func(t *testing.T) {
		// mock service
		o := mock.NewMockOrderService(gomock.NewController(t))
		o.EXPECT().GetAllOrder(gomock.Any()).Return(nil, errors.New("error"))
		p := mock.NewMockPublisherService(gomock.NewController(t))
		h := Handler{
			Service:          o,
			PublisherService: p,
		}

		res := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(res)

		// call
		h.GetOrders(context)

		// assert
		if res.Code != http.StatusInternalServerError {
			t.Error("error")
		}
	})

}
func TestHandler_GetOrderByID(t *testing.T) {
	t.Run("GetOrderByID success", func(t *testing.T) {
		// mock service
		o := mock.NewMockOrderService(gomock.NewController(t))
		o.EXPECT().GetOrderById(gomock.Any()).Return(nil, nil)
		p := mock.NewMockPublisherService(gomock.NewController(t))
		h := Handler{
			Service:          o,
			PublisherService: p,
		}

		res := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(res)
		context.Params = []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}
		// call
		h.GetOrderById(context)

		// assert
		if res.Code != http.StatusOK {
			t.Error("error")
		}
	})
	t.Run("GetOrderByID error", func(t *testing.T) {
		// mock service
		o := mock.NewMockOrderService(gomock.NewController(t))
		o.EXPECT().GetOrderById(gomock.Any()).Return(nil, errors.New("error"))
		p := mock.NewMockPublisherService(gomock.NewController(t))
		h := Handler{
			Service:          o,
			PublisherService: p,
		}

		res := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(res)
		context.Params = []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}
		// call
		h.GetOrderById(context)

		// assert
		if res.Code != http.StatusInternalServerError {
			t.Error("error")
		}
	})

}

func TestHandler_CreateOrder(t *testing.T) {
	t.Run("CreateOrder fail request", func(t *testing.T) {
		// mock service
		o := mock.NewMockOrderService(gomock.NewController(t))
		p := mock.NewMockPublisherService(gomock.NewController(t))
		h := Handler{
			Service:          o,
			PublisherService: p,
		}

		res := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(res)

		// call
		h.CreateOrders(context)

		// assert
		if res.Code != http.StatusBadRequest {
			t.Error("error")
		}
	})
	t.Run("CreateOrder with Create Order error", func(t *testing.T) {
		// mock service
		o := mock.NewMockOrderService(gomock.NewController(t))
		o.EXPECT().CreateOrder(gomock.Any()).Return(errors.New("error"))
		p := mock.NewMockPublisherService(gomock.NewController(t))
		h := Handler{
			Service:          o,
			PublisherService: p,
		}

		res := httptest.NewRecorder()
		httpReq := http.Request{}
		context, _ := gin.CreateTestContext(res)
		context.Request = &httpReq
		context.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(`{"amount":23}`)))

		// call
		h.CreateOrders(context)

		// assert
		if res.Code != http.StatusInternalServerError {
			t.Error("error")
		}
	})
	t.Run("CreateOrder success", func(t *testing.T) {
		// mock service
		o := mock.NewMockOrderService(gomock.NewController(t))
		o.EXPECT().CreateOrder(gomock.Any()).Return(nil)
		p := mock.NewMockPublisherService(gomock.NewController(t))
		h := Handler{
			Service:          o,
			PublisherService: p,
		}

		res := httptest.NewRecorder()
		httpReq := http.Request{}
		context, _ := gin.CreateTestContext(res)
		context.Request = &httpReq
		context.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(`{"amount":23}`)))

		// call
		h.CreateOrders(context)

		// assert
		if res.Code != http.StatusOK {
			t.Error("error")
		}
	})

}