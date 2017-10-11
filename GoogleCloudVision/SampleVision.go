package main

import (
        "fmt"
        "log"
        "os"

        // Imports the Google Cloud Vision API client package.
        vision "cloud.google.com/go/vision/apiv1"
        "golang.org/x/net/context"
)

func main() {
        ctx := context.Background()

        // Creates a client.
        client, err := vision.NewImageAnnotatorClient(ctx)
        if err != nil {
                log.Fatalf("Failed to create client: %v", err)
        }

        // Sets the name of the image file to annotate.
        filename := "../pictures/ife.JPG"

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
        
        fmt.Println("Labels:")
        for _, label := range labels {
                fmt.Println(label.Description)
        }
        annotations, err := client.DetectTexts(ctx, image, nil, 10)
        if err != nil {
                log.Fatalf("Failed to detect text: %v", err)
        }
        faceAnnotations, err := client.DetectFaces(ctx, image, nil, 10)
        if err != nil {
                log.Fatalf("failed to detect faces: %v", err)
        }

        if len(annotations) == 0 {
                fmt.Println("No text found.")
        } else {
                fmt.Println("Text:")
                for _, annotation := range annotations {
                        fmt.Printf( "%q\n", annotation.Description)
                }
        }
         if len(faceAnnotations) == 0 {
                fmt.Println("No text found.")
        } else {
                fmt.Println("Faces:")
               for i, annotation := range faceAnnotations {
                        fmt.Println("Face", i)
                        fmt.Println("Anger:", annotation.AngerLikelihood)
                        fmt.Println("Joy:", annotation.JoyLikelihood)
                        fmt.Println("Surprise:",annotation.SurpriseLikelihood)
                }
        }
}