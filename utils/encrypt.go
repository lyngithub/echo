package utils

import (
	"crypto/md5"
	"echo/common/logger"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"reflect"
	"sort"
	"strings"
)

// Md5S md5加密
func Md5S(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func RemoveEmptyFields(data interface{}) {
	value := reflect.ValueOf(data).Elem()
	typeOfValue := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := typeOfValue.Field(i)

		// 检查字段是否是string类型并且值为空
		if fieldType.Type.Kind() == reflect.String && field.String() == "" {
			// 将字段设置为零值
			field.Set(reflect.Zero(fieldType.Type))
		}
	}
}

func CheckSign(param interface{}, secret string, sign string) bool {
	mp, err := interfaceToMap(param)
	if err != nil {
		return false
	}
	v := SignEncrypt(mp, secret)
	return v == sign
}

func interfaceToMap(param interface{}) (map[string]interface{}, error) {
	x, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	res := make(map[string]interface{})
	err = json.Unmarshal(x, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SignEncrypt 参数加密，生成sign
// 加密方式 md5
func SignEncrypt(params map[string]interface{}, secret string) string {
	sp := AscAscii(params)
	var encryptStr string
	for _, v := range sp {
		val, _ := ToString(params[v])
		encryptStr += v + "=" + val + "&"
	}
	encryptStr += "pk=" + secret
	sign := strings.ToUpper(Md5S(encryptStr))
	logger.AdminLog.Info("[OpenApi] "+"加密结果", zap.String("加密key和sign", fmt.Sprintf("%s    %s", encryptStr, sign)))
	return sign
}

// AscAscii 按照Ascii码进行升序排序
func AscAscii(params map[string]interface{}) []string {
	var ps []string
	for k := range params {
		ps = append(ps, k)
	}
	sort.Strings(ps)
	return ps
}
