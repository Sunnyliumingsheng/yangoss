# yangoss
base on Nginx and docker ,you can quickly build a oss or just store your picture

1. to use this package,you should mkdir a directory to store picture in your server or development enviroment.
for example , I mkdir ~/temp/picture
----
2. docker build -t my-nginx .
docker run -d -v ~/temp/picture:/img -p 80:80 --name testNginx my-nginx
the path "~/temp/picture" can replace as your directory where you like
----
3. after that you can put a img into the ~/temp/picture and use explore or curl to access http://localhost:80/img/yourPictureName
it is easy to work
----
4. ok you can use this package to build a oss
----
config:
