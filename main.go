package main

import (
	"fmt"
	"log"
	"os"
	//"reflect"

	// Imports the Google Cloud Vision API clientt package.
	vision "cloud.google.com/go/vision/apiv1"
	"golang.org/x/net/context"

	"CloudVsion/utils"
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

	to_parse := annotations[0].Description
	full_name := utils.Parse_name(to_parse)

	var name string
	var last_nameP string
	var last_nameF string

	if len(full_name) == 3 {
		name = full_name[2]
		last_nameP = full_name[0]
		last_nameF = full_name[1]
	} else if len(full_name) == 2 {
		name = full_name[1]
		last_nameP = full_name[0]
		last_nameF = " "
	}

	fmt.Println(name)
	fmt.Println(last_nameP)
	fmt.Println(last_nameF)
}
