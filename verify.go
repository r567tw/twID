package twID

import (
	"fmt"
	"regexp"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Verify(IDNumber string) bool {
	// Ref: https://segmentfault.com/a/1190000041413266
	taiwanIDNumber := cases.Title(language.Und).String(IDNumber)
	match, err := regexp.MatchString("^[A-Z][12][0-9]{8}$", taiwanIDNumber)

	if err != nil {
		fmt.Println(err)
		panic("Error")
	}

	return match
}

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
