package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/seyamibrahim/students-api/internal/config"
	"github.com/seyamibrahim/students-api/internal/http/handlers/student"
	"github.com/seyamibrahim/students-api/internal/storage/sqlite"
)

func main() {
	// load config

	cfg := config.MustLoad()



	// database setup

	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))


	// setup router

	router := http.NewServeMux()
	router.HandleFunc("POST /api/students",student.New(storage))
	router.HandleFunc("GET /api/students/{id}",student.GetById(storage))
	router.HandleFunc("GET /api/students",student.GetStudents(storage))
	//setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	fmt.Printf("Server is running on %s\n", cfg.Addr)
	// gracefully shutdown

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start Server")
		}
	}()

	<-done


	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil{
		slog.Error("failed to shutdown server")
	}


	slog.Info("Server shutdown successfully")

}
