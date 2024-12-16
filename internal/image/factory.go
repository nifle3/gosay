package image

import (
	"log"

	"github.com/nifle3/gosay/internal/metadata"
)

type ImageType int

const (
	Go ImageType = iota
	Rust
	Zig
)

type ImageData struct {
	Image    string
	Metadata metadata.Image
}

func New(image ImageType) ImageData {
	var result ImageData

	switch image {
	case Go:
		result = ImageData{Image: gopher, Metadata: gopherMetadata}
	case Rust:
		result = ImageData{Image: ferris, Metadata: ferrisMetadata}
	case Zig:
		result = ImageData{Image: ziggy, Metadata: ziggyMetadata}
	default:
		log.Fatal("Undefined ImageType")
	}

	return result
}
