package main

import (
	"os"
	"image"
    "image/png"
    "image/draw"
    "image/color"
)

func getRGBA(img image.Image) *image.RGBA {
    b    := img.Bounds()
    rgba := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, b.Min, draw.Src)
	return rgba
}

// get just the R, G, B, or A components
func getMask(img image.Image, R uint8, G uint8, B uint8) *image.RGBA {
    temp := getRGBA(img)
    b    := img.Bounds()
    w    := b.Max.X
    h    := b.Max.Y

    for i := 0; i < w; i++ {
        for j := 0; j < h; j++ {
            col := temp.At(i, j)
            r, g, b, a := col.RGBA()
            rf := uint8(r >> 8)
            gf := uint8(g >> 8)
            bf := uint8(b >> 8)
            af := uint8(a >> 8)
            if R != 0 {
                rf = rf >> R
            }
            if G != 0 {
                gf = gf >> G
            }
            if B != 0 {
                bf = bf >> B
            }
            temp.Set(i, j, color.RGBA{rf, gf, bf, af})
        }
    }
    return temp
}

// Returns the image with all colors inverted
func getNeg(img image.Image) *image.RGBA {
    temp := getRGBA(img)
    b    := img.Bounds()
    w    := b.Max.X
    h    := b.Max.Y

    for i := 0; i < w; i++ {
        for j := 0; j < h; j++ {
            col := temp.At(i, j)
            r, g, b, a := col.RGBA()
            rf := 255 - uint8(r >> 8)
            gf := 255 - uint8(g >> 8)
            bf := 255 - uint8(b >> 8)
            af := uint8(a >> 8)
            temp.Set(i, j, color.RGBA{rf , gf, bf, af})
        }
    }
    return temp
}

// Returns the image with a super-imposed "grid" separating "nxn" pixels
func getGrid(img image.Image, n int) {
}

//
func getReduced(img image.Image, degs uint8) *image.RGBA {
    temp := getRGBA(img)
    b    := img.Bounds()
    w    := b.Max.X
    h    := b.Max.Y

    for i := 0; i < w; i++ {
        for j := 0; j < h; j++ {
            col := temp.At(i, j)
            r, g, b, a := col.RGBA()
            rf := (uint8(r >> 8) / degs) * degs
            gf := (uint8(g >> 8) / degs) * degs
            bf := (uint8(b >> 8) / degs) * degs
            af := (uint8(a >> 8) / degs) * degs
            temp.Set(i, j, color.RGBA{rf, gf, bf, af})
        }
    }
    return temp
}

func getEdgesHorz(img image.Image) {
}

func getEdgesVert(img image.Image) {
}

func main() {
    file, err := os.Open("../../Desktop/test.png")
	if err != nil {
		panic(err)
	}
    defer file.Close()

	file2, err := os.Create("../../Desktop/output.png")
	if err != nil {
		panic(err)
	}
	defer file2.Close()

    img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

    out := getMask(img, 1, 0, 0)

    // out.Set(i, j, color.RGBA{   0, 191, 255, 255 })
    // r, g, b, a := img.At(50, 50).RGBA()
	// r>>8, g>>8, b>>8, a>>8

	err = png.Encode(file2, out)
	if err != nil {
		panic(err)
	}
}
