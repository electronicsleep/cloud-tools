package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/spf13/cobra"
)

var aws_ec2Cmd = &cobra.Command{
	Use:   "aws-ec2",
	Short: "aws-ec2 list",
	Long:  `aws-ec2 - list instances`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("cmd: aws-ec2: args: ", args)
		fmt.Print(" env: ", rootOpts.env)
		fmt.Println(" region:", rootOpts.region)
		if len(args) == 0 || args[0] == "list" {
			aws_ec2_list()
		}
	},
}

func init() {
	RootCmd.AddCommand(aws_ec2Cmd)
}

func aws_ec2_list() {
	fmt.Println("aws-ec2: list")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(rootOpts.region))

	if err != nil {
		fmt.Printf("failed to initialize new session: %v", err)
		return
	}

	svc := ec2.NewFromConfig(cfg)

	runningInstances, err := GetRunningInstances(svc)
	if err != nil {
		fmt.Printf("error retrieving instances: %v", err)
		return
	}

	for _, reservation := range runningInstances.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("%s", *instance.PublicIpAddress)
			fmt.Printf(" # %s\n", *instance.InstanceId)
		}
	}
}

func GetRunningInstances(client *ec2.Client) (*ec2.DescribeInstancesOutput, error) {
	result, err := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})

	if err != nil {
		return nil, err
	}

	return result, err
}
