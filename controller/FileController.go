package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const BASE_NAME = "./static/file/"

func RenderView (context *gin.Context) {
	println(">>>> render to file upload view action start <<<<")
	context.Header("Content-Type", "text/html; charset=utf-8")

	context.HTML(200,"fileUpload.html",gin.H{})
}

func FormUpload (context *gin.Context) {
	println(">>>> upload file by form action start <<<<")

	fh,err := context.FormFile("file")
	checkError(err)
	fileName := fh.Filename
	//context.SaveUploadedFile(fh,BASE_NAME + fh.Filename)

	file,err := fh.Open()
	defer file.Close()
	bytes,e := ioutil.ReadAll(file)
	e = ioutil.WriteFile(BASE_NAME + fileName,bytes,0666)
	checkError(e)

	if e != nil {
		context.JSON(200,gin.H{
			"success":false,
		})
	} else {
		context.JSON(200,gin.H{
			"success":true,
		})
	}
}

func MultiUpload(context *gin.Context) {
	println(">>>> upload file by form action start <<<<")
	form,err := context.MultipartForm()
	checkError(err)
	files := form.File["file"]

	var er error
	for _,f := range files {

		// 使用gin自带保存文件方法
		er = context.SaveUploadedFile(f,BASE_NAME + f.Filename)
		checkError(err)
	}
	if er != nil {
		context.JSON(200,gin.H{
			"success":false,
		})
	} else {
		context.JSON(200,gin.H{
			"success":true,
		})
	}

}

func Base64Upload (context *gin.Context) {
	println(">>>> upload file by base64 string action start <<<<")

	bytes,err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		log.Fatal(err)
	}

	strs := strings.Split(string(bytes),",")
	head := strs[0]
	body := strs[1]
	println(head + " | " + body)
	start := strings.LastIndex(head,"/")
	end := strings.LastIndex(head,";")
	tp := head[start + 1:end]

	err = ioutil.WriteFile(BASE_NAME + strconv.Itoa(time.Now().Nanosecond()) + "." + tp,[]byte(body),0666)
	checkError(err)
	//bys,err := base64.StdEncoding.DecodeString(string(bytes))
	//err = ioutil.WriteFile("./static/file/" + strconv.Itoa(time.Now().Nanosecond()),bys,0666)
	if err != nil {
		context.JSON(200,gin.H{
			"success":false,
		})
	} else {
		context.JSON(200,gin.H{
			"success":true,
		})
	}
}

func Download (context *gin.Context) {
	println(">>>> download file by base64 string action start <<<<")

	file, err := os.OpenFile("123.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "hello 沙河"
	file.Write([]byte(str))       //写入字节切片数据
	file.WriteString("\nhello 小王子") //直接写入字符串数据
	var tmp = make([]byte, 128)
	// 循环读取文件
	var content []byte
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}

	context.Writer.Write(content)
	context.Writer.Header().Add("Content-Type", "application/octet-stream")
	context.Writer.Header().Add("Content-Disposition", "attachment;filename="+"测试.txt")
	if err != nil {
		context.JSON(200,gin.H{
			"success":false,
		})
	} else {
		context.JSON(200,gin.H{
			"success":true,
		})
	}
}