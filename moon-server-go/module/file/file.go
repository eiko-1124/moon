package file

import (
	"bytes"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"moon-manage-serve/module/file/entry"
	diarydto "moon-manage-serve/sql/diaryDto"
	filedto "moon-manage-serve/sql/fileDto"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo"
)

var fileTypeMap sync.Map
var allFileQueue map[string]*entry.FileQueue = make(map[string]*entry.FileQueue)

func SetDb(Db *sql.DB) {
	filedto.SetDb(Db)
}

func GetAllFiles(ctx echo.Context) error {
	page := ctx.QueryParam("page")
	cpage, _ := strconv.Atoi(page)
	files, err := filedto.QueryAllFile(cpage)
	if err != nil {
		fmt.Println(err.Error())
	}
	data, _ := json.Marshal(files)
	return ctx.String(http.StatusOK, string(data))
}

func UploadMinorFile(ctx echo.Context) error {
	res := new(entry.UpdateRes)
	name := ctx.FormValue("fileName")
	hasName := true
	var err error
	for hasName {
		hasName, err = filedto.QueryFileName(name)
		if err != nil {
			fmt.Println(err.Error())
			res.Res = "0"
			data, _ := json.Marshal(res)
			return ctx.String(http.StatusOK, string(data))
		}
		if hasName {
			str := ""
			strArray := strings.Split(name, ".")
			length := len(strArray)
			for i := 0; i < length-2; i++ {
				str = str + strArray[i] + "."
			}
			name = str + strArray[length-2] + "(1)." + strArray[length-1]
		}
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
		res.Res = "0"
		data, _ := json.Marshal(res)
		return ctx.String(http.StatusOK, string(data))
	}
	path, err := UpdateFile(name, file)
	if err != nil {
		fmt.Println(err.Error())
		res.Res = "0"
		data, _ := json.Marshal(res)
		return ctx.String(http.StatusOK, string(data))
	}
	link := "http://localhost:1323/" + path
	err = filedto.SaveFile(name, link)
	if err != nil {
		fmt.Println(err.Error())
		res.Res = "0"
		data, _ := json.Marshal(res)
		return ctx.String(http.StatusOK, string(data))
	}
	res.Res = "1"
	res.File.Name = name
	res.File.Link = link
	diarydto.InsertDiary("你上传了文件“"+name+"”", 0)
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func UploadFilePiece(ctx echo.Context) error {
	name := ctx.FormValue("fileName")
	id := ctx.FormValue("id")
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.String(http.StatusOK, "0")
	}
	_, flag := allFileQueue[name]
	if !flag {
		var fileQueue entry.FileQueue
		fileQueue.FileName = strconv.FormatInt(time.Now().Unix(), 10)
		fileQueue.Order = 0
		fileQueue.Queue = make(map[int]*multipart.FileHeader)
		allFileQueue[name] = &fileQueue
	}
	queue := allFileQueue[name]
	nId, _ := strconv.Atoi(id)
	queue.Queue[nId] = file
	for queue.Order == nId {
		_, flag = queue.Queue[nId+1]
		if flag {
			nId++
		}
		queue.Order = nId + 1
		fmt.Println(nId)
	}
	return ctx.String(http.StatusOK, "hello world")
}

