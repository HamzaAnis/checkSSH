# Check SSH

This checks in parallel if the ssh session can be possible on the list of hosts.

Input is yaml file in format

```
ip:
  - 192.168.1.X
  - 192.168.2.X
  - 192.168.1.X
```

The output is the `sucess.log` and `error.log`.

### Import

```
go get github.com/HamzaAnis/checkSSH
```

### Example

```
package main

import "github.com/HamzaAnis/checkSSH"

func main() {
	checkSSH.Perform("ip.yaml",10)
}
```

Here 10 go routines will run and check the sessions in parallel and append to the log files.
