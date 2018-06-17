package message

import "fmt"

//             or              [1] SET SIZE (1..MAX) OF filter Filter,
func readFilterOr(bytes *Bytes) (filteror FilterOr, err error) {
	err = bytes.ReadSubBytes(classContextSpecific, TagFilterOr, filteror.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readFilterOr:\n%s", err.Error())}
		return
	}
	return
}
func (filteror *FilterOr) readComponents(bytes *Bytes) (err error) {
	count := 0
	for bytes.HasMoreData() {
		count++
		var filter Filter
		filter, err = readFilter(bytes)
		if err != nil {
			err = LdapError{fmt.Sprintf("readComponents (filter %d): %s", count, err.Error())}
			return
		}
		*filteror = append(*filteror, filter)
	}
	if len(*filteror) == 0 {
		err = LdapError{"readComponents: expecting at least one Filter"}
		return
	}
	return
}
