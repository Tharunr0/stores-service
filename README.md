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

---

## Example GraphQL Query

graphql
{
  stores {
    id
    name
    revenue
  }
}
Running locally
Requirements

Go 1.22+

PostgreSQL

(Optional) Redis

Docker (optional)

Environment
export POSTGRES_DSN="postgres://postgres:postgres@localhost:5432/stores?sslmode=disable"

Start the server
go run ./cmd/server


Open:

http://localhost:8080


You’ll get GraphQL Playground out of the box.

Docker
docker build -t stores-service .
docker run -p 8080:8080 \
  -e POSTGRES_DSN="postgres://..." \
  stores-service


The resulting image is small, fast, and production-ready.

Engineering decisions
Why GraphQL

Strong contract between frontend and backend

Ideal for multi-client platforms (web + mobile)

Eliminates over-fetching for commerce UIs

Why Go

Simple concurrency model

Predictable performance

Excellent fit for platform and infrastructure work

Why explicit SQL

Full control over queries

Easier performance tuning

No hidden ORM behavior

What I’m proud of

Clarity over cleverness — this code is easy to read and reason about

End-to-end ownership — API, data modeling, infra, and deployment

Production mindset — everything here could ship as-is

Low ceremony — no unnecessary abstractions or frameworks

This reflects how I like to work: ship quickly, keep things understandable, and build software that other engineers enjoy maintaining.

What I’d build next

Redis-backed query caching

Write-side mutations with validation

Background jobs for analytics / notifications

Kubernetes manifests (GKE-ready)

Observability (structured logging + metrics)

About me

I’m a software engineer who enjoys building real products used by real people.
I thrive in low-process environments where ownership, execution, and quality matter more than buzzwords.

If this looks like how you like to build software too, I’d love to work together.

Author: [Your Name]
GitHub: https://github.com/yourusername


---

### Why this README works for hiring managers

- Explains **intent**, not just tools  
- Shows **product thinking**, not tutorial code  
- Communicates **ownership and taste**
- Matches Sticker Mule’s **“no bullshit” culture**
- Makes reviewers think: *“This person could ship on day one.”*

If you want next:
- A **Sticker Mule–specific short version**
- A **GitHub profile cleanup strategy**
- Or a **final polish pass before submission**

Say the word.
