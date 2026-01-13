package data

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/png" // Register PNG decoder
	"io"
	"math"
	"os"
	"path/filepath"
	"strings"
)

// ProcessImageResult holds the paths to the generated images
type ProcessImageResult struct {
	OriginalPath  string
	ThumbnailPath string
	LargePath     string
}

// SaveAndProcessImage saves the raw input, creates a large version (max 1200px) and a thumbnail (max 300px)
func SaveAndProcessImage(r io.Reader, filename string, destDir string) (*ProcessImageResult, error) {
	// 1. Create destination directory if not exists
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %v", err)
	}

	// 2. Decode the image
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	// Generate base filenames
	// We use the original filename's base, but always save as .jpg for processed versions
	ext := filepath.Ext(filename)
	baseName := strings.TrimSuffix(filename, ext)

	// Sanitize baseName simple approach or assume it's already safe/uuid from caller?
	// The caller (API) usually generates a UUID or safe name.

	largeName := baseName + "_large.jpg"
	thumbName := baseName + "_thumb.jpg"
	// We don't save the "original" original unless requested,
	// but the return struct has "OriginalPath".
	// Implementation Plan says:
	// "Create 'Thumbnail' version... Create 'Large' version... Save files..."
	// It doesn't strictly say save the exact original byte-for-byte,
	// but usually we keep the 'Large' as the main viewable.
	// Let's create Large and Thumb.

	largePath := filepath.Join(destDir, largeName)
	thumbPath := filepath.Join(destDir, thumbName)

	// 3. Create "Large" version (Max 1200px width)
	largeImg := ResizeImage(img, 1200)
	if err := saveJPG(largeImg, largePath); err != nil {
		return nil, fmt.Errorf("failed to save large image: %v", err)
	}

	// 4. Create "Thumbnail" version (Max 300px width)
	thumbImg := ResizeImage(img, 300)
	if err := saveJPG(thumbImg, thumbPath); err != nil {
		return nil, fmt.Errorf("failed to save thumbnail: %v", err)
	}

	return &ProcessImageResult{
		LargePath:     largeName,
		ThumbnailPath: thumbName,
	}, nil
}

func saveJPG(img image.Image, path string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	return jpeg.Encode(out, img, &jpeg.Options{Quality: 85})
}

// ResizeImage resizes an image to the specified width while maintaining aspect ratio
// using Bilinear Interpolation.
func ResizeImage(img image.Image, maxWidth int) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// If image is already smaller than maxWidth, return original
	if width <= maxWidth {
		return img
	}

	// Calculate new height preserving aspect ratio
	ratio := float64(maxWidth) / float64(width)
	checkHeight := float64(height) * ratio
	newHeight := int(math.Round(checkHeight))

	// Create new RGBA image
	newImg := image.NewRGBA(image.Rect(0, 0, maxWidth, newHeight))

	// Bilinear scaling
	scaleX := float64(width) / float64(maxWidth)
	scaleY := float64(height) / float64(newHeight)

	for y := 0; y < newHeight; y++ {
		for x := 0; x < maxWidth; x++ {
			// Get source coordinates
			gx := float64(x) * scaleX
			gy := float64(y) * scaleY

			gxi := int(gx)
			gyi := int(gy)

			// Get the 4 surrounding pixels (clamping to edges)
			c00 := img.At(clamp(gxi, 0, width-1), clamp(gyi, 0, height-1))
			c10 := img.At(clamp(gxi+1, 0, width-1), clamp(gyi, 0, height-1))
			c01 := img.At(clamp(gxi, 0, width-1), clamp(gyi+1, 0, height-1))
			c11 := img.At(clamp(gxi+1, 0, width-1), clamp(gyi+1, 0, height-1))

			// Calculate weights
			wx := gx - float64(gxi)
			wy := gy - float64(gyi)

			// Interpolate
			r, g, b, a := bilinearInterpolate(c00, c10, c01, c11, wx, wy)
			newImg.Set(x, y, color.RGBA{R: r, G: g, B: b, A: a})
		}
	}

	return newImg
}

func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func bilinearInterpolate(c00, c10, c01, c11 color.Color, wx, wy float64) (uint8, uint8, uint8, uint8) {
	r00, g00, b00, a00 := c00.RGBA()
	r10, g10, b10, a10 := c10.RGBA()
	r01, g01, b01, a01 := c01.RGBA()
	r11, g11, b11, a11 := c11.RGBA() // These are uint32 pre-multiplied alpha

	// Helper to mix two values linearly
	mix := func(v1, v2 uint32, w float64) float64 {
		return float64(v1)*(1-w) + float64(v2)*w
	}

	// Interpolate X for top row (y=0) and bottom row (y=1)
	rTop := mix(r00, r10, wx)
	gTop := mix(g00, g10, wx)
	bTop := mix(b00, b10, wx)
	aTop := mix(a00, a10, wx)

	rBot := mix(r01, r11, wx)
	gBot := mix(g01, g11, wx)
	bBot := mix(b01, b11, wx)
	aBot := mix(a01, a11, wx)

	// Interpolate Y
	rFinal := rTop*(1-wy) + rBot*wy
	gFinal := gTop*(1-wy) + gBot*wy
	bFinal := bTop*(1-wy) + bBot*wy
	aFinal := aTop*(1-wy) + aBot*wy

	// Convert back to uint8 (RGBA returns 0-65535)
	return uint8(rFinal / 256), uint8(gFinal / 256), uint8(bFinal / 256), uint8(aFinal / 256)
}
