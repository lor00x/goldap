package goldap

import (
	"fmt"
	"goldap/asn1"
)

func ReadOCTETSTRING(bytes *Bytes) (ret OCTETSTRING, err error) {
	tagAndLength, err := bytes.ParseTagAndLength()
	if err != nil {
		return
	}
	err = tagAndLength.Expect(asn1.ClassUniversal, asn1.TagOctetString, asn1.IsNotCompound)
	if err != nil {
		return
	}
	return bytes.ParseOCTETSTRING(tagAndLength.GetLength())
}
func ReadINTEGER(bytes *Bytes) (ret INTEGER, err error) {
	tagAndLength, err := bytes.ParseTagAndLength()
	if err !=nil {
		return
	}
	err = tagAndLength.Expect(asn1.ClassUniversal, asn1.TagInteger, asn1.IsNotCompound)
	if err != nil {
		return
	}
	var integer int32
	integer, err = bytes.ParseInt32(tagAndLength.GetLength())
	return INTEGER(integer), err
}
func ReadBOOLEAN(bytes *Bytes) (ret BOOLEAN, err error) {
	tagAndLength, err := bytes.ParseTagAndLength()
	if err !=nil {
		return
	}
	err = tagAndLength.Expect(asn1.ClassUniversal, asn1.TagBoolean, asn1.IsNotCompound)
	if err != nil {
		return
	}
	var boolean bool
	boolean, err = bytes.ParseBoolean(tagAndLength.GetLength())
	return BOOLEAN(boolean), err
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
	err = bytes.ParseSequence(asn1.ClassUniversal, asn1.TagSequence,
		func(subBytes *Bytes) (subErr error) {
			message.messageID, subErr = ReadMessageID(subBytes)
			if subErr !=nil {
				return
			}
			message.protocolOp, subErr = ReadProtocolOp(subBytes)
			if subErr != nil {
				return
			}
			if subBytes.HasMoreData() {
				var controls Controls
				controls, subErr = ReadControls(subBytes)
				if subErr != nil {
					return
				}
				message.controls = &controls
			}
			return 
		},
	)
	return
}

//        MessageID ::= INTEGER (0 ..  maxInt)
//
//        maxInt INTEGER ::= 2147483647 -- (2^^31 - 1) --
//
func ReadMessageID(bytes *Bytes) (ret MessageID, err error){
	integer, err := ReadINTEGER(bytes)
	if !(integer >= 0 && integer <= maxInt) {
		err = SyntaxError{fmt.Sprintf("Invalid MessageID ! Expected value between 0 and %d. Got %d.", maxInt, integer)}
	}
	ret = MessageID(integer)
	return
}
func ReadProtocolOp(bytes *Bytes) (ret interface{}, err error) {
	tagAndLength, err := bytes.PreviewTagAndLength()
	if err != nil {
		return
	}
	switch tagAndLength.GetTag() {
	case TagBindRequest:
		ret, err = ReadBindRequest(bytes)
	default:
		err = SyntaxError{fmt.Sprintf("Invalid tag value for protocolOp. Got %d.", tagAndLength.GetTag())}
	}
	return
}

