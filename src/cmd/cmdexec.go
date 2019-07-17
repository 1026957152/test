//go:generate goversioninfo
// -icon=testdata/resource/icon.ico
// -manifest=testdata/resource/goversioninfo.exe.manifest

package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/kataras/iris"
	"golang.org/x/crypto/ssh"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"path"
	"runtime"
	"strconv"
	"strings"

	//_ "github.com/ibmdb/go_ibm_db"
	//"github.com/akavel/rsrc"

	"net/http"
	"net/url"
)

type Data struct {
	Xmjp   string  `json:"xmjp"`
	Xb     string  `json:"xb"`
	Grzh   string  `json:"grzh"`
	Grbh   string  `json:"grbh"`
	Zjhlx  string  `json:"zjhlx"`
	Csrq   string  `json:"csrq"`
	Cjgzsj string  `json:"cjgzsj"`
	Grzhzt string  `json:"grzhzt"`
	Sqgzze float64 `json:"sqgzze"`
	Gjjye  float64 `json:"gjjye"`

	Zzdh    string `json:"zzdh"`
	Sjh     string `json:"sjh"`
	Jtdz    string `json:"jtdz"`
	Dwmc    string `json:"dwmc"`
	Dwdjh   string `json:"dwdjh"`
	Khglbbh string `json:"khglbbh"` //	"khglbbh": "09120001",

	Yyzz         string  `json:"yyzz"`         //	"yyzz": "                    ",
	Xzbm         string  `json:"xzbm"`         //"xzbm": "000001",
	Dwdz         string  `json:"dwdz"`         //"dwdz": "榆林市榆阳区文化南路市民大厦8层",
	Dwjezt       string  `json:"dwjezt"`       //"dwjezt": "01",
	Dwjczjy      string  `json:"dwjczjy"`      //"dwjczjy": "0",
	Dwlxfs       string  `json:"dwlxfs"`       //"dwlxfs": "3367893             ",
	Fxr          string  `json:"fxr"`          //"fxr": "01",
	Dwkhrq       string  `json:"dwkhrq"`       //"dwkhrq": "1899-12-31",
	Dwfdrdbr     string  `json:"dwfdrdbr"`     //"dwfdrdbr": "曹文波",
	Dwfdrzjlx    string  `json:"dwfdrzjlx"`    //"dwfdrzjlx": "",
	Dwfdrzjhm    string  `json:"dwfdrzjhm"`    //"dwfdrzjhm": "                    ",
	Cjnyr        string  `json:"cjnyr"`        //"cjnyr": "200512",
	Jznyr        string  `json:"jznyr"`        //"jznyr": "201906",
	Dwjcbl       string  `json:"dwjcbl"`       //"dwjcbl": "0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120",
	Grjcbl       string  `json:"grjcbl"`       //"grjcbl": "0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.050|0.050|0.050|0.050|0.050|0.050|0.050|0.050|0.050|0.050",
	Jccs         int     `json:"jccs"`         //"jccs": 73,
	Lxjcdwmc     string  `json:"lxjcdwmc"`     //"lxjcdwmc": "榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心",
	Lxdwjce      string  `json:"lxdwjce"`      //"lxdwjce": "683.34|683.34|683.34|683.34|683.34|683.34|638.88|638.88|638.88|638.88|638.88|638.88|638.88|638.88|638.88|638.88|638.88|638.88",
	Lxgrjce      string  `json:"lxgrjce"`      //"lxgrjce": "683.34|683.34|683.34|683.34|683.34|683.34|638.88|638.88|266.20|266.20|266.20|266.20|266.20|266.20|266.20|266.20|266.20|266.20",
	Zhycjcrq     string  `json:"zhycjcrq"`     //"zhycjcrq": "2019-07-01",
	Dqgjjnd      string  `json:"dqgjjnd"`      //"dqgjjnd": "",
	Tqljje       float64 `json:"tqljje"`       //"tqljje": 73430.76,
	Lxtqyy       string  `json:"lxtqyy"`       //"lxtqyy": "021|021|021|021|021|021|021|021|021|021|021|021||||||",
	Lxtqsj       string  `json:"lxtqsj"`       //"lxtqsj": "2019-06-05|2019-05-04|2019-04-04|2019-03-04|2019-02-04|2019-01-04|2018-12-04|2018-11-04|2018-10-04|2018-09-04|2018-08-04|2018-07-04||||||",
	Lxtqfs       string  `json:"lxtqfs"`       //"lxtqfs": "|||||||||||",
	Lxtqje       string  `json:"lxtqje"`       //"lxtqje": "1007.85|486.58|2176.03|2176.03|2176.03|2176.03|2176.03|2176.03|2176.03|2176.03|2176.03|2176.03||||||",
	Dkbh         string  `json:"dkbh"`         //"dkbh": "20170105991",
	Dkje         float64 `json:"dkje"`         //"dkje": 500000,
	Dkqx         string  `json:"dkqx"`         //"dkqx": "360",
	Dkhkfs       string  `json:"dkhkfs"`       //"dkhkfs": "01",
	Dkyhke       float64 `json:"dkyhke"`       //"dkyhke": 2176.03,
	Dkqsrq       string  `json:"dkqsrq"`       //"dkqsrq": "2017-12-04",
	Dkdqrq       string  `json:"dkdqrq"`       //"dkdqrq": "2047-12-03",
	Dkjqrq       string  `json:"dkjqrq"`       //"dkjqrq": "1899-12-31",
	Dkye         float64 `json:"dkye"`         //"dkye": 483998.08,
	Gtdkrxm      string  `json:"gtdkrxm"`      //"gtdkrxm": "",
	Gtdkrsfzh    string  `json:"gtdkrsfzh"`    //"gtdkrsfzh": "",
	Gtdkrsfjlgjj string  `json:"gtdkrsfjlgjj"` //"gtdkrsfjlgjj": "",
	Dkzt         string  `json:"dkzt"`         //"dkzt": "02",
	Dklsyqcs     int     `json:"dklsyqcs"`     //"dklsyqcs": 0,
	Zdlxyqcs     string  `json:"zdlxyqcs"`     //"zdlxyqcs": "0",
	Grgfdz       string  `json:"grgfdz"`       //"grgfdz": "榆阳区榆林大道410号亮馨苑小区2幢2单元1303室",
	Fwgmjszj     float64 `json:"fwgmjszj"`     //"fwgmjszj": 640000

	/*	"xmjp": "ZHAOYUAN",
		"xb": "M",
		"grzh": "612000044397",
		"grbh": "AP00044397",
		"zjhlx": "01",
		"csrq": "1984-09-21",
		"cjgzsj": "2012-04-06",
		"grzhzt": "01",
		"sqgzze": 5694.5,
		"gjjye": 1066.13,
		"zzdh": "13468801683         ",
		"sjh": "",
		"jtdz": "",
		"dwmc": "榆林市住房公积金中心",
		"dwdjh": "201000000089",
		"khglbbh": "09120001",
		"zzjgdm": "12312312-3          ",
		"yyzz": "                    ",
		"xzbm": "000001",
		"dwdz": "榆林市榆阳区文化南路市民大厦8层",
		"dwjezt": "01",
		"dwjczjy": "0",
		"dwlxfs": "3367893             ",
		"fxr": "01",
		"dwkhrq": "1899-12-31",
		"dwfdrdbr": "曹文波",
		"dwfdrzjlx": "",
		"dwfdrzjhm": "                    ",
		"cjnyr": "200512",
		"jznyr": "201906",
		"dwjcbl": "0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120",
		"grjcbl": "0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.120|0.050|0.050|0.050|0.050|0.050|0.050|0.050|0.050|0.050|0.050",
		"jccs": 73,
		"lxjcdwmc": "榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心|榆林市住房公积金中心",
		"lxdwjce": "683.34|683.34|683.34|683.34|683.34|683.34|638.88|638.88|638.88|638.88|638.88|638.88|638.88|638.88|638.88|638.88|638.88|638.88",
		"lxgrjce": "683.34|683.34|683.34|683.34|683.34|683.34|638.88|638.88|266.20|266.20|266.20|266.20|266.20|266.20|266.20|266.20|266.20|266.20",
		"zhycjcrq": "2019-07-01",
		"dqgjjnd": "",
		"tqljje": 73430.76,
		"lxtqyy": "021|021|021|021|021|021|021|021|021|021|021|021||||||",
		"lxtqsj": "2019-06-05|2019-05-04|2019-04-04|2019-03-04|2019-02-04|2019-01-04|2018-12-04|2018-11-04|2018-10-04|2018-09-04|2018-08-04|2018-07-04||||||",
		"lxtqfs": "|||||||||||",
		"lxtqje": "1007.85|486.58|2176.03|2176.03|2176.03|2176.03|2176.03|2176.03|2176.03|2176.03|2176.03|2176.03||||||",
		"dkbh": "20170105991",
		"dkje": 500000,
		"dkqx": "360",
		"dkhkfs": "01",
		"dkyhke": 2176.03,
		"dkqsrq": "2017-12-04",
		"dkdqrq": "2047-12-03",
		"dkjqrq": "1899-12-31",
		"dkye": 483998.08,
		"gtdkrxm": "",
		"gtdkrsfzh": "",
		"gtdkrsfjlgjj": "",
		"dkzt": "02",
		"dklsyqcs": 0,
		"zdlxyqcs": "0",
		"grgfdz": "榆阳区榆林大道410号亮馨苑小区2幢2单元1303室",
		"fwgmjszj": 640000
	*/
}
type ResultQuery struct {
	Ret  string `json:"ret"`
	Msg  string `json:"msg"`
	Data []Data `json:"data"`
}
type Numverify struct {
	Valid               bool   `json:"valid"`
	Number              string `json:"number"`
	LocalFormat         string `json:"local_format"`
	InternationalFormat string `json:"international_format"`
	CountryPrefix       string `json:"country_prefix"`
	CountryCode         string `json:"country_code"`
	CountryName         string `json:"country_name"`
	Location            string `json:"location"`
	Carrier             string `json:"carrier"`
	LineType            string `json:"line_type"`
}

