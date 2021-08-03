package jsonplaceholder

import (
	"fmt"
)

func (intergration *JsonPlaceholderIntegration) GetTodos() error {
	result, err := intergration.http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		fmt.Printf("Error %s", err)
		return err
	}
	fmt.Printf("Body : %s", result)

	return nil
}
