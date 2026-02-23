package dammv1_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/ua1984/meteora-go/dammv1"
	"github.com/ua1984/meteora-go/internal/httpclient"
)

type ClientTestSuite struct {
	suite.Suite
}

func TestClient(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

func (s *ClientTestSuite) setupTestServer(method, wantURL string, status int, response any) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Equal(method, r.Method)
		s.Equal(wantURL, r.URL.String())

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(response)
	}))
}

func (s *ClientTestSuite) TestListPools() {
	wantPools := []dammv1.Pool{
		{PoolAddress: "pool1", PoolName: "SOL-USDC"},
		{PoolAddress: "pool2", PoolName: "USDT-USDC"},
	}

	s.Run("should list pools without address", func() {
		// Arrange
		server := s.setupTestServer(http.MethodGet, "/pools", http.StatusOK, wantPools)
		defer server.Close()

		client := dammv1.NewClient(httpclient.New(server.URL, nil))

		// Act
		pools, err := client.ListPools(context.TODO(), "")

		// Assert
		s.NoError(err)
		s.Equal(wantPools, pools)
	})

	s.Run("should list pools with address", func() {
		// Arrange
		address := "pool1"
		server := s.setupTestServer(http.MethodGet, "/pools?address="+address, http.StatusOK, wantPools[:1])
		defer server.Close()

		client := dammv1.NewClient(httpclient.New(server.URL, nil))

		// Act
		pools, err := client.ListPools(context.TODO(), address)

		// Assert
		s.NoError(err)
		s.Equal(wantPools[:1], pools)
	})
}

func (s *ClientTestSuite) TestSearchPools() {
	wantResult := &dammv1.SearchResult{
		Data:       []dammv1.Pool{{PoolAddress: "pool1"}},
		Page:       1,
		TotalCount: 1,
	}

	s.Run("should search pools with params", func() {
		// Arrange
		page := 1
		size := 10
		searchTerm := "SOL"
		sortBy := "tvl"
		sortOrder := "desc"
		params := &dammv1.SearchParams{
			Page:       &page,
			Size:       &size,
			SearchTerm: &searchTerm,
			SortBy:     &sortBy,
			SortOrder:  &sortOrder,
		}

		wantURL := "/pools/search?page=1&search_term=SOL&size=10&sort_by=tvl&sort_order=desc"
		server := s.setupTestServer(http.MethodGet, wantURL, http.StatusOK, wantResult)
		defer server.Close()

		client := dammv1.NewClient(httpclient.New(server.URL, nil))

		// Act
		result, err := client.SearchPools(context.TODO(), params)

		// Assert
		s.NoError(err)
		s.Equal(wantResult, result)
	})

	s.Run("should search pools without params", func() {
		// Arrange
		server := s.setupTestServer(http.MethodGet, "/pools/search", http.StatusOK, wantResult)
		defer server.Close()

		client := dammv1.NewClient(httpclient.New(server.URL, nil))

		// Act
		result, err := client.SearchPools(context.TODO(), nil)

		// Assert
		s.NoError(err)
		s.Equal(wantResult, result)
	})
}

func (s *ClientTestSuite) TestGetPoolsMetrics() {
	// Arrange
	wantMetrics := &dammv1.PoolMetrics{
		DynamicAMMTVL: 1000000,
	}

	server := s.setupTestServer(http.MethodGet, "/pools-metrics", http.StatusOK, wantMetrics)
	defer server.Close()

	client := dammv1.NewClient(httpclient.New(server.URL, nil))

	// Act
	metrics, err := client.GetPoolsMetrics(context.TODO())

	// Assert
	s.NoError(err)
	s.Equal(wantMetrics, metrics)
}

func (s *ClientTestSuite) TestListPoolConfigs() {
	// Arrange
	wantConfigs := []dammv1.PoolConfig{
		{ConfigAddress: "config1", TradeFeeBPS: 25},
	}

	server := s.setupTestServer(http.MethodGet, "/pool-configs", http.StatusOK, wantConfigs)
	defer server.Close()

	client := dammv1.NewClient(httpclient.New(server.URL, nil))

	// Act
	configs, err := client.ListPoolConfigs(context.TODO())

	// Assert
	s.NoError(err)
	s.Equal(wantConfigs, configs)
}

