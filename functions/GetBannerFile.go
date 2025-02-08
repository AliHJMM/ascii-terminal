package functions

import (
	"strings"
)

func GetBannerFile(bannerName string) string {
	bannerFiles := map[string]string{
		"shadow":    "Banners/shadow.txt",
		"standard":  "Banners/standard.txt",
		"thinkertoy": "Banners/thinkertoy.txt",
	}

	bannerName = strings.ToLower(bannerName)

	if strings.HasSuffix(bannerName, ".txt") {
		bannerName = strings.TrimSuffix(bannerName, ".txt") // Fix trimming issue
	}

	if bannerFile, exists := bannerFiles[bannerName]; exists {
		return bannerFile
	}

	return ""
}
