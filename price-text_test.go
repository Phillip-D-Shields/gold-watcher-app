package main

import "testing"

const OPEN_PRICE = "Open: 2721.4600 USD"

func TestApp_GetPriceText(t *testing.T) {
	open, _, _ := myApp.getPriceText()
	if open.Text != OPEN_PRICE {
		t.Errorf("Expected Open: %s, got %s", OPEN_PRICE, open.Text)
	}
}
