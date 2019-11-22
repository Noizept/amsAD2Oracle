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

func getGroups(api *slack.Client) *ldap.SearchResult {
	l, err := ldap.Dial("tcp", LDAPCredentials.server)
	errorCheck(api, err, "pmd", "ADIR_GROUPS_E")
	defer l.Close()
	err = l.Bind(LDAPCredentials.username, LDAPCredentials.password)
	errorCheck(api, err, "pmd", "ADIR_GROUPS_E")

	searchRequest := ldap.NewSearchRequest(
		"DC=office,DC=amsiag,DC=com", // The base dn to search
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		math.MaxInt32,
		0,
		false,
		"(&(objectCategory=group)(objectClass=group))", // The filter to apply
		groupAttributesAD, // A list attributes to retrieve
		nil,
	)

	searchResult, err := l.SearchWithPaging(searchRequest, math.MaxInt32)
	errorCheck(api, err, "pmd", "ADIR_GROUPS_E")
	l.Close()
	return searchResult
}

func insertGroupsOracle(api *slack.Client, searchResult *ldap.SearchResult) {

	db, err := sql.Open("goracle", AmsOracleCredentials)
	errorCheck(api, err, "pmd", "ADIR_GROUPS_E")
	defer db.Close()

	var DESCRIPTION,
		DISPLAYNAME,
		DISTINGUISHEDNAME,
		GROUPTYPE,
		MAIL,
		MANAGEDBY,
		MEMBEROF,
		MSEXCHHIDEFROMADDRESSLISTS,
		OBJECTCATEGORY,
		OBJECTGUID,
		PROXYADDRESSES,
		SAMACCOUNTNAME,
		SAMACCOUNTTYPE,
		WHENCHANGED,
		WHENCREATED,
		INFO,
		MAILNICKNAME,
		MSEXCHALOBJECTVERSION,
		MSEXCHARBITRATIONMAILBOX,
		MSEXCHRECIEPIENTDISPLAYTYPE,
		SHOWINADDRESSBOOK,
		MEMBERS,
		MSEXCHCOMANAGEDBYLINK,
		MSEXCHREQUIREAUTHTOSENDTO,
		GIDNUMBER,
		MSEXCHEXTENSIONATTRIBUTE20,
		DLMEMSUBMITPERMS,
		EXPORTDATETIME []string
	for _, entry := range searchResult.Entries {

		DESCRIPTION = append(DESCRIPTION, html.EscapeString(entry.GetAttributeValue("description")))
		DISPLAYNAME = append(DISPLAYNAME, html.EscapeString(entry.GetAttributeValue("displayName")))
		DISTINGUISHEDNAME = append(DISTINGUISHEDNAME, html.EscapeString(entry.GetAttributeValue("distinguishedName")))
		GROUPTYPE = append(GROUPTYPE, html.EscapeString(entry.GetAttributeValue("groupType")))
		MAIL = append(MAIL, html.EscapeString(entry.GetAttributeValue("mail")))
		MANAGEDBY = append(MANAGEDBY, html.EscapeString(entry.GetAttributeValue("managedBy")))
		MEMBEROF = append(MEMBEROF, html.EscapeString(entry.GetAttributeValue("memberOf")))
		MSEXCHHIDEFROMADDRESSLISTS = append(MSEXCHHIDEFROMADDRESSLISTS, html.EscapeString(entry.GetAttributeValue("msExchHideFromAddressLists")))
		OBJECTCATEGORY = append(OBJECTCATEGORY, html.EscapeString(entry.GetAttributeValue("objectCategory")))
		OBJECTGUID = append(OBJECTGUID, html.EscapeString(fmt.Sprintf("%x", entry.GetRawAttributeValue("objectGUID"))))
		PROXYADDRESSES = append(PROXYADDRESSES, html.EscapeString(entry.GetAttributeValue("proxyAddresses")))
		SAMACCOUNTNAME = append(SAMACCOUNTNAME, html.EscapeString(entry.GetAttributeValue("sAMAccountName")))
		SAMACCOUNTTYPE = append(SAMACCOUNTTYPE, html.EscapeString(entry.GetAttributeValue("sAMAccountType")))
		WHENCHANGED = append(WHENCHANGED, html.EscapeString(entry.GetAttributeValue("whenChanged")))
		WHENCREATED = append(WHENCREATED, html.EscapeString(entry.GetAttributeValue("whenCreated")))
		INFO = append(INFO, html.EscapeString(entry.GetAttributeValue("info")))
		MAILNICKNAME = append(MAILNICKNAME, html.EscapeString(entry.GetAttributeValue("mailNickname")))
		MSEXCHALOBJECTVERSION = append(MSEXCHALOBJECTVERSION, html.EscapeString(entry.GetAttributeValue("msExchALObjectVersion")))
		MSEXCHARBITRATIONMAILBOX = append(MSEXCHARBITRATIONMAILBOX, html.EscapeString(entry.GetAttributeValue("msExchArbitrationMailbox")))
		MSEXCHRECIEPIENTDISPLAYTYPE = append(MSEXCHRECIEPIENTDISPLAYTYPE, html.EscapeString(entry.GetAttributeValue("msExchRecipientDisplayType")))
		SHOWINADDRESSBOOK = append(SHOWINADDRESSBOOK, html.EscapeString(entry.GetAttributeValue("showInAddressBook")))
		MEMBERS = append(MEMBERS, html.EscapeString(entry.GetAttributeValue("member")))
		MSEXCHCOMANAGEDBYLINK = append(MSEXCHCOMANAGEDBYLINK, html.EscapeString(entry.GetAttributeValue("msExchCoManagedByLink")))
		MSEXCHREQUIREAUTHTOSENDTO = append(MSEXCHREQUIREAUTHTOSENDTO, html.EscapeString(entry.GetAttributeValue("msExchRequireAuthToSendTo")))
		GIDNUMBER = append(GIDNUMBER, html.EscapeString(entry.GetAttributeValue("gidNumber")))
		MSEXCHEXTENSIONATTRIBUTE20 = append(MSEXCHEXTENSIONATTRIBUTE20, html.EscapeString(entry.GetAttributeValue("msExchExtensionAttribute20")))
		DLMEMSUBMITPERMS = append(DLMEMSUBMITPERMS, html.EscapeString(entry.GetAttributeValue("dLMemSubmitPerms")))
		EXPORTDATETIME = append(EXPORTDATETIME, time.Now().String())
	}
	_, err = db.Exec(groupTruncate)
	errorCheck(api, err, "pmd", "ADIR_GROUPS_E")

	_, err = db.Exec(groupInsertSQL, DESCRIPTION,
		DISPLAYNAME,
		DISTINGUISHEDNAME,
		GROUPTYPE,
		MAIL,
		MANAGEDBY,
		MEMBEROF,
		MSEXCHHIDEFROMADDRESSLISTS,
		OBJECTCATEGORY,
		OBJECTGUID,
		PROXYADDRESSES,
		SAMACCOUNTNAME,
		SAMACCOUNTTYPE,
		WHENCHANGED,
		WHENCREATED,
		INFO,
		MAILNICKNAME,
		MSEXCHALOBJECTVERSION,
		MSEXCHARBITRATIONMAILBOX,
		MSEXCHRECIEPIENTDISPLAYTYPE,
		SHOWINADDRESSBOOK,
		MEMBERS,
		MSEXCHCOMANAGEDBYLINK,
		MSEXCHREQUIREAUTHTOSENDTO,
		GIDNUMBER,
		MSEXCHEXTENSIONATTRIBUTE20,
		DLMEMSUBMITPERMS,
		EXPORTDATETIME)
	errorCheck(api, err, "pmd", "ADIR_GROUPS_E")

}

// Groups Syncronizhes the ADIR_GROUPS_E table with Active Directory
// It Connects to Active directory and fetches all Users
// Truncates and Inserts into oracle ADIR_GROUPS_E Table
// Sends Log on case of fail via SLACK API
func Groups() {
	api := slack.New(SlackToken)
	searchResult := getGroups(api)
	insertGroupsOracle(api, searchResult)

}