func api(zjhm string, parse bool) (object interface{}, err error) {
	phone := "14158586273"
	// QueryEscape escapes the phone string so
	// it can be safely placed inside a URL query
	safePhone := url.QueryEscape(phone)

	postRequest := map[string]string{
		"appid":    "0001",
		"citybm":   "C61080",
		"sign":     "71c668c9e396f9a0e6513ceaa5635b0b779278bc",
		"xingming": "汪坤", "zjhlx": "0",
	}
	postRequest["zjhm"] = zjhm // "612724198409210339"

	requestBody, err := json.Marshal(postRequest)

	url := fmt.Sprintf("http://10.22.30.75:80/loan_share/public/xfd/gjjgrzhxxcx.service?access_key=YOUR_ACCESS_KEY&number=%s", safePhone)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		//log.Fatal("NewRequest: ", err)
		log.Printf("NewRequest: ", err)

		return ResultQuery{Ret: "5", Msg: "aaaa"}, err
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Do-----client.Do---: ", err)
		return "", err

		//	return "",err
	}
	//log.Printf("Do: %s", resp.Body)
	// Callers should close resp.Body
	// when done reading from it

	// Defer the closing of the body
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		if parse {

			// Fill the record with the data from the JSON
			var record ResultQuery

			// Use json.Decode for reading streams of JSON data

			if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
				log.Println(err)
			}

			fmt.Println("Phone No. = ", record.Ret)
			fmt.Println("Country   = ", record.Msg)
			//	fmt.Println("Location  = ", record.Data[0].Cjgzsj)

			return record, err

		} else {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			log.Printf("============== v%", bodyString)

			log.Printf("============== s%", bodyString)
			return bodyString, err
		}

	} else {
		return "", err

	}

}

