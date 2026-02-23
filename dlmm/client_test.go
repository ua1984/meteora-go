package dlmm_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/ua1984/meteora-go/dlmm"
	"github.com/ua1984/meteora-go/internal/httpclient"
)

type DLMMClientTestSuite struct {
	suite.Suite
}

func TestDLMMClient(t *testing.T) {
	suite.Run(t, new(DLMMClientTestSuite))
}

func ptr[T any](v T) *T {
	return &v
}

func (s *DLMMClientTestSuite) setupTestServer(method, wantURL string, status int, response any) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Equal(method, r.Method)
		s.Equal(wantURL, r.URL.String())
		s.Equal("application/json", r.Header.Get("Accept"))
		s.Equal("Meteora Go SDK/1.0.0", r.Header.Get("User-Agent"))

		w.WriteHeader(status)
		if str, ok := response.(string); ok {
			w.Write([]byte(str))
		} else {
			json.NewEncoder(w).Encode(response)
		}
	}))
}

func (s *DLMMClientTestSuite) TestListPools() {
	tests := []struct {
		name       string
		params     *dlmm.ListPoolsParams
		status     int
		response   any
		wantURL    string
		wantErr    bool
		wantResult *dlmm.PaginatedResponse[dlmm.Pool]
	}{
		{
			name: "should successfully list pools with params",
			params: &dlmm.ListPoolsParams{
				Page:     ptr(1),
				PageSize: ptr(10),
			},
			status: http.StatusOK,
			response: dlmm.PaginatedResponse[dlmm.Pool]{
				Data: []dlmm.Pool{{Address: "pool1"}},
			},
			wantURL: "/pools?page=1&page_size=10",
			wantResult: &dlmm.PaginatedResponse[dlmm.Pool]{
				Data: []dlmm.Pool{{Address: "pool1"}},
			},
		},
		{
			name:       "should successfully list pools without params",
			status:     http.StatusOK,
			response:   dlmm.PaginatedResponse[dlmm.Pool]{Data: []dlmm.Pool{}},
			wantURL:    "/pools",
			wantResult: &dlmm.PaginatedResponse[dlmm.Pool]{Data: []dlmm.Pool{}},
		},
		{
			name:     "should return error on API failure",
			status:   http.StatusInternalServerError,
			response: "Error",
			wantURL:  "/pools",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()
			client := dlmm.NewClient(httpclient.New(server.URL, nil))

			// Act
			resp, err := client.ListPools(context.Background(), tt.params)

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Equal(tt.wantResult, resp)
			}
		})
	}
}

func (s *DLMMClientTestSuite) TestListGroups() {
	tests := []struct {
		name       string
		params     *dlmm.ListGroupsParams
		status     int
		response   any
		wantURL    string
		wantErr    bool
		wantResult *dlmm.PaginatedResponse[dlmm.PoolGroup]
	}{
		{
			name: "should successfully list groups with params",
			params: &dlmm.ListGroupsParams{
				Page: ptr(2),
			},
			status: http.StatusOK,
			response: dlmm.PaginatedResponse[dlmm.PoolGroup]{
				Data: []dlmm.PoolGroup{{GroupName: "SOL-USDC"}},
			},
			wantURL: "/pools/groups?page=2",
			wantResult: &dlmm.PaginatedResponse[dlmm.PoolGroup]{
				Data: []dlmm.PoolGroup{{GroupName: "SOL-USDC"}},
			},
		},
		{
			name:     "should handle error",
			status:   http.StatusNotFound,
			response: "Not Found",
			wantURL:  "/pools/groups",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()
			client := dlmm.NewClient(httpclient.New(server.URL, nil))

			// Act
			resp, err := client.ListGroups(context.Background(), tt.params)

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Equal(tt.wantResult, resp)
			}
		})
	}
}

