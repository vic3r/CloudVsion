package utils

import (
	pb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

func Has_id(labels []*pb.EntityAnnotation) bool {
	for _, label := range labels {
		if label.Description == "identity document" {
			return true
		}
	}
	return false
}
