package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/audrenbdb/goforeground"
	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("OMI AUTO TEST TOOL...")

	// datapath := `E:\MegaDownloader_v1.8_bin\MegaDownloader.exe`
	// fmt.Println(datapath)
	// exec.Command(datapath).Start()

	// robotgo.Sleep(4)
	fpid, err := robotgo.FindIds("Chrome.exe")
	if err != nil {
		fmt.Println(err)
	}
	fzname, err := robotgo.FindName(fpid[0])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fzname)
	goforeground.Activate(int(fpid[0]))
	robotgo.Sleep(1)

	// robotgo.ActivePID(fpid[0])

	// robotgo.MaxWindow(fpid[0])
	// mdata := robotgo.GetActive()
	// // set Window Active
	// robotgo.SetActive(mdata)
	// robotgo.TypeString("pu369")
	// robotgo.KeyTap("space")
	// robotgo.TypeString("golang")
	// robotgo.KeyTap("enter")
	// robotgo.MilliSleep(6 * 1000)

	x, y, w, h := robotgo.GetBounds(fpid[0])
	fmt.Println("GetBounds is: ", x, y, w, h)
	bit1 := robotgo.CaptureScreen(x, y, w, h)
	robotgo.SaveBitmap(bit1, "test2.png")

	// prodcessIds := "notepad.exe"

	// fpid, err := robotgo.FindIds(prodcessIds)

	// if err == nil {
	// 	fmt.Printf("find  %d\n", fpid[0])
	// 	robotgo.ActivePID(fpid[0])

	// 	tl := robotgo.GetTitle(fpid[0])
	// 	fmt.Printf("title is: %s\n", tl)

	// 	// robotgo.MaxWindow(fpid[0])
	// 	// robotgo.MinWindow(fpid[0])
	// 	// robotgo.Sleep(1)

	// 	x, y, w, h := robotgo.GetBounds(fpid[0])
	// 	fmt.Println("GetBounds is: ", x, y, w, h)
	// 	bit1 := robotgo.CaptureScreen(x, y, w, h)
	// 	robotgo.SaveBitmap(bit1, "test2.png")

	// }

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
