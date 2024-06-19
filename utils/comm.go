package utils

import (
	"bufio"
	"crypto/sha256"
	"fmt"

	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func RemoveDuplicates(slice []int64) []int64 {
	encountered := make(map[int64]bool) // 用于存储已经遇到的元素
	result := make([]int64, 0)          // 存储非重复元素的切片
	for _, v := range slice {
		if !encountered[v] {
			encountered[v] = true
			result = append(result, v)
		}
	}

	return result
}

func TrimTrailingZeros(input string) string {
	trimmed := strings.TrimRight(input, "0")
	if strings.HasSuffix(trimmed, ".") {
		return trimmed[:len(trimmed)-1]
	}
	return trimmed
}

func GetSpitDataToInt(idsString string, delimiter string) (ids []int64, err error) {
	if idsString == "" {
		return
	}
	split := strings.Split(idsString, delimiter)
	for _, id := range split {
		var idInt int
		idInt, err = strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, int64(idInt))
	}
	return
}

func BubbleSort(arr []int64) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// 如果前一个元素大于后一个元素，则交换它们
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func ConnectStr(str ...string) string {
	var builder strings.Builder
	for _, s := range str {
		builder.WriteString(s)
	}
	return builder.String()
}

// findMin 返回两个字符串中较小的一个
func FindMin(a, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}
	if LessThen(a, b) {
		return a
	}
	return b
}

// FindMax findMax 返回两个字符串中较大的一个
func FindMax(a, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}
	if LessThen(a, b) {
		return b
	}
	return b
}

// IsOverlap isOverlap 判断两个区间是否重叠
func IsOverlap(min1, max1, min2, max2 string) bool {
	if LessThenAndEqual(min2, max1) && LessThenAndEqual(min1, max2) {
		return true
	}
	return false
}

func ReadWordListFromFile(filePath string) ([]string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var wordList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		wordList = append(wordList, word)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return wordList, nil
}

func RandomWords(wordList []string, numWords int) []string {
	if numWords <= 0 || numWords > len(wordList) {
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	selectedWords := make([]string, numWords)
	for i := 0; i < numWords; i++ {
		randomIndex := rand.Intn(len(wordList))
		selectedWords[i] = wordList[randomIndex]
		wordList = append(wordList[:randomIndex], wordList[randomIndex+1:]...)
	}

	return selectedWords
}

func IsValidPassword(input string) bool {
	if len(input) <= 8 {
		return false
	}
	hasSymbol := false
	hasDigit := false

	for _, char := range input {
		if unicode.IsDigit(char) {
			hasDigit = true
		} else if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			hasSymbol = true
		}

		if hasSymbol && hasDigit {
			return true
		}
	}
	return false
}

func RandomString(length int) string {
	// 定义包含小写字母的字符集合
	const letters = "abcdefghijklmnopqrstuvwxyz123456789"
	rand.Seed(time.Now().UnixNano())
	// 生成随机字符串
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

func Hash256(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	hashBytes := hash.Sum(nil)
	return fmt.Sprintf("%x", hashBytes)
}

/*
统计金额
入参：a:1,b:2,a:3,b:2,c:6
返回：a:4,b:4,c:6
*/
func StatisticalAmount(str string) string {
	m := map[string]string{}
	list := strings.Split(str, ",")
	for _, bean := range list {
		split := strings.Split(bean, ":")
		if 2 == len(split) {
			k, v := split[0], split[1]
			vv, has := m[k]
			if has {
				m[k] = Add(v, vv)
			} else {
				m[k] = v
			}
		}
	}
	result := ""

	// 提取 map 的键到切片 keys 中
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// 对切片进行排序
	sort.Strings(keys)

	//for k, v := range m {
	for _, k := range keys {
		result = ConnectStr(result, ",", k, ":", m[k])
	}
	if 0 < len(result) {
		result = result[1:]
	}
	return result
}

/*
统计金额
入参：1,2,3,4,5
返回：15
*/
func StatisticalAmount2(str string) string {
	result := ""
	list := strings.Split(str, ",")
	for _, bean := range list {
		result = Add(result, bean)
	}
	return result
}
