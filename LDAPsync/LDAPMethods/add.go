package LDAPMethods

import (
	"gopkg.in/ldap.v2"
	"fmt"
	"github.com/spf13/viper"
)

func (adstruct Adstruct) addAccount() {
	tdn := "cn="
	for _, v := range adstruct.DisplayName {
		tdn += v
	}
	tdn += ", " + viper.GetString("slave-dc")
	addreq := ldap.NewAddRequest(tdn)
	addreq.Attribute("uid", adstruct.SAMAccountName)
	addreq.Attribute("objectClass", adstruct.ObjectClass)
	addreq.Attribute("mail", adstruct.Mail)
	addreq.Attribute("memberOf", []string{"cn=Testgroup,ou=Groups,ou=Users,dc=Eresato,dc=com"})
	addreq.Attribute("displayName", adstruct.DisplayName)
	addreq.Attribute("o", adstruct.Company)
	addreq.Attribute("physicalDeliveryOfficeName", adstruct.Department)
	addreq.Attribute("givenName", adstruct.GivenName)
	addreq.Attribute("sn", adstruct.Sn)
	addreq.Attribute("facsimileTelephoneNumber", adstruct.WhenChanged)
	addreq.Attribute("userPassword", []string{"{SASL}" + adstruct.UserPrincipalName[0]})
	err := Slave.Connection.Add(addreq)
	if err != nil {
		DebugAdd(fmt.Sprintf("There was an error while creating a LDAP account"), err, 35)
		fmt.Println("There was an errow hile adding an account, check the logs for more info")
		return
	}
	// connect to default group and then add the new account to it
	groupreq := ldap.NewModifyRequest(viper.GetString("default-group"))
	groupreq.Add("member", []string{tdn})
	err = Slave.Connection.Modify(groupreq)
	if err != nil {
		DebugAdd(fmt.Sprintf("There was an error while adding %s to a group", tdn), err, 48)
		fmt.Println("There was an error while adding a user to a group, check the logs for more info")
		return
	}
	Search = true
	DebugAdd(fmt.Sprintf("Added an account: %s", tdn), nil, 43)
}
