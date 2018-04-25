package LDAPMethods

import (
	"gopkg.in/ldap.v2"
	"fmt"
)

func (ldapstruct Ldapstruct) modifyWhenChanged() {
	modifyReq := &ldap.ModifyRequest{DN: ldapstruct.entryDN}
	modifyReq.Replace("facsimileTelephoneNumber", ldapstruct.FacsimileTelephoneNumber)
	err := Slave.Connection.Modify(modifyReq)
	if err != nil {
		DebugModify(fmt.Sprintf("There was an error while changing whenChanged"), err, 13)
		fmt.Println("There was an error while modifying LDAP, check the logs for more info")
		return
	}
	DebugModify(fmt.Sprintf("whenChangec was modified for user %s\r\n", ldapstruct.DisplayName), nil, 17)
}

func (ldapstruct Ldapstruct) modifyMail() {
	modifyReq := &ldap.ModifyRequest{DN: ldapstruct.entryDN}
	modifyReq.Replace("mail", ldapstruct.Mail)
	err := Slave.Connection.Modify(modifyReq)
	if err != nil {
		DebugModify(fmt.Sprintf("There was an error while changing mail"), err, 25)
		fmt.Println("There was an error while modifying LDAP, check the logs for more info")
		return
	}
	DebugModify(fmt.Sprintf("Mail was modified for user %s", ldapstruct.DisplayName), nil, 29)
}

func (ldapstruct Ldapstruct) modifyDisplayName() {
	modifyReq := &ldap.ModifyRequest{DN: ldapstruct.entryDN}
	modifyReq.Replace("displayname", ldapstruct.DisplayName)
	err := Slave.Connection.Modify(modifyReq)
	if err != nil {
		DebugModify(fmt.Sprintf("There was an error while changing displayName"), err, 37)
		fmt.Println("There was an error while modifying LDAP, check the logs for more info")
		return
	}
	DebugModify(fmt.Sprintf("Display name was modified for user %s", ldapstruct.DisplayName), nil, 41)
}

func (ldapstruct Ldapstruct) modifyDepartment() {
	modifyReq := &ldap.ModifyRequest{DN: ldapstruct.entryDN}
	modifyReq.Replace("physicalDeliveryOfficeName", ldapstruct.PhysicalDeliveryOfficeName)
	err := Slave.Connection.Modify(modifyReq)
	if err != nil {
		DebugModify(fmt.Sprintf("There was an error while changing Department"), err, 49)
		fmt.Println("There was an error while modifying LDAP, check the logs for more info")
		return
	}
	DebugModify(fmt.Sprintf("Department was modified for user %s", ldapstruct.DisplayName), nil, 53)
}

func (ldapstruct Ldapstruct) modifyCompany() {
	modifyReq := &ldap.ModifyRequest{DN: ldapstruct.entryDN}
	modifyReq.Replace("o", ldapstruct.O)
	err := Slave.Connection.Modify(modifyReq)
	if err != nil {
		DebugModify(fmt.Sprintf("There was an error while changing Company"), err, 61)
		fmt.Println("There was an error while modifying LDAP, check the logs for more info")
		return
	}
	DebugModify(fmt.Sprintf("Company was modified for user %s", ldapstruct.DisplayName), nil, 65)
}

func (ldapstruct Ldapstruct) modifyUPN() {
	modifyReq := &ldap.ModifyRequest{DN: ldapstruct.entryDN}
	modifyReq.Replace("userPassword", ldapstruct.UserPassword)
	err := Slave.Connection.Modify(modifyReq)
	if err != nil {
		DebugModify(fmt.Sprintf("There was an error while changing UserPrincipalname"), err, 73)
		fmt.Println("There was an error while modifying LDAP, check the logs for more info")
		return
	}
	DebugModify(fmt.Sprintf("userPrincipalName was modified for user %s", ldapstruct.DisplayName), nil, 77)
}

func (ldapstruct Ldapstruct) modifySn() {
	modifyReq := &ldap.ModifyRequest{DN: ldapstruct.entryDN}
	modifyReq.Replace("sn", ldapstruct.Sn)
	err := Slave.Connection.Modify(modifyReq)
	if err != nil {
		DebugModify(fmt.Sprintf("There was an error while changing Sn"), err, 85)
		fmt.Println("There was an error while modifying LDAP, check the logs for more info")
		return
	}
	DebugModify(fmt.Sprintf("Sn was modified for user %s", ldapstruct.DisplayName), nil, 89)
}

func (ldapstruct Ldapstruct) modifyGivenName() {
	modifyReq := &ldap.ModifyRequest{DN: ldapstruct.entryDN}
	modifyReq.Replace("givenName", ldapstruct.GivenName)
	err := Slave.Connection.Modify(modifyReq)
	if err != nil {
		DebugModify(fmt.Sprintf("There was an error while changing Givenname"), err, 97)
		fmt.Println("There was an error while modifying LDAP, check the logs for more info")
		return
	}
	DebugModify(fmt.Sprintf("GivenName was modified for user %s", ldapstruct.DisplayName), nil, 101)
}

func (ldapstruct Ldapstruct) modifyObjectClass() {
	modifyReq := &ldap.ModifyRequest{DN: ldapstruct.entryDN}
	modifyReq.Replace("objectClass", ldapstruct.ObjectClass)
	err := Slave.Connection.Modify(modifyReq)
	if err != nil {
		DebugModify(fmt.Sprintf("There was an error while changing objectClass"), err, 109)
		fmt.Println("There was an error while modifying LDAP, check the logs for more info")
		return
	}
	DebugModify(fmt.Sprintf("ObjectClass was modified for user %s", ldapstruct.DisplayName), nil, 113)
}

func (ldapstruct Ldapstruct) modifyMemberOf() {
	modifyReq := &ldap.ModifyRequest{DN: ldapstruct.entryDN}
	modifyReq.Replace("memberOf", ldapstruct.MemberOf)
	err := Slave.Connection.Modify(modifyReq)
	if err != nil {
		DebugModify(fmt.Sprintf("There was an error while changing memberOf"), err, 121)
		fmt.Println("There was an error while modifying LDAP, check the logs for more info")
		return
	}
	DebugModify(fmt.Sprintf("memborOf was modified for user %s", ldapstruct.DisplayName), nil, 125)
}
