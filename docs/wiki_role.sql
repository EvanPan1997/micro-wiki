create table micro_wiki.wiki_role
(
    DATA_ID        varchar(32)   null primary key,
    DATA_DATE      varchar(8)    null,
    DATA_DATE_TIME varchar(14)   null,
    crt_user       varchar(32)   null,
    crt_date       varchar(8)    null,
    crt_time       varchar(14)   null,
    chg_user       varchar(32)   null,
    chg_date       varchar(8)    null,
    chg_time       varchar(14)   null,
    apv_user       varchar(32)   null,
    apv_date       varchar(8)    null,
    apv_time       varchar(14)   null,
    remarks        varchar(2000) null,
    role_name      varchar(128)  null,
    role_type      char(4)       null,
    authority      varchar(4)    null,
    status         char          null,
    description    varchar(512)  null,
    rsv1           varchar(128)  null,
    rsv2           varchar(128)  null,
    rsv3           varchar(128)  null,
    rsv4           varchar(128)  null,
    rsv5           varchar(128)  null
);

INSERT INTO micro_wiki.wiki_role (DATA_ID, DATA_DATE, DATA_DATE_TIME, crt_user, crt_date, crt_time, chg_user, chg_date,
                                  chg_time, apv_user, apv_date, apv_time, remarks, role_name, role_type, authority,
                                  status, description, rsv1, rsv2, rsv3, rsv4, rsv5)
VALUES ('admin', '20220618', '202206182135', null, null, null, null, null, null, null, null, null, null, 'admin', '1',
        '1', '1', 'System Admin', null, null, null, null, null);
INSERT INTO micro_wiki.wiki_role (DATA_ID, DATA_DATE, DATA_DATE_TIME, crt_user, crt_date, crt_time, chg_user, chg_date,
                                  chg_time, apv_user, apv_date, apv_time, remarks, role_name, role_type, authority,
                                  status, description, rsv1, rsv2, rsv3, rsv4, rsv5)
VALUES ('user', '20220618', '202206182136', null, null, null, null, null, null, null, null, null, null, 'user', '1',
        '1', '1', 'System User', null, null, null, null, null);
