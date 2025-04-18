package commands

import (
	"context"
	"fmt"
	"kwdb/app/backup"
	"kwdb/internal/helper/flogger"
	"time"
)

const CommandRestore = "RESTORE"

type RestoreCommand struct {
	name       string
	Args       *arguments
	isWritable bool
}

func NewRestoreCommand() *RestoreCommand {
	return &RestoreCommand{
		name:       CommandRestore,
		Args:       new(arguments),
		isWritable: false,
	}
}
func (c *RestoreCommand) CheckArgs() bool {
	return true
}

func (c *RestoreCommand) Execute() ([]byte, error) {

	ctx := context.Background()
	rc, res := backup.Backup(ctx)

	r := 0

	go func() {
		tl := time.NewTimer(3 * time.Second)
		defer tl.Stop()

		for {
			select {
			case msg, ok := <-rc:
				if ok == false {
					break
				}
				_, _ = SetAndRun(msg)
				r++
				continue
			case <-tl.C:
				flogger.Flogger.WriteString("закрыто по таймауту")
				return
			}

			break
		}

		select {
		case code := <-res:
			fmt.Printf("code - %v", code)
			switch code {
			case backup.BACKUP_END_CTX:
				flogger.Flogger.WriteString("закрыто по контексту")
			case backup.BACKUP_END_TIME:
				flogger.Flogger.WriteString("закрыто по таймауту")
			}
		default:
			break
		}
	}()

	return []byte("запущено восстановление..."), nil
}

func (c *RestoreCommand) Name() string {
	return c.name
}

func (c *RestoreCommand) SetArgs(args *arguments) {
	c.Args = args
}

func (c *RestoreCommand) IsWritable() bool {
	return c.isWritable
}
