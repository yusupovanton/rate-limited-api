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
   ```bash
   cd rate-limited-api
   ```

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

### Testing

Ensure you have a proper testing environment set up. To run the tests, execute:

```bash
go clean --testcache && go test ./...
```

## Configuration

The server configuration is managed through environment variables. The following are key variables:

- `LIMIT_PER_USER`: The number of requests per second allowed for one user. Default is 2.
- `RESET_TIME_FRAME`: The timeframe after which the rate limit resets. Default is 30s

For more details on configuration, refer to the `.env.example` file.

## API Reference

### Rate-Limited Request Handler

#### Endpoint: /example

#### Method: `GET`

#### Description:
This handler processes incoming HTTP requests, applying rate limiting based on the client's IP address. Requests that exceed the rate limit are denied to maintain fair usage and system stability.

#### Request:
- **Headers:**
    - None specified. Customize this section based on your actual requirements.
- **Parameters:**
    - None. The handler operates based on the requester's IP address.

#### Responses:
- **200 OK**
    - **Description:** The request has been accepted and processed successfully. This indicates that the request did not exceed the rate limit.
    - **Body:** Not applicable. Customize this section if your handler returns a response body.

- **429 Too Many Requests**
    - **Description:** The request has been denied due to exceeding the rate limit. This status is returned to prevent abuse and ensure service availability.
    - **Body:**
        - The response body for this example is not explicitly defined. You can customize this section to include a JSON response or other content indicating the rate limit has been exceeded.

#### Logging:
- **Request Accepted:**
    - Logs an informational message indicating that the request has been accepted. The client's IP address is included in the log entry for auditing and monitoring purposes.

- **Rate Limit Exceeded:**
    - Logs an error message stating "request denied: rate limit exceeded" along with the client's IP address. This provides clear feedback for system administrators and aids in identifying potentially abusive behavior.

#### Rate Limiting Behavior:
- The rate limiting logic is encapsulated by the `rateLimiter` interface, which assesses whether a request from a given IP address should be allowed based on predefined criteria (e.g., number of requests per time window).

---
