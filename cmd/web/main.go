package main

import (
    "flag"
    "log/slog"
    "net/http"
    "os"
)

type application struct {
    logger *slog.Logger
}

func main() {
    addr := flag.String("addr", ":4000", "HTTP network adress")
    flag.Parse()

    logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelDebug,
        AddSource: true,
    }))

    app := &application{
        logger: logger,
    }

    logger.Info("starting server", "addr", *addr)
    
    err := http.ListenAndServe(*addr, app.routes())
    logger.Error(err.Error())
    os.Exit(1)
    
}