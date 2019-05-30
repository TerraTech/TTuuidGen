[![GoDoc](https://godoc.org/github.com/TerraTech/TTuuidGen/pkg/TTuuid?status.svg)](https://godoc.org/github.com/TerraTech/TTuuidGen/pkg/TTuuid)
[![Build Status](https://travis-ci.org/TerraTech/TTuuidGen.svg?branch=master)](https://travis-ci.org/TerraTech/TTuuidGen)
[![Go Report Card](https://goreportcard.com/badge/github.com/TerraTech/TTuuidGen)](https://goreportcard.com/report/github.com/TerraTech/TTuuidGen)

## TTuuidGen

```
$ TTuuidGen --help
usage: TTuuidGen [-hV] {[--s<1|4>] | [-H <int>] [length]}
  -H int          number of hyphens to add
      --s1        emit standard uuid V1 (time based)
      --s4        emit standard uuid V4 (random based)
  -V, --version   Show version
```
```sh
$ TTuuidGen
92e9e4e569782a1b1400e807957146407c34

$ TTuuidGen --s1
6b91066e-7fb5-11e9-b2b4-166e58793fd2

$ TTuuidGen --s4
59057e9d-3536-45d0-a8e2-1ec993ce27ab

$ TTuuidGen 100
c5cb851df9aa6b5bf349239cb23972a56f82d9da6f6c8f5b2dce3d22a9be7b93fe1f233c2101eae4a4e7a7210ac5f9d5961c

$ TTuuidGen -H 7 100
b8117e62ee57-2f1e8c9f5e01-1371c16dd202-de355ed6b4c8-8f3f217c1e0a-6fd487c189dc-c65fb8dae2c8-94f12e830
```
