package message

func (filter FilterAnd) getFilterTag() int {
	return TagFilterAnd
}
func (filter FilterOr) getFilterTag() int {
	return TagFilterOr
}
func (filter FilterNot) getFilterTag() int {
	return TagFilterNot
}
func (filter FilterEqualityMatch) getFilterTag() int {
	return TagFilterEqualityMatch
}
func (filter FilterSubstrings) getFilterTag() int {
	return TagFilterSubstrings
}
func (filterAnd FilterGreaterOrEqual) getFilterTag() int {
	return TagFilterGreaterOrEqual
}
func (filterAnd FilterLessOrEqual) getFilterTag() int {
	return TagFilterLessOrEqual
}
func (filterAnd FilterPresent) getFilterTag() int {
	return TagFilterPresent
}
func (filterAnd FilterApproxMatch) getFilterTag() int {
	return TagFilterApproxMatch
}
func (filterAnd FilterExtensibleMatch) getFilterTag() int {
	return TagFilterExtensibleMatch
}
