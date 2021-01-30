#!/bin/bash
mysqlconn="/usr/local/mysql-5.7.27/bin/mysql -u root -p"
olddb="mini_sick_2"
newdb="mini_sick_002"

$mysqlconn -e "CREATE DATABASE $newdb"
params=$($mysqlconn -N -e "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE table_schema='$olddb'")
for name in $params; do
  $mysqlconn -e "RENAME TABLE $olddb.$name to $newdb.$name";
done;
$mysqlconn -e "DROP DATABASE $olddb"
