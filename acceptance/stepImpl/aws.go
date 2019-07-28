package stepimpl

import (
	"bitbucket.org/accezz-io/environment-infra-acceptance/services"

	"bitbucket.org/accezz-io/environment-infra-acceptance/utils"
	"github.com/getgauge-contrib/gauge-go/gauge"
	"github.com/getgauge-contrib/gauge-go/testsuit"
)

var _ = gauge.Step("validate elastic-ip <dnsEnvName>", testElasticIP)

func testElasticIP(dnsEnvName string) {

	a := services.NewAws()

	if !utils.IsEnvExist(dnsEnvName) {
		testsuit.T.Errorf("%s env does not exist", dnsEnvName)
	}

	if !a.ValidateElasticIP(utils.GetEnvStripColons(dnsEnvName)) {
		testsuit.T.Errorf("test fail")
	}
}
