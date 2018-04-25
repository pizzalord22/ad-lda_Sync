package LDAPMethods

import (
	"fmt"
)

func (adstruct Adstruct) GetUser(ldapstruct Ldapstructs) int {
	for k, v := range ldapstruct.Ldaps {
		if len(v.Uid) == len(adstruct.SAMAccountName) {
			for key, val := range v.Uid {
				if val == adstruct.SAMAccountName[key] {
					return k
				} else {
					break
				}
			}
		}
	}
	DebugCompare(fmt.Sprintf("User not found %s", adstruct.SAMAccountName), nil, 19)
	return -1
}

func (ldapstruct Ldapstruct) CleanUser() (user bool) {
	for _, v := range Master.Ads {
		if len(ldapstruct.Uid) == len(v.SAMAccountName) {
			for key, val := range ldapstruct.Uid {
				if val == v.SAMAccountName[key] {
					DebugCompare(fmt.Sprintf("user found %s", v.SAMAccountName), nil, 28)
					return true
				} else {
					break
				}
			}
		}
	}
	return false
}

func (adstruct Adstruct) CompareWhenChanged(ldapstruct Ldapstruct) (r bool) {
	r = true
	if len(adstruct.WhenChanged) == len(ldapstruct.FacsimileTelephoneNumber) {
		for k, val := range adstruct.WhenChanged {
			if val != ldapstruct.FacsimileTelephoneNumber[k] {
				DebugCompare(fmt.Sprintf("When changed does not match changing %s to %s for %s", ldapstruct.FacsimileTelephoneNumber, adstruct.WhenChanged, adstruct.DisplayName), nil, 44)
				r = false
				return
			}
		}
	} else {
		DebugCompare(fmt.Sprintf("When changed does not match changing %s to %s for %s", ldapstruct.FacsimileTelephoneNumber, adstruct.WhenChanged, adstruct.DisplayName), nil, 50)
		r = false
		return
	}
	return
}

func (adstruct Adstruct) CompareMail(ldapstruct Ldapstruct) (r bool) {
	r = true
	if len(adstruct.Mail) == len(ldapstruct.Mail) {
		for k, val := range adstruct.Mail {
			if val != ldapstruct.Mail[k] {
				DebugCompare(fmt.Sprintf("Mail does not match changing %s to %s for %s", ldapstruct.Mail, adstruct.Mail, adstruct.DisplayName), nil, 62)
				r = false
				return
			}
		}
	} else {
		DebugCompare(fmt.Sprintf("Mail does not match changing %s to %s for %s", ldapstruct.Mail, adstruct.Mail, adstruct.DisplayName), nil, 68)
		r = false
		return
	}
	return
}

func (adstruct Adstruct) CompareDisplayName(ldapstruct Ldapstruct) (r bool) {
	r = true
	if len(adstruct.DisplayName) == len(ldapstruct.DisplayName) {
		for k, val := range adstruct.DisplayName {
			if val != ldapstruct.DisplayName[k] {
				DebugCompare(fmt.Sprintf("Display does not match changing %s to %s", ldapstruct.DisplayName, adstruct.DisplayName), nil, 80)
				r = false
				return
			}
		}
	} else {
		DebugCompare(fmt.Sprintf("Display does not match changing %s to %s", ldapstruct.DisplayName, adstruct.DisplayName), nil, 86)
		r = false
		return
	}
	return
}

func (adstruct Adstruct) CompareDepartment(ldapstruct Ldapstruct) (r bool) {
	r = true
	if len(adstruct.Department) == len(ldapstruct.PhysicalDeliveryOfficeName) {
		for k, val := range adstruct.Department {
			if val != ldapstruct.PhysicalDeliveryOfficeName[k] {
				DebugCompare(fmt.Sprintf("Department does not match chaning %s to %s for user %s", ldapstruct.PhysicalDeliveryOfficeName, adstruct.Department, adstruct.DisplayName), nil, 98)
				r = false
				return
			}
		}
	} else {
		DebugCompare(fmt.Sprintf("Department does not match chaning %s to %s for user %s", ldapstruct.PhysicalDeliveryOfficeName, adstruct.Department, adstruct.DisplayName), nil, 104)
		r = false
		return
	}
	return
}

