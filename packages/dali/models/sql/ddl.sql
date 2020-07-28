CREATE TABLE IF NOT EXISTS dali_vendor_config
(
    token varchar(64),
    prg_short_code varchar(32),
    vendor_workflow varchar(128),
    workflow_version varchar(16),
    vendor_url varchar(512),
    application_token varchar(64),
    access_token varchar(64),
    timeout_millis integer,
    datetime_created timestamp,
    datetime_modified timestamp,
    PRIMARY KEY (token)
);
CREATE INDEX IF NOT EXISTS dali_vendor_config_index_1 ON dali_vendor_config(prg_short_code);

CREATE TABLE IF NOT EXISTS dali_sys_config
(
    token varchar(64),
    prg_short_code varchar(32),
    config_type varchar(32),
    config_key varchar(128),
    config_value varchar(512),
    datetime_created timestamp,
    datetime_modified timestamp,
    PRIMARY KEY (token)
);
CREATE UNIQUE INDEX IF NOT EXISTS dali_sys_config_index_1 ON dali_sys_config(prg_short_code, config_key);


