package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	cli "github.com/jawher/mow.cli"
	"www.velocidex.com/golang/binparsergen"
	"www.velocidex.com/golang/go-prefetch"
)

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
				binparsergen.FatalIfError(err, fmt.Sprintf("Open file: %v", err))

				prefetch_obj, err := prefetch.LoadPrefetch(fd)
				binparsergen.FatalIfError(err, fmt.Sprintf("Parsing Error: %v", err))

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
			binparsergen.FatalIfError(err, fmt.Sprintf("OpenFile: %v", err))

			header := profile.MAMHeader(reader, 0)
			if header.Signature() == "MAM\x04" {
				// Need to decompress it in memory.
				data := make([]byte, header.UncompressedSize())
				n, err := reader.ReadAt(data, int64(header.Size()))
				if err != io.EOF {
					binparsergen.FatalIfError(err, fmt.Sprintf("Open file: %v", err))
				}

				decompressed, err := prefetch.LZXpressHuffmanDecompress(
					data[:n], int(header.UncompressedSize()))
				binparsergen.FatalIfError(err, fmt.Sprintf("Open file: %v", err))

				os.Stdout.Write(decompressed)
			}
		}
	})

	// Invoke the app passing in os.Args
	app.Run(os.Args)
}
