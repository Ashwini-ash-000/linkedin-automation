package stealth

import "github.com/go-rod/rod"

// Apply applies stealth patches to an active page context
func Apply(page *rod.Page) {
	page.MustEval(`() => {
		Object.defineProperty(navigator, 'webdriver', {
			get: () => undefined
		});
		return true;
	}`)
}
