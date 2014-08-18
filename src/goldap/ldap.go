package goldap

import (
	"errors"
	"fmt"
)

type LdapError struct {
	Msg string
}
func (e LdapError) Error() string { return e.Msg }


func ReadBOOLEAN(bytes *Bytes) (ret BOOLEAN, err error) {
	return ReadTaggedBOOLEAN(bytes, classUniversal, tagBoolean)
}
func ReadTaggedBOOLEAN(bytes *Bytes, class int, tag int) (ret BOOLEAN, err error) {
	tagAndLength, err := bytes.ParseTagAndLength()
	if err != nil {
		return
	}
	err = tagAndLength.Expect(class, tag, isNotCompound)
	if err != nil {
		return
	}
	var boolean bool
	boolean, err = bytes.ParseBool(tagAndLength.GetLength())
	return BOOLEAN(boolean), err
}

func ReadINTEGER(bytes *Bytes) (ret INTEGER, err error) {
	tagAndLength, err := bytes.ParseTagAndLength()
	if err != nil {
		return
	}
	err = tagAndLength.Expect(classUniversal, tagInteger, isNotCompound)
	if err != nil {
		return
	}
	var integer int32
	integer, err = bytes.ParseInt32(tagAndLength.GetLength())
	return INTEGER(integer), err
}

func ReadPositiveINTEGER(bytes *Bytes) (ret INTEGER, err error) {
	ret, err = ReadINTEGER(bytes)
	if err != nil {
		return
	}
	if !(ret >= 0 && ret <= maxInt) {
		err = LdapError{fmt.Sprintf("Invalid INTEGER value %d ! Expected value between 0 and %d", ret, maxInt)}
	}
	return
}

func ReadENUMERATED(bytes *Bytes, allowedValues map[ENUMERATED]string) (ret ENUMERATED, err error) {
	tagAndLength, err := bytes.ParseTagAndLength()
	if err != nil {
		return ret, LdapError{fmt.Sprintf("ReadENUMERATED: %s", err.Error())}
	}
	err = tagAndLength.Expect(classUniversal, tagEnum, isNotCompound)
	if err != nil {
		return ret, LdapError{fmt.Sprintf("ReadENUMERATED: %s", err.Error())}
	}
	var integer int32
	integer, err = bytes.ParseInt32(tagAndLength.GetLength())
	if err != nil {
		return ret, LdapError{fmt.Sprintf("ReadENUMERATED: %s", err.Error())}
	}
	ret = ENUMERATED(integer)
	if _, ok := allowedValues[ret]; !ok {
		return ret, LdapError{fmt.Sprintf("ReadENUMERATED: Invalid ENUMERATED VALUE %d", ret)}
	}
	return
}

func ReadUTF8STRING(bytes *Bytes) (ret UTF8STRING, err error) {
	return ReadTaggedUTF8STRING(bytes, classUniversal, tagUTF8String)
}
func ReadTaggedUTF8STRING(bytes *Bytes, class int, tag int) (ret UTF8STRING, err error) {
	tagAndLength, err := bytes.ParseTagAndLength()
	if err != nil {
		return ret, errors.New(fmt.Sprintf("ReadTaggedUTF8STRING: %s", err.Error()))
	}
	err = tagAndLength.Expect(class, tag, isNotCompound)
	if err != nil {
		return ret, errors.New(fmt.Sprintf("ReadTaggedUTF8STRING: %s", err.Error()))
	}
	var utf8string string
	utf8string, err = bytes.ParseUTF8String(tagAndLength.GetLength())
	if err != nil {
		return ret, errors.New(fmt.Sprintf("ReadTaggedUTF8STRING: %s", err.Error()))
	}
	return UTF8STRING(utf8string), err
}

func ReadOCTETSTRING(bytes *Bytes) (ret OCTETSTRING, err error) {
	return ReadTaggedOCTETSTRING(bytes, classUniversal, tagOctetString)
}
func ReadTaggedOCTETSTRING(bytes *Bytes, class int, tag int) (ret OCTETSTRING, err error) {
	tagAndLength, err := bytes.ParseTagAndLength()
	if err != nil {
		return
	}
	err = tagAndLength.Expect(class, tag, isNotCompound)
	if err != nil {
		return
	}
	var octetstring []byte
	octetstring, err = bytes.ParseOCTETSTRING(tagAndLength.GetLength())
	if err != nil {
		return
	}
	return OCTETSTRING(octetstring), err
}

//   This appendix is normative.
//
//        Lightweight-Directory-Access-Protocol-V3 {1 3 6 1 1 18}
//        -- Copyright (C) The Internet Society (2006).  This version of
//        -- this ASN.1 module is part of RFC 4511; see the RFC itself
//        -- for full legal notices.
//        DEFINITIONS
//        IMPLICIT TAGS
//        EXTENSIBILITY IMPLIED ::=
//
//        BEGIN
//
//        LDAPMessage ::= SEQUENCE {
//             messageID       MessageID,
//             protocolOp      CHOICE {
//                  bindRequest           BindRequest,
//                  bindResponse          BindResponse,
//                  unbindRequest         UnbindRequest,
//                  searchRequest         SearchRequest,
//                  searchResEntry        SearchResultEntry,
//                  searchResDone         SearchResultDone,
//                  searchResRef          SearchResultReference,
//                  modifyRequest         ModifyRequest,
//                  modifyResponse        ModifyResponse,
//                  addRequest            AddRequest,
//                  addResponse           AddResponse,
//                  delRequest            DelRequest,
//                  delResponse           DelResponse,
//                  modDNRequest          ModifyDNRequest,
//                  modDNResponse         ModifyDNResponse,
//                  compareRequest        CompareRequest,
//                  compareResponse       CompareResponse,
//                  abandonRequest        AbandonRequest,
//                  extendedReq           ExtendedRequest,
//                  extendedResp          ExtendedResponse,
//                  ...,
//                  intermediateResponse  IntermediateResponse },
//             controls       [0] Controls OPTIONAL }
//
func NewLDAPMessage() *LDAPMessage { return &LDAPMessage{} }

