package cmd

import (
  "bufio"
  "testing"
  . "github.com/smartystreets/goconvey/convey"
  "strings"
)

func TestCommandExec(t *testing.T) {
  Convey("Command.Exec", t, func(){
    Convey("returns the streamed output", func(){
      cmd := Command("ls")

      out := cmd.Exec(nil)
      lines := bufio.NewReader(out)

      line, _ := lines.ReadString('\n')

      So(line,ShouldNotEqual,nil)
    })

    Convey("takes an input to the command", func(){
      cmd := Command("sed", "s/is/at/")

      input := strings.NewReader("this")

      out := cmd.Exec(input)
      lines := bufio.NewScanner(out)

      for lines.Scan() {
        // Scan for any lines
      }

      line := lines.Text()

      So(line,ShouldEqual,"that")
    })

    Convey("returns the error with the final read", func(){
      cmd := Command("false")

      reader := cmd.Exec(nil)

      var out []byte
      _, err := reader.Read(out)

      So(err, ShouldNotEqual, nil)
    })
  })
}

// Exec creates a command and executices it
// Exec can chain commands: ls . | grep _test

// Command.Exec takes an input to provide to the command
// Command.Exec returns the streamed output as a Reader
