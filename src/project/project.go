package project

import (
  "bytes"
  "encoding/json"
  "fmt"
  "github-cli/src/getConfig"
  "html/template"
  "strings"
)

type Node struct {
  Id string `json: "id"`
  Url string `json: "url"`
  Name string `json: "name"`
  NameWithOwner string `json: "nameWithOwner"`
  Description string `json: "description"`
}

type PageInfo struct {
  EndCursor string `json: "endCursor"`
  HasNextPage bool `json: "hasNextPage"`
  HasPreviousPage bool `json: "hasPreviousPage"`
  StartCursor string `json: "startCursor"`
}

type Repositories struct {
  Nodes []Node `json: "nodes"`
  PageInfo PageInfo `json: "pageInfo"`
  TotalCount int `json: "totalCount"`
}

type ResData struct {
  Data struct {
    User struct {
      StarredRepositories Repositories `json:"starredRepositories"`
      Repositories Repositories `json:"repositories"`
    } `json:"user"`
  } `json:"data"`
}

type Tpl struct {
  Username string
  QueryType string
  Nodes []Node
}

/* 查询条件 */
func starsQuery (config getConfig.Config, queryType, before string) string {
  // login查询条件
  search := ""

  if before != "" {
    search += ", before:\\\"" + before + "\\\""
  }

  // 查询类型
  qType := queryType

  if qType == "stars" {
    qType = "starredRepositories"
  }

  query := `{
  user(login: \"` + config.Username + `\") {
    ` + qType + `(last: 100` + search + `) {
      nodes {
        id
        url
        name
        nameWithOwner
        description
      }
      pageInfo {
        endCursor
        hasNextPage
        hasPreviousPage
        startCursor
      }
      totalCount
    }
  }
}`

  return strings.Replace(query, "\n", "\\n", -1)
}

/**
 * 查询项目
 * @param { getConfig.Config } config: 配置
 * @param { string } dir: 文件（软件）所在的目录
 * @param { 'stars' | 'repositories' } queryType: 查询类型
 * @param { '' | string } port: 查询端口
 * @param { boolean } openBrowser: 是否自动打开浏览器
 */
func Project(config getConfig.Config, dir string, queryType string, port string, openBrowser bool) {
  hasBefore := true
  beforeId := ""
  nodes := []Node {}

  for hasBefore {
    // 查数据
    res := request(config, starsQuery(config, queryType, beforeId))

    // 解析json
    var resData ResData

    json.Unmarshal([]byte(res), &resData)

    // 合并数组
    if queryType == "stars" {
      nodes = append(nodes, reverse(resData.Data.User.StarredRepositories.Nodes)...)
    } else {
      nodes = append(nodes, reverse(resData.Data.User.Repositories.Nodes)...)
    }

    // 判断是否还有数据
    if resData.Data.User.StarredRepositories.PageInfo.HasPreviousPage == false {
      hasBefore = false
    } else {
      beforeId = resData.Data.User.StarredRepositories.PageInfo.StartCursor
    }
  }

  // 拼接html
  tpl := template.New("index.html")
  tpl, _ = tpl.ParseFiles(dir + "/template/index.html")

  var resHtml bytes.Buffer

  err := tpl.Execute(&resHtml, Tpl {
    Username: config.Username,
    QueryType: queryType,
    Nodes: nodes,
  })

  if err != nil {
    fmt.Println(err)
  }

  result := resHtml.String()

  // 启动服务
  Server(result, port, openBrowser)
}