func ReadLDAPMessage(bytes *Bytes) (message LDAPMessage, err error) {
	err = bytes.ParseSubBytes(classUniversal, tagSequence, message.ReadLDAPMessageComponents)
	if err != nil {
		err = errors.New(fmt.Sprintf("ReadLDAPMessage: %s", err.Error()))
	}
	return
}
func (message *LDAPMessage) ReadLDAPMessageComponents(bytes *Bytes) (err error) {
	message.messageID, err = ReadMessageID(bytes)
	if err != nil {
		return
	}
	message.protocolOp, err = ReadProtocolOp(bytes)
	if err != nil {
		return
	}
	if bytes.HasMoreData() {
		var controls Controls
		controls, err = ReadControls(bytes)
		if err != nil {
			return
		}
		message.controls = &controls
	}
	return
}

//        MessageID ::= INTEGER (0 ..  maxInt)
//
//        maxInt INTEGER ::= 2147483647 -- (2^^31 - 1) --
//
func ReadMessageID(bytes *Bytes) (ret MessageID, err error) {
	var integer INTEGER
	integer, err = ReadPositiveINTEGER(bytes)
	if err != nil {
		err = errors.New(fmt.Sprintf("ReadMessageID: %s", err.Error()))
		return
	}
	return MessageID(integer), err
}
func ReadProtocolOp(bytes *Bytes) (ret ProtocolOp, err error) {
	tagAndLength, err := bytes.PreviewTagAndLength()
	if err != nil {
		err = errors.New(fmt.Sprintf("ReadProtocolOp: %s", err.Error()))
		return
	}
	switch tagAndLength.GetTag() {
	case TagBindRequest:
		ret, err = ReadBindRequest(bytes)
	case TagBindResponse:
		ret, err = ReadBindResponse(bytes)
	case TagUnbindRequest:
		ret, err = ReadUnbindRequest(bytes)
	case TagSearchRequest:
		ret, err = ReadSearchRequest(bytes)
	case TagSearchResultEntry:
		ret, err = ReadSearchResultEntry(bytes)
	case TagSearchResultDone:
		ret, err = ReadSearchResultDone(bytes)
	default:
		err = LdapError{fmt.Sprintf("Invalid tag value for protocolOp. Got %d.", tagAndLength.GetTag())}
	}
	if err != nil {
		err = errors.New(fmt.Sprintf("ReadProtocolOp: %s", err.Error()))
	}
	return
}

//        LDAPString ::= OCTET STRING -- UTF-8 encoded,
//                                    -- [ISO10646] characters
func ReadLDAPString(bytes *Bytes) (ldapstring LDAPString, err error) {
	return ReadTaggedLDAPString(bytes, classUniversal, tagOctetString)
}
func ReadTaggedLDAPString(bytes *Bytes, class int, tag int) (ldapstring LDAPString, err error) {
	var utf8string UTF8STRING
	utf8string, err = ReadTaggedUTF8STRING(bytes, class, tag)
	if err != nil {
		err = errors.New(fmt.Sprintf("ReadTaggedLDAPString: %s", err.Error()))
		return
	}
	ldapstring = LDAPString(utf8string)
	return
}

//
//
//
//
//Sermersheim                 Standards Track                    [Page 54]
//
//
//RFC 4511                         LDAPv3                        June 2006
//
//
//        LDAPOID ::= OCTET STRING -- Constrained to <numericoid>
//                                 -- [RFC4512]
func ReadLDAPOID(bytes *Bytes) (ret LDAPOID, err error) {
	var octetstring OCTETSTRING
	octetstring, err = ReadOCTETSTRING(bytes)
	if err != nil {
		return
	}
	// @TODO: check RFC4512 for <numericoid>
	ret = LDAPOID(octetstring)
	return
}

//
//        LDAPDN ::= LDAPString -- Constrained to <distinguishedName>
//                              -- [RFC4514]
func ReadLDAPDN(bytes *Bytes) (ret LDAPDN, err error) {
	var ldapstring LDAPString
	ldapstring, err = ReadLDAPString(bytes)
	if err != nil {
		err = errors.New(fmt.Sprintf("ReadLDAPDN: %s", err.Error()))
		return
	}
	// @TODO: check RFC4514
	ret = LDAPDN(ldapstring)
	return
}

//
//        RelativeLDAPDN ::= LDAPString -- Constrained to <name-component>
//                                      -- [RFC4514]
func ReadRelativeLDAPDN(bytes *Bytes) (ret RelativeLDAPDN, err error) {
	var ldapstring LDAPString
	ldapstring, err = ReadLDAPString(bytes)
	// @TODO: check RFC4514
	ret = RelativeLDAPDN(ldapstring)
	return
}

//
//        AttributeDescription ::= LDAPString
//                                -- Constrained to <attributedescription>
//                                -- [RFC4512]
func ReadAttributeDescription(bytes *Bytes) (ret AttributeDescription, err error) {
	var ldapstring LDAPString
	ldapstring, err = ReadLDAPString(bytes)
	// @TODO: check RFC4512
	ret = AttributeDescription(ldapstring)
	return
}
func ReadTaggedAttributeDescription(bytes *Bytes, class int, tag int) (ret AttributeDescription, err error) {
	var ldapstring LDAPString
	ldapstring, err = ReadTaggedLDAPString(bytes, class, tag)
	// @TODO: check RFC4512
	if err != nil {
		err = errors.New(fmt.Sprintf("ReadTaggedAttributeDescription: %s", err.Error()))
		return
	}
	ret = AttributeDescription(ldapstring)
	return
}

