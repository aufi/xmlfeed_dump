package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	gofeed "github.com/mmcdole/gofeed"
)

func prepareDestination(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
}

func savePage(url string, path string) {
	resp, err := http.Get(url)

	if err != nil {
		log.Panic(err)
	}

	f, err := os.Create(path)
	defer f.Close()

	if err != nil {
		log.Panic(err)
	}

	err = resp.Write(f)

	if err != nil {
		log.Panic(err)
	}
}

func main() {
	feedURLPtr := flag.String("url", "", "RSS or Atom feed URL")
	destDirectoryPtr := flag.String("destination", "", "Local directory path where items should")
	flag.Parse()

	if *feedURLPtr == "" {
		log.Fatal("Provide -url parameter to identify feed to download, to show more options use $ xmlfeed_dump -h")
	}

	if *destDirectoryPtr == "" {
		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		*destDirectoryPtr = path
	}

	feedURL, err := url.Parse(*feedURLPtr)
	if err != nil {
		log.Fatal("Use Feed URL as the firts argument")
		panic(err)
	}

	log.Printf("Starting RSS/Atom feed website dumper for %s", feedURL.String())

	log.Println("Preparing destination directory..")
	destDirectory := *destDirectoryPtr + "/" + feedURL.Hostname()
	prepareDestination(destDirectory)

	log.Println("Fetching the feed..")
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedURL.String())
	if err != nil {
		log.Println("Invalid Feed")
		panic(err)
	}

	log.Println("Have valid feed with name: " + feed.Title)

	log.Println("Fetching items..")
	for i := 0; i < len(feed.Items); i++ {
		item := feed.Items[i]
		log.Println("Storing Item from URL:" + item.Link)
		targetPath := destDirectory + "/" + strings.ReplaceAll(item.Link, "/", "-")
		savePage(item.Link, targetPath)
		log.Println("Stored to: " + targetPath)
	}

	log.Println("Dump completed.")
}
