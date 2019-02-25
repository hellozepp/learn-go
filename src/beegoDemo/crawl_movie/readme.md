项目业务:
使用beegoweb框架爬取豆瓣某网页下面url,添加Redis queue中,并添加redis set url去重,如果该url为电影详情页则存储mysql并出队列,存储所有访问的url
启动步骤:
启动 server
redis-server
启动 web
bee run