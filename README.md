# math-skills

math-skills is a Go program and a solution to an exercise of the same name in the 01-edu curriculum.

## usage

The program reads numbers from a provided text file, line by line, and prints statistical values describing the data set.

The program can be used without compiling it:

```bash
go run . "data.txt"
```

Compiling it to an executable is another option:

```bash
go build
./mathskills data.txt
```

## Docker testing

Make sure Docker is installed and the provided .zip -file is unpacked.

Navigate to the stat-bin folder and run this command:
```bash
./run.sh math-skills
```

The first time, it will build the application, make a data.txt -file and produce an output. On subsequent runs, it will write data.txt over and produce the output.

Navigate to the project folder (math-skills/) and make an executable:
```bash
go build .
```

Move the executable to the stat-bin folder. Now compare the outputs of the two programs:
```bash
./run.sh math-skills
./mathskills data.txt
```

They should be the same.