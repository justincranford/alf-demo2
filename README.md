# Texas Recipes API

A RESTful API for managing Texas recipes, implemented in Go with OpenAPI 3.0 specification.

## Project Structure

- `openapi.yaml` - API specification in OpenAPI 3.0 format
- `server/` - Go implementation of the API
  - `main.go` - Server implementation with handlers
  - `main_test.go` - Integration tests
  - `swagger/` - Swagger UI for API documentation
- `.github/copilot-instructions.md` - Guidelines for generating Texas recipes

## Development Steps

Below is a chronological list of the steps taken to develop this project:

1. Initialize Git repository
   ```
   git init
   ```

2. Create an OpenAPI 3.0 specification for a book store
   ```
   create an openapi 3.0 specification for a book store
   ```

3. Pivot to a beer store API
   ```
   no, make it a beer store
   ```

4. Pivot again to a food recipes API
   ```
   no, food recipes
   ```

5. Commit the OpenAPI specification
   ```
   commit
   ```

6. Initially planned to create a Python server
   ```
   create a python server and tests for this spec
   ```

7. Changed to Go implementation with integration tests
   ```
   implement a server for this specification
   create integration tests for this server
   use go
   ```

8. Run tests to verify functionality
   ```
   run the tests, see if they work
   ```

9. Confirm Go implementation
   ```
   no, i said use golang
   ```

10. Continue implementing in Go
    ```
    keep going with golang
    ```

11. Commit the Go server implementation
    ```
    commit
    ```

12. Start the server and access Swagger UI
    ```
    start the server, and open the swagger ui
    ```

13. Fix any bugs in the server and Swagger UI setup
    ```
    start the server, and open the swagger ui. if there is a bug, fix it.
    ```

14. Commit changes
    ```
    commi
    ```

15. Add Texas-specific recipe instructions
    ```
    add an instructions file, i only want recipes from texas
    ```

16. Create GitHub Copilot instructions
    ```
    i want it to be a copilot-instructions.md file under .github
    ```

17. Commit the instructions
    ```
    commit
    ```

## API Endpoints

- `GET /recipes` - List all Texas recipes
- `POST /recipes` - Add a new Texas recipe
- `GET /recipes/{id}` - Get details of a specific recipe
- `PUT /recipes/{id}` - Update an existing recipe
- `DELETE /recipes/{id}` - Remove a recipe

## Running the Server

```
cd server
go run main.go
```

The server will start on http://localhost:3000, with Swagger UI available at http://localhost:3000/swagger/

## Testing

```
cd server
go test -v
```

## Example Recipe

```json
{
  "id": "tx001",
  "name": "Texas-Style Beef Brisket",
  "ingredients": [
    "1 whole beef brisket (10-12 pounds)",
    "1/4 cup kosher salt",
    "1/4 cup black pepper",
    "2 tablespoons garlic powder",
    "2 tablespoons onion powder",
    "1 tablespoon cayenne pepper"
  ],
  "instructions": "1. Mix all spices together to create a rub. 2. Trim excess fat from brisket, leaving about 1/4 inch fat cap. 3. Apply rub generously to all sides of brisket. 4. Let sit for 1 hour at room temperature. 5. Prepare smoker to 225°F using oak or hickory wood. 6. Smoke brisket fat side up for 6 hours. 7. Wrap in butcher paper and continue cooking until internal temperature reaches 203°F (about 6 more hours). 8. Rest for 1-2 hours before slicing against the grain.",
  "cuisine": "Texas BBQ",
  "prepTime": 60,
  "cookTime": 720
}
```
