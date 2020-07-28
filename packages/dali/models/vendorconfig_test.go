// product_test.go

package models

import (
    "database/sql"
    "io/ioutil"
    "fmt"
    _ "github.com/lib/pq"
    "log"
    "os"
    "strconv"
    "strings"
    "testing"
)

var DB *sql.DB

func TestMain(m *testing.M) {

    portString := os.Getenv("APP_DB_PORT")
    port, _ := strconv.Atoi(portString)
    user := os.Getenv("APP_DB_USERNAME")
    password := os.Getenv("APP_DB_PASSWORD")
    dbname := os.Getenv("APP_DB_NAME")
    host := os.Getenv("APP_DB_HOST")

    connectionString :=
        fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
            host, port, user, password, dbname)

    var err error
    DB, err = sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatal(err)
    }
    err = ensureTablesExists()
    if err == nil {
        code := m.Run()
        cleanUp()
        os.Exit(code)
    }
}

func ensureTablesExists() error {
    sqlStatements, err := ioutil.ReadFile("sql/ddl.sql")
    if err != nil {
        return err
    }
    if _, err := DB.Exec(string(sqlStatements)); err != nil {
        log.Fatal(err)
    }
    return nil
}

func cleanUp() {
    DB.Exec("DELETE FROM dali_vendor_config WHERE prg_short_code = 'unittest_prog'")
}

// TESTS

func TestVendorConfig_CreateVendorConfig(t *testing.T) {
    vc := VendorConfig {
        Token: "",
        PrgShortCode: "unittest_prog",
        VendorWorkflow: "vendorWorkflow",
        WorkflowVersion: "1.0",
        VendorUrl: "http://alloy.co",
        ApplicationToken: "applicationtoken",
        AccessToken: "accesstoken",
        TimeoutMillis: 100,
    }
    err := vc.CreateVendorConfig(DB)
    if err != nil {
        t.Errorf("CreateVendorConfig failed")
        return
    }
    vc2 := VendorConfig{
        Token: vc.Token,
    }
    err2 := vc2.GetVendorConfig(DB)
    if err2 != nil {
        t.Errorf("GetVendorConfig failed")
        return
    }
    if vc2.PrgShortCode != "unittest_prog" {
        t.Errorf("Expected vc2.PrgShortCode = unittest, but got: " + vc2.PrgShortCode)
    }
    if vc2.VendorWorkflow != "vendorWorkflow" {
        t.Errorf("Expected vc2.VendorWorkflow = vendorWorkflow, but got: " + vc2.VendorWorkflow)
    }

}

func TestVendorConfig_GetVendorConfig(t *testing.T) {
    vc := VendorConfig {
        Token: "",
        PrgShortCode: "unittest_prog",
        VendorWorkflow: "vendorWorkflow",
        WorkflowVersion: "1.0",
        VendorUrl: "http://alloy.co",
        ApplicationToken: "applicationtoken",
        AccessToken: "accesstoken",
        TimeoutMillis: 100,
    }
    vc.CreateVendorConfig(DB)

    vc2 := VendorConfig{
        Token: vc.Token,
    }
    err2 := vc2.GetVendorConfig(DB)
    if err2 != nil {
        t.Errorf("GetVendorConfig failed")
        return
    }
    if vc2.PrgShortCode != "unittest_prog" {
        t.Errorf("Expected vc2.PrgShortCode = unittest, but got: " + vc2.PrgShortCode)
    }
    if vc2.VendorWorkflow != "vendorWorkflow" {
        t.Errorf("Expected vc2.VendorWorkflow = vendorWorkflow, but got: " + vc2.VendorWorkflow)
    }
}

