# crypt
Pure go implementation of unix crypt(3). This package only implements the traditional 56 bit DES based algorithm. This implementation is ported mostly from the UnixCrypt.java from [Apache Commons Codec](https://commons.apache.org/proper/commons-codec/), with some modifications in order to follow golang conventions. 

This package is go-routine (concurrent) safe.

Example
-------
```go
import (
        "fmt"
        "github.com/digitive/crypt"
)

func main() {
        hash, err := crypt.Crypt("changeme", "Sa")
        if err != nil {
                fmt.Errorf("error:", err)
                return
        }

        fmt.Println("HASH:", hash)
}
```

Tests & Bug reports
-------------------
Only limited tests have been done. Please report any bug or make fixes via Push Requests.

There is a python script at [resources/gen.py](resources/gen.py) which can be used to generate random passwords and hashes by calling python version (glibc version under the hood) of crypt(3), in order to verify the go version's compatibility to glibc version of crypt(3)

You can modify the gen.py and then run `go test` to test it, e.g 

```bash
resources/gen.py > resources/passwords.txt
go test ./...
```

License
-------

Released under the [MIT License](LICENSE)