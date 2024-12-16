package country_codes

type Codes struct{}

func NewCodes() *Codes {
	return &Codes{}
}

func (c *Codes) IsReal(country string) bool {
	for _, countryCode := range countryCodesList {
		if countryCode == country {
			return true
		}
	}

	return false
}