func UpdateMoodPicture(ctx echo.Context) error {
	res := new(entry.PicRes)
	file, err := ctx.FormFile("file")
	if err != nil {
		res.Res = 0
	} else {
		path, today, err := UpdateMoodImage(file)
		if err != nil {
			res.Res = 0
		} else {
			id, err := filedto.SaveMoodImage(path, today)
			if err != nil {
				res.Res = 0
			} else {
				res.Res = 1
				res.Id = strconv.Itoa(id)
			}
		}
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetArticleCoverSum(ctx echo.Context) error {
	res := new(entry.SumRes)
	sum, err := filedto.QueryArticleCoverSum()
	if err != nil {
		res.Res = 0
		res.Sum = 0
	} else {
		res.Res = 1
		res.Sum = sum
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetArticleCovers(ctx echo.Context) error {
	res := new(entry.AcoverRes)
	page := ctx.QueryParam("page")
	covers, err := filedto.QueryArticleCover(page)
	if err != nil {
		res.Res = 0
		res.Covers = nil
	} else {
		res.Res = 1
		res.Covers = covers
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetMoodImageSum(ctx echo.Context) error {
	res := new(entry.SumRes)
	sum, err := filedto.QueryMoodImageSum()
	if err != nil {
		res.Res = 0
		res.Sum = 0
	} else {
		res.Res = 1
		res.Sum = sum
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetMoodImages(ctx echo.Context) error {
	res := new(entry.ImageRes)
	page := ctx.QueryParam("page")
	images, err := filedto.QueryMoodImages(page)
	if err != nil {
		res.Res = 0
		res.Images = nil
	} else {
		res.Res = 1
		res.Images = images
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetOtherFileSum(ctx echo.Context) error {
	res := new(entry.SumRes)
	sum, err := filedto.QueryOtherFilesSum()
	if err != nil {
		res.Res = 0
		res.Sum = 0
	} else {
		res.Res = 1
		res.Sum = sum
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

func GetOtherFiles(ctx echo.Context) error {
	res := new(entry.OtherFilesRes)
	page := ctx.QueryParam("page")
	files, err := filedto.QueryOtherFiles(page)
	if err != nil {
		res.Res = 0
		res.Files = nil
	} else {
		res.Res = 1
		res.Files = files
	}
	data, _ := json.Marshal(res)
	return ctx.String(http.StatusOK, string(data))
}

//	保存文件
func UpdateFile(name string, file *multipart.FileHeader) (string, error) {
	toDay := time.Now().Format("2006-01-02")
	basePath := "./assets/file/" + toDay + "/"
	src, err := file.Open()
	if err != nil {
		fmt.Println("file open fail", err.Error())
		return "", err
	}
	defer src.Close()
	_, err = CreateExists(basePath)
	if err != nil {
		fmt.Printf("Exists create faied,error[%v]", err.Error())
		return "", err
	}
	dst, err := os.Create(basePath + name)
	if err != nil {
		fmt.Println("create file fail", err.Error())
		return "", err
	}
	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println("save file fail", err.Error())
		return "", err
	}
	defer dst.Close()
	return "static/file/" + toDay + "/" + name, nil
}

// GetFileType 用文件前面几个字节来判断
// fSrc: 文件字节流（就用前面几个字节）
func GetFileType(fSrc []byte) string {
	var fileType string
	fileCode := bytesToHexString(fSrc)

	fileTypeMap.Range(func(key, value interface{}) bool {
		k := key.(string)
		v := value.(string)
		if strings.HasPrefix(fileCode, strings.ToLower(k)) ||
			strings.HasPrefix(k, strings.ToLower(fileCode)) {
			fileType = v
			return false
		}
		return true
	})
	return fileType
}

// 获取前面结果字节的二进制
func bytesToHexString(src []byte) string {
	res := bytes.Buffer{}
	if src == nil || len(src) <= 0 {
		return ""
	}
	temp := make([]byte, 0)
	for _, v := range src {
		sub := v & 0xFF
		hv := hex.EncodeToString(append(temp, sub))
		if len(hv) < 2 {
			res.WriteString(strconv.FormatInt(int64(0), 10))
		}
		res.WriteString(hv)
	}
	return res.String()
}

// PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CreateExists 创建文件夹
func CreateExists(path string) (bool, error) {
	_dir := path
	exist, err := PathExists(_dir)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return false, nil
	}
	if exist {
		return true, nil
	} else {
		fmt.Printf("no dir![%v]\n", _dir)
		// 创建文件夹
		err := os.Mkdir(_dir, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
			return false, nil
		} else {
			return true, nil
		}
	}
}

func UpdateMoodImage(file *multipart.FileHeader) (string, string, error) {
	toDay := time.Now().Format("2006-01-02")
	basePath := "./assets/mood/" + toDay + "/"
	name := strconv.FormatInt(time.Now().Unix(), 10)
	src, err := file.Open()
	if err != nil {
		fmt.Println("file open fail", err.Error())
		return "", "", err
	}
	defer src.Close()
	fSrc, err := ioutil.ReadAll(src)
	if err != nil {
		fmt.Println("file get type fail", err.Error())
		return "", "", err
	}
	fileType := GetFileType(fSrc)
	src, err = file.Open()
	if err != nil {
		fmt.Println("file open fail", err.Error())
		return "", "", err
	}
	defer src.Close()
	_, err = CreateExists(basePath)
	if err != nil {
		fmt.Printf("Exists create faied,error[%v]", err.Error())
		return "", "", err
	}
	dst, err := os.Create(basePath + name + "." + fileType)
	if err != nil {
		fmt.Println("create file fail", err.Error())
		return "", "", err
	}
	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println("save file fail", err.Error())
		return "", "", err
	}
	defer dst.Close()
	return "static/mood/" + toDay + "/" + name + "." + fileType, toDay, nil
}

func init() {
	fileTypeMap.Store("ffd8ffe000104a464946", "jpg")  //JPEG (jpg)
	fileTypeMap.Store("89504e470d0a1a0a0000", "png")  //PNG (png)
	fileTypeMap.Store("47494638396126026f01", "gif")  //GIF (gif)
	fileTypeMap.Store("49492a00227105008037", "tif")  //TIFF (tif)
	fileTypeMap.Store("424d228c010000000000", "bmp")  //16色位图(bmp)
	fileTypeMap.Store("424d8240090000000000", "bmp")  //24位位图(bmp)
	fileTypeMap.Store("424d8e1b030000000000", "bmp")  //256色位图(bmp)
	fileTypeMap.Store("41433130313500000000", "dwg")  //CAD (dwg)
	fileTypeMap.Store("3c21444f435459504520", "html") //HTML (html)   3c68746d6c3e0  3c68746d6c3e0
	fileTypeMap.Store("3c68746d6c3e0", "html")        //HTML (html)   3c68746d6c3e0  3c68746d6c3e0
	fileTypeMap.Store("3c21646f637479706520", "htm")  //HTM (htm)
	fileTypeMap.Store("48544d4c207b0d0a0942", "css")  //css
	fileTypeMap.Store("696b2e71623d696b2e71", "js")   //js
	fileTypeMap.Store("7b5c727466315c616e73", "rtf")  //Rich Text Format (rtf)
	fileTypeMap.Store("38425053000100000000", "psd")  //Photoshop (psd)
	fileTypeMap.Store("46726f6d3a203d3f6762", "eml")  //Email [Outlook Express 6] (eml)
	fileTypeMap.Store("d0cf11e0a1b11ae10000", "doc")  //MS Excel 注意：word、msi 和 excel的文件头一样
	fileTypeMap.Store("d0cf11e0a1b11ae10000", "vsd")  //Visio 绘图
	fileTypeMap.Store("5374616E64617264204A", "mdb")  //MS Access (mdb)
	fileTypeMap.Store("252150532D41646F6265", "ps")
	fileTypeMap.Store("255044462d312e350d0a", "pdf")  //Adobe Acrobat (pdf)
	fileTypeMap.Store("2e524d46000000120001", "rmvb") //rmvb/rm相同
	fileTypeMap.Store("464c5601050000000900", "flv")  //flv与f4v相同
	fileTypeMap.Store("00000020667479706d70", "mp4")
	fileTypeMap.Store("49443303000000002176", "mp3")
	fileTypeMap.Store("000001ba210001000180", "mpg") //
	fileTypeMap.Store("3026b2758e66cf11a6d9", "wmv") //wmv与asf相同
	fileTypeMap.Store("52494646e27807005741", "wav") //Wave (wav)
	fileTypeMap.Store("52494646d07d60074156", "avi")
	fileTypeMap.Store("4d546864000000060001", "mid") //MIDI (mid)
	fileTypeMap.Store("504b0304140000000800", "zip")
	fileTypeMap.Store("526172211a0700cf9073", "rar")
	fileTypeMap.Store("235468697320636f6e66", "ini")
	fileTypeMap.Store("504b03040a0000000000", "jar")
	fileTypeMap.Store("4d5a9000030000000400", "exe")        //可执行文件
	fileTypeMap.Store("3c25402070616765206c", "jsp")        //jsp文件
	fileTypeMap.Store("4d616e69666573742d56", "mf")         //MF文件
	fileTypeMap.Store("3c3f786d6c2076657273", "xml")        //xml文件
	fileTypeMap.Store("494e5345525420494e54", "sql")        //xml文件
	fileTypeMap.Store("7061636b616765207765", "java")       //java文件
	fileTypeMap.Store("406563686f206f66660d", "bat")        //bat文件
	fileTypeMap.Store("1f8b0800000000000000", "gz")         //gz文件
	fileTypeMap.Store("6c6f67346a2e726f6f74", "properties") //bat文件
	fileTypeMap.Store("cafebabe0000002e0041", "class")      //bat文件
	fileTypeMap.Store("49545346030000006000", "chm")        //bat文件
	fileTypeMap.Store("04000000010000001300", "mxp")        //bat文件
	fileTypeMap.Store("504b0304140006000800", "docx")       //docx文件
	fileTypeMap.Store("d0cf11e0a1b11ae10000", "wps")        //WPS文字wps、表格et、演示dps都是一样的
	fileTypeMap.Store("6431303a637265617465", "torrent")
	fileTypeMap.Store("6D6F6F76", "mov")         //Quicktime (mov)
	fileTypeMap.Store("FF575043", "wpd")         //WordPerfect (wpd)
	fileTypeMap.Store("CFAD12FEC5FD746F", "dbx") //Outlook Express (dbx)
	fileTypeMap.Store("2142444E", "pst")         //Outlook (pst)
	fileTypeMap.Store("AC9EBD8F", "qdf")         //Quicken (qdf)
	fileTypeMap.Store("E3828596", "pwl")         //Windows Password (pwl)
	fileTypeMap.Store("2E7261FD", "ram")         //Real Audio (ram)
}
