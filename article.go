package main

import (
  "bytes"
  "github.com/russross/blackfriday"
  "html/template"
  "io/ioutil"
  "path/filepath"
  "regexp"
  "sort"
  "time"
)

type Article struct {
  Metadata *Metadata
  Content  template.HTML
}

type Metadata struct {
  Title       string
  Slug        string
  Description string
  Date        time.Time
  Tags        string
  HeroImage   string
}

type ArticleList []*Article

func (a ArticleList) Len() int           { return len(a) }
func (a ArticleList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ArticleList) Less(i, j int) bool { return a[i].Metadata.Date.After(a[j].Metadata.Date) }

func LoadArticles() ArticleList {
  files, _ := filepath.Glob("articles/*.md")

  articles := make(ArticleList, len(files))

  for i, filename := range files {
    article := parseArticle(filename)
    articles[i] = article
  }

  sort.Sort(articles)

  return articles
}

func parseArticle(filename string) *Article {
  fileContent, _ := ioutil.ReadFile(filename)

  metadata, content := parseFileContent(fileContent)

  return &Article{Metadata: metadata, Content: content}
}

func parseFileContent(content []byte) (*Metadata, template.HTML) {
  splited := bytes.Split(content, []byte("---"))
  return parseMetadata(splited[1]), parseBody(splited[2])
}

func parseMetadata(metadata []byte) *Metadata {
  titleRegex := regexp.MustCompile("title:\\s(.+)")
  title := titleRegex.FindStringSubmatch(string(metadata))[1]

  descriptionRegex := regexp.MustCompile("description:\\s(.+)")
  description := descriptionRegex.FindStringSubmatch(string(metadata))[1]

  dateRegex := regexp.MustCompile("date:\\s(.+)")
  date := dateRegex.FindStringSubmatch(string(metadata))[1]
  parsedDate, _ := time.Parse("2006-01-02", date)

  tagsRegex := regexp.MustCompile("tags:\\s(.+)")
  tags := tagsRegex.FindStringSubmatch(string(metadata))[1]

  slugRegex := regexp.MustCompile("slug:\\s(.+)")
  slug := slugRegex.FindStringSubmatch(string(metadata))[1]

  return &Metadata{
    Title:       title,
    Slug:        slug,
    Description: description,
    Date:        parsedDate,
    Tags:        tags}
}

func parseBody(content []byte) template.HTML {
  parsedMarkdown := blackfriday.MarkdownCommon(content)
  return template.HTML(parsedMarkdown)
}
