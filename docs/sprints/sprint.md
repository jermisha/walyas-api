# Sprin1

What comes next?
Step 1 — Initialize Walya Go module
Step 2 — Implement domain entities
Step 3 — Build multi-tenant portfolio access
Step 4 — Add PostgreSQL persistence
Step 5 — Integrate market-data providers


Sprint overview
These are proposed learning phases, not fixed calendar weeks. Each sprint ends with runnable code, automated tests, and a review of our engineering choices.
Sprint 1 — Go Backend Engineering (In progress)
100 engineering decisions · Portfolio API, testing, concurrency, PostgreSQL, security boundaries, HTTP, Docker
Outcome: A reliable, testable portfolio API.
Sprint 2 — Identity, Investment Profiles & Portfolio Management
Users, brokerage accounts, positions, transactions, authentication, investment preferences, and tenant isolation.
Outcome: Multi-user portfolio management.
Sprint 3 — Financial Common Data Model & Ontology
Alpaca and IBKR adapters, provider-independent market data, financial entity identity, FIBO concepts, ISO 20022, FIX, OpenFIGI, and schema evolution.
Outcome: Walyas Financial Common Schema (WFCS).
Sprint 4 — Market Data Ingestion & Processing
Historical and streaming ingestion, pagination, rate limits, deduplication, event time, backfills, storage, and data quality.
Outcome: Reproducible financial-data pipelines.
Sprint 5 — Quantitative Analytics & Backtesting
Returns, indicators, strategy interfaces, backtesting, transaction costs, look-ahead bias, performance metrics, and risk calculations.
Outcome: A deterministic strategy research engine.
Sprint 6 — Order Management & Paper Trading
Order lifecycle, broker adapters, idempotency, reconciliation, risk checks, audit logs, and paper trading.
Outcome: A controlled paper-trading platform.
Sprint 7 — Distributed Systems & Cloud Deployment
Kafka, outbox pattern, event contracts, distributed consistency, observability, Terraform, Kubernetes, and AWS.
Outcome: An independently scalable event-driven architecture where justified.
Sprint 8 — AI Enrichment & Knowledge Graphs
LLM enrichment, financial document processing, entity resolution, semantic search, RAG, provenance, and knowledge graphs.
Outcome: Explainable financial intelligence.
Sprint 9 — Supply Chain Management
Suppliers, products, procurement, inventory, shipments, GS1 identifiers, EPCIS events, and cross-domain relationships.
Outcome: Supply chain visibility connected to financial analytics.
Sprint 10 — Full Platform Integration
Next.js frontend, dashboards, API security, performance, end-to-end tests, operational controls, and production readiness.
Outcome: Integrated Walyas platform.