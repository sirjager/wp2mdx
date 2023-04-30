package service

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type WordressData map[string]interface{}

func (w *WordressSite) DownloadData(directory string) (err error) {
	if w.perPage < 1 || w.perPage > 100 {
		w.perPage = 50
	}
	allPosts := []Post{}
	allMedia := []Media{}
	allUsers := []User{}
	allPages := []Page{}
	allTags := []Tag{}
	allCategories := []Category{}

	var mu sync.Mutex
	var wg sync.WaitGroup

	// delay := time.Millisecond * 700

	// Create a custom HTTP client with a custom Transport
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	for _, collection := range w.Collections {
		wg.Add(1)
		go func(collection string) {
			defer wg.Done()
			var totalPages int = 1 // after first request this will update from response header of 1st request
			for i := 1; i <= totalPages; i++ {
				// time.Sleep(delay)
				pageURL := fmt.Sprintf("%s/wp-json/wp/v2/%s?per_page=%d&page=%d&_embed", w.siteUrl, collection, w.perPage, i)

				req, err := http.NewRequest("GET", pageURL, nil)
				if err != nil {
					w.logger.Error().Err(err).Msgf("failed to create request for URL: %s", pageURL)
					return
				}

				// Set the User-Agent header
				req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
				req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")

				res, err := client.Do(req)

				if err != nil {
					w.logger.Error().Err(err).Msgf("download: failed fetching: %s", pageURL)
					return
				}
				defer res.Body.Close()

				if i == 1 {
					_totalPages := res.Header.Get("x-wp-totalpages")
					totalPages, _ = strconv.Atoi(_totalPages)
					// _totalItems := res.Header.Get("x-wp-total")
				}

				if res.StatusCode != http.StatusOK {
					w.logger.Error().Msgf("download: failed fetching: %s (status: %s)", pageURL, res.Status)
					return
				}

				switch collection {
				case "posts":
					var itemList []Post
					err = json.NewDecoder(res.Body).Decode(&itemList)
					if err != nil {
						w.logger.Error().Err(err).Msgf("download: failed decoding posts: %s", pageURL)
						return
					}
					mu.Lock()
					allPosts = append(allPosts, itemList...)
					mu.Unlock()
				case "pages":
					var itemList []Page
					err = json.NewDecoder(res.Body).Decode(&itemList)
					if err != nil {
						w.logger.Error().Err(err).Msgf("download: failed decoding pages:  %s", pageURL)
						return
					}
					mu.Lock()
					allPages = append(allPages, itemList...)
					mu.Unlock()

				case "users":
					var itemList []User
					err = json.NewDecoder(res.Body).Decode(&itemList)
					if err != nil {
						w.logger.Error().Err(err).Msgf("download: failed decoding users: %s", pageURL)
						return
					}
					mu.Lock()
					allUsers = append(allUsers, itemList...)
					mu.Unlock()
				case "tags":
					var itemList []Tag
					err = json.NewDecoder(res.Body).Decode(&itemList)
					if err != nil {
						w.logger.Error().Err(err).Msgf("download: failed decoding tags: %s", pageURL)
						return
					}
					mu.Lock()
					allTags = append(allTags, itemList...)
					mu.Unlock()
				case "categories":
					var itemList []Category
					err = json.NewDecoder(res.Body).Decode(&itemList)
					if err != nil {
						w.logger.Error().Err(err).Msgf("download: failed decoding categories:  %s", pageURL)
						return
					}
					mu.Lock()
					allCategories = append(allCategories, itemList...)
					mu.Unlock()
				case "media":
					var itemList []Media
					err = json.NewDecoder(res.Body).Decode(&itemList)
					if err != nil {
						w.logger.Error().Err(err).Msgf("download: failed decoding media:  %s", pageURL)
						return
					}
					mu.Lock()
					allMedia = append(allMedia, itemList...)
					mu.Unlock()
				}
			}
		}(collection)

	}

	wg.Wait()

	// Manipulate Data If Needed
	allPosts = modifyPosts(allPosts)

	var data WordressData = WordressData{
		"posts":      allPosts,
		"pages":      allPages,
		"users":      allUsers,
		"media":      allMedia,
		"tags":       allTags,
		"categories": allCategories,
	}

	var wg2 sync.WaitGroup

	for _, _collection := range w.Collections {
		wg2.Add(1)
		go func(collection string) {
			defer wg2.Done()
			filename := collection + ".json"
			items := data[collection]
			jsonData, err := json.MarshalIndent(items, "", " ")
			if err != nil {
				w.logger.Error().Err(err).Msg("failed to marshal posts to JSON")
				return
			}
			writeDataToFile(directory, filename, jsonData)
		}(_collection)
	}

	wg2.Wait()

	w.AllPosts = allPosts
	w.AllMedia = allMedia
	w.AllPages = allPages
	w.AllUsers = allUsers
	w.AllTags = allTags
	w.AllCategories = allCategories
	w.HasData = true

	return
}

func modifyPosts(posts []Post) (modifided []Post) {
	for _, p := range posts {
		if p.Type == "post" {
			p.Type = "article"
		}
		if p.Status == "publish" {
			p.Status = "published"
		}
		modifided = append(modifided, p)
	}
	return
}
