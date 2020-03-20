package main

import (
  "fmt"
  "github.com/AlecAivazis/survey/v2"
  "github-cli/src/getConfig"
  "github-cli/src/project"
  "os"
)

func main()  {
  fmt.Println("repositories(r/repo): 查询自己的项目\nstars      (s/stars): 查询关注的项目") // 输出说明

  dir, _ := os.Getwd()               // 文件目录
  config := getConfig.GetConfig(dir) // 获取配置

  queryType := "" // 查询参数
  port := ""      // 端口号
  openBrowser := true // 打开浏览器

  survey.AskOne(&survey.Input { Message: "你要查询什么: ", }, &queryType)
  survey.AskOne(&survey.Input { Message: "端口号（随机）: ", }, &port)
  survey.AskOne(&survey.Confirm{ Message: "自动打开浏览器吗？", Default: true }, &openBrowser)

  if queryType == "s" {
    queryType = "stars"
  } else if queryType == "r" || queryType == "repo" {
    queryType = "repositories"
  }

  switch queryType {
    case "stars":        // 查询stars
      project.Project(config, dir, queryType, port, openBrowser)
      break

    case "repositories": // 查询repositories
      project.Project(config, dir, queryType, port, openBrowser)
      break
  }
}