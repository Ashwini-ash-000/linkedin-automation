package main

import (
	"github.com/yourusername/linkedin-automation/internal/browser"
	"github.com/yourusername/linkedin-automation/internal/config"
	"github.com/yourusername/linkedin-automation/internal/logger"
	"github.com/yourusername/linkedin-automation/internal/scheduler"
	"github.com/yourusername/linkedin-automation/internal/store"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize logger
	log := logger.New(cfg)

	// Initialize browser session
	session := browser.NewSession(cfg, log)
	defer session.Close()

	// Initialize persistent store (SQLite)
	st, err := store.New("state.db")
	if err != nil {
		log.Error("Failed to initialize state store")
		return
	}

	// Run scheduled automation job
	scheduler.Run(cfg, func() error {
		return runBot(session, st, log)
	})
}
