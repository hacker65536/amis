package cmd

import (
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var (
	ec2svc *ec2.EC2
)

func init() {
	ec2svc = ec2.New(cfg)
}

func describe(opt string) {

	// Example sending a request using the DescribeImagesRequest method.
	input := &ec2.DescribeImagesInput{
		Filters: filter(opt),
	}
	req := ec2svc.DescribeImagesRequest(input)
	resp, err := req.Send()
	if err == nil {
		//		fmt.Println(resp)
	}

	img := sortImg(resp.Images)

	for _, v := range img {

		fmt.Println(aws.StringValue(v.Description), (aws.StringValue(v.ImageId)))
	}
}

func sortImg(obj []ec2.Image) []ec2.Image {
	sort.Slice(
		obj,
		func(i, j int) bool {
			return aws.StringValue(obj[i].CreationDate) > aws.StringValue(obj[j].CreationDate)
		},
	)
	return obj

}
