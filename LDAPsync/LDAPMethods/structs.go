package LDAPMethods

import "gopkg.in/ldap.v2"

var Master Ldapstructs
var Slave Ldapstructs
//var GroupSync Ldapstructs
var SyncChecks SyncCheck
var CNGroups Groups
var Search bool
// debug variables
var DebugMainState bool
var DebugCompareState bool
var DebugModifyState bool
var DebugDeleteState bool
var DebugInitialState bool
var DebugAddState bool

type Ldapstructs struct {
	Connection    ldap.Conn
	Ldaps         []Ldapstruct
	Ads           []Adstruct
	Dc            string
	Filter        string
	Attributes    []string
	SearchResults ldap.SearchResult
	EntryDN       [][]string
}

type Adstruct struct {
	Mail              []string
	DisplayName       []string
	LockoutTime       []string
	SAMAccountName    []string
	WhenChanged       []string
	Department        []string
	Company           []string
	UserPrincipalName []string
	Sn                []string
	GivenName         []string
	ObjectClass       []string
	MemberOf          []string
}

type Ldapstruct struct {
	Mail                       []string
	DisplayName                []string
	LockoutTime                []string
	Uid                        []string
	FacsimileTelephoneNumber   []string
	PhysicalDeliveryOfficeName []string
	O                          []string
	UserPassword               []string
	Sn                         []string
	GivenName                  []string
	ObjectClass                []string
	MemberOf                   []string
	entryDN                    string
}

type SyncCheck struct {
	WhenChanged       bool
	Mail              bool
	DisplayName       bool
	Department        bool
	Company           bool
	UserPrincipalName bool
	Sn                bool
	GivenName         bool
	ObjectClass       bool
	MemberOf          bool
}

type UpdateLdap struct {
	add    []*ldap.AddRequest
	modify []*ldap.ModifyRequest
	delete []*ldap.DelRequest
}

type Groups struct {
	Group []string
}
