-- 组织
drop table if exists app_org;
create table app_org
(
    id          varchar(30) primary key,
    pid         varchar(30),
    name        varchar(30),
    create_time timestamp(0) not null default now()
);

comment on table app_org is '组织表 ';
comment on column app_org.pid is '上级组织 ';
comment on column app_org.name is '用户昵称 ';
comment on column app_org.create_time is '创建时间';

-- 初始化组织
insert into app_org(id, pid, name)
values (1, '', '根组织');

-- 用户
drop table if exists app_user;
create table if not exists app_user
(
    id          varchar(30) primary key,
    name        varchar(50)  not null,
    account     varchar(50)  not null unique,
    phone       varchar(11)  not null default '',
    email       varchar(100) not null default '',
    password    varchar(100) not null default '',
    picture     varchar(500) not null default '',
    gender      int          not null default 0,
    age         int          not null default 0,
    status      boolean      not null default true,
    create_time timestamp(0)          default now() not null
);
create index user_account_key on app_user (account);
create index user_phone_key on app_user (phone);
create index user_email_key on app_user (email);
create index ape on app_user (account, phone, email);
comment on table app_user is '系统的账号表 ';
comment on column app_user.account is '用户注册账号 对应 ';
comment on column app_user.name is '用户昵称 ';
comment on column app_user.picture is '用户头像 ';
comment on column app_user.password is '用户密码';
comment on column app_user.gender is '性别';
comment on column app_user.age is '年龄';
comment on column app_user.phone is '手机';
comment on column app_user.email is '邮箱';
comment on column app_user.create_time is '创建时间';

-- 初始化root用户
insert into app_user(id, name, picture, account, password, gender, age, phone, email)
values (1, 'root', 'https://im-1252940994.cos.ap-nanjing.myqcloud.com/go.jpg', 'root',
        'PXeFkSq39sRSVXsSaxLWUUWDt45I8tw9mgcY8GE3B/r3VylOko0q727qPChy+uibqcFuCy6w67ruaQ3AyHHeDg==', 0, 12,
        '15527574117', '1219449282@qq.com');

-- 角色
drop table if exists app_role;
create table app_role
(
    id          varchar(30) primary key,
    name        varchar(30),
    role_key    varchar(50)  not null,
    status      boolean      not null default true,
    create_time timestamp(0) not null default now()
);

comment on table app_role is '角色 ';
comment on column app_role.name is '角色名称 ';
comment on column app_role.role_key is '角色标识符 ';
comment on column app_role.create_time is '创建时间';

-- 初始化角色
insert into app_role(id, name, role_key)
values (1, '超级管理员', 'root');
insert into app_role(id, name, role_key)
values (2, '管理员', 'admin');
insert into app_role(id, name, role_key)
values (3, '普通用户', 'normal');

-- 组织角色
drop table if exists app_org_role;
create table app_org_role
(
    id          varchar(30) primary key,
    org_id      varchar(30)  not null,
    role_id     varchar(50)  not null,
    create_time timestamp(0) not null default now()
);
comment on table app_org_role is '组织角色表';
comment on column app_org_role.org_id is '角色所属组织id';
comment on column app_org_role.role_id is '角色id ';
comment on column app_org_role.create_time is '创建时间';

-- 初始化组织角色
insert into app_org_role(id, org_id, role_id)
values (1, 1, 1);

insert into app_org_role(id, org_id, role_id)
values (2, 1, 2);

insert into app_org_role(id, org_id, role_id)
values (3, 1, 3);


-- 组织用户
drop table if exists app_org_user;
create table app_org_user
(
    id           varchar(30) primary key,
    org_id       varchar(30)  not null,
    user_id      varchar(30)  not null,
    first_choice bool         not null default false,
    create_time  timestamp(0) not null default now()
);
comment on table app_org_user is '用户组织关系表';
comment on column app_org_user.org_id is '组织id';
comment on column app_org_user.user_id is '用户id ';
comment on column app_org_user.first_choice is '当前组织是否为首选组织';
comment on column app_org_user.create_time is '创建时间';

-- 初始化组织用户
insert into app_org_user(id, org_id, user_id, first_choice)
values (1, 1, 1, true);

-- 组织用户角色
drop table if exists app_org_user_role;
create table app_org_user_role
(
    id           varchar(30) primary key,
    org_id       varchar(30)  not null,
    user_id      varchar(30)  not null,
    role_id      varchar(50)  not null,
    first_choice bool         not null default false,
    create_time  timestamp(0) not null default now()
);
comment on table app_org_user_role is '组织用户角色分配表';
comment on column app_org_user_role.org_id is '角色所属组织id';
comment on column app_org_user_role.role_id is '角色id ';
comment on column app_org_user_role.user_id is '用户id ';
comment on column app_org_user_role.first_choice is '是否用户首选角色';
comment on column app_org_user_role.create_time is '创建时间';

-- 初始化 root 角色分配
insert into app_org_user_role(id, org_id, user_id, role_id, first_choice)
values (1, 1, 1, 1, true);

insert into app_org_user_role(id, org_id, user_id, role_id, first_choice)
values (2, 1, 1, 2, false);

insert into app_org_user_role(id, org_id, user_id, role_id, first_choice)
values (3, 1, 1, 3, false);

