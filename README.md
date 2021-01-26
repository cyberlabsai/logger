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
    fmt.Println("Without extension")
}
```

In this second example the main package uses the logger.

```go
package main

import "github.com/cyberlabsai/logger"

func main() {
    logger.Info("With extension")
}



```