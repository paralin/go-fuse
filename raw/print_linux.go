package raw
import (
	"fmt"
	"syscall"
)

func init() {
	OpenFlagNames[syscall.O_DIRECT] = "DIRECT"
	OpenFlagNames[syscall.O_LARGEFILE] = "LARGEFILE"
	OpenFlagNames[syscall_O_NOATIME] = "NOATIME"

	initFlagNames[CAP_XTIMES] = "XTIMES"
	initFlagNames[CAP_VOL_RENAME] = "VOL_RENAME"
	initFlagNames[CAP_CASE_INSENSITIVE] = "CASE_INSENSITIVE"
}

func (a *Attr) String() string {
	return fmt.Sprintf(
		"{M0%o SZ=%d L=%d "+
			"%d:%d "+
			"%d*%d %d:%d "+
			"A %d.%09d "+
			"M %d.%09d "+
			"C %d.%09d}",
		a.Mode, a.Size, a.Nlink,
		a.Uid, a.Gid,
		a.Blocks, a.Blksize,
		a.Rdev, a.Ino, a.Atime, a.Atimensec, a.Mtime, a.Mtimensec,
		a.Ctime, a.Ctimensec)
}

func (me *GetAttrIn) String() string {
	return fmt.Sprintf("{Fh %d}", me.Fh_)
}

func (me *ReadIn) String() string {
	return fmt.Sprintf("{Fh %d off %d sz %d %s L %d %s}",
		me.Fh, me.Offset, me.Size,
		FlagString(readFlagNames, int(me.ReadFlags), ""),
		me.LockOwner,
		FlagString(OpenFlagNames, int(me.Flags), "RDONLY"))
}

func (me *WriteIn) String() string {
	return fmt.Sprintf("{Fh %d off %d sz %d %s L %d %s}",
		me.Fh, me.Offset, me.Size,
		FlagString(writeFlagNames, int(me.WriteFlags), ""),
		me.LockOwner,
		FlagString(OpenFlagNames, int(me.Flags), "RDONLY"))
}