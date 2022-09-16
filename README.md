# File Organizer

This project is a CLI that organize the files saved within a folder without any directory hierarcy,then copying them to a diferent folder where will be created subdirectories using the `year/month` notation, this info is get from the metadata of each file in the source folder.

## Usage

This is a CLI expect the following arguments:

* **src** This is the path where there are the files to be orgnized
* **dest** This is the path where the files will be orgnized using a subdirectories strucutre of `year/month`
* **w** Number of workers to parallelize file processing

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

In order to build it just run `make` from the root folder of this repo.

Also there is a `.vscode/launch.json` file with the config to debug it using VSCode.

## TODO

* Add unit tests.
* Use Cobra framework instead of just `flags`.
