package main

import (
    "fmt"
    "io/ioutil"
    "net/http"

    "golang.org/x/text/transform"
    "golang.org/x/text/encoding/traditionalchinese"

    // MIT
    "github.com/parnurzeal/gorequest"
)

func getParkInfo() (string, error) {
    //return getParkInfoByFile()
    return getParkInfoByURL()
}

func getParkInfoByFile() (string, error) {
    s, err := ioutil.ReadFile("data/ParkInfo.txt")
    if err != nil {
        return "", err
    }

    // covert big5 to utf8
    ns, _, err := transform.String(traditionalchinese.Big5.NewDecoder(), string(s))
    if err != nil {
        return "", err
    }

    return ns, nil
}

func getParkInfoByURL() (string, error) {
    url := "http://www.parkinginfo.ntpc.gov.tw/parkingrealInfo/"

    request := gorequest.New()
    resp, body, errs := request.Get(url).End()

    if len(errs) > 0 {
        fmt.Println(errs)
        return "", errs[0]
    }

    if resp.StatusCode != http.StatusOK {
        err := fmt.Errorf("network error(%d)", resp.StatusCode)

        return "", err
    }

    // covert big5 to utf8
    s, _, err := transform.String(traditionalchinese.Big5.NewDecoder(), body)
    if err != nil {
        return "", err
    }

    return s, nil
}