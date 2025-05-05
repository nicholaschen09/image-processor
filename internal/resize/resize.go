package resize

import (
    "image"
    "image/draw"
)

func Resize(img image.Image, width int, height int) image.Image {
    // Create a new image with the specified dimensions
    newImg := image.NewRGBA(image.Rect(0, 0, width, height))
    
    // Resize the original image to the new dimensions
    draw.BiLinear.Scale(newImg, newImg.Rect, img, img.Bounds(), draw.Over, nil)
    
    return newImg
}