# Golang + MySQL


## Description

This is a learning project for building a restful API using Golang and several supporting packages. This is the Categories API, which allows you to manage categories for various purposes in your application.


## Features

- Endpoints for creating, reading, updating, and deleting data.
- Connection to a MySQL database.
- Easy-to-follow directory structure for larger project development.
- Utilizes the [httprouter](https://github.com/julienschmidt/httprouter) package for efficient HTTP routing.


## Installation

1. Ensure that Go is installed on your system. You can visit https://golang.org for more information.

2. Clone this repository to your local directory:
```shell
git clone https://github.com/andromeda951/belajar-golang-restful-api.git
```

3. Navigate to the project directory:
```shell
cd belajar-golang-restful-api
```

4. Install the required dependencies:
```shell
go mod download
```


## Configuration

1. Create an `.env` file containing the application configuration. Example configuration:
```shell
# Port
PORT=8080

# MySQL configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=secret
DB_NAME=database_name

# API key
API_KEY=secret
```

2. Adjust the configuration values according to your environment.


## Usage

1. Run the application using the following command:
```shell
go run main.go
```

2. The application will run on `http://localhost:8080/api`.

3. You can access the API using tools such as cURL or Postman.


## Endpoints

The following endpoints are available in the Categories API:

### Get All Categories

- **Endpoint:** `/categories`
- **Method:** GET
- **Description:** Retrieves all categories from the database.
- **Response:** Returns an array of category objects.

### Get Category by ID

- **Endpoint:** `/categories/{id}`
- **Method:** GET
- **Description:** Retrieves a specific category based on the provided ID.
- **Parameters:**
  - `{id}` (Path Parameter): The unique identifier of the category.
- **Response:** Returns the category object corresponding to the provided ID.

### Create Category

- **Endpoint:** `/categories`
- **Method:** POST
- **Description:** Creates a new category.
- **Request Body:** JSON object containing the category name.
- **Response:** Returns the created category object along with the ID automatically incremented.

### Update Category

- **Endpoint:** `/categories/{id}`
- **Method:** PUT
- **Description:** Updates an existing category based on the provided ID.
- **Parameters:**
  - `{id}` (Path Parameter): The unique identifier of the category to be updated.
- **Request Body:** JSON object containing the updated category name.
- **Response:** Returns the updated category object.

### Delete Category

- **Endpoint:** `/categories/{id}`
- **Method:** DELETE
- **Description:** Deletes a specific category based on the provided ID.
- **Parameters:**
  - `{id}` (Path Parameter): The unique identifier of the category to be deleted.
- **Response:** Returns a success message indicating the deletion of the category.

Please note that the examples provided above are for simple purposes. You should adjust the endpoint details based on the specific requirements and design of your Categories API.


## Directory Structure

```shell
├── app
│   ├── database.go
│   └── router.go
├── controller
│   ├── category_controller_impl.go
│   └── category_controller.go
├── exception
│   ├── error_handler.go
│   └── not_fount_error.go
├── helper
│   ├── error.go
│   ├── json.go
│   ├── model.go
│   └── tx.go
├── middleware
│   └── auth_middleware.go
├── model
│   ├── domain
│   │   └── category.go
│   ├── web
│   │   ├── category_create_request.go
│   │   ├── category_response.go
│   │   ├── category_update_request.go
│   │   └── web_response.go
├── repository
│   ├── category_repository_impl.go
│   └── category_repository.go
├── service
│   ├── category_service_impl.go
│   └── category_service.go
├── test
│   └── category_test.go
├── .env.example
├── .gitignore
├── apispec.json
├── belajar_golang_restful_api.sql
├── go.mod
├── main.go
└── README.md
```
