## ams Custom LDAP to Oracle Implementation

# Step 1

Install the Package

`go get github.com/Noizept/amsAD2Oracle`

# Step 2

Setup configuration structs 


Example usage with concurrent processing

```
func main() {
	var wg sync.WaitGroup
	fmt.Println(time.Now().String())

	wg.Add(3)
	
    go users(&wg)
	go groups(&wg)
	go contacts(&wg)
	wg.Wait()

	fmt.Println(time.Now().String())
}

func contacts(wg *sync.WaitGroup) {
	conf := amsAD2Oracle.AMSETL{
		SlackConfig:          amsAD2Oracle.SlackConf{Token: "mytoken", Channel: "pmd"},
		AmsOracleCredentials: "MY_SCHEMA/PASSWORD@SERVER.AMS.COM",
		LdapConfig:           amsAD2Oracle.LdapConnection{Username: "USER_LDAP", Password: "PASSWORD", Server: "ams-ldap:389"},
		LdapOptions:          amsAD2Oracle.LdapFilter{BaseDN: "OU=SPSAddressList,DC=office,DC=amsiag,DC=com", ObjectFilter: "(&(objectClass=contact))", AttributesRetrieval: amsAD2Oracle.ContactAttributesAD},
	}

	contacts:= amsAD2Oracle.Contact{}
	contacts.SetConfigs(conf)
	contacts.Sync()
	wg.Done()
}
.
.
.
.


}
```
