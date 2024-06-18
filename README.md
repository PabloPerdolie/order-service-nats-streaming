# Order Service with NATS Streaming

This repository contains an experimental order service built using GoLang, NATS Streaming, Docker, and PostgreSQL. It is designed to handle order processing with reliable message streaming and persistence.

## Features

- Order processing service
- NATS Streaming for messaging
- PostgreSQL for persistent storage
- Docker for containerization

## Technologies Used

- **Programming Language**: Go (Golang)
- **Messaging**: NATS Streaming
- **Database**: PostgreSQL
- **Containerization**: Docker

## Getting Started

#### To run PostgreSql and nats-streaming:
```bash
docker compose up
```
#### To run nats-streaming publisher to publish at least ONE order:
```bash
go run .\nats-streaming-publish\publisher.go
```
#### To run main service:
```bash
go run .\order-service\cmd\main.go
```
### Prerequisites

- Docker and Docker Compose
- Go 1.16+
- NATS Streaming

### Installation
Clone the repository:
```bash
git clone https://github.com/PabloPerdolie/order-service-nats-streaming.git
```

Thank you for checking out my project! Feel free to reach out if you have any questions or feedback.


   



