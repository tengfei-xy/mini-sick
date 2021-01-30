#!/bin/bash
/usr/local/mysql-5.7.27/bin/mysqldump -B mini_sick_000 -d --compact -u root -p >  ~/mini_sick/option/backup/sql/base_struct.sql 
