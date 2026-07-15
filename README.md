# Weather Service

A small Go service that retrieves weather information from the National Weather Service API and exposes a local Swagger UI for exploring the API.

## Prerequisites

- Go 1.25+
- Internet access for the outbound weather API calls

## Running the service

From the repository root, run:

```bash
go run .
```

The service will start on:

- http://localhost:8080/swagger-ui/

## API endpoint

The service exposes:

```text
GET /currentweather/{lat}/{long}
```

Example:

### PowerShell

```powershell
$headers = @{ "X-My-User-Agent" = "my-app" }
Invoke-RestMethod -Uri "http://localhost:8080/currentweather/33.0811/-97.5631" -Headers $headers
```

### Bash

```bash
curl -H "X-My-User-Agent: my-app" "http://localhost:8080/currentweather/33.0811/-97.5631"
```

The `X-My-User-Agent` header is required for requests to this service.

## Swagger UI

Swagger UI is served locally from the `resources/` directory.

Open:

```text
http://localhost:8080/swagger-ui/
```

## Project structure

- `main.go` - application entrypoint
- `web/` - HTTP server setup
- `mgr/` - business logic for weather processing
- `repo/` - repositories for calling the weather API
- `models/` - request/response models
- `resources/` - Swagger UI assets and OpenAPI spec

## Notes

- `lat` and `long` should be decimal values.
