## Step 1: Creating a Basic Web Server

### Requirements:
1. The server must include a router (using either the `chi` or `httprouter` package) and a handler named `healthcheck` that returns "ok".
2. It should accept flags for routing port.
3. Logging functionality must be included.

### Good to Have:
- Addressing the issue of hosting the API documentation.

### Solution:
1. Utilize the [swaggo/swag](https://github.com/swaggo/swag) package to generate the Swagger YAML file.
2. Embed the logo and index.html into the codebase.
3. Use Spotlight Web Components to host the API documentation.

## Step 2: Returning HTTP JSON Responses

### Requirements:
1. Create a helper function that writes a JSON response, marshaled from a Go struct, with pretty printing enabled.
2. Implement handlers for the following routes:
    - GET /v1/movies
    - POST /v1/movies
    - GET /v1/movie/:id
    - PUT /v1/movie/:id
    - DELETE /v1/movie/:id
3. Develop a basic error response handler to manage 500, 404, and 405 errors.

### Problem:
Addressing potential issues with response enveloping.

### Solution:
Utilize an enveloping mechanism to wrap the response for better management.

## Step 3: Accepting, Parsing, and Validating HTTP JSON Requests

### Requirements:
1. Handle bad requests, including scenarios such as:
    - Receiving a non-valid JSON request body.
    - Malformed or erroneous JSON.
    - Input body not matching expected types.
    - No input body provided.

### Solutions:
- Leverage the `Decode()` method to catch and triage errors related to JSON parsing.
- Implement a `readJSON()` helper function in the `cmd/api/helpers.go` file to manage JSON input.

## Additional: create smoke tests for all endpoints

## Additional: create a makefile




