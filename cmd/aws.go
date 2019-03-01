package cmd

import "github.com/aws/aws-sdk-go-v2/service/ec2"

func initec2svc() {
	ec2svc = ec2.New(cfg)
}
