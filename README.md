# Simple Movie API

This is a simple movie API implemented in Go using the Gorilla Mux router. It allows you to perform basic CRUD operations (Create, Read, Update, Delete) on movies.

## Endpoints (use postman to test it)

### Get All Movies

- **Endpoint:** `/movies`
- **Method:** GET
- **Description:** Retrieve a list of all movies.
- **Response:** Returns a JSON array containing all movies.

### Get Movie by ID

- **Endpoint:** `/movies/{id}`
- **Method:** GET
- **Description:** Retrieve a specific movie by its ID.
- **Response:** Returns a JSON object representing the movie if found, or a 404 Not Found response if no movie with the specified ID exists.

### Create Movie

- **Endpoint:** `/movies`
- **Method:** POST
- **Description:** Create a new movie.
- **Request Body:** Expects a JSON object representing the movie to be created. The `id` field will be automatically generated.
- **Response:** Returns a JSON object representing the created movie.

### Delete Movie

- **Endpoint:** `/movies/{id}`
- **Method:** DELETE
- **Description:** Delete a movie by its ID.
- **Response:** If a movie with the specified ID is found and successfully deleted, returns a 200 OK response. If no movie with the specified ID exists, returns a JSON array containing all movies.

## Movie JSON Structure

A movie object has the following properties:

- `id` (string): The unique identifier of the movie.
- `title` (string): The title of the movie.
- `genre` (string): The genre of the movie.
- `director` (object): Information about the director of the movie, including the `name` and `age` properties.

## Usage

1. Make sure you have Go installed on your machine.
2. Clone this repository.
3. Navigate to the project directory.
4. Build and run the project with the following command:
   ```shell
   go run main.go
   ```
