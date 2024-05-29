# Go CRUD Application

## Overview

This project is a simple CRUD (Create, Read, Update, Delete) application built using Go. It demonstrates how to build a RESTful API using Go with basic CRUD functionalities for managing data entries.

## Project Structure

- **cmd/**: Contains the main application entry point.
  - **main.go**: The main file that initializes and starts the application.
  
- **pkg/**: Contains the application's core logic and router setup.
  - **books/controller**: Contains the HTTP handlers for each endpoint.
    - **controller.go**: Defines the functions to handle HTTP requests.
  - **common/models**: Contains the data models.
    - **models.go**: Defines the data structures used in the application.
  - **common/db**: Contains creation of the database and related operations.
  - **books/**: Contains the all the api and database interactions.
    - **add_book.go**: Defines database operations regarding adding a book into the database.
    - **delete_book.go**: Defines database operations regarding removing a book from the database.
    - **get_books.go**: Defines database operations regarding fetching all books from the database.
    - **get_books.go**: Defines database operations regarding fetching th book using id from the database.
    - **update_book.go**: Defines database operations regarding updating the book using id from the database.

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
