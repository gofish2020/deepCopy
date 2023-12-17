package deepcopy

import (
	"testing"
)

type User struct {
	A int
	b int

	C   *interface{}
	Sli []string

	M map[interface{}]interface{}

	Ptr *User
}

func TestCopy(t *testing.T) {

	a := (interface{}(1))

	srcUser := &User{
		A:   100,
		b:   20,
		C:   &a,
		Sli: []string{"1", "2"},

		M: map[interface{}]interface{}{"test1": 1, "test2": "xxxxx", 3: 3},

		Ptr: &User{
			A:   200,
			b:   30,
			M:   map[interface{}]interface{}{"test1": 1, "test2": "xxxxx", 4: 4},
			Sli: []string{"1", "2"},
			Ptr: (*User)(nil),
		},
	}

	dest := Copy(srcUser)
	u := dest.(*User)
	u.A = 10000
	t.Logf("%#v", u)
	t.Logf("%#v", srcUser)

}

type CInfo struct {
	Id int
}
type Info struct {
	A string

	B struct {
		Id int
	}
	CInfo
}

func TestStructAnonymous(t *testing.T) {

	info := Info{
		A: "xxx",
		B: struct{ Id int }{
			Id: 3,
		},
		CInfo: CInfo{
			Id: 55555,
		},
	}
	copyInfo := Copy(info)
	t.Logf("%#v", copyInfo)
}
