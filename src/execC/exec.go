package execC

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"

	//    "strconv"
	"os/exec"
	"os/user"
	"time"
)

func Chrome_on() {

	f, err := exec.LookPath("chromium-browser")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(f)

	cmd := exec.Command("/bin/sh", "-c", "export DISPLAY=:0.0; echo $DISPLAY; chromium-browser --incognito http://localhost:10080/report/index/STORAGE00000001")
	//cmd := exec.Command("/bin/sh","-c","chromium-browser --incognito --kiosk  http://localhost:10080/report/index/STORAGE00000001 ")
	//cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	user, err := user.Lookup("pi")
	if err == nil {
		log.Printf("uid=%s,gid=%s", user.Uid, user.Gid)

		//     uid, _ := strconv.Atoi(user.Uid)
		//       gid, _ := strconv.Atoi(user.Gid)

		//       cmd.SysProcAttr = &syscall.SysProcAttr{}
		//     cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
	}
	tm := time.AfterFunc(10*time.Second, func() {
		//	syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	})

	log.Printf("timer=%v", tm)

	//  go func(){
	//    time.Sleep(1*time.Second)
	//  cmd.Process.Kill()
	//}()

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}
func Chrome_off() {

	f, err := exec.LookPath("chromium-browser")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(f)

	cmd := exec.Command("/bin/sh", "-c", "export DISPLAY=:0.0; echo $DISPLAY; chromium-browser --incognito http://localhost:10080/report/index/STORAGE00000001")
	//cmd := exec.Command("/bin/sh","-c","chromium-browser --incognito --kiosk  http://localhost:10080/report/index/STORAGE00000001 ")
	//	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	user, err := user.Lookup("pi")
	if err == nil {
		log.Printf("uid=%s,gid=%s", user.Uid, user.Gid)

		//     uid, _ := strconv.Atoi(user.Uid)
		//       gid, _ := strconv.Atoi(user.Gid)

		//       cmd.SysProcAttr = &syscall.SysProcAttr{}
		//     cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
	}
	tm := time.AfterFunc(10*time.Second, func() {
		//	syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	})

	log.Printf("timer=%v", tm)

	//  go func(){
	//    time.Sleep(1*time.Second)
	//  cmd.Process.Kill()
	//}()

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	f, err := exec.LookPath("chromium-browser")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(f)

	cmd := exec.Command("/bin/sh", "-c", "export DISPLAY=:0.0; echo $DISPLAY; chromium-browser --incognito http://localhost:10080/report/index/STORAGE00000001")
	//cmd := exec.Command("/bin/sh","-c","chromium-browser --incognito --kiosk  http://localhost:10080/report/index/STORAGE00000001 ")
	//	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	user, err := user.Lookup("pi")
	if err == nil {
		log.Printf("uid=%s,gid=%s", user.Uid, user.Gid)

		//     uid, _ := strconv.Atoi(user.Uid)
		//       gid, _ := strconv.Atoi(user.Gid)

		//       cmd.SysProcAttr = &syscall.SysProcAttr{}
		//     cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
	}
	tm := time.AfterFunc(10*time.Second, func() {
		//	syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	})

	log.Printf("timer=%v", tm)

	//  go func(){
	//    time.Sleep(1*time.Second)
	//  cmd.Process.Kill()
	//}()

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}

func main___() {
	cmdStr := "sudo docker run -v ~/exp/a.out:/a.out ubuntu:14.04 /a.out -m 10m"
	out, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()
	fmt.Printf("%s", out)
}

func DockerCompose(fileName string) {

	f, err := exec.LookPath("docker-compose")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(f)

	cmd := exec.Command("/bin/sh", "-c", "docker-compose up")
	//cmd := exec.Command("/bin/sh","-c","chromium-browser --incognito --kiosk  http://localhost:10080/report/index/STORAGE00000001 ")
	//	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	user, err := user.Lookup("pi")
	if err == nil {
		log.Printf("uid=%s,gid=%s", user.Uid, user.Gid)

		//     uid, _ := strconv.Atoi(user.Uid)
		//       gid, _ := strconv.Atoi(user.Gid)

		//       cmd.SysProcAttr = &syscall.SysProcAttr{}
		//     cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
	}
	tm := time.AfterFunc(10*time.Second, func() {
		//	syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	})

	log.Printf("timer=%v", tm)

	//  go func(){
	//    time.Sleep(1*time.Second)
	//  cmd.Process.Kill()
	//}()

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}

// Prepare some data to insert into the template.
type Recipient struct {
	Name, Gift, Url string
	Attended        bool
}

func Service(fileName string) {

	var script = `
	cd /tmp
	sudo useradd echoservice -s /sbin/nologin -M
	wget {{.Url}}
	sudo mv {{.Name}} /lib/systemd/system/.
	sudo chmod 755 /lib/systemd/system/{{.Name}}
	`
	var recipients = []Recipient{
		{"simple-api.service",
			"bone china tea set",
			"https://raw.githubusercontent.com/1026957152/test/master/src/simple-api.service",
			true},

		{"Uncle John", "moleskin pants", "", false},
		{"Cousin Rodney", "", "", false},
	}

	// Create a new template and parse the letter into it.
	ttemplate := template.Must(template.New("letter").Parse(script))

	var tpl bytes.Buffer
	err := ttemplate.Execute(&tpl, recipients[0])

	cmd := exec.Command("/bin/sh", "-c", tpl.String())
	//cmd := exec.Command("/bin/sh","-c","chromium-browser --incognito --kiosk  http://localhost:10080/report/index/STORAGE00000001 ")
	//	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	user, err := user.Lookup("pi")
	if err == nil {
		log.Printf("uid=%s,gid=%s", user.Uid, user.Gid)

		//     uid, _ := strconv.Atoi(user.Uid)
		//       gid, _ := strconv.Atoi(user.Gid)

		//       cmd.SysProcAttr = &syscall.SysProcAttr{}
		//     cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
	}
	tm := time.AfterFunc(10*time.Second, func() {
		//	syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	})

	log.Printf("timer=%v", tm)

	//  go func(){
	//    time.Sleep(1*time.Second)
	//  cmd.Process.Kill()
	//}()

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	systemctl("")
}

func systemctl(fileName string) {

	var systemctl = `
sudo systemctl enable {{.Name}}
sudo systemctl start {{.Gift}}
`
	/*	var systemctl = `
	sudo systemctl enable {{.Name}}
	sudo systemctl start {{.Gift}}
	sudo journalctl -f -u {{.Gift}}
	`*/
	var recipients = []Recipient{
		{"simple-api.service",
			"simple-api",
			"https://raw.githubusercontent.com/1026957152/test/master/src/simple-api.service",
			true},
		{"Uncle John", "moleskin pants", "", false},
		{"Cousin Rodney", "", "", false},
	}

	// Create a new template and parse the letter into it.
	ttemplate := template.Must(template.New("letter").Parse(systemctl))
	var tpl bytes.Buffer
	err := ttemplate.Execute(&tpl, recipients[0])

	cmd := exec.Command("/bin/sh", "-c", tpl.String())
	//cmd := exec.Command("/bin/sh","-c","chromium-browser --incognito --kiosk  http://localhost:10080/report/index/STORAGE00000001 ")
	//	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	user, err := user.Lookup("pi")
	if err == nil {
		log.Printf("uid=%s,gid=%s", user.Uid, user.Gid)

		//     uid, _ := strconv.Atoi(user.Uid)
		//       gid, _ := strconv.Atoi(user.Gid)
		//       cmd.SysProcAttr = &syscall.SysProcAttr{}
		//     cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
	}

	//  go func(){
	//    time.Sleep(1*time.Second)
	//  cmd.Process.Kill()
	//}()

	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("执行结束啊啊啊")

}
