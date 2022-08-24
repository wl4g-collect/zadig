package stepcontroller

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/koderover/zadig/pkg/config"
	"github.com/koderover/zadig/pkg/setting"
	helmtool "github.com/koderover/zadig/pkg/tool/helmclient"
	"github.com/koderover/zadig/pkg/tool/log"
	helmclient "github.com/mittwald/go-helm-client"
)

func init() {
	log.Init(&log.Config{
		Level:       "debug",
		SendToFile:  false,
		Filename:    filepath.Join(os.ExpandEnv(""), "log"),
		Development: config.Mode() != setting.ReleaseMode,
		MaxSize:     5,
	})
}

func TestHelmClient(t *testing.T) {
	fmt.Printf("Into test helm client ....")

	helmClient, err := helmtool.NewClientFromNamespace("", "test1")
	if err != nil {
		fmt.Printf("helmClient: %v, erorr - %v", helmClient, err)
	}

	// historys, err2 := helmClient.ListReleaseHistory("demo", 10)
	// fmt.Printf("err2: %v", err2)
	// for _, his := range historys {
	// 	fmt.Printf("err2: %v", his)
	// }

	chartPath := "/mnt/disk1/__wanglsir_Documents/other-workspace/wl4g-projects/safecloud-charts/all-app-stack"
	if valuesFile, err := os.Open(chartPath + "/values.yaml"); err == nil {
		if replacedMergedValuesYaml, err := ioutil.ReadAll(valuesFile); err == nil {
			chartSpec := helmclient.ChartSpec{
				ReleaseName: "safecloud",
				ChartName:   chartPath,
				Namespace:   "test1",
				ReuseValues: true,
				Version:     "",
				ValuesYaml:  string(replacedMergedValuesYaml),
				SkipCRDs:    false,
				UpgradeCRDs: true,
				Timeout:     time.Second * time.Duration(10),
				Wait:        true,
				Replace:     true,
				MaxHistory:  10,
			}
			// fmt.Printf("chartSpec: %v", chartSpec)

			if release, err2 := helmClient.InstallOrUpgradeChart(context.TODO(), &chartSpec); err2 != nil {
				fmt.Printf("failed to upgrade helm chart")
			} else {
				fmt.Printf("release: %v", release)
			}
		}
	}

}
