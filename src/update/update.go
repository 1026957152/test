package update

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type UpdateConfig struct {
	Images []string
}

type updateConfig struct {
}

func DownloadFile_(filepath string, url string) (config UpdateConfig, err error) {

	var steps map[string]string = make(map[string]string)
	steps["step1"] = "下载其他"
	steps["step1"] = "安装服务"
	steps["step2"] = "下载mysql"
	steps["step3"] = "下载mqtt"
	steps["step4"] = "下载本地应用"

	steps["step2"] = "加载 local 配置信息"

	steps["step2"] = "启动串口中兴网络"
	steps["step2"] = "启动二维码扫描"
	steps["step2"] = "启动rfid 射频识别"
	steps["step2"] = "启动web 服务后台"
	steps["step2"] = "启动 摄像头 服务"

	steps["step2"] = "启动mqtt命令控制"

	//  mqtt: 下发配置命令。 1 更改local配置信息

	steps["step2"] = "启动本地 mqtt命令控制"

	log.Printf("Create the file")

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		panic(err)
		return config, err
	}
	log.Printf("out, err := os.Create(filepath)")

	defer out.Close()
	log.Printf("os.Create(filepath) docker")

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("获取远程资源 失败 %s", url)

		return config, err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {

		log.Printf("获取远程资源 失败 %s", url)

		return config, fmt.Errorf("bad status: %s", resp.Status)
	}

	log.Printf("获取远程资源 成功  %s", url)

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return config, err
	}

	//	cfg, err := config.ReadDefault(filepath)

	source, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", config.Images)

	return config, nil
}
