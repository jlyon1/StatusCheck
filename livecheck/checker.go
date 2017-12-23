package livecheck

import (
  "time"
  "fmt"
  "StatusCheck/database"
)

func registerCheck(d time.Duration, Database database.DB, addr string) {
  for x := range time.Tick(d) {
    _= x;
    fmt.Println(x);
  }
}
