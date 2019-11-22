package adsync

import (
	"database/sql"
	"fmt"
	"html"
	"math"
	"time"

	ldap "github.com/go-ldap/ldap/v3"
	"github.com/nlopes/slack"
	_ "gopkg.in/goracle.v2"
)

func getContacts(api *slack.Client) *ldap.SearchResult {
	l, err := ldap.Dial("tcp", LDAPCredentials.server)
	errorCheck(api, err, SlackConfig.channel, "ADIR_CONTACTS_E")
	defer l.Close()
	err = l.Bind(LDAPCredentials.username, LDAPCredentials.password)
	errorCheck(api, err, SlackConfig.channel, "ADIR_CONTACTS_E")

	searchRequest := ldap.NewSearchRequest(
		"OU=SPSAddressList,DC=office,DC=amsiag,DC=com", // The base dn to search
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		math.MaxInt32,
		0,
		false,
		"(&(objectClass=contact))", // The filter to apply
		contactAttributesAD,        // A list attributes to retrieve
		nil,
	)

	searchResult, err := l.SearchWithPaging(searchRequest, math.MaxInt32)
	errorCheck(api, err, SlackConfig.channel, "ADIR_CONTACTS_E")
	l.Close()
	return searchResult
}

func insertContactsOracle(api *slack.Client, searchResult *ldap.SearchResult) {

	db, err := sql.Open("goracle", AmsOracleCredentials)
	errorCheck(api, err, SlackConfig.channel, "ADIR_CONTACTS_E")
	defer db.Close()

	var DISPLAYNAME,
		DISTINGUISHEDNAME,
		GIVENNAME,
		CONTACTNAME,
		OBJECTGUID,
		SN,
		TELEPHONENUMBER,
		WHENCREATED,
		WHENCHANGED,
		CN,
		COMPANY,
		MAIL,
		MAILNICKNAME,
		MEMBEROF,
		MSEXCHHIDEFROMADDRESSLISTS,
		OBJECTCATEGORY,
		PROXYADDRESSES,
		MSEXCHEXTENSIONATTRIBUTE20,
		MSEXCHREQUIREAUTHTOSENDTO,
		EXPORTDATETIME []string
	for _, entry := range searchResult.Entries {
		DISPLAYNAME = append(DISPLAYNAME, html.EscapeString(entry.GetAttributeValue("displayName")))
		DISTINGUISHEDNAME = append(DISTINGUISHEDNAME, html.EscapeString(entry.GetAttributeValue("distinguishedName")))
		GIVENNAME = append(GIVENNAME, html.EscapeString(entry.GetAttributeValue("givenName")))
		CONTACTNAME = append(CONTACTNAME, html.EscapeString(entry.GetAttributeValue("name")))
		OBJECTGUID = append(OBJECTGUID, html.EscapeString(fmt.Sprintf("%x", entry.GetRawAttributeValue("objectGUID"))))
		SN = append(SN, html.EscapeString(entry.GetAttributeValue("sn")))
		TELEPHONENUMBER = append(TELEPHONENUMBER, html.EscapeString(entry.GetAttributeValue("telephoneNumber")))
		WHENCHANGED = append(WHENCHANGED, html.EscapeString(entry.GetAttributeValue("whenChanged")))
		WHENCREATED = append(WHENCREATED, html.EscapeString(entry.GetAttributeValue("whenCreated")))
		CN = append(CN, html.EscapeString(entry.GetAttributeValue("cn")))
		COMPANY = append(COMPANY, html.EscapeString(entry.GetAttributeValue("company")))
		MAIL = append(MAIL, html.EscapeString(entry.GetAttributeValue("mail")))
		MAILNICKNAME = append(MAILNICKNAME, html.EscapeString(entry.GetAttributeValue("mailNickname")))
		MEMBEROF = append(MEMBEROF, html.EscapeString(entry.GetAttributeValue("memberOf")))
		MSEXCHHIDEFROMADDRESSLISTS = append(MSEXCHHIDEFROMADDRESSLISTS, html.EscapeString(entry.GetAttributeValue("msExchHideFromAddressLists")))
		OBJECTCATEGORY = append(OBJECTCATEGORY, html.EscapeString(entry.GetAttributeValue("objectCategory")))
		PROXYADDRESSES = append(PROXYADDRESSES, html.EscapeString(entry.GetAttributeValue("proxyAddresses")))
		MSEXCHEXTENSIONATTRIBUTE20 = append(MSEXCHEXTENSIONATTRIBUTE20, html.EscapeString(entry.GetAttributeValue("msExchExtensionAttribute20")))
		MSEXCHREQUIREAUTHTOSENDTO = append(MSEXCHREQUIREAUTHTOSENDTO, html.EscapeString(entry.GetAttributeValue("msExchRequireAuthToSendTo")))
		EXPORTDATETIME = append(EXPORTDATETIME, time.Now().String())
	}
	_, err = db.Exec(contactTruncate)
	errorCheck(api, err, SlackConfig.channel, "ADIR_CONTACTS_E")

	_, err = db.Exec(contactInsertSQL,
		DISPLAYNAME,
		DISPLAYNAME,
		DISTINGUISHEDNAME,
		GIVENNAME,
		CONTACTNAME,
		OBJECTGUID,
		SN,
		TELEPHONENUMBER,
		WHENCREATED,
		WHENCHANGED,
		CN,
		COMPANY,
		MAIL,
		MAILNICKNAME,
		MEMBEROF,
		MSEXCHHIDEFROMADDRESSLISTS,
		OBJECTCATEGORY,
		PROXYADDRESSES,
		MSEXCHEXTENSIONATTRIBUTE20,
		MSEXCHREQUIREAUTHTOSENDTO,
		EXPORTDATETIME)
	errorCheck(api, err, SlackConfig.channel, "ADIR_CONTACTS_E")

}

// Contacts Syncronizhes the ADIR_CONTACTS_E table with Active Directory
// It Connects to Active directory and fetches all Users
// Truncates and Inserts into oracle ADIR_GROUPS_E Table
// Sends Log on case of fail via SLACK API
func Contacts() {
	api := slack.New(SlackConfig.token)
	searchResult := getContacts(api)
	insertContactsOracle(api, searchResult)

}
