package system

import "flag"

var root string

var gitDir = flag.String("dir", "", "repository dir name")
var port = flag.Int64("t", 5, "git pull frequency(/s)")

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	root = *gitDir
}

func git_pull(dir string) {
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
}