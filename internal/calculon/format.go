package calculon

import (
    "fmt"
    "strconv"
)

func FormatByteToBitString(input byte) string {
    binary := strconv.FormatUint(uint64(input), 2)
    return fmt.Sprintf("%08s", binary)
}