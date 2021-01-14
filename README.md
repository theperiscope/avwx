# AVWX

## Getting Started

avwx Тhe repository contains two command-line tools for accessing the most recent raw METAR and TAF
information from the [Text Data Server](https://aviationweather.gov/dataserver) provided by the Aviation Weather Center.

### Prerequisites

* Go

### Build

* Clone this repository
* Run `build.cmd`

### Usage

* Single station

  ```text
  avwx metar KORD
  avwx taf KORD
  ```

* Multiple Stations

  ```text
  avwx metar KORD PHOG
  avwx taf KORD PHOG
  ```

* Partial Station Name

  ```text
  avwx metar PH*
  avwx taf PH*
  ```

* Entire State
  _use two-letter U.S. state or two-letter Canadian province_

  ```text
  avwx metar @il
  avwx taf @il
  ```

* Country
  _use two-letter country abbreviation_

  ```text
  avwx metar ~au
  avwx taf ~au
  ```

* Mix & Match

  ```text
  avwx metar KORD CY* ~au @hi
  avwx taf KORD CY* ~au @hi
  ```
