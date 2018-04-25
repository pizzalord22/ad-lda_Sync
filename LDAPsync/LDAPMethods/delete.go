package LDAPMethods

import (
	"gopkg.in/ldap.v2"
	"github.com/spf13/viper"
	"fmt"
)

func (ldapstruct Ldapstruct) deleteAccount() {
	tdn := "cn="
	for _, v := range ldapstruct.DisplayName {
		tdn += v
	}
	tdn += ", " + viper.GetString("slave-dc")
	delReq := ldap.DelRequest{DN: tdn}
	delReq.Controls = []ldap.Control{}
	err := Slave.Connection.Del(&delReq)
	if err != nil {
		DebugDelete("There was an error on deleting from LDAP", err, 19)
		fmt.Println("There was an error while deleting an user check the logs for info")
		return
	}
	DebugDelete(fmt.Sprintf("Deleted account %s", tdn), nil, 23)
}

func (ldapstructs Ldapstructs) CleanSlave() {
	DebugDelete(fmt.Sprintf("STarting cleaning the slave"), nil, 27)
	for _, v := range ldapstructs.Ldaps {
		clean := v.CleanUser()
		if clean == false {
			v.deleteAccount()
		}
	}
	DebugDelete(fmt.Sprintf("Done cleaning the slave"), nil, 34)
}
