package main

import (
  "html/template"
  "log"
  "net/http"
)

type Blog struct {
  Articles ArticleList
}

func (b *Blog) homeHandler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }

  template, _ := template.ParseFiles("layouts/home.html")

  template.Execute(w, b.Articles)
}

func (b *Blog) getArticle(slug string) *Article {
  for _, article := range b.Articles {
    if article.Metadata.Slug == slug {
      return article
    }
  }

  return nil
}

func (b *Blog) articleHandler(w http.ResponseWriter, r *http.Request) {
  slug := r.URL.Path[len("/articles/"):]

  article := b.getArticle(slug)

  if article == nil {
    http.NotFound(w, r)
    return
  }

  template, _ := template.ParseFiles("layouts/article.html")

  template.Execute(w, article)
}

func main() {
  blog := Blog{Articles: LoadArticles()}

  fs := http.FileServer(http.Dir("public"))
  http.Handle("/public/", http.StripPrefix("/public/", fs))

  http.HandleFunc("/", blog.homeHandler)
  http.HandleFunc("/articles/", blog.articleHandler)

  log.Println("Listening...")
  http.ListenAndServe(":3000", Log(http.DefaultServeMux))
}
