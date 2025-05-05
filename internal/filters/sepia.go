package filters

import (
    "image"
    "image/color"
)

// Sepia applies a sepia tone to the given image.
func Sepia(img image.Image) *image.RGBA {
    bounds := img.Bounds()
    sepiaImg := image.NewRGBA(bounds)

    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            originalColor := img.At(x, y)
            r, g, b, a := originalColor.RGBA()

            // Convert to sepia
            tr := (0.393 * float64(r>>8)) + (0.769 * float64(g>>8)) + (0.189 * float64(b>>8))
            tg := (0.349 * float64(r>>8)) + (0.686 * float64(g>>8)) + (0.168 * float64(b>>8))
            tb := (0.272 * float64(r>>8)) + (0.534 * float64(g>>8)) + (0.131 * float64(b>>8))

            // Clamp values to 0-255
            if tr > 255 {
                tr = 255
            }
            if tg > 255 {
                tg = 255
            }
            if tb > 255 {
                tb = 255
            }

            sepiaColor := color.RGBA{
                R: uint8(tr),
                G: uint8(tg),
                B: uint8(tb),
                A: uint8(a >> 8),
            }

            sepiaImg.Set(x, y, sepiaColor)
        }
    }

    return sepiaImg
}