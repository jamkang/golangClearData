package realization

type SetEle interface {
	SetId(id string)
	SetName(name string)
}

type GetEle interface {
	GetUserData() GetEle
}

type User struct {
	Id   string
	Name string
}

func (self *User) SetId(id string) {
	self.Id = id
}

func (self *User) SetName(Name string) {
	self.Name = Name
}

func (self *User) GetUserData() SetEle {
	return self
}
