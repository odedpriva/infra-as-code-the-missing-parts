package stepimpl

import (
	"fmt"
	"os"

	cutils "bitbucket.org/accezz-io/devops-common/utils"
	"bitbucket.org/accezz-io/environment-infra-acceptance/utils"
	"github.com/getgauge-contrib/gauge-go/gauge"
	m "github.com/getgauge-contrib/gauge-go/models"
	"github.com/getgauge-contrib/gauge-go/testsuit"
)

func init() {
	cutils.SetLogrusLogLevelFromEnvironmentVariableWithDefult("ACCEPTANCE_LOG_LEVEL", "INFO")
}

var _ = gauge.Step("setup validate", func() {
	fmt.Print("seems to be working... wow")
})

var _ = gauge.Step("not implemented", func() {
	fmt.Print("not implemented")
})

var _ = gauge.Step("gauge environment", func() {
	for _, pair := range os.Environ() {
		fmt.Println(pair)
	}
})

var _ = gauge.Step("validate environment variables <table>", func(tbl *m.Table) {
	for _, row := range tbl.Rows {
		envVar := row.Cells[0]
		if !utils.IsEnvExist(envVar) {
			testsuit.T.Errorf("%s env does not exist", envVar)
		}
	}
})
