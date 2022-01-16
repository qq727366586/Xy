/**
 *@Author luojunying
 *@Date 2022-01-16 19:32
 */
package Xy

import (
	"net/http"
	"strings"
)

type Core struct {
	router map[string]*Tree
}

func initRouter() map[string]*Tree {
	router := map[string]*Tree{}
	router[http.MethodGet] = NewTree()
	router[http.MethodPost] = NewTree()
	router[http.MethodPut] = NewTree()
	router[http.MethodPatch] = NewTree()
	router[http.MethodHead] = NewTree()
	router[http.MethodDelete] = NewTree()
	router[http.MethodConnect] = NewTree()
	router[http.MethodTrace] = NewTree()
	router[http.MethodOptions] = NewTree()
	return router
}

func NewCore() *Core {
	return &Core{
		router: initRouter(),
	}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[http.MethodGet].AddRouter(url, handler)
}

func (c *Core) Post(url string, handler ControllerHandler) {
	c.router[http.MethodPost].AddRouter(url, handler)
}

func (c *Core) Put(url string, handler ControllerHandler) {
	c.router[http.MethodPut].AddRouter(url, handler)
}

func (c *Core) Patch(url string, handler ControllerHandler) {
	c.router[http.MethodPatch].AddRouter(url, handler)
}

func (c *Core) Head(url string, handler ControllerHandler) {
	c.router[http.MethodHead].AddRouter(url, handler)
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	c.router[http.MethodDelete].AddRouter(url, handler)
}

func (c *Core) Connect(url string, handler ControllerHandler) {
	c.router[http.MethodConnect].AddRouter(url, handler)
}

func (c *Core) Trace(url string, handler ControllerHandler) {
	c.router[http.MethodTrace].AddRouter(url, handler)
}

func (c *Core) Options(url string, handler ControllerHandler) {
	c.router[http.MethodOptions].AddRouter(url, handler)
}

func (c *Core) Any(url string, handler ControllerHandler) {
	c.router[http.MethodGet].AddRouter(url, handler)
	c.router[http.MethodPost].AddRouter(url, handler)
	c.router[http.MethodPut].AddRouter(url, handler)
	c.router[http.MethodPatch].AddRouter(url, handler)
	c.router[http.MethodHead].AddRouter(url, handler)
	c.router[http.MethodDelete].AddRouter(url, handler)
	c.router[http.MethodConnect].AddRouter(url, handler)
	c.router[http.MethodTrace].AddRouter(url, handler)
	c.router[http.MethodOptions].AddRouter(url, handler)
}

func (c *Core) FindRouter(r *http.Request) ControllerHandler {
	uri := r.URL.Path
	method := strings.ToUpper(r.Method)
	if methodTree, ok := c.router[method]; ok {
		return methodTree.FindHandler(uri)
	}
	return nil
}

//实现net/http包下的ServeHttp
func (c *Core) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := NewContext(r, rw)
	//找到路由
    handler := c.FindRouter(r)
	if handler == nil {
		ctx.Json(http.StatusNotFound, "no match handler")
		return
	}
	//处理业务逻辑
	handler(ctx)
}
