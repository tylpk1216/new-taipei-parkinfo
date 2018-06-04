package main

import (
    "fmt"
    "strings"
)

type parkItem struct {
    name string
    leftCount string
    allCount string
    updateDate string
    updateTime string
}

func getPadSpaceForName(length int) string {
    max := 16

    // chinese utf8 is 3
    word := length / 3

    var b []byte
    // When showing chinese word, it only needs two bytes.
    limit := max * 2 - word * 2
    for i := 0; i < limit; i++ {
        b = append(b, ' ')
    }

    return string(b)
}

func formatParkInfo(item *parkItem) string {
    yi := strings.Index(item.updateDate, "年")
    mi := strings.Index(item.updateDate, "月")
    di := strings.Index(item.updateDate, "日")
    hi := strings.Index(item.updateTime, "時")
    mini := strings.Index(item.updateTime, "分")
    si := strings.Index(item.updateTime, "秒")

    var year, month, day, hour, min, sec string

    if yi != -1 && mi != -1 && di != -1 &&
       hi != -1 && mini != -1 && si != -1 {

        year = item.updateDate[:yi]
        month = item.updateDate[yi+len("年"):mi]
        day = item.updateDate[mi+len("月"):di]
        hour = item.updateTime[:hi]
        min = item.updateTime[hi+len("時"):mini]
        sec = item.updateTime[mini+len("分"):si]
    }

    space := getPadSpaceForName(len(item.name))

    return fmt.Sprintf(
        "%s%s %-4s %-4s %s/%s/%s %s:%s:%s",
        item.name, space, item.leftCount, item.allCount,
        year, month, day, hour, min, sec,
    )
}

//-----------------------------------------------------------------------------
/*
 --- 2018/5/31 format
 --- 有顏色的才會有 span tag
<tr>
    <td class='Point' width='18%'>新店安坑停車場(岳洋)</td>
    <td class='Text_Center' width='22%'>新北市新店區安康路2段341巷旁</td>
    <td class='Text_Center' width='13%'>
       <span style='color:red;'>3002</span>
    </td>
    <td class='Text_Center' width='11%'>42</td>
    <td class='Text_Center' width='11%'>107年05月30日</td>
    <td class='Text_Center' width='19%'>18時54分01秒</td>
</tr>
*/
func findParkInfo(s string, name string) string {
    var item parkItem
    item.name = name

    c := ""

    i := strings.Index(s, name)
    if i == -1 {
        return c
    }

    // "name"xxxxxxxx
    ns := s[i:]

    // move to leftCount. it maybe has "span tag" or not.
    count := 0
    for {
        i = strings.Index(ns, ">")
        if i == -1 {
            return c
        }

        ns = ns[i+1:]
        count++

        if count == 4 {
            i = 0
            break
        }
    }

    // if it has "span tag", move to the first byte of leftCount.
    j := strings.Index(ns, "<span")
    if j == 0 {
        i = strings.Index(ns, ">")
        if i == -1 {
            return c
        }

        ns = ns[i+1:]
        i = 0
    }

    // find the boundary of LeftCount
    i1 := strings.Index(ns, "<")
    if i1 == -1 {
        return c
    }

    item.leftCount = ns[i:i1]

    ns = ns[i1+1:]

    var msg [3]string
    count = 0
    for {
        // "td tag" begin
        i := strings.Index(ns, "'>")
        if i == -1{
            return c
        }

        ns = ns[i+2:]

        // "td tag" end(</td>)
        i1 := strings.Index(ns, "<")
        if i1 == -1 {
            return c
        }

        msg[count] = ns[:i1]
        count++

        ns = ns[i1+1:]

        if count == 3 {
            break
        }
    }

    item.allCount = msg[0]
    item.updateDate = msg[1]
    item.updateTime = msg[2]

    return formatParkInfo(&item)
}
