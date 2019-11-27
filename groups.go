package adsync

import (
	"database/sql"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/nlopes/slack"
	_ "gopkg.in/goracle.v2"
	"html"
	"time"
)

//Group Table ETL
type Group struct {
	ams AMSETL
}

func (group *Group) Sync(){
	api := slack.New(group.ams.SlackConfig.Token)

	searchResult,err := group.ams.GetEntries()
	errorCheck(api,err, group.ams.SlackConfig.Channel,"ADIR_GROUPS_E")

	err = group.insertOracle(searchResult)
	errorCheck(api,err, group.ams.SlackConfig.Channel,"ADIR_GROUPS_E")


}


//SetConfigs initiates the configurations for the Group Struct
func (group *Group) SetConfigs(conf AMSETL){
	group.ams = conf
}

func (group *Group)  insertOracle (searchResult *ldap.SearchResult) (err error) {

	db, err := sql.Open("goracle", group.ams.AmsOracleCredentials)
	if err!= nil {
		return err
	}
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
	if err!= nil {
		return err
	}
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
	if err!= nil {
		return err
	}
	return nil
}