func (s *DLMMClientTestSuite) TestGetGroup() {
	tests := []struct {
		name              string
		lexicalOrderMints string
		params            *dlmm.GetGroupParams
		status            int
		response          any
		wantURL           string
		wantErr           bool
		wantResult        *dlmm.PaginatedResponse[dlmm.Pool]
	}{
		{
			name:              "should successfully get group",
			lexicalOrderMints: "mint1-mint2",
			params: &dlmm.GetGroupParams{
				PageSize: ptr(5),
			},
			status: http.StatusOK,
			response: dlmm.PaginatedResponse[dlmm.Pool]{
				Data: []dlmm.Pool{{Address: "poolA"}},
			},
			wantURL: "/pools/groups/mint1-mint2?page_size=5",
			wantResult: &dlmm.PaginatedResponse[dlmm.Pool]{
				Data: []dlmm.Pool{{Address: "poolA"}},
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()
			client := dlmm.NewClient(httpclient.New(server.URL, nil))

			// Act
			resp, err := client.GetGroup(context.Background(), tt.lexicalOrderMints, tt.params)

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Equal(tt.wantResult, resp)
			}
		})
	}
}

func (s *DLMMClientTestSuite) TestGetPool() {
	tests := []struct {
		name       string
		address    string
		status     int
		response   any
		wantURL    string
		wantErr    bool
		wantResult *dlmm.Pool
	}{
		{
			name:       "should successfully get pool",
			address:    "addr123",
			status:     http.StatusOK,
			response:   dlmm.Pool{Address: "addr123"},
			wantURL:    "/pools/addr123",
			wantResult: &dlmm.Pool{Address: "addr123"},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()
			client := dlmm.NewClient(httpclient.New(server.URL, nil))

			// Act
			resp, err := client.GetPool(context.Background(), tt.address)

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Equal(tt.wantResult, resp)
			}
		})
	}
}

func (s *DLMMClientTestSuite) TestGetOHLCV() {
	tests := []struct {
		name       string
		address    string
		params     *dlmm.OHLCVParams
		status     int
		response   any
		wantURL    string
		wantErr    bool
		wantResult *dlmm.OHLCVResponse
	}{
		{
			name:    "should successfully get ohlcv with resolution",
			address: "pool1",
			params: &dlmm.OHLCVParams{
				TimeframeBasedParams: dlmm.TimeframeBasedParams{
					Timeframe: ptr("1H"),
				},
			},
			status: http.StatusOK,
			response: dlmm.OHLCVResponse{
				Data: []dlmm.OHLCV{{Timestamp: 1000}},
			},
			wantURL: "/pools/pool1/ohlcv?timeframe=1H",
			wantResult: &dlmm.OHLCVResponse{
				Data: []dlmm.OHLCV{{Timestamp: 1000}},
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()
			client := dlmm.NewClient(httpclient.New(server.URL, nil))

			// Act
			resp, err := client.GetOHLCV(context.Background(), tt.address, tt.params)

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Equal(tt.wantResult, resp)
			}
		})
	}
}

func (s *DLMMClientTestSuite) TestGetVolumeHistory() {
	tests := []struct {
		name       string
		address    string
		params     *dlmm.VolumeHistoryParams
		status     int
		response   any
		wantURL    string
		wantErr    bool
		wantResult *dlmm.VolumeHistoryResponse
	}{
		{
			name:    "should successfully get volume history",
			address: "poolX",
			params: &dlmm.VolumeHistoryParams{
				TimeframeBasedParams: dlmm.TimeframeBasedParams{
					Timeframe: ptr("24h"),
				},
			},
			status: http.StatusOK,
			response: dlmm.VolumeHistoryResponse{
				Data: []dlmm.VolumeHistory{{Timestamp: 2000}},
			},
			wantURL: "/pools/poolX/volume/history?timeframe=24h",
			wantResult: &dlmm.VolumeHistoryResponse{
				Data: []dlmm.VolumeHistory{{Timestamp: 2000}},
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()
			client := dlmm.NewClient(httpclient.New(server.URL, nil))

			// Act
			resp, err := client.GetVolumeHistory(context.Background(), tt.address, tt.params)

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Equal(tt.wantResult, resp)
			}
		})
	}
}

func (s *DLMMClientTestSuite) TestGetProtocolMetrics() {
	tests := []struct {
		name       string
		status     int
		response   any
		wantURL    string
		wantErr    bool
		wantResult *dlmm.ProtocolMetrics
	}{
		{
			name:       "should successfully get protocol metrics",
			status:     http.StatusOK,
			response:   dlmm.ProtocolMetrics{TotalVolume: 999.9},
			wantURL:    "/stats/protocol_metrics",
			wantResult: &dlmm.ProtocolMetrics{TotalVolume: 999.9},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()
			client := dlmm.NewClient(httpclient.New(server.URL, nil))

			// Act
			resp, err := client.GetProtocolMetrics(context.Background())

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Equal(tt.wantResult, resp)
			}
		})
	}
}

