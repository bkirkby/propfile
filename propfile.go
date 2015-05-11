package propfile

import (
  "bufio"
  "os"
  "strings"
)

func ReadFileInto( propMap map[string]string, fileName string) {
  f, err := os.Open( fileName)
  if err != nil {
  }
  defer f.Close()

  scanner := bufio.NewScanner( f)
  for scanner.Scan() {
    line := scanner.Text()
    key := line[:strings.Index( line, "=")]
    val := line[strings.Index( line, "=")+1:]
    propMap[key] = val
  }
}
