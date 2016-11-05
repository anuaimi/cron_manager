# Spec

## Overiew
RESTful API to manage one-time and repeating jobs.  This is a replacement for CRON with a number of advantages including: 
- don't need ssh to server manage jobs.  instead can use a RESTful API that is language independant.
- tracks history of jobs so it is easy to which jobs failed and when
- ability to retry failed jobs a number of times
- can notify results using a variety of methods
- integration with Promethius to allow monitoring
- can scale beyond a single server

## Details
- there are a log of similarities between a QUEUE and CRON.  Should not have two completely separate tools for this

 
