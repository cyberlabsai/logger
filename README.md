<img src="https://raw.githubusercontent.com/Yoonit-Labs/android-yoonit-camera/development/logo_cyberlabs.png" width="300">

# logger

Logger is a Logrus extension to quickly use as a Singleton when it is an advantage to your project.

## How it works

The logger initializes independently with the function init(), than you can use their functions in any package.

In case of your main package doesn't use the logger, you must import the package with undercore.
```go
package main

import _ "github.com/cyberlabsai/logger" 

func main() {
    fmt.Println("Without use extension on package main")
}
```

In this second example the main package uses the logger.

```go
package main

import "github.com/cyberlabsai/logger"

func main() {
    logger.Info("with extension")

    logger.Error("shit happens!")
}
```

Logging with custom fields.

```go
package main

import "github.com/cyberlabsai/logger"

func main() {
    // Note that the F is just a easy access to a map.
    logger.WithFields(log.F{
		"user_id": 1,
		"func":    "customMessage",
		"message": "my custom message",
	}).Info("When i log with custom fields")
}
```