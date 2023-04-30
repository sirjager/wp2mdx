package service

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func writeDataToFile(filename string, jsonData []byte) (err error) {
	dirPath := "./raw"
	filePath := fmt.Sprintf("%s/%s", dirPath, filename)
	err = os.MkdirAll(dirPath, os.ModePerm) // Create the directory if it doesn't exist
	if err != nil {
		return
	}

	file, err := os.Create(filePath)
	if err != nil {

		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return
	}
	return
}

func hTMLToPlainText(htmlString string) (plainText string, err error) {
	htmlString, err = deleteElementByID(htmlString, "toc_container")
	if err != nil {
		return "", err
	}

	// Parse the HTML string
	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		return "", err
	}

	// Extract plain text from the parsed HTML
	var extractPlainText func(*html.Node, *bytes.Buffer)
	extractPlainText = func(n *html.Node, buf *bytes.Buffer) {
		if n.Type == html.TextNode {
			buf.WriteString(n.Data)
		} else if n.Type == html.ElementNode {
			switch n.Data {
			case "h1":
				buf.WriteString("# ")
			case "h2":
				buf.WriteString("## ")
			case "h3":
				buf.WriteString("### ")
			case "h4":
				buf.WriteString("#### ")
			case "h5":
				buf.WriteString("##### ")
			case "h6":
				buf.WriteString("###### ")
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractPlainText(c, buf)
		}
	}

	var buf bytes.Buffer
	extractPlainText(doc, &buf)
	plainText = strings.TrimSpace(buf.String())

	return plainText, nil
}

func HtmlToMarkDown(rawHtmlContent string) (markdown string, err error) {
	htmlString, err := deleteElementByID(rawHtmlContent, "toc_container")
	if err != nil {
		return "", err
	}

	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		return "", err
	}

	var visitor func(*html.Node)
	visitor = func(node *html.Node) {
		switch node.Type {
		case html.ElementNode:
			switch node.Data {
			case "h1":
				markdown += "# " + getNodeText(node) + "\n"
			case "h2":
				markdown += "## " + getNodeText(node) + "\n"
			case "h3":
				markdown += "### " + getNodeText(node) + "\n"
			case "h4":
				markdown += "#### " + getNodeText(node) + "\n"
			case "h5":
				markdown += "##### " + getNodeText(node) + "\n"
			case "h6":
				markdown += "###### " + getNodeText(node) + "\n"
			case "ul":
				visitor(node.FirstChild)
				markdown += "\n"
			case "li":
				markdown += "- " + getNodeText(node) + "\n"
			case "a":
				href := getAttributeValue(node, "href")
				markdown += fmt.Sprintf("[%s](%s)", getNodeText(node), href)
			case "table":
				visitor(node.FirstChild)
				markdown += "\n"
			case "tr":
				visitor(node.FirstChild)
				markdown += "\n"
			case "th", "td":
				markdown += "| " + getNodeText(node) + " "
				// Add support for other major tags if needed
			}
		case html.TextNode:
			markdown += node.Data
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			visitor(child)
		}
	}

	visitor(doc)
	return markdown, nil
}

func getAttributeValue(node *html.Node, attrName string) string {
	for _, attr := range node.Attr {
		if attr.Key == attrName {
			return attr.Val
		}
	}
	return ""
}

func getNodeText(node *html.Node) string {
	var text strings.Builder
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.TextNode {
			text.WriteString(child.Data)
		} else if child.Type == html.ElementNode && child.Data == "br" {
			text.WriteString("\n")
		} else {
			text.WriteString(getNodeText(child))
		}
	}
	return text.String()
}

func deleteElementByID(htmlStr string, id string) (string, error) {
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return "", err
	}

	var traverse func(*html.Node) bool
	traverse = func(node *html.Node) bool {
		if node.Type == html.ElementNode && node.Data == "div" {
			for _, attr := range node.Attr {
				if attr.Key == "id" && attr.Val == id {
					// Remove the element from its parent node
					parent := node.Parent
					if parent != nil {
						parent.RemoveChild(node)
					}

					// Return true to stop further traversal
					return true
				}
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			if traverse(child) {
				break
			}
		}

		return false
	}

	traverse(doc)

	var output strings.Builder
	err = html.Render(&output, doc)
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
