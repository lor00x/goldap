package goldap

import ()

type OCTETSTRING []byte
type UTF8STRING string
type INTEGER int32 // In this RFC the max INTEGER value is 2^31 - 1, so int32 is enough
type BOOLEAN bool

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
type LDAPMessage struct {
	messageID  MessageID
	protocolOp interface{}
	controls   *Controls
}

//        MessageID ::= INTEGER (0 ..  maxInt)
//
type MessageID INTEGER

//        maxInt INTEGER ::= 2147483647 -- (2^^31 - 1) --
const maxInt = INTEGER(2147483647)

//
//        LDAPString ::= OCTET STRING -- UTF-8 encoded,
//                                    -- [ISO10646] characters
type LDAPString OCTETSTRING

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
type LDAPOID OCTETSTRING

//
//        LDAPDN ::= LDAPString -- Constrained to <distinguishedName>
//                              -- [RFC4514]
type LDAPDN LDAPString

//
//        RelativeLDAPDN ::= LDAPString -- Constrained to <name-component>
//                                      -- [RFC4514]
type RelativeLDAPDN LDAPString

//
//        AttributeDescription ::= LDAPString
//                                -- Constrained to <attributedescription>
//                                -- [RFC4512]
type AttributeDescription LDAPString

//
//        AttributeValue ::= OCTET STRING
type AttributeValue OCTETSTRING

//
//        AttributeValueAssertion ::= SEQUENCE {
//             attributeDesc   AttributeDescription,
//             assertionValue  AssertionValue }
type AttributeValueAssertion struct {
	attributeDesc  AttributeDescription
	assertionValue AssertionValue
}

//
//        AssertionValue ::= OCTET STRING
type AssertionValue OCTETSTRING

//
//        PartialAttribute ::= SEQUENCE {
//             type       AttributeDescription,
//             vals       SET OF value AttributeValue }
type PartialAttribute struct {
	type_ AttributeDescription
	vals  []AttributeValue
}

//
//        Attribute ::= PartialAttribute(WITH COMPONENTS {
//             ...,
//             vals (SIZE(1..MAX))})
type Attribute struct {
	PartialAttribute
	components interface{}
}

//
//        MatchingRuleId ::= LDAPString
type MatchingRuleId LDAPString

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
type LDAPResult struct {
	resultCode        int
	matchedDN         LDAPDN
	diagnosticMessage LDAPString
	referral          *Referral
}

//        Referral ::= SEQUENCE SIZE (1..MAX) OF uri URI
type Referral []URI

//
//        URI ::= LDAPString     -- limited to characters permitted in
//                               -- URIs
type URI LDAPString

//
//        Controls ::= SEQUENCE OF control Control
type Controls []Control

