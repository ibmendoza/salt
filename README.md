A simple NaCl library to generate an encrypted token with claims like that of JWT. NaCl already provides 
[encrypt-then-MAC](https://www.daemonology.net/blog/2009-06-24-encrypt-then-mac.html) capability so you don't need HMAC to sign the token

Claims: set exp using the following convenience functions:

- ExpiresInSeconds
- ExpiresInMinutes
- ExpiresInHours
- ExpiresInDays
- ExpiresInMonths

Generate NaCl key: call GenerateKey()

Generate a Token
----------------

```go
func Sign(claims map[string]interface{}, naclKey string) (string, error)
```

Rules: 

- claims is a map or the equivalent of object in JavaScript
- naclKey is used to encrypt the claims
- call GenerateKey() to generate naclKey

Verify a Token
--------------

```go
func Verify(token, naclKey string) (map[string]interface{}, error)
```

Returns the corresponding claims as map[string]interface{} if token is valid


Example
-------

```go
package main

import (
	"fmt"
	"github.com/ibmendoza/salt"
	"time"
)

func main() {

	claims := make(map[string]interface{})
	claims["sub"] = "1234567890"
	claims["name"] = "John Doe"
	claims["admin"] = true
	claims["exp"] = salt.ExpiresInSeconds(1)
	//claims["exp"] = jwt.ExpiresInMinutes(1)
	//claims["exp"] = jwt.ExpiresInHours(1)

	key, _ := salt.GenerateKey()
	fmt.Println(key)

	token, _ := salt.Sign(claims, key)

	fmt.Println(token)

	fmt.Println(salt.Verify(token, key))

	timer1 := time.NewTimer(time.Second * 2) 
	<-timer1.C

	fmt.Println(salt.Verify(token, key)) //expired after 2secs
}
```

Author
-----

Isagani Mendoza (http://itjumpstart.wordpress.com)

License
-------

MIT

Want JWT instead?
--------------------------------

See http://github.com/ibmendoza/jwt
