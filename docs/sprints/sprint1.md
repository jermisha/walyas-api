Initialize Walya Go module
Sprint 1 — Detailed 100-decision roadmap
This is our immediate priority. We will keep the questions tied to real code changes rather than abstract quizzes.

| Decisions | Engineering focus | Deliverables |
|--------|--------|--------|
| 1–10 | Domain design, Go structs, interfaces, maps, mutexes | Portfolio domain and repository foundation |
| 11–20 | Defensive copies, optimistic locking, services, testing | Version-aware repository and service tests |
| 21–25 | Tenant-safe writes, committed versions, HTTP preconditions, error mapping | Ownership and concurrency contracts |
| **26–35** | Strict JSON parsing, API validation, request/response design | Hardened HTTP API |
| 36–45 | Table-driven tests, fuzzing, race tests, benchmarks | Comprehensive Go testing |
| 46–55 | PostgreSQL, migrations, transactions, isolation | Persistent repository |
| 56–65 | Authentication boundaries, authorization, tenant isolation | Secure multi-user API |
| 66–75 | Configuration, structured logging, context, shutdown | Operationally reliable server |
| 76–85 | Retry strategies, rate limits, idempotency, resilience | Reusable infrastructure components |
| 86–95 | Docker, integration tests, CI, deployment configuration | Reproducible build and test pipeline |
| 96–100 | Refactoring, architecture review, performance, release | Sprint 1 release candidate |