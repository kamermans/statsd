# statsd

## Introduction

This is a fork of [alexcesaro/statsd](https://github.com/alexcesaro/statsd) client that adds support for logging statsd messages to syslog.

## Documentation

All of the upstream documentation is valid:
https://godoc.org/gopkg.in/alexcesaro/statsd.v2

> Note: Only `kamermans/statsd` should be imported when using this fork, do not import `alexcesaro/statsd.v2`!

To log statsd messages to syslog, you can pass the `Syslog(pri syslog.Priority)` config option to `statsd.New()` or `Client.Clone()`:

```go
c, err := statsd.New(
    statsd.Address("127.0.0.1:514"),
    statsd.TagsFormat(statsd.InfluxDB),
    statsd.Syslog(syslog.LOG_LOCAL5|syslog.LOG_EMERG),
)
if err != nil {
    panic(err)
}

defer c.Close()

c.Increment("foo.counter")
```

### Syslog Facilities and Severities
Syslog requires a *facility* and a *severtity*.  The combination (bitwise `OR`) of these is called the *priority*.

All of the available Syslog facilities and severities are available in the `syslog` package.

**Facilities**
- **kern**: `syslog.LOG_KERN`
- **user**: `syslog.LOG_USER`
- **mail**: `syslog.LOG_MAIL`
- **daemon**: `syslog.LOG_DAEMON`
- **auth**: `syslog.LOG_AUTH`
- **syslog**: `syslog.LOG_SYSLOG`
- **lpr**: `syslog.LOG_LPR`
- **news**: `syslog.LOG_NEWS`
- **uucp**: `syslog.LOG_UUCP`
- **cron**: `syslog.LOG_CRON`
- **authpriv**: `syslog.LOG_AUTHPRIV`
- **ftp**: `syslog.LOG_FTP`
- **local0**: `syslog.LOG_LOCAL0`
- **local1**: `syslog.LOG_LOCAL1`
- **local2**: `syslog.LOG_LOCAL2`
- **local3**: `syslog.LOG_LOCAL3`
- **local4**: `syslog.LOG_LOCAL4`
- **local5**: `syslog.LOG_LOCAL5`
- **local6**: `syslog.LOG_LOCAL6`
- **local7**: `syslog.LOG_LOCAL7`

**Severities**
- **emerg**: `syslog.LOG_EMERG`
- **alert**: `syslog.LOG_ALERT`
- **crit**: `syslog.LOG_CRIT`
- **err**: `syslog.LOG_ERR`
- **warning**: `syslog.LOG_WARNING`
- **notice**: `syslog.LOG_NOTICE`
- **info**: `syslog.LOG_INFO`
- **debug**: `syslog.LOG_DEBUG`

There is a helper function called `statsd.SyslogPriority(facility, severity string)` that will return the `syslog.priority` for the text versions of the facility and severity above:

```go
pri, err := statsd.SyslogPriority("local1", "info")
if err != nil {
    log.Fatalf("Unable to get syslog priority: %v", pri)
}

c, err := statsd.New(
    statsd.Address("127.0.0.1:514"),
    statsd.TagsFormat(statsd.InfluxDB),
    statsd.Syslog(pri),
)
```

## Download

    dep ensure -add github.com/kamermans/statsd

## Example

See the [examples in the documentation](https://godoc.org/gopkg.in/alexcesaro/statsd.v2#example-package).


## License

[MIT](LICENSE)