func (s *ClientTestSuite) TestGetFeeConfig() {
	// Arrange
	configAddr := "config1"
	wantConfigs := []dammv1.FeeConfig{
		{ConfigAddress: configAddr, FeePercentage: "0.25"},
	}

	server := s.setupTestServer(http.MethodGet, "/fee-config/"+configAddr, http.StatusOK, wantConfigs)
	defer server.Close()

	client := dammv1.NewClient(httpclient.New(server.URL, nil))

	// Act
	configs, err := client.GetFeeConfig(context.TODO(), configAddr)

	// Assert
	s.NoError(err)
	s.Equal(wantConfigs, configs)
}

func (s *ClientTestSuite) TestListPoolsWithFarm() {
	wantPools := []dammv1.Pool{
		{PoolAddress: "pool1", PoolName: "SOL-USDC"},
	}

	s.Run("should list pools with farm with params", func() {
		// Arrange
		page := 1
		size := 10
		params := &dammv1.PaginationParams{
			Page: &page,
			Size: &size,
		}

		server := s.setupTestServer(http.MethodGet, "/farm?page=1&size=10", http.StatusOK, wantPools)
		defer server.Close()

		client := dammv1.NewClient(httpclient.New(server.URL, nil))

		// Act
		pools, err := client.ListPoolsWithFarm(context.TODO(), params)

		// Assert
		s.NoError(err)
		s.Equal(wantPools, pools)
	})
}

func (s *ClientTestSuite) TestListAlphaVaults() {
	wantVaults := []dammv1.AlphaVault{
		{VaultAddress: "vault1"},
	}

	s.Run("should list alpha vaults without params", func() {
		// Arrange
		server := s.setupTestServer(http.MethodGet, "/alpha-vault", http.StatusOK, wantVaults)
		defer server.Close()

		client := dammv1.NewClient(httpclient.New(server.URL, nil))

		// Act
		vaults, err := client.ListAlphaVaults(context.TODO(), nil)

		// Assert
		s.NoError(err)
		s.Equal(wantVaults, vaults)
	})

	s.Run("should list alpha vaults with filters", func() {
		// Arrange
		params := &dammv1.AlphaVaultParams{
			VaultAddress: []string{"vault1"},
			PoolAddress:  []string{"pool1"},
			BaseMint:     []string{"mint1"},
		}

		wantURL := "/alpha-vault?base_mint=mint1&pool_address=pool1&vault_address=vault1"
		server := s.setupTestServer(http.MethodGet, wantURL, http.StatusOK, wantVaults)
		defer server.Close()

		client := dammv1.NewClient(httpclient.New(server.URL, nil))

		// Act
		vaults, err := client.ListAlphaVaults(context.TODO(), params)

		// Assert
		s.NoError(err)
		s.Equal(wantVaults, vaults)
	})
}

func (s *ClientTestSuite) TestListAlphaVaultConfigs() {
	// Arrange
	wantConfigs := &dammv1.AlphaVaultConfigs{
		ProrataConfigs: []dammv1.ProrataConfig{{Address: "config1"}},
	}

	server := s.setupTestServer(http.MethodGet, "/alpha-vault-configs", http.StatusOK, wantConfigs)
	defer server.Close()

	client := dammv1.NewClient(httpclient.New(server.URL, nil))

	// Act
	configs, err := client.ListAlphaVaultConfigs(context.TODO())

	// Assert
	s.NoError(err)
	s.Equal(wantConfigs, configs)
}

func (s *ClientTestSuite) TestGetPoolsByVaultLP() {
	// Arrange
	address := "vaultLP1"
	wantPools := []dammv1.Pool{
		{PoolAddress: "pool1"},
	}

	server := s.setupTestServer(http.MethodPost, "/get_pools_by_a_vault_lp?address="+address, http.StatusOK, wantPools)
	defer server.Close()

	client := dammv1.NewClient(httpclient.New(server.URL, nil))

	// Act
	pools, err := client.GetPoolsByVaultLP(context.TODO(), address)

	// Assert
	s.NoError(err)
	s.Equal(wantPools, pools)
}
