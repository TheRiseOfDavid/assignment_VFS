package objectmanager

import (
	"fmt"
	objectbaseinfo "virtual_file_system/object_base_info"
)

type ObjectManager[T objectbaseinfo.IdentifyMethod] struct {
	*objectbaseinfo.ObjectBaseInfo
	Members *[]T
}

func (obj *ObjectManager[T]) SelectByName(name string) (*T, error) {
	errMsg := fmt.Errorf("Error: The [%s] doesn't exist.\n", name)
	if obj == nil || obj.members == nil {
		return nil, errMsg
	}
	fmt.Println("go")
	for i, member := range *obj.members {

		if member.GetName() == name {
			return &(*obj.members)[i], nil
		}
	}
	return nil, errMsg
}

func (obj *ObjectManager[T]) IsExists(name string) bool {
	if obj == nil || obj.members == nil {
		return false
	}

	for _, member := range *obj.members {
		if member.GetName() == name {
			return true
		}
	}
	return false
}
