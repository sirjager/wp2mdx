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

	websiteURL := os.Getenv("WORDPRESS_SITE")

	if !service.ValidateURL(websiteURL) {
		log.Fatalf("Invalid website url : %s", websiteURL)
	}

	folderName, err := service.GenerateFolderName(websiteURL)
	if err != nil {
		log.Fatalf("Error generating folder name : %s", websiteURL)
	}

	directory := "./temp/" + folderName

	site := service.NewWordpressSite(os.Getenv("WORDPRESS_SITE"), logger)

	defer func() {
		elapsedTime := time.Since(startTime)
		logger.Info().Str("time taken", elapsedTime.String()).Msg("")
	}()

	site.DownloadData(directory)

	if err := site.LoadData(directory); err != nil {
		if err := site.DownloadData(directory); err != nil {
			logger.Error().Err(err).Msg("")
			return
		}
	}

	if !site.HasData {
		if err := site.DownloadData(directory); err != nil {
			logger.Error().Err(err).Msg("")
			return
		}
	}

	for _, post := range site.AllPosts {
		post.BuildMarkdown(site, directory)
	}

}
