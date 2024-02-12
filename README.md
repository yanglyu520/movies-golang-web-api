# Movies web api with golang

This is a movies standard web api developed with golang. This includes steps to demonstrate how it is developed step by step.

## Step 1: Create a basic web server
### Acceptance criteria:
1. It must include a server, a router(use chi or httprouter package), a handler called healthcheck that return text ok
2. It must accept flags for routing port
3. It must have a logger
                                                  
### Good to have:
### Issue:
Find a place to host the apidocs
### Solution:
1. Use [swaggo/swag](https://github.com/swaggo/swag) package to create the swag yaml file
2. Embed the logo, index.html into the code
3. Using spotlight webcomponents to have the apidoc

## Step 2: Return HTTP JSON response
### Acceptance criteria:
1. Create a helper function that writes jsonResponse marshalled from a golang struct, pretty print the JSON response
2. Create a GET /v1/movies handler
3. Create a POST /v1/movies handler
4. Create a GET /v1/movie/:id handler, create a helper function that reads and validate the id
5. Create a PUT /v1/movie/:id handler
6. Create a DELETE /v1/movie/:id handler
7. Create basic errorResponse that deal with 500, 404, 405 error

### Problem:
### Solution:
1. Use envelop to wrap the response

## Step 3: Accept, Parse and Validate HTTP JSON Request
### Acceptance criteria
1. Managing bad requests like below scenarios:
- it receives a non valid JSON request body, like XML or some random bytes, 
- malformed or contains an error, 
- input bodydoes not match the types we are trying to decode into
- no input body

### Solutions:
- All these errors will be caught by the Decode() method, and we just need to triage the errors and replace them with clearer and easy to action error messages to help the client debug exactly what is wrong with their JSON
- create a readJSON() helper function in the cmd/api/helpers.go file




