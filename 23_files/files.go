package main

import (
	"fmt"
	"os"
)

// Advise do not use ReadFile because it loads entire data in once if the file size is large it can consume entire rss of your application which will lead to application crash....

func main() {
	fmt.Println("File System")

	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }

	// defer f.Close()

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }

	// fmt.Println("File Name", fileInfo.Name())
	// fmt.Println("File or Folder", fileInfo.IsDir())
	// fmt.Println("File Size", fileInfo.Size())
	// fmt.Println("File Mode", fileInfo.Mode())
	// fmt.Println("File Modified", fileInfo.ModTime())

	// buf := make([]byte, 20)

	// d, err := f.Read(buf)
	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	println("Data Buffer", d, string(buf[i]))
	// }

	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }

	// fmt.Print(string(data))

	//...........Read Folders
	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), "---", fi.IsDir())
	// }

	//...........Create a File

	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("Hi Golang")
	// f.WriteString(" Nice Language")

	// bytes := []byte("Hello Golang")
	// f.Write(bytes)

	// ................ Read and Write to another file (streaming fashion)

	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example3.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	by, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(by)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }

	// writer.Flush()

	// ..............Delete a file

	// sourceFile, err := os.Open("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File Deleted Successfully")
}
