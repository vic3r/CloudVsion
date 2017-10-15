package main

import (
	"fmt"
	"log"
	"os"
	//"reflect"

	// Imports the Google Cloud Vision API client package.
	vision "cloud.google.com/go/vision/apiv1"
	"golang.org/x/net/context"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("Usage: ./SampleVision imagefile")
	}

	// Sets the name of the image file to annotate.
	filename := os.Args[1]

	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}
	//Detect the entities that are inside the image (labels)
	labels, err := client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect labels: %v", err)
	}

	proceed := false
	for _, label := range labels {
		fmt.Println(label.Description)
		if label.Description == "identity document" {
			proceed = true
		}
	}

	if proceed == false {
		log.Fatalf("This is not an identity document")
	}

	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect text: %v", err)
	}

	if len(annotations) == 0 {
		fmt.Println("No text found.")
	} else {
		for _, annotation := range annotations {

			fmt.Printf("%q\n", annotation.Description)
		}
	}
}
