package twID

import (
	"fmt"
	"regexp"
	"strings"
)

func Verify(IDNumber string) bool {
	taiwanIDNumber := strings.Title(IDNumber)

	match, err := regexp.MatchString("^[A-Z][12][0-9]{8}$", taiwanIDNumber)

	if err != nil {
		fmt.Println(err)
		panic("Verify ERROR")
	}

	fmt.Println(taiwanIDNumber)
	return match
}

// 需不需要再輸入時幫忙轉開頭大寫, verify 給個預設值這樣 如果他要求一定要大寫則就直接送 如果沒有則幫他轉一轉就可以
//  #將i轉為首字母大寫，其餘字母小寫
// i=i.capitalize()
// if not re.match('^[A-Z][12][0-9]{8}$',i):
// 	return '此身分證格式不正確,需為0-9,A-Z的字串'
// else:
// 	a=[]
// 	a.extend('10987654932210898765431320')
// 	c=int(a[ord(i[0])-65])+int(i[9])
// 	for x in range(1,9):
// 		c+=int(i[x])*(9-x)

// 	if c%10!=0:
// 		return '此身分證格式不正確,不通過驗證'
// 	else:
// 		return  '此身分證格式正確'
