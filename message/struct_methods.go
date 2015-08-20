package message

func (l *LDAPOID) String() string {
	return string(*l)
}

func (l *LDAPOID) Bytes() []byte {
	return []byte(*l)
}

func (l *OCTETSTRING) String() string {
	return string(*l)
}

func (l *OCTETSTRING) Bytes() []byte {
	return []byte(*l)
}

func (l *INTEGER) Int() int {
	return int(*l)
}

func (l *ENUMERATED) Int() int {
	return int(*l)
}

func (l *BOOLEAN) Bool() bool {
	return bool(*l)
}

func (l *LDAPMessage) MessageID() MessageID {
	return l.messageID
}

func (l *LDAPMessage) Controls() *Controls {
	return l.controls
}

func (l *LDAPMessage) ProtocolOp() ProtocolOp {
	return l.protocolOp
}

func (b *BindRequest) Name() LDAPDN {
	return b.name
}

func (b *BindRequest) Authentication() AuthenticationChoice {
	return b.authentication
}

func (e *ExtendedRequest) RequestName() LDAPOID {
	return e.requestName
}

func (e *ExtendedRequest) RequestValue() *OCTETSTRING {
	return e.requestValue
}

func (s *SearchRequest) BaseObject() LDAPDN {
	return s.baseObject
}
func (s *SearchRequest) Scope() ENUMERATED {
	return s.scope
}
func (s *SearchRequest) DerefAliases() ENUMERATED {
	return s.derefAliases
}
func (s *SearchRequest) SizeLimit() INTEGER {
	return s.sizeLimit
}

func (s *SearchRequest) TimeLimit() INTEGER {
	return s.timeLimit
}
func (s *SearchRequest) TypesOnly() BOOLEAN {
	return s.typesOnly
}
func (s *SearchRequest) Attributes() AttributeSelection {
	return s.attributes
}

func (s *SearchRequest) Filter() Filter {
	return s.filter
}

func (c *CompareRequest) Entry() LDAPDN {
	return c.entry
}

func (c *CompareRequest) Ava() *AttributeValueAssertion {
	return &c.ava
}

func (a *AttributeValueAssertion) AttributeDesc() AttributeDescription {
	return a.attributeDesc
}

func (a *AttributeValueAssertion) AssertionValue() AssertionValue {
	return a.assertionValue
}

func (a *AddRequest) Entry() LDAPDN {
	return a.entry
}

func (a *AddRequest) Attributes() AttributeList {
	return a.attributes
}

func (a *Attribute) Type_() AttributeDescription {
	return a.type_
}
func (a *Attribute) Vals() []AttributeValue {
	return a.vals
}

func (m *ModifyRequest) Object() LDAPDN {
	return m.object
}
func (m *ModifyRequest) Changes() []ModifyRequestChange {
	return m.changes
}

func (m *ModifyRequestChange) Operation() ENUMERATED {
	return m.operation
}

func (m *ModifyRequestChange) Modification() *PartialAttribute {
	return &m.modification
}

func (p *PartialAttribute) Type_() AttributeDescription {
	return p.type_
}
func (p *PartialAttribute) Vals() []AttributeValue {
	return p.vals
}

func (c *Control) ControlType() LDAPOID {
	return c.controlType
}

func (c *Control) Criticality() BOOLEAN {
	return c.criticality
}

func (c *Control) ControlValue() *OCTETSTRING {
	return c.controlValue
}

func (a *FilterEqualityMatch) AttributeDesc() AttributeDescription {
	return a.attributeDesc
}

func (a *FilterEqualityMatch) AssertionValue() AssertionValue {
	return a.assertionValue
}

func (a *FilterGreaterOrEqual) AttributeDesc() AttributeDescription {
	return a.attributeDesc
}

func (a *FilterGreaterOrEqual) AssertionValue() AssertionValue {
	return a.assertionValue
}

func (a *FilterLessOrEqual) AttributeDesc() AttributeDescription {
	return a.attributeDesc
}

func (a *FilterLessOrEqual) AssertionValue() AssertionValue {
	return a.assertionValue
}

func (a *FilterApproxMatch) AttributeDesc() AttributeDescription {
	return a.attributeDesc
}

func (a *FilterApproxMatch) AssertionValue() AssertionValue {
	return a.assertionValue
}

func (s *FilterSubstrings) Type_() AttributeDescription {
	return s.type_
}

func (s *FilterSubstrings) Substrings() []Substring {
	return s.substrings
}
