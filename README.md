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

