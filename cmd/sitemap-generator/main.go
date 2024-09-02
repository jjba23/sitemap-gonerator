// Package main is the entry point of the sitemap-generator
package main

import (
	"log"

	"github.com/averageflow/sitemap-generator.git/internal/sitemapgenerator"
)

func main() {
	config := sitemapgenerator.ParseConfigFlags()
	if config == nil {
		return
	}

	sitemapData := sitemapgenerator.CrawlWebsite(config)

	if err := sitemapgenerator.ReplaceSitemapFileWithNewData(sitemapData); err != nil {
		log.Fatalln(err.Error())
		return
	}
}
