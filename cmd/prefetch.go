package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	cli "github.com/jawher/mow.cli"
	"www.velocidex.com/golang/go-prefetch"
)

func FatalIfError(err error, format string, args ...interface{}) {
	if err != nil {
		format += ": %v\n"
		args = append(args, err)
		fmt.Printf(format, args...)
		os.Exit(1)
	}
}

func main() {
	// create an app
	app := cli.App("prefetch", "parse pf files")

	// Specify the action to execute when the app is invoked correctly
	app.Command("parse", "Parse pf files", func(cmd *cli.Cmd) {
		file := cmd.StringsArg("PF_FILE", nil, "Prefix .pf files")
		cmd.Spec = "PF_FILE..."

		cmd.Action = func() {
			for _, arg := range *file {
				fd, err := os.Open(arg)
				FatalIfError(err, fmt.Sprintf("Open file: %v", err))

				prefetch_obj, err := prefetch.LoadPrefetch(fd)
				FatalIfError(err, fmt.Sprintf("Parsing Error: %v", err))

				serialized_content, _ := json.MarshalIndent(prefetch_obj, " ", " ")
				fmt.Println(string(serialized_content))
			}
		}
	})

	app.Command("extract", "Export a compressed pf file", func(cmd *cli.Cmd) {
		file := cmd.StringArg("PF_FILE", "", "Prefix .pf files")

		cmd.Action = func() {
			profile := prefetch.NewPrefetchProfile()
			reader, err := os.Open(*file)
			FatalIfError(err, fmt.Sprintf("OpenFile: %v", err))

			header := profile.MAMHeader(reader, 0)
			if header.Signature() == "MAM\x04" {
				// Need to decompress it in memory.
				data := make([]byte, header.UncompressedSize())
				n, err := reader.ReadAt(data, int64(header.Size()))
				if err != io.EOF {
					FatalIfError(err, fmt.Sprintf("Open file: %v", err))
				}

				decompressed, err := prefetch.LZXpressHuffmanDecompressWithFallback(
					data[:n], int(header.UncompressedSize()))
				FatalIfError(err, fmt.Sprintf("Open file: %v", err))

				os.Stdout.Write(decompressed)
			}
		}
	})

	// Invoke the app passing in os.Args
	app.Run(os.Args)
}
