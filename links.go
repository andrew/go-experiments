// spider a site given a url

package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "code.google.com/p/go.net/html"
  "strings"
)

func main() {
  urls := parse(download("http://google.com/"))
  for _, v := range urls {
    parse(download(v))
  }
}

func download(url string) string{
  res, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  response, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
    log.Fatal(err)
  }

  s := string(response[:])
  return s
}

func parse(body string) []string{
    s := []string{}
    doc, err := html.Parse(strings.NewReader(body))
    if err != nil {
        log.Fatal(err)
    }
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            for _, a := range n.Attr {
                if a.Key == "href" {
                    fmt.Println(a.Val)
                    s = append(s, a.Val)
                    break
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }

    }
    f(doc)
    return s
}
