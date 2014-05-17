package captcha

import (
	"image"
	"image/color"
	"image/png"
	"image/jpeg"
	"image/gif"
	"os"
)

type captcha struct{
	width int
	height int
	imgType string
	name string
}

func NewCaptcha () *captcha {
	return &captcha{imgType:"png"}
}

/*********************************
*初始化
*params: 
*return: void
********************************/
func (i *captcha) init(width, height int) {
	i.width = width
	i.height = height
	if fd, err := os.Create(fmt.Sprintf("%03d.png", imgcounter)); err !=nil {
		panic("文件创建错误：",err)
	}else{
		i.fd = fd
	}
}

/*********************************
*画水平线
*params: 
*return: void
********************************/
func (i *captcha) DrawHorizLine(color color.Color, fromX, toX, y int) {
	// 遍历画每个点
	for x := fromX; x <= toX; x++ {
	    i.Set(x, y, color)
	}
}

/*********************************
*产生图片
*params: 
*return: void
********************************/
func (i *captcha) GenarateImage(color color.Color, fromX, toX, y int) {
	img := image.NewNRGBA(image.Rect(0, 0, i.width, i.height))
	for y := 0; y < dy; y++ {
        for x := 0; x < dx; x++ {
            if x%8 == 0 {
                // 设置某个点的颜色，依次是 RGBA
                img.Set(x, y, color.RGBA{uint8(x % 256), uint8(y % 256), 0, 255})
            }
        }
    }
    return img
}

func (ih *ImageHandler) rotate(im image.Image) image.Image {
	if ih.Rotate == 0 {
		return im
	}
	var rotated *image.NRGBA
	// trigonometric (i.e counter clock-wise)
	switch ih.Rotate {
		case 90:
			newH, newW := im.Bounds().Dx(), im.Bounds().Dy()
			rotated = image.NewNRGBA(image.Rect(0, 0, newW, newH))
			for y := 0; y < newH; y++ {
				for x := 0; x < newW; x++ {
					rotated.Set(x, y, im.At(newH-1-y, x))
				}
			}
		case -90:
			newH, newW := im.Bounds().Dx(), im.Bounds().Dy()
			rotated = image.NewNRGBA(image.Rect(0, 0, newW, newH))
			for y := 0; y < newH; y++ {
				for x := 0; x < newW; x++ {
					rotated.Set(x, y, im.At(y, newW-1-x))
				}
			}
		case 180, -180:
			newW, newH := im.Bounds().Dx(), im.Bounds().Dy()
			rotated = image.NewNRGBA(image.Rect(0, 0, newW, newH))
			for y := 0; y < newH; y++ {
				for x := 0; x < newW; x++ {
					rotated.Set(x, y, im.At(newW-1-x, newH-1-y))
				}
			}
	}
	return rotated
}