//        LDAPString ::= OCTET STRING -- UTF-8 encoded,
//                                    -- [ISO10646] characters
func ReadLDAPString(bytes *Bytes) (ret LDAPString, err error) {
	tagAndLength, err := bytes.ParseTagAndLength()
	if err != nil {
		return
	}
	err = tagAndLength.Expect(asn1.ClassUniversal, asn1.TagOctetString, asn1.IsNotCompound)
	if err != nil {
		return
	}
	var utf8string UTF8STRING
	utf8string, err = bytes.ParseUTF8STRING(tagAndLength.GetLength())
	if err != nil {
		return
	}
	return LDAPString(utf8string), err
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
func ReadControls(bytes *Bytes) (controls Controls, err error){
	err = bytes.ParseSequence(asn1.ClassUniversal, asn1.TagSequence,
		func(subBytes *Bytes) (subErr error) {
			tagAndLength, subErr := subBytes.ParseTagAndLength()
			if subErr != nil {
				return
			}
			subErr = tagAndLength.Expect(asn1.ClassUniversal, asn1.TagSequence, asn1.IsCompound)
			if subErr != nil {
				return
			}
			for subBytes.HasMoreData(){
				var control Control
				control, subErr = ReadControl(subBytes)
				if subErr != nil {
					return
				}
				controls = append(controls, control) 
			}
			return
		},
	)
	return
}
//
//        Control ::= SEQUENCE {
//             controlType             LDAPOID,
//             criticality             BOOLEAN DEFAULT FALSE,
//             controlValue            OCTET STRING OPTIONAL }
func ReadControl(bytes *Bytes) (control Control, err error){
	err = bytes.ParseSequence(asn1.ClassUniversal, asn1.TagSequence,
		func(subBytes *Bytes) (subErr error) {
			control.controlType, subErr = ReadLDAPOID(subBytes)
			if subErr != nil {
				return
			}
			control.criticality, subErr = ReadBOOLEAN(subBytes)
			if subErr != nil {
				return
			}
			if subBytes.HasMoreData() {
				var octetstring OCTETSTRING
				octetstring, subErr = ReadOCTETSTRING(subBytes)
				if subErr != nil {
					return
				}
				control.controlValue = &octetstring 
			}
			return
		},
	)
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
	err = bytes.ParseSequence(asn1.ClassApplication, TagBindRequest,
		func(subBytes *Bytes) (subErr error){
			bindrequest.version, subErr = ReadINTEGER(subBytes)
			if !(bindrequest.version >= BindRequestVersionMin && bindrequest.version <= BindRequestVersionMax){
				err = SyntaxError{fmt.Sprintf("BindRequest: invalid version %d. Must be between 1 and 127", bindrequest.version)}
				return
			} 
			bindrequest.name, subErr = ReadLDAPDN(subBytes)
			if subErr != nil {
				return
			}
			bindrequest.authentication, subErr = ReadAuthenticationChoice(subBytes)
			return
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
	tagAndLength, err := bytes.PreviewTagAndLength()
	if err != nil {
		return
	}
	err = tagAndLength.ExpectClass(asn1.ClassContextSpecific)
	if err != nil {
		return
	}
	switch tagAndLength.GetTag() {
	case TagAuthenticationChoiceSimple:
		ret, err = ReadAuthenticationChoiceSimple(bytes)
	case TagAuthenticationChoiceSaslCredentials:
		ret, err = ReadSaslCredentials(bytes)
	default:
		err = SyntaxError{fmt.Sprintf("Invalid tag value for AuthenticationChoice. Got %d.", tagAndLength.GetTag())}
	}
	return
}

func ReadAuthenticationChoiceSimple(bytes *Bytes) (ret OCTETSTRING, err error){
	tagAndLength, err := bytes.ParseTagAndLength()
	err = tagAndLength.ExpectClass(asn1.ClassContextSpecific)
	return bytes.ParseOCTETSTRING(tagAndLength.GetLength()) // ReadOCTETSTRING(bytes)
}

//
//        SaslCredentials ::= SEQUENCE {
//             mechanism               LDAPString,
//             credentials             OCTET STRING OPTIONAL }
//
func ReadSaslCredentials(bytes *Bytes) (authentication SaslCredentials, err error) {
	authentication = SaslCredentials{}
	err = bytes.ParseSequence(asn1.ClassContextSpecific, TagAuthenticationChoiceSaslCredentials,
		func(subBytes *Bytes) (subErr error) {
			authentication.mechanism, subErr = ReadLDAPString(subBytes)
			if subErr != nil {
				return
			}
			if subBytes.HasMoreData() {
				var credentials OCTETSTRING
				credentials, subErr = ReadOCTETSTRING(subBytes)
				if subErr != nil {
					return
				}
				authentication.credentials = &credentials
			}
			return 
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
