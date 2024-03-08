# Rate-limited API Server

## Overview

This API server provides a robust and efficient way to handle HTTP requests with rate limiting to ensure fair usage and system stability.

## Features

- **Rate Limiting**: Ensures that users adhere to request limits, enhancing the overall system stability and fairness.
- **Dependency Injection**: Simplifies the management of dependencies and promotes cleaner, more modular code.
- **Environment-Based Configuration**: Allows easy configuration of the server through environment variables, facilitating deployment across different environments.
- **Structured Logging**: Utilizes structured logging for easier monitoring and debugging.
- **.env File Support**: Supports loading configurations from `.env` files for development convenience.

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yusupovanton/rate-limited-api
   ```
2. Navigate to the project directory:

3. Load environment variables:
    - Copy the `.env.example` file to `.env` and adjust the variables to match your setup.

4. Build the server (optional):
   ```bash
   go build -o apiserver
   ```

### Running the Server

- Directly with Go:
  ```bash
  go run cmd/service/main.go
  ```
- Using the built executable:
  ```bash
  ./apiserver
  ```

## Configuration

The server configuration is managed through environment variables. The following are key variables:

- `LIMIT_PER_USER`: The number of requests per second allowed for one user. Default is 2.
- `RESET_TIME_FRAME`: The timeframe after which the rate limit resets. Default is 30s
- `PORT`: The port at which the server will be run. Default is :8080

For more details on configuration, refer to the `.env.example` file.

## API Reference

### Rate-Limited Request Handler

#### Endpoint: /example

#### Method: `GET`

#### Description:
This handler processes incoming HTTP requests, applying rate limiting based on the client's IP address. Requests that exceed the rate limit are denied to maintain fair usage and system stability.

#### Request:
- **Parameters:**
    - None. The handler operates based on the requester's IP address.

#### Responses:
- **200 OK**
    - **Description:** The request has been accepted and processed successfully. This indicates that the request did not exceed the rate limit.

- **429 Too Many Requests**
    - **Description:** The request has been denied due to exceeding the rate limit. This status is returned to prevent abuse and ensure service availability.

- **405 Method Not Allowed**
    - **Description:** You are using something other than GET

#### Logging:
- **Request Accepted:**
    - Logs an informational message indicating that the request has been accepted. The client's IP address is included in the log entry for auditing and monitoring purposes.

- **Rate Limit Exceeded:**
    - Logs an error message stating "request denied: rate limit exceeded" along with the client's IP address. This provides clear feedback for system administrators and aids in identifying potentially abusive behavior.

#### Rate Limiting Behavior:
- The rate limiting logic is encapsulated by the `rateLimiter` interface, which assesses whether a request from a given IP address should be allowed based on predefined criteria (e.g., number of requests per time window).

---

### Testing

Ensure you have a proper testing environment set up. To run the tests, execute:

```bash
go clean --testcache && go test ./...
```

You should also be able to test the service manually. Run the server as described above (or use your IDE of choice) and 
run this curl `curl --location 'localhost:8080/example'`. If you set everything up correctly, you should be blocked after 
n successful requests and unblocked again after t, where n and t are set properly. Please also note you should use your port
if you specify it otherwise in your config file.

--- 