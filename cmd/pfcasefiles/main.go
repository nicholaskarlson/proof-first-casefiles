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
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *kitDir == "" || *outDir == "" {
		return 2
	}

	if err := ResetOutDir(*outDir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}

	out, err := ProcessKit(*kitDir)
	if err != nil {
		if werr := WriteText(filepath.Join(*outDir, "error.txt"), err.Error()+"\n"); werr != nil {
			fmt.Fprintln(os.Stderr, werr)
			return 1
		}
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
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *kitDir == "" || *outDir == "" {
		return 2
	}

	if err := ResetOutDir(*outDir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}

	out, err := ProcessKit(*kitDir)
	if err != nil {
		if werr := WriteText(filepath.Join(*outDir, "error.txt"), err.Error()+"\n"); werr != nil {
			fmt.Fprintln(os.Stderr, werr)
			return 1
		}
		return 0
	}

	rep := VerifyReport{
		OK:         true,
		KitTitle:   out.Index.Title,
		FileCount:  len(out.Index.Files),
		TotalBytes: out.Index.TotalBytes,
	}

	b, err := json.MarshalIndent(rep, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	if err := WriteText(filepath.Join(*outDir, "verify_report.json"), string(b)+"\n"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}

func cmdDemo(args []string) int {
	fs := flag.NewFlagSet("demo", flag.ContinueOnError)
	outBase := fs.String("out", "", "output directory")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *outBase == "" {
		return 2
	}

	if err := ResetOutDir(*outBase); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}

	cases, err := ListCases("fixtures/input")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	for _, c := range cases {
		inKit := filepath.Join("fixtures/input", c, "kit")
		expDir := filepath.Join("fixtures/expected", c)
		outDir := filepath.Join(*outBase, c)
		if err := ResetOutDir(outDir); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
		}

		out, err := ProcessKit(inKit)
		if err != nil {
			if werr := WriteText(filepath.Join(outDir, "error.txt"), err.Error()+"\n"); werr != nil {
				fmt.Fprintln(os.Stderr, werr)
				return 1
			}
		} else {
			if werr := WriteOutputs(outDir, out); werr != nil {
				fmt.Fprintln(os.Stderr, werr)
				return 1
			}
		}

		if err := DiffTrees(expDir, outDir); err != nil {
			fmt.Fprintf(os.Stderr, "demo mismatch %s: %v\n", c, err)
			return 1
		}
	}

	fmt.Println("OK")
	return 0
}
