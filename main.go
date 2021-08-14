package main

/*
importing necessary packages
*/
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"

	Custom "gohartojson/src/custompackages"
)

var mapstringinterface map[string]interface{}
var arrayinterface []interface{}

func main() {
	// fmt.Println("test")
	filename := "test_api.har"
	var filedata []byte
	filedata, errobj := readfile(filename)
	if errobj != nil {
		pc, filename, lineno, okstatus := runtime.Caller(0)
		logerror(fmt.Sprint(pc), filename, fmt.Sprint(lineno), okstatus, errobj)
		return
	}
	var filecont map[string]interface{}
	errobj = json.Unmarshal(filedata, &filecont)
	if errobj != nil {
		pc, filename, lineno, okstatus := runtime.Caller(0)
		logerror(fmt.Sprint(pc), filename, fmt.Sprint(lineno), okstatus, errobj)
		return
	}

	log, _ := ConvertData(filecont["log"], "Map")
	_, entries := ConvertData(log["entries"], "Interface")
	objarray := processinterface_array(entries)
	// fmt.Println(errobj)
	// fmt.Println(objarray)
	var MObj Custom.FullObj
	MObj.Info.Name = "Data"
	MObj.Info.Schema = "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	MObj.Item = objarray
	writefile(MObj)
}

func writefile(objarray Custom.FullObj) {
	filename := "test_api.json"
	filedata, errobj := json.Marshal(objarray)
	if errobj != nil {
		pc, filename, lineno, okstatus := runtime.Caller(0)
		logerror(fmt.Sprint(pc), filename, fmt.Sprint(lineno), okstatus, errobj)
		return
	}
	errobj = ioutil.WriteFile(filename, filedata, 0644)
	if errobj != nil {
		pc, filename, lineno, okstatus := runtime.Caller(0)
		logerror(fmt.Sprint(pc), filename, fmt.Sprint(lineno), okstatus, errobj)
		return
	}
}

func ConvertData(Data interface{}, Pty string) (Res map[string]interface{}, Resa []interface{}) {
	a, errobj := json.Marshal(Data)
	if errobj != nil {
		pc, filename, lineno, okstatus := runtime.Caller(0)
		logerror(fmt.Sprint(pc), filename, fmt.Sprint(lineno), okstatus, errobj)
		return
	}
	if Pty == "Map" {
		errobj = json.Unmarshal(a, &Res)
	}
	if Pty == "Interface" {
		errobj = json.Unmarshal(a, &Resa)
	}
	if errobj != nil {
		pc, filename, lineno, okstatus := runtime.Caller(0)
		logerror(fmt.Sprint(pc), filename, fmt.Sprint(lineno), okstatus, errobj)
		return
	}
	return
}

func readfile(filename string) (filedata []byte, errobj error) {
	filedata, errobj = ioutil.ReadFile(filename)
	return
}

func processinterface_array(apidata []interface{}) (collection []Custom.ParaCol) {
	for _, val := range apidata {
		// fmt.Println(key, " : ", reflect.TypeOf(val), " : ")
		a, _ := ConvertData(val, "Map")
		// fmt.Println("aaaaaaaaaaaaa : ", request["request"], "aaaaaaaaaaaaa", reflect.TypeOf(request["request"]))
		request, _ := ConvertData(a["request"], "Map")
		// fmt.Println("aaaaaaaaaaaaa : ", request, "aaaaaaaaaaaaa", reflect.TypeOf(request))
		var obj Custom.ParaCol
		var nm string
		obj.Request.Method = getMethod(request)
		obj.Request.Url.Raw, obj.Request.Url.ProtoCol, obj.Request.Url.Host, obj.Request.Url.Port, obj.Request.Url.Path, nm = getUrlData(request)
		obj.Name = obj.Request.Method + " " + nm + " successfully"
		obj.Request.Header = getHeader(request)
		obj.Request.Body.Options.Raw.Language = "json"
		obj.Request.Body.Mode = "raw"
		obj.Request.Body.Raw = getBody(request)
		collection = append(collection, obj)
	}
	return
}
func getBody(data map[string]interface{}) (bbody string) {
	if data["postData"] != nil {
		a := data["postData"].(map[string]interface{})
		fmt.Println(a)
		fmt.Println(a["text"])
		bbody = fmt.Sprint(a["text"])
	}
	return
}

func getHeader(data map[string]interface{}) (Header []Custom.HeaderCol) {
	_, dataobj := ConvertData(data["headers"], "Interface")
	// fmt.Println(dataobj)
	for i := 0; i < len(dataobj); i = i + 1 {
		// fmt.Println(dataobj[i])
		hdt := dataobj[i].(map[string]interface{})
		var hd Custom.HeaderCol
		hd.Key = fmt.Sprint(hdt["name"])
		hd.Value = fmt.Sprint(hdt["value"])
		Header = append(Header, hd)
	}
	// fmt.Println(Header)
	return
}

func getUrlData(data map[string]interface{}) (Raw string, ProtoCol string, Host []string, Port string, path []string, nm string) {
	Raw, ProtoCol = "", ""
	// Host, path = []string, []string
	Raw = fmt.Sprint(data["url"])
	ProtoCol = "http"
	Host, Port, path, nm = getURLDetails(Raw)
	// fmt.Println(Host, " : ", path)
	return
}

func getURLDetails(url string) (Host []string, Port string, path []string, nm string) {
	urlparts := strings.Split(url, "/")
	// fmt.Println(urlparts)
	pts := urlparts[3:]
	for i := 0; i < len(pts); i = i + 1 {
		path = append(path, pts[i])
		nm = nm + "/" + pts[i]
	}
	hst := strings.Split(urlparts[2], ":")
	if len(hst) > 1 {
		Port = hst[1]
	}
	mainhst := strings.Split(hst[0], ".")
	for j := 0; j < len(mainhst); j = j + 1 {
		Host = append(Host, mainhst[j])
		//)
	}
	return
}
func getMethod(data map[string]interface{}) (method string) {
	method = ""
	// fmt.Println(data["method"], " : ", reflect.TypeOf(data["method"]))
	method = fmt.Sprint(data["method"])
	return
}

func logerror(pc string, filename string, lineno string, okstatus bool, errobj error) {
	fmt.Println(pc, " : ", filename, " : ", lineno, " : ", okstatus, " : ", errobj)
}
