# new-taipei-parkinfo
Find the available parking in New Taipei, Taiwan.


### Why to write this tool even we have web system to query?
Sometimes I want to know whether I can go to somewhere by car tomorrow. But I need to make an decision today. This will effect my time to get up. So I need history parking information to help me to make an decision. 

I just wite the crawler. I will add SQLite function to save history information, and use Vue.js to show these information.

### Add your favorite park to main.go
```
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
```
