create table wiki_role
(
    DATA_ID        VARCHAR(32),
    DATA_DATE      VARCHAR(8),
    DATA_DATE_TIME VARCHAR(14),
    crt_user       VARCHAR(32),
    crt_date       VARCHAR(8),
    crt_time       VARCHAR(14),
    chg_user       VARCHAR(32),
    chg_date       VARCHAR(8),
    chg_time       VARCHAR(14),
    apv_user       VARCHAR(32),
    apv_date       VARCHAR(8),
    apv_time       VARCHAR(14),
    remarks        VARCHAR(2000),
    role_name      VARCHAR(128),
    role_type      CHAR(4),
    authority      VARCHAR(4),
    status         CHAR(1),
    description    VARCHAR(512),
    rsv1           VARCHAR(128),
    rsv2           VARCHAR(128),
    rsv3           VARCHAR(128),
    rsv4           VARCHAR(128),
    rsv5           VARCHAR(128)
);

INSERT INTO micro_wiki.wiki_role (DATA_ID, DATA_DATE, DATA_DATE_TIME, crt_user, crt_date, crt_time, chg_user, chg_date, chg_time, apv_user, apv_date, apv_time, remarks, role_name, role_type, authority, status, description, rsv1, rsv2, rsv3, rsv4, rsv5) VALUES ('admin', '20220618', '202206182135', null, null, null, null, null, null, null, null, null, null, 'admin', '1', '1', '1', 'System Admin', null, null, null, null, null);
INSERT INTO micro_wiki.wiki_role (DATA_ID, DATA_DATE, DATA_DATE_TIME, crt_user, crt_date, crt_time, chg_user, chg_date, chg_time, apv_user, apv_date, apv_time, remarks, role_name, role_type, authority, status, description, rsv1, rsv2, rsv3, rsv4, rsv5) VALUES ('user', '20220618', '202206182136', null, null, null, null, null, null, null, null, null, null, 'user', '1', '1', '1', 'System User', null, null, null, null, null);