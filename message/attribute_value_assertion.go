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
