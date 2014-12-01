package godeal

import (
	"testing"
	"fmt"
	"github.com/bmizerany/assert"
)

func TestModel(t *testing.T) {
	p := &Position{116.46, 39.92, "北京", "beijing..."}
	u, u2 := &User {
		Id:0, Account:"13811820678", Name:"lyso", Gender:Male, Phone:"13811820678",
		Pass:"pass", Token:"token",
	},
	&User {
		Id:0, Account:"13811820679", Name:"lyso2", Gender:Female, Phone:"13811820679",
		Pass:"pass", Token:"token",
	}
	fmt.Println("Avatar: |" + u.Avatar + "|")
	u2.Account = "leo2"
	fmt.Println(p)
	fmt.Println(u2)

	newp, e := FromJson(p.String(), &Position{})
	assert.Equal(t, e, nil, "")
	assert.Equal(t, newp, p, "")
	newu, e := FromJson(u.String(), &User{})
	assert.Equal(t, e, nil, "")
	fmt.Println(newp)
	fmt.Println(newu)
}
