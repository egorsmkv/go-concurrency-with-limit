# Go Concurrency with Limit

## A goal

To provide a simple example on the question "how to do a limited concurrency in Go?".

### What does this example do?

1. Read a file with domains of mail providers.
2. Extract subdomain, domain and TLD for each domain and get their IP addresses via DNS concurrently.
3. Print results.

## Run

```
go run main.go
```

## Benchmark

### Results for `6100` domains

Concurrency=500:

```
Time start: 2020-09-11 17:16:00.857152197 +0300 EEST m=+0.000333288
Time end: 2020-09-11 17:16:13.591289057 +0300 EEST m=+12.734470178
Time difference: 12.73413689s
```

Concurrency=100:

```
Time start: 2020-09-11 17:15:03.606068609 +0300 EEST m=+0.000312989
Time end: 2020-09-11 17:15:30.193861802 +0300 EEST m=+26.588106172
Time difference: 26.587793183s
```

Concurrency=1:

```
No results, because I am not a crazy man.
```

### Results for `500` domains

Concurrency=100:

```
Time start: 2020-09-11 16:58:32.909812603 +0300 EEST m=+0.000545235
Time end: 2020-09-11 16:58:42.947877372 +0300 EEST m=+10.038609975
Time difference: 10.03806474s
```

Concurrency=10:

```
Time start: 2020-09-11 16:57:09.509066626 +0300 EEST m=+0.000300896
Time end: 2020-09-11 16:57:27.794322435 +0300 EEST m=+18.285556776
Time difference: 18.28525588s
```

Concurrency=1:

```
Time start: 2020-09-11 16:59:06.880638779 +0300 EEST m=+0.000315233
Time end: 2020-09-11 17:00:11.641919047 +0300 EEST m=+64.761595541
Time difference: 1m4.761280308s
```
