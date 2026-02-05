package main

import (
	"Argus/pkg"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func snap_handler(_ *cobra.Command, args []string) {
	start_time := time.Now()

	target_dir := args[0]

	hash_storage := pkg.Scan_folder(target_dir)
	str := ""
	for k, v := range hash_storage.Get_files() {
		str += k + " " + v + "\n"
	}
	file_name := target_dir + "argus_snapshot.ags"
	err := os.WriteFile(file_name, []byte(str), 0644)
	if err != nil {
		fmt.Println(err)
	}

	println("Argus is done in", time.Now().Sub(start_time), "ms")
	os.Exit(0)
}

func compare_handler(_ *cobra.Command, args []string) {
	start_time := time.Now()

	current_storage := pkg.Scan_folder(args[0])
	current_map := current_storage.Get_files()

	snapshot_bytes, _ := os.ReadFile(args[1])
	snapshot_map := make(map[string]string)

	lines := strings.Split(string(snapshot_bytes), "\n")
	for _, line := range lines {
		f := strings.Fields(line)
		if len(f) == 2 {
			snapshot_map[f[0]] = f[1]
		}
	}

	println("+--------------+")

	for path, current_hash := range current_map {

		file_info, err := os.Stat(path)
		if err != nil || file_info.IsDir() {
			continue
		}

		if path == args[1] {
			continue
		}

		if old_hash, exists := snapshot_map[path]; exists {
			if old_hash != current_hash {
				fmt.Printf("File tampered : %s\n", path)
			}
		} else {
			fmt.Printf("New file : %s\n", path)
		}
	}
	println("+--------------+")
	println("Argus is done in", time.Now().Sub(start_time), "ms")
	os.Exit(0)
}

func main() {
	var argus_cmd = &cobra.Command{
		Use:   "argus",
		Short: "Argus checks file hashes to ensure integrity",
	}

	var snap_cmd = &cobra.Command{
		Use:   "snap [folder path]",
		Short: "Snapshot the hashes of a folder",
		Args:  cobra.ExactArgs(1),
		Run:   snap_handler,
	}

	var compare_cmd = &cobra.Command{
		Use:   "compare [folder path] [snapshot path]",
		Short: "Compare current files against the .ags snapshot",
		Args:  cobra.ExactArgs(2),
		Run:   compare_handler,
	}

	argus_cmd.AddCommand(snap_cmd)
	argus_cmd.AddCommand(compare_cmd)

	// If error when executing the command
	if err := argus_cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
