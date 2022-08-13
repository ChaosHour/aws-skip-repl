#! /usr/bin/env bash
#
# Created by Kurt Larsen to fix a replication issue.
#
#set -xv

if [ $# -eq 0 ]; then
    printf "No arguments provided\n"
    printf "Usage: %s <AWS-EndPoint>\n" $0
    exit 1
fi


# set the variables for your log.
log=repl.log

# create log file or overrite if already present
printf "Log File - " >> ${log}

# append date to log file
date >> ${log}


until
    mysql -t -h ${1} -e "show slave status\G" | grep -i "Slave_SQL_Running: Yes";
  do
    mysql -t -h ${1} -e "CALL mysql.rds_skip_repl_error";
  sleep 0.1;
done >> ${log}
