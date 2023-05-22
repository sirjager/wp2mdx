package service

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func (p *Post) BuildMarkdown(w *WordressSite, directory string) (err error) {
	directory = directory + "/markdown"
	// Build the markdown content
	var builder strings.Builder

	// Write Markdown to a file at the finish of this function
	defer func() {
		content := builder.String()
		filename := fmt.Sprintf("%s.mdx", p.Slug)

		if err := os.MkdirAll(directory, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		filepath := filepath.Join(directory, filename)

		file, err := os.Create(filepath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		_, err = io.WriteString(file, content)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Front Matter Start
	builder.WriteString("---\n")

	// Dates and status
	builder.WriteString(fmt.Sprintf("type: '%s'\n", p.Type))
	if p.Status == "publish" {
		p.Status = "published"
	}
	builder.WriteString(fmt.Sprintf("status: '%s'\n", p.Status))
	builder.WriteString(fmt.Sprintf("published: %s\n", p.Date))
	builder.WriteString(fmt.Sprintf("modifided: %s\n", p.Modified))

	// Slug, Title And Description
	builder.WriteString(fmt.Sprintf("slug: '%s'\n", p.Slug))
	builder.WriteString(fmt.Sprintf("title: '%s'\n", p.Title.Rendered))

	description, err := hTMLToPlainText(p.Excerpt.Rendered)
	if err != nil {
		return err
	}
	description = strings.ReplaceAll(description, " […]\n", "")
	builder.WriteString(fmt.Sprintf("description: '%s'\n", description))

	// Author
	builder.WriteString("author: \n")
	builder.WriteString(fmt.Sprintf(" name: '%s'\n", p.Embed.Author[0].Name))
	builder.WriteString(fmt.Sprintf(" slug: '%s'\n", p.Embed.Author[0].Slug))
	builder.WriteString(fmt.Sprintf(" about: '%s'\n", p.Embed.Author[0].Description))
	builder.WriteString(fmt.Sprintf(" image: '%s'\n", p.Embed.Author[0].AvatarURLs.Size96))

	// Tags
	tags := w.GetTags(p)
	if len(tags) > 0 {
		builder.WriteString("tags: \n")
		for _, tag := range tags {
			builder.WriteString(fmt.Sprintf(" - '%s' \n", tag.Slug))
		}
	} else {
		builder.WriteString("tags: [] \n")
	}

	// Categories
	categories := w.GetCategories(p)
	if len(categories) > 0 {
		builder.WriteString("categories: \n")
		for _, cat := range categories {
			builder.WriteString(fmt.Sprintf("  - '%s' \n", cat.Slug))
		}
	} else {
		builder.WriteString("categories: [] \n")
	}

	// Images
	builder.WriteString("image: \n")
	if len(p.Embed.FeaturedMedia) > 0 {
		// image title
		builder.WriteString(fmt.Sprintf("  title: '%s'\n", p.Embed.FeaturedMedia[0].Title.Rendered))
		// image caption
		caption, err := hTMLToPlainText(p.Embed.FeaturedMedia[0].Caption.Rendered)
		if err != nil {
			return err
		}
		builder.WriteString(fmt.Sprintf("  caption: '%s'\n", caption))
		altText := p.Embed.FeaturedMedia[0].AltText
		if altText == "" {
			altText = p.Title.Rendered
		}
		builder.WriteString(fmt.Sprintf("  alt: '%s'\n", altText))
	} else {
		builder.WriteString(fmt.Sprintf("  title: '%s'\n", p.Title.Rendered))
		builder.WriteString(fmt.Sprintf("  caption: '%s'\n", p.Title.Rendered))
		builder.WriteString(fmt.Sprintf("  alt: '%s'\n", description))

	}
	builder.WriteString(fmt.Sprintf("  template: 'default/%s'\n", p.Slug))
	builder.WriteString(fmt.Sprintf("  preview: '/articles/%s/preview.avif'\n", p.Slug))
	builder.WriteString(fmt.Sprintf("  featured: '/articles/%s/featured.avif'\n", p.Slug))

	// Front Matter End
	builder.WriteString("---\n")

	builder.WriteString("\n")

	// content, err := HtmlToMarkDown(p.Content.Rendered)
	// if err != nil {
	// 	w.logger.Error().Err(err).Msgf("markdown: failed", p.Slug)
	// 	return
	// }

	// content = strings.ReplaceAll(content, "\n\n\n", "\n")

	builder.WriteString(p.Title.Rendered)

	return
}
