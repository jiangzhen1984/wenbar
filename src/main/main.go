

package main

import (
    "gotom"
    "main/service"
)



func main() {
    ws.InitDB(ws.DBConfiguration{DBUrl:"localhost"})
    gotom.InitServer(conf)
}
