package translateShell

import "fmt"

type Count struct {
	Bing   int64
	Google int64
	Cache  int64
}

func (c *Count) SetBing() {
	c.Bing++
}
func (c *Count) SetGoogle() {
	c.Google++
}
func (c *Count) SetCache() {
	c.Cache++
}
func (c *Count) GetAll() {
	fmt.Printf("从bing获取:%d条\n从google获取:%d条\n从cache获取:%d条\n", c.Bing, c.Google, c.Cache)
}
