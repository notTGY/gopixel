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
  uuid := "<YOUR MINECRAFT UUID>"
  playerResponse, _ := gopixel.QueryPlayerApi(uuid, token)
  fmt.Printf("%v", playerResponse)
}
```
