package main

import "github.com/Sp92535/web-scraper/scraper"

func main() {
	scraper.InitDB()
	scraper.StartScraper()
	scraper.ExportToJSON("results.json")
}