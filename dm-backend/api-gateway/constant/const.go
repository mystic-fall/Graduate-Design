package constant

const (
	Student   = "student"
	Building  = "building"
	Dormitory = "dormitory"
	Live      = "live"
	DM        = "dormitoryManager"
	Admin     = "admin"
)

const DefaultSheet = "Sheet1"

var (
	HeaderMaps = map[string]map[string]string{
		Student: map[string]string{
			"A1": "学生ID",
			"B1": "学生编号",
			"C1": "姓名",
			"D1": "性别",
			"E1": "密码",
		},
		DM: map[string]string{
			"A1": "学生ID",
			"B1": "学生编号",
			"C1": "姓名",
			"D1": "性别",
			"E1": "密码",
		},
		Admin: map[string]string{
			"A1": "学生ID",
			"B1": "学生编号",
			"C1": "姓名",
			"D1": "性别",
			"E1": "密码",
		},
		Dormitory: map[string]string{
			"A1": "学生ID",
			"B1": "学生编号",
			"C1": "姓名",
			"D1": "性别",
			"E1": "密码",
		},
		Building: map[string]string{
			"A1": "学生ID",
			"B1": "学生编号",
			"C1": "姓名",
			"D1": "性别",
			"E1": "密码",
		},
		Live: map[string]string{
			"A1": "学生ID",
			"B1": "学生编号",
			"C1": "姓名",
			"D1": "性别",
			"E1": "密码",
		},
	}
)
