package main 
 
import ( 
    "net/http" 
    "io/ioutil" 
    "regexp" 
    "text/template" 
) 
 
type Page struct { 
    Title string 
    Body  []byte 
} 

type error interface {
	Error() string
}
 
func (p *Page) save() error { 
    filename := p.Title + ".txt" 
    return ioutil.WriteFile(filename, p.Body, 0600) 
} 
 
func loadPage(title string) (*Page, error) { 
    filename := title + ".txt" 
    body, err := ioutil.ReadFile(filename) 
    if err != nil { 
        return nil, err 
    } 
    return &Page{Title: title, Body: body}, nil 
} 
 
func viewHandler(w http.ResponseWriter, r *http.Request, title string) { 
    p, err := loadPage(title) 
    if err != nil { 
        http.Redirect(w, r, "/edit/"+title, http.StatusFound) 
        return 
    } 
    renderTemplate(w, "view", p) 
} 
 
func editHandler(w http.ResponseWriter, r *http.Request, title string) { 
    p, err := loadPage(title) 
    if err != nil { 
        p = &Page{Title: title} 
    } 
    renderTemplate(w, "edit", p) 
} 
 
func saveHandler(w http.ResponseWriter, r *http.Request, title string) { 
    body := r.FormValue("body") 
    p := &Page{Title: title, Body: []byte(body)} 
    err := p.save() 
    if err != nil { 
        http.Error(w, err.Error(), http.StatusInternalServerError) 
        return 
    } 
    http.Redirect(w, r, "/view/"+title, http.StatusFound) 
} 

func menuHandler(w http.ResponseWriter, r*http.Request, title string) {
	p, err := loadPage(title)

	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "menu", p)
}

func foodHandler(w http.ResponseWriter, r*http.Request, title string) {
	//body := r.FormValue("food")
	//fmt.Printf("%s", body)
}
 
var templates = make(map[string]*template.Template) 
 
func init() { 
    for _, tmpl := range []string{"edit", "view", "menu"} { 
    	filename := tmpl + ".html" 
    	body, err := ioutil.ReadFile(filename) 
    	if err != nil { 
    	} 
	t := template.New(tmpl)
	t.Parse(string(body))
        //t := template.Must(template.Parse(tmpl + ".html")) 
        templates[tmpl] = t 
    } 
} 
 
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) { 
    err := templates[tmpl].Execute(w, p) 
    if err != nil { 
        http.Error(w, err.Error(), http.StatusInternalServerError) 
    } 
} 
 
const lenPath = len("/view/") 
 
var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$") 
 
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc { 
    return func(w http.ResponseWriter, r *http.Request) { 
        title := r.URL.Path[lenPath:] 
        if !titleValidator.MatchString(title) { 
            http.NotFound(w, r) 
            return 
        } 
        fn(w, r, title) 
    } 
} 

func chihuoHandler(w http.ResponseWriter, r *http.Request, title string) {
	w.Header().Add("Set-Cookie", "123")
}
 
func main() { 
    http.HandleFunc("/view/", makeHandler(viewHandler)) 
    http.HandleFunc("/edit/", makeHandler(editHandler)) 
    http.HandleFunc("/save/", makeHandler(saveHandler)) 
    http.HandleFunc("/menu/", makeHandler(menuHandler)) 
    http.HandleFunc("/food/", makeHandler(foodHandler)) 
    http.ListenAndServe(":7777", nil) 
} 
