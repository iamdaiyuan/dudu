0 */2 * * * /sbin/ntpdate time.windows.com
10 0 * * * killall go
10 0 * * * ps -ef|grep /tmp/go-build |awk '{print $2}'|xargs -i kill {}
20 0 * * * /data/app/redis-3.2.1/src/redis-cli -a smart2016 flushall
10 2 * * * nohup go run /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/usa/listmain.go  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/usa/listmain.log 2>&1 &
20 2 * * * nohup go run /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/usa/asinmain.go  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/usa/asinmain.log 2>&1 &
10 2 * * * nohup go run /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/jp/listmain.go  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/jp/listmain.log 2>&1 &
20 2 * * * nohup go run /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/jp/asinmain.go  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/jp/asinmain.log 2>&1 &
