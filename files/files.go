package files

//go:generate go run files-gogen.go

func init() {

}

// func NewViewsFS() http.FileSystem {
// 	return fsutil.NewHTTPFuncFS(func(name string) (http.File, error) {
// 		return EmbeddedAssets.Open("/views" + path.Clean("/"+name))
// 	})
// }

// func NewIncludesFS() http.FileSystem {
// 	return fsutil.NewHTTPFuncFS(func(name string) (http.File, error) {
// 		return EmbeddedAssets.Open("/includes" + path.Clean("/"+name))
// 	})
// }

// func NewTmplStore() tmpl.Store {
// 	return &tmpl.HFSStore{
// 		FileSystems: map[string]http.FileSystem{
// 			tmpl.ViewsCategory:    NewViewsFS(),
// 			tmpl.IncludesCategory: NewIncludesFS(),
// 		},
// 	}
// }
