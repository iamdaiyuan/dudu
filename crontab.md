0 */2 * * * /sbin/ntpdate time.windows.com
10 0 * * * killall go
10 0 * * * ps -ef|grep /tmp/go-build |awk '{print $2}'|xargs -i kill {}
20 0 * * * /data/app/redis-3.2.1/src/redis-cli -a smart2016 flushall
30 0 * * * mysqldump -uroot -psmart2016 -q smartdb $(date -d yesterday +\%Y\%m\%d) > /data/www/web/go/src/github.com/hunterhug/GoAmazonWeb/file/sql/usa$(date -d yesterday +\%Y\%m\%d).sql
50 0 * * * mysqldump -uroot -psmart2016 -q smartdb Asin$(date -d yesterday +\%Y\%m\%d) > /data/www/web/go/src/github.com/hunterhug/GoAmazonWeb/file/sql/usaasin$(date -d yesterday +\%Y\%m\%d).sql
10 1 * * * mysqldump -uroot -psmart2016 -q jp_smartdb $(date -d yesterday +\%Y\%m\%d) > /data/www/web/go/src/github.com/hunterhug/GoAmazonWeb/file/sql/jp$(date -d yesterday +\%Y\%m\%d).sql
30 1 * * * mysqldump -uroot -psmart2016 -q jp_smartdb Asin$(date -d yesterday +\%Y\%m\%d) > /data/www/web/go/src/github.com/hunterhug/GoAmazonWeb/file/sql/jpasin$(date -d yesterday +\%Y\%m\%d).sql
0 2 * * * nohup go run /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/usa/ippool.go  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/usa/ippool.log 2>&1 &
5 2 * * * nohup go run /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/usa/urlpool.go  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/usa/urlpool.log 2>&1 &
10 2 * * * nohup go run /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/usa/listmain.go  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/usa/listmain.log 2>&1 &
20 2 * * * nohup go run /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/usa/asinmain.go  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/usa/asinmain.log 2>&1 &
0 2 * * * nohup go run /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/jp/ippool.go  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/jp/ippool.log 2>&1 &
5 2 * * * nohup go run /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/jp/urlpool.go  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/jp/urlpool.log 2>&1 &
10 2 * * * nohup go run /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/jp/listmain.go  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/jp/listmain.log 2>&1 &
20 2 * * * nohup go run /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/jp/asinmain.go  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/jp/asinmain.log 2>&1 &
