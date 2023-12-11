package objectbaseinfo

type ObjectBaseInfo struct {
	name        string
	description string
	created     string
}

func (obj *ObjectBaseInfo) GetName() string {
	return obj.name
}
