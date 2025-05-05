package resize

import (
    "image"
    "golang.org/x/image/draw"
)

// Resize resizes an image to the specified dimensions
func Resize(img image.Image, width, height int) image.Image {
    if width == 0 && height == 0 {
        return img
    }
    
    // Calculate new dimensions while preserving aspect ratio
    bounds := img.Bounds()
    origWidth := bounds.Dx()
    origHeight := bounds.Dy()
    
    if width == 0 {
        width = int(float64(height) * float64(origWidth) / float64(origHeight))
    } else if height == 0 {
        height = int(float64(width) * float64(origHeight) / float64(origWidth))
    }
    
    // Create new image with the calculated dimensions
    dst := image.NewRGBA(image.Rect(0, 0, width, height))
    
    // Use high-quality resampling
    draw.BiLinear.Scale(dst, dst.Bounds(), img, bounds, draw.Over, nil)
    
    return dst
}