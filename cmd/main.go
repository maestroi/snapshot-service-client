package main

import (
	"archive/tar"
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/klauspost/pgzip"
)

type SnapshotInfo struct {
	LastModified string `json:"last_modified"`
	Size         int64  `json:"size"`
	URL          string `json:"url"`
}

func main() {
	overwrite := flag.Bool("f", false, "Overwrite existing files without prompt")
	flag.Parse()

	args := flag.Args()
	if len(args) < 3 {
		fmt.Println("Please provide protocol, network, and target directory")
		os.Exit(1)
	}

	protocol := args[0]
	network := args[1]
	targetDir := args[2]

	start := time.Now()
	url := fmt.Sprintf("http://localhost:8080/api/v1/files/%s/%s/latest", protocol, network)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching snapshot:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	snapshotInfo := SnapshotInfo{}
	err = json.NewDecoder(resp.Body).Decode(&snapshotInfo)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		os.Exit(1)
	}

	fmt.Printf("Latest snapshot size: %.2f GB\n", float64(snapshotInfo.Size)/1e9)
	fmt.Println("Starting download... This may take a while.")

	resp, err = http.Get(snapshotInfo.URL)
	if err != nil {
		fmt.Println("Error downloading snapshot:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	gr, err := pgzip.NewReader(resp.Body)
	if err != nil {
		fmt.Println("Error initializing gzip reader:", err)
		os.Exit(1)
	}

	tr := tar.NewReader(gr)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			fmt.Println("Error reading tar:", err)
			os.Exit(1)
		}

		target := filepath.Join(targetDir, header.Name)

		switch header.Typeflag {
		case tar.TypeDir: // if a dir, create it
			fmt.Printf("Creating directory: %s\n", target)
			if err := os.MkdirAll(target, 0755); err != nil {
				fmt.Println("Error creating directory:", err)
				os.Exit(1)
			}

		case tar.TypeReg: // if a file, extract it
			if _, err := os.Stat(target); err == nil && !*overwrite {
				reader := bufio.NewReader(os.Stdin)
				fmt.Printf("File %s already exists. Do you want to overwrite it? [y/N]: ", target)
				response, _ := reader.ReadString('\n')
				response = strings.ToLower(strings.TrimSpace(response))
				if response != "y" && response != "yes" {
					fmt.Println("Skipping file on user request.")
					continue
				}
			}

			fmt.Printf("Extracting file: %s\n", target)
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				fmt.Println("Error extracting file:", err)
				os.Exit(1)
			}

			if _, err := io.Copy(f, tr); err != nil {
				f.Close()
				fmt.Println("Error extracting file:", err)
				os.Exit(1)
			}
			f.Close()
		}
	}

	fmt.Println("Snapshot extraction completed.")
	elapsed := time.Since(start)
	fmt.Printf("Snapshot extraction took %s\n", elapsed)
}
