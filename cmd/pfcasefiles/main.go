package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func usage() {
	fmt.Fprintln(os.Stderr, "pfcasefiles demo|render|verify ...")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	switch os.Args[1] {
	case "demo":
		os.Exit(cmdDemo(os.Args[2:]))
	case "render":
		os.Exit(cmdRender(os.Args[2:]))
	case "verify":
		os.Exit(cmdVerify(os.Args[2:]))
	default:
		usage()
		os.Exit(2)
	}
}

func cmdRender(args []string) int {
	fs := flag.NewFlagSet("render", flag.ContinueOnError)
	kitDir := fs.String("kit", "", "kit directory")
	outDir := fs.String("out", "", "output directory")
	_ = fs.Parse(args)
	if *kitDir == "" || *outDir == "" {
		return 2
	}

	_ = ResetOutDir(*outDir)

	out, err := ProcessKit(*kitDir)
	if err != nil {
		_ = WriteText(filepath.Join(*outDir, "error.txt"), err.Error()+"\n")
		return 0
	}

	if err := WriteOutputs(*outDir, out); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}

func cmdVerify(args []string) int {
	fs := flag.NewFlagSet("verify", flag.ContinueOnError)
	kitDir := fs.String("kit", "", "kit directory")
	outDir := fs.String("out", "", "output directory")
	_ = fs.Parse(args)
	if *kitDir == "" || *outDir == "" {
		return 2
	}

	_ = ResetOutDir(*outDir)

	out, err := ProcessKit(*kitDir)
	if err != nil {
		_ = WriteText(filepath.Join(*outDir, "error.txt"), err.Error()+"\n")
		return 0
	}

	rep := VerifyReport{
		OK:         true,
		KitTitle:   out.Index.Title,
		FileCount:  len(out.Index.Files),
		TotalBytes: out.Index.TotalBytes,
	}

	b, _ := json.MarshalIndent(rep, "", "  ")
	_ = WriteText(filepath.Join(*outDir, "verify_report.json"), string(b)+"\n")
	return 0
}

func cmdDemo(args []string) int {
	fs := flag.NewFlagSet("demo", flag.ContinueOnError)
	outBase := fs.String("out", "", "output directory")
	_ = fs.Parse(args)
	if *outBase == "" {
		return 2
	}

	_ = ResetOutDir(*outBase)

	cases, err := ListCases("fixtures/input")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	for _, c := range cases {
		inKit := filepath.Join("fixtures/input", c, "kit")
		expDir := filepath.Join("fixtures/expected", c)
		outDir := filepath.Join(*outBase, c)
		_ = ResetOutDir(outDir)

		out, err := ProcessKit(inKit)
		if err != nil {
			_ = WriteText(filepath.Join(outDir, "error.txt"), err.Error()+"\n")
		} else {
			_ = WriteOutputs(outDir, out)
		}

		if err := DiffTrees(expDir, outDir); err != nil {
			fmt.Fprintf(os.Stderr, "demo mismatch %s: %v\n", c, err)
			return 1
		}
	}

	fmt.Println("OK")
	return 0
}
