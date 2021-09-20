package baseball

// import (
// 	"fmt"
// 	"math"
// 	"math/rand"
// 	"strconv"
// 	"time"
// )

// var diffArr = []int{3, 4}

// func Contains(arr []int, b int) bool {
// 	for i := range arr {
// 		if arr[i] == b {
// 			return true
// 		}
// 	}
// 	return false
// }

// func MakeAnswer(diff int) (ansArr []int, answer int) {
// 	ansArr = make([]int, diff)
// 	answer = 0

// 	setting := rand.NewSource(time.Now().UnixNano())
// 	setRand := rand.New(setting)

// 	for i := range ansArr {
// 		ansArr[i] = setRand.Intn(10)
// 		ansArrPart := ansArr[:i]
// 		for Contains(ansArrPart, ansArr[i]) {
// 			// fmt.Print("숫자를 재설정합니다.\n")
// 			ansArr[i] = setRand.Intn(10)
// 		}
// 		// fmt.Printf("%v에 %d가 들어가지 않음\n", ansArrPart, ansArr[i])

// 		answer += ansArr[i] * int(math.Pow10(diff-1-i))
// 	}
// 	fmt.Println("랜덤 숫자 생성 완료. 숫자야구 시작!")
// 	return
// }

// func getUserInput(userInput *string, diff int) {

// 	var a string

// 	for {
// 		fmt.Scan(&a)
// 		input_int, err2 := strconv.Atoi(a)
// 		if err2 != nil {
// 			fmt.Println("잘못된 입력입니다. 다시 입력해 주세요.")
// 			continue
// 		} else if input_int >= int(math.Pow10(diff)) {
// 			fmt.Println("범위가 아닙니다. 다시 입력해 주세요.")
// 			continue
// 		} else {
// 			check := false
// 			for i := 1; i < diff; i++ {
// 				for j := 0; j < i; j++ {
// 					if a[i] == a[j] {
// 						check = true
// 					}
// 				}
// 			}
// 			if check {
// 				fmt.Println("중복된 값이 있습니다. 다시 입력해 주세요.")
// 				continue
// 			} else {
// 				*userInput = a
// 				break
// 			}
// 		}
// 	}
// 	*userInput = a
// 	return
// }

// func CheckNumber(input string, arr []int) (ball int, strike int) {
// 	ball, strike = 0, 0
// 	for i := 0; i < len(input); i++ {
// 		temp, _ := strconv.Atoi(input[i : i+1])
// 		if temp == arr[i] {
// 			strike++
// 		} else if Contains(arr, temp) {
// 			ball++
// 		}
// 	}
// 	return
// }

// func main() {
// 	var diff int

// 	fmt.Print("난이도를 입력해 주세요(3 or 4): ")
// 	fmt.Scanln(&diff)

// 	if !Contains(diffArr, diff) {
// 		fmt.Println("잘못된 입력입니다.")
// 		return
// 	}

// 	fmt.Printf("\n난이도가 %d로 설정되었습니다.\n", diff)

// 	ansArr, _ := MakeAnswer(diff)

// 	// fmt.Printf("답은 %d\n", answer)
// 	// for i := range ansArr {
// 	// 	fmt.Printf("%d번째 자릿수는 %d\n", i+1, ansArr[i])
// 	// }

// 	input := ""
// 	count := 0
// 	for {
// 		count++
// 		fmt.Printf("%d회차 : %d자리 숫자를 입력해 주세요 (중복 제외) : ", count, diff)
// 		getUserInput(&input, diff)
// 		fmt.Printf("입력한 수는 %s입니다.\n", input)
// 		ball, strike := CheckNumber(input, ansArr)
// 		if strike == diff {
// 			fmt.Println("정답!")
// 			break
// 		} else if ball+strike == 0 {
// 			fmt.Println("OUT")
// 		} else {
// 			fmt.Printf("%d Ball %d Strike\n", ball, strike)
// 		}
// 	}
// 	fmt.Printf("%d번만에 맞추셨습니다. 숫자야구 프로그램이 종료됩니다.", count)
// }
