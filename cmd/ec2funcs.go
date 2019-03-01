package cmd

import "fmt"

func describe() {

	// Example sending a request using the DescribeImagesRequest method.
	req := client.DescribeImagesRequest(params)
	resp, err := req.Send()
	if err == nil {
		fmt.Println(resp)
	}
}
