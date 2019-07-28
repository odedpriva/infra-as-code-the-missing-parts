package services

import (
	"errors"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"bitbucket.org/accezz-io/environment-infra-acceptance/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

type laws struct{}

func NewAws() *laws {
	return &laws{}
}

func (a *laws) ValidateElasticIP(lbDNS string) bool {

	var (
		found = false
	)

	service, err := elbV2AWSService()
	if err != nil {
		logrus.Error(err)
		return false
	}
	lbs, err := service.DescribeLoadBalancers(nil)
	if err != nil {
		logrus.Error(err)
		return false
	}

	for _, lb := range lbs.LoadBalancers {
		if lbDNS == *lb.DNSName {
			found = true
			for _, az := range lb.AvailabilityZones {
				if len(az.LoadBalancerAddresses) == 0 {
					logrus.Error(fmt.Errorf("%s does not have elastic ip attached to it", lbDNS))
					return false
				}
			}
		}
	}

	if !found {
		logrus.Error(fmt.Errorf("could not find lb %s", lbDNS))
		return false
	}

	return true
}

// service utility functions

func awsSession() (*session.Session, error) {

	c, err := setAwsConfig()

	if err != nil {
		return nil, err
	}

	return session.NewSession(c)
}

// elbV2AWSService returns aws elbv2 session
func elbV2AWSService() (*elbv2.ELBV2, error) {

	sess, err := awsSession()
	if err != nil {
		return nil, err
	}

	return elbv2.New(sess), nil
}

func setAwsConfig() (*aws.Config, error) {

	if !utils.IsEnvExist("ENVIRONMENT_REGION") {
		return nil, errors.New("ENVIRONMENT_REGION does not exist")
	}

	return &aws.Config{
		Region: aws.String(os.Getenv("ENVIRONMENT_REGION")),
	}, nil

}
