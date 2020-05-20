package order

// CheckErr 检查数据库操作
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
