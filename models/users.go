package models

import "time"

//create table user(id int(11) not null auto_increment primary key,
//username char(20) not null unique,
//password char(16) not null,
//register_time datetime not null
//)ENGINE=InnoDB, default charset=utf8;

type User struct {
	ID            int
	UserName      string
	PassWord      string
	Register_time time.Time
}
