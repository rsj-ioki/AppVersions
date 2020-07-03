package presentation

import (
	"stefma.guru/appVersions/usecase"
)

func HelpText(domain string) string {
  return  `Please add a 'ios' or 'android' query to the url.
The output can be changed via the 'format' [json, pretty] query parameter.

Examples:
https://` + domain + `?android=com.ioki.hamburg,com.ioki.wittlich
https://` + domain + `?android=com.ioki.hamburg&format=json
https://` + domain + `?ios=ioki-hamburg/id1400408720,ioki-wittlich/id1377071496
https://` + domain + `?ios=ioki-hamburg/id1400408720&format=json
https://` + domain + `?android=com.ioki.hamburg,com.ioki.wittlich&ios=ioki-hamburg/id1400408720&format=json`
}

func FormatOutput(format string, androidApps []usecase.App, iosApps []usecase.App) string {
  switch format {
  case "json":
    return formatToJson(androidApps, iosApps)
  case "pretty":
    return formatToPretty(androidApps, iosApps)
  default:
    return formatToPretty(androidApps, iosApps)
  }
}
