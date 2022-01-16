/**
 *@Author luojunying
 *@Date 2022-01-16 19:32
 */
package framework

import (
	"fmt"
	"net/http"
)

type Core struct {

}

func NewCore() *Core {
	return &Core{}
}

//实现net/http包下的ServeHttp
func (c *Core) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("init....")
}
