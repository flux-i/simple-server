# Concurrent API Requests Handling in Go

This project demonstrates how to handle concurrent API requests efficiently in Go, while ensuring optimal resource utilization. It utilizes Goroutines, channels, and mutexes to handle multiple API requests concurrently and avoid redundant calculations.

## Project Structure

- `cmd`: Contains the main application entry point.
- `src`: Contains the source code of the project.
  - `entities`: Defines the entities used in the project, such as API structs and request structs.
  - `handlers`: Contains the handlers for different API endpoints.
  - `managers`: Includes the logic for managing API requests, ensuring concurrency safety.
  - `requests`: Defines common request structures and functions.
  - `routes`: Defines the API routes and paths.
  - `usecases`: Implements the use cases for handling API requests.

## Usage

To run the application, navigate to the `cmd` directory and execute `go run main.go`.

## Dependencies

This project manages its dependencies using Go Modules. Ensure you have Go installed and run `go mod tidy` to install the required dependencies.

## License

MIT

**Free software, Hell Yeah!**
