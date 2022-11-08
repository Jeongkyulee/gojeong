package gojeong

import (
	"bufio"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"os"
	"strings"
)

func Excelsearch(path string, filename string, dirname string) []string {
	fileopen, err := os.Open(strings.Replace(path+"\\"+dirname+"\\"+filename, "\\", "/", 100))
	ErrCheck(err)
	defer func() {
		fileopen.Close()
	}()
	readline := bufio.NewScanner(fileopen)
	var bufResult string
	var bufNow string
	bufCount := 0
	countStart := 0
	countOnOff := 0
	// 결과가 들어가는 배열 resultEnd
	var resultEnd []string
	for readline.Scan() {
		tempstring, _, _ := transform.String(korean.EUCKR.NewDecoder(), string(readline.Text()))
		//      for i := 0; i < len(temprune); i++ {
		//         tempstring = tempstring + string(temprune[i])
		//      }
		if strings.Contains(tempstring, "[START]") {
			countStart += 1
			countOnOff = 1
		}
		if countOnOff == 1 {
			if strings.Contains(tempstring, "결과") {
				bufResult = tempstring
			}
			if bufCount == 1 {
				if !strings.Contains(tempstring, "[END]") {
					bufNow = bufNow + tempstring + "\r\n"
				}
			}
			if strings.Contains(tempstring, "현황") {
				bufCount = 1
			}
		}
		if strings.Contains(tempstring, "[END]") {
			countOnOff = 0
			bufCount = 0
			resultEnd = append(resultEnd, bufResult, strings.TrimRight(bufNow, "\r\n"))
			// string variable 초기화
			bufResult = ""
			bufNow = ""
		}
	}
	return resultEnd
}