var raw string = `

<h1>Welcome</h1>

`

func main_________() {
	con := "HOSTNAME=host;DATABASE=name;PORT=number;UID=username;PWD=password"
	db, err := sql.Open("go_ibm_db", con)
	if err != nil {

		fmt.Println(err)
	}
	db.Close()
}

func profileByUsername(ctx iris.Context) {
	// .Params are used to get dynamic path parameters.
	username := ctx.Params().Get("username")
	ctx.ViewData("Username", username)
	// renders "./views/users/profile.html"
	// with {{ .Username }} equals to the username dynamic path parameter.
	ctx.View("users/profile.html")
}
func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}

func DataMasking(s string) (r string) {

	s = strings.TrimSpace(s)

	rune_str := []rune(s)
	temp := SubString(s, 0, len(rune_str)/2)

	var buffer bytes.Buffer
	buffer.WriteString(temp)
	for i := 0; i < len(rune_str)/2; i++ {
		buffer.WriteString("*")
	}
	return buffer.String()

}
func main() {

	app := iris.Default()
	app.RegisterView(iris.HTML("E:/go/src/test/src/cmd/views", ".html").Reload(true))
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("E:/go/src/test/src/cmd/static"))))
	//	app.StaticEmbedded("/static", "./static", iris.Application.Asset, AssetNames)
	app.StaticWeb("/static", "E:/go/src/test/src/cmd/static")
	//app.Favicon("./web/ico/one.ico")
	app.HandleMany("GET", "/ /OpenAPI调用说明", func(ctx iris.Context) {

		User := make(map[string]string)
		User["Firstname"] = "https://graph.qq.com/user/get_simple_userinfo?access_token=1234ABD1234ABD&amp;oauth_consumer_key=12345&amp;openid=B08D412EEC4000FFC37CAABBDC1234CC&amp;format=json"
		ctx.ViewData("User", User)
		isHTTPS := ctx.Request().TLS != nil
		u := "/api/get_simple_userinfo?name=612724198409210339&access_token=1234ABD1234ABD&oauth_consumer_key=12345&openid=B08D412EEC4000FFC37CAABBDC1234CC&format=json"
		if isHTTPS {
			url := "https://" + ctx.Request().Host + u
			ctx.ViewData("Url", url)
		} else {
			url := "http://" + ctx.Request().Host + u
			ctx.ViewData("Url", url)
		}

		ctx.View("users/login.html")

	})
	app.Handle("GET", "/wiki/获取用户OpenID", func(ctx iris.Context) {
		User := make(map[string]string)

		fmt.Println("URL URL " + ctx.Request().Host)
		User["Firstname"] = "https://graph.qq.com/user/get_simple_userinfo?access_token=1234ABD1234ABD&amp;oauth_consumer_key=12345&amp;openid=B08D412EEC4000FFC37CAABBDC1234CC&amp;format=json"
		ctx.ViewData("User", User)

		url := ctx.Request().Host + "api/get_simple_userinfo?name=612724198409210339&access_token=1234ABD1234ABD&oauth_consumer_key=12345&openid=B08D412EEC4000FFC37CAABBDC1234CC&format=json"
		ctx.ViewData("Url", url)

		ctx.View("users/获取用户OpenID.html")
	})
	app.Handle("GET", "/wiki/使用Implicit_Grant方式获取Access_Token", func(ctx iris.Context) {
		User := make(map[string]string)
		User["Firstname"] = "https://graph.qq.com/user/get_simple_userinfo?access_token=1234ABD1234ABD&amp;oauth_consumer_key=12345&amp;openid=B08D412EEC4000FFC37CAABBDC1234CC&amp;format=json"
		ctx.ViewData("User", User)
		ctx.View("users/使用Implicit_Grant方式获取Access_Token.html")
	})

	app.Get("/profile/{username:string}", profileByUsername)

	// This handler will match /user/john but will not match neither /user/ or /user.
	app.Get("/api/get_simple_userinfo", func(ctx iris.Context) {
		fmt.Println("URL URL " + ctx.Request().Host)

		name := ctx.URLParam("name")
		ctx.URLParam("access_token")
		ctx.URLParam("oauth_consumer_key")
		ctx.URLParam("format")

		/*	https://graph.qq.com/user/get_simple_userinfo?access_token=1234ABD1234ABD&amp;oauth_consumer_key=12345&amp;
			openid=B08D412EEC4000FFC37CAABBDC1234CC&amp;format=json
		*/
		//ctx.Writef("Hello %s", name)
		body, err := api(name, true) //"612724198409210339")

		if err == nil {
			if v2, ok := body.(string); ok {
				println(v2)
				ctx.WriteString(v2)

			} else if v3, ok2 := body.(ResultQuery); ok2 {

				if v3.Ret == "0" {
					v3.Data[0].Xmjp = DataMasking(v3.Data[0].Xmjp)
					v3.Data[0].Grzh = DataMasking(v3.Data[0].Grzh)
					v3.Data[0].Dwmc = DataMasking(v3.Data[0].Dwmc)
					v3.Data[0].Dwdz = DataMasking(v3.Data[0].Dwdz)
					v3.Data[0].Zzdh = DataMasking(v3.Data[0].Zzdh)
					v3.Data[0].Grgfdz = DataMasking(v3.Data[0].Grgfdz)
					v3.Data[0].Lxjcdwmc = ""

					v3.Data[0].Dwfdrdbr = DataMasking(v3.Data[0].Dwfdrdbr)

				}
				ctx.JSON(v3)

			}
		} else {
			ctx.JSON(ResultQuery{Ret: "5", Msg: "aaaa"})

		}

		//	ctx.HTML(raw+body)
		//	ctx.JSON(iris.Map{"message": "Hello iris web framework."})

	})

	// This handler will match /user/john but will not match neither /user/ or /user.
	app.Get("/api/{name}", func(ctx iris.Context) {

		/*	https://graph.qq.com/user/get_simple_userinfo?access_token=1234ABD1234ABD&amp;oauth_consumer_key=12345&amp;
			openid=B08D412EEC4000FFC37CAABBDC1234CC&amp;format=json
		*/
		name := ctx.Params().Get("name")
		//ctx.Writef("Hello %s", name)
		body, _ := api(name, false) //"612724198409210339")
		//	ctx.HTML(raw+body)
		//	ctx.JSON(iris.Map{"message": "Hello iris web framework."})
		ctx.WriteString(body.(string))
	})

	// Method:   GET
	// Resource: http://localhost:8080/
	app.Handle("GET", "/0000000000000===", func(ctx iris.Context) {
		ctx.HTML("Hello world!")
	})

	// Method:   GET
	// Resource: http://localhost:8080/
	app.Handle("GET", "/0000000000000", func(ctx iris.Context) {
		ctx.WriteString("Hello world!")
	})

	app.Get("/users/{id:uint64}", func(ctx iris.Context) {
		//id := ctx.Params().GetUint64Default("id", 0)
		// [...]
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello iris web framework."})
	})
	fmt.Printf("------app.Run(iris.Add--------------")

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"))
	fmt.Printf("--------------------")
}

