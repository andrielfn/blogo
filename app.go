package main

import (
  "html/template"
  "log"
  "net/http"
  "path"
)

type Blog struct {
  Articles ArticleList
}

func renderTemplate(name string) *template.Template {
  layout := path.Join("templates", "layout.html")
  partial := path.Join("templates", name+".html")
  return template.Must(template.ParseFiles(layout, partial))
}

func (b *Blog) homeHandler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }

  renderTemplate("home").ExecuteTemplate(w, "layout", b.Articles)
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

  // var home = template.Must(template.ParseFiles("templates/layout.html", "templates/article.html"))
  renderTemplate("article").ExecuteTemplate(w, "layout", article)
}

func (b *Blog) aboutHandler(w http.ResponseWriter, r *http.Request) {
  renderTemplate("about").ExecuteTemplate(w, "layout", nil)
}

func main() {
  blog := Blog{Articles: LoadArticles()}

  fs := http.FileServer(http.Dir("public"))
  http.Handle("/public/", http.StripPrefix("/public/", fs))

  http.HandleFunc("/", blog.homeHandler)
  http.HandleFunc("/articles/", blog.articleHandler)
  http.HandleFunc("/about", blog.aboutHandler)

  log.Println("Listening...")
  http.ListenAndServe(":3000", Log(http.DefaultServeMux))
}
