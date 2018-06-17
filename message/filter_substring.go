package message

import "fmt"

//             substrings      [4] SubstringFilter,
func readFilterSubstrings(bytes *Bytes) (filtersubstrings FilterSubstrings, err error) {
	var substringfilter SubstringFilter
	substringfilter, err = readTaggedSubstringFilter(bytes, classContextSpecific, TagFilterSubstrings)
	if err != nil {
		err = LdapError{fmt.Sprintf("readFilterSubstrings:\n%s", err.Error())}
		return
	}
	filtersubstrings = FilterSubstrings(substringfilter)
	return
}

//
//        SubstringFilter ::= SEQUENCE {
//             type           AttributeDescription,
//             substrings     SEQUENCE SIZE (1..MAX) OF substring CHOICE {
//                  initial [0] AssertionValue,  -- can occur at most once
//                  any     [1] AssertionValue,
//                  final   [2] AssertionValue } -- can occur at most once
//             }
func readTaggedSubstringFilter(bytes *Bytes, class int, tag int) (substringfilter SubstringFilter, err error) {
	err = bytes.ReadSubBytes(class, tag, substringfilter.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readTaggedSubstringFilter:\n%s", err.Error())}
		return
	}
	return
}
func (substringfilter *SubstringFilter) readComponents(bytes *Bytes) (err error) {
	substringfilter.type_, err = readAttributeDescription(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	err = substringfilter.readSubstrings(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	return
}
func (substringfilter *SubstringFilter) readSubstrings(bytes *Bytes) (err error) {
	err = bytes.ReadSubBytes(classUniversal, tagSequence, substringfilter.readSubstringsComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readSubstrings:\n%s", err.Error())}
		return
	}
	return
}
func (substringfilter *SubstringFilter) readSubstringsComponents(bytes *Bytes) (err error) {
	var foundInitial = 0
	var foundFinal = 0
	var tagAndLength TagAndLength
	for bytes.HasMoreData() {
		tagAndLength, err = bytes.PreviewTagAndLength()
		if err != nil {
			err = LdapError{fmt.Sprintf("readSubstringsComponents:\n%s", err.Error())}
			return
		}
		var assertionvalue AssertionValue
		switch tagAndLength.Tag {
		case TagSubstringInitial:
			foundInitial++
			if foundInitial > 1 {
				err = LdapError{"readSubstringsComponents: initial can occur at most once"}
				return
			}
			assertionvalue, err = readTaggedAssertionValue(bytes, classContextSpecific, TagSubstringInitial)
			if err != nil {
				err = LdapError{fmt.Sprintf("readSubstringsComponents:\n%s", err.Error())}
				return
			}
			substringfilter.substrings = append(substringfilter.substrings, SubstringInitial(assertionvalue))
		case TagSubstringAny:
			assertionvalue, err = readTaggedAssertionValue(bytes, classContextSpecific, TagSubstringAny)
			if err != nil {
				err = LdapError{fmt.Sprintf("readSubstringsComponents:\n%s", err.Error())}
				return
			}
			substringfilter.substrings = append(substringfilter.substrings, SubstringAny(assertionvalue))
		case TagSubstringFinal:
			foundFinal++
			if foundFinal > 1 {
				err = LdapError{"readSubstringsComponents: final can occur at most once"}
				return
			}
			assertionvalue, err = readTaggedAssertionValue(bytes, classContextSpecific, TagSubstringFinal)
			if err != nil {
				err = LdapError{fmt.Sprintf("readSubstringsComponents:\n%s", err.Error())}
				return
			}
			substringfilter.substrings = append(substringfilter.substrings, SubstringFinal(assertionvalue))
		default:
			err = LdapError{fmt.Sprintf("readSubstringsComponents: invalid tag %d", tagAndLength.Tag)}
			return
		}
	}
	if len(substringfilter.substrings) == 0 {
		err = LdapError{"readSubstringsComponents: expecting at least one substring"}
		return
	}
	return
}
