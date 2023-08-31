package main

// 饿汉式单例
// 注意定义非导出类型
type databaseConn struct {
}

var dbConn *databaseConn

func init() {
	dbConn = &databaseConn{}
}

// GetInstance获取实例
func Db() *databaseConn {
	return dbConn
}
