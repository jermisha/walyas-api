# The Walyas Project 

# walyas-api initial structure

walyas-api/
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── identity/
│   │   ├── user.go
│   │   └── user_test.go
│   │
│   ├── portfolio/
│   │   ├── portfolio.go
│   │   ├── portfolio_test.go
│   │   ├── repository.go
│   │   ├── service.go
│   │   ├── service_test.go
│   │   │
│   │   ├── memory/
│   │   │   ├── repository.go
│   │   │   └── repository_test.go
│   │   │
│   │   └── postgres/
│   │       ├── repository.go
│   │       └── repository_integration_test.go
│   │
│   ├── marketdata/
│   │   ├── provider.go
│   │   ├── service.go
│   │   ├── service_test.go
│   │   ├── alpaca/
│   │   │   ├── client.go
│   │   │   └── client_test.go
│   │   └── ibkr/
│   │       ├── client.go
│   │       └── client_test.go
│   │
│   ├── trading/
│   │   ├── strategy.go
│   │   ├── strategy_test.go
│   │   ├── backtest.go
│   │   └── backtest_test.go
│   │
│   ├── analytics/
│   │   ├── returns.go
│   │   ├── returns_test.go
│   │   ├── risk.go
│   │   └── risk_test.go
│   │
│   ├── platform/
│   │   ├── database/
│   │   └── httpserver/
│   │
│   └── testutil/
│       ├── fixtures.go
│       └── factories.go
│
├── tests/
│   ├── e2e/
│   │   └── portfolio_test.go
│   ├── performance/
│   │   └── ingestion_test.go
│   └── fixtures/
│       └── market_bars.json
│
├── go.mod
└── README.md