func New_mqtt(appID string, status map[string]string, server string) {
	log.Printf("MAIN 主程序继续  New_mqtt")
	//	log.Printf("MAIN 主程序继续 New_mqtt", cf.AppID, status, cf.Server)

	deviceId := status["deviceEui"]
	//appId := appID
	log.Printf("9999999999999 %s-------%s", appID, deviceId)

	defer func() {
		fmt.Println("Mqqt 错误 找值，defer end...")
	}()
	defer func() {

		if r := recover(); r != nil {
			fmt.Printf("MQTT 捕获到的错误：%s\n", r)
		}
	}()

	var thingName string = "Led"
	var region string = "us-west-2"

	opts := MQTT.NewClientOptions().AddBroker(server)
	opts.SetClientID("mac-go")
	opts.SetUsername("11")
	opts.SetPassword(region)
	opts.SetUsername(thingName)

	//opts.SetDefaultPublishHandler(f)

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	//Client := c
	//
	r := strings.NewReplacer("<DevID>", deviceId, "<AppID>", appID)

	rt := r.Replace("DOWNLINK_Messages_t_down")
	/*	SubscribeHandler := nil
		log.Printf("%s", rt)
		if token := c.Subscribe(rt, 0, SubscribeHandler); token.Wait() &&
			token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)

		}*/

	//v3/app1/devices/dev1/down

	//v3/{application id}/devices/{device id}/down/ack

	rt = r.Replace("<AppID>/devices/<DevID>/events/activations/errors")
	log.Printf("%s", rt)
	//Downlink Messages
	if token := c.Subscribe(rt, 0, nil); token.Wait() &&
		token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	log.Printf(" 结束mqtt 初始化")

	time.Sleep(3 * time.Second)

} //

