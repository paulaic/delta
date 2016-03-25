package delta_test

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"testing"

	"github.com/GeoNet/delta/meta"
)

func TestConsistency(t *testing.T) {

	files := map[string]struct {
		f string
		l meta.List
	}{
		"connections": {f: "../install/connections.csv", l: &meta.ConnectionList{}},
		"cameras":     {f: "../install/cameras.csv", l: &meta.InstalledCameraList{}},
		"dataloggers": {f: "../install/dataloggers.csv", l: &meta.DeployedDataloggerList{}},
		"metsensors":  {f: "../install/metsensors.csv", l: &meta.InstalledMetSensorList{}},
		"radomes":     {f: "../install/radomes.csv", l: &meta.InstalledRadomeList{}},
		"receivers":   {f: "../install/receivers.csv", l: &meta.DeployedReceiverList{}},
		"recorders":   {f: "../install/recorders.csv", l: &meta.InstalledRecorderList{}},
		"sensors":     {f: "../install/sensors.csv", l: &meta.InstalledSensorList{}},
		"firmware":    {f: "../install/firmware.csv", l: &meta.FirmwareHistoryList{}},
		"networks":    {f: "../network/networks.csv", l: &meta.NetworkList{}},
		"stations":    {f: "../network/stations.csv", l: &meta.StationList{}},
		"sites":       {f: "../network/sites.csv", l: &meta.SiteList{}},
		"marks":       {f: "../network/marks.csv", l: &meta.MarkList{}},
		"mounts":      {f: "../network/mounts.csv", l: &meta.MountList{}},
	}

	for k, v := range files {
		if err := meta.LoadList(v.f, v.l); err != nil {
			t.Fatal(err)
		}

		raw, err := ioutil.ReadFile(v.f)
		if err != nil {
			t.Fatal(err)
		}

		var buf bytes.Buffer
		if err := csv.NewWriter(&buf).WriteAll(meta.EncodeList(v.l)); err != nil {
			t.Fatal(err)
		}

		if string(raw) != buf.String() {
			t.Error(k + ": **** csv file mismatch **** : " + v.f)
		}
	}
}
