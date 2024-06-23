# onemonth

A program that iterate the day of the month. I created it to generate the date text for the diary I am generating every month.

- [Installation](#installation)
- [Usage](#usage)
    - [Output the date of the current month](#output-the-date-of-the-current-month)
    - [Output the date of the specified month](#output-the-date-of-the-specified-month)
    - [Output the date of the current month of the specified year](#output-the-date-of-the-current-month-of-the-specified-year)
    - [Specify output format](#specify-output-format)
    - [Load configuration file](#load-configuration-file)
- [Use in Golang code](#use-in-golang-code)

## Installation

```sh
go install github.com/yomashishi/onemonth/cmd@latest
```

## Usage

### Output the date of the current month

```sh
onemonth
```

Output:

Partially omitted.

```txt
2022-11-01
2022-11-02
2022-11-03
2022-11-04
2022-11-05
...
2022-11-26
2022-11-27
2022-11-28
2022-11-29
2022-11-30
```

### Output the date of the specified month

```sh
onemonth -month 1
# Shorthand
onemonth -m 1
```

### Output the date of the current month of the specified year

```sh
onemonth -year 2023
# Shorthand
onemonth -y 2023
```

### Specify output format

Conform to Golang time format.

see: <https://pkg.go.dev/time#pkg-constants>

```sh
onemonth -layout 'Jan _2'
# shorthand
onemonth -l 'Jan _2'
```

Output:

Partially omitted.

```txt
Nov  1
Nov  2
Nov  3
Nov  4
Nov  5
...
Nov 26
Nov 27
Nov 28
Nov 29
Nov 30
```

### Load configuration file

Read the following files saved in `$HOME/.config/onemonth` directory.
Read only the first line.

- `layout.time`:
    - [Specify output format](#specify-output-format)
- `layout.output`:
    - Example when you want to insert two line breaks for each month:

    ```txt
    %s\n\n
    ```

## Use in Golang code

```sh
go get github.com/yomashishi/onemonth
```

```golang
func main() {
	month, err := onemonth.NewOneMonth(2022, 1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	month.Iterate(func(day time.Time) {
		// some operation
	})
}
```
