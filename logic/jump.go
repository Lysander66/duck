package logic

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	jsonRe    = regexp.MustCompile("```json\n({[\\s\\S]*?})\n```")
	commentRe = regexp.MustCompile(`.*,*\s*(//.*)`)
	imgRe     = regexp.MustCompile(`"(logo|icon)":\s?"([^,]*)",?`)
)

func Jump()  {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		return 
	}
	for _, file := range fileInfos {
		if !file.IsDir()&&strings.HasSuffix(file.Name(),".md"){
			filename:=path.Join(dir,file.Name())
			formatDoc(filename)
		}
	}
	time.Sleep(time.Second)
	fmt.Println()
	fmt.Println("You jump，I jump！")
}

func formatDoc(filename string) {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	content,ok := format(string(file))
	if !ok {
		return
	}
	os.Truncate(filename, 0)
	write(filename, content)
	fmt.Println("-- -- >",filename)
}

func format(content string) (string,bool) {
	s := content
	result := jsonRe.FindAllStringSubmatch(content, -1)
	if len(result)==0 {
		return s,false
	}
	for _, v := range result {
		segment := v[1]
		res := commentRe.FindAllStringSubmatch(segment, -1)
		var maxLength int
		for _, k := range res {
			contentLength := utf8.RuneCountInString(k[0]) - utf8.RuneCountInString(k[1])
			if contentLength > maxLength {
				maxLength = contentLength
			}
		}
		for _, k := range res {
			row, comment := k[0], k[1]
			if strings.Contains(row, `"logo"`) || strings.Contains(row, `"icon"`) {
				link := imgRe.FindStringSubmatch(row)
				segment = strings.ReplaceAll(segment, link[2], "https://example.png")
				//fmt.Println(fmt.Sprintf("replace: %s", link[2]))
				continue
			}
			contentLength := utf8.RuneCountInString(k[0]) - utf8.RuneCountInString(k[1])
			spacing := maxLength - contentLength + 5
			row = strings.ReplaceAll(row, comment, strings.Repeat(" ", spacing)+comment)
			segment = strings.ReplaceAll(segment, k[0], row)
		}
		s = strings.ReplaceAll(s, v[1], segment)
	}
	return s,true
}

func write(filename, content string) {
	var f *os.File
	var err1 error
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		f, err1 = os.Create(filename)
	} else {
		f, err1 = os.OpenFile(filename, os.O_RDWR|os.O_APPEND, os.ModePerm)
	}
	if err1 != nil {
		panic(err1)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString(content)
	w.Flush()
}
