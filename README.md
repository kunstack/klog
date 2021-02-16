# klog
klog is a structured logger for Go (golang), completely API compatible with the standard library logger.

- [Document](https://godoc.org/github.com/kunstack/klog)

Only one file, can be easily introduced into the project, fast and lightweight. Inspired by [Logrus](https://github.com/sirupsen/logrus)

# Fully compatible with the Golang log library, it is easy to replace the official log package

Example:

    package main

    import (
        log "github.com/kunstack/klog"
        "os"
    )

    func main() {
        log.Println("this is test string")
        log.Printf("%d this is a test number",22222)
        log.Panic("oh ! panic happen")
    }

# Examples

    package main

    import (
        log "github.com/kunstack/klog"
        "os"
    )

    func main() {
        // Create a new logger
        logger := log.New(os.Stderr, "prefix", log.LstdFlags|log.Ldebug)


        // Print a debug log
        logger.Debug("This is debug msg (Debug)")
        logger.Debugln("This is debug msg (Debugln) ")
        logger.Debugf("This is debug number %d", 2233)

        // Print a log log, level `INFO`
        logger.Info("this is Info msg (Info)")
        logger.Infoln("this is Info msg (Infoln)")
        logger.Infof("%s %d","this is a Info number")

        // Print a log, log level `WARN`
        logger.Warn("...")
        logger.Warnln("...")
        logger.Warnf("...")

        // print a log , log level `ERROR`
        logger.Error("...")
        logger.Errorln("...")
        logger.Errorf("...")

        // Log and throw a panic  log level `PANIC` 
        logger.Panic("...")
        logger.Patnicln("...")
        logger.Patnicf("...")

        // Print the error log and exit the  program, log level  `ERROR`
        logger.Fatal("...")
        logger.Fataln("...")
        logger.Fatalf("...")
    }


# Modify the default log output

The log is output to `os.Stderr` by default. You can use SetOutput to modify the default output object.

    package main

    import (
            "os"

            log "github.com/kunstack/klog"
    )

    func main() {
            outFile, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
            if err != nil {
                    log.Fatal(err.Error())
            }


            log.SetLevel(log.DebugLevel)   //Set the default log display level. If you do not set all levels of logs, they will be output and the log level will not be displayed (in order to be consistent with the official log package)

            // Modify the default log output object
            log.SetOutput(outFile)

            log.Info("xxx")
    }


# Log level and modification of Flag

    const (
        Ldate         = 1 << iota // the date in the local time zone: 2009/01/23  
        Ltime                     // the time in the local time zone: 01:23:23
        Lmicroseconds             // microsecond resolution: 01:23:23.123123.  assumes Ltime.
        Llongfile                 // full file name and line number: /a/b/c/d.go:23
        Lshortfile                // final file name element and line number: d.go:23. overrides Llongfile
        LUTC                      // if Ldate or Ltime is set, use UTC rather than the local time zone
        Ldebug                    // debug level
        Linfo                     // info level
        Lwarn                     // Warn level
        Lerror                    // Error level
        Lpanic                    // panic level
        LstdFlags = Ldate | Ltime | Llongfile // initial values for the standard logger

    )

By default, all levels of output are printed and the log level is not displayed (the goal is to keep it consistent with the standard log package, if you want to display the log level can be set by)

    package main

    import (
            "os"

            log "github.com/kunstack/klog"
    )

    func main() {
            log.SetFlag(log.Flags()  | log.Linfo)   //Set the default log time format, the log level is info, and the log below the info level (debug) will no longer be printed.
            log.Debug("this is debug msg")  //Will not print
            log.Info("xxx")  // print xxx 
    }

