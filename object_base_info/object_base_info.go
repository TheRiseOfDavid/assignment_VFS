package objectbaseinfo

import "fmt"

type ObjectBaseInfo struct {
	Name        string
	Description string
	Created     string
}

func (ObjectBaseInfo) selectByName(objects []*ObjectBaseInfo, name string) (*ObjectBaseInfo, error) {
	for _, obj := range objects {
		if obj.Name == name {
			return obj, nil
		}
	}
	return nil, fmt.Errorf("Error: The [%s] doesn't exist.\n", name)
}

func (ObjectBaseInfo) IsExists(objects []*ObjectBaseInfo, name string) bool {
	for _, obj := range objects {
		if obj.Name == name {
			return true
		}
	}
	return false
}
