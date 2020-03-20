package getConfig

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
)

type Config struct {
  Username string `json:"username"`
  Token  string `json:"token"`
}

/**
 * @param { string } dir: 文件（软件）所在的目录
 */
func GetConfig(dir string) Config {
  // 读取文件
  jsonFile, err := os.Open(dir + "/config.json")

  if err != nil {
    fmt.Println(err)
  }

  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)

  // 解析json
  var config Config

  json.Unmarshal([]byte(byteValue), &config)

  return config
}