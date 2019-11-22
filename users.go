package adsync

import (
	"database/sql"
	"fmt"
	"html"
	"math"
	"strconv"
	"time"

	ldap "github.com/go-ldap/ldap/v3"
	"github.com/nlopes/slack"
	_ "gopkg.in/goracle.v2"
)

func getUsers(api *slack.Client) *ldap.SearchResult {
	l, err := ldap.Dial("tcp", LDAPCredentials.server)
	errorCheck(api, err, SlackConfig.channel, "ADIR_USERS_E")
	defer l.Close()
	err = l.Bind(LDAPCredentials.username, LDAPCredentials.password)

	errorCheck(api, err, SlackConfig.channel, "ADIR_USERS_E")

	searchRequest := ldap.NewSearchRequest(
		"OU=ams user,DC=office,DC=amsiag,DC=com", // The base dn to search
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		math.MaxInt32,
		0,
		false,
		"(&(objectClass=person))", // The filter to apply
		usersAttributesAD,         // A list attributes to retrieve
		nil,
	)

	searchResult, err := l.SearchWithPaging(searchRequest, math.MaxInt32)
	errorCheck(api, err, SlackConfig.channel, "ADIR_USERS_E")
	l.Close()
	return searchResult
}

func insertUsersOracle(api *slack.Client, searchResult *ldap.SearchResult) {

	db, err := sql.Open("goracle", AmsOracleCredentials)
	errorCheck(api, err, SlackConfig.channel, "ADIR_USERS_E")
	defer db.Close()
	var ACCOUNTEXPIRES []uint64
	var C,
		CO,
		COMPANY,
		COUNTRYCODE,
		DEPARTMENT,
		DESCRIPTION,
		DISPLAYNAME,
		DISTINGUISHEDNAME,
		EMPLOYEEID,
		EXTENSIONATTRIBUTE1,
		EXTENSIONATTRIBUTE2,
		EXTENSIONATTRIBUTE3,
		EXTENSIONATTRIBUTE4,
		EXTENSIONATTRIBUTE5,
		EXTENSIONATTRIBUTE8,
		EXTENSIONATTRIBUTE10,
		EXTENSIONATTRIBUTE12,
		EXTENSIONATTRIBUTE15,
		FACSIMILETELEPHONENUMBER,
		GIVENNAME,
		IPPHONE,
		L,
		MAIL,
		MEMBEROF,
		MOBILE,
		MSEXCHHIDEFROMADDRESSLISTS,
		OBJECTCATEGORY,
		OBJECTGUID,
		OTHERTELEPHONE,
		PHYSICALDELIVERYOFFICENAME,
		POSTALCODE,
		PROXYADDRESSES,
		ROOMNUMBER,
		SAMACCOUNTNAME,
		SN,
		STREETADDRESS,
		TELEPHONENUMBER,
		TITLE,
		WHENCHANGED,
		WHENCREATED,
		SAMACCOUNTTYPE,
		MSEXCHEXTENSIONATTRIBUTE20,
		INFO,
		LASTLOGON,
		EXTENSIONATTRIBUTE9,
		MSEXCHREQUIREAUTHTOSENDTO,
		EXPORTDATETIME []string
	for _, entry := range searchResult.Entries {
		accExpeires, _ := strconv.ParseUint(entry.GetAttributeValue("accountExpires"), 10, 32)
		ACCOUNTEXPIRES = append(ACCOUNTEXPIRES, accExpeires)
		C = append(C, html.EscapeString(entry.GetAttributeValue("c")))
		CO = append(CO, html.EscapeString(entry.GetAttributeValue("co")))
		COMPANY = append(COMPANY, html.EscapeString(entry.GetAttributeValue("company")))
		COUNTRYCODE = append(COUNTRYCODE, html.EscapeString(entry.GetAttributeValue("countryCode")))
		DEPARTMENT = append(DEPARTMENT, html.EscapeString(entry.GetAttributeValue("department")))
		DESCRIPTION = append(DESCRIPTION, html.EscapeString(entry.GetAttributeValue("description")))
		DISPLAYNAME = append(DISPLAYNAME, html.EscapeString(entry.GetAttributeValue("displayName")))
		DISTINGUISHEDNAME = append(DISTINGUISHEDNAME, html.EscapeString(entry.GetAttributeValue("distinguishedName")))
		EMPLOYEEID = append(EMPLOYEEID, html.EscapeString(entry.GetAttributeValue("employeeID")))
		EXTENSIONATTRIBUTE1 = append(EXTENSIONATTRIBUTE1, html.EscapeString(entry.GetAttributeValue("extensionAttribute1")))
		EXTENSIONATTRIBUTE2 = append(EXTENSIONATTRIBUTE2, html.EscapeString(entry.GetAttributeValue("extensionAttribute2")))
		EXTENSIONATTRIBUTE3 = append(EXTENSIONATTRIBUTE3, html.EscapeString(entry.GetAttributeValue("extensionAttribute3")))
		EXTENSIONATTRIBUTE4 = append(EXTENSIONATTRIBUTE4, html.EscapeString(entry.GetAttributeValue("extensionAttribute4")))
		EXTENSIONATTRIBUTE5 = append(EXTENSIONATTRIBUTE5, html.EscapeString(entry.GetAttributeValue("extensionAttribute5")))
		EXTENSIONATTRIBUTE8 = append(EXTENSIONATTRIBUTE8, html.EscapeString(entry.GetAttributeValue("extensionAttribute8")))
		EXTENSIONATTRIBUTE10 = append(EXTENSIONATTRIBUTE10, html.EscapeString(entry.GetAttributeValue("extensionAttribute10")))
		EXTENSIONATTRIBUTE12 = append(EXTENSIONATTRIBUTE12, html.EscapeString(entry.GetAttributeValue("extensionAttribute12")))
		EXTENSIONATTRIBUTE15 = append(EXTENSIONATTRIBUTE15, html.EscapeString(entry.GetAttributeValue("extensionAttribute15")))
		FACSIMILETELEPHONENUMBER = append(FACSIMILETELEPHONENUMBER, html.EscapeString(entry.GetAttributeValue("facsimileTelephoneNumber")))
		GIVENNAME = append(GIVENNAME, html.EscapeString(entry.GetAttributeValue("givenName")))
		IPPHONE = append(IPPHONE, html.EscapeString(entry.GetAttributeValue("ipPhone")))
		L = append(L, html.EscapeString(entry.GetAttributeValue("l")))
		MAIL = append(MAIL, html.EscapeString(entry.GetAttributeValue("mail")))
		MEMBEROF = append(MEMBEROF, html.EscapeString(entry.GetAttributeValue("memberOf")))
		MOBILE = append(MOBILE, html.EscapeString(entry.GetAttributeValue("mobile")))
		MSEXCHHIDEFROMADDRESSLISTS = append(MSEXCHHIDEFROMADDRESSLISTS, html.EscapeString(entry.GetAttributeValue("msExchHideFromAddressLists")))
		OBJECTCATEGORY = append(OBJECTCATEGORY, html.EscapeString(entry.GetAttributeValue("objectCategory")))
		OBJECTGUID = append(OBJECTGUID, html.EscapeString(fmt.Sprintf("%x", entry.GetRawAttributeValue("objectGUID"))))
		OTHERTELEPHONE = append(OTHERTELEPHONE, html.EscapeString(entry.GetAttributeValue("otherTelephone")))
		PHYSICALDELIVERYOFFICENAME = append(PHYSICALDELIVERYOFFICENAME, html.EscapeString(entry.GetAttributeValue("physicalDeliveryOfficeName")))
		POSTALCODE = append(POSTALCODE, html.EscapeString(entry.GetAttributeValue("postalCode")))
		PROXYADDRESSES = append(PROXYADDRESSES, html.EscapeString(entry.GetAttributeValue("proxyAddresses")))
		ROOMNUMBER = append(ROOMNUMBER, html.EscapeString(entry.GetAttributeValue("roomNumber")))
		SAMACCOUNTNAME = append(SAMACCOUNTNAME, html.EscapeString(entry.GetAttributeValue("sAMAccountName")))
		SN = append(SN, html.EscapeString(entry.GetAttributeValue("sn")))
		STREETADDRESS = append(STREETADDRESS, html.EscapeString(entry.GetAttributeValue("streetAddress")))
		TELEPHONENUMBER = append(TELEPHONENUMBER, html.EscapeString(entry.GetAttributeValue("telephoneNumber")))
		TITLE = append(TITLE, html.EscapeString(entry.GetAttributeValue("title")))
		WHENCHANGED = append(WHENCHANGED, html.EscapeString(entry.GetAttributeValue("whenChanged")))
		WHENCREATED = append(WHENCREATED, html.EscapeString(entry.GetAttributeValue("whenCreated")))
		SAMACCOUNTTYPE = append(SAMACCOUNTTYPE, html.EscapeString(entry.GetAttributeValue("sAMAccountType")))
		MSEXCHEXTENSIONATTRIBUTE20 = append(MSEXCHEXTENSIONATTRIBUTE20, html.EscapeString(entry.GetAttributeValue("msExchExtensionAttribute20")))
		INFO = append(INFO, html.EscapeString(entry.GetAttributeValue("info")))
		LASTLOGON = append(LASTLOGON, html.EscapeString(entry.GetAttributeValue("lastLogon")))
		EXTENSIONATTRIBUTE9 = append(EXTENSIONATTRIBUTE9, html.EscapeString(entry.GetAttributeValue("extensionAttribute9")))
		MSEXCHREQUIREAUTHTOSENDTO = append(MSEXCHREQUIREAUTHTOSENDTO, html.EscapeString(entry.GetAttributeValue("msExchRequireAuthToSendTo")))
		EXPORTDATETIME = append(EXPORTDATETIME, time.Now().String())
	}
	_, err = db.Exec(usersTruncate)
	errorCheck(api, err, SlackConfig.channel, "ADIR_USERS_E")

	_, err = db.Exec(usersInsertSQL, ACCOUNTEXPIRES,
		C,
		CO,
		COMPANY,
		COUNTRYCODE,
		DEPARTMENT,
		DESCRIPTION,
		DISPLAYNAME,
		DISTINGUISHEDNAME,
		EMPLOYEEID,
		EXTENSIONATTRIBUTE1,
		EXTENSIONATTRIBUTE2,
		EXTENSIONATTRIBUTE3,
		EXTENSIONATTRIBUTE4,
		EXTENSIONATTRIBUTE5,
		EXTENSIONATTRIBUTE8,
		EXTENSIONATTRIBUTE10,
		EXTENSIONATTRIBUTE12,
		EXTENSIONATTRIBUTE15,
		FACSIMILETELEPHONENUMBER,
		GIVENNAME,
		IPPHONE,
		L,
		MAIL,
		MEMBEROF,
		MOBILE,
		MSEXCHHIDEFROMADDRESSLISTS,
		OBJECTCATEGORY,
		OBJECTGUID,
		OTHERTELEPHONE,
		PHYSICALDELIVERYOFFICENAME,
		POSTALCODE,
		PROXYADDRESSES,
		ROOMNUMBER,
		SAMACCOUNTNAME,
		SN,
		STREETADDRESS,
		TELEPHONENUMBER,
		TITLE,
		WHENCHANGED,
		WHENCREATED,
		SAMACCOUNTTYPE,
		MSEXCHEXTENSIONATTRIBUTE20,
		INFO,
		LASTLOGON,
		EXTENSIONATTRIBUTE9,
		MSEXCHREQUIREAUTHTOSENDTO,
		EXPORTDATETIME)
	errorCheck(api, err, SlackConfig.channel, "ADIR_USERS_E")

}

// Users Syncronizhes the ADIR_USER_E table with Active Directory
// It Connects to Active directory and fetches all Users
// Truncates and Inserts into oracle ADIR_USER_E Table
// Sends Log on case of fail via SLACK API
func Users() {
	api := slack.New(SlackConfig.token)
	searchResult := getUsers(api)
	insertUsersOracle(api, searchResult)

}