//
//        AttributeValue ::= OCTET STRING
func ReadAttributeValue(bytes *Bytes) (ret AttributeValue, err error) {
	var octetstring OCTETSTRING
	octetstring, err = ReadOCTETSTRING(bytes)
	if err != nil {
		return
	}
	ret = AttributeValue(octetstring)
	return
}

//
//        AttributeValueAssertion ::= SEQUENCE {
//             attributeDesc   AttributeDescription,
//             assertionValue  AssertionValue }
func ReadTaggedAttributeValueAssertion(bytes *Bytes, class int, tag int) (ret AttributeValueAssertion, err error){
	err = bytes.ParseSubBytes(class, tag, ret.ReadAttributeValueAssertionComponents)
	return
}

func (attributevalueassertion *AttributeValueAssertion) ReadAttributeValueAssertionComponents(bytes *Bytes) (err error) {
	attributevalueassertion.attributeDesc, err = ReadAttributeDescription(bytes)
	if err != nil {
		return
	}
	attributevalueassertion.assertionValue, err = ReadAssertionValue(bytes)
	if err != nil {
		return
	}
	return
}

//
//        AssertionValue ::= OCTET STRING
func ReadAssertionValue(bytes *Bytes) (assertionvalue AssertionValue, err error) {
	return ReadTaggedAssertionValue(bytes, classUniversal, tagOctetString)
}
func ReadTaggedAssertionValue(bytes *Bytes, class int, tag int) (assertionvalue AssertionValue, err error) {
	var octetstring OCTETSTRING
	octetstring, err = ReadTaggedOCTETSTRING(bytes, class, tag)
	if err != nil {
		return
	}
	assertionvalue = AssertionValue(octetstring)
	return
}

//
//        PartialAttribute ::= SEQUENCE {
//             type       AttributeDescription,
//             vals       SET OF value AttributeValue }
func ReadPartialAttribute(bytes *Bytes) (partialattribute PartialAttribute, err error){
	err = bytes.ParseSubBytes(classUniversal, tagSequence, partialattribute.ReadPartialAttributeComponents)
	return
}

func (partialattribute *PartialAttribute) ReadPartialAttributeComponents(bytes *Bytes) (err error){
	partialattribute.type_, err = ReadAttributeDescription(bytes)
	if err != nil {
		return
	}
	err = bytes.ParseSubBytes(classUniversal, tagSet, partialattribute.ReadPartialAttributeValsComponents)
	if err != nil {
		return
	}
	return
}
func (partialattribute *PartialAttribute) ReadPartialAttributeValsComponents(bytes *Bytes) (err error){
	for bytes.HasMoreData(){
		var attributevalue AttributeValue
		attributevalue, err = ReadAttributeValue(bytes)
		if err != nil {
			return
		}
		partialattribute.vals = append(partialattribute.vals, attributevalue)
	}
	return
}
//
//        Attribute ::= PartialAttribute(WITH COMPONENTS {
//             ...,
//             vals (SIZE(1..MAX))})
//
//        MatchingRuleId ::= LDAPString
func ReadTaggedMatchingRuleId(bytes *Bytes, class int, tag int) (matchingruleid MatchingRuleId, err error) {
	var ldapstring LDAPString
	ldapstring, err = ReadTaggedLDAPString(bytes, class, tag)
	if err != nil {
		return
	}
	matchingruleid = MatchingRuleId(ldapstring)
	return

}

//
//        LDAPResult ::= SEQUENCE {
//             resultCode         ENUMERATED {
//                  success                      (0),
//                  operationsError              (1),
//                  protocolError                (2),
//                  timeLimitExceeded            (3),
//                  sizeLimitExceeded            (4),
//                  compareFalse                 (5),
//                  compareTrue                  (6),
//                  authMethodNotSupported       (7),
//                  strongerAuthRequired         (8),
//                       -- 9 reserved --
//                  referral                     (10),
//                  adminLimitExceeded           (11),
//                  unavailableCriticalExtension (12),
//                  confidentialityRequired      (13),
//                  saslBindInProgress           (14),
//
//
//
//Sermersheim                 Standards Track                    [Page 55]
//
//
//RFC 4511                         LDAPv3                        June 2006
//
//
//                  noSuchAttribute              (16),
//                  undefinedAttributeType       (17),
//                  inappropriateMatching        (18),
//                  constraintViolation          (19),
//                  attributeOrValueExists       (20),
//                  invalidAttributeSyntax       (21),
//                       -- 22-31 unused --
//                  noSuchObject                 (32),
//                  aliasProblem                 (33),
//                  invalidDNSyntax              (34),
//                       -- 35 reserved for undefined isLeaf --
//                  aliasDereferencingProblem    (36),
//                       -- 37-47 unused --
//                  inappropriateAuthentication  (48),
//                  invalidCredentials           (49),
//                  insufficientAccessRights     (50),
//                  busy                         (51),
//                  unavailable                  (52),
//                  unwillingToPerform           (53),
//                  loopDetect                   (54),
//                       -- 55-63 unused --
//                  namingViolation              (64),
//                  objectClassViolation         (65),
//                  notAllowedOnNonLeaf          (66),
//                  notAllowedOnRDN              (67),
//                  entryAlreadyExists           (68),
//                  objectClassModsProhibited    (69),
//                       -- 70 reserved for CLDAP --
//                  affectsMultipleDSAs          (71),
//                       -- 72-79 unused --
//                  other                        (80),
//                  ...  },
//             matchedDN          LDAPDN,
//             diagnosticMessage  LDAPString,
//             referral           [3] Referral OPTIONAL }
func ReadTaggedLDAPResult(bytes *Bytes, class int, tag int) (ret LDAPResult, err error){
	err = bytes.ParseSubBytes(class, tag, ret.ReadLDAPResultComponents)
	if err != nil {
		err = fmt.Errorf("ReadLDAPResult: %s", err.Error())
	}
	return
}
func ReadLDAPResult(bytes *Bytes) (ldapresult LDAPResult, err error) {
	return ReadTaggedLDAPResult(bytes, classUniversal, tagSequence)
}
func (ldapresult *LDAPResult) ReadLDAPResultComponents(bytes *Bytes) (err error) {
	ldapresult.resultCode, err = ReadENUMERATED(bytes, EnumeratedLDAPResultCode)
	if err != nil {
		return
	}
	ldapresult.matchedDN, err = ReadLDAPDN(bytes)
	if err != nil {
		return
	}
	ldapresult.diagnosticMessage, err = ReadLDAPString(bytes)
	if err != nil {
		return
	}
	if bytes.HasMoreData() {
		var referral Referral
		referral, err = ReadReferral(bytes)
		if err != nil {
			return
		}
		ldapresult.referral = &referral
	}
	return
}

