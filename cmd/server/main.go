package main

import (
	"fmt"
)

func main() {
	server := &http.Server{
		ReadHeaderTimeout: 30 * time.Second,
		Addr:              ":8080",
		Handler:           mux,
	}

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
