0 */2 * * * /sbin/ntpdate time.windows.com
10 0 * * * killall go
10 0 * * * ps -ef|grep /tmp/go-build |awk '{print $2}'|xargs -i kill {}
20 0 * * * /data/app/redis-3.2.1/src/redis-cli -a smart2016 flushall
30 0 * * * mysqldump -uroot -psmart2016 -q smartdb $(date -d yesterday +\%Y\%m\%d) > /data/www/web/go/src/github.com/hunterhug/GoAmazonWeb/file/sql/usa$(date -d yesterday +\%Y\%m\%d).sql.gz
50 0 * * * mysqldump -uroot -psmart2016 -q smartdb Asin$(date -d yesterday +\%Y\%m\%d) > /data/www/web/go/src/github.com/hunterhug/GoAmazonWeb/file/sql/usaasin$(date -d yesterday +\%Y\%m\%d).sql.gz
10 1 * * * mysqldump -uroot -psmart2016 -q jp_smartdb $(date -d yesterday +\%Y\%m\%d) > /data/www/web/go/src/github.com/hunterhug/GoAmazonWeb/file/sql/jp$(date -d yesterday +\%Y\%m\%d).sql.gz
30 1 * * * mysqldump -uroot -psmart2016 -q jp_smartdb Asin$(date -d yesterday +\%Y\%m\%d) > /data/www/web/go/src/github.com/hunterhug/GoAmazonWeb/file/sql/jpasin$(date -d yesterday +\%Y\%m\%d).sql.gz
0 2 * * * /data/www/web/go/src/github.com/hunterhug/AmazonGo/crontab/usa/ip.sh  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/usa/ippool.log 2>&1 &
5 2 * * * /data/www/web/go/src/github.com/hunterhug/AmazonGo/crontab/usa/urlpool.sh  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/usa/urlpool.log 2>&1 &
10 2 * * * /data/www/web/go/src/github.com/hunterhug/AmazonGo/crontab/usa/helpspider.sh  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/usa/listmain.log 2>&1 &
20 2 * * * /data/www/web/go/src/github.com/hunterhug/AmazonGo/crontab/usa/asinspider.sh  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/usa/asinmain.log 2>&1 &
0 2 * * * /data/www/web/go/src/github.com/hunterhug/AmazonGo/crontab/jp/ip.sh  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/jp/ippool.log 2>&1 &
5 2 * * * /data/www/web/go/src/github.com/hunterhug/AmazonGo/crontab/jp/urlpool.sh  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/jp/urlpool.log 2>&1 &
10 2 * * * /data/www/web/go/src/github.com/hunterhug/AmazonGo/crontab/jp/helpspider.sh  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/jp/listmain.log 2>&1 &
20 2 * * * /data/www/web/go/src/github.com/hunterhug/AmazonGo/crontab/jp/asinspider.sh  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/jp/asinmain.log 2>&1 &
0 2 * * * /data/www/web/go/src/github.com/hunterhug/AmazonGo/crontab/de/ip.sh  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/de/ippool.log 2>&1 &
5 2 * * * /data/www/web/go/src/github.com/hunterhug/AmazonGo/crontab/de/urlpool.sh  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/de/urlpool.log 2>&1 &
10 2 * * * /data/www/web/go/src/github.com/hunterhug/AmazonGo/crontab/de/helpspider.sh  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/de/listmain.log 2>&1 &
20 2 * * * /data/www/web/go/src/github.com/hunterhug/AmazonGo/crontab/de/asinspider.sh  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/de/asinmain.log 2>&1 &
0 2 * * * /data/www/web/go/src/github.com/hunterhug/AmazonGo/crontab/uk/ip.sh  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/uk/ippool.log 2>&1 &
5 2 * * * /data/www/web/go/src/github.com/hunterhug/AmazonGo/crontab/uk/urlpool.sh  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/uk/urlpool.log 2>&1 &
10 2 * * * /data/www/web/go/src/github.com/hunterhug/AmazonGo/crontab/uk/helpspider.sh  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/uk/listmain.log 2>&1 &
20 2 * * * /data/www/web/go/src/github.com/hunterhug/AmazonGo/crontab/uk/asinspider.sh  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/uk/asinmain.log 2>&1 &
