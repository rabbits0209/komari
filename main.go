package main

import (
	"log"
	"log/slog"

	"github.com/rabbits0209/komari/cmd"
	"github.com/rabbits0209/komari/utils"
	logutil "github.com/rabbits0209/komari/utils/log"
)

func main() {
	if utils.VersionHash == "unknown" {
		logutil.SetupGlobalLogger(slog.LevelDebug)
	} else {
		logutil.SetupGlobalLogger(slog.LevelInfo)
	}

	log.Printf("Komari Monitor %s (hash: %s)", utils.CurrentVersion, utils.VersionHash)

	cmd.Execute()
}
