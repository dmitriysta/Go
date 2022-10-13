package main

import "fmt"

type Hen interface {
  getCountOfEggsPerMonth() int
  getDescription() string

}
