## ams Custom LDAP to Oracle Implementation

# Step 1

Install the Package

`go get github.com/Noizept/amsAD2Oracle`

# Step 2

Setup configuration structs 


Example usage with concurrent processing

```
package main

import (
	"fmt"
	amsAD2Oracle "github.com/Noizept/amsAD2Oracle"
	"sync"
	"time"
)

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
		SlackConfig:          amsAD2Oracle.SlackConf{
			Token: "mytoken", 
			Channel: "pmd",
		},
		AmsOracleCredentials: "my_schema/PWD@server.AMS.COM",
		LdapConfig:           amsAD2Oracle.LdapConnection{
			Username: "lda_usr", 
			Password: "ldappw", 
			Server: "ams-ldap.com:389",
		},
		LdapOptions:          amsAD2Oracle.LdapFilter{
			BaseDN: "OU=SPSAddressList,DC=office,DC=amsiag,DC=com",
			ObjectFilter: "(&(objectClass=contact))", 
			AttributesRetrieval: amsAD2Oracle.ContactAttributesAD,
		},
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