/* ```


go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo

E:\go\src\test\cmd>%GOPATH%\bin\rsrc.exe -manifest shell.manifest -o
rsrc.syso


GOOS="windows" GOARCH="386" go build -o app.exe

```
*/
func Decode(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}
func Gbk2Utf(s string) (string, error) {
	/* 由gbk ---> utf8  */
	bs := []byte(s)
	reader := transform.NewReader(bytes.NewReader(bs), simplifiedchinese.GBK.NewDecoder())
	res, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func Utf2Gbk(s string) (string, error) {
	/* utf8 ---> gbk  */
	bs := []byte(s)
	reader := transform.NewReader(bytes.NewReader(bs), simplifiedchinese.GBK.NewEncoder())
	res, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
func m() {

	//exec.Command("netsh", "interface", "ipv6", "set", "privacy", "state=disable").Run()
	//main___()
	/*	cmd := exec.Command("java") // or whatever the program is
		cmd.Dir = "C:/"         // or whatever directory it's in
		out, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("%s", out);
		}*/

	//cmd := exec.Command(`.\\echo.bat`, "&whoami", "a-z", "A-Z")
	cmd := exec.Command(`.\\echo.bat`, "&whoami")

	cmd.Dir = "C:/"
	cmd.Stdin = os.Stdin
	//cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	out, err := cmd.Output() //　//运行命令并返回其标准输出

	//out, err := cmd.CombinedOutput()   //并返回标准输出和标准错误
	//err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s", out)
	}

	_, filename, _, ok := runtime.Caller(1)
	var cwdPath string
	if ok {
		cwdPath = path.Join(path.Dir(filename), "") // the the main function file directory
	} else {
		cwdPath = "./"
	}
	fmt.Println("cwd path...", cwdPath)

	runSsh()
	/*	err = cmd.Wait()
		if err != nil {
			log.Printf("Command finished with error: %v", err)
		}*/
}

func SSHConnect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: auth,
		// Timeout:             30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}

type PowerShell struct {
	powerShell string
}

func New() *PowerShell {
	ps, _ := exec.LookPath("powershell.exe")
	return &PowerShell{
		powerShell: ps,
	}
}

func (p *PowerShell) Execute(args ...string) (stdOut string, stdErr string, err error) {
	//args = append([]string{"-NoProfile", "-NonInteractive"}, args...)
	args = append([]string{}, args...)

	cmd := exec.Command(p.powerShell, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	stdt, err := Gbk2Utf(stdout.String())
	if err != nil {
		fmt.Printf(" 转换失败")
	}
	stdr, err := Gbk2Utf(stderr.String())
	if err != nil {
		fmt.Printf(" 错误输出 转换失败")

	}

	stdOut, stdErr = stdt, stdr
	return
}

func ain___() {
	posh := New()
	//stdout, stderr, err := posh.Execute("$OutputEncoding = [Console]::OutputEncoding; (Get-VMSwitch).Name")
	stdout, stderr, err := posh.Execute("mkdir 'c:\\windows\\我爱你'")

	fmt.Println(stdout)
	fmt.Println(stderr)

	if err != nil {
		fmt.Println(err)
	}
}

func runSsh() {

	var stdOut, stdErr bytes.Buffer

	session, err := SSHConnect("root", "silence@110!", "192.168.10.90", 22)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.Stdout = &stdOut
	session.Stderr = &stdErr
	session.Run("ls")
	//	session.Run("if [ -d liujx/project ]; then echo 0; else echo 1; fi")
	fmt.Printf("%s, %s\n", stdOut.String(), stdErr.String())
	ret, err := strconv.Atoi(strings.Replace(stdOut.String(), "\n", "", -1))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d, %s\n", ret, stdErr.String())

}
