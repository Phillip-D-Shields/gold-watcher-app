package main

import (
	"time"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (app *Config) makeUI() {
	// get current price of gold
	openPrice, currentPrice, priceChange := app.getPriceText()
	// store price data
	priceContent := container.NewGridWithColumns(3, openPrice, currentPrice, priceChange)

	app.PriceContainer = priceContent

	// get toolbar
	toolbar := app.getToolBar()
	app.Toolbar = toolbar

	priceTabContent := app.pricesTab()

	// app tabs
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Prices", theme.HomeIcon(), priceTabContent),
		container.NewTabItemWithIcon("Holdings", theme.StorageIcon(), canvas.NewText("holdings will go here", nil)),
	)

	tabs.SetTabLocation(container.TabLocationTop)
	// add price data to the UI
	finalContent := container.NewVBox(priceContent, toolbar, tabs)

	app.MainWindow.SetContent(finalContent)

	go func() {
		for range time.Tick(time.Second * 5) {
			app.refreshPriceContent()
		}
	}()
}

func (app *Config) refreshPriceContent() {
	app.InfoLog.Print("Refreshing price content")
	openPrice, currentPrice, priceChange := app.getPriceText()
	app.PriceContainer.Objects[0] = openPrice
	app.PriceContainer.Objects[1] = currentPrice
	app.PriceContainer.Objects[2] = priceChange
	app.PriceContainer.Refresh()

	chart := app.getChart()
	app.PriceChartContainer.Objects[0] = chart
	app.PriceChartContainer.Refresh()
}
