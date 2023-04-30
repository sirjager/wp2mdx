package main

import (
	"log"
	"os"
	"time"
	"wordpress/service"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

var logger zerolog.Logger

var startTime time.Time

func init() {
	startTime = time.Now()
	logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: false})
	logger = logger.With().Timestamp().Logger()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	site := service.NewWordpressSite(os.Getenv("WORDPRESS_SITE"), logger)

	defer func() {
		elapsedTime := time.Since(startTime)
		logger.Info().Str("time taken", elapsedTime.String()).Msg("")
	}()

	// site.DownloadData()

	if err := site.LoadData(); err != nil {
		if err := site.DownloadData(); err != nil {
			logger.Error().Err(err).Msg("")
			return
		}
	}

	if !site.HasData {
		if err := site.DownloadData(); err != nil {
			logger.Error().Err(err).Msg("")
			return
		}
	}

	for _, post := range site.AllPosts {
		post.BuildMarkdown(site, "blog")
	}

}
