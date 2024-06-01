## Ejaw Test Case

This project is a sample web application for managing users, sellers, products, customers, and orders. It demonstrates the implementation of a RESTful API using Go, PostgreSQL, and Docker.

## Project Structure

```
ejaw_test_case/
├── cmd/
├── internal/
├── pkg/
├── .dockerignore
├── .env.example
├── .gitignore
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum

```

## Setup
1. Clone the repository: 
```
git clone https://github.com/timofeyka25/ejaw_test_case.git
cd ejaw_test_case
```
2. Copy the example environment variables file and update it with your settings:
```
cp .env.example .env
```
3. Start the application using Docker Compose:
```
docker-compose up --build
```

## Database Schema
![Diagram](images/database%20schema.png)

## API endpoints
    User Authentication:
        POST /auth/signin: Sign in a user.
        POST /auth/signup: Sign up a new user.

    Products:
        GET /products/all: Get a list of products.
        GET /products/{id}: Get a specific product by ID.
        POST /products/: Create a new product.
        PUT /products/{id}: Update an existing product.
        DELETE /products/{id}: Delete a product.

    Sellers:
        GET /sellers/all: Get a list of sellers.
        GET /sellers/{id}: Get a specific seller by ID.
        POST /sellers/: Create a new seller.
        PUT /sellers/{id}: Update an existing seller.
        DELETE /sellers/{id}: Delete a seller.
