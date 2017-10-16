package main

import (
	"log"
	"os"
	//"reflect"

	// Imports the Google Cloud Vision API clientt package.
	vision "cloud.google.com/go/vision/apiv1"
	"golang.org/x/net/context"

	"CloudVsion/utils"
	"CloudVsion/parser"
	"strings"
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

	//Detect if labels has an id object
	if !utils.Has_id(labels){
		log.Fatalf("This is not an identity document")
	}

	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect text: %v", err)
	}

	to_parse := strings.ToTitle(annotations[0].Description)
	parser.Parse_JSON(to_parse)
}
