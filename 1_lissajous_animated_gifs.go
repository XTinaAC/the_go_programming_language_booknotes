/*
	Create a sequence of bit-mapped Lissajous figures, 
	& encode it as a GIF animation.
	
	利萨茹（Lissajous）曲线（又称利萨茹图形、鲍迪奇(Bowditch)曲线)
		是两个沿着互相垂直方向的正弦振动的合成轨迹:
	x(θ) = a * sin(θ)
	y(θ) = b * sin(nθ + 𝝋), n是两个正弦振动的频率比。
*/
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"bytes"
	"math"
	"math/rand"
)

//【composite literals】：a compact notation for instantiating any of Go's
//【composite types】（e.g., slices, structs）from a sequence of elements
var palette = []color.Color{ color.White, color.RGBA{ 0xff, 0, 0, 0xff }, color.RGBA{ 0, 0xff, 0, 0xff }, color.RGBA{ 0, 0, 0xff, 0xff } }

//【constants】：values are fixed at compile time; must be of type number/string/boolean
const (
	whiteIndex = 0 	// 1st color in palette
	redIndex = 1	// 2nd color in palette
	greenIndex = 2	// 3rd color in palette
	blueIndex = 3 	// 4th color in palette
)

func main() {
	// 控制台标准输出出现乱码-待解决...
	// lissajous(os.Stdout) 

	// 改为文件输出：https://blog.csdn.net/ocean_this_is_it/article/details/129850517
	buf := &bytes.Buffer{}
	lissajous(buf)
	if err := ioutil.WriteFile("lissajous.gif", buf.Bytes(), 0666); err != nil {
		panic(err)
	}
}

func lissajous(out io.Writer) {
	const (
		cycles = 5		// number of complete x oscillator revolutions
		res = 0.001		// angular resolution
		size = 100		// image canvas covers [-size ... +size]
		nframes = 64 	// number of animation frames
		delay = 8 		// delay between frames (in 10ms units)
	)
	
	// all other fields than LoopCount of the【struct literal】have the ZERO value
	anim := gif.GIF{ LoopCount: nframes }

	// frequency of the y oscillator (relative to the x oscillator) is a random number between 0 and 3
	freq := rand.Float64() * 3.0

	// phase of the y oscillator (relative to the x oscillator) is initially 0 but increases 0.1 with each frame
	phase := 0.0

	// the outer loop - producing 64 frames of the animation
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		// the inner loop - running 2 oscillators until the x oscillator has completed 5 cycles
		for t := 0.0; t < cycles*2*math.Pi; t += res {

			// x(t) = sin(t)
			x := math.Sin(t)

			// y(t) = sin(t * freq + phase)
			y := math.Sin(t*freq + phase)

			// rotating the color index
			var colorIndex uint8 = whiteIndex;
			if i%30 <= 10 {
				colorIndex = redIndex
			} else if i%30 <= 20 {
				colorIndex = greenIndex
			} else {
				colorIndex = blueIndex
			}

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}

		phase += 0.1

		// access individual fields of a struct using dot notation
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	// Encode the sequence of frames & delays into GIF format, and write it to the output stream.
	gif.EncodeAll(out, &anim)
}