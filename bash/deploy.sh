#!/bin/sh

cd /etc/binh/cronmailfe && git pull && npm i && npm run build && cp -r build/* /www/wwwroot/beta.cronmail.net && chown -R www-data:www-data /www/wwwroot/beta.cronmail.net