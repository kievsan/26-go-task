package main

// 26.1 Программа-аналог cat

import (
	"fmt"
	"flag"
	"os"
	"io"
	"bufio"
	"time"
	"strings"
)

func getFile(fileName string) (resultStr string) {
	file, err := os.Open(fileName)
	res := err == nil
	
	if res {
		fmt.Printf("\nФайл %v открыт для чтения\n", fileName)
	} else {
		fmt.Printf("\nОшибка доступа к файлу %v:\n\t%v\n", fileName, err)
		file, err = os.Create(fileName)
		
		if err != nil {
		fmt.Printf("\nНе смогли создать файл %v\n%v\n", fileName, err)
		} else {
		writeFile(file)
		}    
	}  
	
	err = file.Close()
	if err != nil {
		panic(err)
	}
	
	file, _ = os.Open(fileName)
	defer file.Close()
	resultStr = ""  
	if res {
		resultStr = getResultStr(file)
	}
	
	return
}

func writeFile(file *os.File) {
	justNow := strings.Split(time.Now().String(), " ")
	date := justNow[0]
	time := strings.Split(justNow[1], ".")[0]
	marker := date + " " + time + "\t"
	
	writer := bufio.NewWriter(file)      
	size, err := writer.WriteString(marker + "File created...\n")
	
	if res := err == nil; !res {
		fmt.Println("\nОшибка записи в файл:\n\t", err)
	} else {
		writer.Flush()
		fmt.Printf("\nФайл %v ДОСТУПЕН!\tЗаписаны первые %v байт\n", file.Name(), size) 
	}
}

func getResultStr(file *os.File) (resultStr string) {
	resultBytes, err := io.ReadAll(file)
	if err == nil {
		resultStr = string(resultBytes)
		fmt.Printf("Файл %v прочитан!\n", file.Name())
	} else {
		fmt.Printf("Ошибка чтения файла %v:\n\t%v\n", file.Name(), err)
	}
	return
}


func main() {
	fmt.Println("Программа-аналог cat")

  // go run main.go file1 file2 file3
  // go run main.go log1.txt log2.txt log.txt
	
		flag.Parse()
	args := flag.Args()
	argsCount := len(args)  
	fmt.Println(argsCount, args)  

	contents := make([]string, argsCount)
	if argsCount == 0 {
		return
	} else {
		contents[0] = getFile(args[0])
	}

	count := 1
	if argsCount > 1 {
		contents[1] = getFile(args[1])
		count = 2
	}
	content := strings.Join(contents[:count], "")

	if argsCount < 3 {
		fmt.Println("\n" + content)
	} else {
		cat(content, args[2])
	}
}

func cat(content, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
    fmt.Println("\nНе смогли создать файл ", fileName)
    panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)      
	size, err := writer.WriteString(content)
	
	if res := err == nil; !res {
		fmt.Println("\nОшибка записи в файл:\n\t", err)
	} else {
		writer.Flush()
		fmt.Printf("\nФайл %v записан: %v байт\n", fileName, size) 
	}
}
