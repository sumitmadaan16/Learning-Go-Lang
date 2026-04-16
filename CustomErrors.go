package main

import "fmt"

// we create a struct to create a customError on
type customError struct {
	Code    int
	Message string
}

// here (e customError) is called receiver. and Error() is method inside error interface we implement that and override this metod
func (e customError) Error() string {
	return fmt.Sprintf("Error Code is: %d, and Error Message is: %s ", e.Code, e.Message) // elaborate the error anyway we want it to look like
}

func CheckForCustomError(code int) (string, error) {
	if code == 400 {
		return "", &customError{Code: code, Message: "Invalid request"}
	} // passing customError as refrence just to save memory and any possible annomaly as passing as value can create multiple prblems like value not updated etc.
	return "Success", nil
}

func CustomErrorDemo() {
	if _, err := CheckForCustomError(400); err != nil {
		fmt.Println(err)
	}
	var strMsg, err = CheckForCustomError(200)
	{
		fmt.Println(err)
		fmt.Println(strMsg)
	}
}
