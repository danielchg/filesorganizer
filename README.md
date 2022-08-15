# File Organizer

This project is a CLI that organize the files within a folder copying then to a diferent folder where will be create subdirectories using the `year/month` notation, this info is get from the metadata of each file in the source folder.

## Usage

This is a CLI that expect just two arguments:

* **src** This is the path where there are the files to be orgnized
* **dest** This is the path where the files will be orgnized using a subdirectories strucutre of `year/month`

Example:

```bash
./fileorganizer --src ./samples --dest ./test
```

This will generate the following folder organization where will be copied each file:

```bash
./test
└── 2022
    ├── August
    └── July
```

## Development

In order to build it just run `make build` from the root folder of this repo.

Also there are the `.vscode` foler with the launcher to debug it from VSCode.

## TODO

* Add unit tests.
* Add Go routines to improve the performance.
* Use Cobra framework instead of just `flags`.
* Create GitHub Action with GoReleaser to build from a pipeline a new version.
