package update

import (
	"fmt"
	"github.com/cavaliercoder/grab"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
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

func Install(filepath string, url string) (config string, err error) {

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
	/*	_, err = io.Copy(out, resp.Body)
		if err != nil {
			return config, err
		}
	*/

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)

	//	cfg, err := config.ReadDefault(filepath)

	log.Printf("即将放回 的信息   %s", responseString)

	return responseString, nil
}

func main() {
	// create client
	client := grab.NewClient()
	req, _ := grab.NewRequest(".", "http://www.golang-book.com/public/pdf/gobook.pdf")

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Download saved to ./%v \n", resp.Filename)

	// Output:
	// Downloading http://www.golang-book.com/public/pdf/gobook.pdf...
	//   200 OK
	//   transferred 42970 / 2893557 bytes (1.49%)
	//   transferred 1207474 / 2893557 bytes (41.73%)
	//   transferred 2758210 / 2893557 bytes (95.32%)
	// Download saved to ./gobook.pdf
}
