package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("OMI AUTO TEST TOOL...")
	fmt.Println("請選擇要執行function 1. 操作截圖 2. open csv 3. 啟動固定視窗")

	var a int
	fmt.Scanln(&a)
	switch a {
	case 1:
		robotgo.MoveMouse(762, 1425)
		robotgo.MouseClick()
		robotgo.Sleep(1)
		robotgo.KeyTap("f2", "shift")
	case 2:
		ReadCsv()
	case 3:
		prodcessIds := "notepad.exe"
		fpid, err := robotgo.FindIds(prodcessIds)

		if err == nil {
			fmt.Printf("find  %d\n", fpid[0])
			robotgo.ActivePID(fpid[0])

			mdata := robotgo.GetActive()
			robotgo.SetActive(mdata)

			tl := robotgo.GetTitle(fpid[0])
			fmt.Printf("title is: %s\n", tl)

			// robotgo.MaxWindow(fpid[0])
			// robotgo.MinWindow(fpid[0])
			// robotgo.Sleep(1)

			x, y, w, h := robotgo.GetBounds(fpid[0])
			fmt.Println("GetBounds is: ", x, y, w, h)
			bit1 := robotgo.CaptureScreen(x, y, w, h)
			robotgo.SaveBitmap(bit1, "test2.png")

		}
	default:
	}

}

func ReadCsv() {
	FilePath := "OMI.csv"
	file, err := os.OpenFile(FilePath, os.O_RDONLY, 0777)
	if err != nil {
		log.Fatalln("找不到檔案路徑", FilePath, err)
	}

	csv_data := csv.NewReader(file)
	csv_data.Comma = ','
	m := map[string][]string{}

	for {
		// var action []string
		record, err := csv_data.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		m[record[0]] = record

	}

	fmt.Println("請選擇要執行 OMI FUCNTION")
	for key, value := range m {
		fmt.Printf("FUNCTION %s, Task Step : %s\n", key, value)

	}
	var task string
	fmt.Scanln(&task)
	for _, value := range m[task][1:] {
		exe_robotkey(value)
	}

}

func exe_robotkey(key string) {
	if strings.ContainsAny(key, "()") {
		go_key := key[0:strings.Index(key, "(")]
		value := key[strings.Index(key, "(")+1 : strings.Index(key, ")")]

		fmt.Println(go_key, value)
		robotgo_function(go_key, value)
	} else {
		fmt.Println(key)
	}

}

func robotgo_function(key string, value string) {
	switch key {
	case "key":
		robotgo.TypeStr(value)
	case "sleep":
		v, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		robotgo.Sleep(v)

	default:
		fmt.Println("fuck")
	}

}
