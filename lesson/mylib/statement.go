package mylib

import (
	"fmt"
	"log"
	"os"
	"time"
)

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func Iflesson() {
	result := by2(10)

	if result == "ok" {
		fmt.Println("great")
	}

	if result2 := by2(10); result2 == "ok" {
		fmt.Println("great 2")
	}
	/*
		num := 6
		if num % 2 == 0 {
			fmt.Println("by 2")
		} else if num % 3 == 0 {
			fmt.Println("by 3")
		} else {
			fmt.Println("else")
		}
	*/

	x, y := 11, 10
	if x == 10 && y == 10 {
		fmt.Println("&&")
	}

	if x == 10 || y == 10 {
		fmt.Println("||")
	}

}

func Forlesson() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}

		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)

		sum := 1
		for sum < 10 {
			sum += sum
			fmt.Println(sum)
		}
		fmt.Println(sum)

		/* 無限ループ
		for {
			fmt.Println("hello")
		}
		*/
	}

}
func Rangelesson() {
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	for i, v := range l {
		fmt.Println(i, v)
	}

	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for _, v := range m {
		fmt.Println(v)
	}

}

func getOsName() string {
	return "mac"
}

func Switchlesson() {
	//os := getOsName()
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows!!")
	default:
		fmt.Println("Default!!", os)
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	}
}

func foo() {
	defer fmt.Println("world foo")

	fmt.Println("hello foo")
}

func Deferlesson() {
	foo()

	defer fmt.Println("world")

	fmt.Println("hello")

	file, _ := os.Open("./main.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}

/*
func LoggingSettings(logFile string){
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
*/

func Loggerlesson() {
	//LoggingSettings("test.log")
	_, err := os.Open("dddd")
	if err != nil {
		log.Fatalln("Exit", err)
	}

	log.Println("logging!")
	log.Printf("%T %v ", "test", "test")

	log.Fatalln("error!!") // 最初のFatalの時点で実行は終了する
	log.Fatalf("%T %v", "test", "teT")
}

func ErrHandlinglesson() {
	file, err := os.Open("./mai.go")
	if err != nil {
		log.Fatal("Error!")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data) // 2つ目のerrだが、1つ目のerrに上書きしている
	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))

	if err := os.Chdir("test"); err != nil {
		log.Fatalln("Error")
	}
}

func thirdPartyConnectDB() {
	panic("Unable to connect database!") //goではpanicを自分で使わないように推奨されている
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func PanicRecoverlesson() {
	save()
	fmt.Println("OK?")
}
