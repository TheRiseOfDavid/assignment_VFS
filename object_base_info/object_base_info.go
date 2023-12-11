package objectbaseinfo

type ObjectBaseInfo struct {
	Name        string
	Description string
	Created     string
}

// 你會注意到，go 的泛型超級難用
// 棄用，等 go 技術文件有更好的 example 再來更新
// https://go.dev/doc/tutorial/generics
//https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/generics
// func (ObjectBaseInfo) SelectByName(objects []*ObjectBaseInfo, name string) (*ObjectBaseInfo, error) {
// 	for _, obj := range objects {
// 		if obj.Name == name {
// 			return obj, nil
// 		}
// 	}
// 	return nil, fmt.Errorf("Error: The [%s] doesn't exist.\n", name)
// }

// func (ObjectBaseInfo) IsExists(objects []ObjectBaseInfo, name string) bool {
// 	for _, obj := range objects {
// 		if obj.Name == name {
// 			return true
// 		}
// 	}
// 	return false
// }
