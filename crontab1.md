* */1 * * * /sbin/ntpdate time.windows.com
* 1 * * * killall go
* 1 * * * ps -ef|grep /tmp/go-build |awk '{print $2}'|xargs -i kill {}
20 1 * * * /data/app/redis-3.2.1/src/redis-cli -a smart2016 flushall
10 2 * * * nohup go run /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/usa/listmain.go  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/sh/usa/listmain.log 2>&1 &
20 2 * * * nohup go run /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/usa/asinmain.go  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/sh/usa/asinmain.log 2>&1 &
10 2 * * * nohup go run /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/jp/listmain.go  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/sh/jp/listmain.log 2>&1 &
20 2 * * * nohup go run /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/jp/asinmain.go  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/sh/jp/asinmain.log 2>&1 &
