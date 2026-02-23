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

func (s *DLMMClientTestSuite) TestGetClosedPositions() {
	tests := []struct {
		name       string
		wallet     string
		params     *dlmm.GetClosedPositionsParams
		status     int
		response   any
		wantURL    string
		wantErr    bool
		wantResult *dlmm.ClosedPositionsCursorResponse
	}{
		{
			name:   "should successfully get closed positions with params",
			wallet: "wallet1",
			params: &dlmm.GetClosedPositionsParams{
				Limit: ptr(10),
				Pool:  ptr("pool1"),
			},
			status: http.StatusOK,
			response: dlmm.ClosedPositionsCursorResponse{
				Limit: 10,
				Data:  []dlmm.ClosedPosition{{PositionAddress: "pos1", PoolAddress: "pool1"}},
			},
			wantURL: "/wallets/wallet1/closed_positions?limit=10&pool=pool1",
			wantResult: &dlmm.ClosedPositionsCursorResponse{
				Limit: 10,
				Data:  []dlmm.ClosedPosition{{PositionAddress: "pos1", PoolAddress: "pool1"}},
			},
		},
		{
			name:   "should successfully get closed positions without params",
			wallet: "wallet2",
			status: http.StatusOK,
			response: dlmm.ClosedPositionsCursorResponse{
				Limit: 10,
				Data:  []dlmm.ClosedPosition{},
			},
			wantURL: "/wallets/wallet2/closed_positions",
			wantResult: &dlmm.ClosedPositionsCursorResponse{
				Limit: 10,
				Data:  []dlmm.ClosedPosition{},
			},
		},
		{
			name:     "should return error on API failure",
			wallet:   "wallet3",
			status:   http.StatusBadRequest,
			response: "Bad Request",
			wantURL:  "/wallets/wallet3/closed_positions",
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
			resp, err := client.GetClosedPositions(context.Background(), tt.wallet, tt.params)

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

func (s *DLMMClientTestSuite) TestGetOpenPositions() {
	tests := []struct {
		name       string
		wallet     string
		params     *dlmm.GetOpenPositionsParams
		status     int
		response   any
		wantURL    string
		wantErr    bool
		wantResult *dlmm.OpenPositionsResponse
	}{
		{
			name:   "should successfully get open positions with pool filter",
			wallet: "wallet1",
			params: &dlmm.GetOpenPositionsParams{
				Pool: ptr("pool1,pool2"),
			},
			status: http.StatusOK,
			response: dlmm.OpenPositionsResponse{
				TotalPools:     1,
				TotalPositions: 2,
				Data: []dlmm.PositionsByPool{{
					PoolAddress: "pool1",
					Name:        "SOL-USDC",
					Positions:   []dlmm.OpenPosition{{PositionAddress: "posA"}},
				}},
			},
			wantURL: "/wallets/wallet1/open_positions?pool=pool1%2Cpool2",
			wantResult: &dlmm.OpenPositionsResponse{
				TotalPools:     1,
				TotalPositions: 2,
				Data: []dlmm.PositionsByPool{{
					PoolAddress: "pool1",
					Name:        "SOL-USDC",
					Positions:   []dlmm.OpenPosition{{PositionAddress: "posA"}},
				}},
			},
		},
		{
			name:   "should successfully get open positions without params",
			wallet: "wallet2",
			status: http.StatusOK,
			response: dlmm.OpenPositionsResponse{
				TotalPools:     0,
				TotalPositions: 0,
				Data:           []dlmm.PositionsByPool{},
			},
			wantURL: "/wallets/wallet2/open_positions",
			wantResult: &dlmm.OpenPositionsResponse{
				TotalPools:     0,
				TotalPositions: 0,
				Data:           []dlmm.PositionsByPool{},
			},
		},
		{
			name:     "should return error on API failure",
			wallet:   "wallet3",
			status:   http.StatusBadRequest,
			response: "Bad Request",
			wantURL:  "/wallets/wallet3/open_positions",
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
			resp, err := client.GetOpenPositions(context.Background(), tt.wallet, tt.params)

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

func (s *DLMMClientTestSuite) TestGetPositionHistoricalEvents() {
	tests := []struct {
		name       string
		address    string
		params     *dlmm.GetPositionHistoricalEventsParams
		status     int
		response   any
		wantURL    string
		wantErr    bool
		wantResult *dlmm.GetPositionHistoricalEventsResponse
	}{
		{
			name:    "should successfully get historical events with params",
			address: "pos1",
			params: &dlmm.GetPositionHistoricalEventsParams{
				EventType:      ptr(dlmm.PositionEventTypeAdd),
				OrderDirection: ptr(dlmm.PositionEventOrderDirectionDesc),
			},
			status: http.StatusOK,
			response: dlmm.GetPositionHistoricalEventsResponse{
				Events: []dlmm.PositionEvent{{Signature: "sig1", EventType: "add"}},
			},
			wantURL: "/positions/pos1/historical?event_type=add&order_direction=desc",
			wantResult: &dlmm.GetPositionHistoricalEventsResponse{
				Events: []dlmm.PositionEvent{{Signature: "sig1", EventType: "add"}},
			},
		},
		{
			name:    "should successfully get historical events without params",
			address: "pos2",
			status: http.StatusOK,
			response: dlmm.GetPositionHistoricalEventsResponse{
				Events: []dlmm.PositionEvent{},
			},
			wantURL: "/positions/pos2/historical",
			wantResult: &dlmm.GetPositionHistoricalEventsResponse{
				Events: []dlmm.PositionEvent{},
			},
		},
		{
			name:    "should return error on API failure",
			address: "pos3",
			status:  http.StatusInternalServerError,
			response: "Error",
			wantURL: "/positions/pos3/historical",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()
			client := dlmm.NewClient(httpclient.New(server.URL, nil))

			// Act
			resp, err := client.GetPositionHistoricalEvents(context.Background(), tt.address, tt.params)

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

func (s *DLMMClientTestSuite) TestGetPositionTotalClaimFees() {
	tests := []struct {
		name       string
		address    string
		status     int
		response   any
		wantURL    string
		wantErr    bool
		wantResult []dlmm.PositionTotalClaimFees
	}{
		{
			name:    "should successfully get total claim fees",
			address: "pos1",
			status:  http.StatusOK,
			response: []dlmm.PositionTotalClaimFees{
				{PoolAddress: "pool1", TotalFeeX: "100", TotalFeeY: "200"},
			},
			wantURL: "/positions/pos1/total_claim_fees",
			wantResult: []dlmm.PositionTotalClaimFees{
				{PoolAddress: "pool1", TotalFeeX: "100", TotalFeeY: "200"},
			},
		},
		{
			name:     "should return error on API failure",
			address:  "pos2",
			status:   http.StatusNotFound,
			response: "Not Found",
			wantURL:  "/positions/pos2/total_claim_fees",
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
			resp, err := client.GetPositionTotalClaimFees(context.Background(), tt.address)

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

func (s *DLMMClientTestSuite) TestGetPoolPositionPnL() {
	tests := []struct {
		name        string
		poolAddress string
		params      *dlmm.GetPoolPositionPnLParams
		status      int
		response    any
		wantURL     string
		wantErr     bool
		wantResult  *dlmm.GetPoolPositionPnLResponse
	}{
		{
			name:        "should successfully get pool position pnl with all params",
			poolAddress: "pool1",
			params: &dlmm.GetPoolPositionPnLParams{
				User:     "user123",
				Status:   ptr(dlmm.PositionStatusOpen),
				Page:     ptr(1),
				PageSize: ptr(20),
			},
			status: http.StatusOK,
			response: dlmm.GetPoolPositionPnLResponse{
				Page:        1,
				PageSize:    20,
				TotalCount:  1,
				TokenXPrice: "1.0",
				TokenYPrice: "2.0",
				RewardTokenXPrice: "3.0",
				RewardTokenYPrice: "4.0",
				Positions: []dlmm.PositionPnLData{{PositionAddress: "posA", IsClosed: false}},
			},
			wantURL: "/positions/pool1/pnl?page=1&page_size=20&status=open&user=user123",
			wantResult: &dlmm.GetPoolPositionPnLResponse{
				Page:        1,
				PageSize:    20,
				TotalCount:  1,
				TokenXPrice: "1.0",
				TokenYPrice: "2.0",
				RewardTokenXPrice: "3.0",
				RewardTokenYPrice: "4.0",
				Positions: []dlmm.PositionPnLData{{PositionAddress: "posA", IsClosed: false}},
			},
		},
		{
			name:        "should return error on API failure",
			poolAddress: "pool2",
			params: &dlmm.GetPoolPositionPnLParams{
				User: "user456",
			},
			status:   http.StatusBadRequest,
			response: "Bad Request",
			wantURL:  "/positions/pool2/pnl?user=user456",
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
			resp, err := client.GetPoolPositionPnL(context.Background(), tt.poolAddress, tt.params)

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
