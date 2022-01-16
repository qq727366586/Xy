/**
 *@Author luojunying
 *@Date 2022-01-16 19:48
 */
package Xy

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Context struct {
	request  *http.Request
	responseWriter http.ResponseWriter
	//是否超时标记
	hasTimeout bool
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request: r,
		responseWriter: w,
	}
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

//query start
func (c *Context) QueryAll() map[string][]string {
	if c.request != nil {
		return c.request.URL.Query()
	}
	return map[string][]string{}
}

func (c *Context) QueryInt(key string, def int) int {
	params := c.QueryAll()
	if val, ok := params[key]; ok {
		lens := len(val)
		if lens > 0 {
			intVal, err := strconv.Atoi(val[lens-1])
			if err != nil {
				return def
			}
			return intVal
		}
	}
	return def
}

func (c *Context) QueryString(key string, def string) string {
	params := c.QueryAll()
	if val, ok := params[key]; ok {
		lenS := len(val)
		if lenS > 0 {
			return val[lenS-1]
		}
	}
	return def
}

func (c *Context) QueryArray(key string, def []string) []string {
	params := c.QueryAll()
	if val, ok := params[key]; ok {
		return val
	}
	return def
}
//query end


//form start
func (c *Context) FormAll() map[string][]string {
	if c.request != nil {
		return c.request.PostForm
	}
	return map[string][]string{}
}

func (c *Context) FormInt(key string, def int) int {
	params := c.FormAll()
	if val, ok := params[key]; ok {
		lens := len(val)
		if lens > 0 {
			intVal, err := strconv.Atoi(val[lens-1])
			if err != nil {
				return def
			}
			return intVal
		}
	}
	return def
}

func (c *Context) FormString(key string, def string) string {
	params := c.FormAll()
	if val, ok := params[key]; ok {
		lenS := len(val)
		if lenS > 0 {
			return val[lenS-1]
		}
	}
	return def
}

func (c *Context) FormArray(key string, def []string) []string {
	params := c.FormAll()
	if val, ok := params[key]; ok {
		return val
	}
	return def
}
//form End

//application/json
func (c *Context) BindJson(obj interface{}) error {
	if c.request != nil {
		body, err := ioutil.ReadAll(c.request.Body)
		if err != nil {
			return err
		}
		c.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		if err = json.Unmarshal(body, obj); err != nil {
			return err
		}

	} else {
		return errors.New("ctx.request empty")
	}
	return nil
}
//application/json End

//response
func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}
	ctx.responseWriter.Header().Set("Content-Type", "application/json")
	ctx.responseWriter.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}
	ctx.responseWriter.Write(byt)
	return nil
}
//response End


//net.context imp start
func (c *Context) Deadline() (deadline time.Time, ok bool){
	return
}

func (c *Context)Done() <-chan struct{} {
	return nil
}

func (c *Context)Err() error {
	return nil
}

func (c *Context)Value(key interface{}) interface{}{
	return nil
}
//net.context imp end


