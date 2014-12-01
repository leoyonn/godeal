package godeal

import (
	//	sjson "github.com/bitly/go-simplejson"
	json "encoding/json"
	. "fmt"
	"time"
)

type Jsonable interface {
	ToJson() string
	FromJson() Jsonable
}

type Gender byte

const (
	Male Gender = 1 + iota
	Female
)

var chs = [...]string{"男", "女"}

var ens = [...]string{"Male", "Female"}

func (gender Gender) Ch() string {
	return chs[gender]
}

func (gender Gender) En() string {
	return ens[gender]
}

type Position struct {
	Lon     float64   `json:"lon"`
	Lat     float64   `json:"lat"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
}

type User struct {
	Id          int64   `json:"id"`
	Account     string  `json:"account"`
	Name        string  `json:"name"`
	Desc        string  `json:"desc"`
	Gender      Gender  `json:"gender"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Avatar      string  `json:"avatar"`
	Pass        string  `json:"password"`
	Token       string  `json:"token"`
	CreateAt    time.Time    `json:"createAt"`
}

func (p Position) ToJson() (string, error) {
	v, e := json.Marshal(p)
	if e != nil {
		return "", e
	}
	return string(v), nil
}

func (p Position) String() string {
	v, e := p.ToJson()
	if e != nil {
		return Sprintf("%v", p)
	}
	return v
}

func (u User) ToJson() (string, error) {
	v, e := json.Marshal(u)
	if e != nil {
		return "", e
	}
	return string(v), nil
}

func (u User) String() string {
	v, e := u.ToJson()
	if e != nil {
		return Sprintf("%v", u)
	}
	return v
}

func FromJson(str string, v interface{}) (interface{}, error) {
	e := json.Unmarshal([]byte(str), v)
	return v, e
}

