package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"simple-bank/src/api"
	"simple-bank/src/dao"
	mockdb "simple-bank/tests/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	account := dao.Account{
		ID:       123,
		Owner:    "example",
		Currency: "USD",
		Balance:  500000.00,
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	store := mockdb.NewMockStore(controller)
	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	server := api.NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/api/v1/accounts/%d", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)

	require.NoError(t, err)

	server.Router.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
}
