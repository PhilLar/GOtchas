package main

import (
	"github.com/PhilLar/GOtchas/deps-injection_uber-fx/logger"
)

// unmodularized

//func main() {
//	logger := log.New(os.Stdout, "[ACME] ", 0)
//
//	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		logger.Println("Handler called")
//		io.WriteString(w, "Hello, LOGGER")
//	})
//
//	http.Handle("/", handler)
//
//	http.ListenAndServe(":8080", nil)
//}

// modularized

func main() {
	logger := logger.NewLogger()
}










// 1) use closure to add count to logger