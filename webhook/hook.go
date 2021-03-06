package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/ghaoo/githook/utils"
	"github.com/ghaoo/githook/utils/logger"
	"github.com/ghaoo/githook"
)

var log *logger.Log

var root string

var port = flag.String("port", "9900", "web hook listen port")
var gitDir = flag.String("dir", "", "repository dir name")
var storeDir = flag.String("store", githook.DEFAULT_STORE, "store dir name")

func main() {

	if !flag.Parsed() {
		flag.Parse()
	}

	root = *gitDir

	http.HandleFunc("/", webhook)
	log.Infof("监听端口：%s", *port)

	http.ListenAndServe(":"+*port, nil)
}

type Commits struct {
	Message string
	Url     string
}

type Res struct {
	Commits []Commits

	Repository struct {
		Name      string
		UpdatedAt time.Time
	}
}

func webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("read body err, %v\n", err)
			return
		}
		defer r.Body.Close()

		var res Res

		err = json.Unmarshal(body, &res)
		if err != nil {
			log.Errorf("unmarshal json err, %v\n", err)
			return
		}

		if root == "" {
			root = res.Repository.Name
		}

		if err = githook.Pull(root); err != nil {
			log.Errorf("更新失败：%v", err)
			return
		}

		res.Repository.UpdatedAt = time.Now()

		store, err := utils.NewStore(*storeDir, "hook.json")

		if err != nil {
			log.Errorf("创建储存引擎失败：%v", err)
			return
		}

		if err = store.Store(&res); err != nil {
			log.Errorf("储存数据失败：%v", err)
			return
		}
	} else {
		w.Write([]byte("123"))
	}

}

func init() {

	// 初始化
	log = logger.NewLog(1000)

	// 设置log级别
	log.SetLevel("Debug")

	// 设置输出引擎
	log.SetEngine("file", `{"level":4, "spilt":"size", "filename":"`+*storeDir+`/hook.log", "maxsize":10}`)
}

