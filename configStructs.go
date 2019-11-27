package adsync

import (
	ldap "github.com/go-ldap/ldap/v3"
	"math"
)

// SlackConfiguration struct for Slack notifications
type SlackConf struct {
	Token   string
	Channel string
}

// LdapConnection configuration credentials
type LdapConnection struct {
	Username string
	Password string
	Server   string
}

// LdapFilter query information
type LdapFilter struct {
	BaseDN              string
	ObjectFilter        string
	AttributesRetrieval []string
}

// AMSETL struct for ldap  extracting
type AMSETL struct {
	SlackConfig          SlackConf
	AmsOracleCredentials string
	LdapConfig           LdapConnection
	LdapOptions          LdapFilter
}

// GetEntries Retrieves Entries from Active directory
func (amsetl *AMSETL) GetEntries() (*ldap.SearchResult,error) {
	l, err := ldap.Dial("tcp", amsetl.LdapConfig.Server)
	if err != nil {
		return nil,err
	}
	defer l.Close()
	err = l.Bind(amsetl.LdapConfig.Username, amsetl.LdapConfig.Password)
	if err != nil {
		return nil,err
	}
	searchRequest := ldap.NewSearchRequest(
		amsetl.LdapOptions.BaseDN, // The base dn to search
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		math.MaxInt32,
		0,
		false,
		amsetl.LdapOptions.ObjectFilter,        // The filter to apply
		amsetl.LdapOptions.AttributesRetrieval, // A list attributes to retrieve
		nil,
	)
	searchResult, err := l.SearchWithPaging(searchRequest, math.MaxInt32)
	if err != nil {
		return nil,err
	}
	l.Close()
	return searchResult,nil
}
