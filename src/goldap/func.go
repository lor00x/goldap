package goldap

import (
	"fmt"
	"goldap/asn1"
)

func NewOCTETSTRING(bytes []byte) *OCTETSTRING {
	octetstring := OCTETSTRING(bytes)
	return &octetstring
}
func ReadOCTETSTRING(bytes *Bytes) *OCTETSTRING {
	tagAndLength := bytes.ParseTagAndLength()
	tagAndLength.Expect(asn1.ClassUniversal, asn1.TagOctetString, asn1.IsNotCompound)
	octetstring := bytes.ParseOCTETSTRING(tagAndLength.GetLength())
	return &octetstring
}
func ReadINTEGER(bytes *Bytes) INTEGER {
	tagAndLength := bytes.ParseTagAndLength()
	tagAndLength.Expect(asn1.ClassUniversal, asn1.TagInteger, asn1.IsNotCompound)
	integer := bytes.ParseInt32(tagAndLength.GetLength())
	return INTEGER(integer)
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
func (m *LDAPMessage) ReadSequenceTag(bytes *Bytes) {
	tagAndLength := bytes.ParseTagAndLength()
	tagAndLength.Expect(asn1.ClassUniversal, asn1.TagSequence, asn1.IsCompound)
}

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

func ReadLDAPMessage(bytes *Bytes) (message *LDAPMessage) {
	defer func() {
		if e := recover(); e != nil {
			message = nil
			fmt.Printf("ERROR at offset %d: %v\n", bytes.offset, e)
		}
		return
	}()
	message = NewLDAPMessage()
	message.ReadSequenceTag(bytes)
	message.ReadMessageID(bytes)
	message.ReadProtocolOp(bytes)
	//	message.ReadControls(bytes)
	return
}

//        MessageID ::= INTEGER (0 ..  maxInt)
//
//        maxInt INTEGER ::= 2147483647 -- (2^^31 - 1) --
//
func (m *LDAPMessage) ReadMessageID(bytes *Bytes) {
	tagAndLength := bytes.ParseTagAndLength()
	tagAndLength.Expect(asn1.ClassUniversal, asn1.TagInteger, asn1.IsNotCompound)
	id := INTEGER(bytes.ParseInt32(tagAndLength.GetLength()))
	if !(id >= 0 && id <= maxInt) {
		panic(fmt.Sprintf("Invalid MessageID ! Expected value between 0 and %d. Got %d.", maxInt, id))
	}
	m.messageID = MessageID(id)
}
func (m *LDAPMessage) ReadProtocolOp(bytes *Bytes) (err error) {
	tagAndLength := bytes.ParseTagAndLength()
	// tagAndLength, offset := asn1.ParseTagAndLength(bytes.bytes, bytes.offset)
	if asn1.ClassApplication != tagAndLength.GetClass() {
		panic(fmt.Sprintf("Invalid tag class for protocolOp. Expected %d. Got %d.", asn1.ClassApplication, tagAndLength.GetClass()))
	}
	// bytes.offset = offset
	switch tagAndLength.GetTag() {
	case TagBindRequest:
		m.protocolOp, err = ReadBindRequest(bytes)
	default:
		panic(fmt.Sprintf("Invalid tag value for protocolOp. Got %d.", tagAndLength.GetTag()))
	}
	return
}

//        LDAPString ::= OCTET STRING -- UTF-8 encoded,
//                                    -- [ISO10646] characters
func ReadLDAPString(bytes *Bytes) LDAPString {
	tagAndLength := bytes.ParseTagAndLength()
	tagAndLength.Expect(asn1.ClassUniversal, asn1.TagOctetString, asn1.IsNotCompound)
	return LDAPString(bytes.ParseUTF8STRING(tagAndLength.GetLength()))
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
//
//        LDAPDN ::= LDAPString -- Constrained to <distinguishedName>
//                              -- [RFC4514]
func ReadLDAPDN(bytes *Bytes) LDAPDN {
	// @TODO: check RFC4514
	return LDAPDN(ReadLDAPString(bytes))
}

//
//        RelativeLDAPDN ::= LDAPString -- Constrained to <name-component>
//                                      -- [RFC4514]
func ReadRelativeLDAPDN(bytes *Bytes) RelativeLDAPDN {
	// @TODO: check RFC4514
	return RelativeLDAPDN(ReadLDAPString(bytes))
}

//
//        AttributeDescription ::= LDAPString
//                                -- Constrained to <attributedescription>
//                                -- [RFC4512]
func ReadAttributeDescription(bytes *Bytes) AttributeDescription {
	// @TODO: check RFC4512
	return AttributeDescription(ReadLDAPString(bytes))
}

//
//        AttributeValue ::= OCTET STRING
//
//        AttributeValueAssertion ::= SEQUENCE {
//             attributeDesc   AttributeDescription,
//             assertionValue  AssertionValue }
//
//        AssertionValue ::= OCTET STRING
//
//        PartialAttribute ::= SEQUENCE {
//             type       AttributeDescription,
//             vals       SET OF value AttributeValue }
//
//        Attribute ::= PartialAttribute(WITH COMPONENTS {
//             ...,
//             vals (SIZE(1..MAX))})
//
//        MatchingRuleId ::= LDAPString
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
//
//        Referral ::= SEQUENCE SIZE (1..MAX) OF uri URI
//
//        URI ::= LDAPString     -- limited to characters permitted in
//                               -- URIs
//
//        Controls ::= SEQUENCE OF control Control
//
//        Control ::= SEQUENCE {
//             controlType             LDAPOID,
//             criticality             BOOLEAN DEFAULT FALSE,
//             controlValue            OCTET STRING OPTIONAL }
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
func ReadBindRequest(bytes *Bytes) (bindrequest *BindRequest, err error) {
	bindrequest = &BindRequest{}
	err = bytes.ParseSequence(asn1.ClassApplication, TagBindRequest,
		func(subBytes *Bytes) {
			bindrequest.version = ReadINTEGER(subBytes)
			if !(bindrequest.version >= BindRequestVersionMin && bindrequest.version <= BindRequestVersionMax){
				err = SyntaxError{fmt.Sprintf("BindRequest: invalid version %d. Must be between 1 and 127", bindrequest.version)}
				return
			} 
			bindrequest.name = ReadLDAPDN(subBytes)
			bindrequest.authentication, err = ReadAuthenticationChoice(subBytes)
		},
	)
	return
}

//
//        AuthenticationChoice ::= CHOICE {
//             simple                  [0] OCTET STRING,
//                                     -- 1 and 2 reserved
//             sasl                    [3] SaslCredentials,
//             ...  }
func ReadAuthenticationChoice(bytes *Bytes) (ret interface{}, err error) {
	tagAndLength := bytes.PreviewTagAndLength()
	if asn1.ClassContextSpecific != tagAndLength.GetClass() {
		err = SyntaxError{fmt.Sprintf("Invalid tag class for AuthenticationChoice. Expected %d. Got %d.", asn1.ClassApplication, tagAndLength.GetClass())}
	}
	switch tagAndLength.GetTag() {
	case TagAuthenticationChoiceSimple:
		ret = ReadOCTETSTRING(bytes)
	case TagAuthenticationChoiceSaslCredentials:
		ret, err = ReadSaslCredentials(bytes)
	default:
		panic(fmt.Sprintf("Invalid tag value for AuthenticationChoice. Got %d.", tagAndLength.GetTag()))
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
	err = bytes.ParseSequence(asn1.ClassContextSpecific, TagAuthenticationChoiceSaslCredentials,
		func(subBytes *Bytes) {
			authentication.mechanism = ReadLDAPString(subBytes)
			if subBytes.HasMoreData() {
				authentication.credentials = ReadOCTETSTRING(subBytes)
			}
		},
	)
	return
}

//        BindResponse ::= [APPLICATION 1] SEQUENCE {
//             COMPONENTS OF LDAPResult,
//             serverSaslCreds    [7] OCTET STRING OPTIONAL }
//
//        UnbindRequest ::= [APPLICATION 2] NULL
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
//
//        AttributeSelection ::= SEQUENCE OF selector LDAPString
//                       -- The LDAPString is constrained to
//                       -- <attributeSelector> in Section 4.5.1.8
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
//
//        SubstringFilter ::= SEQUENCE {
//             type           AttributeDescription,
//             substrings     SEQUENCE SIZE (1..MAX) OF substring CHOICE {
//                  initial [0] AssertionValue,  -- can occur at most once
//                  any     [1] AssertionValue,
//                  final   [2] AssertionValue } -- can occur at most once
//             }
//
//        MatchingRuleAssertion ::= SEQUENCE {
//             matchingRule    [1] MatchingRuleId OPTIONAL,
//             type            [2] AttributeDescription OPTIONAL,
//             matchValue      [3] AssertionValue,
//             dnAttributes    [4] BOOLEAN DEFAULT FALSE }
//
//        SearchResultEntry ::= [APPLICATION 4] SEQUENCE {
//             objectName      LDAPDN,
//             attributes      PartialAttributeList }
//
//        PartialAttributeList ::= SEQUENCE OF
//                             partialAttribute PartialAttribute
//
//        SearchResultReference ::= [APPLICATION 19] SEQUENCE
//                                  SIZE (1..MAX) OF uri URI
//
//        SearchResultDone ::= [APPLICATION 5] LDAPResult
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
