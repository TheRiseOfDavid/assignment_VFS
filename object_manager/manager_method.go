package objectmanager

import objectbaseinfo "virtual_file_system/object_base_info"

type ManagerInterface[T objectbaseinfo.ObjectBaseInfo] interface {
	SelectByName(name string) (*T, error)
	IsExists(name string) bool
}
