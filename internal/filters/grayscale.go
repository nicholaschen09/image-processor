package filters

import (
    "image"
    "image/color"
)

// Grayscale is a filter that implements the Filter interface
type Grayscale struct{}

// Apply converts the image to grayscale
func (g Grayscale) Apply(img image.Image) image.Image {
    bounds := img.Bounds()
    grayImg := image.NewRGBA(bounds)
    
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            oldColor := img.At(x, y)
            r, g, b, a := oldColor.RGBA()
            
            // Calculate gray using standard luminance formula
            gray := uint8((0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 256.0)
            
            grayImg.Set(x, y, color.RGBA{gray, gray, gray, uint8(a >> 8)})
        }
    }
    
    return grayImg
}