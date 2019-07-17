package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"

	"golang.org/x/crypto/ssh"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"path"
	"runtime"
	"strconv"
	"strings"
	//"github.com/akavel/rsrc"
)

/* ```


%GOPATH%\bin\rsrc.exe -manifest shell.manifest -o rsrc.syso
%GOPATH%\bin\rsrc.exe -manifest E:\go\src\test\shell.manifest -o E:\go\src\test\rsrc.syso



E:\go\src\test>%GOPATH%\bin\rsrc.exe -manifest E:\go\src\test\shell.manifest -o
rsrc.syso



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
func main() {
	//exec.Command("netsh", "interface", "ipv6", "set", "privacy", "state=disable").Run()
	main___()
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

func main___() {
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
