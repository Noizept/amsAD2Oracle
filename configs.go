package adsync

// SlackToken Slack API Token
var SlackToken string = "myToken"

// AmsOracleCredentials Credentials to connect into database format USER/PASSWORD@SERVER
var AmsOracleCredentials string = "Connection-String"

// LDAPConnection Connection struct for LDAP
type LDAPConnection struct {
	username string
	password string
	server   string
}

// SetUsername set LDAPCredentials username
func (conn *LDAPConnection) SetUsername(name string) {
	conn.username = name
}

// SetPassword set LDAPCredentials password
func (conn *LDAPConnection) SetPassword(password string) {
	conn.password = password
}

// SetServer set LDAPCredentials server
func (conn *LDAPConnection) SetServer(server string) {
	conn.server = server
}

// LDAPCredentials Default
var LDAPCredentials *LDAPConnection = &LDAPConnection{username: "ldapUsername", password: "ldapPassword", server: "ldapServer"}
