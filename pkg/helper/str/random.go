package str

import (
	"fmt"
	"math/rand"
)

// RandType ...
type RandType int

const (
	// 大写字母
	T_CAPITAL RandType = iota + 1
	// 小写字母
	T_LOWERCASE
	// 数字
	T_NUMBERIC
	// 小写字母+数字
	T_LOWERCASE_NUMBERIC
	// 大写字母+数字
	T_CAPITAL_NUMBERIC
	// 大写字母+小写字母
	T_CAPITAL_LOWERCASE
	// 数字+字母
	T_ALL
)

const (
	StrNumberic    = `0123456789`
	StrLowercase   = `abcdefghijklmnopqrstuvwxyz`
	StrCapital     = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	StrPunctuation = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

func (rt RandType) String() string {
	switch rt {
	case T_CAPITAL:
		return StrCapital
	case T_LOWERCASE:
		return StrLowercase
	case T_NUMBERIC:
		return StrNumberic
	case T_CAPITAL_LOWERCASE:
		return fmt.Sprint(StrCapital, StrLowercase)
	case T_CAPITAL_NUMBERIC:
		return fmt.Sprint(StrCapital, StrNumberic)
	case T_LOWERCASE_NUMBERIC:
		return fmt.Sprint(StrLowercase, StrNumberic)
	case T_ALL:
		return fmt.Sprint(StrCapital, StrLowercase, StrNumberic)
	}
	return ""
}

func getRandType() RandType {
	var ct = []RandType{
		// 大写字母
		T_CAPITAL,
		// 小写字母
		T_LOWERCASE,
		// 数字
		T_NUMBERIC,
		// 小写字母+数字
		T_LOWERCASE_NUMBERIC,
		// 大写字母+数字
		T_CAPITAL_NUMBERIC,
		// 大写字母+小写字母
		T_CAPITAL_LOWERCASE,
		// 数字+字母
		T_ALL,
	}
	l := len(ct)
	i := rand.Intn(l)
	return ct[i]
}

func getCharactorFromStr(str string) string {
	strLen := len(str)
	return string(([]rune(str))[rand.Intn(strLen-1)])
}

//func getFillStr(rt RandType) string {
//	switch rt {
//	case T_CAPITAL:
//		return StrCapital
//	case T_LOWERCASE:
//		return StrLowercase
//	case T_NUMBERIC:
//		return StrNumberic
//	case T_CAPITAL_LOWERCASE:
//		return fmt.Sprint(StrCapital, StrLowercase)
//	case T_CAPITAL_NUMBERIC:
//		return fmt.Sprint(StrCapital, StrNumberic)
//	case T_LOWERCASE_NUMBERIC:
//		return fmt.Sprint(StrLowercase, StrNumberic)
//	case T_ALL:
//		return fmt.Sprint(StrCapital, StrLowercase, StrNumberic)
//	}
//	return ""
//}

func RandBetween(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// Rand 默认随机生成6-32位的随机字符串(长度类型皆随机)
// 如果传入不同的参数,则分别对应不同的函数
func Rand(args ...interface{}) string {
	switch len(args) {
	case 1:
		return Random(args[0].(int))
	case 2:
		if r, ok := args[1].(RandType); ok {
			return Random(args[0].(int), r)
		}
		return RandomBetween(args[0].(int), args[1].(int))
	case 3:
		return RandomBetween(args[0].(int), args[1].(int), args[2].(RandType))
	default:
		var rt = getRandType()
		var length = RandBetween(6, 32)
		return randomReal(length, rt)
	}
}

// Random 随机生成指定长度的随机字符串(类型可选或随机)
func Random(length int, fill ...RandType) string {
	var rt RandType
	if len(fill) > 0 {
		rt = fill[0]
	} else {
		rt = getRandType()
	}
	return randomReal(length, rt)
}

// RandomBetween 随机生成指定长度区间的随机字符串(类型可选或随机)
func RandomBetween(min, max int, fill ...RandType) string {
	var rt RandType
	if len(fill) > 0 {
		rt = fill[0]
	} else {
		rt = getRandType()
	}
	var length = RandBetween(min, max)
	return randomReal(length, rt)
}

func randomReal(length int, fill RandType) string {
	str := fill.String()
	var res string
	var i = length
	for i > 0 {
		res += getCharactorFromStr(str)
		i--
	}
	return res
}

func randomRealDo(length int, source string) string {
	var res string
	var i = length
	for i > 0 {
		res += getCharactorFromStr(source)
		i--
	}
	return res
}

func RandCapital(length ...int) string {
	if len(length) > 0 {
		return Random(length[0], T_CAPITAL)
	}
	return RandomBetween(6, 32, T_CAPITAL)
}

func RandLowercase(length ...int) string {
	if len(length) > 0 {
		return Random(length[0], T_LOWERCASE)
	}
	return RandomBetween(6, 32, T_LOWERCASE)
}

func RandString(length ...int) string {
	if len(length) > 0 {
		return Random(length[0], T_CAPITAL_LOWERCASE)
	}
	return Random(RandBetween(6, 32), T_CAPITAL_LOWERCASE)
}

func RandNumberic(length ...int) string {
	if len(length) > 0 {
		return Random(length[0], T_NUMBERIC)
	}
	return Random(RandBetween(6, 32), T_NUMBERIC)
}

func RandAll(length ...int) string {
	if len(length) > 0 {
		return Random(length[0], T_ALL)
	}
	return Random(RandBetween(6, 32), T_ALL)
}

func RandCustom(length int, source string) string {
	return randomRealDo(length, source)
}
