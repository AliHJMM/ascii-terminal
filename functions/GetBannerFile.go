package functions

import (
	"strings"
)

func GetBannerFile(bannerName string) string {
	bannerFiles := map[string]string{
		"shadow":    "shadow.txt",
		"standard":  "standard.txt",
		"thinkertoy": "thinkertoy.txt",
	}

	bannerName = strings.ToLower(bannerName)

	if strings.HasSuffix(bannerName, ".txt") {
		bannerName = strings.TrimRight(bannerName, ".txt")
	}

	if bannerFile, exists := bannerFiles[bannerName]; exists {
		return bannerFile
	}

	return ""
}
