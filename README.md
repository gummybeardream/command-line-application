# Housing CSV to JSON CLI
Go command line application for outputting JSON lines files from csv files

## How to Use
To start, clone the repo and open the directory. To test the CLI tool, use

```bash
go run . CSV_FILE_PATH OUTPUT_FILE_PATH
```

The input file should be .csv. The output file will be .jl or .json. 

To compile the code into an executable binary, use the command:
```bash
go build -o command-line-application .
```

and run it like:
```bash
./command-line-application CSV_FILE_PATH OUTPUT_FILE_PATH
```

## Examples
There needs to be 7 fields for the 7 columns in the CSV file. For example: 

```csv
"value","income","age","rooms","bedrooms","pop","hh"
452600,8.3252,41,880,129,322,126
358500,8.3014,21,7099,1106,2401,1138
...
```
```json
{
    "houses": [
        {
            "value": 452600,
            "income": 8.3252,
            "age": 41,
            "rooms": 880,
            "bedrooms": 129,
            "pop": 322,
            "hh": 126
        },
        {
            "value": 358500,
            "income": 8.3014,
            "age": 21,
            "rooms": 7099,
            "bedrooms": 1106,
            "pop": 2401,
            "hh": 1138
        },
    ]
    ...
}
```

## AI Disclosure
I used ChatGPT to organize the project and guide me through the logic without generating code. I used it to teach me syntax and look up Go fucntions, packages, and libraries. 