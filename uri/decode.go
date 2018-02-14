package uri

import "strings"

// Naive and basic implementation of decodeuricomponent from javascript
// Handles only simple cases
// In a future can be extendad with
// https://chromium.googlesource.com/v8/v8/+/4.3.49/src/uri.js?autodive=0%2F%2F
// Or replaced with proper implementation

var replacer = strings.NewReplacer(
	"%2D", "-",
	"%2E", ".",
	"%23", "#",
	"%2F", "/",
)

func URIDecodeComponent(source string) (string, error) {
	return replacer.Replace(source), nil
}
