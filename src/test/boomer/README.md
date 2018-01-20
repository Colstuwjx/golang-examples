# boomer

### Write DNS test client

like the `boomer.go` said, we made dns query tests.

### Run it!

single test task via:

```
go build -o a.out boomer.go
./a.out --run-tasks dns
```

we can also simply start the master and slave via:

```
# run locust master
pip install locust
locust -f dummy.py --master --master-bind-host=127.0.0.1 --master-bind-port=5557

# run slaves
go run boomer.go --master-host=127.0.0.1 --master-port=5557
```

### K8Sized

TODO.

### Thanks

thanks boomer [author](https://github.com/myzhan/boomer): http://myzhan.github.io/2016/03/01/write-a-load-testing-tool-in-golang/
