package service

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"sync"
)

func (w *WordressSite) LoadData(directory string) (err error) {
	var wg sync.WaitGroup
	for _, _col := range w.Collections {
		wg.Add(1)
		go func(collection string) {
			defer wg.Done()
			filePath := filepath.Join(directory+"/", collection+".json")
			// Open the file
			file, err := os.Open(filePath)
			if err != nil {
				w.logger.Error().Err(err).Msgf("failed to open %s", filePath)
				return
			}
			defer file.Close()
			// Read the file contents
			data, err := io.ReadAll(file)
			if err != nil {
				w.logger.Error().Err(err).Msgf("failed to read %s", filePath)
				return
			}

			switch collection {
			case "posts":
				err = json.Unmarshal(data, &w.AllPosts)
				if err != nil {
					w.logger.Error().Err(err).Msgf("failed unmarshalling : %s", collection)
					return
				}
			case "pages":
				err = json.Unmarshal(data, &w.AllPages)
				if err != nil {
					w.logger.Error().Err(err).Msgf("failed unmarshalling : %s", collection)
					return
				}
			case "media":
				err = json.Unmarshal(data, &w.AllMedia)
				if err != nil {
					w.logger.Error().Err(err).Msgf("failed unmarshalling : %s", collection)
					return
				}
			case "users":
				err = json.Unmarshal(data, &w.AllUsers)
				if err != nil {
					w.logger.Error().Err(err).Msgf("failed unmarshalling : %s", collection)
					return
				}
			case "tags":
				err = json.Unmarshal(data, &w.AllTags)
				if err != nil {
					w.logger.Error().Err(err).Msgf("failed unmarshalling : %s", collection)
					return
				}
			case "categories":
				err = json.Unmarshal(data, &w.AllCategories)
				if err != nil {
					w.logger.Error().Err(err).Msgf("failed unmarshalling : %s", collection)
					return
				}
			}

		}(_col)

	}

	wg.Wait()
	w.HasData = true
	return
}

func (w *WordressSite) GetCategories(post *Post) (categories []Category) {
	for _, catID := range post.Categories {
		for _, category := range w.AllCategories {
			if category.ID == catID {
				categories = append(categories, category)
			}
		}
	}
	return
}

func (w *WordressSite) GetTags(post *Post) (tags []Tag) {
	for _, tagID := range post.Tags {
		for _, tag := range w.AllTags {
			if tag.ID == tagID {
				tags = append(tags, tag)
			}
		}
	}
	return
}
