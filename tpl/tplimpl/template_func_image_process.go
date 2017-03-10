// Copyright 2017 The Hugo Authors. All rights reserved.
//
// Portions Copyright The Go Authors.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tplimpl

import (
	"image"
	"time"
	"strconv"
	"math/rand"
	_nfntResize "github.com/nfnt/resize"
	_bildAdjust "github.com/anthonynsimon/bild/adjust"
	_bildBlend "github.com/anthonynsimon/bild/blend"
	_bildBlur "github.com/anthonynsimon/bild/blur"
	_bildTransform "github.com/anthonynsimon/bild/transform"
	_bildChannel "github.com/anthonynsimon/bild/channel"
	_bildEffect "github.com/anthonynsimon/bild/effect"
	_bildImgIo "github.com/anthonynsimon/bild/imgio"

	// help needed: implement _smartCrop struggled installing smartcrop
	// _smartCrop "github.com/muesli/smartcrop"
)

// File Handler

func ipOpen(path string) image.Image {

	img, err := _bildImgIo.Open(path)
	if err != nil {
		panic(err)
	}

	return img;

}

func ipSave(img image.Image) string {

	// general:
	// help needed: method only works when --renderToDisk is enabled
	// help needed: ipSave should have an optional parameter for filename
	// help needed: caching strategy?

	// filename:
	// help needed: should get the md5 checksum of the file content
	// help needed: should take the file extension from the image instead of fix jpg
	// help needed: should prefix with a custom filepath from config.toml
	filename := strconv.Itoa(rand.Intn(50000)) + strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + ".jpg";

	// help needed: should take public from variables not fix set to public
	publicFilename := "public/" + filename;

	// help needed: should not take _bildImgTo.JPEG as a fix value
	if err := _bildImgIo.Save(publicFilename, img, _bildImgIo.JPEG); err != nil {
		panic(err)
	}

	/*
	Should return:
	{
	  "Permalink" : "https://example.com/[images/]<filename>.jpg",
	  "RelPermalink" : "/[<defined path in config.toml images/>]<filename>.jpg",
	  "Width" : 400,
	  "Height" : 200
	}
	 */

	return filename;

}

// Image Transformation

func ipCrop(img image.Image, x0 int, y0 int, width int, height int) image.Image {
	rect := image.Rect(x0, y0, (x0 + width), (y0 + height))
	return _bildTransform.Crop(img, rect)
}

func ipSmartCrop(img image.Image, width int, height int) image.Image {

	return img;

	// Help needed: unfortunately I receive the following error
	// trying to install the 'smartcrop'-package
	// exec: "gcc": executable file not found in %PATH%

	//settings := smartcrop.CropSettings{
	//	FaceDetection:                    true,
	//	FaceDetectionHaarCascadeFilepath: "./files/aarcascade_frontalface_alt.xml",
	//}
	//analyzer := smartcrop.NewAnalyzerWithCropSettings(settings)
	//topCrop, _ := analyzer.FindBestCrop(img, width, height)
	//
	//return ipCrop(img, topCrop.X, topCrop.Y, topCrop.Width, topCrop.Height);
}

func ipResize(img image.Image, width uint, height uint, interpolation string) image.Image {

	switch interpolation {
	case "Bilinear" :
		return _nfntResize.Resize(width, height, img, _nfntResize.Bilinear);
		break
	case "NearestNeighbor" :
		return _nfntResize.Resize(width, height, img, _nfntResize.NearestNeighbor);
		break
	case "Bicubic" :
		return _nfntResize.Resize(width, height, img, _nfntResize.Bicubic);
		break
	case "MitchellNetravali" :
		return _nfntResize.Resize(width, height, img, _nfntResize.MitchellNetravali);
		break
	case "Lanczos2" :
		return _nfntResize.Resize(width, height, img, _nfntResize.Lanczos2);
		break
	case "Lanczos3" :
		return _nfntResize.Resize(width, height, img, _nfntResize.Lanczos3);
		break
	default :
		return _nfntResize.Resize(width, height, img, _nfntResize.NearestNeighbor);
		break
	}

}

func ipThumbnail(img image.Image, maxWidth uint, maxHeight uint, interpolation string) image.Image {

	switch interpolation {
	case "Bilinear" :
		return _nfntResize.Thumbnail(maxWidth, maxHeight, img, _nfntResize.Bilinear);
		break
	case "NearestNeighbor" :
		return _nfntResize.Thumbnail(maxWidth, maxHeight, img, _nfntResize.NearestNeighbor);
		break
	case "Bicubic" :
		return _nfntResize.Thumbnail(maxWidth, maxHeight, img, _nfntResize.Bicubic);
		break
	case "MitchellNetravali" :
		return _nfntResize.Thumbnail(maxWidth, maxHeight, img, _nfntResize.MitchellNetravali);
		break
	case "Lanczos2" :
		return _nfntResize.Thumbnail(maxWidth, maxHeight, img, _nfntResize.Lanczos2);
		break
	case "Lanczos3" :
		return _nfntResize.Thumbnail(maxWidth, maxHeight, img, _nfntResize.Lanczos3);
		break
	default :
		return _nfntResize.Thumbnail(maxWidth, maxHeight, img, _nfntResize.NearestNeighbor);
		break
	}

}

