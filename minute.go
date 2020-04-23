package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	separator := string(filepath.Separator)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	checkErr(err)
	unixTimeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	fileName := "sample-" + unixTimeStamp + ".txt"
	filePath := dir + separator + fileName
	f, er := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	checkErr(er)
	defer f.Close()
	fmt.Println("filePath: ", filePath)
}

func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}