func (adstruct Adstruct) CompareCompany(ldapstruct Ldapstruct) (r bool) {
	r = true
	if len(adstruct.Company) == len(ldapstruct.O) {
		for k, val := range adstruct.Company {
			if val != ldapstruct.O[k] {
				DebugCompare(fmt.Sprintf("Company does not match changing %s to %s for user %s", ldapstruct.O, adstruct.Company, adstruct.DisplayName), nil, 116)
				r = false
				return
			}
		}
	} else {
		DebugCompare(fmt.Sprintf("Company does not match changing %s to %s for user %s", ldapstruct.O, adstruct.Company, adstruct.DisplayName), nil, 122)
		r = false
		return
	}
	return
}

func (adstruct Adstruct) CompareUPN(ldapstruct Ldapstruct) (r bool) {
	r = true
	if len(adstruct.UserPrincipalName) == len(ldapstruct.UserPassword) {
		for k, val := range adstruct.UserPrincipalName {
			if k == 0 {
				if "{SASL}"+val != ldapstruct.UserPassword[k] {
					DebugCompare(fmt.Sprintf("The password for user %s did not match the AD", adstruct.DisplayName), nil, 135)
					r = false
					return
				}
			} else {
				if val != ldapstruct.UserPassword[k] {
					DebugCompare(fmt.Sprintf("The password for user %s did not match the AD", adstruct.DisplayName), nil, 141)
					r = false
					return
				}
			}
		}
	} else {
		r = false
		DebugCompare(fmt.Sprintf("The password for user %s did not match the AD", adstruct.DisplayName), nil, 149)
		return
	}
	return
}

func (adstruct Adstruct) CompareSn(ldapstruct Ldapstruct) (r bool) {
	r = true
	if len(adstruct.Sn) == len(ldapstruct.Sn) {
		for k, val := range adstruct.Sn {
			if val != ldapstruct.Sn[k] {
				DebugCompare(fmt.Sprintf("The surname did not match changing %s to %s for user %s", ldapstruct.Sn, adstruct.Sn, adstruct.DisplayName), nil, 160)
				r = false
				return
			}
		}
	} else {
		DebugCompare(fmt.Sprintf("The surname did not match changing %s to %s for user %s", ldapstruct.Sn, adstruct.Sn, adstruct.DisplayName), nil, 166)
		r = false
		return
	}
	return
}

func (adstruct Adstruct) CompareGivenName(ldapstruct Ldapstruct) (r bool) {
	r = true
	if len(adstruct.GivenName) == len(ldapstruct.GivenName) {
		for k, val := range adstruct.GivenName {
			if val != ldapstruct.GivenName[k] {
				DebugCompare(fmt.Sprintf("Given name does not match changing %s to %s for user %s", ldapstruct.GivenName, adstruct.GivenName, adstruct.DisplayName), nil, 178)
				r = false
				return
			}
		}
	} else {
		DebugCompare(fmt.Sprintf("Given name does not match changing %s to %s for user %s", ldapstruct.GivenName, adstruct.GivenName, adstruct.DisplayName), nil, 184)
		r = false
		return
	}
	return
}

func (adstruct Adstruct) CompareObjectClass(ldapstruct Ldapstruct) (r bool) {
	r = true
	if len(adstruct.ObjectClass) <= len(ldapstruct.ObjectClass) {
		for k, val := range adstruct.ObjectClass {
			if val != ldapstruct.ObjectClass[k] {
				DebugCompare(fmt.Sprintf("ObjectClass does not match changing %s to %s for user %s", ldapstruct.ObjectClass, adstruct.ObjectClass, adstruct.DisplayName), nil, 196)
				r = false
				return
			}
		}
	} else {
		DebugCompare(fmt.Sprintf("ObjectClass does not match changing %s to %s for user %s", ldapstruct.ObjectClass, adstruct.ObjectClass, adstruct.DisplayName), nil, 202)
		r = false
		return
	}
	return
}

