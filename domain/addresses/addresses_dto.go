package addresses

type Address struct {
	AID           int64  `json:"aid"`
	ZipCode       int64  `json:"zip_code"`
	StreetName    string `json:"street_name"`
	PidReferences int64  `json:"p___id"`
}
