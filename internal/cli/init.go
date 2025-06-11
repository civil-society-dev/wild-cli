package cli

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

//go:embed scaffold scaffold/.gitignore scaffold/.env.example scaffold/.wildcloud
var scaffoldFS embed.FS

const wildcloudDir = ".wildcloud"

func runInit(args []string) error {
	var update bool
	
	for _, arg := range args {
		switch arg {
		case "--update":
			update = true
		case "-h", "--help":
			showInitHelp()
			return nil
		default:
			return fmt.Errorf("unknown flag: %s", arg)
		}
	}
	
	if !update {
		if err := checkEmptyDirectory(); err != nil {
			return err
		}
	}
	
	fmt.Println("Initializing Wild-Cloud project in", getCurrentDir())
	
	if update {
		fmt.Println("Updating scaffold files (preserving existing non-scaffold files)")
	} else {
		fmt.Println("Creating .wildcloud directory and scaffold files")
	}
	
	// Note: .wildcloud directory will be created by the scaffold files
	
	if err := createScaffoldFiles(update); err != nil {
		return fmt.Errorf("failed to create scaffold files: %v", err)
	}
	
	fmt.Println("")
	fmt.Println("Wild-Cloud project initialized successfully!")
	fmt.Println("")
	fmt.Println("Next steps:")
	fmt.Println("1. Review and customize the configuration files")
	fmt.Println("2. Set up your .wildcloud/config.yaml with your Wild-Cloud repository path")
	fmt.Println("3. Start using wild-app-* commands to manage your applications")
	
	return nil
}

func showInitHelp() {
	fmt.Println("Usage: wild init [--update]")
	fmt.Println("")
	fmt.Println("Initialize a new Wild-Cloud project by creating .wildcloud directory and scaffold files.")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  --update    Update existing files with scaffold files (overwrite)")
	fmt.Println("  -h, --help  Show this help message")
	fmt.Println("")
	fmt.Println("By default, this command will only run in an empty directory.")
	fmt.Println("Use --update to overwrite existing scaffold files while preserving other files.")
}

func checkEmptyDirectory() error {
	entries, err := os.ReadDir(".")
	if err != nil {
		return fmt.Errorf("failed to read current directory: %v", err)
	}
	
	if len(entries) > 0 {
		return fmt.Errorf("current directory is not empty\nUse --update flag to overwrite existing scaffold files while preserving other files")
	}
	
	return nil
}

func getCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return "."
	}
	return dir
}

func createWildcloudDir() error {
	return os.MkdirAll(wildcloudDir, 0755)
}

func createScaffoldFiles(update bool) error {
	return fs.WalkDir(scaffoldFS, "scaffold", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		
		// Skip the root scaffold directory itself
		if path == "scaffold" {
			return nil
		}
		
		// Get relative path from scaffold root
		relPath := strings.TrimPrefix(path, "scaffold/")
		
		// Files go directly to current directory, not inside .wildcloud
		destPath := relPath
		
		if d.IsDir() {
			// Create directory
			return os.MkdirAll(destPath, 0755)
		}
		
		// Handle file
		if !update {
			if fileExists(destPath) {
				fmt.Printf("Skipping existing file: %s\n", relPath)
				return nil
			}
		}
		
		if update && fileExists(destPath) {
			fmt.Printf("Updating: %s\n", relPath)
		} else {
			fmt.Printf("Creating: %s\n", relPath)
		}
		
		// Read file from embedded filesystem
		content, err := scaffoldFS.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read embedded file %s: %v", path, err)
		}
		
		// Ensure destination directory exists
		if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
			return fmt.Errorf("failed to create directory for %s: %v", relPath, err)
		}
		
		// Write file
		if err := os.WriteFile(destPath, content, 0644); err != nil {
			return fmt.Errorf("failed to create %s: %v", relPath, err)
		}
		
		return nil
	})
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}