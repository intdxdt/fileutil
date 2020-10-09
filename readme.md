# FileUtil
File Utilities in Go.

## API
 - Create multiple directories  : `MakeDirs(directory string) error`
 - Read all file                : `ReadAllOfFile(fileName string) (string, error)`
 - Save text to file            : `SaveText(destination, data string) error`
 - Copy file                    : `CopyFile(destination, source string) error`
 - Directory Tree               : `Tree(directory string, ignoreDirectories []string) ([]string, error)`
 - File Stat                    : `NewStat(fname string) (*Stat, error)`
