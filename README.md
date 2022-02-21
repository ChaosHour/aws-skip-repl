# aws-skip-repl

This is kinda a waste of time as I could and should just create the event from below.  But, once again I am trying to learn.
I have not found any functions in the aws-sdk-go to do the same as of yet.


```
CREATE EVENT repl_error_skipper 
ON SCHEDULE 
EVERY 15 MINUTE
COMMENT 'Calling rds_skip_repl_error to skip replication error'
Do
CALL mysql.rds_skip_repl_error;
/*also you can add other logic */
```