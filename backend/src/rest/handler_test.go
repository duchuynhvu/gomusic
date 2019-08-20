package rest

import (
	"dblayer"
	"encoding/json"
	"errors"
	"models"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandler_GetProducts(t *testing.T) {
	//need enable the test mode
	gin.SetMode(gin.TestMode)
	//initialize the mockdblayer
	mockdblayer := dblayer.NewMockDBLayerWithData()
	h := NewHandlerWithDB(mockdblayer)
	const productURL string = "/product"
	//error message type
	type errMSG struct {
		Error string `json:"error"`
	}
	//use table driven testing
	tests := []struct {
		name             string
		inErr            error
		outStatusCode    int
		expectedRespBody interface{}
	}{
		{
			"getproductsnoerrors",
			nil,
			http.StatusOK,
			mockdblayer.GetMockProductsData(),
		},
		{
			"getproductswitherror",
			errors.New("get products error"),
			http.StatusInternalServerError,
			errMSG{Error: "get products error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//set the input error
			mockdblayer.SetError(tt.inErr)
			//create a test request
			req := httptest.NewRequest(http.MethodGet, productURL, nil)
			//create an http response recorder
			w := httptest.NewRecorder()
			//create a fresh gin engine object from the response recorder,
			//we will ignore the context value
			_, engine := gin.CreateTestContext(w)
			//configure the get request
			engine.GET(productURL, h.GetProducts)
			//serve the request
			engine.ServeHTTP(w, req)
			//test the output
			response := w.Result()
			if response.StatusCode != tt.outStatusCode {
				t.Errorf("Received Status code %d does not match expected status code %d", response.StatusCode, tt.outStatusCode)
			}
			//Since we don't know the data type to expect from the http response, we'll use interface{} as the type
			var respBody interface{}
			//if an error was injected, then the response should decode to an error message type
			if tt.inErr != nil {
				var errmsg errMSG
				json.NewDecoder(response.Body).Decode(&errmsg)
				respBody = errmsg
			} else {
				//if an error was not injected, the response should decode to a slice of products data types
				products := []models.Product{}
				json.NewDecoder(response.Body).Decode(&products)
				respBody = products
			}
			if !reflect.DeepEqual(respBody, tt.expectedRespBody) {
				t.Errorf("Received HTTP response body %+v does not match expected HTTP response Body %+v", respBody, tt.expectedRespBody)
			}
		})
	}
}
