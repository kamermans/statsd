package statsd

import (
	"fmt"
	"log/syslog"
	"os"
	"time"
)

var (
	syslogFacilities = map[string]syslog.Priority{
		"kern":     syslog.LOG_KERN,
		"user":     syslog.LOG_USER,
		"mail":     syslog.LOG_MAIL,
		"daemon":   syslog.LOG_DAEMON,
		"auth":     syslog.LOG_AUTH,
		"syslog":   syslog.LOG_SYSLOG,
		"lpr":      syslog.LOG_LPR,
		"news":     syslog.LOG_NEWS,
		"uucp":     syslog.LOG_UUCP,
		"cron":     syslog.LOG_CRON,
		"authpriv": syslog.LOG_AUTHPRIV,
		"ftp":      syslog.LOG_FTP,
		"local0":   syslog.LOG_LOCAL0,
		"local1":   syslog.LOG_LOCAL1,
		"local2":   syslog.LOG_LOCAL2,
		"local3":   syslog.LOG_LOCAL3,
		"local4":   syslog.LOG_LOCAL4,
		"local5":   syslog.LOG_LOCAL5,
		"local6":   syslog.LOG_LOCAL6,
		"local7":   syslog.LOG_LOCAL7,
	}

	syslogSeverities = map[string]syslog.Priority{
		"emerg":   syslog.LOG_EMERG,
		"alert":   syslog.LOG_ALERT,
		"crit":    syslog.LOG_CRIT,
		"err":     syslog.LOG_ERR,
		"warning": syslog.LOG_WARNING,
		"notice":  syslog.LOG_NOTICE,
		"info":    syslog.LOG_INFO,
		"debug":   syslog.LOG_DEBUG,
	}

	syslogTag = "statsd"
)

// syslogHeader returns a syslog header; formatting implementation
// is from the core Go syslog package.
// @see https://golang.org/src/log/syslog/syslog.go
func SyslogHeader(pri syslog.Priority) string {

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	// Ensure there is only one newline at the end of the message
	return fmt.Sprintf("<%d>%s %s %s[%d]: ",
		pri, time.Now().Format(time.RFC3339), hostname,
		syslogTag, os.Getpid())
}

// SyslogPriority gets the syslog priority from the string facility and severity
func SyslogPriority(facility, severity string) (syslog.Priority, error) {

	facilityPri, ok := syslogFacilities[facility]
	if !ok {
		return 0, fmt.Errorf("Invalid syslog facility '%v'", facility)
	}

	severityPri, ok := syslogSeverities[severity]
	if !ok {
		return 0, fmt.Errorf("Invalid syslog severity '%v'", severity)
	}

	return facilityPri | severityPri, nil
}