//
//        Control ::= SEQUENCE {
//             controlType             LDAPOID,
//             criticality             BOOLEAN DEFAULT FALSE,
//             controlValue            OCTET STRING OPTIONAL }
type Control struct {
	controlType  LDAPOID
	criticality  BOOLEAN
	controlValue *OCTETSTRING
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
const TagBindRequest = 0
const BindRequestVersionMin = 0
const BindRequestVersionMax = 127

type BindRequest struct {
	version        INTEGER
	name           LDAPDN
	authentication AuthenticationChoice
}

//
//        AuthenticationChoice ::= CHOICE {
//             simple                  [0] OCTET STRING,
//                                     -- 1 and 2 reserved
//             sasl                    [3] SaslCredentials,
//             ...  }
const TagAuthenticationChoiceSimple = 0
const TagAuthenticationChoiceSaslCredentials = 3

type AuthenticationChoice interface{}

//
//        SaslCredentials ::= SEQUENCE {
//             mechanism               LDAPString,
//             credentials             OCTET STRING OPTIONAL }
type SaslCredentials struct {
	mechanism   LDAPString
	credentials *OCTETSTRING
}

//
//        BindResponse ::= [APPLICATION 1] SEQUENCE {
//             COMPONENTS OF LDAPResult,
//             serverSaslCreds    [7] OCTET STRING OPTIONAL }
type BindResponse struct {
	LDAPResult
	serverSaslCreds *OCTETSTRING
}

//
//        UnbindRequest ::= [APPLICATION 2] NULL
type UnbindRequest struct {
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
type SearchRequest struct {
	baseObject   LDAPDN
	scope        int
	derefAliases int
	sizeLimit    INTEGER
	timeLimit    INTEGER
	typesOnly    BOOLEAN
	filter       Filter
	attributes   AttributeSelection
}

//
//        AttributeSelection ::= SEQUENCE OF selector LDAPString
//                       -- The LDAPString is constrained to
//                       -- <attributeSelector> in Section 4.5.1.8
type AttributeSelection []LDAPString

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
type Filter struct {
	and             *[]Filter
	or              *[]Filter
	not             *Filter
	equalityMatch   *AttributeValueAssertion
	substrings      *SubstringFilter
	greaterOfEqual  *AttributeValueAssertion
	lessOrEqual     *AttributeValueAssertion
	present         *AttributeDescription
	approxMatch     *AttributeValueAssertion
	extensibleMatch *MatchingRuleAssertion
}

//
//        SubstringFilter ::= SEQUENCE {
//             type           AttributeDescription,
//             substrings     SEQUENCE SIZE (1..MAX) OF substring CHOICE {
//                  initial [0] AssertionValue,  -- can occur at most once
//                  any     [1] AssertionValue,
//                  final   [2] AssertionValue } -- can occur at most once
//             }
type SubstringFilter struct {
	type_      AttributeDescription
	substrings []struct {
		initial *AssertionValue
		any     *AssertionValue
		final   *AssertionValue
	}
}

//
//        MatchingRuleAssertion ::= SEQUENCE {
//             matchingRule    [1] MatchingRuleId OPTIONAL,
//             type            [2] AttributeDescription OPTIONAL,
//             matchValue      [3] AssertionValue,
//             dnAttributes    [4] BOOLEAN DEFAULT FALSE }
type MatchingRuleAssertion struct {
	matchingRule *MatchingRuleId
	type_        *AttributeDescription
	matchValue   AssertionValue
	dnAttributes BOOLEAN
}

//
//        SearchResultEntry ::= [APPLICATION 4] SEQUENCE {
//             objectName      LDAPDN,
//             attributes      PartialAttributeList }
type SearchResultEntry struct {
	objectName LDAPDN
	attributes PartialAttributeList
}

//
//        PartialAttributeList ::= SEQUENCE OF
//                             partialAttribute PartialAttribute
type PartialAttributeList []PartialAttribute

//
//        SearchResultReference ::= [APPLICATION 19] SEQUENCE
//                                  SIZE (1..MAX) OF uri URI
type SearchResultReference []URI

//
//        SearchResultDone ::= [APPLICATION 5] LDAPResult
type SearchResultDone LDAPResult

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
type ModifyRequest struct {
	object  LDAPDN
	changes []struct {
		operation    int
		modification PartialAttribute
	}
}

//
//        ModifyResponse ::= [APPLICATION 7] LDAPResult
type ModifyResponse LDAPResult

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
type AddRequest struct {
	entry      LDAPDN
	attributes AttributeList
}

//
//        AttributeList ::= SEQUENCE OF attribute Attribute
type AttributeList []Attribute

//
//        AddResponse ::= [APPLICATION 9] LDAPResult
type AddResponse LDAPResult

//
//        DelRequest ::= [APPLICATION 10] LDAPDN
type DelRequest LDAPDN

//
//        DelResponse ::= [APPLICATION 11] LDAPResult
type DelResponse LDAPResult

//
//        ModifyDNRequest ::= [APPLICATION 12] SEQUENCE {
//             entry           LDAPDN,
//             newrdn          RelativeLDAPDN,
//             deleteoldrdn    BOOLEAN,
//             newSuperior     [0] LDAPDN OPTIONAL }
//
type ModifyDNRequest struct {
	entry        LDAPDN
	newrdn       RelativeLDAPDN
	deleteoldrdn BOOLEAN
	newSuperior  *LDAPDN
}

//        ModifyDNResponse ::= [APPLICATION 13] LDAPResult
type ModifyDNResponse LDAPResult

//
//        CompareRequest ::= [APPLICATION 14] SEQUENCE {
//             entry           LDAPDN,
//             ava             AttributeValueAssertion }
type CompareRequest struct {
	entry LDAPDN
	ava   AttributeValueAssertion
}

//
//        CompareResponse ::= [APPLICATION 15] LDAPResult
type CompareResponse LDAPResult

//
//        AbandonRequest ::= [APPLICATION 16] MessageID
type AbandonRequest MessageID

//
//        ExtendedRequest ::= [APPLICATION 23] SEQUENCE {
//             requestName      [0] LDAPOID,
//             requestValue     [1] OCTET STRING OPTIONAL }
type ExtendedRequest struct {
	requestName  LDAPOID
	requestValue *OCTETSTRING
}

//
//        ExtendedResponse ::= [APPLICATION 24] SEQUENCE {
//             COMPONENTS OF LDAPResult,
//             responseName     [10] LDAPOID OPTIONAL,
//             responseValue    [11] OCTET STRING OPTIONAL }
type ExtendedResponse struct {
	LDAPResult
	responseName  *LDAPOID
	responseValue *OCTETSTRING
}

//
//        IntermediateResponse ::= [APPLICATION 25] SEQUENCE {
//             responseName     [0] LDAPOID OPTIONAL,
//             responseValue    [1] OCTET STRING OPTIONAL }
type IntermediateResponse struct {
	responseName  *LDAPOID
	responseValue *OCTETSTRING
}

//
//        END
//
