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

### Hourly forecast interpretation

The service calls the National Weather Service hourly forecast API and uses the first forecast period in the response as the current weather snapshot. It maps that period to the response payload as follows:

- `forecast` uses the first period's `shortForecast` value.
- `tempCharacteristic` is derived from the first period's `temperature` value in Fahrenheit:
  - `Hot` when the temperature is greater than 75°F
  - `Cold` when the temperature is less than 50°F
  - `Moderate` for all other temperatures (50°F to 75°F inclusive)

Example response:

```json
{
  "city": "Dallas",
  "state": "TX",
  "forecast": "Partly Cloudy",
  "tempCharacteristic": "Moderate"
}
```

Example:

### PowerShell

```powershell
Invoke-RestMethod -Uri "http://localhost:8080/currentweather/33.0811/-97.5631"
```

### Bash

```bash
curl "http://localhost:8080/currentweather/33.0811/-97.5631"
```

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
