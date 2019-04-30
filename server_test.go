package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestServerServiceHandler_ChangeApp(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/app_change", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.ChangeApp(ctx, "1234", "24")

	if err != nil {
		t.Errorf("Server.ChangeApp returned %+v, ", err)
	}
}

func TestServerServiceHandler_ListApps(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/app_change_list", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"1": {"APPID": "1","name": "LEMP","short_name": "lemp","deploy_name": "LEMP on CentOS 6 x64","surcharge": 0}}`
		fmt.Fprint(writer, response)
	})

	application, err := client.Server.ListApps(ctx, "1234")

	if err != nil {
		t.Errorf("Server.ListApps returned %+v, ", err)
	}

	expected := []Application{
		{
			AppID:      "1",
			Name:       "LEMP",
			ShortName:  "lemp",
			DeployName: "LEMP on CentOS 6 x64",
			Surcharge:  0,
		},
	}

	if !reflect.DeepEqual(application, expected) {
		t.Errorf("Server.ListApps returned %+v, expected %+v", application, expected)
	}
}

func TestServerServiceHandler_AppInfo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/get_app_info", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"app_info": "test"}`
		fmt.Fprint(writer, response)
	})

	appInfo, err := client.Server.AppInfo(ctx, "1234")

	if err != nil {
		t.Errorf("Server.AppInfo returned %+v, ", err)
	}

	expected := &ServerAppInfo{AppInfo: "test"}

	if !reflect.DeepEqual(appInfo, expected) {
		t.Errorf("Server.AppInfo returned %+v, expected %+v", appInfo, expected)
	}
}

func TestServerServiceHandler_EnableBackup(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/backup_enable", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.EnableBackup(ctx, "1234")

	if err != nil {
		t.Errorf("Server.EnableBackup returned %+v, ", err)
	}
}

func TestServerServiceHandler_DisableBackup(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/backup_disable", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.DisableBackup(ctx, "1234")

	if err != nil {
		t.Errorf("Server.DisableBackup returned %+v, ", err)
	}
}

func TestServerServiceHandler_GetBackupSchedule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/backup_get_schedule", func(writer http.ResponseWriter, request *http.Request) {
		response := `{ "enabled": true,"cron_type": "weekly","next_scheduled_time_utc": "2016-05-07 08:00:00","hour": 8,"dow": 6,"dom": 0}`
		fmt.Fprint(writer, response)
	})

	backup, err := client.Server.GetBackupSchedule(ctx, "1234")

	if err != nil {
		t.Errorf("Server.GetBackupSchedule returned %+v, ", err)
	}

	expected := &BackupSchedule{
		Enabled:  true,
		CronType: "weekly",
		NextRun:  "2016-05-07 08:00:00",
		Hour:     8,
		Dow:      6,
		Dom:      0,
	}

	if !reflect.DeepEqual(backup, expected) {
		t.Errorf("Server.GetBackupSchedule returned %+v, expected %+v", backup, expected)
	}
}

func TestServerServiceHandler_SetBackupSchedule(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/backup_set_schedule", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	bs := &BackupSchedule{
		CronType: "",
		Hour:     23,
		Dow:      2,
		Dom:      3,
	}

	err := client.Server.SetBackupSchedule(ctx, "1234", bs)

	if err != nil {
		t.Errorf("Server.SetBackupSchedule returned %+v, ", err)
	}
}

func TestServerServiceHandler_RestoreBackup(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/restore_backup", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.RestoreBackup(ctx, "1234", "45a31f4")

	if err != nil {
		t.Errorf("Server.RestoreBackup returned %+v, ", err)
	}
}

func TestServerServiceHandler_RestoreSnapshot(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/restore_snapshot", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.RestoreSnapshot(ctx, "1234", "45a31f4")

	if err != nil {
		t.Errorf("Server.RestoreSnapshot returned %+v, ", err)
	}
}

func TestServerServiceHandler_SetLabel(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/label_set", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.SetLabel(ctx, "1234", "new-label")

	if err != nil {
		t.Errorf("Server.SetLabel returned %+v, ", err)
	}
}

func TestServerServiceHandler_SetTag(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/tag_set", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.SetTag(ctx, "1234", "new-tag")

	if err != nil {
		t.Errorf("Server.SetTag returned %+v, ", err)
	}
}

func TestServerServiceHandler_Neighbors(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/neighbors", func(writer http.ResponseWriter, request *http.Request) {
		response := `[2345,1234]`
		fmt.Fprint(writer, response)
	})

	neighbors, err := client.Server.Neighbors(ctx, "1234")

	if err != nil {
		t.Errorf("Server.Neighbors returned %+v, ", err)
	}

	expected := []int{2345, 1234}

	if !reflect.DeepEqual(neighbors, expected) {
		t.Errorf("Server.Neighbors returned %+v, expected %+v", neighbors, expected)
	}
}

func TestServerServiceHandler_EnablePrivateNetwork(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/private_network_enable", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.EnablePrivateNetwork(ctx, "1234", "45a31f4")

	if err != nil {
		t.Errorf("Server.EnablePrivateNetwork returned %+v, ", err)
	}
}

func TestServerServiceHandler_DisablePrivateNetwork(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/private_network_disable", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.Server.DisablePrivateNetwork(ctx, "1234", "45a31f4")

	if err != nil {
		t.Errorf("Server.DisablePrivateNetwork returned %+v, ", err)
	}
}

func TestServerServiceHandler_ListPrivateNetworks(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/server/private_networks", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"net539626f0798d7": {"NETWORKID": "net539626f0798d7","mac_address": "5a:02:00:00:24:e9","ip_address": "10.99.0.3"}}`
		fmt.Fprint(writer, response)
	})

	privateNetwork, err := client.Server.ListPrivateNetworks(ctx, "12345")

	if err != nil {
		t.Errorf("Server.ListPrivateNetworks return %+v, ", err)
	}

	expected := []PrivateNetwork{
		{
			NetworkID:  "net539626f0798d7",
			MacAddress: "5a:02:00:00:24:e9",
			IPAddress:  "10.99.0.3",
		},
	}

	if !reflect.DeepEqual(privateNetwork, expected) {
		t.Errorf("Server.ListPrivateNetworks returned %+v, expected %+v", privateNetwork, expected)
	}
}
