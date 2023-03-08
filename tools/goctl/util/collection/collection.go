package collection

type Collection struct {
	data []string
}

func NewSet() *Collection {
	return &Collection{}
}

func (c *Collection) AddStr(str ...string) {
	for _, v := range str {
		c.data = append(c.data, v)
	}
}

func (c *Collection) KeysStr() []string {
	return c.data
}
