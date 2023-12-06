# gomux

A series of exercises to demonstrate how to build a http service using http.ServeMux

- Fork the repo
- Do each exercise is in a branch
- After each exercise, merge the branch to `main`
- Start the next exercise

## Exercise 1 - create a server with a health endpoint 

Create an http server. The server handles `/health` and returns `{"status": "ok"}` upon success. 

### Reading

- [`NewServeMux`](https://pkg.go.dev/net/http#NewServeMux)
- [`ServeMux`](https://pkg.go.dev/net/http#ServeMux)
- [`HandleFunc`](https://pkg.go.dev/net/http#ServeMux.HandleFunc) 

## Exercise 2 - inject the server starting time to the health response

Continue to modify your ex01 code, add the start time in the health response, e.g. 

``` json
{
    "status": "ok",
    "startedAt": "time when the server started"
}
```

Since you need to record the time at the beginning of `main`, you need to inject
the time into the response when each health request is served. You will need to
define a `struct` that implements the
[`Handler`](https://pkg.go.dev/net/http#Handler) interface. The `struct` is
known as a handler or a controller. Initialise the handler with the server start
time so that it is available when serving each request.

### Reading

- [Handler](https://pkg.go.dev/net/http#Handler)
- [Handle](https://pkg.go.dev/net/http#ServeMux.Handle)

## Exercise 3 - HTTP methods and error response

Only accept HTTP get method on `/health` and return 401 for other http methods.

### Reading

- [HTTPResponseWriter](https://pkg.go.dev/net/http#ResponseWriter)


## Exercise 4 - Middlewares

A [Handler](https://pkg.go.dev/net/http#Handler) is an object that implements
the [Handler](https://pkg.go.dev/net/http#Handler) interface by implementing the
`ServeHTTP` method. What `mux.Handle(path, handler)` does is to call
`handler.ServeHTTP` when a request comes at `path`

Middlewares allow you to chain the handlers for a request. A middleware is
useful to apply common logic for all requests.

A middleware is a function that takes in a handler and returns a handler.

Using [HandlerFunc](https://pkg.go.dev/net/http#HandlerFunc), you can turn a
`HandleFunc` to [Handler](https://pkg.go.dev/net/http#Handler). Here is a
middlware that logs a request.

``` go

// a middleware that logs every request
func logRequestHandler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        start := time.Now()

        log.Printf("%s %s %s", req.Method, req.RequestURI, start)

        next.ServeHTTP(w, req)
    })
}

// Use the middleware before processing a request 
mux.Handle("/health", logRequestHandler(healthHandler))

```

Integrate the log request handler in your code.

### Reading

- [HandlerFunc](https://pkg.go.dev/net/http#HandlerFunc)

