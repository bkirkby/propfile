package propfile

import (
  "testing"
  "gopkg.in/check.v1"
  "fmt"
  "os"
)

func Test(t *testing.T) { check.TestingT(t) }

type PropFileSuite struct {
}

var _ = check.Suite(&PropFileSuite{})

func getTestFile()(string) {
  return os.TempDir()+"/propfile.test"
}

func (s *PropFileSuite) SetUpSuite( c *check.C) {
  f, err := os.OpenFile( getTestFile(), os.O_RDWR | os.O_CREATE, 0666)
  if err != nil {
    panic( fmt.Sprintf("unable to create file: %s\n", getTestFile()))
  }
  defer f.Close()

  _,err = f.WriteString("testKey=testVal=\n")
  if err !=nil {
    fmt.Println( err.Error())
  }
}

func (s *PropFileSuite) TearDownSuite( c *check.C) {
  os.Remove( getTestFile())
}

func (s *PropFileSuite) TestProps( c *check.C) {
  propMap := make(map[string]string)
  testFile := getTestFile()
  ReadFileInto( propMap, testFile)
  c.Assert( propMap["testKey"], check.Equals, "testVal=")
}
