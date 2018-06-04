package main

import (
    "fmt"
)

func main() {
    s, err := getParkInfo()
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(findParkInfo(s, "新北市藝文中心平面停車場"))
    fmt.Println(findParkInfo(s, "板橋高中地下停車場"))
    fmt.Println(findParkInfo(s, "市民廣場地下停車場"))
    fmt.Println(findParkInfo(s, "寶橋立體停車場"))
}