# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased] - TBD

## [2.0.0] - 2026-02-23

### Added

- Added DLMM portfolio endpoints: `/portfolio`, `/portfolio/open`, and `/portfolio/total`
- Added DLMM positions endpoints: `/positions/{address}/historical`, `/positions/{address}/total_claim_fees`, and `/positions/{pool_address}/pnl`
- Added DLMM wallet endpoints: `/wallets/{wallet}/closed_positions` and `/wallets/{wallet}/open_positions`
- Added DAMM v2 wallet endpoints: `/wallets/{wallet}/closed_positions` and `/wallets/{wallet}/open_positions`

### Changed

- Remove DLMM legacy API support

### Fixed

- DLMM endpoints now use correct parameters
- DAMM v2 endpoints now use correct parameters
- DAMM v1 endpoints now use correct parameters
- Stake2Earn endpoints now use correct parameters
- API error responses now correctly parse the `{"message": "..."}` JSON body instead of returning the raw body string

## [1.1.0] - 2026-02-17

### Added

- Added helper functions for easier pointer parameter passing
- Added retry configuration to the HTTP client

## [1.0.0] - 2026-02-14

### Added

- HTTP client implementation for Meteora API communication
- DLMM (Dynamic Liquidity Market Maker) API support
- DAMM v1 protocol implementation
- DAMM v2 protocol implementation
- Stake2Earn APIs integration
- DynamicVault APIs support
- Meteora Go client library with comprehensive endpoint coverage
- Basic usage examples demonstrating DLMM, DAMM, Stake2Earn, and Dynamic Vault endpoints
- Full README documentation for the Go client library
- GitHub Actions workflow for automated testing across multiple OS and Go versions
- MIT License
- Test coverage for all client modules

### Changed

- Refactored client test files to use external test packages for better encapsulation

[2.0.0]: https://github.com/ua1984/meteora-go/compare/v1.1.0...v2.0.0
[1.1.0]: https://github.com/ua1984/meteora-go/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/ua1984/meteora-go/commit/7fcde5e295143aca2565a255337cc4bd2ba9e73a
