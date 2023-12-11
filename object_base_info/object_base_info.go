package objectbaseinfo

type ObjectBaseInfo struct {
	Name        string
	Description string
	Created     string
}

func (obj ObjectBaseInfo) GetName() string {
	return obj.Name
}
