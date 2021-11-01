package project

import (
  "fmt"
  "io"
  "net"
  "net/http"
  "os/exec"
  "runtime"
  "strconv"
)

/* 打开浏览器 */
func openUrlInBrowser(url string) {
  switch runtime.GOOS {
    case "darwin":
      exec.Command("open", url).Start()
      break

    case "windows":
      exec.Command("cmd", "/c", "start", url).Start()
      break

    case "linux":
      exec.Command("xdg-open", url).Start()
      break
  }
}

/* 获取本机可用端口 */
func getFreePort() (int, error) {
  addr, err := net.ResolveTCPAddr("tcp", "localhost:0")

  if err != nil {
    return 0, err
  }

  l, err := net.ListenTCP("tcp", addr)

  if err != nil {
    return 0, err
  }

  defer l.Close()

  return l.Addr().(*net.TCPAddr).Port, nil
}

/* 启动http服务 */
func Server(htmlResult string, port string, openBrowser bool) {
  // 获取端口
  usePort := port

  if usePort == "" {
    freePort, err := getFreePort()

    if err != nil {
      fmt.Println(err)
    }

    usePort = strconv.Itoa(freePort)
  }

  // 启动地址
  url := "http://localhost:" + usePort + "/"

  fmt.Println("启动服务: " + url)

  if openBrowser {
    openUrlInBrowser(url)
  }

  // 启动服务
  http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request)  {
    io.WriteString(w, htmlResult)
  })

  err := http.ListenAndServe(":" + usePort, nil)

  if err != nil {
    fmt.Println(err)
  }
}