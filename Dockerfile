#from scratch
from debian:wheezy
run apt-get update
run apt-get install -y postgresql postgresql-contrib
add go-webapp /
cmd ["/go-webapp"]
