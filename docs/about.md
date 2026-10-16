# The Walyas Project
The Walyas Project as the official name of our multi-user algorithmic trading and portfolio analytics platform.
The naming convention going forward:

| Component | Name |
|--------|--------|
| Overall initiative | **The Walyas Project** |
| Application | Walyas |
| Repository | `walyas` |
| Go backend | `walyas-api` |
| Future frontend | `walyas-web` |
| Future distributed services | `walyas-market-data`, `walyas-trading`, etc. |

The Walyas Project
│
├── walyas-api          # Go backend (current project)
├── walyas-web          # Next.js frontend (future)
├── walyas-infra        # Terraform, Kubernetes, AWS (future)
└── walyas-research     # Trading algorithms and experiments (future)

# Go Project Structure
Recommended naming for The Walyas Project

| Future component | Name | Responsibility |
|--------|--------|--------|
| API gateway/backend | `walyas-api` | Authentication, REST API |
| Portfolio | `walyas-portfolio` | Portfolios, positions, valuation |
| Market data | `walyas-marketdata` | Alpaca, IBKR, price ingestion |
| Trading | `walyas-trading` | Strategies, orders, execution |
| Analytics | `walyas-analytics` | Risk, performance, backtesting |
| Infrastructure | `walyas-infra` | Terraform, Kubernetes, AWS |