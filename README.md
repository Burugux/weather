[![Build Status](https://travis-ci.org/Burugux/weather.svg?branch=master)](https://travis-ci.org/Burugux/weather)
# Weather

Get today's forecast from the command line!

## Pre-requisite
You need to create a dark sky secret key, do that [here.](https://darksky.net/dev)
Then add it to your `.bashrc` or `.profile`.

```console
export API_KEY=<YOUR_SECRET_KEY>
```

## Installation
```console
$ go get github.com/Burugux/weather
```

## Usage
```console
$ weather
Current location: Africa/Nairobi
Today's forecast:Mostly Cloudy
Current temp:19.56°C
Feels like:19.56°C
```