// TestGetPortfolio tests the GetPortfolio method.
func (s *DLMMClientTestSuite) TestGetPortfolio() {
	tests := []struct {
		name       string
		params     *dlmm.GetPortfolioParams
		status     int
		response   any
		wantURL    string
		wantErr    bool
		wantResult *dlmm.GetPortfolioResponse
	}{
		{
			name: "should successfully get portfolio with params",
			params: &dlmm.GetPortfolioParams{
				User:     "user123",
				Page:     ptr(1),
				PageSize: ptr(20),
				DaysBack: ptr(30),
			},
			status: http.StatusOK,
			response: dlmm.GetPortfolioResponse{
				Page:     1,
				PageSize: 20,
				Pools:    []dlmm.PoolPortfolioItem{{PoolAddress: "pool1"}},
			},
			wantURL: "/portfolio?days_back=30&page=1&page_size=20&user=user123",
			wantResult: &dlmm.GetPortfolioResponse{
				Page:     1,
				PageSize: 20,
				Pools:    []dlmm.PoolPortfolioItem{{PoolAddress: "pool1"}},
			},
		},
		{
			name: "should handle error",
			params: &dlmm.GetPortfolioParams{
				User: "invalid_user",
			},
			status:   http.StatusBadRequest,
			response: "Invalid user",
			wantURL:  "/portfolio?user=invalid_user",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()
			client := dlmm.NewClient(httpclient.New(server.URL, nil))

			// Act
			resp, err := client.GetPortfolio(context.Background(), tt.params)

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Equal(tt.wantResult, resp)
			}
		})
	}
}

// TestGetOpenPortfolio tests the GetOpenPortfolio method.
func (s *DLMMClientTestSuite) TestGetOpenPortfolio() {
	tests := []struct {
		name       string
		params     *dlmm.GetOpenPortfolioParams
		status     int
		response   any
		wantURL    string
		wantErr    bool
		wantResult *dlmm.GetOpenPortfolioResponse
	}{
		{
			name: "should successfully get open portfolio with params",
			params: &dlmm.GetOpenPortfolioParams{
				User:          "user123",
				Page:          ptr(1),
				PageSize:      ptr(20),
				SortDirection: ptr(dlmm.SortDirectionAsc),
				SortBy:        ptr(dlmm.SortByCurrentBalances),
			},
			status: http.StatusOK,
			response: dlmm.GetOpenPortfolioResponse{
				Page:     1,
				PageSize: 20,
				Pools:    []dlmm.PoolOpenPortfolioItem{{PoolAddress: "pool1"}},
			},
			wantURL: "/portfolio/open?page=1&page_size=20&sort_by=current_balances&sort_direction=asc&user=user123",
			wantResult: &dlmm.GetOpenPortfolioResponse{
				Page:     1,
				PageSize: 20,
				Pools:    []dlmm.PoolOpenPortfolioItem{{PoolAddress: "pool1"}},
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()
			client := dlmm.NewClient(httpclient.New(server.URL, nil))

			// Act
			resp, err := client.GetOpenPortfolio(context.Background(), tt.params)

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Equal(tt.wantResult, resp)
			}
		})
	}
}

// TestGetPortfolioTotal tests the GetPortfolioTotal method.
func (s *DLMMClientTestSuite) TestGetPortfolioTotal() {
	tests := []struct {
		name       string
		user       string
		status     int
		response   any
		wantURL    string
		wantErr    bool
		wantResult *dlmm.PortfolioTotalResponse
	}{
		{
			name:   "should successfully get portfolio total",
			user:   "user123",
			status: http.StatusOK,
			response: dlmm.PortfolioTotalResponse{
				TotalPnLUsd: "100.5",
			},
			wantURL: "/portfolio/total?user=user123",
			wantResult: &dlmm.PortfolioTotalResponse{
				TotalPnLUsd: "100.5",
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()
			client := dlmm.NewClient(httpclient.New(server.URL, nil))

			// Act
			resp, err := client.GetPortfolioTotal(context.Background(), tt.user)

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Equal(tt.wantResult, resp)
			}
		})
	}
}
