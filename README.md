# coffer

## performance
```
/Volumes % sync; dd if=/dev/zero of=/Volumes/localhost/tempfile bs=1M count=1024; sync
1024+0 records in
1024+0 records out
1073741824 bytes transferred in 1.640524 secs (654511500 bytes/sec)
/Volumes % sync; dd if=/Volumes/localhost/tempfile of=/dev/null bs=1M count=1024; sync
1024+0 records in
1024+0 records out
1073741824 bytes transferred in 0.105351 secs (10192042069 bytes/sec)
```

```
/Volumes % sync; dd if=/dev/zero of=/Users/x/Documents/tempfile bs=1M count=1024; sync
1024+0 records in
1024+0 records out
1073741824 bytes transferred in 0.157545 secs (6815461132 bytes/sec)
/Volumes % sync; dd if=/Users/x/Documents/tempfile of=/dev/null bs=1M count=1024; sync
1024+0 records in
1024+0 records out
1073741824 bytes transferred in 0.075648 secs (14193922166 bytes/sec)
```