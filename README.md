[![Build Status](https://travis-ci.com/edualb/godmitri.svg?branch=master)](https://travis-ci.com/edualb/godmitri)

<div align="center">
  <img src="https://user-images.githubusercontent.com/39157101/89476373-e52d3980-d760-11ea-99b4-d35b675c3a8d.png" width="200">
</div>
<div align="center">
  Icon made by <a href="http://www.freepik.com/" title="Freepik">Freepik</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a>
</div>
<br>

#### Table Of Contents:

[Introduction](#introduction)

[Installation](#installation)

[Getting start](#getting-start)

## **Introduction**

This project provides utilities for general chemistry. The name of the project was inspired by periodic table creator Dmitri Ivanovic Mendeleev.

Key features:

  * Get all elements properties from the periodic table :heavy_check_mark:

If you want to contribute with a new feature or found out some bug which you can solve, check our [contributing guideline](https://github.com/edualb/godmitri/blob/master/CONTRIBUTING.md)

## **Installation**

* init the go module in your project

  ```
  $ go mod init example.com/myProject
  ```

* Just import the `godmitri` library in your project

  ```go
  import "github.com/edualb/godmitri/element"
  ```

## **Getting start**

```go
package main

import (
	"fmt"

	"github.com/edualb/godmitri/element"
)

func main() {
    // Creates an instance of Aluminium
	a := element.Aluminium{}

    // Prints the atomic weight from aluminium
	fmt.Printf("%f\n", a.GetAtomicWeight())
}
```
