# Go CRUD Application

## Overview

This project is a simple CRUD (Create, Read, Update, Delete) application built using Go. It demonstrates how to build a RESTful API using Go with basic CRUD functionalities for managing data entries.

## Project Structure

- **cmd/**: Contains the main application entry point.
  - **main.go**: The main file that initializes and starts the application.
  
- **pkg/**: Contains the application's core logic and helper functions.
  - **handlers/**: Contains the HTTP handlers for each endpoint.
    - **handlers.go**: Defines the functions to handle HTTP requests.
  - **models/**: Contains the data models.
    - **models.go**: Defines the data structures used in the application.
  - **repository/**: Contains the database interactions.
    - **repository.go**: Implements CRUD operations for the database.
  - **router/**: Contains the router setup.
    - **router.go**: Defines the routes and associates them with handlers.

- **vendor/**: Contains the project's dependencies. (Managed by Go modules)

## Setup Instructions

1. **Clone the Repository**
    ```sh
    git clone https://github.com/rsengaravua/go-crud.git
    cd go-crud
    ```

2. **Install Dependencies**
    ```sh
    go mod tidy
    ```

3. **Run the Application**
    ```sh
    make server
    ```

4. **Build the Application**
    ```sh
    go build -o go-crud cmd/main.go
    ```

## API Endpoints

### Domain: localhost:8080
- **GET /v1/AllBooks**: Retrieve all items.
- **GET /v1/book/{id}**: Retrieve a specific item by ID.
- **POST /v1/book**: Create a new item.
- **PUT /v1/book/{id}**: Update an existing item by ID.
- **DELETE /v1/book/{id}**: Delete an item by ID.

## Files

- **Makefile**: Contains build and run scripts.
- **go.mod**: Defines the module and its dependencies.
- **go.sum**: Checksums for module dependencies.
