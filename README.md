# split-file

Can be used to split a big file in multiple parts to make transport easier.

## Usage
### Split files
Pass the filename with the -f flag and the wanted part size in Mb,Kb or bytes. Only 1 unit is necessary.
Start the program using:
```
split-file -f {filename} -gb {size in Gb} -mb {size in Mb} -kb {size in Kb} -b {size in b}
```

### Merge files
Pass the filename with the -f flag and use the -m flag to merge.
Start the program using:
```
split-file -f {filename} -m
```

### Clean files
Removes the partial files.
```
split-file -f {filename} -c
```

## Help
```
split-file -h
  -b int
        Size in byte
  -c    Clean all file parts
  -f string
        Filename (without .number when merging)
  -gb int
        Size in Gb
  -kb int
        Size in Kb
  -m    Merge parts
  -mb int
        Size in Mb
```