# Zerodha-clone APP

The Zerodha-clone APP is a RESTful APP that allows you to manage stock data. The API is written in Go and uses the Gorilla Mux library for routing.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites

- Go (version 1.16 or higher recommended)
- Gorilla Mux Go library

### Installing

1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. Run the command `go run main.go` to start the server.

## Endpoints

The API contains the following endpoints:

### Welcome

- **URL**: `/`
- **Method**: `GET` or `OPTIONS`
- **Purpose**: This endpoint provides a welcome message when the API server is accessed at the root URL.

### Get Stock

- **URL**: `/api/stock/{id}`
- **Method**: `GET` or `OPTIONS`
- **Purpose**: This endpoint retrieves the stock with the specified `id`.

### Get All Stock

- **URL**: `/api/stock`
- **Method**: `GET` or `OPTIONS`
- **Purpose**: This endpoint retrieves all the stock items available in the database.

### Create Stock

- **URL**: `/api/newStock`
- **Method**: `POST` or `OPTIONS`
- **Purpose**: This endpoint creates a new stock item.

### Update Stock

- **URL**: `/api/stock/{id}`
- **Method**: `PUT` or `OPTIONS`
- **Purpose**: This endpoint updates the stock with the specified `id`.

### Delete Stock

- **URL**: `/api/stock/{id}`
- **Method**: `DELETE` or `OPTIONS`
- **Purpose**: This endpoint deletes the stock with the specified `id`.

## Middleware

The `middleware` package contains the handlers for each endpoint. You will find functions such as `GetStock`, `GetAllStock`, `CreateStock`, `UpdateStock`, and `DeleteStock` here. These functions are responsible for carrying out the logic when the corresponding endpoints are hit.


## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgements

- Gorilla Mux for providing the router used in this project.
