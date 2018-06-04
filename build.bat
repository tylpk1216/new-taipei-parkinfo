@ECHO OFF

SET PATH=%PATH%;C:\Go\bin
SET GOPATH=%CD%\..\..\

go build -o ParkInfo.exe main.go fetchdata.go passdata.go

