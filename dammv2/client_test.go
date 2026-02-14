package dammv2_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/ua1984/meteora-go/dammv2"
	"github.com/ua1984/meteora-go/internal/httpclient"
)

type DammV2ClientTestSuite struct {
	suite.Suite
}

func TestDammV2Client(t *testing.T) {
	suite.Run(t, new(DammV2ClientTestSuite))
}

func (s *DammV2ClientTestSuite) setupTestServer(method, wantURL string, status int, response any) *httptest.Server {
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

func (s *DammV2ClientTestSuite) TestListPools() {
	tests := []struct {
		name       string
		params     *dammv2.ListPoolsParams
		response   any
		status     int
		wantErr    bool
		wantURL    string
		wantResult int // Number of pools in response
	}{
		{
			name: "should successfully list pools with params",
			params: &dammv2.ListPoolsParams{
				Page:      ptr(1),
				Limit:     ptr(10),
				SortBy:    ptr("tvl"),
				SortOrder: ptr("desc"),
			},
			response: dammv2.PaginatedResponse[dammv2.Pool]{
				Data:        []dammv2.Pool{{Address: "pool1"}, {Address: "pool2"}},
				Total:       2,
				CurrentPage: 1,
				PageSize:    10,
			},
			status:     http.StatusOK,
			wantURL:    "/pools?limit=10&page=1&sort_by=tvl&sort_order=desc",
			wantResult: 2,
		},
		{
			name:       "should return error on API failure",
			status:     http.StatusInternalServerError,
			response:   "Internal Server Error",
			wantErr:    true,
			wantURL:    "/pools",
			wantResult: 0,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()

			httpClient := httpclient.New(server.URL, nil)
			client := dammv2.NewClient(httpClient)

			// Act
			resp, err := client.ListPools(context.TODO(), tt.params)

			// Assert
			if tt.wantErr {
				s.Error(err)
				s.Nil(resp)
			} else {
				s.NoError(err)
				s.NotNil(resp)
				s.Len(resp.Data, tt.wantResult)
			}
		})
	}
}

func (s *DammV2ClientTestSuite) TestListGroups() {
	tests := []struct {
		name       string
		params     *dammv2.ListGroupsParams
		response   any
		status     int
		wantErr    bool
		wantURL    string
		wantResult int
	}{
		{
			name: "should successfully list groups",
			params: &dammv2.ListGroupsParams{
				Page: ptr(1),
			},
			response: dammv2.PaginatedResponse[dammv2.PoolGroup]{
				Data: []dammv2.PoolGroup{{GroupName: "group1"}},
			},
			status:     http.StatusOK,
			wantURL:    "/pools/groups?page=1",
			wantResult: 1,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()

			httpClient := httpclient.New(server.URL, nil)
			client := dammv2.NewClient(httpClient)

			// Act
			resp, err := client.ListGroups(context.TODO(), tt.params)

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

func (s *DammV2ClientTestSuite) TestGetGroup() {
	tests := []struct {
		name              string
		lexicalOrderMints string
		params            *dammv2.GetGroupParams
		response          any
		status            int
		wantErr           bool
		wantURL           string
		wantResult        int
	}{
		{
			name:              "should successfully get group pools",
			lexicalOrderMints: "mint1-mint2",
			response: dammv2.PaginatedResponse[dammv2.Pool]{
				Data: []dammv2.Pool{{Address: "pool1"}},
			},
			status:     http.StatusOK,
			wantURL:    "/pools/groups/mint1-mint2",
			wantResult: 1,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()

			httpClient := httpclient.New(server.URL, nil)
			client := dammv2.NewClient(httpClient)

			// Act
			resp, err := client.GetGroup(context.TODO(), tt.lexicalOrderMints, tt.params)

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

func (s *DammV2ClientTestSuite) TestGetPool() {
	tests := []struct {
		name     string
		address  string
		response any
		status   int
		wantErr  bool
		wantURL  string
	}{
		{
			name:    "should successfully get pool",
			address: "pool123",
			response: dammv2.Pool{
				Address: "pool123",
			},
			status:  http.StatusOK,
			wantURL: "/pools/pool123",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()

			httpClient := httpclient.New(server.URL, nil)
			client := dammv2.NewClient(httpClient)

			// Act
			resp, err := client.GetPool(context.TODO(), tt.address)

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Equal(tt.address, resp.Address)
			}
		})
	}
}

func (s *DammV2ClientTestSuite) TestGetOHLCV() {
	tests := []struct {
		name     string
		address  string
		params   *dammv2.OHLCVParams
		response any
		status   int
		wantErr  bool
		wantURL  string
	}{
		{
			name:    "should successfully get OHLCV",
			address: "pool123",
			params: &dammv2.OHLCVParams{
				Resolution: ptr("1H"),
			},
			response: dammv2.OHLCVResponse{
				Data: []dammv2.OHLCV{{Timestamp: 123456}},
			},
			status:  http.StatusOK,
			wantURL: "/pools/pool123/ohlcv?resolution=1H",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()

			httpClient := httpclient.New(server.URL, nil)
			client := dammv2.NewClient(httpClient)

			// Act
			resp, err := client.GetOHLCV(context.TODO(), tt.address, tt.params)

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.NotEmpty(resp.Data)
			}
		})
	}
}

func (s *DammV2ClientTestSuite) TestGetVolumeHistory() {
	tests := []struct {
		name     string
		address  string
		params   *dammv2.VolumeHistoryParams
		response any
		status   int
		wantErr  bool
		wantURL  string
	}{
		{
			name:    "should successfully get volume history",
			address: "pool123",
			response: dammv2.VolumeHistoryResponse{
				Data: []dammv2.VolumeHistory{{Timestamp: 123456}},
			},
			status:  http.StatusOK,
			wantURL: "/pools/pool123/volume/history",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()

			httpClient := httpclient.New(server.URL, nil)
			client := dammv2.NewClient(httpClient)

			// Act
			resp, err := client.GetVolumeHistory(context.TODO(), tt.address, tt.params)

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.NotEmpty(resp.Data)
			}
		})
	}
}

func (s *DammV2ClientTestSuite) TestGetProtocolMetrics() {
	tests := []struct {
		name     string
		response any
		status   int
		wantErr  bool
		wantURL  string
	}{
		{
			name: "should successfully get protocol metrics",
			response: dammv2.ProtocolMetrics{
				TotalTVL: 1000,
			},
			status:  http.StatusOK,
			wantURL: "/stats/protocol_metrics",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			server := s.setupTestServer(http.MethodGet, tt.wantURL, tt.status, tt.response)
			defer server.Close()

			httpClient := httpclient.New(server.URL, nil)
			client := dammv2.NewClient(httpClient)

			// Act
			resp, err := client.GetProtocolMetrics(context.TODO())

			// Assert
			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Equal(1000.0, resp.TotalTVL)
			}
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}
