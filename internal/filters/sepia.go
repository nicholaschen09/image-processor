package filters

import (
    "image"
    "image/color"
    "math"
)

// Sepia is a filter that implements the Filter interface
type Sepia struct{}

// Apply converts the image to sepia tone
func (s Sepia) Apply(img image.Image) image.Image {
    bounds := img.Bounds()
    sepiaImg := image.NewRGBA(bounds)
    
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            oldColor := img.At(x, y)
            r, g, b, a := oldColor.RGBA()
            
            // Convert to 8-bit components
            r8 := float64(r >> 8)
            g8 := float64(g >> 8)
            b8 := float64(b >> 8)
            
            // Apply sepia formula
            newR := uint8(math.Min(0.393*r8+0.769*g8+0.189*b8, 255))
            newG := uint8(math.Min(0.349*r8+0.686*g8+0.168*b8, 255))
            newB := uint8(math.Min(0.272*r8+0.534*g8+0.131*b8, 255))
            
            sepiaImg.Set(x, y, color.RGBA{newR, newG, newB, uint8(a >> 8)})
        }
    }
    
    return sepiaImg
}