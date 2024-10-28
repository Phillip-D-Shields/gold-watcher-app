package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (app *Config) pricesTab() *fyne.Container {
	chart := app.getChart()
	chartContainer := container.NewVBox(chart)
	app.PriceChartContainer = chartContainer

	return chartContainer
}

func (app *Config) getChart() *canvas.Image {
	apiURL := fmt.Sprintf("https://goldprice.org/charts/gold_3d_b_o_%s_x.png", strings.ToLower(currency))

	var img *canvas.Image
	err := app.downloadFile(apiURL, "gold_chart.png")
	if err != nil {
		img = canvas.NewImageFromResource(resourceUnreachablePng)
	} else {
		img = canvas.NewImageFromFile("gold_chart.png")
	}

	img.SetMinSize(fyne.Size{Width: 770, Height: 410})
	img.FillMode = canvas.ImageFillOriginal

	return img
}

func (app *Config) downloadFile(URL, filename string) error {
	// get response bytes
	response, err := app.HTTPClient.Get(URL)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("failed to download file")
	}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return err
	}

	out, err := os.Create(fmt.Sprintf("./%s", filename))
	if err != nil {
		return err
	}

	defer out.Close()

	err = png.Encode(out, img)
	if err != nil {
		return err
	}

	return nil
}
