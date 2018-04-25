package LDAPMethods

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/ldap.v2"
	"strings"
)

func (master *Ldapstructs) BindConnection(name string) {
	ip := viper.GetString(name + "-ip")
	port := viper.GetInt(name + "-port")
	principal := viper.GetString(name + "-principal")
	password := viper.GetString(name + "-password")
	conn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		DebugInitial(fmt.Sprintf("There was an error while Dialing"), err, 17)
		panic("There was an error while connnecting check the logs for info")
	}
	err = conn.Bind(principal, password)
	if err != nil {
		DebugInitial(fmt.Sprintf("There was an error while binding the connection"), err, 22)
		panic("There was an error while connnecting check the logs for info")
	}
	master.Connection = *conn
	DebugInitial(fmt.Sprintf("A connection was made using ip: %s, port: %d, principal: %s, password: %s", ip, port, principal, password), nil, 26)
}

func (master *Ldapstructs) SearchLDAP() {
	searchRequest := ldap.NewSearchRequest(
		master.Dc,
		ldap.ScopeSingleLevel,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		master.Filter,
		master.Attributes,
		nil,
	)
	sr, err := master.Connection.Search(searchRequest)
	if err != nil {
		DebugInitial(fmt.Sprintf("There was an error while searching trought the system"), err, 43)
		fmt.Println("SearchLDAP gave an error check the logs for more info")
	}
	master.SearchResults = *sr
	DebugInitial(fmt.Sprintf("Searched LDAP, DC: %s, filter: %s, attributes: %s", master.Dc, master.Filter, master.Attributes), nil, 47)
}

func (master *Ldapstructs) SortAd() {
	myArray := [][]string{}
	tmpstruct := Adstruct{}
	var add bool
	for _, entry := range master.SearchResults.Entries {
		add = true
		myArray = [][]string{}
		myArray = append(myArray, entry.GetAttributeValues("sn"))
		myArray = append(myArray, []string{"organizationalPerson", "top", "person", "inetOrgPerson"})
		myArray = append(myArray, []string{"cn=Testgroup,ou=Groups,ou=Users,dc=Eresato,dc=com"})
		myArray = append(myArray, entry.GetAttributeValues("mail"))
		myArray = append(myArray, entry.GetAttributeValues("givenName"))
		myArray = append(myArray, entry.GetAttributeValues("displayName"))
		myArray = append(myArray, entry.GetAttributeValues("whenChanged"))
		myArray = append(myArray, entry.GetAttributeValues("userPrincipalName"))
		myArray = append(myArray, entry.GetAttributeValues("sAMAccountName"))
		myArray = append(myArray, entry.GetAttributeValues("company"))
		myArray = append(myArray, entry.GetAttributeValues("department"))
		master.EntryDN = append(master.EntryDN, entry.GetAttributeValues("cn"), entry.GetAttributeValues("ou"), entry.GetAttributeValues("dc"))
		for _, v := range myArray {
			if len(v) == 0 {
				add = false
				DebugInitial(fmt.Sprintf("There was an empty field for user: %s, %s %s", myArray[6], myArray[5], myArray[0]), nil, 72)
				fmt.Sprintf("There was an empty field while sorting the AD, check the logs for more info\r\n")
				break
			}
		}
		if add == true {
			tmpstruct.Sn = myArray[0]
			tmpstruct.ObjectClass = myArray[1]
			tmpstruct.MemberOf = myArray[2]
			tmpstruct.Mail = myArray[3]
			tmpstruct.GivenName = myArray[4]
			tmpstruct.DisplayName = myArray[5]
			tmpstruct.WhenChanged = myArray[6]
			tmpstruct.UserPrincipalName = myArray[7]
			tmpstruct.SAMAccountName = myArray[8]
			tmpstruct.Company = myArray[9]
			tmpstruct.Department = myArray[10]
			master.Ads = append(master.Ads, tmpstruct)
		}
	}
	DebugInitial(fmt.Sprintf("Done sorting AD"), nil, 92)
}

func (master *Ldapstructs) ClearStruct() {
	master.Ads = []Adstruct{}
	master.Ldaps = []Ldapstruct{}
	DebugInitial(fmt.Sprintf("Cleared the structs"), nil, 98)
}

func (master *Ldapstructs) SortLdap() {
	tmpstruct := Ldapstruct{}
	for _, entry := range master.SearchResults.Entries {
		tmpstruct.Sn = entry.GetAttributeValues("sn")
		tmpstruct.ObjectClass = entry.GetAttributeValues("objectClass")
		tmpstruct.MemberOf = entry.GetAttributeValues("memberOf")
		tmpstruct.Mail = entry.GetAttributeValues("mail")
		tmpstruct.LockoutTime = entry.GetAttributeValues("lockoutTime")
		tmpstruct.GivenName = entry.GetAttributeValues("givenName")
		tmpstruct.DisplayName = entry.GetAttributeValues("displayName")
		tmpstruct.O = entry.GetAttributeValues("o")
		tmpstruct.UserPassword = entry.GetAttributeValues("userPassword")
		tmpstruct.Uid = entry.GetAttributeValues("uid")
		tmpstruct.PhysicalDeliveryOfficeName = entry.GetAttributeValues("physicalDeliveryOfficeName")
		tmpstruct.FacsimileTelephoneNumber = entry.GetAttributeValues("facsimileTelephoneNumber")
		tmpstruct.entryDN = entry.DN
		master.Ldaps = append(master.Ldaps, tmpstruct)
	}
	DebugInitial(fmt.Sprintf("Sorted LDAP"), nil, 119)
}

func (master *Ldapstructs) SetAttributes(name string) {
	attributes := strings.Split(viper.GetString(name+"-Attributes"), ", ")
	for _, v := range attributes {
		master.Attributes = append(master.Attributes, v)
	}
	DebugInitial(fmt.Sprintf("Set the attributes for %s", name), nil, 127)
}

func (master *Ldapstructs) SetDC(name string) {
	master.Dc = viper.GetString(name + "-Dc")
	DebugInitial(fmt.Sprintf("set DC for %s", name), nil, 132)
}

func (master *Ldapstructs) SetFilter(name string) {
	master.Filter = viper.GetString(name + "-Filter")
	DebugInitial(fmt.Sprintf("set Filter for %s", name), nil, 137)
}

func (master *Ldapstructs) Initialize(name string) {
	DebugInitial(fmt.Sprintf("Initializing %s", name), nil, 141)
	master.BindConnection(name)
	master.SetDC(name)
	master.SetFilter(name)
	master.SetAttributes(name)
	DebugInitial(fmt.Sprintf("Finished initializing %s", name), nil, 146)
}

func (groups *Groups) Initialize() {
	grouping := viper.GetStringSlice("groups")
	for _, v := range grouping {
		groups.Group = append(groups.Group, v)
	}
	DebugInitial(fmt.Sprintf("set groups for current session"), nil, 154)
}