func ipFlip(img image.Image, axis string) image.Image {
	switch axis {
	case "Horizontal" :
		return _bildTransform.FlipH(img)
	case "Vertical" :
		return _bildTransform.FlipV(img)
	default :
		return img
	}
}

func ipRotate(img image.Image, angle float64) image.Image {
	return _bildTransform.Rotate(img, angle, nil);
}

func ipShear(img image.Image, axis string, angle float64) image.Image {
	switch axis {
	case "Horizontal" :
		return _bildTransform.ShearH(img, angle)
	case "Vertical" :
		return _bildTransform.ShearV(img, angle)
	default :
		return img
	}
}

func ipTranslate(img image.Image, dx int, dy int) image.Image {
	return _bildTransform.Translate(img, dx, dy)
}

// Image Adjustments

func ipBrightness(img image.Image, change float64) image.Image {
	return _bildAdjust.Brightness(img, change);
}

func ipContrast(img image.Image, change float64) image.Image {
	return _bildAdjust.Contrast(img, change);
}

func ipGamma(img image.Image, gamma float64) image.Image {
	return _bildAdjust.Gamma(img, gamma);
}

func ipHue(img image.Image, change int) image.Image {
	return _bildAdjust.Hue(img, change);
}

func ipSaturation(img image.Image, change float64) image.Image {
	return _bildAdjust.Saturation(img, change);
}

// Blend Modes

func ipBlendMode(background image.Image, foreground image.Image, blendmode string) image.Image {
	switch blendmode {
	case "Add" :
		return _bildBlend.Add(background, foreground)
	case "ColorBurn" :
		return _bildBlend.ColorBurn(background, foreground)
	case "ColorDodge" :
		return _bildBlend.ColorDodge(background, foreground)
	case "Darken" :
		return _bildBlend.Darken(background, foreground)
	case "Difference" :
		return _bildBlend.Difference(background, foreground)
	case "Divide" :
		return _bildBlend.Divide(background, foreground)
	case "Exclusion" :
		return _bildBlend.Exclusion(background, foreground)
	case "Lighten" :
		return _bildBlend.Lighten(background, foreground)
	case "LinearBurn" :
		return _bildBlend.LinearBurn(background, foreground)
	case "LinearLight" :
		return _bildBlend.LinearLight(background, foreground)
	case "Multiply" :
		return _bildBlend.Multiply(background, foreground)
	case "Normal" :
		return _bildBlend.Normal(background, foreground)
	case "Overlay" :
		return _bildBlend.Overlay(background, foreground)
	case "Screen" :
		return _bildBlend.Screen(background, foreground)
	case "SoftLight" :
		return _bildBlend.SoftLight(background, foreground)
	case "Subtract" :
		return _bildBlend.Subtract(background, foreground)
	default :
		return background
	}
}

func ipBlendModeOpacity(background image.Image, foreground image.Image, percentage float64) image.Image {
	return _bildBlend.Opacity(background, foreground, percentage)
}

// Blur

func ipBlur(img image.Image, method string, radius float64) image.Image {
	switch method {
	case "Box" :
		return _bildBlur.Box(img, radius)
	case "Gaussian" :
		return _bildBlur.Gaussian(img, radius)
	default :
		return img
	}
}

// Channel Extraction

func ipExtractChannel(img image.Image, chnl string) image.Image {
	switch chnl {
	case "Red" :
		return _bildChannel.Extract(img, _bildChannel.Red)
	case "Green" :
		return _bildChannel.Extract(img, _bildChannel.Green)
	case "Blue" :
		return _bildChannel.Extract(img, _bildChannel.Blue)
	case "Alpha" :
		return _bildChannel.Extract(img, _bildChannel.Alpha)
	default :
		return img
	}
}

// Painting

func ipFloodFill(img image.Image, point [2]int, clr [4]uint8, fuzzy uint8) image.Image {

	// help needed: struggling to provide a proper color RGBA to FloodFill

	// lot of troubles with this function color.NRGBA - need help.
	//return paint.FloodFill(img, image.Pt(point[0], point[1]), color.NRGBA(clr[0], clr[1], clr[2], clr[3]), fuzzy)
	return img;
}

// Image Effects

func ipDilate(img image.Image, radius float64) image.Image {
	return _bildEffect.Dilate(img, radius);
}

func ipEdgeDetection(img image.Image, radius float64) image.Image {
	return _bildEffect.Dilate(img, radius);
}

func ipEmboss(img image.Image) image.Image {
	return _bildEffect.Emboss(img);
}

func ipErode(img image.Image, radius float64) image.Image {
	return _bildEffect.Erode(img, radius);
}

func ipGrayscale(img image.Image) image.Image {
	return _bildEffect.Grayscale(img);
}

func ipInvert(img image.Image) image.Image {
	return _bildEffect.Invert(img);
}

func ipMedian(img image.Image, radius float64) image.Image {
	return _bildEffect.Median(img, radius);
}

func ipSepia(img image.Image) image.Image {
	return _bildEffect.Sepia(img);
}

func ipSharpen(img image.Image) image.Image {
	return _bildEffect.Sharpen(img);
}

func ipSobel(img image.Image) image.Image {
	return _bildEffect.Sobel(img);
}

func ipUnsharpMask(img image.Image, radius float64, amount float64) image.Image {
	return _bildEffect.UnsharpMask(img, radius, amount);
}

