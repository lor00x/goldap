package message

import "fmt"

func (attributevalueassertion *AttributeValueAssertion) readComponents(bytes *Bytes) (err error) {
	attributevalueassertion.attributeDesc, err = readAttributeDescription(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	attributevalueassertion.assertionValue, err = readAssertionValue(bytes)
	if err != nil {
		err = LdapError{fmt.Sprintf("readComponents:\n%s", err.Error())}
		return
	}
	return
}

//
//        AttributeValueAssertion ::= SEQUENCE {
//             attributeDesc   AttributeDescription,
//             assertionValue  AssertionValue }
func (a AttributeValueAssertion) write(bytes *Bytes) (size int) {
	size += a.assertionValue.write(bytes)
	size += a.attributeDesc.write(bytes)
	size += bytes.WriteTagAndLength(classUniversal, isCompound, tagSequence, size)
	return
}
func (a AttributeValueAssertion) writeTagged(bytes *Bytes, class int, tag int) (size int) {
	size += a.assertionValue.write(bytes)
	size += a.attributeDesc.write(bytes)
	size += bytes.WriteTagAndLength(class, isCompound, tag, size)
	return
}

//
//        AttributeValueAssertion ::= SEQUENCE {
//             attributeDesc   AttributeDescription,
//             assertionValue  AssertionValue }
func (a AttributeValueAssertion) size() (size int) {
	size += a.attributeDesc.size()
	size += a.assertionValue.size()
	size += sizeTagAndLength(tagSequence, size)
	return
}
func (a AttributeValueAssertion) sizeTagged(tag int) (size int) {
	size += a.attributeDesc.size()
	size += a.assertionValue.size()
	size += sizeTagAndLength(tag, size)
	return
}
