# Stores Service

A minimal, production-oriented **Go + GraphQL commerce service** 

This project demonstrates how I design, build, and ship real backend services: clean architecture, readable code, and infrastructure that’s easy to run locally and deploy in production.

---

## Why this project exists
  
This service models a small but realistic slice of that ecosystem: a backend service that powers storefront data via a GraphQL API.

The focus here is not flashy features—it’s **clarity, correctness, and scalability**.

---

## Tech Stack

- **Go** — backend service and business logic
- **GraphQL (gqlgen)** — strongly typed API for web & mobile clients
- **PostgreSQL** — primary data store
- **Redis** — caching layer (ready for use)
- **Docker** — containerized for production
- **Cloud-ready** — designed to run cleanly on Kubernetes / GCP

---

## Project Structure
stores-service/
├── cmd/server/ # Application entrypoint
├── internal/
│ ├── db/ # PostgreSQL connection handling
│ ├── cache/ # Redis client setup
│ ├── graph/ # GraphQL schema & resolvers
│ └── store/ # Domain models
├── Dockerfile
└── go.mod

This structure keeps boundaries explicit and avoids framework magic.

---

## What this service does

- Exposes a **GraphQL API** to fetch store data
- Reads from **PostgreSQL** using explicit SQL
- Designed to support **web (React)** and **mobile (React Native / Expo)** clients
- Easy to extend with:
  - caching
  - background jobs
  - additional services (Ship, Give, Notify–style tools)



