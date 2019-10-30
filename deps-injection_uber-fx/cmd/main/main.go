package main

import (
	"github.com/PhilLar/GOtchas/deps-injection_uber-fx/handlers"
	"github.com/PhilLar/GOtchas/deps-injection_uber-fx/logger"
	"github.com/PhilLar/GOtchas/deps-injection_uber-fx/server"
	"go.uber.org/fx"
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

//func main() {
//	log := logger.NewLogger()
//	handler := handlers.NewHandler(log)
//
//	http.Handle("/", handler)
//
//	http.ListenAndServe(":8080", nil)
//
//}

// using fx framework
func main() {
	app := fx.New(
		server.Module,
		logger.Module,
		fx.Provide(
			handlers.NewHandler,
			),
		)
	app.Run()
}











// 1) use closure to add count to logger