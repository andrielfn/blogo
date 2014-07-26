package main

import (
  "bytes"
  "github.com/russross/blackfriday"
  "html/template"
  "io/ioutil"
  "log"
  "net/http"
  "path"
  "regexp"
  "time"
)

type Header struct {
  Title       string
  Description string
  Date        time.Time
  Tags        string
}

type Post struct {
  Header Header
  Body   template.HTML
}

func main() {
  fs := http.FileServer(http.Dir("public"))
  http.Handle("/public/", http.StripPrefix("/public/", fs))

  http.HandleFunc("/post/", postHandler)

  log.Println("Listening...")
  http.ListenAndServe(":3001", nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/post/"):]

  post, err := loadPost(title)
  if err != nil {
    http.NotFound(w, r)
    return
  }

  template, err := template.ParseFiles("layouts/post.html")

  if err != nil {
    http.NotFound(w, r)
    return
  }

  template.Execute(w, post)
}

func loadPost(title string) (*Post, error) {
  filename := getPostFilenameFor(title)
  post, err := parsePostFor(filename)

  return post, err
}

func getPostFilenameFor(s string) string {
  return path.Join("posts", s+".md")
}

func parsePostFor(filename string) (*Post, error) {
  fileContent, err := ioutil.ReadFile(filename)

  if err != nil {
    return nil, err
  }

  header, body := parseFile(fileContent)

  return &Post{Header: header, Body: template.HTML(body)}, nil
}

func parseFile(content []byte) (Header, []byte) {
  splited := bytes.Split(content, []byte("---"))
  header := parseHeader(splited[1])

  return header, renderMarkdown(splited[2])
}

func parseHeader(content []byte) Header {
  titleRegex := regexp.MustCompile("title:\\s(.+)")
  title := titleRegex.FindStringSubmatch(string(content))[1]

  descriptionRegex := regexp.MustCompile("description:\\s(.+)")
  description := descriptionRegex.FindStringSubmatch(string(content))[1]

  dateRegex := regexp.MustCompile("date:\\s(.+)")
  date := dateRegex.FindStringSubmatch(string(content))[1]
  parsedDate, _ := time.Parse("2006-01-02", date)

  tagsRegex := regexp.MustCompile("tags:\\s(.+)")
  tags := tagsRegex.FindStringSubmatch(string(content))[1]

  return Header{
    Title:       title,
    Description: description,
    Date:        parsedDate,
    Tags:        tags}
}

func renderMarkdown(input []byte) []byte {
  // set up the HTML renderer
  htmlFlags := 0
  htmlFlags |= blackfriday.HTML_GITHUB_BLOCKCODE
  htmlFlags |= blackfriday.HTML_USE_XHTML
  htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
  htmlFlags |= blackfriday.HTML_SMARTYPANTS_FRACTIONS
  htmlFlags |= blackfriday.HTML_SMARTYPANTS_LATEX_DASHES
  // htmlFlags |= blackfriday.HTML_SANITIZE_OUTPUT
  renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")

  // set up the parser
  extensions := 0
  extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
  extensions |= blackfriday.EXTENSION_TABLES
  extensions |= blackfriday.EXTENSION_FENCED_CODE
  extensions |= blackfriday.EXTENSION_AUTOLINK
  extensions |= blackfriday.EXTENSION_STRIKETHROUGH
  extensions |= blackfriday.EXTENSION_SPACE_HEADERS
  extensions |= blackfriday.EXTENSION_HEADER_IDS

  return blackfriday.Markdown(input, renderer, extensions)
}
