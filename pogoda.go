package main

import (
	"./mypkg"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type ViewData struct {
	Title   string
	Message string
}

func main() {
	jsonFile, err := os.Open("pogoda.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened pogoda.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var weather mypkg.StructWeather

	json.Unmarshal(byteValue, &weather)
	fmt.Printf("Now: %v\n", weather.Now)
	fmt.Printf("Fact.Condition: %v\n", mypkg.GetCondition(weather.Fact.Condition))

	data := ViewData{
		Title:   "Some Title",
		Message: "Some Content",
	}

	tmpl, _ := template.ParseFiles("tmpl/index.html")
	//tmpl.Execute(os.Stdout, data)

	var html bytes.Buffer
	if err := tmpl.Execute(&html, data); err != nil {
		fmt.Println(err)
	}

	result := html.String()

	fmt.Println("\nHTML output result: \n\n", result)
	fmt.Println(".................................................")
	fmt.Println(MinifyHTML([]byte(result)))
}

func MinifyHTML(html []byte) string {
	minifiedHTML := ""
	scanner := bufio.NewScanner(bytes.NewReader(CleanHtml(html)))
	for scanner.Scan() {
		lineTrimmed := strings.TrimSpace(scanner.Text())
		minifiedHTML += lineTrimmed
		if len(lineTrimmed) > 0 {
			minifiedHTML += " "
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	minifiedHTML = CleanJsonBrackets(minifiedHTML)

	return minifiedHTML
}

func CleanHtml(content []byte) []byte {
	cleanComments := regexp.MustCompile(`<!--[^>]*-->`)
	cleanSpaces := regexp.MustCompile(`>\s{1,}<`)
	noComment := cleanComments.ReplaceAll(content, []byte(""))
	return cleanSpaces.ReplaceAll(noComment, []byte("><"))
}

func CleanJsonBrackets(html string) string {
	htmlClear := strings.ReplaceAll(html, " } <", "}<")
	htmlClear = strings.ReplaceAll(htmlClear, " { ", "{")
	return htmlClear
}
