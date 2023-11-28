# gomux
A series of exercises to demonstrate how to build a http service using http.ServeMux

- Fork the repo
- Each exercise is in a branch named with a prefix of `ex`. Use `git branch` to see the exercises. 
- After each exercise, merge the branch to `main`
- Merge `main` to the next exercise
- Start the next exercise

## Exercise 1 - create a server with a health endpoint 

Create a http server. The server handles `/health` and returns `{"status": "ok"}` upon success. 

### Reading

- (`NewServeMux`)[https://pkg.go.dev/net/http#NewServeMux]
- (`ServeMux`)[https://pkg.go.dev/net/http#ServeMux]
- (`HandleFunc`)[https://pkg.go.dev/net/http#ServeMux.HandleFunc] 

## Exercise 2 - inject the server starting time to the health response

Continue to modify your ex01 code, add the start time in the health response, e.g. 

``` json
{
    "status": "ok",
    "startedAt": "time when the server started"
}
```

Since you need to record the time at the beginning of `main`, you need to inject
the time when a health request is served. You will need to define a `struct`
that implements the `Handler` interface. The `struct` is known as a handler or a
controller. Initialise the handler with the server start time so that it is
available when serving each request.

### Reading

- (Handler)[https://pkg.go.dev/net/http#Handler]
- (Handle)[https://pkg.go.dev/net/http#ServeMux.Handle]
