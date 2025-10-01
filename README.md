# Udhhar App

**Udhhar** is a simple yet powerful app to track money lent or borrowed among friends. Built with **Go**, **Gin**, and **Redis**, it helps you manage personal debts efficiently with speed and reliability.

---

## Features

* **Track Transactions** – Record money borrowed or lent with friends.
* **Status Management** – Mark transactions as `pending`, `completed`, or `overdue`.
* **Fast Access** – Redis caching ensures quick retrieval of frequent queries.
* **RESTful API** – Clean and robust backend endpoints using Go Gin.
* **JSON Support** – All requests and responses are in JSON format for easy integration.

---

## Tech Stack

* **Backend:** Go (Gin framework)
* **Database:** MongoDB
* **Caching:** Redis
* **Environment Management:** `godotenv`

---

## Getting Started

### Prerequisites

* Go 1.25+
* MongoDB
* Redis

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/<username>/udhaar-app.git
   cd udhaar-app
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up environment variables in `.env` file:

   ```env
   PORT=8080
   MONGO_URI=<your_mongo_uri>
   REDIS_ADDR=<your_redis_address>
   REDIS_PASSWORD=<your_redis_password>
   ```

4. Run the application:

   ```bash
   go run main.go
   ```

---

## API Endpoints

* `POST /transaction` – Create a new transaction
* `GET /transaction/:id` – Get a transaction by ID
* `PUT /transaction/:id` – Update transaction status/details
* `DELETE /transaction/:id` – Delete a transaction

All endpoints accept and return **JSON** payloads.

---

## Future Improvements

* **Authentication** – Add user accounts and secure endpoints.
* **Frontend Dashboard** – Visualize transactions and summaries.
* **Notifications** – Remind users of pending debts.

---

## Contributing

Contributions are welcome! Please fork the repository, make changes, and submit a pull request.

---
