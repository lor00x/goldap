package message

import (
	"errors"
	"fmt"
)

type LdapError struct {
	Msg string
}

func (e LdapError) Error() string { return e.Msg }

func readBOOLEAN(bytes Bytes) (ret BOOLEAN, err error) {
	return readTaggedBOOLEAN(bytes, classUniversal, tagBoolean)
}
func readTaggedBOOLEAN(bytes Bytes, class int, tag int) (ret BOOLEAN, err error) {
	tagAndLength, err := bytes.ParseTagAndLength()
	if err != nil {
		return
	}
	err = tagAndLength.Expect(class, tag, isNotCompound)
	if err != nil {
		return
	}
	var boolean bool
	boolean, err = bytes.ParseBool(tagAndLength.Length)
	return BOOLEAN(boolean), err
}

func readINTEGER(bytes Bytes) (ret INTEGER, err error) {
	return readTaggedINTEGER(bytes, classUniversal, tagInteger)
}
func readTaggedINTEGER(bytes Bytes, class int, tag int) (ret INTEGER, err error) {
	tagAndLength, err := bytes.ParseTagAndLength()
	if err != nil {
		return
	}
	err = tagAndLength.Expect(class, tag, isNotCompound)
	if err != nil {
		return
	}
	var integer int32
	integer, err = bytes.ParseInt32(tagAndLength.Length)
	return INTEGER(integer), err
}

func readPositiveINTEGER(bytes Bytes) (ret INTEGER, err error) {
	return readTaggedPositiveINTEGER(bytes, classUniversal, tagInteger)
}
func readTaggedPositiveINTEGER(bytes Bytes, class int, tag int) (ret INTEGER, err error) {
	ret, err = readTaggedINTEGER(bytes, class, tag)
	if err != nil {
		return
	}
	if !(ret >= 0 && ret <= maxInt) {
		err = LdapError{fmt.Sprintf("Invalid INTEGER value %d ! Expected value between 0 and %d", ret, maxInt)}
	}
	return
}

func readENUMERATED(bytes Bytes, allowedValues map[ENUMERATED]string) (ret ENUMERATED, err error) {
	tagAndLength, err := bytes.ParseTagAndLength()
	if err != nil {
		return ret, LdapError{fmt.Sprintf("readENUMERATED: %s", err.Error())}
	}
	err = tagAndLength.Expect(classUniversal, tagEnum, isNotCompound)
	if err != nil {
		return ret, LdapError{fmt.Sprintf("readENUMERATED: %s", err.Error())}
	}
	var integer int32
	integer, err = bytes.ParseInt32(tagAndLength.Length)
	if err != nil {
		return ret, LdapError{fmt.Sprintf("readENUMERATED: %s", err.Error())}
	}
	ret = ENUMERATED(integer)
	if _, ok := allowedValues[ret]; !ok {
		return ret, LdapError{fmt.Sprintf("readENUMERATED: Invalid ENUMERATED VALUE %d", ret)}
	}
	return
}

func readUTF8STRING(bytes Bytes) (ret UTF8STRING, err error) {
	return readTaggedUTF8STRING(bytes, classUniversal, tagUTF8String)
}
func readTaggedUTF8STRING(bytes Bytes, class int, tag int) (ret UTF8STRING, err error) {
	tagAndLength, err := bytes.ParseTagAndLength()
	if err != nil {
		return ret, errors.New(fmt.Sprintf("readTaggedUTF8STRING: %s", err.Error()))
	}
	err = tagAndLength.Expect(class, tag, isNotCompound)
	if err != nil {
		return ret, errors.New(fmt.Sprintf("readTaggedUTF8STRING: %s", err.Error()))
	}
	var utf8string string
	utf8string, err = bytes.ParseUTF8String(tagAndLength.Length)
	if err != nil {
		return ret, errors.New(fmt.Sprintf("readTaggedUTF8STRING: %s", err.Error()))
	}
	return UTF8STRING(utf8string), err
}

