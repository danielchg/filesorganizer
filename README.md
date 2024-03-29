# File Organizer

![CI Status](https://github.com/danielchg/filesorganizer/actions/workflows/goreleaser.yaml/badge.svg)
![CI Status](https://github.com/danielchg/filesorganizer/actions/workflows/tests.yaml/badge.svg)

This project is a CLI that organizes the files saved within a folder without any directory hierarchy, then copies them to a different folder where will be created subdirectories using the `year/month` notation, this info is obtained from the metadata of each file in the source folder.

## Usage

This is a CLI expect the following arguments:

* **src** This is the path where there are the files to be orgnized
* **dest** This is the path where the files will be orgnized using a subdirectories strucutre of `year/month`
* **w** Number of workers to parallelize file processing

Example:

```bash
./fileorganizer --src ./samples --dest ./test
```

All the files should be organized as below:

```bash
$ tree ./test
./test
├── 2019
│   ├── August
│   │   └── 2019-08-18 18.14.01-1-1.jpg
│   ├── February
│   │   └── 2019-02-06 15.21.06.jpg
│   ├── January
│   │   ├── 2019-01-27 06.08.07-1-1.jpg
│   │   └── 2019-01-27 14.03.37.jpg
│   └── March
│       └── 2019-03-28 22.16.18-1.jpg
└── 2022
    └── September
        ├── 2022-07-25 13.42.15.jpg
        └── 2022-08-03 12.03.09.jpg

7 directories, 7 files
```

![terminal](./assets/terminal.gif)

## Development

In order to build it just run `make` from the root folder of this repo.

Also there is a `.vscode/launch.json` file with the config to debug it using VSCode.

## TODO

* Use Cobra framework instead of just `flags`.
