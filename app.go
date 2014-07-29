package main

import (
  "html/template"
  "log"
  "net/http"
)

type Blog struct {
  Articles map[string]*Article
}

func (b *Blog) homeHandler(w http.ResponseWriter, r *http.Request) {
  template, _ := template.ParseFiles("layouts/home.html")

  template.Execute(w, b.Articles)
}

func (b *Blog) getArticle(slug string) *Article {
  if article := b.Articles[slug]; article != nil {
    return article
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
  http.ListenAndServe(":3001", nil)
}
