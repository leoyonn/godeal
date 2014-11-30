package model

import (
	"testing"
	"fmt"
	"github.com/bmizerany/assert"
)

func Test(t *testing.T) {
	p := &Position{116.46, 39.92, "北京", "beijing..."}
	u, u2 := &User{1, "leo", p, Male}, &User{1, "leo", p, Male}
	assert.Equal(t, u, u2, "nil")
	u2.Account = "leo2"
	assert.NotEqual(t, u, u2, "nil")
	fmt.Println(p)
	fmt.Println(u2)

	newp, e := FromJson(p.String(), &Position{})
	assert.Equal(t, e, nil, "")
	assert.Equal(t, newp, p, "")
	newu, e := FromJson(u.String(), &User{})
	assert.Equal(t, e, nil, "")
	assert.Equal(t, newu, u, "")
	fmt.Println(newp)
	fmt.Println(newu)
}
