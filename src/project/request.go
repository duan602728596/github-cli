package project

import (
  "github-cli/src/getConfig"
  "io/ioutil"
  "net/http"
  "strings"
)

/**
 * 发送http请求
 * @param { getConfig.Config } config: 配置
 * @param { string } query: GraphQL查询条件
 */
func request(config getConfig.Config, query string) []byte {
  // 请求body
  value := `{ "query": "` + query + `" }`
  body := ioutil.NopCloser(strings.NewReader(value))

  // 发送请求
  client := &http.Client {}
  req, _ := http.NewRequest("POST", "https://api.github.com/graphql", body)

  req.Header.Add("Authorization", "bearer " + config.Token)
  req.Header.Add("Content-Type", "application/json")

  resp, _ := client.Do(req)

  defer resp.Body.Close()

  resBody, _ := ioutil.ReadAll(resp.Body)

  return resBody
}