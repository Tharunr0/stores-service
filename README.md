# Stores Service

A minimal, production-oriented **Go + GraphQL commerce service** inspired by Sticker Mule’s **Stores** platform.

This project demonstrates how I design, build, and ship real backend services: clean architecture, readable code, and infrastructure that’s easy to run locally and deploy in production.

---

## Why this project exists

Sticker Mule builds simple, powerful tools that help people earn money online.  
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

