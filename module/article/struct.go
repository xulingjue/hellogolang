package article

import (
//"fmt"
)

type article struct {
	Idarticle    int
	Content      string
	CreateTime   string
	Readnum      int
	Favnum       int
	articleClass articleClass
}

type articleClass struct {
	IdarticleClass int
	Name           string
}
