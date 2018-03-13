# logger
golang logger.

## Installing

### Using *go get*

    $ go get github.com/gzw13999/logger

## Example
    import (
        "github.com/gzw13999/logger"
    )

    func main() {
        logger.LogInit("./logs")
		
		logger.LogInit("./logs")
		logger.Run("run", "run", 1)
		logger.Error("err message")
        ...
    }
