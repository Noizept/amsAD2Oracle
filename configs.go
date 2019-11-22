package adsync

// SlackToken Slack API Token
var SlackToken string = "myToken"

// AmsOracleCredentials Credentials to connect into database format USER/PASSWORD@SERVER
var AmsOracleCredentials string = "Connection-String"
// SlackConfiguration struct for Slack notifications
type SlackConfiguration struct {
	token string
	channel string
}


// SetToken set SlackConfig token
func (slack *SlackConfiguration) SetToken(token string) {
	slack.token = token
}

// SetChannel set SlackConfig channel
func (slack *SlackConfiguration) SetChannel(channel string) {
	slack.channel = channel
}


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
// SlackConfig Default
var SlackConfig *SlackConfiguration = &SlackConfiguration{token: "myToken", channel: "myChanell"}
