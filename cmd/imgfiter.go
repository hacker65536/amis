package cmd

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var (
	filterTmp = []ec2.Filter{
		{
			Name: aws.String("is-public"),
			Values: []string{
				"true",
			},
		},
		{
			Name: aws.String("virtualization-type"),
			Values: []string{
				"hvm",
			},
		},
		{
			Name: aws.String("root-device-type"),
			Values: []string{
				"ebs",
			},
		},
		{
			Name: aws.String("architecture"),
			Values: []string{
				"x86_64",
			},
		},
		{
			Name: aws.String("image-type"),
			Values: []string{
				"machine",
			},
		},
		{
			Name: aws.String("ena-support"),
			Values: []string{
				"true",
			},
		},
		{
			Name: aws.String("sriov-net-support"),
			Values: []string{
				"simple",
			},
		},
		{
			Name: aws.String("state"),
			Values: []string{
				"available",
			},
		},
		{
			Name: aws.String("block-device-mapping.volume-type"),
			Values: []string{
				"gp2",
			},
		},
	}
)

func filter(opt string) []ec2.Filter {
	fs := []ec2.Filter{}

	switch opt {
	case "amz2":
		f := []ec2.Filter{
			{
				Name: aws.String("owner-id"),
				Values: []string{
					"137112412989",
				},
			},
			{
				Name: aws.String("manifest-location"),
				Values: []string{
					"amazon/amzn2-ami-hvm*",
				},
			},
			{
				Name: aws.String("description"),
				Values: []string{
					"Amazon Linux 2 AMI 2.0*",
				},
			},
		}
		fs = append(filterTmp, f...)

	case "cent7":
		f := ec2.Filter{
			Name: aws.String("product-code"),
			Values: []string{
				"aw0evgkw8e5c1q413zgy5pjce",
			},
		}
		fs = append(filterTmp, f)
	case "ubuntu18":
		f := []ec2.Filter{
			{
				Name: aws.String("owner-id"),
				Values: []string{
					"099720109477",
				},
			},
			{
				Name: aws.String("manifest-location"),
				Values: []string{
					"099720109477/ubuntu/images/hvm-ssd/ubuntu-bionic*",
				},
			},
		}
		fs = append(filterTmp, f...)

	}

	return fs

}
