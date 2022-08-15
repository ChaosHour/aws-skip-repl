# aws-skip-repl

You need the mysql-client installed to use the bash script.

**Note:**

> replit.sh is a bash script to help you skip AWS RDS Aurora Replication records. 
> Only use this if you are sure it's ok to skip those records.


```bash

You can put it in a screen session to keep it running.

screen -d -RR scan


If you have the watch command installed, you can use it to keep the script running.

watch -n 1800 -c ./rel2.sh

If you don't have the watch command, you can use a while true loop to keep the script running.
while :; do ./repl2.sh ; sleep 1800 ; done

Or you can use cron to run it every 30 minutes.

*/30 * * * * /path/to/repl2.sh

```

```bash
./replit.sh          
 No arguments provided
Usage: ./replit.sh <AWS-EndPoint>
```


> Create a AWS RDS Aurora Event to skip the Replication records.

**Note:**

Login access to the Aurora database is required if you plan on creating the Event.


```SQL
CREATE EVENT repl_error_skipper 
ON SCHEDULE 
EVERY 15 MINUTE
COMMENT 'Calling rds_skip_repl_error to skip replication error'
Do
CALL mysql.rds_skip_repl_error;

```