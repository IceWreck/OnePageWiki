package config

import (
	"os"

	"github.com/IceWreck/OnePageWiki/logger"
	"github.com/joho/godotenv"
)

// Default values for env vars
var (
	Credentials             = map[string]string{}
	Port             string = ":8000"
	WikiTitle               = "OnePageWiki"
	MarkdownLocation        = "./wiki.md"
)

func init() {
	err := godotenv.Load("config.env")
	if err != nil {
		logger.Fatal("Error loading config.env file")
	} else {
		Port = ":" + os.Getenv("WIKIPORT")
		WikiTitle = os.Getenv("WIKITITLE")
		MarkdownLocation = os.Getenv("WIKIMARKDOWNFILE")
		user, pass := os.Getenv("WIKIUSER"), os.Getenv("WIKIPASS")
		Credentials[user] = pass
	}
}
