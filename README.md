# gopixel
Hypixel api golang client

This packages primary use is in
[https://foragingupdate.com](https://foragingupdate.com).

So, if you need more data over skyblock profiles,
file in pull request. No NITs, typos, tests, etc.


## Example Usage

```
package main

import (
  "fmt"
	"github.com/nottgy/gopixel"
)

func main() {
  token := "<YOUR TOKEN>"
  uuid := "3a1bb62230b04dcb9df18a6dba4909c4"
	playerResponse, _ := gopixel.QueryPlayerApi(uuid, token)
  fmt.Printf("%v", playerResponse)
}
```
