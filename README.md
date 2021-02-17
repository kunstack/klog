# klog
klog is a structured logger for Go (golang), completely API compatible with the standard library logger.

It can be introduced into the project easily, quickly and lightly. Inspired by [Logrus](https://github.com/sirupsen/logrus)

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
