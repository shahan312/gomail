package gomail

// Address struct
type Address struct {
	Name    string
	Address string
}

// SetAddressListHeader sets an address to the given header field
func (m *Message) SetAddressListHeader(field string, list []Address) {
	val := make([]string, len(list), len(list))

	for i := range list {
		val[i] = m.FormatAddress(list[i].Address, list[i].Name)
	}

	m.header[field] = val
}

// Flag to turn ExportBcc on / off
var ExportBcc = false
