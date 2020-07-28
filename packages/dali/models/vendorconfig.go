// vendorconfig.go

package models

import (
    "database/sql"
    guuid "github.com/google/uuid"
    "time"
)

type VendorConfig struct {
    Token            string    `json:"token"`
    PrgShortCode     string    `json:"prg_short_code"`
    VendorWorkflow   string    `json:"vendor_workflow"`
    WorkflowVersion  string    `json:"workflow_version"`
    VendorUrl        string    `json:"vendor_url"`
    ApplicationToken string    `json:"application_token"`
    AccessToken      string    `json:"access_token"`
    TimeoutMillis    int       `json:"timeout_millis"`
    DatetimeCreated  time.Time `json:"datetime_created"`
    DatetimeModified time.Time `json:"datetime_modified"`
}

func (vc *VendorConfig) GetVendorConfig(db *sql.DB) error {
    sql := `SELECT prg_short_code,
              vendor_workflow, 
              workflow_version, 
              vendor_url, 
              application_token, 
              access_token,
              timeout_millis,
              datetime_created, 
              datetime_modified 
            FROM dali_vendor_config WHERE token =$1`
    return db.QueryRow(sql, vc.Token).Scan(
        &vc.PrgShortCode,
        &vc.VendorWorkflow,
        &vc.WorkflowVersion,
        &vc.VendorUrl,
        &vc.ApplicationToken,
        &vc.AccessToken,
        &vc.TimeoutMillis,
        &vc.DatetimeCreated,
        &vc.DatetimeModified)
}

func (vc *VendorConfig) UpdateVendorConfig(db *sql.DB) error {
    now := time.Now()
    vc.DatetimeModified = now
    sql := `UPDATE dali_vendor_config SET 
                vendor_workflow=$1,
                workflow_version=$2, 
                vendor_url=$3, 
                application_token=$4, 
                access_token=$5, 
                timeout_millis=$6,
                datetime_modified=$7
            WHERE token=$8`
    _, err :=
        db.Exec(sql,
            vc.VendorWorkflow,
            vc.WorkflowVersion,
            vc.VendorUrl,
            vc.ApplicationToken,
            vc.AccessToken,
            vc.TimeoutMillis,
            vc.DatetimeModified,
            vc.Token)

    return err
}

func (vc *VendorConfig) DeleteVendorConfig(db *sql.DB) error {
    _, err := db.Exec("DELETE FROM dali_vendor_config  WHERE token=$1", vc.Token)

    return err
}

func (vc *VendorConfig) CreateVendorConfig(db *sql.DB) error {
    uuid := guuid.New()
    now := time.Now()
    vc.Token = uuid.String()
    vc.DatetimeCreated = now
    vc.DatetimeModified = now
    _, err := db.Exec(
        `INSERT INTO dali_vendor_config 
                  (token, 
                  prg_short_code, 
                  vendor_workflow, 
                  workflow_version, 
                  vendor_url, 
                  application_token, 
                  access_token, 
                  timeout_millis, 
                  datetime_created, 
                  datetime_modified)
                VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
        vc.Token,
        vc.PrgShortCode,
        vc.VendorWorkflow,
        vc.WorkflowVersion,
        vc.VendorUrl,
        vc.ApplicationToken,
        vc.AccessToken,
        vc.TimeoutMillis,
        vc.DatetimeCreated,
        vc.DatetimeModified)

    if err != nil {
        return err
    }
    return nil
}

func GetVendorConfigs(db *sql.DB, start, count int) ([]VendorConfig, error) {
    sql := `SELECT token, 
                prg_short_code,
                vendor_workflow,
                workflow_version,
                vendor_url,
                application_token, 
                access_token, 
                timeout_millis, 
                datetime_created, 
                datetime_modified 
            FROM dali_vendor_config
            ORDER BY datetime_created 
                LIMIT $1 OFFSET $2`
    rows, err := db.Query(
        sql, count, start)

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    vendorConfigs := []VendorConfig{}

    for rows.Next() {
        var vc VendorConfig
        if err := rows.Scan(
            &vc.Token,
            &vc.PrgShortCode,
            &vc.VendorWorkflow,
            &vc.WorkflowVersion,
            &vc.VendorUrl,
            &vc.ApplicationToken,
            &vc.AccessToken,
            &vc.TimeoutMillis,
            &vc.DatetimeCreated,
            &vc.DatetimeModified); err != nil {
            return nil, err
        }
        vendorConfigs = append(vendorConfigs, vc)
    }

    return vendorConfigs, nil
}