//
//        Referral ::= SEQUENCE SIZE (1..MAX) OF uri URI
func ReadReferral(bytes *Bytes) (referral Referral, err error) {
	err = bytes.ParseSubBytes(classUniversal, tagSequence, referral.ReadReferralComponents)
	return
}
func (referral *Referral) ReadReferralComponents(bytes *Bytes) (err error) {
	for bytes.HasMoreData() {
		var uri URI
		uri, err = ReadURI(bytes)
		if err != nil {
			return
		}
		*referral = append(*referral, uri)
	}
	if len(*referral) == 0 {
		return LdapError{"ReadReferral: expecting at least one URI"}
	}
	return
}

//
//        URI ::= LDAPString     -- limited to characters permitted in
//                               -- URIs
func ReadURI(bytes *Bytes) (uri URI, err error) {
	var ldapstring LDAPString
	ldapstring, err = ReadLDAPString(bytes)
	// @TODO: check permitted chars in URI
	if err != nil {
		return
	}
	uri = URI(ldapstring)
	return
}

//
//        Controls ::= SEQUENCE OF control Control
func ReadControls(bytes *Bytes) (controls Controls, err error) {
	err = bytes.ParseSubBytes(classUniversal, tagSequence, controls.ReadControlsComponents)
	return
}
func (controls *Controls) ReadControlsComponents(bytes *Bytes) (err error) {
	for bytes.HasMoreData() {
		var control Control
		control, err = ReadControl(bytes)
		if err != nil {
			return
		}
		*controls = append(*controls, control)
	}
	return
}

//
//        Control ::= SEQUENCE {
//             controlType             LDAPOID,
//             criticality             BOOLEAN DEFAULT FALSE,
//             controlValue            OCTET STRING OPTIONAL }
func ReadControl(bytes *Bytes) (control Control, err error) {
	err = bytes.ParseSubBytes(classUniversal, tagSequence, control.ReadControlComponents)
	return
}
func (control *Control) ReadControlComponents(bytes *Bytes) (err error) {
	control.controlType, err = ReadLDAPOID(bytes)
	if err != nil {
		return
	}
	control.criticality, err = ReadBOOLEAN(bytes)
	if err != nil {
		return
	}
	if bytes.HasMoreData() {
		var octetstring OCTETSTRING
		octetstring, err = ReadOCTETSTRING(bytes)
		if err != nil {
			return
		}
		control.controlValue = &octetstring
	}
	return
}

//
//
//
//
//Sermersheim                 Standards Track                    [Page 56]
//
//
//RFC 4511                         LDAPv3                        June 2006
//
//
//        BindRequest ::= [APPLICATION 0] SEQUENCE {
//             version                 INTEGER (1 ..  127),
//             name                    LDAPDN,
//             authentication          AuthenticationChoice }
func ReadBindRequest(bytes *Bytes) (bindrequest BindRequest, err error) {
	err = bytes.ParseSubBytes(classApplication, TagBindRequest, bindrequest.ReadBindRequestComponents)
	if err != nil {
		err = errors.New(fmt.Sprintf("ReadBindRequest: %s", err.Error()))
	}
	return
}
func (bindrequest *BindRequest) ReadBindRequestComponents(bytes *Bytes) (err error) {
	bindrequest.version, err = ReadINTEGER(bytes)
	if !(bindrequest.version >= BindRequestVersionMin && bindrequest.version <= BindRequestVersionMax) {
		err = LdapError{fmt.Sprintf("Invalid version %d. Must be between %d and %d", bindrequest.version, BindRequestVersionMin, BindRequestVersionMax)}
		return
	}
	bindrequest.name, err = ReadLDAPDN(bytes)
	if err != nil {
		return
	}
	bindrequest.authentication, err = ReadAuthenticationChoice(bytes)
	return
}

//
//        AuthenticationChoice ::= CHOICE {
//             simple                  [0] OCTET STRING,
//                                     -- 1 and 2 reserved
//             sasl                    [3] SaslCredentials,
//             ...  }
func ReadAuthenticationChoice(bytes *Bytes) (ret interface{}, err error) {
	tagAndLength, err := bytes.PreviewTagAndLength()
	if err != nil {
		err = errors.New(fmt.Sprintf("ReadAuthenticationChoice: %s", err.Error()))
		return
	}
	err = tagAndLength.ExpectClass(classContextSpecific)
	if err != nil {
		err = errors.New(fmt.Sprintf("ReadAuthenticationChoice: %s", err.Error()))
		return
	}
	switch tagAndLength.GetTag() {
	case TagAuthenticationChoiceSimple:
		ret, err = ReadTaggedOCTETSTRING(bytes, classContextSpecific, TagAuthenticationChoiceSimple)
	case TagAuthenticationChoiceSaslCredentials:
		ret, err = ReadSaslCredentials(bytes)
	default:
		err = LdapError{fmt.Sprintf("Invalid tag value for AuthenticationChoice. Got %d.", tagAndLength.GetTag())}
	}
	if err != nil {
		err = errors.New(fmt.Sprintf("ReadAuthenticationChoice: %s", err.Error()))
	}
	return
}

