package vodka

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

const (
	get  = "GET"
	put  = "PUT"
	post = "POST"
)

var (
	routeTable map[string]*Route = make(map[string]*Route)
)

// HandlerFunc for handler functions.
type HandlerFunc func(c *Context)

// Route definition.
type Route struct {
	path    string
	param   map[int]string
	method  string
	handler HandlerFunc
}

// Engine for the framework
//Will add more functionality to this structure as this code grows.
type Engine struct {
}

//Default vodka engine.
func Default() *Engine {
	engine := Engine{}
	return &engine
}

// Run the api engine
func (e *Engine) Run(port int64) {
	http.HandleFunc("/", handleRequest)
	startHTTPServer(port)
}

func startHTTPServer(port int64) {

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}

// GET adds HTTP get routes.
func (e *Engine) GET(path string, handler HandlerFunc) {

	route := Route{
		method:  http.MethodGet,
		handler: handler,
		param:   map[int]string{},
		path:    path,
	}
	sPath := strings.Split(path, "/")
	re := regexp.MustCompile(`\{\w*\}`)
	params := re.FindAllString(path, -1)

	for _, p := range params {
		route.path = strings.Replace(path, p, `\w*`, -1)
	}
	//Revisit this logic. I have done string manipulation, just to unblock myself.
	for i, p := range strings.Split(route.path, "/") {
		if p == `\w*` {
			route.param[i] = strings.Trim(strings.Trim(sPath[i], "{"), "}")
		}
	}

	routeTable[route.path] = &route
}

// POST route mapping
func (e *Engine) POST(path string, handler HandlerFunc) {
	route := Route{
		path:    path,
		method:  http.MethodPost,
		handler: handler,
	}
	routeTable[path] = &route
}

// Just implamented HTTP GET and POST methods.
func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetRequest(w, r)
		return
	case "POST":
		handlePost(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("HTTP methos not allowed."))
		return
	}
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	context := Context{
		Writer:  w,
		Request: r,
	}
	route := getRoute(r)
	if route == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	sPath := strings.Split(r.RequestURI, "/")
	pathParams := make(map[string]string)
	for i, p := range route.param {
		pathParams[p] = sPath[i]
	}
	context.PathParams = pathParams
	route.handler(&context)
}

func getRoute(r *http.Request) *Route {
	for _, p := range routeTable {
		re := regexp.MustCompile("(?i)" + p.path)
		if re.MatchString(r.RequestURI) == true {
			route := routeTable[p.path]
			if route.method == r.Method {
				return route
			}
			return nil
		}
	}
	return nil
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	//TODO: In this excercise this was not needed, hence I have not implemented it.
	//Partially implimenting this to return 404 id url is not found.
	route := getRoute(r)
	if route == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
