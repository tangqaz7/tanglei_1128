package func_judge
import (
	. "tanglei_1128/lv1/data"
)
//判断是否存在用户
func IsExist(user string) bool {
	//如果长度为0说明尚未有用户注册
	if len(Slice) == 0 {
		return false
	} else {
		//遍历切片
		for _, v := range Slice {
			// return v.Name == user //此时只能和第一个比较，所以第一个之后全为false
			if v.Name == user {
				return true
			}
		}
	}
	return false
}
//判断密码是否正确
func IsRight(user string, passwd string) bool {
	for _, v := range Slice {
		if v.Name == user {
			//先确认姓名一致，密码相同返回true
			return v.Passwd == passwd
		}
	}
	return false
}