//
//        SaslCredentials ::= SEQUENCE {
//             mechanism               LDAPString,
//             credentials             OCTET STRING OPTIONAL }
//
func ReadSaslCredentials(bytes *Bytes) (authentication SaslCredentials, err error) {
	authentication = SaslCredentials{}
	err = bytes.ParseSubBytes(classContextSpecific, TagAuthenticationChoiceSaslCredentials, authentication.ReadSaslCredentialsComponents)
	return
}
func (authentication *SaslCredentials) ReadSaslCredentialsComponents(bytes *Bytes) (err error) {
	authentication.mechanism, err = ReadLDAPString(bytes)
	if err != nil {
		return
	}
	if bytes.HasMoreData() {
		var credentials OCTETSTRING
		credentials, err = ReadOCTETSTRING(bytes)
		if err != nil {
			return
		}
		authentication.credentials = &credentials
	}
	return
}

//        BindResponse ::= [APPLICATION 1] SEQUENCE {
//             COMPONENTS OF LDAPResult,
//             serverSaslCreds    [7] OCTET STRING OPTIONAL }
func ReadBindResponse(bytes *Bytes) (bindresponse BindResponse, err error) {
	err = bytes.ParseSubBytes(classApplication, TagBindResponse, bindresponse.ReadBindResponseComponents)
	return
}

func (bindresponse *BindResponse) ReadBindResponseComponents(bytes *Bytes) (err error) {
	bindresponse.ReadLDAPResultComponents(bytes)
	if bytes.HasMoreData() {
		var serverSaslCreds OCTETSTRING
		serverSaslCreds, err = ReadTaggedOCTETSTRING(bytes, classContextSpecific, TagBindResponseServerSaslCreds)
		bindresponse.serverSaslCreds = &serverSaslCreds
	}
	return
}

//
//        UnbindRequest ::= [APPLICATION 2] NULL
func ReadUnbindRequest(bytes *Bytes) (unbindrequest UnbindRequest, err error) {
	var tagAndLength tagAndLength
	tagAndLength, err = bytes.ParseTagAndLength()
	if err != nil {
		return
	}
	err = tagAndLength.Expect(classApplication, TagUnbindRequest, isNotCompound)
	if err != nil {
		return
	}
	if tagAndLength.length != 0 {
		err = LdapError{"Unbind request: expecting NULL"}
	}
	return
}

//
//        SearchRequest ::= [APPLICATION 3] SEQUENCE {
//             baseObject      LDAPDN,
//             scope           ENUMERATED {
//                  baseObject              (0),
//                  singleLevel             (1),
//                  wholeSubtree            (2),
//                  ...  },
//             derefAliases    ENUMERATED {
//                  neverDerefAliases       (0),
//                  derefInSearching        (1),
//                  derefFindingBaseObj     (2),
//                  derefAlways             (3) },
//             sizeLimit       INTEGER (0 ..  maxInt),
//             timeLimit       INTEGER (0 ..  maxInt),
//             typesOnly       BOOLEAN,
//             filter          Filter,
//             attributes      AttributeSelection }
func ReadSearchRequest(bytes *Bytes) (searchrequest SearchRequest, err error) {
	err = bytes.ParseSubBytes(classApplication, TagSearchRequest, searchrequest.ReadSearchRequestComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("ReadSearchRequest: %s", err.Error())}
	}
	return
}
func (searchrequest *SearchRequest) ReadSearchRequestComponents(bytes *Bytes) (err error) {
	searchrequest.baseObject, err = ReadLDAPDN(bytes)
	if err != nil {
		return
	}
	searchrequest.scope, err = ReadENUMERATED(bytes, EnumeratedSearchRequestScope)
	if err != nil {
		return
	}
	searchrequest.derefAliases, err = ReadENUMERATED(bytes, EnumeratedSearchRequestDerefAliases)
	if err != nil {
		return
	}
	searchrequest.sizeLimit, err = ReadPositiveINTEGER(bytes)
	if err != nil {
		return
	}
	searchrequest.timeLimit, err = ReadPositiveINTEGER(bytes)
	if err != nil {
		return
	}
	searchrequest.typesOnly, err = ReadBOOLEAN(bytes)
	if err != nil {
		return
	}
	searchrequest.filter, err = ReadFilter(bytes)
	if err != nil {
		return
	}
	searchrequest.attributes, err = ReadAttributeSelection(bytes)
	if err != nil {
		return
	}
	return
}

//
//        AttributeSelection ::= SEQUENCE OF selector LDAPString
//                       -- The LDAPString is constrained to
//                       -- <attributeSelector> in Section 4.5.1.8
func ReadAttributeSelection(bytes *Bytes) (attributeSelection AttributeSelection, err error) {
	err = bytes.ParseSubBytes(classUniversal, tagSequence, attributeSelection.ReadAttributeSelectionComponents)
	return
}
func (attributeSelection *AttributeSelection) ReadAttributeSelectionComponents(bytes *Bytes) (err error) {
	for bytes.HasMoreData() {
		var ldapstring LDAPString
		ldapstring, err = ReadLDAPString(bytes)
		if err != nil {
			return
		}
		*attributeSelection = append(*attributeSelection, ldapstring)
	}
	return
}

