package dlmm

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
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
		params     *ListPoolsParams
		status     int
		response   any
		wantURL    string
		wantErr    bool
		wantResult *PaginatedResponse[Pool]
	}{
		{
			name: "should successfully list pools with params",
			params: &ListPoolsParams{
				Page:  ptr(1),
				Limit: ptr(10),
			},
			status: http.StatusOK,
			response: PaginatedResponse[Pool]{
				Data: []Pool{{Address: "pool1"}},
			},
			wantURL: "/pools?limit=10&page=1",
			wantResult: &PaginatedResponse[Pool]{
				Data: []Pool{{Address: "pool1"}},
			},
		},
		{
			name:       "should successfully list pools without params",
			status:     http.StatusOK,
			response:   PaginatedResponse[Pool]{Data: []Pool{}},
			wantURL:    "/pools",
			wantResult: &PaginatedResponse[Pool]{Data: []Pool{}},
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
			client := NewClient(httpclient.New(server.URL, nil), nil)

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
		params     *ListGroupsParams
		status     int
		response   any
		wantURL    string
		wantErr    bool
		wantResult *PaginatedResponse[PoolGroup]
	}{
		{
			name: "should successfully list groups with params",
			params: &ListGroupsParams{
				Page: ptr(2),
			},
			status: http.StatusOK,
			response: PaginatedResponse[PoolGroup]{
				Data: []PoolGroup{{GroupName: "SOL-USDC"}},
			},
			wantURL: "/pools/groups?page=2",
			wantResult: &PaginatedResponse[PoolGroup]{
				Data: []PoolGroup{{GroupName: "SOL-USDC"}},
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
			client := NewClient(httpclient.New(server.URL, nil), nil)

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
		params            *GetGroupParams
		status            int
		response          any
		wantURL           string
		wantErr           bool
		wantResult        *PaginatedResponse[Pool]
	}{
		{
			name:              "should successfully get group",
			lexicalOrderMints: "mint1-mint2",
			params: &GetGroupParams{
				Limit: ptr(5),
			},
			status: http.StatusOK,
			response: PaginatedResponse[Pool]{
				Data: []Pool{{Address: "poolA"}},
			},
			wantURL: "/pools/groups/mint1-mint2?limit=5",
			wantResult: &PaginatedResponse[Pool]{
				Data: []Pool{{Address: "poolA"}},
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()
			client := NewClient(httpclient.New(server.URL, nil), nil)

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
		wantResult *Pool
	}{
		{
			name:       "should successfully get pool",
			address:    "addr123",
			status:     http.StatusOK,
			response:   Pool{Address: "addr123"},
			wantURL:    "/pools/addr123",
			wantResult: &Pool{Address: "addr123"},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()
			client := NewClient(httpclient.New(server.URL, nil), nil)

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
		params     *OHLCVParams
		status     int
		response   any
		wantURL    string
		wantErr    bool
		wantResult *OHLCVResponse
	}{
		{
			name:    "should successfully get ohlcv with resolution",
			address: "pool1",
			params: &OHLCVParams{
				Resolution: ptr("1H"),
			},
			status: http.StatusOK,
			response: OHLCVResponse{
				Data: []OHLCV{{Timestamp: 1000}},
			},
			wantURL: "/pools/pool1/ohlcv?resolution=1H",
			wantResult: &OHLCVResponse{
				Data: []OHLCV{{Timestamp: 1000}},
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()
			client := NewClient(httpclient.New(server.URL, nil), nil)

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
		params     *VolumeHistoryParams
		status     int
		response   any
		wantURL    string
		wantErr    bool
		wantResult *VolumeHistoryResponse
	}{
		{
			name:    "should successfully get volume history",
			address: "poolX",
			params: &VolumeHistoryParams{
				Limit: ptr(20),
			},
			status: http.StatusOK,
			response: VolumeHistoryResponse{
				Data: []VolumeHistory{{Timestamp: 2000}},
			},
			wantURL: "/pools/poolX/volume/history?limit=20",
			wantResult: &VolumeHistoryResponse{
				Data: []VolumeHistory{{Timestamp: 2000}},
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()
			client := NewClient(httpclient.New(server.URL, nil), nil)

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
		wantResult *ProtocolMetrics
	}{
		{
			name:       "should successfully get protocol metrics",
			status:     http.StatusOK,
			response:   ProtocolMetrics{TotalVolume: 999.9},
			wantURL:    "/stats/protocol_metrics",
			wantResult: &ProtocolMetrics{TotalVolume: 999.9},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()
			client := NewClient(httpclient.New(server.URL, nil), nil)

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

func (s *DLMMClientTestSuite) TestListAllPairs() {
	tests := []struct {
		name       string
		status     int
		response   any
		wantURL    string
		wantErr    bool
		wantResult []LegacyPair
	}{
		{
			name:       "should successfully list all pairs from legacy api",
			status:     http.StatusOK,
			response:   []LegacyPair{{Address: "legacy1"}},
			wantURL:    "/pair/all",
			wantResult: []LegacyPair{{Address: "legacy1"}},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()
			client := NewClient(nil, httpclient.New(server.URL, nil))

			// Act
			resp, err := client.ListAllPairs(context.Background())

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
