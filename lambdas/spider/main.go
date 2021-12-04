// Copyright 2021 - present prtech.news. All rights reserved.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/prtech-news/common"
	"log"
	"sort"
)

const htmlOutputBucket = "prtech.news"
const configBucket ="prtech.news.config"
const configBucketKey = "config.json"
const region = "us-east-1"
const key = "index.html"

type Event struct {}

type Configuration struct {
	Urls []string `json:"urls"`
	Phrases []string `json:"phrases"`
}

var s3session *session.Session

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		log.Println("Unable to create aws session")
	}
	s3session = sess
}

func HandleRequest(ctx context.Context, event *Event) (string, error) {
	// Download configuration object
	downloader := s3manager.NewDownloader(s3session)
	buf := aws.NewWriteAtBuffer([]byte{})
	_, err := downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(configBucket),
		Key:    aws.String(configBucketKey),
	})
	if err != nil {
		log.Printf(
			"Unable to download config from object key: %s, bucket: %s\n",
			configBucketKey,
			urlSourceBucket,
		)
	}
	var config *Configuration
	err = json.Unmarshal(buf.Bytes(), &config)
	if err != nil {
		log.Println("Error unmarshaling config")
		return nil, errors.New("Error unmarshaling config")
	}

	// Parse RSS Feeds
	feedParser := &common.RSSFeedParser{nil}
	articles := common.ParseRSSFeedsAsync(feedParser, config.Urls)
	// Filter articles
	filtered := common.FilterByTitle(articles, config.Phrases)
	// Sort DESC
	sortByPubDateDesc(filtered)
	// Server Side HTML templating
	htmlBytes := common.CreateHtmlFromArticles(filtered)

	log.Printf("Uploading %s into bucket %s\n", key, htmlOutputBucket)
	uploader := s3manager.NewUploader(s3session)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(htmlOutputBucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(htmlBytes),
		ContentType: aws.String("text/html"),
	})

	return "Finished", nil
}

func sortByPubDateDesc(articles []*Article) {
	sort.Slice(articles, func(i, j int) bool {
    return articles[i].PubDateParsed.After(articles[j].PubDateParsed)
	})
}

func main() {
	//lambda.Start(HandleRequest)
	HandleRequest(nil, &Event{})
}