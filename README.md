# aws-skip-repl

You need the mysql-client installed to use the bash script.

Note:
replit.sh a bash script to help you skip AWS RDS Aurora Replication records. 
Only use this if you are sure it's ok to skip those records.


```bash

./replit.sh          
No arguments provided
Usage: ./replit.sh <AWS-EndPoint>
```

Or


Create a AWS RDS Aurora Event to skip the Replication records.

Note:
Login access to the Aurora database is required if you plan on creating the Event.

```SQL
CREATE EVENT repl_error_skipper 
ON SCHEDULE 
EVERY 15 MINUTE
COMMENT 'Calling rds_skip_repl_error to skip replication error'
Do
CALL mysql.rds_skip_repl_error;

```