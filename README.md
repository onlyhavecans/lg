# lg (Log to DayOne)

This is a port of [Brett Terpstra's ruby script](https://brettterpstra.com/2012/01/16/logging-with-day-one-geek-style/) to golang.

I wanted to have a version where I did not necessitate gem handling. This also moves to the newer DayOne CLI.

This script works with the Day One command line utility;
It parses an input string for a [date string] at the
beginning to parse natural language dates

## Install
Requires golang to build!
```
brew install golang
go get onlyhavecans.works/amy/lg
```

If you would like to start posting pre-built binaries tweet me at @onlyhavecans

## Example usage:

```shell script
lg "This is a entry."
lg "[yesterday 3pm] Something I did yesterday at 3:00PM"
lg "[1 hour ago] something I did an hour ago"
```

I primarily use this as an alfred extension using the following applescript. You may need to adjust your path to the script

```applescript
on alfred_script(q)
    do shell script "~/go/bin/lg \"" & q & "\""
end alfred_script
```
