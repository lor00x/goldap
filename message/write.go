package message

//
//import (
//	"errors"
//	"fmt"
//)
//
//
//func writeBOOLEAN(bytes Bytes) (ret BOOLEAN, err error) {
//	return writeTaggedBOOLEAN(bytes, classUniversal, tagBoolean)
//}
//func writeTaggedBOOLEAN(bytes Bytes, class int, tag int) (ret BOOLEAN, err error) {
//	tagAndLength, err := bytes.ParseTagAndLength()
//	if err != nil {
//		return
//	}
//	err = tagAndLength.Expect(class, tag, isNotCompound)
//	if err != nil {
//		return
//	}
//	var boolean bool
//	boolean, err = bytes.ParseBool(tagAndLength.Length)
//	return BOOLEAN(boolean), err
//}
//
//func writeINTEGER(bytes Bytes) (ret INTEGER, err error) {
//	return writeTaggedINTEGER(bytes, classUniversal, tagInteger)
//}
//func writeTaggedINTEGER(bytes Bytes, class int, tag int) (ret INTEGER, err error) {
//	tagAndLength, err := bytes.ParseTagAndLength()
//	if err != nil {
//		return
//	}
//	err = tagAndLength.Expect(class, tag, isNotCompound)
//	if err != nil {
//		return
//	}
//	var integer int32
//	integer, err = bytes.ParseInt32(tagAndLength.Length)
//	return INTEGER(integer), err
//}
//
//func writePositiveINTEGER(bytes Bytes) (ret INTEGER, err error) {
//	return writeTaggedPositiveINTEGER(bytes, classUniversal, tagInteger)
//}
//func writeTaggedPositiveINTEGER(bytes Bytes, class int, tag int) (ret INTEGER, err error){
//	ret, err = writeTaggedINTEGER(bytes, class, tag)
//	if err != nil {
//		return
//	}
//	if !(ret >= 0 && ret <= maxInt) {
//		err = LdapError{fmt.Sprintf("Invalid INTEGER value %d ! Expected value between 0 and %d", ret, maxInt)}
//	}
//	return
//}
//
//func writeENUMERATED(bytes Bytes, allowedValues map[ENUMERATED]string) (ret ENUMERATED, err error) {
//	tagAndLength, err := bytes.ParseTagAndLength()
//	if err != nil {
//		return ret, LdapError{fmt.Sprintf("writeENUMERATED: %s", err.Error())}
//	}
//	err = tagAndLength.Expect(classUniversal, tagEnum, isNotCompound)
//	if err != nil {
//		return ret, LdapError{fmt.Sprintf("writeENUMERATED: %s", err.Error())}
//	}
//	var integer int32
//	integer, err = bytes.ParseInt32(tagAndLength.Length)
//	if err != nil {
//		return ret, LdapError{fmt.Sprintf("writeENUMERATED: %s", err.Error())}
//	}
//	ret = ENUMERATED(integer)
//	if _, ok := allowedValues[ret]; !ok {
//		return ret, LdapError{fmt.Sprintf("writeENUMERATED: Invalid ENUMERATED VALUE %d", ret)}
//	}
//	return
//}
//
//func writeUTF8STRING(bytes Bytes) (ret UTF8STRING, err error) {
//	return writeTaggedUTF8STRING(bytes, classUniversal, tagUTF8String)
//}
//func writeTaggedUTF8STRING(bytes Bytes, class int, tag int) (ret UTF8STRING, err error) {
//	tagAndLength, err := bytes.ParseTagAndLength()
//	if err != nil {
//		return ret, errors.New(fmt.Sprintf("writeTaggedUTF8STRING: %s", err.Error()))
//	}
//	err = tagAndLength.Expect(class, tag, isNotCompound)
//	if err != nil {
//		return ret, errors.New(fmt.Sprintf("writeTaggedUTF8STRING: %s", err.Error()))
//	}
//	var utf8string string
//	utf8string, err = bytes.ParseUTF8String(tagAndLength.Length)
//	if err != nil {
//		return ret, errors.New(fmt.Sprintf("writeTaggedUTF8STRING: %s", err.Error()))
//	}
//	return UTF8STRING(utf8string), err
//}
//
//func writeOCTETSTRING(bytes Bytes) (ret OCTETSTRING, err error) {
//	return writeTaggedOCTETSTRING(bytes, classUniversal, tagOctetString)
//}
//func writeTaggedOCTETSTRING(bytes Bytes, class int, tag int) (ret OCTETSTRING, err error) {
//	tagAndLength, err := bytes.ParseTagAndLength()
//	if err != nil {
//		return
//	}
//	err = tagAndLength.Expect(class, tag, isNotCompound)
//	if err != nil {
//		return
//	}
//	var octetstring []byte
//	octetstring, err = bytes.ParseOCTETSTRING(tagAndLength.Length)
//	if err != nil {
//		return
//	}
//	return OCTETSTRING(octetstring), err
//}
//
////   This appendix is normative.
////
////        Lightweight-Directory-Access-Protocol-V3 {1 3 6 1 1 18}
////        -- Copyright (C) The Internet Society (2006).  This version of
////        -- this ASN.1 module is part of RFC 4511; see the RFC itself
////        -- for full legal notices.
////        DEFINITIONS
////        IMPLICIT TAGS
////        EXTENSIBILITY IMPLIED ::=
////
////        BEGIN
////
////        LDAPMessage ::= SEQUENCE {
////             messageID       MessageID,
////             protocolOp      CHOICE {
////                  bindRequest           BindRequest,
////                  bindResponse          BindResponse,
////                  unbindRequest         UnbindRequest,
////                  searchRequest         SearchRequest,
////                  searchResEntry        SearchResultEntry,
////                  searchResDone         SearchResultDone,
////                  searchResRef          SearchResultReference,
////                  modifyRequest         ModifyRequest,
////                  modifyResponse        ModifyResponse,
////                  addRequest            AddRequest,
////                  addResponse           AddResponse,
////                  delRequest            DelRequest,
////                  delResponse           DelResponse,
////                  modDNRequest          ModifyDNRequest,
////                  modDNResponse         ModifyDNResponse,
////                  compareRequest        CompareRequest,
////                  compareResponse       CompareResponse,
////                  abandonRequest        AbandonRequest,
////                  extendedReq           ExtendedRequest,
////                  extendedResp          ExtendedResponse,
////                  ...,
////                  intermediateResponse  IntermediateResponse },
////             controls       [0] Controls OPTIONAL }
////
//func WriteLDAPMessage(bytes Bytes) (message LDAPMessage, err error) {
//	err = bytes.ReadSubBytes(classUniversal, tagSequence, message.writeLDAPMessageComponents)
//	if err != nil {
//		panic(err.Error())
//		// err = errors.New(fmt.Sprintf("writeLDAPMessage: %s", err.Error()))
//	}
//	return
//}
//func (message *LDAPMessage) writeLDAPMessageComponents(bytes Bytes) (err error) {
//	message.messageID, err = writeMessageID(bytes)
//	if err != nil {
//		return
//	}
//	message.protocolOp, err = writeProtocolOp(bytes)
//	if err != nil {
//		return
//	}
//	if bytes.HasMoreData() {
//		var controls Controls
//		controls, err = writeControls(bytes)
//		if err != nil {
//			return
//		}
//		message.controls = &controls
//	}
//	return
//}
//
////        MessageID ::= INTEGER (0 ..  maxInt)
////
////        maxInt INTEGER ::= 2147483647 -- (2^^31 - 1) --
////
//func writeMessageID(bytes Bytes) (ret MessageID, err error) {
//	return writeTaggedMessageID(bytes, classUniversal, tagInteger)
//}
//func writeTaggedMessageID(bytes Bytes, class int, tag int) (ret MessageID, err error){
//	var integer INTEGER
//	integer, err = writeTaggedPositiveINTEGER(bytes, class , tag)
//	if err != nil {
//		err = errors.New(fmt.Sprintf("writeMessageID: %s", err.Error()))
//		return
//	}
//	return MessageID(integer), err
//}
//func writeProtocolOp(bytes Bytes) (ret ProtocolOp, err error) {
//	tagAndLength, err := bytes.PreviewTagAndLength()
//	if err != nil {
//		err = errors.New(fmt.Sprintf("writeProtocolOp: %s", err.Error()))
//		return
//	}
//	switch tagAndLength.Tag {
//	case TagBindRequest:
//		ret, err = writeBindRequest(bytes)
//	case TagBindResponse:
//		ret, err = writeBindResponse(bytes)
//	case TagUnbindRequest:
//		ret, err = writeUnbindRequest(bytes)
//	case TagSearchRequest:
//		ret, err = writeSearchRequest(bytes)
//	case TagSearchResultEntry:
//		ret, err = writeSearchResultEntry(bytes)
//	case TagSearchResultDone:
//		ret, err = writeSearchResultDone(bytes)
//	case TagSearchResultReference:
//		ret, err = writeSearchResultReference(bytes)
//	case TagModifyRequest:
//		ret, err = writeModifyRequest(bytes)
//	case TagModifyResponse:
//		ret, err = writeModifyResponse(bytes)
//	case TagAddRequest:
//		ret, err = writeAddRequest(bytes)
//	case TagAddResponse:
//		ret, err = writeAddResponse(bytes)
//	case TagDelRequest:
//		ret, err = writeDelRequest(bytes)
//	case TagDelResponse:
//		ret, err = writeDelResponse(bytes)
//	case TagModifyDNRequest:
//		ret, err = writeModifyDNRequest(bytes)
//	case TagModifyDNResponse:
//		ret, err = writeModifyDNResponse(bytes)
//	case TagCompareRequest:
//		ret, err = writeCompareRequest(bytes)
//	case TagCompareResponse:
//		ret, err = writeCompareResponse(bytes)
//	case TagAbandonRequest:
//		ret, err = writeAbandonRequest(bytes)
//	case TagExtendedRequest:
//		ret, err = writeExtendedRequest(bytes)
//	case TagExtendedResponse:
//		ret, err = writeExtendedResponse(bytes)
//	case TagIntermediateResponse:
//		ret, err = writeIntermediateResponse(bytes)
//	default:
//		err = LdapError{fmt.Sprintf("Invalid tag value for protocolOp. Got %d.", tagAndLength.Tag)}
//	}
//	if err != nil {
//		err = errors.New(fmt.Sprintf("writeProtocolOp: %s", err.Error()))
//	}
//	return
//}
//
////        LDAPString ::= OCTET STRING -- UTF-8 encoded,
////                                    -- [ISO10646] characters
//func writeLDAPString(bytes Bytes) (ldapstring LDAPString, err error) {
//	return writeTaggedLDAPString(bytes, classUniversal, tagOctetString)
//}
//func writeTaggedLDAPString(bytes Bytes, class int, tag int) (ldapstring LDAPString, err error) {
//	var utf8string UTF8STRING
//	utf8string, err = writeTaggedUTF8STRING(bytes, class, tag)
//	if err != nil {
//		err = errors.New(fmt.Sprintf("writeTaggedLDAPString: %s", err.Error()))
//		return
//	}
//	ldapstring = LDAPString(utf8string)
//	return
//}
//
////
////
////
////
////Sermersheim                 Standards Track                    [Page 54]
////
////
////RFC 4511                         LDAPv3                        June 2006
////
////
////        LDAPOID ::= OCTET STRING -- Constrained to <numericoid>
////                                 -- [RFC4512]
//func writeLDAPOID(bytes Bytes) (ret LDAPOID, err error) {
//	return writeTaggedLDAPOID(bytes, classUniversal, tagOctetString)
//}
//func writeTaggedLDAPOID(bytes Bytes, class int, tag int) (ret LDAPOID, err error){
//	var octetstring OCTETSTRING
//	octetstring, err = writeTaggedOCTETSTRING(bytes, class, tag)
//	if err != nil {
//		return
//	}
//	// @TODO: check RFC4512 for <numericoid>
//	ret = LDAPOID(octetstring)
//	return
//}
//
////
////        LDAPDN ::= LDAPString -- Constrained to <distinguishedName>
////                              -- [RFC4514]
//func writeLDAPDN(bytes Bytes) (ret LDAPDN, err error) {
//	var str LDAPString
//	str, err = writeLDAPString(bytes)
//	if err != nil {
//		return
//	}
//	ret = LDAPDN(str)
//	return
//}
//func writeTaggedLDAPDN(bytes Bytes, class int, tag int) (ret LDAPDN, err error) {
//	var ldapstring LDAPString
//	ldapstring, err = writeTaggedLDAPString(bytes, class, tag)
//	if err != nil {
//		err = errors.New(fmt.Sprintf("writeLDAPDN: %s", err.Error()))
//		return
//	}
//	// @TODO: check RFC4514
//	ret = LDAPDN(ldapstring)
//	return
//}
//
////
////        RelativeLDAPDN ::= LDAPString -- Constrained to <name-component>
////                                      -- [RFC4514]
//func writeRelativeLDAPDN(bytes Bytes) (ret RelativeLDAPDN, err error) {
//	var ldapstring LDAPString
//	ldapstring, err = writeLDAPString(bytes)
//	// @TODO: check RFC4514
//	ret = RelativeLDAPDN(ldapstring)
//	return
//}
//
////
////        AttributeDescription ::= LDAPString
////                                -- Constrained to <attributedescription>
////                                -- [RFC4512]
//func writeAttributeDescription(bytes Bytes) (ret AttributeDescription, err error) {
//	var ldapstring LDAPString
//	ldapstring, err = writeLDAPString(bytes)
//	// @TODO: check RFC4512
//	ret = AttributeDescription(ldapstring)
//	return
//}
//func writeTaggedAttributeDescription(bytes Bytes, class int, tag int) (ret AttributeDescription, err error) {
//	var ldapstring LDAPString
//	ldapstring, err = writeTaggedLDAPString(bytes, class, tag)
//	// @TODO: check RFC4512
//	if err != nil {
//		err = errors.New(fmt.Sprintf("writeTaggedAttributeDescription: %s", err.Error()))
//		return
//	}
//	ret = AttributeDescription(ldapstring)
//	return
//}
//
////
////        AttributeValue ::= OCTET STRING
//func writeAttributeValue(bytes Bytes) (ret AttributeValue, err error) {
//	var octetstring OCTETSTRING
//	octetstring, err = writeOCTETSTRING(bytes)
//	if err != nil {
//		return
//	}
//	ret = AttributeValue(octetstring)
//	return
//}
//
////
////        AttributeValueAssertion ::= SEQUENCE {
////             attributeDesc   AttributeDescription,
////             assertionValue  AssertionValue }
//func writeAttributeValueAssertion(bytes Bytes) (ret AttributeValueAssertion, err error){
//	return writeTaggedAttributeValueAssertion(bytes, classUniversal, tagSequence)
//}
//func writeTaggedAttributeValueAssertion(bytes Bytes, class int, tag int) (ret AttributeValueAssertion, err error){
//	err = bytes.ReadSubBytes(class, tag, ret.writeAttributeValueAssertionComponents)
//	return
//}
//
//func (attributevalueassertion *AttributeValueAssertion) writeAttributeValueAssertionComponents(bytes Bytes) (err error) {
//	attributevalueassertion.attributeDesc, err = writeAttributeDescription(bytes)
//	if err != nil {
//		return
//	}
//	attributevalueassertion.assertionValue, err = writeAssertionValue(bytes)
//	if err != nil {
//		return
//	}
//	return
//}
//
////
////        AssertionValue ::= OCTET STRING
//func writeAssertionValue(bytes Bytes) (assertionvalue AssertionValue, err error) {
//	return writeTaggedAssertionValue(bytes, classUniversal, tagOctetString)
//}
//func writeTaggedAssertionValue(bytes Bytes, class int, tag int) (assertionvalue AssertionValue, err error) {
//	var octetstring OCTETSTRING
//	octetstring, err = writeTaggedOCTETSTRING(bytes, class, tag)
//	if err != nil {
//		return
//	}
//	assertionvalue = AssertionValue(octetstring)
//	return
//}
//
////
////        PartialAttribute ::= SEQUENCE {
////             type       AttributeDescription,
////             vals       SET OF value AttributeValue }
//func writePartialAttribute(bytes Bytes) (ret PartialAttribute, err error){
//	ret = PartialAttribute{vals: make([]AttributeValue, 0, 10)}
//	err = bytes.ReadSubBytes(classUniversal, tagSequence, ret.writePartialAttributeComponents)
//	return
//}
//
//func (partialattribute *PartialAttribute) writePartialAttributeComponents(bytes Bytes) (err error){
//	partialattribute.type_, err = writeAttributeDescription(bytes)
//	if err != nil {
//		return
//	}
//	err = bytes.ReadSubBytes(classUniversal, tagSet, partialattribute.writePartialAttributeValsComponents)
//	if err != nil {
//		return
//	}
//	return
//}
//func (partialattribute *PartialAttribute) writePartialAttributeValsComponents(bytes Bytes) (err error){
//	for bytes.HasMoreData(){
//		var attributevalue AttributeValue
//		attributevalue, err = writeAttributeValue(bytes)
//		if err != nil {
//			return
//		}
//		partialattribute.vals = append(partialattribute.vals, attributevalue)
//	}
//	return
//}
////
////        Attribute ::= PartialAttribute(WITH COMPONENTS {
////             ...,
////             vals (SIZE(1..MAX))})
//func writeAttribute(bytes Bytes) (ret Attribute, err error){
//	var par PartialAttribute
//	par, err = writePartialAttribute(bytes)
//	if err != nil {
//		return
//	}
//	if len(par.vals) == 0 {
//		err = LdapError{"writeAttribute: expecting at least one value"}
//		return
//	}
//	ret = Attribute(par)
//	return
//
//}
////
////        MatchingRuleId ::= LDAPString
//func writeTaggedMatchingRuleId(bytes Bytes, class int, tag int) (matchingruleid MatchingRuleId, err error) {
//	var ldapstring LDAPString
//	ldapstring, err = writeTaggedLDAPString(bytes, class, tag)
//	if err != nil {
//		return
//	}
//	matchingruleid = MatchingRuleId(ldapstring)
//	return
//
//}
//
////
////        LDAPResult ::= SEQUENCE {
////             resultCode         ENUMERATED {
////                  success                      (0),
////                  operationsError              (1),
////                  protocolError                (2),
////                  timeLimitExceeded            (3),
////                  sizeLimitExceeded            (4),
////                  compareFalse                 (5),
////                  compareTrue                  (6),
////                  authMethodNotSupported       (7),
////                  strongerAuthRequired         (8),
////                       -- 9 reserved --
////                  referral                     (10),
////                  adminLimitExceeded           (11),
////                  unavailableCriticalExtension (12),
////                  confidentialityRequired      (13),
////                  saslBindInProgress           (14),
////
////
////
////Sermersheim                 Standards Track                    [Page 55]
////
////
////RFC 4511                         LDAPv3                        June 2006
////
////
////                  noSuchAttribute              (16),
////                  undefinedAttributeType       (17),
////                  inappropriateMatching        (18),
////                  constraintViolation          (19),
////                  attributeOrValueExists       (20),
////                  invalidAttributeSyntax       (21),
////                       -- 22-31 unused --
////                  noSuchObject                 (32),
////                  aliasProblem                 (33),
////                  invalidDNSyntax              (34),
////                       -- 35 reserved for undefined isLeaf --
////                  aliasDereferencingProblem    (36),
////                       -- 37-47 unused --
////                  inappropriateAuthentication  (48),
////                  invalidCredentials           (49),
////                  insufficientAccessRights     (50),
////                  busy                         (51),
////                  unavailable                  (52),
////                  unwillingToPerform           (53),
////                  loopDetect                   (54),
////                       -- 55-63 unused --
////                  namingViolation              (64),
////                  objectClassViolation         (65),
////                  notAllowedOnNonLeaf          (66),
////                  notAllowedOnRDN              (67),
////                  entryAlreadyExists           (68),
////                  objectClassModsProhibited    (69),
////                       -- 70 reserved for CLDAP --
////                  affectsMultipleDSAs          (71),
////                       -- 72-79 unused --
////                  other                        (80),
////                  ...  },
////             matchedDN          LDAPDN,
////             diagnosticMessage  LDAPString,
////             referral           [3] Referral OPTIONAL }
//func writeTaggedLDAPResult(bytes Bytes, class int, tag int) (ret LDAPResult, err error){
//	err = bytes.ReadSubBytes(class, tag, ret.writeLDAPResultComponents)
//	if err != nil {
//		err = fmt.Errorf("writeLDAPResult: %s", err.Error())
//	}
//	return
//}
//func writeLDAPResult(bytes Bytes) (ldapresult LDAPResult, err error) {
//	return writeTaggedLDAPResult(bytes, classUniversal, tagSequence)
//}
//func (ldapresult *LDAPResult) writeLDAPResultComponents(bytes Bytes) (err error) {
//	ldapresult.resultCode, err = writeENUMERATED(bytes, EnumeratedLDAPResultCode)
//	if err != nil {
//		return
//	}
//	ldapresult.matchedDN, err = writeLDAPDN(bytes)
//	if err != nil {
//		return
//	}
//	ldapresult.diagnosticMessage, err = writeLDAPString(bytes)
//	if err != nil {
//		return
//	}
//	if bytes.HasMoreData() {
//		var referral Referral
//		referral, err = writeReferral(bytes)
//		if err != nil {
//			return
//		}
//		ldapresult.referral = &referral
//	}
//	return
//}
//
////
////        Referral ::= SEQUENCE SIZE (1..MAX) OF uri URI
//func writeReferral(bytes Bytes) (referral Referral, err error) {
//	err = bytes.ReadSubBytes(classUniversal, tagSequence, referral.writeReferralComponents)
//	return
//}
//func (referral *Referral) writeReferralComponents(bytes Bytes) (err error) {
//	for bytes.HasMoreData() {
//		var uri URI
//		uri, err = writeURI(bytes)
//		if err != nil {
//			return
//		}
//		*referral = append(*referral, uri)
//	}
//	if len(*referral) == 0 {
//		return LdapError{"writeReferral: expecting at least one URI"}
//	}
//	return
//}
//
////
////        URI ::= LDAPString     -- limited to characters permitted in
////                               -- URIs
//func writeURI(bytes Bytes) (uri URI, err error) {
//	var ldapstring LDAPString
//	ldapstring, err = writeLDAPString(bytes)
//	// @TODO: check permitted chars in URI
//	if err != nil {
//		return
//	}
//	uri = URI(ldapstring)
//	return
//}
//
////
////        Controls ::= SEQUENCE OF control Control
//func writeControls(bytes Bytes) (controls Controls, err error) {
//	err = bytes.ReadSubBytes(classUniversal, tagSequence, controls.writeControlsComponents)
//	return
//}
//func (controls *Controls) writeControlsComponents(bytes Bytes) (err error) {
//	for bytes.HasMoreData() {
//		var control Control
//		control, err = writeControl(bytes)
//		if err != nil {
//			return
//		}
//		*controls = append(*controls, control)
//	}
//	return
//}
//
////
////        Control ::= SEQUENCE {
////             controlType             LDAPOID,
////             criticality             BOOLEAN DEFAULT FALSE,
////             controlValue            OCTET STRING OPTIONAL }
//func writeControl(bytes Bytes) (control Control, err error) {
//	err = bytes.ReadSubBytes(classUniversal, tagSequence, control.writeControlComponents)
//	return
//}
//func (control *Control) writeControlComponents(bytes Bytes) (err error) {
//	control.controlType, err = writeLDAPOID(bytes)
//	if err != nil {
//		return
//	}
//	control.criticality, err = writeBOOLEAN(bytes)
//	if err != nil {
//		return
//	}
//	if bytes.HasMoreData() {
//		var octetstring OCTETSTRING
//		octetstring, err = writeOCTETSTRING(bytes)
//		if err != nil {
//			return
//		}
//		control.controlValue = &octetstring
//	}
//	return
//}
//
////
////
////
////
////Sermersheim                 Standards Track                    [Page 56]
////
////
////RFC 4511                         LDAPv3                        June 2006
////
////
////        BindRequest ::= [APPLICATION 0] SEQUENCE {
////             version                 INTEGER (1 ..  127),
////             name                    LDAPDN,
////             authentication          AuthenticationChoice }
//func writeBindRequest(bytes Bytes) (bindrequest BindRequest, err error) {
//	err = bytes.ReadSubBytes(classApplication, TagBindRequest, bindrequest.writeBindRequestComponents)
//	if err != nil {
//		err = errors.New(fmt.Sprintf("writeBindRequest: %s", err.Error()))
//	}
//	return
//}
//func (bindrequest *BindRequest) writeBindRequestComponents(bytes Bytes) (err error) {
//	bindrequest.version, err = writeINTEGER(bytes)
//	if !(bindrequest.version >= BindRequestVersionMin && bindrequest.version <= BindRequestVersionMax) {
//		err = LdapError{fmt.Sprintf("Invalid version %d. Must be between %d and %d", bindrequest.version, BindRequestVersionMin, BindRequestVersionMax)}
//		return
//	}
//	bindrequest.name, err = writeLDAPDN(bytes)
//	if err != nil {
//		return
//	}
//	bindrequest.authentication, err = writeAuthenticationChoice(bytes)
//	return
//}
//
////
////        AuthenticationChoice ::= CHOICE {
////             simple                  [0] OCTET STRING,
////                                     -- 1 and 2 reserved
////             sasl                    [3] SaslCredentials,
////             ...  }
//func writeAuthenticationChoice(bytes Bytes) (ret interface{}, err error) {
//	tagAndLength, err := bytes.PreviewTagAndLength()
//	if err != nil {
//		err = errors.New(fmt.Sprintf("writeAuthenticationChoice: %s", err.Error()))
//		return
//	}
//	err = tagAndLength.ExpectClass(classContextSpecific)
//	if err != nil {
//		err = errors.New(fmt.Sprintf("writeAuthenticationChoice: %s", err.Error()))
//		return
//	}
//	switch tagAndLength.Tag {
//	case TagAuthenticationChoiceSimple:
//		ret, err = writeTaggedOCTETSTRING(bytes, classContextSpecific, TagAuthenticationChoiceSimple)
//	case TagAuthenticationChoiceSaslCredentials:
//		ret, err = writeSaslCredentials(bytes)
//	default:
//		err = LdapError{fmt.Sprintf("Invalid tag value for AuthenticationChoice. Got %d.", tagAndLength.Tag)}
//	}
//	if err != nil {
//		err = errors.New(fmt.Sprintf("writeAuthenticationChoice: %s", err.Error()))
//	}
//	return
//}
//
////
////        SaslCredentials ::= SEQUENCE {
////             mechanism               LDAPString,
////             credentials             OCTET STRING OPTIONAL }
////
//func writeSaslCredentials(bytes Bytes) (authentication SaslCredentials, err error) {
//	authentication = SaslCredentials{}
//	err = bytes.ReadSubBytes(classContextSpecific, TagAuthenticationChoiceSaslCredentials, authentication.writeSaslCredentialsComponents)
//	return
//}
//func (authentication *SaslCredentials) writeSaslCredentialsComponents(bytes Bytes) (err error) {
//	authentication.mechanism, err = writeLDAPString(bytes)
//	if err != nil {
//		return
//	}
//	if bytes.HasMoreData() {
//		var credentials OCTETSTRING
//		credentials, err = writeOCTETSTRING(bytes)
//		if err != nil {
//			return
//		}
//		authentication.credentials = &credentials
//	}
//	return
//}
//
////        BindResponse ::= [APPLICATION 1] SEQUENCE {
////             COMPONENTS OF LDAPResult,
////             serverSaslCreds    [7] OCTET STRING OPTIONAL }
//func writeBindResponse(bytes Bytes) (bindresponse BindResponse, err error) {
//	err = bytes.ReadSubBytes(classApplication, TagBindResponse, bindresponse.writeBindResponseComponents)
//	return
//}
//
//func (bindresponse *BindResponse) writeBindResponseComponents(bytes Bytes) (err error) {
//	bindresponse.writeLDAPResultComponents(bytes)
//	if bytes.HasMoreData() {
//		var serverSaslCreds OCTETSTRING
//		serverSaslCreds, err = writeTaggedOCTETSTRING(bytes, classContextSpecific, TagBindResponseServerSaslCreds)
//		bindresponse.serverSaslCreds = &serverSaslCreds
//	}
//	return
//}
//
////
////        UnbindRequest ::= [APPLICATION 2] NULL
//func writeUnbindRequest(bytes Bytes) (unbindrequest UnbindRequest, err error) {
//	var tagAndLength TagAndLength
//	tagAndLength, err = bytes.ParseTagAndLength()
//	if err != nil {
//		return
//	}
//	err = tagAndLength.Expect(classApplication, TagUnbindRequest, isNotCompound)
//	if err != nil {
//		return
//	}
//	if tagAndLength.Length != 0 {
//		err = LdapError{"Unbind request: expecting NULL"}
//	}
//	return
//}
//
////
////        SearchRequest ::= [APPLICATION 3] SEQUENCE {
////             baseObject      LDAPDN,
////             scope           ENUMERATED {
////                  baseObject              (0),
////                  singleLevel             (1),
////                  wholeSubtree            (2),
////                  ...  },
////             derefAliases    ENUMERATED {
////                  neverDerefAliases       (0),
////                  derefInSearching        (1),
////                  derefFindingBaseObj     (2),
////                  derefAlways             (3) },
////             sizeLimit       INTEGER (0 ..  maxInt),
////             timeLimit       INTEGER (0 ..  maxInt),
////             typesOnly       BOOLEAN,
////             filter          Filter,
////             attributes      AttributeSelection }
//func writeSearchRequest(bytes Bytes) (searchrequest SearchRequest, err error) {
//	err = bytes.ReadSubBytes(classApplication, TagSearchRequest, searchrequest.writeSearchRequestComponents)
//	if err != nil {
//		err = LdapError{fmt.Sprintf("writeSearchRequest: %s", err.Error())}
//	}
//	return
//}
//func (searchrequest *SearchRequest) writeSearchRequestComponents(bytes Bytes) (err error) {
//	searchrequest.baseObject, err = writeLDAPDN(bytes)
//	if err != nil {
//		return
//	}
//	searchrequest.scope, err = writeENUMERATED(bytes, EnumeratedSearchRequestScope)
//	if err != nil {
//		return
//	}
//	searchrequest.derefAliases, err = writeENUMERATED(bytes, EnumeratedSearchRequestDerefAliases)
//	if err != nil {
//		return
//	}
//	searchrequest.sizeLimit, err = writePositiveINTEGER(bytes)
//	if err != nil {
//		return
//	}
//	searchrequest.timeLimit, err = writePositiveINTEGER(bytes)
//	if err != nil {
//		return
//	}
//	searchrequest.typesOnly, err = writeBOOLEAN(bytes)
//	if err != nil {
//		return
//	}
//	searchrequest.filter, err = writeFilter(bytes)
//	if err != nil {
//		return
//	}
//	searchrequest.attributes, err = writeAttributeSelection(bytes)
//	if err != nil {
//		return
//	}
//	return
//}
//
////
////        AttributeSelection ::= SEQUENCE OF selector LDAPString
////                       -- The LDAPString is constrained to
////                       -- <attributeSelector> in Section 4.5.1.8
//func writeAttributeSelection(bytes Bytes) (attributeSelection AttributeSelection, err error) {
//	err = bytes.ReadSubBytes(classUniversal, tagSequence, attributeSelection.writeAttributeSelectionComponents)
//	return
//}
//func (attributeSelection *AttributeSelection) writeAttributeSelectionComponents(bytes Bytes) (err error) {
//	for bytes.HasMoreData() {
//		var ldapstring LDAPString
//		ldapstring, err = writeLDAPString(bytes)
//		// @TOTO: check <attributeSelector> in Section 4.5.1.8
//		if err != nil {
//			return
//		}
//		*attributeSelection = append(*attributeSelection, ldapstring)
//	}
//	return
//}
//
////
////        Filter ::= CHOICE {
////             and             [0] SET SIZE (1..MAX) OF filter Filter,
////             or              [1] SET SIZE (1..MAX) OF filter Filter,
////             not             [2] Filter,
////             equalityMatch   [3] AttributeValueAssertion,
////
////
////
////Sermersheim                 Standards Track                    [Page 57]
////
////
////RFC 4511                         LDAPv3                        June 2006
////
////
////             substrings      [4] SubstringFilter,
////             greaterOrEqual  [5] AttributeValueAssertion,
////             lessOrEqual     [6] AttributeValueAssertion,
////             present         [7] AttributeDescription,
////             approxMatch     [8] AttributeValueAssertion,
////             extensibleMatch [9] MatchingRuleAssertion,
////             ...  }
//func writeFilter(bytes Bytes) (filter Filter, err error) {
//	var tagAndLength TagAndLength
//	tagAndLength, err = bytes.PreviewTagAndLength()
//	if err != nil {
//		return
//	}
//	err = tagAndLength.ExpectClass(classContextSpecific)
//	if err != nil {
//		return
//	}
//	switch tagAndLength.Tag {
//	case TagFilterAnd:
//		filter, err = writeFilterAnd(bytes)
//	case TagFilterOr:
//		filter, err = writeFilterOr(bytes)
//	case TagFilterNot:
//		filter, err = writeFilterNot(bytes)
//	case TagFilterEqualityMatch:
//		filter, err = writeFilterEqualityMatch(bytes)
//	case TagFilterSubstrings:
//		filter, err = writeFilterSubstrings(bytes)
//	case TagFilterGreaterOrEqual:
//		filter, err = writeFilterGreaterOrEqual(bytes)
//	case TagFilterLessOrEqual:
//		filter, err = writeFilterLessOrEqual(bytes)
//	case TagFilterPresent:
//		filter, err = writeFilterPresent(bytes)
//	case TagFilterApproxMatch:
//		filter, err = writeFilterApproxMatch(bytes)
//	case TagFilterExtensibleMatch:
//		filter, err = writeFilterExtensibleMatch(bytes)
//	default:
//		err = LdapError{fmt.Sprintf("Invalid tag value for filter: %d.", tagAndLength.Tag)}
//	}
//	if err != nil {
//		err = LdapError{fmt.Sprintf("writeFilter: %s.", err.Error())}
//	}
//	return
//}
//
////             and             [0] SET SIZE (1..MAX) OF filter Filter,
//func writeFilterAnd(bytes Bytes) (filterand FilterAnd, err error) {
//	err = bytes.ReadSubBytes(classContextSpecific, TagFilterAnd, filterand.writeFilterAndComponents)
//	return
//}
//func (filterand *FilterAnd) writeFilterAndComponents(bytes Bytes) (err error) {
//	for bytes.HasMoreData() {
//		var filter Filter
//		filter, err = writeFilter(bytes)
//		if err != nil {
//			return
//		}
//		*filterand = append(*filterand, filter)
//	}
//	if len(*filterand) == 0 {
//		err = LdapError{"writeFilterAnd: expecting at least one Filter"}
//	}
//	return
//}
//
////             or              [1] SET SIZE (1..MAX) OF filter Filter,
//func writeFilterOr(bytes Bytes) (filteror FilterOr, err error) {
//	err = bytes.ReadSubBytes(classContextSpecific, TagFilterOr, filteror.writeFilterOrComponents)
//	return
//}
//
//func (filteror *FilterOr) writeFilterOrComponents(bytes Bytes) (err error) {
//	for bytes.HasMoreData() {
//		var filter Filter
//		filter, err = writeFilter(bytes)
//		if err != nil {
//			return
//		}
//		*filteror = append(*filteror, filter)
//	}
//	if len(*filteror) == 0 {
//		err = LdapError{"writeFilterOr: expecting at least one Filter"}
//	}
//	return
//}
//
////             not             [2] Filter,
//func writeFilterNot(bytes Bytes) (filternot FilterNot, err error) {
//	err = bytes.ReadSubBytes(classContextSpecific, TagFilterNot, filternot.writeFilterNotComponents)
//	return
//}
//
//func (filternot *FilterNot) writeFilterNotComponents(bytes Bytes) (err error) {
//	var tagAndLength TagAndLength
//	tagAndLength, err = bytes.ParseTagAndLength()
//	if err != nil {
//		return
//	}
//	err = tagAndLength.ExpectTag(TagFilterNot)
//	if err != nil {
//		return
//	}
//	filternot.Filter, err =  writeFilter(bytes)
//	if err != nil {
//		return
//	}
//	return
//}
//
////             equalityMatch   [3] AttributeValueAssertion,
//func writeFilterEqualityMatch(bytes Bytes) (ret FilterEqualityMatch, err error) {
//	var attributevalueassertion AttributeValueAssertion
//	attributevalueassertion, err = writeTaggedAttributeValueAssertion(bytes, classContextSpecific, TagFilterEqualityMatch)
//	if err != nil {
//		return
//	}
//	ret = FilterEqualityMatch(attributevalueassertion)
//	return
//}
//
////             substrings      [4] SubstringFilter,
//func writeFilterSubstrings(bytes Bytes) (filtersubstrings FilterSubstrings, err error) {
//	var substringfilter SubstringFilter
//	substringfilter, err = writeTaggedSubstringFilter(bytes, classContextSpecific, TagFilterSubstrings)
//	if err != nil {
//		return
//	}
//	filtersubstrings = FilterSubstrings(substringfilter)
//	return
//}
//
////             greaterOrEqual  [5] AttributeValueAssertion,
//func writeFilterGreaterOrEqual(bytes Bytes) (ret FilterGreaterOrEqual, err error) {
//	var attributevalueassertion AttributeValueAssertion
//	attributevalueassertion, err = writeTaggedAttributeValueAssertion(bytes, classContextSpecific, TagFilterGreaterOrEqual)
//	if err != nil {
//		return
//	}
//	ret = FilterGreaterOrEqual(attributevalueassertion)
//	return
//}
//
////             lessOrEqual     [6] AttributeValueAssertion,
//func writeFilterLessOrEqual(bytes Bytes) (ret FilterLessOrEqual, err error) {
//	var attributevalueassertion AttributeValueAssertion
//	attributevalueassertion, err = writeTaggedAttributeValueAssertion(bytes, classContextSpecific, TagFilterLessOrEqual)
//	if err != nil {
//		return
//	}
//	ret = FilterLessOrEqual(attributevalueassertion)
//	return
//}
//
////             present         [7] AttributeDescription,
//func writeFilterPresent(bytes Bytes) (ret FilterPresent, err error) {
//	var attributedescription AttributeDescription
//	attributedescription, err = writeTaggedAttributeDescription(bytes, classContextSpecific, TagFilterPresent)
//	if err != nil {
//		return ret, LdapError{fmt.Sprintf("writeFilterPresent: %s", err.Error())}
//	}
//	ret = FilterPresent(attributedescription)
//	return
//}
//
////             approxMatch     [8] AttributeValueAssertion,
//func writeFilterApproxMatch(bytes Bytes) (ret FilterApproxMatch, err error) {
//	var attributevalueassertion AttributeValueAssertion
//	attributevalueassertion, err = writeTaggedAttributeValueAssertion(bytes, classContextSpecific, TagFilterApproxMatch)
//	if err != nil {
//		return
//	}
//	ret = FilterApproxMatch(attributevalueassertion)
//	return
//}
//
////             extensibleMatch [9] MatchingRuleAssertion,
//func writeFilterExtensibleMatch(bytes Bytes) (filterextensiblematch FilterExtensibleMatch, err error) {
//	var matchingruleassertion MatchingRuleAssertion
//	matchingruleassertion, err = writeTaggedMatchingRuleAssertion(bytes, classContextSpecific, TagFilterExtensibleMatch)
//	if err != nil {
//		return
//	}
//	filterextensiblematch = FilterExtensibleMatch(matchingruleassertion)
//	return
//}
//
////
////        SubstringFilter ::= SEQUENCE {
////             type           AttributeDescription,
////             substrings     SEQUENCE SIZE (1..MAX) OF substring CHOICE {
////                  initial [0] AssertionValue,  -- can occur at most once
////                  any     [1] AssertionValue,
////                  final   [2] AssertionValue } -- can occur at most once
////             }
//func writeTaggedSubstringFilter(bytes Bytes, class int, tag int) (substringfilter SubstringFilter, err error) {
//	err = bytes.ReadSubBytes(class, tag, substringfilter.writeSubstringFilterComponents)
//	return
//}
//func (substringfilter SubstringFilter) writeSubstringFilterComponents(bytes Bytes) (err error) {
//	substringfilter.type_, err = writeAttributeDescription(bytes)
//	if err != nil {
//		return
//	}
//	substringfilter.substrings, err = writeSubstringFilterSubstrings(bytes)
//	if err != nil {
//		return
//	}
//	return
//}
//
//func writeSubstringFilterSubstrings(bytes Bytes) (substrings SubstringFilterSubstrings, err error) {
//	err = bytes.ReadSubBytes(classUniversal, tagSequence, substrings.writeSubstringFilterSubstringsComponents)
//	return
//}
//
//func (substrings *SubstringFilterSubstrings) writeSubstringFilterSubstringsComponents(bytes Bytes) (err error) {
//	var foundInitial = 0
//	var foundFinal = 0
//	var tagAndLength TagAndLength
//	for bytes.HasMoreData() {
//		tagAndLength, err = bytes.PreviewTagAndLength()
//		if err != nil {
//			return
//		}
//		var assertionvalue AssertionValue
//		switch tagAndLength.Tag {
//			case TagSubstringInitial:
//				foundInitial++
//				if foundInitial > 1 {
//					return LdapError{"writeSubstring: initial can occur at most once"}
//				}
//				assertionvalue, err = writeTaggedAssertionValue(bytes, classContextSpecific, TagSubstringInitial)
//				if err != nil {
//					return
//				}
//				*substrings = append(*substrings, SubstringInitial(assertionvalue))
//			case TagSubstringAny:
//				assertionvalue, err = writeTaggedAssertionValue(bytes, classContextSpecific, TagSubstringAny)
//				if err != nil {
//					return
//				}
//				*substrings = append(*substrings, SubstringAny(assertionvalue))
//			case TagSubstringFinal:
//				foundFinal++
//				if foundFinal > 1 {
//					return LdapError{"writeSubstring: final can occur at most once"}
//				}
//				assertionvalue, err = writeTaggedAssertionValue(bytes, classContextSpecific, TagSubstringFinal)
//				if err != nil {
//					return
//				}
//				*substrings = append(*substrings, SubstringFinal(assertionvalue))
//			default:
//				return LdapError{fmt.Sprintf("writeSubstring: invalid tag %d", tagAndLength.Tag)}
//		}
//	}
//	if len(*substrings) == 0 {
//		err = LdapError{"writeSubstringFilterSubstrings: expecting at least one substring"}
//	}
//	return
//}
//
////
////        MatchingRuleAssertion ::= SEQUENCE {
////             matchingRule    [1] MatchingRuleId OPTIONAL,
////             type            [2] AttributeDescription OPTIONAL,
////             matchValue      [3] AssertionValue,
////             dnAttributes    [4] BOOLEAN DEFAULT FALSE }
//func writeTaggedMatchingRuleAssertion(bytes Bytes, class int, tag int) (ret MatchingRuleAssertion, err error) {
//	err = bytes.ReadSubBytes(class, tag, ret.writeMatchingRuleAssertionComponents)
//	return
//}
//func (matchingruleassertion MatchingRuleAssertion) writeMatchingRuleAssertionComponents(bytes Bytes) (err error) {
//	err = matchingruleassertion.writeMatchingRule(bytes)
//	if err != nil {
//		return LdapError{fmt.Sprintf("writeMatchingRuleAssertionComponents: %s", err.Error())}
//	}
//	err = matchingruleassertion.writeType(bytes)
//	if err != nil {
//		return LdapError{fmt.Sprintf("writeMatchingRuleAssertionComponents: %s", err.Error())}
//	}
//	matchingruleassertion.matchValue, err = writeTaggedAssertionValue(bytes, classContextSpecific, TagMatchingRuleAssertionMatchValue)
//	if err != nil {
//		return LdapError{fmt.Sprintf("writeMatchingRuleAssertionComponents: %s", err.Error())}
//	}
//	matchingruleassertion.dnAttributes, err = writeTaggedBOOLEAN(bytes, classContextSpecific, TagMatchingRuleAssertionDnAttributes)
//	if err != nil {
//		return LdapError{fmt.Sprintf("writeMatchingRuleAssertionComponents: %s", err.Error())}
//	}
//	return
//}
//func (matchingruleassertion MatchingRuleAssertion) writeMatchingRule(bytes Bytes) (err error) {
//	var tagAndLength TagAndLength
//	tagAndLength, err = bytes.PreviewTagAndLength()
//	if err != nil {
//		return LdapError{fmt.Sprintf("writeMatchingRuleAssertionMatchingRule: %s", err.Error())}
//	}
//	if tagAndLength.Tag == TagMatchingRuleAssertionMatchingRule {
//		var matchingRule MatchingRuleId
//		matchingRule, err = writeTaggedMatchingRuleId(bytes, classContextSpecific, TagMatchingRuleAssertionMatchingRule)
//		if err != nil {
//			return LdapError{fmt.Sprintf("writeMatchingRuleAssertionMatchingRule: %s", err.Error())}
//		}
//		matchingruleassertion.matchingRule = &matchingRule
//	}
//	return
//}
//func (matchingruleassertion MatchingRuleAssertion) writeType(bytes Bytes) (err error) {
//	var tagAndLength TagAndLength
//	tagAndLength, err = bytes.PreviewTagAndLength()
//	if err != nil {
//		return LdapError{fmt.Sprintf("writeMatchingRuleAssertionType: %s", err.Error())}
//	}
//	if tagAndLength.Tag == TagMatchingRuleAssertionType {
//		var attributedescription AttributeDescription
//		attributedescription, err = writeTaggedAttributeDescription(bytes, classContextSpecific, TagMatchingRuleAssertionType)
//		if err != nil {
//			return LdapError{fmt.Sprintf("writeMatchingRuleAssertionType: %s", err.Error())}
//		}
//		matchingruleassertion.type_ = &attributedescription
//	}
//	return
//}
//
////
////        SearchResultEntry ::= [APPLICATION 4] SEQUENCE {
////             objectName      LDAPDN,
////             attributes      PartialAttributeList }
//func writeSearchResultEntry(bytes Bytes) (searchresultentry SearchResultEntry, err error){
//	err = bytes.ReadSubBytes(classApplication, TagSearchResultEntry, searchresultentry.writeSearchResultEntryComponents)
//	return
//}
//func (searchresultentry *SearchResultEntry) writeSearchResultEntryComponents(bytes Bytes) (err error){
//	searchresultentry.objectName, err = writeLDAPDN(bytes)
//	if err != nil {
//		return
//	}
//	searchresultentry.attributes, err = writePartialAttributeList(bytes)
//	if err != nil {
//		return
//	}
//	return
//}
////
////        PartialAttributeList ::= SEQUENCE OF
////                             partialAttribute PartialAttribute
//func writePartialAttributeList(bytes Bytes) (ret PartialAttributeList, err error){
//	ret = PartialAttributeList(make([]PartialAttribute, 0, 10))
//	err = bytes.ReadSubBytes(classUniversal, tagSequence, ret.writePartialAttributeListComponents)
//	return ret, err
//}
//func (partialattributelist *PartialAttributeList) writePartialAttributeListComponents(bytes Bytes) (err error){
//	for bytes.HasMoreData() {
//		var partialattribute PartialAttribute
//		partialattribute, err = writePartialAttribute(bytes)
//		if err != nil {
//			return
//		}
//		*partialattributelist = append(*partialattributelist, partialattribute)
//	}
//	return
//}
//
////
////        SearchResultReference ::= [APPLICATION 19] SEQUENCE
////                                  SIZE (1..MAX) OF uri URI
//func writeSearchResultReference(bytes Bytes) (ret SearchResultReference, err error){
//	err = bytes.ReadSubBytes(classApplication, TagSearchResultReference, ret.writeComponents)
//	return
//}
//func (s *SearchResultReference) writeComponents(bytes Bytes) (err error){
//	for bytes.HasMoreData() {
//		var uri URI
//		uri, err = writeURI(bytes)
//		if err != nil {
//			return
//		}
//		*s = append(*s, uri)
//	}
//	if(len(*s) == 0){
//		err = LdapError{"SearchResultReference: expecting at least one URI"}
//	}
//	return
//}
////
////        SearchResultDone ::= [APPLICATION 5] LDAPResult
//func writeSearchResultDone(bytes Bytes) (ret SearchResultDone, err error){
//	var ldapresult LDAPResult
//	ldapresult, err = writeTaggedLDAPResult(bytes, classApplication, TagSearchResultDone)
//	if err != nil {
//		return
//	}
//	ret = SearchResultDone(ldapresult)
//	return
//}
////
////        ModifyRequest ::= [APPLICATION 6] SEQUENCE {
////             object          LDAPDN,
////             changes         SEQUENCE OF change SEQUENCE {
////                  operation       ENUMERATED {
////                       add     (0),
////                       delete  (1),
////                       replace (2),
////                       ...  },
////                  modification    PartialAttribute } }
//func writeModifyRequest(bytes Bytes) (ret ModifyRequest, err error){
//	err = bytes.ReadSubBytes(classApplication, TagModifyRequest, ret.writeComponents)
//	return
//}
//func (m *ModifyRequest) writeComponents(bytes Bytes) (err error){
//	m.object, err = writeLDAPDN(bytes)
//	if err != nil {
//		return
//	}
//	err = bytes.ReadSubBytes(classUniversal, tagSequence, m.writeChanges)
//	return
//}
//func (m *ModifyRequest) writeChanges(bytes Bytes) (err error) {
//	for bytes.HasMoreData(){
//		var c ModifyRequestChange
//		c, err = writeModifyRequestChange(bytes)
//		if err != nil {
//			return
//		}
//		m.changes = append(m.changes, c)
//	}
//	return
//}
//func writeModifyRequestChange(bytes Bytes) (ret ModifyRequestChange, err error){
//	err = bytes.ReadSubBytes(classUniversal, tagSequence, ret.writeComponents)
//	return
//}
//func (m *ModifyRequestChange) writeComponents(bytes Bytes) (err error){
//	m.operation, err = writeENUMERATED(bytes, EnumeratedModifyRequestChangeOpration)
//	if err != nil {
//		return
//	}
//	m.modification, err = writePartialAttribute(bytes)
//	return
//}
////
////        ModifyResponse ::= [APPLICATION 7] LDAPResult
//func writeModifyResponse(bytes Bytes) (ret ModifyResponse, err error){
//	var res LDAPResult
//	res, err = writeTaggedLDAPResult(bytes, classApplication, TagModifyResponse)
//	if err != nil {
//		return
//	}
//	ret = ModifyResponse(res)
//	return
//}
////
////
////
////
////
////
////Sermersheim                 Standards Track                    [Page 58]
////
////
////RFC 4511                         LDAPv3                        June 2006
////
////
////        AddRequest ::= [APPLICATION 8] SEQUENCE {
////             entry           LDAPDN,
////             attributes      AttributeList }
//func writeAddRequest(bytes Bytes) (ret AddRequest, err error){
//	err = bytes.ReadSubBytes(classApplication, TagAddRequest, ret.writeComponents)
//	return
//}
//func (req *AddRequest) writeComponents(bytes Bytes) (err error){
//	req.entry, err = writeLDAPDN(bytes)
//	if err != nil {
//		return
//	}
//	req.attributes, err = writeAttributeList(bytes)
//	return
//}
//
////
////        AttributeList ::= SEQUENCE OF attribute Attribute
//func writeAttributeList(bytes Bytes) (ret AttributeList, err error){
//	err = bytes.ReadSubBytes(classUniversal, tagSequence, ret.writeComponents)
//	return
//}
//func (list *AttributeList) writeComponents(bytes Bytes) (err error){
//	for bytes.HasMoreData(){
//		var attr Attribute
//		attr, err = writeAttribute(bytes)
//		if err != nil {
//			return
//		}
//		*list = append(*list, attr)
//	}
//	return
//}
////
////        AddResponse ::= [APPLICATION 9] LDAPResult
//func writeAddResponse(bytes Bytes) (ret AddResponse, err error){
//	var res LDAPResult
//	res, err = writeTaggedLDAPResult(bytes, classApplication, TagAddResponse)
//	if err != nil {
//		return
//	}
//	ret = AddResponse(res)
//	return
//}
////
////        DelRequest ::= [APPLICATION 10] LDAPDN
//func writeDelRequest(bytes Bytes) (ret DelRequest, err error){
//	var res LDAPDN
//	res, err = writeTaggedLDAPDN(bytes, classApplication, TagDelRequest)
//	if err != nil {
//		return
//	}
//	ret = DelRequest(res)
//	return
//}
//
////
////        DelResponse ::= [APPLICATION 11] LDAPResult
//func writeDelResponse(bytes Bytes) (ret DelResponse, err error){
//	var res LDAPResult
//	res, err = writeTaggedLDAPResult(bytes, classApplication, TagDelResponse)
//	if err != nil {
//		return
//	}
//	ret = DelResponse(res)
//	return
//}
//
////
////        ModifyDNRequest ::= [APPLICATION 12] SEQUENCE {
////             entry           LDAPDN,
////             newrdn          RelativeLDAPDN,
////             deleteoldrdn    BOOLEAN,
////             newSuperior     [0] LDAPDN OPTIONAL }
//func writeModifyDNRequest(bytes Bytes) (ret ModifyDNRequest, err error){
//	err = bytes.ReadSubBytes(classApplication, TagModifyDNRequest, ret.writeComponents)
//	return
//}
//func (req *ModifyDNRequest) writeComponents(bytes Bytes) (err error){
//	req.entry, err = writeLDAPDN(bytes)
//	if err != nil {
//		return
//	}
//	req.newrdn, err = writeRelativeLDAPDN(bytes)
//	if err != nil {
//		return
//	}
//	req.deleteoldrdn, err = writeBOOLEAN(bytes)
//	if err != nil {
//		return
//	}
//	if bytes.HasMoreData() {
//		var ldapdn LDAPDN
//		ldapdn, err = writeTaggedLDAPDN(bytes, classContextSpecific, TagModifyDNRequestNewSuperior)
//		if err != nil {
//			return
//		}
//		req.newSuperior = &ldapdn
//	}
//	return
//}
//
////
////        ModifyDNResponse ::= [APPLICATION 13] LDAPResult
//func writeModifyDNResponse(bytes Bytes) (ret ModifyDNResponse, err error){
//	var res LDAPResult
//	res, err = writeTaggedLDAPResult(bytes, classApplication, TagModifyDNResponse)
//	if err != nil {
//		return
//	}
//	ret = ModifyDNResponse(res)
//	return
//}
//
////
////        CompareRequest ::= [APPLICATION 14] SEQUENCE {
////             entry           LDAPDN,
////             ava             AttributeValueAssertion }
//func writeCompareRequest(bytes Bytes) (ret CompareRequest, err error){
//	err = bytes.ReadSubBytes(classApplication, TagCompareRequest, ret.writeComponents)
//	return
//}
//func (req *CompareRequest) writeComponents(bytes Bytes) (err error){
//	req.entry, err = writeLDAPDN(bytes)
//	if err != nil {
//		return
//	}
//	req.ava, err = writeAttributeValueAssertion(bytes)
//	return
//}
//
////
////        CompareResponse ::= [APPLICATION 15] LDAPResult
//func writeCompareResponse(bytes Bytes) (ret CompareResponse, err error){
//	var res LDAPResult
//	res, err = writeTaggedLDAPResult(bytes, classApplication, TagCompareResponse)
//	if err != nil {
//		return
//	}
//	ret = CompareResponse(res)
//	return
//}
//
////
////        AbandonRequest ::= [APPLICATION 16] MessageID
//func writeAbandonRequest(bytes Bytes) (ret AbandonRequest, err error){
//	var mes MessageID
//	mes, err = writeTaggedMessageID(bytes, classApplication, TagAbandonRequest)
//	if err != nil {
//		return
//	}
//	ret = AbandonRequest(mes)
//	return
//}
//
////
////        ExtendedRequest ::= [APPLICATION 23] SEQUENCE {
////             requestName      [0] LDAPOID,
////             requestValue     [1] OCTET STRING OPTIONAL }
//func writeExtendedRequest(bytes Bytes) (ret ExtendedRequest, err error){
//	err = bytes.ReadSubBytes(classApplication, TagExtendedRequest, ret.writeComponents)
//	return
//}
//func (req *ExtendedRequest) writeComponents(bytes Bytes) (err error){
//	req.requestName, err = writeTaggedLDAPOID(bytes, classContextSpecific, TagExtendedRequestName)
//	if err != nil {
//		return
//	}
//	if bytes.HasMoreData() {
//		var str OCTETSTRING
//		str, err = writeTaggedOCTETSTRING(bytes, classContextSpecific, TagExtendedRequestValue)
//		if err != nil {
//			return
//		}
//		req.requestValue = &str
//	}
//	return
//}
//
////
////        ExtendedResponse ::= [APPLICATION 24] SEQUENCE {
////             COMPONENTS OF LDAPResult,
////             responseName     [10] LDAPOID OPTIONAL,
////             responseValue    [11] OCTET STRING OPTIONAL }
//func writeExtendedResponse(bytes Bytes) (ret ExtendedResponse, err error){
//	err = bytes.ReadSubBytes(classApplication, TagExtendedResponse, ret.writeComponents)
//	return
//}
//func (res *ExtendedResponse) writeComponents(bytes Bytes) (err error){
//	res.writeLDAPResultComponents(bytes)
//	if bytes.HasMoreData() {
//		var oid LDAPOID
//		oid, err = writeTaggedLDAPOID(bytes, classContextSpecific, TagExtendedResponseName)
//		if err != nil {
//			return
//		}
//		res.responseName = &oid
//	}
//	if bytes.HasMoreData() {
//		var str OCTETSTRING
//		str, err = writeTaggedOCTETSTRING(bytes, classContextSpecific, TagExtendedResponseValue)
//		if err != nil {
//			return
//		}
//		res.responseValue = &str
//	}
//	return
//}
//
////
////        IntermediateResponse ::= [APPLICATION 25] SEQUENCE {
////             responseName     [0] LDAPOID OPTIONAL,
////             responseValue    [1] OCTET STRING OPTIONAL }
//func writeIntermediateResponse(bytes Bytes) (ret IntermediateResponse, err error){
//	err = bytes.ReadSubBytes(classApplication, TagIntermediateResponse, ret.writeComponents)
//	return
//}
//func (res *IntermediateResponse) writeComponents(bytes Bytes) (err error){
//	if bytes.HasMoreData() {
//		var oid LDAPOID
//		oid, err = writeTaggedLDAPOID(bytes, classContextSpecific, TagIntermediateResponseName)
//		if err != nil {
//			return
//		}
//		res.responseName = &oid
//	}
//	if bytes.HasMoreData() {
//		var str OCTETSTRING
//		str, err = writeTaggedOCTETSTRING(bytes, classContextSpecific, TagIntermediateResponseValue)
//		if err != nil {
//			return
//		}
//		res.responseValue = &str
//	}
//	return
//}
//
////
////        END
////
