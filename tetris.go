package main
import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"
	"os"
	"fmt"
	"image"
	"image/draw"
	_ "golang.org/x/image/bmp"
)

func appMain(driver gxui.Driver) {
	file := "./img/block.bmp"
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Failed to open image '%s': %v\n", file, err)
		os.Exit(1)
	}

	source, _, err := image.Decode(f)
	if err != nil {
		fmt.Printf("Failed to read image '%s': %v\n", file, err)
		os.Exit(1)
	}

	theme := flags.CreateTheme(driver)
	img := theme.CreateImage()

	//mx := source.Bounds().Max
	window := theme.CreateWindow(1000, 1000, "Tetris")
	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(img)

	rgba := image.NewRGBA(source.Bounds())
	draw.Draw(rgba, source.Bounds(), source, image.ZP, draw.Src)
	texture := driver.CreateTexture(rgba, 1)
	img.SetTexture(texture)

	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}
