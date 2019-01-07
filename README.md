# gifstitch

This project is a simple program to combine multiple source GIF animation files into a single merged animation. A list of source files is provided, as well as an optional list of loop counters, and a master file is aggregated by appending each animation's frames `n` number of times where `n` is either 1 or the provided number of loops.

## Usage

To install, you can run:

```bash
go install github.com/haleyrc/gifstitch
```

Then in the directory with your source files, run:

```bash
gifstitch --files FILELIST [--loops LOOPLIST -o OUTPUTFILE]
```

where `FILELIST` is an ordered, comma-separated list of file names to append and `LOOPLIST` is an ordered, comma-separated list of loop counters for each file.

Note that if `--loops` is provided, the number of counters provided _MUST_ be equal to the number of source files provided.

`OUTPUTFILE` is the name of the file to create. If no output file is given, the default name of `merged.gif` will be used.

## TODO

- [ ] Better support for files outside of the working directory.
- [ ] Piped input support to help with scripting.
- [ ] A server implementation demo.
