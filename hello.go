package gojeong

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ErrCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func PathRead(dirname string) (string, string) {
	workPath, err := os.Getwd()
	ErrCheck(err)
	fmt.Println(workPath, reflect.TypeOf(workPath))
	targetPath, err := os.Getwd()
	ErrCheck(err)
	targetPath = targetPath + "\\" + dirname
	fmt.Println(targetPath, reflect.TypeOf(targetPath))
	return workPath, targetPath
}

func ListFile(targetPath string) []string {
	files, err := ioutil.ReadDir(targetPath)
	ErrCheck(err)
	listFile := []string{}
	for _, file := range files {
		time.Sleep(10 * time.Microsecond)
		if !file.IsDir() && strings.Contains(file.Name(), "txt") {
			listFile = append(listFile, file.Name())
		}
	}
	return listFile
}

func CheckRegexp(listPublic []string, str string) []string {
	checkReg, _ := regexp.Compile(str)
	listFile := []string{}
	for _, targetName := range listPublic {
		listFile = append(listFile, checkReg.FindString(targetName))
	}
	return listFile
}

func LimitDate(YYYYMMDD int) {
	fmt.Println("Limit Date Start")
	nowTime, _ := strconv.Atoi(strings.Replace(strings.Split(time.Now().String(), " ")[0], "-", "", -1))
	if YYYYMMDD < nowTime {
		fmt.Println("limit Date Over Exit")
		os.Exit(3)
	}
	fmt.Println("Limit Date Pass")
}

func StrigToHangul(urlEncordString string) string {
	hangulString, err := url.QueryUnescape(urlEncordString)
	ErrCheck(err)
	return hangulString
}
