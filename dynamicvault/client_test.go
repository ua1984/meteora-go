package dynamicvault_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/ua1984/meteora-go/dynamicvault"
	"github.com/ua1984/meteora-go/internal/httpclient"
)

type DynamicVaultClientTestSuite struct {
	suite.Suite
}

func TestDynamicVaultClient(t *testing.T) {
	suite.Run(t, new(DynamicVaultClientTestSuite))
}

func (s *DynamicVaultClientTestSuite) setupTestServer(method, wantURL string, status int, response any) *httptest.Server {
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

func (s *DynamicVaultClientTestSuite) TestListVaultInfo() {
	tests := []struct {
		name       string
		response   any
		status     int
		wantErr    bool
		wantURL    string
		wantResult int
	}{
		{
			name: "should successfully list vault info",
			response: []dynamicvault.VaultInfo{
				{Symbol: "SOL", Pubkey: "pubkey1"},
				{Symbol: "USDC", Pubkey: "pubkey2"},
			},
			status:     http.StatusOK,
			wantURL:    "/vault_info",
			wantResult: 2,
		},
		{
			name:       "should return error on API failure",
			status:     http.StatusInternalServerError,
			response:   "Internal Server Error",
			wantErr:    true,
			wantURL:    "/vault_info",
			wantResult: 0,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()

			httpClient := httpclient.New(server.URL, nil)
			client := dynamicvault.NewClient(httpClient)

			// Act
			resp, err := client.ListVaultInfo(context.TODO())

			// Assert
			if tt.wantErr {
				s.Error(err)
				s.Nil(resp)
			} else {
				s.NoError(err)
				s.Len(resp, tt.wantResult)
			}
		})
	}
}

func (s *DynamicVaultClientTestSuite) TestListVaultAddresses() {
	tests := []struct {
		name       string
		response   any
		status     int
		wantErr    bool
		wantURL    string
		wantResult int
	}{
		{
			name: "should successfully list vault addresses",
			response: []dynamicvault.VaultAddress{
				{Symbol: "SOL", VaultAddress: "addr1"},
			},
			status:     http.StatusOK,
			wantURL:    "/vault_addresses",
			wantResult: 1,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()

			httpClient := httpclient.New(server.URL, nil)
			client := dynamicvault.NewClient(httpClient)

			// Act
			resp, err := client.ListVaultAddresses(context.TODO())

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Len(resp, tt.wantResult)
			}
		})
	}
}

func (s *DynamicVaultClientTestSuite) TestGetVaultState() {
	tests := []struct {
		name      string
		tokenMint string
		response  any
		status    int
		wantErr   bool
		wantURL   string
	}{
		{
			name:      "should successfully get vault state",
			tokenMint: "mint123",
			response: dynamicvault.VaultState{
				TokenAddress: "mint123",
				Pubkey:       "vault123",
			},
			status:  http.StatusOK,
			wantURL: "/vault_state/mint123",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()

			httpClient := httpclient.New(server.URL, nil)
			client := dynamicvault.NewClient(httpClient)

			// Act
			resp, err := client.GetVaultState(context.TODO(), tt.tokenMint)

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Equal(tt.tokenMint, resp.TokenAddress)
			}
		})
	}
}

func (s *DynamicVaultClientTestSuite) TestGetAPYState() {
	tests := []struct {
		name      string
		tokenMint string
		response  any
		status    int
		wantErr   bool
		wantURL   string
	}{
		{
			name:      "should successfully get APY state",
			tokenMint: "mint123",
			response: dynamicvault.APYState{
				ClosestAPY: []dynamicvault.APYBreakdown{{StrategyName: "Strategy 1", APY: 5.5}},
			},
			status:  http.StatusOK,
			wantURL: "/apy_state/mint123",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()

			httpClient := httpclient.New(server.URL, nil)
			client := dynamicvault.NewClient(httpClient)

			// Act
			resp, err := client.GetAPYState(context.TODO(), tt.tokenMint)

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.NotEmpty(resp.ClosestAPY)
				s.Equal(5.5, resp.ClosestAPY[0].APY)
			}
		})
	}
}

func (s *DynamicVaultClientTestSuite) TestGetAPYByTimeRange() {
	tests := []struct {
		name      string
		tokenMint string
		start     int64
		end       int64
		response  any
		status    int
		wantErr   bool
		wantURL   string
	}{
		{
			name:      "should successfully get APY by time range",
			tokenMint: "mint123",
			start:     1000,
			end:       2000,
			response: []dynamicvault.APYEntry{
				{Timestamp: 1500, APY: 4.2},
			},
			status:  http.StatusOK,
			wantURL: "/apy_filter/mint123/1000/2000",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()

			httpClient := httpclient.New(server.URL, nil)
			client := dynamicvault.NewClient(httpClient)

			// Act
			resp, err := client.GetAPYByTimeRange(context.TODO(), tt.tokenMint, tt.start, tt.end)

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Len(resp, 1)
				s.Equal(4.2, resp[0].APY)
			}
		})
	}
}

func (s *DynamicVaultClientTestSuite) TestGetVirtualPrice() {
	tests := []struct {
		name      string
		tokenMint string
		strategy  string
		response  any
		status    int
		wantErr   bool
		wantURL   string
	}{
		{
			name:      "should successfully get virtual price",
			tokenMint: "mint123",
			strategy:  "strat123",
			response: []dynamicvault.VirtualPrice{
				{Price: "1.005", Timestamp: 12345},
			},
			status:  http.StatusOK,
			wantURL: "/virtual_price/mint123/strat123",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()

			httpClient := httpclient.New(server.URL, nil)
			client := dynamicvault.NewClient(httpClient)

			// Act
			resp, err := client.GetVirtualPrice(context.TODO(), tt.tokenMint, tt.strategy)

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Len(resp, 1)
				s.Equal("1.005", resp[0].Price)
			}
		})
	}
}
