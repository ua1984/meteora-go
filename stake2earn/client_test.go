package stake2earn_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/ua1984/meteora-go/internal/httpclient"
	"github.com/ua1984/meteora-go/stake2earn"
)

type Stake2EarnClientTestSuite struct {
	suite.Suite
}

func TestStake2EarnClient(t *testing.T) {
	suite.Run(t, new(Stake2EarnClientTestSuite))
}

func (s *Stake2EarnClientTestSuite) setupTestServer(method, wantURL string, status int, response any) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Equal(method, r.Method)
		s.Equal(wantURL, r.URL.String())
		s.Equal("application/json", r.Header.Get("Accept"))

		w.WriteHeader(status)
		if str, ok := response.(string); ok {
			w.Write([]byte(str))
		} else {
			json.NewEncoder(w).Encode(response)
		}
	}))
}

func (s *Stake2EarnClientTestSuite) TestGetAnalytics() {
	tests := []struct {
		name       string
		response   any
		status     int
		wantErr    bool
		wantURL    string
		wantResult float64 // TotalStakedAmountUSD
	}{
		{
			name: "should successfully get analytics",
			response: stake2earn.Analytics{
				TotalFeeVaults:       5,
				TotalStakedAmountUSD: 1234.56,
			},
			status:     http.StatusOK,
			wantURL:    "/analytics/all",
			wantResult: 1234.56,
		},
		{
			name:     "should return error on API failure",
			status:   http.StatusInternalServerError,
			response: "Internal Server Error",
			wantErr:  true,
			wantURL:  "/analytics/all",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()

			httpClient := httpclient.New(server.URL, nil)
			client := stake2earn.NewClient(httpClient)

			// Act
			resp, err := client.GetAnalytics(context.TODO())

			// Assert
			if tt.wantErr {
				s.Error(err)
				s.Nil(resp)
			} else {
				s.NoError(err)
				s.NotNil(resp)
				s.Equal(tt.wantResult, resp.TotalStakedAmountUSD)
			}
		})
	}
}

func (s *Stake2EarnClientTestSuite) TestListVaults() {
	tests := []struct {
		name       string
		response   any
		status     int
		wantErr    bool
		wantURL    string
		wantResult int // Number of vaults
	}{
		{
			name: "should successfully list all vaults",
			response: stake2earn.VaultListResponse{
				Total: 2,
				Data: []stake2earn.Vault{
					{VaultAddress: "vault1"},
					{VaultAddress: "vault2"},
				},
			},
			status:     http.StatusOK,
			wantURL:    "/vault/all",
			wantResult: 2,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()

			httpClient := httpclient.New(server.URL, nil)
			client := stake2earn.NewClient(httpClient)

			// Act
			resp, err := client.ListVaults(context.TODO())

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Len(resp.Data, tt.wantResult)
			}
		})
	}
}

func (s *Stake2EarnClientTestSuite) TestFilterVaults() {
	tests := []struct {
		name       string
		params     *stake2earn.FilterParams
		response   any
		status     int
		wantErr    bool
		wantURL    string
		wantResult int
	}{
		{
			name: "should successfully filter vaults with params",
			params: &stake2earn.FilterParams{
				Page:       ptr(1),
				Size:       ptr(10),
				SortBy:     ptr("total_staked_amount_usd"),
				SortOrder:  ptr("desc"),
				SearchTerm: ptr("SOL"),
			},
			response: stake2earn.VaultListResponse{
				Total: 1,
				Data:  []stake2earn.Vault{{VaultAddress: "vault1"}},
			},
			status:     http.StatusOK,
			wantURL:    "/vault/filter?page=1&search_term=SOL&size=10&sort_by=total_staked_amount_usd&sort_order=desc",
			wantResult: 1,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()

			httpClient := httpclient.New(server.URL, nil)
			client := stake2earn.NewClient(httpClient)

			// Act
			resp, err := client.FilterVaults(context.TODO(), tt.params)

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Len(resp.Data, tt.wantResult)
			}
		})
	}
}

func (s *Stake2EarnClientTestSuite) TestGetVault() {
	tests := []struct {
		name     string
		address  string
		response any
		status   int
		wantErr  bool
		wantURL  string
	}{
		{
			name:    "should successfully get a single vault",
			address: "vault123",
			response: stake2earn.Vault{
				VaultAddress: "vault123",
			},
			status:  http.StatusOK,
			wantURL: "/vault/vault123",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()

			httpClient := httpclient.New(server.URL, nil)
			client := stake2earn.NewClient(httpClient)

			// Act
			resp, err := client.GetVault(context.TODO(), tt.address)

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Equal(tt.address, resp.VaultAddress)
			}
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}
