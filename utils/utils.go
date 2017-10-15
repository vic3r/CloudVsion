package utils

import (
	pb "google.golang.org/genproto/googleapis/cloud/vision/v1"
	"strings"
)

func Has_id(labels []*pb.EntityAnnotation) bool {
	for _, label := range labels {
		if label.Description == "identity document" {
			return true
		}
	}
	return false
}

func Parse_name(to_parse string) []string {
	to_parse = strings.Split(to_parse, "NOMBRE\n")[1]
	to_parse = strings.Split(to_parse, "\nDOMICILIO")[0]
	return strings.Split(to_parse, "\n")
}