//
//        Filter ::= CHOICE {
//             and             [0] SET SIZE (1..MAX) OF filter Filter,
//             or              [1] SET SIZE (1..MAX) OF filter Filter,
//             not             [2] Filter,
//             equalityMatch   [3] AttributeValueAssertion,
//
//
//
//Sermersheim                 Standards Track                    [Page 57]
//
//
//RFC 4511                         LDAPv3                        June 2006
//
//
//             substrings      [4] SubstringFilter,
//             greaterOrEqual  [5] AttributeValueAssertion,
//             lessOrEqual     [6] AttributeValueAssertion,
//             present         [7] AttributeDescription,
//             approxMatch     [8] AttributeValueAssertion,
//             extensibleMatch [9] MatchingRuleAssertion,
//             ...  }
func ReadFilter(bytes *Bytes) (filter Filter, err error) {
	var tagAndLength tagAndLength
	tagAndLength, err = bytes.PreviewTagAndLength()
	if err != nil {
		return
	}
	err = tagAndLength.ExpectClass(classContextSpecific)
	if err != nil {
		return
	}
	switch tagAndLength.tag {
	case TagFilterAnd:
		filter, err = ReadFilterAnd(bytes)
	case TagFilterOr:
		filter, err = ReadFilterOr(bytes)
	case TagFilterNot:
		filter, err = ReadFilterNot(bytes)
	case TagFilterEqualityMatch:
		filter, err = ReadFilterEqualityMatch(bytes)
	case TagFilterSubstrings:
		filter, err = ReadFilterSubstrings(bytes)
	case TagFilterGreaterOrEqual:
		filter, err = ReadFilterGreaterOrEqual(bytes)
	case TagFilterLessOrEqual:
		filter, err = ReadFilterLessOrEqual(bytes)
	case TagFilterPresent:
		filter, err = ReadFilterPresent(bytes)
	case TagFilterApproxMatch:
		filter, err = ReadFilterApproxMatch(bytes)
	case TagFilterExtensibleMatch:
		filter, err = ReadFilterExtensibleMatch(bytes)
	default:
		err = LdapError{fmt.Sprintf("Invalid tag value for filter: %d.", tagAndLength.GetTag())}
	}
	if err != nil {
		err = LdapError{fmt.Sprintf("ReadFilter: %s.", err.Error())}
	}
	return
}

//             and             [0] SET SIZE (1..MAX) OF filter Filter,
func ReadFilterAnd(bytes *Bytes) (filterand FilterAnd, err error) {
	err = bytes.ParseSubBytes(classContextSpecific, TagFilterAnd, filterand.ReadFilterAndComponents)
	return
}
func (filterand *FilterAnd) ReadFilterAndComponents(bytes *Bytes) (err error) {
	for bytes.HasMoreData() {
		var filter Filter
		filter, err = ReadFilter(bytes)
		if err != nil {
			return
		}
		*filterand = append(*filterand, filter)
	}
	if len(*filterand) == 0 {
		err = LdapError{"ReadFilterAnd: expecting at least one Filter"}
	}
	return
}

//             or              [1] SET SIZE (1..MAX) OF filter Filter,
func ReadFilterOr(bytes *Bytes) (filteror FilterOr, err error) {
	err = bytes.ParseSubBytes(classContextSpecific, TagFilterOr, filteror.ReadFilterOrComponents)
	return
}

func (filteror *FilterOr) ReadFilterOrComponents(bytes *Bytes) (err error) {
	for bytes.HasMoreData() {
		var filter Filter
		filter, err = ReadFilter(bytes)
		if err != nil {
			return
		}
		*filteror = append(*filteror, filter)
	}
	if len(*filteror) == 0 {
		err = LdapError{"ReadFilterOr: expecting at least one Filter"}
	}
	return
}

//             not             [2] Filter,
func ReadFilterNot(bytes *Bytes) (filternot FilterNot, err error) {
	err = bytes.ParseSubBytes(classContextSpecific, TagFilterNot, filternot.ReadFilterNotComponents)
	return
}

func (filternot *FilterNot) ReadFilterNotComponents(bytes *Bytes) (err error) {
	var tagAndLength tagAndLength
	tagAndLength, err = bytes.ParseTagAndLength()
	if err != nil {
		return
	}
	err = tagAndLength.ExpectTag(TagFilterNot)
	if err != nil {
		return
	}
	filternot.Filter, err =  ReadFilter(bytes)
	if err != nil {
		return
	}
	return
}

//             equalityMatch   [3] AttributeValueAssertion,
func ReadFilterEqualityMatch(bytes *Bytes) (ret FilterEqualityMatch, err error) {
	var attributevalueassertion AttributeValueAssertion
	attributevalueassertion, err = ReadTaggedAttributeValueAssertion(bytes, classContextSpecific, TagFilterEqualityMatch)
	if err != nil {
		return
	}
	ret = FilterEqualityMatch(attributevalueassertion)
	return
}

//             substrings      [4] SubstringFilter,
func ReadFilterSubstrings(bytes *Bytes) (filtersubstrings FilterSubstrings, err error) {
	var substringfilter SubstringFilter
	substringfilter, err = ReadTaggedSubstringFilter(bytes, classContextSpecific, TagFilterSubstrings)
	if err != nil {
		return
	}
	filtersubstrings = FilterSubstrings(substringfilter)
	return
}

//             greaterOrEqual  [5] AttributeValueAssertion,
func ReadFilterGreaterOrEqual(bytes *Bytes) (ret FilterGreaterOrEqual, err error) {
	var attributevalueassertion AttributeValueAssertion
	attributevalueassertion, err = ReadTaggedAttributeValueAssertion(bytes, classContextSpecific, TagFilterGreaterOrEqual)
	if err != nil {
		return
	}
	ret = FilterGreaterOrEqual(attributevalueassertion)
	return
}

//             lessOrEqual     [6] AttributeValueAssertion,
func ReadFilterLessOrEqual(bytes *Bytes) (ret FilterLessOrEqual, err error) {
	var attributevalueassertion AttributeValueAssertion
	attributevalueassertion, err = ReadTaggedAttributeValueAssertion(bytes, classContextSpecific, TagFilterLessOrEqual)
	if err != nil {
		return
	}
	ret = FilterLessOrEqual(attributevalueassertion)
	return
}

