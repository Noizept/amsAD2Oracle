## ams Custom LDAP to Oracle Implementation

# Step 1

Install the Package

`go get github.com/Noizept/amsAD2Oracle`

# Step 2

Overwrite the constants :

- SlackConfig
- AmsOracleCredentials
- LDAPCredentials

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

	amsAD2Oracle.SlackConfig.SetToken("TOKEN")
	amsAD2Oracle.SlackConfig.SetChannel("pmd")

	amsAD2Oracle.AmsOracleCredentials = "MY_USERSCHEMA/MYPASSWORD@SERVER.AMS.COM"

	amsAD2Oracle.LDAPCredentials.SetServer("AMS-LDAP.com:389")
	amsAD2Oracle.LDAPCredentials.SetUsername("ldap_bind")
	amsAD2Oracle.LDAPCredentials.SetPassword("LDAP_PASSWORD")

	var wg sync.WaitGroup
	fmt.Println(time.Now().String())

	wg.Add(3)
	go users(&wg)
	go groups(&wg)
	go contacts(&wg)
	wg.Wait()

	fmt.Println(time.Now().String())

}

func users(wg *sync.WaitGroup) {
	amsAD2Oracle.Users()
	wg.Done()

}

func contacts(wg *sync.WaitGroup) {
	amsAD2Oracle.Contacts()
	wg.Done()

}
func groups(wg *sync.WaitGroup) {
	amsAD2Oracle.Groups()
	wg.Done()

}
```
