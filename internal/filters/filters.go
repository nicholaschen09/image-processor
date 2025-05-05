package filters

import "image"

// Filter defines an interface for image filters.
type Filter interface {
	Apply(image.Image) image.Image
}

// ApplyFilter applies the specified filter to the given image.
func ApplyFilter(img image.Image, filterType string) image.Image {
	var filter Filter

	switch filterType {
	case "grayscale":
		filter = Grayscale{}
	case "sepia":
		filter = Sepia{}
	default:
		return img // No filter applied
	}

	return filter.Apply(img)
}