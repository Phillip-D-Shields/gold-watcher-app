package main

import "fyne.io/fyne/v2/container"

func (app *Config) makeUI() {
	// get current price of gold
	openPrice, currentPrice, priceChange := app.getPriceText()
	// store price data
	priceContent := container.NewGridWithColumns(3, openPrice, currentPrice, priceChange)

	app.PriceContainer = priceContent

	// get toolbar
	toolbar := app.getToolBar(app.MainWindow)

	// add price data to the UI
	finalContent := container.NewVBox(priceContent, toolbar)

	app.MainWindow.SetContent(finalContent)
}
