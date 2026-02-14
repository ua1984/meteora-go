package dammv2

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
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
		params     *ListPoolsParams
		response   any
		status     int
		wantErr    bool
		wantURL    string
		wantResult int // Number of pools in response
	}{
		{
			name: "should successfully list pools with params",
			params: &ListPoolsParams{
				Page:      ptr(1),
				Limit:     ptr(10),
				SortBy:    ptr("tvl"),
				SortOrder: ptr("desc"),
			},
			response: PaginatedResponse[Pool]{
				Data:        []Pool{{Address: "pool1"}, {Address: "pool2"}},
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
			client := NewClient(httpClient)

			// Act
			resp, err := client.ListPools(s.T().Context(), tt.params)

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
		params     *ListGroupsParams
		response   any
		status     int
		wantErr    bool
		wantURL    string
		wantResult int
	}{
		{
			name: "should successfully list groups",
			params: &ListGroupsParams{
				Page: ptr(1),
			},
			response: PaginatedResponse[PoolGroup]{
				Data: []PoolGroup{{GroupName: "group1"}},
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
			client := NewClient(httpClient)

			// Act
			resp, err := client.ListGroups(s.T().Context(), tt.params)

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
		params            *GetGroupParams
		response          any
		status            int
		wantErr           bool
		wantURL           string
		wantResult        int
	}{
		{
			name:              "should successfully get group pools",
			lexicalOrderMints: "mint1-mint2",
			response: PaginatedResponse[Pool]{
				Data: []Pool{{Address: "pool1"}},
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
			client := NewClient(httpClient)

			// Act
			resp, err := client.GetGroup(s.T().Context(), tt.lexicalOrderMints, tt.params)

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
			response: Pool{
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
			client := NewClient(httpClient)

			// Act
			resp, err := client.GetPool(s.T().Context(), tt.address)

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
		params   *OHLCVParams
		response any
		status   int
		wantErr  bool
		wantURL  string
	}{
		{
			name:    "should successfully get OHLCV",
			address: "pool123",
			params: &OHLCVParams{
				Resolution: ptr("1H"),
			},
			response: OHLCVResponse{
				Data: []OHLCV{{Timestamp: 123456}},
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
			client := NewClient(httpClient)

			// Act
			resp, err := client.GetOHLCV(s.T().Context(), tt.address, tt.params)

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
		params   *VolumeHistoryParams
		response any
		status   int
		wantErr  bool
		wantURL  string
	}{
		{
			name:    "should successfully get volume history",
			address: "pool123",
			response: VolumeHistoryResponse{
				Data: []VolumeHistory{{Timestamp: 123456}},
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
			client := NewClient(httpClient)

			// Act
			resp, err := client.GetVolumeHistory(s.T().Context(), tt.address, tt.params)

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
			response: ProtocolMetrics{
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
			client := NewClient(httpClient)

			// Act
			resp, err := client.GetProtocolMetrics(s.T().Context())

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
