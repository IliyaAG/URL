package main

import (
	"fmt"
	"strings"
)

var url string = ""

func main() {
	fmt.Println("enter your url :")
	fmt.Scanln(&url)
	pipedLink := convertToPipedLink(url)
	fmt.Println(pipedLink)
}
func convertToPipedLink(url string) string {
	videoID := extractVideoID(url)
	pipedLink := fmt.Sprintf("https://piped.video/v/%s", videoID)
	return pipedLink
}
func extractVideoID(url string) string {
	startIndex := strings.Index(url, "?v=")
	if startIndex == -1 {
		return ""
	}
	videoID := url[startIndex+3:]
	endIndex := strings.Index(videoID, "&")
	if endIndex != -1 {
		videoID = videoID[:endIndex]
	}
	return videoID
}