func TestGetVendorConfigs(t *testing.T) {
    // Remove any unittest data in the table
    cleanUp()

    vc := VendorConfig {
        Token: "",
        PrgShortCode: "unittest_prog",
        VendorWorkflow: "vendorWorkflow1",
        WorkflowVersion: "1.0",
        VendorUrl: "http://alloy.co",
        ApplicationToken: "applicationtoken",
        AccessToken: "accesstoken",
        TimeoutMillis: 100,
    }
    vc.CreateVendorConfig(DB)

    vc = VendorConfig {
        Token: "",
        PrgShortCode: "unittest_prog",
        VendorWorkflow: "vendorWorkflow2",
        WorkflowVersion: "2.0",
        VendorUrl: "http://alloy.co",
        ApplicationToken: "applicationtoken",
        AccessToken: "accesstoken",
        TimeoutMillis: 100,
    }
    vc.CreateVendorConfig(DB)

    vc = VendorConfig {
        Token: "",
        PrgShortCode: "unittest_prog",
        VendorWorkflow: "vendorWorkflow3",
        WorkflowVersion: "3.0",
        VendorUrl: "http://alloy.co",
        ApplicationToken: "applicationtoken",
        AccessToken: "accesstoken",
        TimeoutMillis: 100,
    }
    vc.CreateVendorConfig(DB)

    vendorConfigs, err := GetVendorConfigs(DB, 0, 3)
    if err != nil {
       t.Errorf("GetVendorConfigs failed")
       fmt.Println(err)
       return
    }
    if len(vendorConfigs) != 3 {
       t.Errorf("Number of rows returned != 3")
    }
    for i := 0; i < len(vendorConfigs); i++ {
       fmt.Println(vendorConfigs[i])
    }

    vc1 := vendorConfigs[0]
    if vc1.VendorWorkflow != "vendorWorkflow1" {
       t.Errorf("Expected vc1.VendorWorflow = vendorWorkflow1, but got: " + vc1.VendorWorkflow)
    }
    vc2 := vendorConfigs[1]
    if vc2.VendorWorkflow != "vendorWorkflow2" {
       t.Errorf("Expected vc2.VendorWorflow = vendorWorkflow2, but got: " + vc2.VendorWorkflow)
    }
    vc3 := vendorConfigs[2]
    if vc3.VendorWorkflow != "vendorWorkflow3" {
       t.Errorf("Expected vc3.VendorWorflow = vendorWorkflow3, but got: " + vc3.VendorWorkflow)
    }
}

func TestVendorConfig_UpdateVendorConfig(t *testing.T) {
    vc := VendorConfig {
        Token: "",
        PrgShortCode: "unittest_prog",
        VendorWorkflow: "vendorWorkflow",
        WorkflowVersion: "1.0",
        VendorUrl: "http://alloy.co",
        ApplicationToken: "applicationtoken",
        AccessToken: "accesstoken",
        TimeoutMillis: 100,
    }
    vc.CreateVendorConfig(DB)

    vc.VendorUrl = "http://new_alloy.co"
    err := vc.UpdateVendorConfig(DB)
    if err !=  nil {
        t.Errorf("UpdateVendorConfig failed")
        return
    }

    vc2 := VendorConfig{
        Token: vc.Token,
    }
    err2 := vc2.GetVendorConfig(DB)
    if err2 != nil {
        t.Errorf("GetVendorConfig failed")
    }
    if vc2.VendorUrl != "http://new_alloy.co" {
        t.Errorf("Expected vc2.VendorUrl = http://new_alloy.co but got: %s", vc2.VendorUrl)
    }
}

func TestVendorConfig_DeleteVendorConfig(t *testing.T) {
    vc := VendorConfig {
        Token: "",
        PrgShortCode: "unittest_prog",
        VendorWorkflow: "vendorWorkflow",
        WorkflowVersion: "1.0",
        VendorUrl: "http://alloy.co",
        ApplicationToken: "applicationtoken",
        AccessToken: "accesstoken",
        TimeoutMillis: 100,
    }
    vc.CreateVendorConfig(DB)
    token := vc.Token

    vc.DeleteVendorConfig(DB)

    vc2 := VendorConfig{
        Token: token,
    }
    err2 := vc2.GetVendorConfig(DB)
    if err2 == nil {
        t.Errorf("DeleteVendorConfig failed")
    }
    if !strings.Contains(err2.Error(), "no rows in result set")  {
        t.Errorf("Expected: no rows in result set")
    }
}



