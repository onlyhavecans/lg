# Day One Logger

This is a port of [Brett Terpstra's ruby script](https://brettterpstra.com/2012/01/16/logging-with-day-one-geek-style/) to golang.

I wanted to have a version where I did not necessitate gem handling. This also moves to the newer DayOne CLI.

Most all information is a direct copy of the blog post above. 

This script works with the Day One command line utility
It parses an input string for a [date string] at the
beginning to parse natural language dates

Example usage:
dayonelogger "This is a entry."
dayonelogger "[yesterday 3pm] Something I did yesterday at 3:00PM"
dayonelogger "[1 hour ago] something I did an hour ago"


I highly recommend you rename the binary if you like to something short like "lg" or script this into alfred using the following applescript

```applescript
on alfred_script(q)
    do shell script "~/bin/dayonelogger \"" & q & "\""
end alfred_script
```
