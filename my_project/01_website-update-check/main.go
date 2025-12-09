package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"website-checker/config"
	"website-checker/logger"
	"website-checker/monitor"
)

func main() {
	logger.Init("development")
	defer logger.Sync()
	cfg, err := config.Load()
	if err != nil {
        logger.Fatal("Config yüklenemedi:", err)
    }
	logger.Init(cfg.App.Env)


	logger.Info("Uygulama baslatiliyor",
		"env", cfg.App.Env,
		"interval", cfg.Monitor.Interval,
	)

    
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
        sig := <-sigChan
        logger.Info("Signal alındı:", sig)
        cancel()
    }()

	
	m := monitor.New(cfg.Monitor.URL, cfg.Monitor.Interval, cfg.Monitor.Selector)
	m.Start(ctx)
	logger.Info("Uygulama kapatıldı")
}