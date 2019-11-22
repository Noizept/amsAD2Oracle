package adsync

var groupAttributesAD = []string{
	"description",
	"displayName",
	"distinguishedName",
	"groupType",
	"mail",
	"managedBy",
	"memberOf",
	"msExchHideFromAddressLists",
	"objectCategory",
	"objectGUID",
	"proxyAddresses",
	"sAMAccountName",
	"sAMAccountType",
	"whenChanged",
	"whenCreated",
	"info",
	"mailNickname",
	"msExchALObjectVersion",
	"msExchArbitrationMailbox",
	"msExchRecipientDisplayType",
	"showInAddressBook",
	"member",
	"msExchCoManagedByLink",
	"msExchRequireAuthToSendTo",
	"gidNumber",
	"msExchExtensionAttribute20",
	"dLMemSubmitPerms",
}

var groupInsertSQL = `INSERT INTO ADIR_OWNER.ADIR_GROUPS_E (
	DESCRIPTION,
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
	EXPORTDATETIME
) VALUES (
	:DESCRIPTION,
	:DISPLAYNAME,
	:DISTINGUISHEDNAME,
	:GROUPTYPE,
	:MAIL,
	:MANAGEDBY,
	:MEMBEROF,
	:MSEXCHHIDEFROMADDRESSLISTS,
	:OBJECTCATEGORY,
	:OBJECTGUID,
	:PROXYADDRESSES,
	:SAMACCOUNTNAME,
	:SAMACCOUNTTYPE,
	:WHENCHANGED,
	:WHENCREATED,
	:INFO,
	:MAILNICKNAME,
	:MSEXCHALOBJECTVERSION,
	:MSEXCHARBITRATIONMAILBOX,
	:MSEXCHRECIEPIENTDISPLAYTYPE,
	:SHOWINADDRESSBOOK,
	:MEMBERS,
	:MSEXCHCOMANAGEDBYLINK,
	:MSEXCHREQUIREAUTHTOSENDTO,
	:GIDNUMBER,
	:MSEXCHEXTENSIONATTRIBUTE20,
	:DLMEMSUBMITPERMS,
	:EXPORTDATETIME
)`

var groupTruncate ="TRUNCATE TABLE ADIR_GROUPS_E"