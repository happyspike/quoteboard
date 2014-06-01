rm -rf quoteboardold
mv quoteboard quoteboardold
mv quotebaordnew quoteboard
for KILLPID in `ps -A | grep quoteboard | grep -v grep  | awk ' { print $1;}'`; do kill -9 $KILLPID; done;
quoteboard/quoteboard &