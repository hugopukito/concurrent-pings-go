# Program that run ping command on ips extract from file

## Concurrent

First loop threw ips, doing pings concurrently, not waiting for the finish of the pings on the previous ip to start

## Sequential

Same loop threw ips, waiting for pings to finish on every ips.