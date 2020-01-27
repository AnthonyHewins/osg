package investigate

import (
	"regexp"
)

// TODO tries? Lots of room for improvement here
var url_regex = regexp.MustCompile(
	"(?i)((ht|f)tps?://)?([a-z][a-z0-9]+[.])+(com?|rs|me|edu|gov|mil|net|org|biz|info|name|museum|us|ca|uk)(:[0-9]{1,6})?(/($|[a-z0-9.,;?'\\+&%$#=~_-]+))*",
)

// Be sure to use stemming, ie dont use analytics, use analytic
// Also it's case insensitive by default
var suspicious_words = regexp.MustCompile(
	"(?i)(analytic|telemetry)+",
)