/*func (adstruct Adstruct) CompareMemberOf(ldapstruct Ldapstruct) (r bool) {
	r = false
	if len(ldapstruct.MemberOf) > 0 {
		r = true
	} else {
		fmt.Println(len(ldapstruct.MemberOf))
	}
	return
}*/

func (master Ldapstructs) Compare() {
	DebugCompare(fmt.Sprintf("Comparing all AD and LDAP results"), nil, 220)
	var skip bool
	for k, v := range master.Ads {
		skip = false
		index := v.GetUser(Slave)
		if index == -1 {
			DebugCompare(fmt.Sprintf("Adding a new account to Ldap with display name: %s", v.DisplayName), nil, 226)
			v.addAccount()
			skip = true
		}
		if skip == false {
			ldapIndex := Slave.Ldaps[index]
			SyncChecks.WhenChanged = v.CompareWhenChanged(ldapIndex)
			SyncChecks.Mail = v.CompareMail(ldapIndex)
			SyncChecks.DisplayName = v.CompareDisplayName(ldapIndex)
			SyncChecks.Department = v.CompareDepartment(ldapIndex)
			SyncChecks.Company = v.CompareCompany(ldapIndex)
			SyncChecks.UserPrincipalName = v.CompareUPN(ldapIndex)
			SyncChecks.Sn = v.CompareSn(ldapIndex)
			SyncChecks.GivenName = v.CompareGivenName(ldapIndex)
			SyncChecks.ObjectClass = v.CompareObjectClass(ldapIndex)
			SyncChecks.CheckSync(k, index)
		}
	}
	DebugCompare(fmt.Sprintf("Done comparing all AD LDAP results"), nil, 244)
}

func (sync SyncCheck) CheckSync(mindex int, sindex int) {
	if !sync.WhenChanged {
		Slave.Ldaps[sindex].FacsimileTelephoneNumber = Master.Ads[mindex].WhenChanged
		Slave.Ldaps[sindex].modifyWhenChanged()
	}
	if !sync.Mail {
		Slave.Ldaps[sindex].Mail = Master.Ads[mindex].Mail
		Slave.Ldaps[sindex].modifyMail()
	}
	if !sync.DisplayName {
		Slave.Ldaps[sindex].DisplayName = Master.Ads[mindex].DisplayName
		Slave.Ldaps[sindex].modifyDisplayName()
	}
	if !sync.Department {
		Slave.Ldaps[sindex].PhysicalDeliveryOfficeName = Master.Ads[mindex].Department
		Slave.Ldaps[sindex].modifyDepartment()
	}
	if !sync.Company {
		Slave.Ldaps[sindex].O = Master.Ads[mindex].Company
		Slave.Ldaps[sindex].modifyCompany()
	}
	if !sync.UserPrincipalName {
		Slave.Ldaps[sindex].UserPassword = []string{"{SASL}" + Master.Ads[mindex].UserPrincipalName[0]}
		Slave.Ldaps[sindex].modifyUPN()
	}
	if !sync.Sn {
		Slave.Ldaps[sindex].Sn = Master.Ads[mindex].Sn
		Slave.Ldaps[sindex].modifySn()
	}
	if !sync.GivenName {
		Slave.Ldaps[sindex].GivenName = Master.Ads[mindex].GivenName
		Slave.Ldaps[sindex].modifyGivenName()
	}
	if !sync.ObjectClass {
		Slave.Ldaps[sindex].ObjectClass = Master.Ads[mindex].ObjectClass
		Slave.Ldaps[sindex].modifyObjectClass()
	}
}
