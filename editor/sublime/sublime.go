package sublime

import (
	"fmt"
	"github.com/quarnster/completion/editor"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

var verbose = true

func init() {
	editor.Register(&Sublime{})
}

type Sublime struct{}

func (s *Sublime) Install() error {
	var (
		st_paths []string
		files    = []string{
			"3rdparty/jsonrpc.py",
			"editor/sublime/plugin.py",
		}
	)

	if u, err := user.Current(); err != nil {
		return err
	} else {
		switch runtime.GOOS {
		case "darwin":
			st_paths = append(st_paths, filepath.Join(u.HomeDir, "Library", "Application Support", "Sublime Text 2", "Packages"))
			st_paths = append(st_paths, filepath.Join(u.HomeDir, "Library", "Application Support", "Sublime Text 3", "Packages"))
		case "linux":
			st_paths = append(st_paths, filepath.Join(u.HomeDir, ".config", "sublime-text-2", "Packages"))
			st_paths = append(st_paths, filepath.Join(u.HomeDir, ".config", "sublime-text-3", "Packages"))
		default:
			return fmt.Errorf("Don't know where to install Sublime Text files on OS %s", runtime.GOOS)
		}
	}
	for i := range st_paths {
		if fi, err := os.Stat(st_paths[i]); err != nil || !fi.IsDir() {
			continue
		} else {
			p := filepath.Join(st_paths[i], "completion")
			os.Mkdir(p, 0755)
			for _, f := range files {
				if err := editor.Copy(f, filepath.Join(p, filepath.Base(f))); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (s *Sublime) Name() string {
	return "st"
}

func (s *Sublime) Description() string {
	return "Sublime Text 2 and Sublime Text 3"
}

func (s *Sublime) Uninstall() error {
	return fmt.Errorf("Sublime.Uninstall not implemented")
}
