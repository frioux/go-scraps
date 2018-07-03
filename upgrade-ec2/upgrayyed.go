package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var svc *ec2.EC2
var mapping = map[ec2.InstanceType]string{
	ec2.InstanceTypeM4Large:    "m5.large",
	ec2.InstanceTypeM4Xlarge:   "m5.xlarge",
	ec2.InstanceTypeM42xlarge:  "m5.2xlarge",
	ec2.InstanceTypeM44xlarge:  "m5.4xlarge",
	ec2.InstanceTypeM410xlarge: "m5.12xlarge",
}

func main() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	cfg.Region = endpoints.UsWest1RegionID
	svc = ec2.New(cfg)

	filter := &ec2.DescribeInstancesInput{
		Filters: []ec2.Filter{{
			Name:   aws.String("instance-state-name"),
			Values: []string{"stopped"},
		}, {
			Name: aws.String("instance-type"),
			Values: []string{
				"m4.large",
				"m4.xlarge",
				"m4.2xlarge",
				"m4.4xlarge",
				"m4.10xlarge",
			},
		}},
	}

	for {
		fmt.Println(time.Now().String(), "checking for more stopped sandboxes")
		resp, err := svc.DescribeInstancesRequest(filter).Send()
		if err != nil {
			panic("unable to load SDK config, " + err.Error())
		}

		for _, reservation := range resp.Reservations {
			for _, instance := range reservation.Instances {
				fmt.Println("Upgrading", *instance.InstanceId)

				m5(instance.InstanceId, instance.InstanceType)
			}

		}
		time.Sleep(time.Minute)
	}
}

func m5(instanceID *string, oldType ec2.InstanceType) {
	_, err := svc.ModifyInstanceAttributeRequest(
		&ec2.ModifyInstanceAttributeInput{
			InstanceId: instanceID,
			EnaSupport: &ec2.AttributeBooleanValue{Value: aws.Bool(true)},
		},
	).Send()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't enable ENA: %s\n", err)
	}

	instanceType := &ec2.AttributeValue{
		Value: aws.String(mapping[oldType]),
	}
	_, err = svc.ModifyInstanceAttributeRequest(
		&ec2.ModifyInstanceAttributeInput{
			InstanceId:   instanceID,
			InstanceType: instanceType,
		},
	).Send()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't change instance type: %s\n", err)
	}

	res, err := svc.StartInstancesRequest(
		&ec2.StartInstancesInput{
			InstanceIds: []string{*instanceID},
		},
	).Send()
	fmt.Println(res)
}
