package myexception

import "fmt"

func DeferTest() {
	if err := test(); err != nil {
		fmt.Printf("mimic error: %v\n", err)
	}
}

func mimicError(key string) error {
	return fmt.Errorf("mimic error: %s", key)
}

func test() (err error) {
	defer func() {
		fmt.Println("start defer")

		if err != nil {
			fmt.Println("defer error:", err)
		}
	}()

	fmt.Println("start test")

	err = mimicError("2")

	fmt.Println("end test")

	return
}