func readOCTETSTRING(bytes Bytes) (ret OCTETSTRING, err error) {
	return readTaggedOCTETSTRING(bytes, classUniversal, tagOctetString)
}
func readTaggedOCTETSTRING(bytes Bytes, class int, tag int) (ret OCTETSTRING, err error) {
	tagAndLength, err := bytes.ParseTagAndLength()
	if err != nil {
		return
	}
	err = tagAndLength.Expect(class, tag, isNotCompound)
	if err != nil {
		return
	}
	var octetstring []byte
	octetstring, err = bytes.ParseOCTETSTRING(tagAndLength.Length)
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

func ReadLDAPMessage(bytes Bytes) (message LDAPMessage, err error) {
	err = bytes.ReadSubBytes(classUniversal, tagSequence, message.readLDAPMessageComponents)
	if err != nil {
		//panic(err.Error())
		err = errors.New(fmt.Sprintf("readLDAPMessage: %s", err.Error()))
	}
	return
}
func (message *LDAPMessage) readLDAPMessageComponents(bytes Bytes) (err error) {
	message.messageID, err = readMessageID(bytes)
	if err != nil {
		return
	}
	message.protocolOp, err = readProtocolOp(bytes)
	if err != nil {
		return
	}
	if bytes.HasMoreData() {
		var controls Controls
		controls, err = readControls(bytes)
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
func readMessageID(bytes Bytes) (ret MessageID, err error) {
	return readTaggedMessageID(bytes, classUniversal, tagInteger)
}
func readTaggedMessageID(bytes Bytes, class int, tag int) (ret MessageID, err error) {
	var integer INTEGER
	integer, err = readTaggedPositiveINTEGER(bytes, class, tag)
	if err != nil {
		err = errors.New(fmt.Sprintf("readMessageID: %s", err.Error()))
		return
	}
	return MessageID(integer), err
}
func readProtocolOp(bytes Bytes) (ret ProtocolOp, err error) {
	tagAndLength, err := bytes.PreviewTagAndLength()
	if err != nil {
		err = errors.New(fmt.Sprintf("readProtocolOp: %s", err.Error()))
		return
	}
	switch tagAndLength.Tag {
	case TagBindRequest:
		ret, err = readBindRequest(bytes)
	case TagBindResponse:
		ret, err = readBindResponse(bytes)
	case TagUnbindRequest:
		ret, err = readUnbindRequest(bytes)
	case TagSearchRequest:
		ret, err = readSearchRequest(bytes)
	case TagSearchResultEntry:
		ret, err = readSearchResultEntry(bytes)
	case TagSearchResultDone:
		ret, err = readSearchResultDone(bytes)
	case TagSearchResultReference:
		ret, err = readSearchResultReference(bytes)
	case TagModifyRequest:
		ret, err = readModifyRequest(bytes)
	case TagModifyResponse:
		ret, err = readModifyResponse(bytes)
	case TagAddRequest:
		ret, err = readAddRequest(bytes)
	case TagAddResponse:
		ret, err = readAddResponse(bytes)
	case TagDelRequest:
		ret, err = readDelRequest(bytes)
	case TagDelResponse:
		ret, err = readDelResponse(bytes)
	case TagModifyDNRequest:
		ret, err = readModifyDNRequest(bytes)
	case TagModifyDNResponse:
		ret, err = readModifyDNResponse(bytes)
	case TagCompareRequest:
		ret, err = readCompareRequest(bytes)
	case TagCompareResponse:
		ret, err = readCompareResponse(bytes)
	case TagAbandonRequest:
		ret, err = readAbandonRequest(bytes)
	case TagExtendedRequest:
		ret, err = readExtendedRequest(bytes)
	case TagExtendedResponse:
		ret, err = readExtendedResponse(bytes)
	case TagIntermediateResponse:
		ret, err = readIntermediateResponse(bytes)
	default:
		err = LdapError{fmt.Sprintf("Invalid tag value for protocolOp. Got %d.", tagAndLength.Tag)}
	}
	if err != nil {
		err = errors.New(fmt.Sprintf("readProtocolOp: %s", err.Error()))
	}
	return
}

//        LDAPString ::= OCTET STRING -- UTF-8 encoded,
//                                    -- [ISO10646] characters
func readLDAPString(bytes Bytes) (ldapstring LDAPString, err error) {
	return readTaggedLDAPString(bytes, classUniversal, tagOctetString)
}
func readTaggedLDAPString(bytes Bytes, class int, tag int) (ldapstring LDAPString, err error) {
	var utf8string UTF8STRING
	utf8string, err = readTaggedUTF8STRING(bytes, class, tag)
	if err != nil {
		err = errors.New(fmt.Sprintf("readTaggedLDAPString: %s", err.Error()))
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
func readLDAPOID(bytes Bytes) (ret LDAPOID, err error) {
	return readTaggedLDAPOID(bytes, classUniversal, tagOctetString)
}
func readTaggedLDAPOID(bytes Bytes, class int, tag int) (ret LDAPOID, err error) {
	var octetstring OCTETSTRING
	octetstring, err = readTaggedOCTETSTRING(bytes, class, tag)
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
func readLDAPDN(bytes Bytes) (ret LDAPDN, err error) {
	var str LDAPString
	str, err = readLDAPString(bytes)
	if err != nil {
		return
	}
	ret = LDAPDN(str)
	return
}
func readTaggedLDAPDN(bytes Bytes, class int, tag int) (ret LDAPDN, err error) {
	var ldapstring LDAPString
	ldapstring, err = readTaggedLDAPString(bytes, class, tag)
	if err != nil {
		err = errors.New(fmt.Sprintf("readLDAPDN: %s", err.Error()))
		return
	}
	// @TODO: check RFC4514
	ret = LDAPDN(ldapstring)
	return
}

//
//        RelativeLDAPDN ::= LDAPString -- Constrained to <name-component>
//                                      -- [RFC4514]
func readRelativeLDAPDN(bytes Bytes) (ret RelativeLDAPDN, err error) {
	var ldapstring LDAPString
	ldapstring, err = readLDAPString(bytes)
	// @TODO: check RFC4514
	ret = RelativeLDAPDN(ldapstring)
	return
}

//
//        AttributeDescription ::= LDAPString
//                                -- Constrained to <attributedescription>
//                                -- [RFC4512]
func readAttributeDescription(bytes Bytes) (ret AttributeDescription, err error) {
	var ldapstring LDAPString
	ldapstring, err = readLDAPString(bytes)
	// @TODO: check RFC4512
	ret = AttributeDescription(ldapstring)
	return
}
func readTaggedAttributeDescription(bytes Bytes, class int, tag int) (ret AttributeDescription, err error) {
	var ldapstring LDAPString
	ldapstring, err = readTaggedLDAPString(bytes, class, tag)
	// @TODO: check RFC4512
	if err != nil {
		err = errors.New(fmt.Sprintf("readTaggedAttributeDescription: %s", err.Error()))
		return
	}
	ret = AttributeDescription(ldapstring)
	return
}

//
//        AttributeValue ::= OCTET STRING
func readAttributeValue(bytes Bytes) (ret AttributeValue, err error) {
	var octetstring OCTETSTRING
	octetstring, err = readOCTETSTRING(bytes)
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
func readAttributeValueAssertion(bytes Bytes) (ret AttributeValueAssertion, err error) {
	return readTaggedAttributeValueAssertion(bytes, classUniversal, tagSequence)
}
func readTaggedAttributeValueAssertion(bytes Bytes, class int, tag int) (ret AttributeValueAssertion, err error) {
	err = bytes.ReadSubBytes(class, tag, ret.readAttributeValueAssertionComponents)
	return
}

func (attributevalueassertion *AttributeValueAssertion) readAttributeValueAssertionComponents(bytes Bytes) (err error) {
	attributevalueassertion.attributeDesc, err = readAttributeDescription(bytes)
	if err != nil {
		return
	}
	attributevalueassertion.assertionValue, err = readAssertionValue(bytes)
	if err != nil {
		return
	}
	return
}

//
//        AssertionValue ::= OCTET STRING
func readAssertionValue(bytes Bytes) (assertionvalue AssertionValue, err error) {
	return readTaggedAssertionValue(bytes, classUniversal, tagOctetString)
}
func readTaggedAssertionValue(bytes Bytes, class int, tag int) (assertionvalue AssertionValue, err error) {
	var octetstring OCTETSTRING
	octetstring, err = readTaggedOCTETSTRING(bytes, class, tag)
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
func readPartialAttribute(bytes Bytes) (ret PartialAttribute, err error) {
	ret = PartialAttribute{vals: make([]AttributeValue, 0, 10)}
	err = bytes.ReadSubBytes(classUniversal, tagSequence, ret.readPartialAttributeComponents)
	return
}

func (partialattribute *PartialAttribute) readPartialAttributeComponents(bytes Bytes) (err error) {
	partialattribute.type_, err = readAttributeDescription(bytes)
	if err != nil {
		return
	}
	err = bytes.ReadSubBytes(classUniversal, tagSet, partialattribute.readPartialAttributeValsComponents)
	if err != nil {
		return
	}
	return
}
func (partialattribute *PartialAttribute) readPartialAttributeValsComponents(bytes Bytes) (err error) {
	for bytes.HasMoreData() {
		var attributevalue AttributeValue
		attributevalue, err = readAttributeValue(bytes)
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
func readAttribute(bytes Bytes) (ret Attribute, err error) {
	var par PartialAttribute
	par, err = readPartialAttribute(bytes)
	if err != nil {
		return
	}
	if len(par.vals) == 0 {
		err = LdapError{"readAttribute: expecting at least one value"}
		return
	}
	ret = Attribute(par)
	return

}

//
//        MatchingRuleId ::= LDAPString
func readTaggedMatchingRuleId(bytes Bytes, class int, tag int) (matchingruleid MatchingRuleId, err error) {
	var ldapstring LDAPString
	ldapstring, err = readTaggedLDAPString(bytes, class, tag)
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
func readTaggedLDAPResult(bytes Bytes, class int, tag int) (ret LDAPResult, err error) {
	err = bytes.ReadSubBytes(class, tag, ret.readLDAPResultComponents)
	if err != nil {
		err = fmt.Errorf("readLDAPResult: %s", err.Error())
	}
	return
}
func readLDAPResult(bytes Bytes) (ldapresult LDAPResult, err error) {
	return readTaggedLDAPResult(bytes, classUniversal, tagSequence)
}
func (ldapresult *LDAPResult) readLDAPResultComponents(bytes Bytes) (err error) {
	ldapresult.resultCode, err = readENUMERATED(bytes, EnumeratedLDAPResultCode)
	if err != nil {
		return
	}
	ldapresult.matchedDN, err = readLDAPDN(bytes)
	if err != nil {
		return
	}
	ldapresult.diagnosticMessage, err = readLDAPString(bytes)
	if err != nil {
		return
	}
	if bytes.HasMoreData() {
		var referral Referral
		referral, err = readReferral(bytes)
		if err != nil {
			return
		}
		ldapresult.referral = &referral
	}
	return
}

//
//        Referral ::= SEQUENCE SIZE (1..MAX) OF uri URI
func readReferral(bytes Bytes) (referral Referral, err error) {
	err = bytes.ReadSubBytes(classUniversal, tagSequence, referral.readReferralComponents)
	return
}
func (referral *Referral) readReferralComponents(bytes Bytes) (err error) {
	for bytes.HasMoreData() {
		var uri URI
		uri, err = readURI(bytes)
		if err != nil {
			return
		}
		*referral = append(*referral, uri)
	}
	if len(*referral) == 0 {
		return LdapError{"readReferral: expecting at least one URI"}
	}
	return
}

//
//        URI ::= LDAPString     -- limited to characters permitted in
//                               -- URIs
func readURI(bytes Bytes) (uri URI, err error) {
	var ldapstring LDAPString
	ldapstring, err = readLDAPString(bytes)
	// @TODO: check permitted chars in URI
	if err != nil {
		return
	}
	uri = URI(ldapstring)
	return
}

//
//        Controls ::= SEQUENCE OF control Control
func readControls(bytes Bytes) (controls Controls, err error) {
	err = bytes.ReadSubBytes(classUniversal, tagSequence, controls.readControlsComponents)
	return
}
func (controls *Controls) readControlsComponents(bytes Bytes) (err error) {
	for bytes.HasMoreData() {
		var control Control
		control, err = readControl(bytes)
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
func readControl(bytes Bytes) (control Control, err error) {
	err = bytes.ReadSubBytes(classUniversal, tagSequence, control.readControlComponents)
	return
}
func (control *Control) readControlComponents(bytes Bytes) (err error) {
	control.controlType, err = readLDAPOID(bytes)
	if err != nil {
		return
	}
	control.criticality, err = readBOOLEAN(bytes)
	if err != nil {
		return
	}
	if bytes.HasMoreData() {
		var octetstring OCTETSTRING
		octetstring, err = readOCTETSTRING(bytes)
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
func readBindRequest(bytes Bytes) (bindrequest BindRequest, err error) {
	err = bytes.ReadSubBytes(classApplication, TagBindRequest, bindrequest.readBindRequestComponents)
	if err != nil {
		err = errors.New(fmt.Sprintf("readBindRequest: %s", err.Error()))
	}
	return
}
func (bindrequest *BindRequest) readBindRequestComponents(bytes Bytes) (err error) {
	bindrequest.version, err = readINTEGER(bytes)
	if !(bindrequest.version >= BindRequestVersionMin && bindrequest.version <= BindRequestVersionMax) {
		err = LdapError{fmt.Sprintf("Invalid version %d. Must be between %d and %d", bindrequest.version, BindRequestVersionMin, BindRequestVersionMax)}
		return
	}
	bindrequest.name, err = readLDAPDN(bytes)
	if err != nil {
		return
	}
	bindrequest.authentication, err = readAuthenticationChoice(bytes)
	return
}

//
//        AuthenticationChoice ::= CHOICE {
//             simple                  [0] OCTET STRING,
//                                     -- 1 and 2 reserved
//             sasl                    [3] SaslCredentials,
//             ...  }
func readAuthenticationChoice(bytes Bytes) (ret interface{}, err error) {
	tagAndLength, err := bytes.PreviewTagAndLength()
	if err != nil {
		err = errors.New(fmt.Sprintf("readAuthenticationChoice: %s", err.Error()))
		return
	}
	err = tagAndLength.ExpectClass(classContextSpecific)
	if err != nil {
		err = errors.New(fmt.Sprintf("readAuthenticationChoice: %s", err.Error()))
		return
	}
	switch tagAndLength.Tag {
	case TagAuthenticationChoiceSimple:
		ret, err = readTaggedOCTETSTRING(bytes, classContextSpecific, TagAuthenticationChoiceSimple)
	case TagAuthenticationChoiceSaslCredentials:
		ret, err = readSaslCredentials(bytes)
	default:
		err = LdapError{fmt.Sprintf("Invalid tag value for AuthenticationChoice. Got %d.", tagAndLength.Tag)}
	}
	if err != nil {
		err = errors.New(fmt.Sprintf("readAuthenticationChoice: %s", err.Error()))
	}
	return
}

//
//        SaslCredentials ::= SEQUENCE {
//             mechanism               LDAPString,
//             credentials             OCTET STRING OPTIONAL }
//
func readSaslCredentials(bytes Bytes) (authentication SaslCredentials, err error) {
	authentication = SaslCredentials{}
	err = bytes.ReadSubBytes(classContextSpecific, TagAuthenticationChoiceSaslCredentials, authentication.readSaslCredentialsComponents)
	return
}
func (authentication *SaslCredentials) readSaslCredentialsComponents(bytes Bytes) (err error) {
	authentication.mechanism, err = readLDAPString(bytes)
	if err != nil {
		return
	}
	if bytes.HasMoreData() {
		var credentials OCTETSTRING
		credentials, err = readOCTETSTRING(bytes)
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
func readBindResponse(bytes Bytes) (bindresponse BindResponse, err error) {
	err = bytes.ReadSubBytes(classApplication, TagBindResponse, bindresponse.readBindResponseComponents)
	return
}

func (bindresponse *BindResponse) readBindResponseComponents(bytes Bytes) (err error) {
	bindresponse.readLDAPResultComponents(bytes)
	if bytes.HasMoreData() {
		var serverSaslCreds OCTETSTRING
		serverSaslCreds, err = readTaggedOCTETSTRING(bytes, classContextSpecific, TagBindResponseServerSaslCreds)
		bindresponse.serverSaslCreds = &serverSaslCreds
	}
	return
}

//
//        UnbindRequest ::= [APPLICATION 2] NULL
func readUnbindRequest(bytes Bytes) (unbindrequest UnbindRequest, err error) {
	var tagAndLength TagAndLength
	tagAndLength, err = bytes.ParseTagAndLength()
	if err != nil {
		return
	}
	err = tagAndLength.Expect(classApplication, TagUnbindRequest, isNotCompound)
	if err != nil {
		return
	}
	if tagAndLength.Length != 0 {
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
func readSearchRequest(bytes Bytes) (searchrequest SearchRequest, err error) {
	err = bytes.ReadSubBytes(classApplication, TagSearchRequest, searchrequest.readSearchRequestComponents)
	if err != nil {
		err = LdapError{fmt.Sprintf("readSearchRequest: %s", err.Error())}
	}
	return
}
func (searchrequest *SearchRequest) readSearchRequestComponents(bytes Bytes) (err error) {
	searchrequest.baseObject, err = readLDAPDN(bytes)
	if err != nil {
		return
	}
	searchrequest.scope, err = readENUMERATED(bytes, EnumeratedSearchRequestScope)
	if err != nil {
		return
	}
	searchrequest.derefAliases, err = readENUMERATED(bytes, EnumeratedSearchRequestDerefAliases)
	if err != nil {
		return
	}
	searchrequest.sizeLimit, err = readPositiveINTEGER(bytes)
	if err != nil {
		return
	}
	searchrequest.timeLimit, err = readPositiveINTEGER(bytes)
	if err != nil {
		return
	}
	searchrequest.typesOnly, err = readBOOLEAN(bytes)
	if err != nil {
		return
	}
	searchrequest.filter, err = readFilter(bytes)
	if err != nil {
		return
	}
	searchrequest.attributes, err = readAttributeSelection(bytes)
	if err != nil {
		return
	}
	return
}

//
//        AttributeSelection ::= SEQUENCE OF selector LDAPString
//                       -- The LDAPString is constrained to
//                       -- <attributeSelector> in Section 4.5.1.8
func readAttributeSelection(bytes Bytes) (attributeSelection AttributeSelection, err error) {
	err = bytes.ReadSubBytes(classUniversal, tagSequence, attributeSelection.readAttributeSelectionComponents)
	return
}
func (attributeSelection *AttributeSelection) readAttributeSelectionComponents(bytes Bytes) (err error) {
	for bytes.HasMoreData() {
		var ldapstring LDAPString
		ldapstring, err = readLDAPString(bytes)
		// @TOTO: check <attributeSelector> in Section 4.5.1.8
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
func readFilter(bytes Bytes) (filter Filter, err error) {
	var tagAndLength TagAndLength
	tagAndLength, err = bytes.PreviewTagAndLength()
	if err != nil {
		return
	}
	err = tagAndLength.ExpectClass(classContextSpecific)
	if err != nil {
		return
	}
	switch tagAndLength.Tag {
	case TagFilterAnd:
		filter, err = readFilterAnd(bytes)
	case TagFilterOr:
		filter, err = readFilterOr(bytes)
	case TagFilterNot:
		filter, err = readFilterNot(bytes)
	case TagFilterEqualityMatch:
		filter, err = readFilterEqualityMatch(bytes)
	case TagFilterSubstrings:
		filter, err = readFilterSubstrings(bytes)
	case TagFilterGreaterOrEqual:
		filter, err = readFilterGreaterOrEqual(bytes)
	case TagFilterLessOrEqual:
		filter, err = readFilterLessOrEqual(bytes)
	case TagFilterPresent:
		filter, err = readFilterPresent(bytes)
	case TagFilterApproxMatch:
		filter, err = readFilterApproxMatch(bytes)
	case TagFilterExtensibleMatch:
		filter, err = readFilterExtensibleMatch(bytes)
	default:
		err = LdapError{fmt.Sprintf("Invalid tag value for filter: %d.", tagAndLength.Tag)}
	}
	if err != nil {
		err = LdapError{fmt.Sprintf("readFilter: %s.", err.Error())}
	}
	return
}

//             and             [0] SET SIZE (1..MAX) OF filter Filter,
func readFilterAnd(bytes Bytes) (filterand FilterAnd, err error) {
	err = bytes.ReadSubBytes(classContextSpecific, TagFilterAnd, filterand.readFilterAndComponents)
	return
}
func (filterand *FilterAnd) readFilterAndComponents(bytes Bytes) (err error) {
	for bytes.HasMoreData() {
		var filter Filter
		filter, err = readFilter(bytes)
		if err != nil {
			return
		}
		*filterand = append(*filterand, filter)
	}
	if len(*filterand) == 0 {
		err = LdapError{"readFilterAnd: expecting at least one Filter"}
	}
	return
}

//             or              [1] SET SIZE (1..MAX) OF filter Filter,
func readFilterOr(bytes Bytes) (filteror FilterOr, err error) {
	err = bytes.ReadSubBytes(classContextSpecific, TagFilterOr, filteror.readFilterOrComponents)
	return
}

func (filteror *FilterOr) readFilterOrComponents(bytes Bytes) (err error) {
	for bytes.HasMoreData() {
		var filter Filter
		filter, err = readFilter(bytes)
		if err != nil {
			return
		}
		*filteror = append(*filteror, filter)
	}
	if len(*filteror) == 0 {
		err = LdapError{"readFilterOr: expecting at least one Filter"}
	}
	return
}

//             not             [2] Filter,
func readFilterNot(bytes Bytes) (filternot FilterNot, err error) {
	err = bytes.ReadSubBytes(classContextSpecific, TagFilterNot, filternot.readFilterNotComponents)
	return
}

func (filternot *FilterNot) readFilterNotComponents(bytes Bytes) (err error) {
	var tagAndLength TagAndLength
	tagAndLength, err = bytes.ParseTagAndLength()
	if err != nil {
		return
	}
	err = tagAndLength.ExpectTag(TagFilterNot)
	if err != nil {
		return
	}
	filternot.Filter, err = readFilter(bytes)
	if err != nil {
		return
	}
	return
}

//             equalityMatch   [3] AttributeValueAssertion,
func readFilterEqualityMatch(bytes Bytes) (ret FilterEqualityMatch, err error) {
	var attributevalueassertion AttributeValueAssertion
	attributevalueassertion, err = readTaggedAttributeValueAssertion(bytes, classContextSpecific, TagFilterEqualityMatch)
	if err != nil {
		return
	}
	ret = FilterEqualityMatch(attributevalueassertion)
	return
}

//             substrings      [4] SubstringFilter,
func readFilterSubstrings(bytes Bytes) (filtersubstrings FilterSubstrings, err error) {
	var substringfilter SubstringFilter
	substringfilter, err = readTaggedSubstringFilter(bytes, classContextSpecific, TagFilterSubstrings)
	if err != nil {
		return
	}
	filtersubstrings = FilterSubstrings(substringfilter)
	return
}

//             greaterOrEqual  [5] AttributeValueAssertion,
func readFilterGreaterOrEqual(bytes Bytes) (ret FilterGreaterOrEqual, err error) {
	var attributevalueassertion AttributeValueAssertion
	attributevalueassertion, err = readTaggedAttributeValueAssertion(bytes, classContextSpecific, TagFilterGreaterOrEqual)
	if err != nil {
		return
	}
	ret = FilterGreaterOrEqual(attributevalueassertion)
	return
}

//             lessOrEqual     [6] AttributeValueAssertion,
func readFilterLessOrEqual(bytes Bytes) (ret FilterLessOrEqual, err error) {
	var attributevalueassertion AttributeValueAssertion
	attributevalueassertion, err = readTaggedAttributeValueAssertion(bytes, classContextSpecific, TagFilterLessOrEqual)
	if err != nil {
		return
	}
	ret = FilterLessOrEqual(attributevalueassertion)
	return
}

//             present         [7] AttributeDescription,
func readFilterPresent(bytes Bytes) (ret FilterPresent, err error) {
	var attributedescription AttributeDescription
	attributedescription, err = readTaggedAttributeDescription(bytes, classContextSpecific, TagFilterPresent)
	if err != nil {
		return ret, LdapError{fmt.Sprintf("readFilterPresent: %s", err.Error())}
	}
	ret = FilterPresent(attributedescription)
	return
}

//             approxMatch     [8] AttributeValueAssertion,
func readFilterApproxMatch(bytes Bytes) (ret FilterApproxMatch, err error) {
	var attributevalueassertion AttributeValueAssertion
	attributevalueassertion, err = readTaggedAttributeValueAssertion(bytes, classContextSpecific, TagFilterApproxMatch)
	if err != nil {
		return
	}
	ret = FilterApproxMatch(attributevalueassertion)
	return
}

//             extensibleMatch [9] MatchingRuleAssertion,
func readFilterExtensibleMatch(bytes Bytes) (filterextensiblematch FilterExtensibleMatch, err error) {
	var matchingruleassertion MatchingRuleAssertion
	matchingruleassertion, err = readTaggedMatchingRuleAssertion(bytes, classContextSpecific, TagFilterExtensibleMatch)
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
func readTaggedSubstringFilter(bytes Bytes, class int, tag int) (substringfilter SubstringFilter, err error) {
	err = bytes.ReadSubBytes(class, tag, substringfilter.readSubstringFilterComponents)
	return
}
func (substringfilter SubstringFilter) readSubstringFilterComponents(bytes Bytes) (err error) {
	substringfilter.type_, err = readAttributeDescription(bytes)
	if err != nil {
		return
	}
	substringfilter.substrings, err = readSubstringFilterSubstrings(bytes)
	if err != nil {
		return
	}
	return
}

func readSubstringFilterSubstrings(bytes Bytes) (substrings SubstringFilterSubstrings, err error) {
	err = bytes.ReadSubBytes(classUniversal, tagSequence, substrings.readSubstringFilterSubstringsComponents)
	return
}

func (substrings *SubstringFilterSubstrings) readSubstringFilterSubstringsComponents(bytes Bytes) (err error) {
	var foundInitial = 0
	var foundFinal = 0
	var tagAndLength TagAndLength
	for bytes.HasMoreData() {
		tagAndLength, err = bytes.PreviewTagAndLength()
		if err != nil {
			return
		}
		var assertionvalue AssertionValue
		switch tagAndLength.Tag {
		case TagSubstringInitial:
			foundInitial++
			if foundInitial > 1 {
				return LdapError{"readSubstring: initial can occur at most once"}
			}
			assertionvalue, err = readTaggedAssertionValue(bytes, classContextSpecific, TagSubstringInitial)
			if err != nil {
				return
			}
			*substrings = append(*substrings, SubstringInitial(assertionvalue))
		case TagSubstringAny:
			assertionvalue, err = readTaggedAssertionValue(bytes, classContextSpecific, TagSubstringAny)
			if err != nil {
				return
			}
			*substrings = append(*substrings, SubstringAny(assertionvalue))
		case TagSubstringFinal:
			foundFinal++
			if foundFinal > 1 {
				return LdapError{"readSubstring: final can occur at most once"}
			}
			assertionvalue, err = readTaggedAssertionValue(bytes, classContextSpecific, TagSubstringFinal)
			if err != nil {
				return
			}
			*substrings = append(*substrings, SubstringFinal(assertionvalue))
		default:
			return LdapError{fmt.Sprintf("readSubstring: invalid tag %d", tagAndLength.Tag)}
		}
	}
	if len(*substrings) == 0 {
		err = LdapError{"readSubstringFilterSubstrings: expecting at least one substring"}
	}
	return
}

//
//        MatchingRuleAssertion ::= SEQUENCE {
//             matchingRule    [1] MatchingRuleId OPTIONAL,
//             type            [2] AttributeDescription OPTIONAL,
//             matchValue      [3] AssertionValue,
//             dnAttributes    [4] BOOLEAN DEFAULT FALSE }
func readTaggedMatchingRuleAssertion(bytes Bytes, class int, tag int) (ret MatchingRuleAssertion, err error) {
	err = bytes.ReadSubBytes(class, tag, ret.readMatchingRuleAssertionComponents)
	return
}
func (matchingruleassertion MatchingRuleAssertion) readMatchingRuleAssertionComponents(bytes Bytes) (err error) {
	err = matchingruleassertion.readMatchingRule(bytes)
	if err != nil {
		return LdapError{fmt.Sprintf("readMatchingRuleAssertionComponents: %s", err.Error())}
	}
	err = matchingruleassertion.readType(bytes)
	if err != nil {
		return LdapError{fmt.Sprintf("readMatchingRuleAssertionComponents: %s", err.Error())}
	}
	matchingruleassertion.matchValue, err = readTaggedAssertionValue(bytes, classContextSpecific, TagMatchingRuleAssertionMatchValue)
	if err != nil {
		return LdapError{fmt.Sprintf("readMatchingRuleAssertionComponents: %s", err.Error())}
	}
	matchingruleassertion.dnAttributes, err = readTaggedBOOLEAN(bytes, classContextSpecific, TagMatchingRuleAssertionDnAttributes)
	if err != nil {
		return LdapError{fmt.Sprintf("readMatchingRuleAssertionComponents: %s", err.Error())}
	}
	return
}
func (matchingruleassertion MatchingRuleAssertion) readMatchingRule(bytes Bytes) (err error) {
	var tagAndLength TagAndLength
	tagAndLength, err = bytes.PreviewTagAndLength()
	if err != nil {
		return LdapError{fmt.Sprintf("readMatchingRuleAssertionMatchingRule: %s", err.Error())}
	}
	if tagAndLength.Tag == TagMatchingRuleAssertionMatchingRule {
		var matchingRule MatchingRuleId
		matchingRule, err = readTaggedMatchingRuleId(bytes, classContextSpecific, TagMatchingRuleAssertionMatchingRule)
		if err != nil {
			return LdapError{fmt.Sprintf("readMatchingRuleAssertionMatchingRule: %s", err.Error())}
		}
		matchingruleassertion.matchingRule = &matchingRule
	}
	return
}
func (matchingruleassertion MatchingRuleAssertion) readType(bytes Bytes) (err error) {
	var tagAndLength TagAndLength
	tagAndLength, err = bytes.PreviewTagAndLength()
	if err != nil {
		return LdapError{fmt.Sprintf("readMatchingRuleAssertionType: %s", err.Error())}
	}
	if tagAndLength.Tag == TagMatchingRuleAssertionType {
		var attributedescription AttributeDescription
		attributedescription, err = readTaggedAttributeDescription(bytes, classContextSpecific, TagMatchingRuleAssertionType)
		if err != nil {
			return LdapError{fmt.Sprintf("readMatchingRuleAssertionType: %s", err.Error())}
		}
		matchingruleassertion.type_ = &attributedescription
	}
	return
}

//
//        SearchResultEntry ::= [APPLICATION 4] SEQUENCE {
//             objectName      LDAPDN,
//             attributes      PartialAttributeList }
func readSearchResultEntry(bytes Bytes) (searchresultentry SearchResultEntry, err error) {
	err = bytes.ReadSubBytes(classApplication, TagSearchResultEntry, searchresultentry.readSearchResultEntryComponents)
	return
}
func (searchresultentry *SearchResultEntry) readSearchResultEntryComponents(bytes Bytes) (err error) {
	searchresultentry.objectName, err = readLDAPDN(bytes)
	if err != nil {
		return
	}
	searchresultentry.attributes, err = readPartialAttributeList(bytes)
	if err != nil {
		return
	}
	return
}

//
//        PartialAttributeList ::= SEQUENCE OF
//                             partialAttribute PartialAttribute
func readPartialAttributeList(bytes Bytes) (ret PartialAttributeList, err error) {
	ret = PartialAttributeList(make([]PartialAttribute, 0, 10))
	err = bytes.ReadSubBytes(classUniversal, tagSequence, ret.readPartialAttributeListComponents)
	return ret, err
}
func (partialattributelist *PartialAttributeList) readPartialAttributeListComponents(bytes Bytes) (err error) {
	for bytes.HasMoreData() {
		var partialattribute PartialAttribute
		partialattribute, err = readPartialAttribute(bytes)
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
func readSearchResultReference(bytes Bytes) (ret SearchResultReference, err error) {
	err = bytes.ReadSubBytes(classApplication, TagSearchResultReference, ret.readComponents)
	return
}
func (s *SearchResultReference) readComponents(bytes Bytes) (err error) {
	for bytes.HasMoreData() {
		var uri URI
		uri, err = readURI(bytes)
		if err != nil {
			return
		}
		*s = append(*s, uri)
	}
	if len(*s) == 0 {
		err = LdapError{"SearchResultReference: expecting at least one URI"}
	}
	return
}

//
//        SearchResultDone ::= [APPLICATION 5] LDAPResult
func readSearchResultDone(bytes Bytes) (ret SearchResultDone, err error) {
	var ldapresult LDAPResult
	ldapresult, err = readTaggedLDAPResult(bytes, classApplication, TagSearchResultDone)
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
func readModifyRequest(bytes Bytes) (ret ModifyRequest, err error) {
	err = bytes.ReadSubBytes(classApplication, TagModifyRequest, ret.readComponents)
	return
}
func (m *ModifyRequest) readComponents(bytes Bytes) (err error) {
	m.object, err = readLDAPDN(bytes)
	if err != nil {
		return
	}
	err = bytes.ReadSubBytes(classUniversal, tagSequence, m.readChanges)
	return
}
func (m *ModifyRequest) readChanges(bytes Bytes) (err error) {
	for bytes.HasMoreData() {
		var c ModifyRequestChange
		c, err = readModifyRequestChange(bytes)
		if err != nil {
			return
		}
		m.changes = append(m.changes, c)
	}
	return
}
func readModifyRequestChange(bytes Bytes) (ret ModifyRequestChange, err error) {
	err = bytes.ReadSubBytes(classUniversal, tagSequence, ret.readComponents)
	return
}
func (m *ModifyRequestChange) readComponents(bytes Bytes) (err error) {
	m.operation, err = readENUMERATED(bytes, EnumeratedModifyRequestChangeOpration)
	if err != nil {
		return
	}
	m.modification, err = readPartialAttribute(bytes)
	return
}

//
//        ModifyResponse ::= [APPLICATION 7] LDAPResult
func readModifyResponse(bytes Bytes) (ret ModifyResponse, err error) {
	var res LDAPResult
	res, err = readTaggedLDAPResult(bytes, classApplication, TagModifyResponse)
	if err != nil {
		return
	}
	ret = ModifyResponse(res)
	return
}

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
func readAddRequest(bytes Bytes) (ret AddRequest, err error) {
	err = bytes.ReadSubBytes(classApplication, TagAddRequest, ret.readComponents)
	return
}
func (req *AddRequest) readComponents(bytes Bytes) (err error) {
	req.entry, err = readLDAPDN(bytes)
	if err != nil {
		return
	}
	req.attributes, err = readAttributeList(bytes)
	return
}

//
//        AttributeList ::= SEQUENCE OF attribute Attribute
func readAttributeList(bytes Bytes) (ret AttributeList, err error) {
	err = bytes.ReadSubBytes(classUniversal, tagSequence, ret.readComponents)
	return
}
func (list *AttributeList) readComponents(bytes Bytes) (err error) {
	for bytes.HasMoreData() {
		var attr Attribute
		attr, err = readAttribute(bytes)
		if err != nil {
			return
		}
		*list = append(*list, attr)
	}
	return
}

//
//        AddResponse ::= [APPLICATION 9] LDAPResult
func readAddResponse(bytes Bytes) (ret AddResponse, err error) {
	var res LDAPResult
	res, err = readTaggedLDAPResult(bytes, classApplication, TagAddResponse)
	if err != nil {
		return
	}
	ret = AddResponse(res)
	return
}

//
//        DelRequest ::= [APPLICATION 10] LDAPDN
func readDelRequest(bytes Bytes) (ret DelRequest, err error) {
	var res LDAPDN
	res, err = readTaggedLDAPDN(bytes, classApplication, TagDelRequest)
	if err != nil {
		return
	}
	ret = DelRequest(res)
	return
}

//
//        DelResponse ::= [APPLICATION 11] LDAPResult
func readDelResponse(bytes Bytes) (ret DelResponse, err error) {
	var res LDAPResult
	res, err = readTaggedLDAPResult(bytes, classApplication, TagDelResponse)
	if err != nil {
		return
	}
	ret = DelResponse(res)
	return
}

//
//        ModifyDNRequest ::= [APPLICATION 12] SEQUENCE {
//             entry           LDAPDN,
//             newrdn          RelativeLDAPDN,
//             deleteoldrdn    BOOLEAN,
//             newSuperior     [0] LDAPDN OPTIONAL }
func readModifyDNRequest(bytes Bytes) (ret ModifyDNRequest, err error) {
	err = bytes.ReadSubBytes(classApplication, TagModifyDNRequest, ret.readComponents)
	return
}
func (req *ModifyDNRequest) readComponents(bytes Bytes) (err error) {
	req.entry, err = readLDAPDN(bytes)
	if err != nil {
		return
	}
	req.newrdn, err = readRelativeLDAPDN(bytes)
	if err != nil {
		return
	}
	req.deleteoldrdn, err = readBOOLEAN(bytes)
	if err != nil {
		return
	}
	if bytes.HasMoreData() {
		var ldapdn LDAPDN
		ldapdn, err = readTaggedLDAPDN(bytes, classContextSpecific, TagModifyDNRequestNewSuperior)
		if err != nil {
			return
		}
		req.newSuperior = &ldapdn
	}
	return
}

//
//        ModifyDNResponse ::= [APPLICATION 13] LDAPResult
func readModifyDNResponse(bytes Bytes) (ret ModifyDNResponse, err error) {
	var res LDAPResult
	res, err = readTaggedLDAPResult(bytes, classApplication, TagModifyDNResponse)
	if err != nil {
		return
	}
	ret = ModifyDNResponse(res)
	return
}

//
//        CompareRequest ::= [APPLICATION 14] SEQUENCE {
//             entry           LDAPDN,
//             ava             AttributeValueAssertion }
func readCompareRequest(bytes Bytes) (ret CompareRequest, err error) {
	err = bytes.ReadSubBytes(classApplication, TagCompareRequest, ret.readComponents)
	return
}
func (req *CompareRequest) readComponents(bytes Bytes) (err error) {
	req.entry, err = readLDAPDN(bytes)
	if err != nil {
		return
	}
	req.ava, err = readAttributeValueAssertion(bytes)
	return
}

//
//        CompareResponse ::= [APPLICATION 15] LDAPResult
func readCompareResponse(bytes Bytes) (ret CompareResponse, err error) {
	var res LDAPResult
	res, err = readTaggedLDAPResult(bytes, classApplication, TagCompareResponse)
	if err != nil {
		return
	}
	ret = CompareResponse(res)
	return
}

//
//        AbandonRequest ::= [APPLICATION 16] MessageID
func readAbandonRequest(bytes Bytes) (ret AbandonRequest, err error) {
	var mes MessageID
	mes, err = readTaggedMessageID(bytes, classApplication, TagAbandonRequest)
	if err != nil {
		return
	}
	ret = AbandonRequest(mes)
	return
}

//
//        ExtendedRequest ::= [APPLICATION 23] SEQUENCE {
//             requestName      [0] LDAPOID,
//             requestValue     [1] OCTET STRING OPTIONAL }
func readExtendedRequest(bytes Bytes) (ret ExtendedRequest, err error) {
	err = bytes.ReadSubBytes(classApplication, TagExtendedRequest, ret.readComponents)
	return
}
func (req *ExtendedRequest) readComponents(bytes Bytes) (err error) {
	req.requestName, err = readTaggedLDAPOID(bytes, classContextSpecific, TagExtendedRequestName)
	if err != nil {
		return
	}
	if bytes.HasMoreData() {
		var str OCTETSTRING
		str, err = readTaggedOCTETSTRING(bytes, classContextSpecific, TagExtendedRequestValue)
		if err != nil {
			return
		}
		req.requestValue = &str
	}
	return
}

//
//        ExtendedResponse ::= [APPLICATION 24] SEQUENCE {
//             COMPONENTS OF LDAPResult,
//             responseName     [10] LDAPOID OPTIONAL,
//             responseValue    [11] OCTET STRING OPTIONAL }
func readExtendedResponse(bytes Bytes) (ret ExtendedResponse, err error) {
	err = bytes.ReadSubBytes(classApplication, TagExtendedResponse, ret.readComponents)
	return
}
func (res *ExtendedResponse) readComponents(bytes Bytes) (err error) {
	res.readLDAPResultComponents(bytes)
	if bytes.HasMoreData() {
		var oid LDAPOID
		oid, err = readTaggedLDAPOID(bytes, classContextSpecific, TagExtendedResponseName)
		if err != nil {
			return
		}
		res.responseName = &oid
	}
	if bytes.HasMoreData() {
		var str OCTETSTRING
		str, err = readTaggedOCTETSTRING(bytes, classContextSpecific, TagExtendedResponseValue)
		if err != nil {
			return
		}
		res.responseValue = &str
	}
	return
}

//
//        IntermediateResponse ::= [APPLICATION 25] SEQUENCE {
//             responseName     [0] LDAPOID OPTIONAL,
//             responseValue    [1] OCTET STRING OPTIONAL }
func readIntermediateResponse(bytes Bytes) (ret IntermediateResponse, err error) {
	err = bytes.ReadSubBytes(classApplication, TagIntermediateResponse, ret.readComponents)
	return
}
func (res *IntermediateResponse) readComponents(bytes Bytes) (err error) {
	if bytes.HasMoreData() {
		var oid LDAPOID
		oid, err = readTaggedLDAPOID(bytes, classContextSpecific, TagIntermediateResponseName)
		if err != nil {
			return
		}
		res.responseName = &oid
	}
	if bytes.HasMoreData() {
		var str OCTETSTRING
		str, err = readTaggedOCTETSTRING(bytes, classContextSpecific, TagIntermediateResponseValue)
		if err != nil {
			return
		}
		res.responseValue = &str
	}
	return
}

//
//        END
//
