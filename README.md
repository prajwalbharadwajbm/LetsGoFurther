# Greenlight API

This project is a follow-along implementation from the "Let's Go Further" book, focusing on building a production-ready REST API using Go.

## Project Overview

Greenlight is a REST API for a movie database application, built using modern Go practices and following production-ready patterns.

## Current Implementation Status

- Basic API server setup with configuration management
- HTTP router using `httprouter`
- Structured error handling
- JSON response handling
- Health check endpoint
- Basic movie endpoints (structure in place)

## Key Learnings

1. **Project Structure**
   - Clean separation of concerns using the `cmd/api` pattern
   - Modular code organization for better maintainability

2. **Configuration Management**
   - Using `flag` package for command-line configuration
   - Environment-specific settings (development/staging/production)

3. **HTTP Server Setup**
   - Custom server configuration with timeouts
   - Proper error handling
   - Structured logging

4. **API Response Handling**
   - Standardized JSON response format
   - Proper content type headers
   - UTF-8 encoding compliance

5. **Routing**
   - Using `httprouter` for efficient request routing
   - Clean route definitions
   - URL parameter handling

## Running the Project

1. Clone the repository
2. Navigate to the `greenlight` directory
3. Run the server:

```bash
go run ./cmd/api
```
The server will start on port 4000 by default.

## API Documentation

The API documentation is available through Swagger UI when running the server. Visit `/swagger/index.html` to view the interactive API documentation.

To regenerate the API documentation, run:

```bash
swag init -g ./cmd/api/docs.go
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.

