#!/bin/bash

GOOS=linux GOARCH=amd64 go build
ssh root@quoteboard.happyspike.net rm -r quoteboardnew
ssh root@quoteboard.happyspike.net mkdir quoteboardnew
scp quoteboard root@quoteboard.happyspike.net:~/quoteboardnew/quoteboard
scp -r public root@quoteboard.happyspike.net:~/quoteboardnew/
ssh root@quoteboard.happyspike.net rm -f deploylocal.sh
scp scripts/deploylocal.sh root@quoteboard.happyspike.net:~/
ssh root@quoteboard.happyspike.net bash deploylocal.sh
# rm quoteboard
echo 'deployment complete!'