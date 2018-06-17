package message

import "fmt"

//
//        MatchingRuleAssertion ::= SEQUENCE {
//             matchingRule    [1] MatchingRuleId OPTIONAL,
//             type            [2] AttributeDescription OPTIONAL,
//             matchValue      [3] AssertionValue,
//             dnAttributes    [4] BOOLEAN DEFAULT FALSE }
func readTaggedMatchingRuleAssertion(bytes *Bytes, class int, tag int) (ret MatchingRuleAssertion, err error) {
	err = bytes.ReadSubBytes(class, tag, ret.readComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readTaggedMatchingRuleAssertion:\n%s", err.Error())}
		return
	}
	return
}
func (matchingruleassertion *MatchingRuleAssertion) readComponents(bytes *Bytes) (err error) {
	err = matchingruleassertion.readMatchingRule(bytes)
	if err != nil {
		return LdapError{fmt.Sprintf("readComponents: %s", err.Error())}
	}
	err = matchingruleassertion.readType(bytes)
	if err != nil {
		return LdapError{fmt.Sprintf("readComponents: %s", err.Error())}
	}
	matchingruleassertion.matchValue, err = readTaggedAssertionValue(bytes, classContextSpecific, TagMatchingRuleAssertionMatchValue)
	if err != nil {
		return LdapError{fmt.Sprintf("readComponents: %s", err.Error())}
	}
	matchingruleassertion.dnAttributes, err = readTaggedBOOLEAN(bytes, classContextSpecific, TagMatchingRuleAssertionDnAttributes)
	if err != nil {
		return LdapError{fmt.Sprintf("readComponents: %s", err.Error())}
	}
	return
}
func (matchingruleassertion *MatchingRuleAssertion) readMatchingRule(bytes *Bytes) (err error) {
	var tagAndLength TagAndLength
	tagAndLength, err = bytes.PreviewTagAndLength()
	if err != nil {
		return LdapError{fmt.Sprintf("readMatchingRule: %s", err.Error())}
	}
	if tagAndLength.Tag == TagMatchingRuleAssertionMatchingRule {
		var matchingRule MatchingRuleId
		matchingRule, err = readTaggedMatchingRuleId(bytes, classContextSpecific, TagMatchingRuleAssertionMatchingRule)
		if err != nil {
			return LdapError{fmt.Sprintf("readMatchingRule: %s", err.Error())}
		}
		matchingruleassertion.matchingRule = matchingRule.Pointer()
	}
	return
}
func (matchingruleassertion *MatchingRuleAssertion) readType(bytes *Bytes) (err error) {
	var tagAndLength TagAndLength
	tagAndLength, err = bytes.PreviewTagAndLength()
	if err != nil {
		return LdapError{fmt.Sprintf("readType: %s", err.Error())}
	}
	if tagAndLength.Tag == TagMatchingRuleAssertionType {
		var attributedescription AttributeDescription
		attributedescription, err = readTaggedAttributeDescription(bytes, classContextSpecific, TagMatchingRuleAssertionType)
		if err != nil {
			return LdapError{fmt.Sprintf("readType: %s", err.Error())}
		}
		matchingruleassertion.type_ = attributedescription.Pointer()
	}
	return
}
func (m MatchingRuleAssertion) writeTagged(bytes *Bytes, class int, tag int) (size int) {
	if m.dnAttributes != BOOLEAN(false) {
		size += m.dnAttributes.writeTagged(bytes, classContextSpecific, TagMatchingRuleAssertionDnAttributes)
	}
	size += m.matchValue.writeTagged(bytes, classContextSpecific, TagMatchingRuleAssertionMatchValue)
	if m.type_ != nil {
		size += m.type_.writeTagged(bytes, classContextSpecific, TagMatchingRuleAssertionType)
	}
	if m.matchingRule != nil {
		size += m.matchingRule.writeTagged(bytes, classContextSpecific, TagMatchingRuleAssertionMatchingRule)
	}
	size += bytes.WriteTagAndLength(class, isCompound, tag, size)
	return
}

//
//        MatchingRuleAssertion ::= SEQUENCE {
//             matchingRule    [1] MatchingRuleId OPTIONAL,
//             type            [2] AttributeDescription OPTIONAL,
//             matchValue      [3] AssertionValue,
//             dnAttributes    [4] BOOLEAN DEFAULT FALSE }
func (m MatchingRuleAssertion) write(bytes *Bytes) (size int) {
	return m.writeTagged(bytes, classUniversal, tagSequence)
}
