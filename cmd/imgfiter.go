package cmd

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var filterList map[string]string = map[string]string{
	"architecture":                     "x86_64",
	"block-device-mapping.volume-type": "gp2",
	"ena-support":                      "true",
	"image-type":                       "machine",
	"is-public":                        "true",
	"root-device-type":                 "ebs",
	"sriov-net-support":                "simple",
	"state":                            "available",
	"virtualization-type":              "hvm",
}

var (
	filterTmp = []ec2.Filter{}
)

func init() {
	for k, v := range filterList {
		filterTmp = append(filterTmp, []ec2.Filter{
			{
				Name:   aws.String(k),
				Values: []string{v},
			},
		}...)
	}
}

func filter(opt string) []ec2.Filter {
	fs := []ec2.Filter{}

	switch opt {
	case "amz":
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
					"amazon/amzn-ami-hvm*",
				},
			},
			{
				Name: aws.String("description"),
				Values: []string{
					"Amazon Linux AMI*",
				},
			},
		}
		fs = append(filterTmp, f...)

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

	case "cent6":
		f := ec2.Filter{
			Name: aws.String("product-code"),
			Values: []string{
				"6x5jmcajty9edm3f211pqjfn2",
			},
		}
		fs = append(filterTmp, f)

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
