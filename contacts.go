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

//Contact Table ETL
type Contact struct {
	ams AMSETL
}

func (contact *Contact) Sync(){
	api := slack.New(contact.ams.SlackConfig.Token)

	searchResult,err := contact.ams.GetEntries()
	errorCheck(api,err, contact.ams.SlackConfig.Channel,"ADIR_GROUPS_E")

	err = contact.insertOracle(searchResult)
	errorCheck(api,err, contact.ams.SlackConfig.Channel,"ADIR_CONTACTS_E")
}


//SetConfigs initiates the configurations for the Contact Struct
func (contact *Contact) SetConfigs(conf AMSETL){
	contact.ams = conf
}

func (contact *Contact)  insertOracle (searchResult *ldap.SearchResult) (err error) {
	db, err := sql.Open("goracle", contact.ams.AmsOracleCredentials)
	if err!= nil {
		return err
	}
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
	if err!= nil {
		return err
	}
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
	if err!= nil {
		return err
	}
	return nil
}