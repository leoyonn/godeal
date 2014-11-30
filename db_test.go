package godeal

import (
	"testing"
	. "fmt"
	"github.com/bmizerany/assert"
)

func TestDb(t *testing.T) {
	u, u2 := &User {
		Id:0, Account:"13811820678", Name:"lyso", Gender:Male, Phone:"13811820678",
		Pass:"pass", Token:"token",
	},
	&User {
		Id:0, Account:"13811820679", Name:"lyso2", Gender:Female, Phone:"13811820679",
		Pass:"pass2", Token:"token2",
	}
	Println(u)
	Println(u2)

	assert.Equal(t, u.Id, int64(0))
	e := DelByAccount(u.Account); assert.T(t, e == nil)
	e = DelByAccount(u2.Account); assert.T(t, e == nil)
	id, e := reg(u.Phone, u.Token); assert.T(t, e == nil); assert.T(t, id > 0);
	id, e = reg(u.Phone, u.Token); assert.T(t, e != nil); assert.T(t, id == 0);
	id, e = reg(u2.Phone, u2.Token); assert.T(t, e == nil); assert.T(t, id > 0);
	e = setPass(u.Account, u.Token, u.Pass, "badpass"); assert.T(t, e != nil);
	e = setPass(u.Account, u.Token, u.Pass, ""); assert.T(t, e == nil);
	e = setPass(u.Account, "badtoken", "newpass", u.Pass); assert.T(t, e != nil);
	e = setPass(u.Account, "", "newpass", u.Pass); assert.T(t, e != nil);
	e = setPass(u.Account, u.Token, "newpass", ""); assert.T(t, e != nil);
	e = setPass(u.Account, u.Token, "newpass", "badpass"); assert.T(t, e != nil);
	e = setPass(u.Account, u.Token, "newpass", u.Pass); assert.T(t, e == nil);

	u3, e := login(u.Account, "badpass"); assert.T(t, e != nil)
	u3, e = login(u.Account, "newpass"); assert.T(t, e == nil)

	Println(u3)

}