//             present         [7] AttributeDescription,
func ReadFilterPresent(bytes *Bytes) (ret FilterPresent, err error) {
	var attributedescription AttributeDescription
	attributedescription, err = ReadTaggedAttributeDescription(bytes, classContextSpecific, TagFilterPresent)
	if err != nil {
		return ret, LdapError{fmt.Sprintf("ReadFilterPresent: %s", err.Error())} 
	}
	ret = FilterPresent(attributedescription)
	return
}

//             approxMatch     [8] AttributeValueAssertion,
func ReadFilterApproxMatch(bytes *Bytes) (ret FilterApproxMatch, err error) {
	var attributevalueassertion AttributeValueAssertion
	attributevalueassertion, err = ReadTaggedAttributeValueAssertion(bytes, classContextSpecific, TagFilterApproxMatch)
	if err != nil {
		return
	}
	ret = FilterApproxMatch(attributevalueassertion)
	return
}

//             extensibleMatch [9] MatchingRuleAssertion,
func ReadFilterExtensibleMatch(bytes *Bytes) (filterextensiblematch FilterExtensibleMatch, err error) {
	var matchingruleassertion MatchingRuleAssertion
	matchingruleassertion, err = ReadTaggedMatchingRuleAssertion(bytes, classContextSpecific, TagFilterExtensibleMatch)
	if err != nil {
		return
	}
	filterextensiblematch = FilterExtensibleMatch(matchingruleassertion)
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
func ReadTaggedSubstringFilter(bytes *Bytes, class int, tag int) (substringfilter SubstringFilter, err error) {
	err = bytes.ParseSubBytes(class, tag, substringfilter.ReadSubstringFilterComponents)
	return
}
func (substringfilter SubstringFilter) ReadSubstringFilterComponents(bytes *Bytes) (err error) {
	substringfilter.type_, err = ReadAttributeDescription(bytes)
	if err != nil {
		return
	}
	substringfilter.substrings, err = ReadSubstringFilterSubstrings(bytes)
	if err != nil {
		return
	}
	return
}

func ReadSubstringFilterSubstrings(bytes *Bytes) (substrings SubstringFilterSubstrings, err error) {
	err = bytes.ParseSubBytes(classUniversal, tagSequence, substrings.ReadSubstringFilterSubstringsComponents)
	return
}

func (substrings *SubstringFilterSubstrings) ReadSubstringFilterSubstringsComponents(bytes *Bytes) (err error) {
	var foundInitial = 0
	var foundFinal = 0
	for bytes.HasMoreData() {
		var tagAndLength tagAndLength
		tagAndLength, err = bytes.PreviewTagAndLength()
		if err != nil {
			return
		}
		var assertionvalue AssertionValue
		switch tagAndLength.tag {
			case TagSubstringInitial:
				foundInitial++
				if foundInitial > 1 {
					return LdapError{"ReadSubstring: initial can occur at most once"}
				}
				assertionvalue, err = ReadTaggedAssertionValue(bytes, classContextSpecific, TagSubstringInitial)
				if err != nil {
					return
				}
				*substrings = append(*substrings, SubstringInitial(assertionvalue))
			case TagSubstringAny:
				assertionvalue, err = ReadTaggedAssertionValue(bytes, classContextSpecific, TagSubstringAny)
				if err != nil {
					return
				}
				*substrings = append(*substrings, SubstringAny(assertionvalue))
			case TagSubstringFinal:
				foundFinal++
				if foundFinal > 1 {
					return LdapError{"ReadSubstring: final can occur at most once"}
				}
				assertionvalue, err = ReadTaggedAssertionValue(bytes, classContextSpecific, TagSubstringFinal)
				if err != nil {
					return
				}
				*substrings = append(*substrings, SubstringFinal(assertionvalue))
			default:
				return LdapError{fmt.Sprintf("ReadSubstring: invalid tag %d", tagAndLength.tag)}
		}
	}
	if len(*substrings) == 0 {
		err = LdapError{"ReadSubstringFilterSubstrings: expecting at least one substring"}
	}
	return
}

//
//        MatchingRuleAssertion ::= SEQUENCE {
//             matchingRule    [1] MatchingRuleId OPTIONAL,
//             type            [2] AttributeDescription OPTIONAL,
//             matchValue      [3] AssertionValue,
//             dnAttributes    [4] BOOLEAN DEFAULT FALSE }
func ReadTaggedMatchingRuleAssertion(bytes *Bytes, class int, tag int) (ret MatchingRuleAssertion, err error) {
	err = bytes.ParseSubBytes(class, tag, ret.ReadMatchingRuleAssertionComponents)
	return
}
func (matchingruleassertion MatchingRuleAssertion) ReadMatchingRuleAssertionComponents(bytes *Bytes) (err error) {
	err = matchingruleassertion.ReadMatchingRule(bytes)
	if err != nil {
		return LdapError{fmt.Sprintf("ReadMatchingRuleAssertionComponents: %s", err.Error())}
	}
	err = matchingruleassertion.ReadType(bytes)
	if err != nil {
		return LdapError{fmt.Sprintf("ReadMatchingRuleAssertionComponents: %s", err.Error())}
	}
	matchingruleassertion.matchValue, err = ReadTaggedAssertionValue(bytes, classContextSpecific, TagMatchingRuleAssertionMatchValue)
	if err != nil {
		return LdapError{fmt.Sprintf("ReadMatchingRuleAssertionComponents: %s", err.Error())}
	}
	matchingruleassertion.dnAttributes, err = ReadTaggedBOOLEAN(bytes, classContextSpecific, TagMatchingRuleAssertionDnAttributes)
	if err != nil {
		return LdapError{fmt.Sprintf("ReadMatchingRuleAssertionComponents: %s", err.Error())}
	}
	return
}
func (matchingruleassertion MatchingRuleAssertion) ReadMatchingRule(bytes *Bytes) (err error) {
	var tagAndLength tagAndLength
	tagAndLength, err = bytes.PreviewTagAndLength()
	if err != nil {
		return LdapError{fmt.Sprintf("ReadMatchingRuleAssertionMatchingRule: %s", err.Error())}
	}
	if tagAndLength.tag == TagMatchingRuleAssertionMatchingRule {
		var matchingRule MatchingRuleId
		matchingRule, err = ReadTaggedMatchingRuleId(bytes, classContextSpecific, TagMatchingRuleAssertionMatchingRule)
		if err != nil {
			return LdapError{fmt.Sprintf("ReadMatchingRuleAssertionMatchingRule: %s", err.Error())}
		}
		matchingruleassertion.matchingRule = &matchingRule
	}
	return
}
func (matchingruleassertion MatchingRuleAssertion) ReadType(bytes *Bytes) (err error) {
	var tagAndLength tagAndLength
	tagAndLength, err = bytes.PreviewTagAndLength()
	if err != nil {
		return LdapError{fmt.Sprintf("ReadMatchingRuleAssertionType: %s", err.Error())}
	}
	if tagAndLength.tag == TagMatchingRuleAssertionType {
		var attributedescription AttributeDescription
		attributedescription, err = ReadTaggedAttributeDescription(bytes, classContextSpecific, TagMatchingRuleAssertionType)
		if err != nil {
			return LdapError{fmt.Sprintf("ReadMatchingRuleAssertionType: %s", err.Error())}
		}
		matchingruleassertion.type_ = &attributedescription
	}
	return
}

//
//        SearchResultEntry ::= [APPLICATION 4] SEQUENCE {
//             objectName      LDAPDN,
//             attributes      PartialAttributeList }
func ReadSearchResultEntry(bytes *Bytes) (searchresultentry SearchResultEntry, err error){
	err = bytes.ParseSubBytes(classApplication, TagSearchResultEntry, searchresultentry.ReadSearchResultEntryComponents)
	return
}
func (searchresultentry *SearchResultEntry) ReadSearchResultEntryComponents(bytes *Bytes) (err error){
	searchresultentry.objectName, err = ReadLDAPDN(bytes)
	if err != nil {
		return
	}
	searchresultentry.attributes, err = ReadPartialAttributeList(bytes)
	if err != nil {
		return
	}
	return
}
//
//        PartialAttributeList ::= SEQUENCE OF
//                             partialAttribute PartialAttribute
func ReadPartialAttributeList(bytes *Bytes) (ret PartialAttributeList, err error){
	err = bytes.ParseSubBytes(classUniversal, tagSequence, ret.ReadPartialAttributeListComponents)
	return
}
func (partialattributelist *PartialAttributeList) ReadPartialAttributeListComponents(bytes *Bytes) (err error){
	for bytes.HasMoreData() {
		var partialattribute PartialAttribute
		partialattribute, err = ReadPartialAttribute(bytes)
		if err != nil {
			return
		}
		*partialattributelist = append(*partialattributelist, partialattribute)
	}
	return
}

//
//        SearchResultReference ::= [APPLICATION 19] SEQUENCE
//                                  SIZE (1..MAX) OF uri URI
//
//        SearchResultDone ::= [APPLICATION 5] LDAPResult
func ReadSearchResultDone(bytes *Bytes) (ret SearchResultDone, err error){
	var ldapresult LDAPResult
	ldapresult, err = ReadTaggedLDAPResult(bytes, classApplication, TagSearchResultDone)
	if err != nil {
		return
	}
	ret = SearchResultDone(ldapresult)
	return
}
//
//        ModifyRequest ::= [APPLICATION 6] SEQUENCE {
//             object          LDAPDN,
//             changes         SEQUENCE OF change SEQUENCE {
//                  operation       ENUMERATED {
//                       add     (0),
//                       delete  (1),
//                       replace (2),
//                       ...  },
//                  modification    PartialAttribute } }
//
//        ModifyResponse ::= [APPLICATION 7] LDAPResult
//
//
//
//
//
//
//Sermersheim                 Standards Track                    [Page 58]
//
//
//RFC 4511                         LDAPv3                        June 2006
//
//
//        AddRequest ::= [APPLICATION 8] SEQUENCE {
//             entry           LDAPDN,
//             attributes      AttributeList }
//
//        AttributeList ::= SEQUENCE OF attribute Attribute
//
//        AddResponse ::= [APPLICATION 9] LDAPResult
//
//        DelRequest ::= [APPLICATION 10] LDAPDN
//
//        DelResponse ::= [APPLICATION 11] LDAPResult
//
//        ModifyDNRequest ::= [APPLICATION 12] SEQUENCE {
//             entry           LDAPDN,
//             newrdn          RelativeLDAPDN,
//             deleteoldrdn    BOOLEAN,
//             newSuperior     [0] LDAPDN OPTIONAL }
//
//        ModifyDNResponse ::= [APPLICATION 13] LDAPResult
//
//        CompareRequest ::= [APPLICATION 14] SEQUENCE {
//             entry           LDAPDN,
//             ava             AttributeValueAssertion }
//
//        CompareResponse ::= [APPLICATION 15] LDAPResult
//
//        AbandonRequest ::= [APPLICATION 16] MessageID
//
//        ExtendedRequest ::= [APPLICATION 23] SEQUENCE {
//             requestName      [0] LDAPOID,
//             requestValue     [1] OCTET STRING OPTIONAL }
//
//        ExtendedResponse ::= [APPLICATION 24] SEQUENCE {
//             COMPONENTS OF LDAPResult,
//             responseName     [10] LDAPOID OPTIONAL,
//             responseValue    [11] OCTET STRING OPTIONAL }
//
//        IntermediateResponse ::= [APPLICATION 25] SEQUENCE {
//             responseName     [0] LDAPOID OPTIONAL,
//             responseValue    [1] OCTET STRING OPTIONAL }
//
//        END
//
