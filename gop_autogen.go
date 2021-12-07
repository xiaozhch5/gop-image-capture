package main

import (
	fmt "fmt"
	strings "strings"
	json "encoding/json"
	template "html/template"
	io "io"
	ioutil "io/ioutil"
	log "log"
	http "net/http"
	os "os"
)

type Hit struct {
	ID              int    "json:\"id\""
	PageURL         string "json:\"pageURL\""
	Type            string "json:\"type\""
	Tags            string "json:\"tags\""
	PreviewURL      string "json:\"previewURL\""
	PreviewWidth    int    "json:\"previewWidth\""
	PreviewHeight   int    "json:\"previewHeight\""
	WebformatURL    string "json:\"webformatURL\""
	WebformatWidth  int    "json:\"webformatWidth\""
	WebformatHeight int    "json:\"webformatHeight\""
	LargeImageURL   string "json:\"largeImageURL\""
	FullHDURL       string "json:\"fullHDURL\""
	ImageURL        string "json:\"imageURL\""
	ImageWidth      int    "json:\"imageWidth\""
	ImageHeight     int    "json:\"imageHeight\""
	ImageSize       int    "json:\"imageSize\""
	Views           int    "json:\"views\""
	Downloads       int    "json:\"downloads\""
	Likes           int    "json:\"likes\""
	Comments        int    "json:\"comments\""
	UserID          int    "json:\"user_id\""
	User            string "json:\"user\""
	UserImageURL    string "json:\"userImageURL\""
}
type Result struct {
	Total     int   "json:\"total\""
	TotalHits int   "json:\"totalHits\""
	Hits      []Hit "json:\"hits\""
}

func CreatePathIfNotExists(path string) {
//line C:\bigdata\image-capture\main.gop:46
	_, err := os.Stat(path)
//line C:\bigdata\image-capture\main.gop:47
	if os.IsNotExist(err) {
//line C:\bigdata\image-capture\main.gop:48
		os.Mkdir(path, os.ModePerm)
	}
}
func DownloadImages(w http.ResponseWriter, r *http.Request) {
//line C:\bigdata\image-capture\main.gop:54
	if r.Method == "GET" {
//line C:\bigdata\image-capture\main.gop:55
		t, _ := template.ParseFiles("downloadImages.gtpl")
//line C:\bigdata\image-capture\main.gop:56
		log.Println(t.Execute(w, nil))
	} else {
//line C:\bigdata\image-capture\main.gop:58
		err := r.ParseForm()
//line C:\bigdata\image-capture\main.gop:59
		if err != nil {
//line C:\bigdata\image-capture\main.gop:60
			log.Fatal("ParseForm: ", err)
		}
//line C:\bigdata\image-capture\main.gop:63
		fmt.Println("key:", r.Form["key"])
//line C:\bigdata\image-capture\main.gop:64
		fmt.Println("q:", r.Form["q"])
//line C:\bigdata\image-capture\main.gop:65
		fmt.Println("image_type:", r.Form["image_type"])
//line C:\bigdata\image-capture\main.gop:66
		fmt.Println("save_path:", r.Form["save_path"])
//line C:\bigdata\image-capture\main.gop:68
		_, keyOk := r.Form["key"]
//line C:\bigdata\image-capture\main.gop:69
		_, qOk := r.Form["q"]
//line C:\bigdata\image-capture\main.gop:70
		_, imageTypeOk := r.Form["image_type"]
//line C:\bigdata\image-capture\main.gop:71
		_, savePathOk := r.Form["save_path"]
//line C:\bigdata\image-capture\main.gop:72
		if !keyOk || !qOk || !imageTypeOk || !savePathOk {
//line C:\bigdata\image-capture\main.gop:73
			fmt.Fprintf(w, "请输入正确的参数")
//line C:\bigdata\image-capture\main.gop:74
			return
		}
//line C:\bigdata\image-capture\main.gop:77
		resp, err := http.Get("https://pixabay.com/api/?key=" + r.Form["key"][0] + "&q=" + r.Form["q"][0] + "&image_type=" + r.Form["image_type"][0])
//line C:\bigdata\image-capture\main.gop:79
		if err != nil {
//line C:\bigdata\image-capture\main.gop:80
			panic(err)
		}
//line C:\bigdata\image-capture\main.gop:83
		var res Result
//line C:\bigdata\image-capture\main.gop:85
		defer resp.Body.Close()
//line C:\bigdata\image-capture\main.gop:87
		body, _ := ioutil.ReadAll(resp.Body)
//line C:\bigdata\image-capture\main.gop:89
		json.Unmarshal(body, &res)
		for
//line C:\bigdata\image-capture\main.gop:90
		i, hit := range res.Hits {
//line C:\bigdata\image-capture\main.gop:91
			fmt.Println(i, hit.LargeImageURL)
//line C:\bigdata\image-capture\main.gop:93
			var iURLSplit = strings.Split(hit.LargeImageURL, "/")
//line C:\bigdata\image-capture\main.gop:95
			imageResp, err := http.Get(hit.LargeImageURL)
//line C:\bigdata\image-capture\main.gop:97
			defer resp.Body.Close()
//line C:\bigdata\image-capture\main.gop:99
			CreatePathIfNotExists(r.Form["save_path"][0])
//line C:\bigdata\image-capture\main.gop:101
			if err != nil {
//line C:\bigdata\image-capture\main.gop:102
				panic(err)
			}
//line C:\bigdata\image-capture\main.gop:105
			out, err := os.Create(r.Form["save_path"][0] + iURLSplit[len(iURLSplit)-1])
//line C:\bigdata\image-capture\main.gop:106
			if err != nil {
//line C:\bigdata\image-capture\main.gop:107
				panic(err)
			}
//line C:\bigdata\image-capture\main.gop:109
			defer out.Close()
//line C:\bigdata\image-capture\main.gop:110
			_, err = io.Copy(out, imageResp.Body)
//line C:\bigdata\image-capture\main.gop:111
			if err != nil {
//line C:\bigdata\image-capture\main.gop:112
				panic(err)
			}
		}
//line C:\bigdata\image-capture\main.gop:115
		fmt.Fprintf(w, "图片下载完成,下载路径为:"+r.Form["save_path"][0])
	}
}
func main() {
//line C:\bigdata\image-capture\main.gop:119
	http.HandleFunc("/", DownloadImages)
//line C:\bigdata\image-capture\main.gop:120
	err := http.ListenAndServe(":8888", nil)
//line C:\bigdata\image-capture\main.gop:122
	if err != nil {
//line C:\bigdata\image-capture\main.gop:123
		log.Fatal("ListenAndServe: ", err)
	}
}
