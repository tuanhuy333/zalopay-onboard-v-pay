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
	t.Run("CreateOrder fail authentication", func(t *testing.T) {
		// mock service
		o := mock.NewMockOrderService(gomock.NewController(t))
		p := mock.NewMockPublisherService(gomock.NewController(t))
		h := Handler{
			Service:          o,
			PublisherService: p,
		}

		res := httptest.NewRecorder()
		httpReq := http.Request{}
		context, _ := gin.CreateTestContext(res)
		context.Request = &httpReq
		context.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(`{"amount":23,"appID":2000,"mac":"abc"}`)))

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
		context.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(`{"amount":23,"appID":2000,"mac":"d3af0db2f2acd7ec5c9114188f12288e650b669921011124e834418f82da97d0"}`)))

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
		context.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(`{"amount":23,"appID":2000,"mac":"d3af0db2f2acd7ec5c9114188f12288e650b669921011124e834418f82da97d0"}`)))

		// call
		h.CreateOrders(context)

		// assert
		if res.Code != http.StatusOK {
			t.Error("error")
		}
	})

}
