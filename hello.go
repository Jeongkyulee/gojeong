package gojeong

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"runtime"
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

func StringToSHA(str string) string {
	hashCode := sha256.New()
	hashCode.Write([]byte(str))
	mdCode := hashCode.Sum(nil)
	stringSHA := hex.EncodeToString(mdCode)
	return stringSHA
}

func CMDcommand(command string) string {
	exeCommand := []string{"cmd", "/c"}
	cmd := exec.Command(exeCommand[0], exeCommand[1], command)
	cmdPath, err := os.Getwd()
	ErrCheck(err)
	cmd.Dir = cmdPath
	Output, err := cmd.Output()
	ErrCheck(err)
	encordingString, _, err := transform.String(korean.EUCKR.NewDecoder(), string(Output))
	ErrCheck(err)
	return encordingString
}

func BashCommand(command string) string {
	exeCommand := []string{"bash", "-c"}
	cmd := exec.Command(exeCommand[0], exeCommand[1], command)
	cmdPath, err := os.Getwd()
	ErrCheck(err)
	cmd.Dir = cmdPath
	Output, err := cmd.Output()
	ErrCheck(err)
	encordingString, _, err := transform.String(korean.EUCKR.NewDecoder(), string(Output))
	ErrCheck(err)
	return encordingString
}

func CheckOS() string {
	os := runtime.GOOS
	return os
}
