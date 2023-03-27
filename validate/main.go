package main

import (
        "fmt"
        "context"
        "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
        EmployeeID string `json:"employeeID"`
		Name string `json:"name"`
		Amount string `json:"amount`
		Email string `json:"email"`
		ManagerEmail string `json:"managerEmail"`
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
        return fmt.Sprintf("Hello %s!", event.Name ), nil
}

func main() {
        lambda.Start(HandleRequest)
}