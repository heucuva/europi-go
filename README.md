# EuroPiGo

[![Go Reference](https://pkg.go.dev/badge/github.com/awonak/EuroPiGo.svg)](https://pkg.go.dev/github.com/awonak/EuroPiGo)

Alternate firmware for the [Allen Synthesis EuroPi](https://github.com/Allen-Synthesis/EuroPi) written in Go.

## Getting started

Install Go

[https://go.dev/doc/install](https://go.dev/doc/install)

Install TinyGo

[https://tinygo.org/getting-started/install](https://tinygo.org/getting-started/install)

Install the TinyGo VSCode plugin

[https://marketplace.visualstudio.com/items?itemName=tinygo.vscode-tinygo](https://marketplace.visualstudio.com/items?itemName=tinygo.vscode-tinygo)

## Build the example

Use the `tinygo flash` command while your pico is in BOOTSEL mode to compile the script and copy it to your attached EuroPi pico.

```shell
tinygo flash --target pico examples/diagnostics.go
```

After the first time you flash your TinyGo program you will no longer need to reboot in BOOTSEL mode to flash your script.

> **_NOTE:_** If your script throws a panic, you will need to reflash using BOOTSEL mode.

## Serial printing

When your EuroPi pico is connected via USB you can view printed serial output using a tool like `minicom`.

For example, a line in your code like:

```go
log.Printf("K1: %2.2f\n\r", e.K1.ReadVoltage())
```

You can launch minicom to view the printed output:

```shell
minicom -b 115200 -o -D /dev/ttyACM0
```

## VSCode build task

Add the TinyGo flash command as your default build task:

```plaintext
Ctrl + Shift + P > Tasks: Configure Default Build Task
```

Use the following example task configuration to set tinygo flash as your default build command:

```json
{
    "version": "2.0.0",
    "tasks": [
        {
            "label": "tinygo flash",
            "type": "shell",
            "command": "tinygo flash --target pico -size short -opt 1 ${workspaceRoot}/examples",
            "group": {
                "kind": "build",
                "isDefault": true
            },
        }
    ]
}
```

Now you can build and flash your project using `Ctrl + Shift + B` or search for the command:

```plaintext
Ctrl + Shift + P > Tasks: Run Build Task
```

## Debugging using picoprobe

TODO

## Debugging using a browser

**NOTE**: This is a beta feature! Please do not expect it to function properly, if at all.

If your EuroPi project is set up via the `europi.Bootstrap()` execution paradigm, you may operate and visualize it via a browser instead of having to flash it to a EuroPi hardware module or build it via TinyGo.

There are some caveats that must be understood before using this functionality:
1. This builds via Pure Go and only simulates (to the best of ability) a RP2040 / Pico running the software.
1. The visualization web page is terrible. While some care was taken to make it minimally usable, it's not pretty and is rife with bugs.
1. The websocket interface presented can currently only bind to port 8080 on "all" network interfaces.
1. Projects which do not utilize the `europi.Bootstrap()` execution paradigm will not be able to partake in this functionality.
1. [pprof](https://github.com/google/pprof/blob/main/README.md) support is enabled while the websocket interface is available, but it does not represent how the code will run on a RP2040 / Pico or within a EuroPi module. It is specifically provided here for debugging multiprocessing functionality and gross abuses of memory and CPU. 

### How to get the browser support to work

First, you need [Go installed](https://go.dev/doc/install). Preferably 1.18 or newer.

Second, set up your module project to build with the `europi.Bootstrap()` execution paradigm and add the following line to the `europi.Bootstrap()` parameter list:

```go
europi.AttachNonPicoWS(true),
```

Then, execute your module using Go (not TinyGo) and pass a buildflag tag parameter matching [which revision of hardware](hardware/README.md#revision-list) you want to run, like so:

```bash
go run -tags=revision1 internal/projects/randomskips/randomskips.go
```

Finally, open up a browser to [localhost:8080](http://localhost:8080/) to see your module in action!

## Why should I use the TinyGo version of the EuroPi firmware?

You probably shouldn't. This is my passion project because I love Go. You should probably be using the official [EuroPi firmware](https://github.com/Allen-Synthesis/EuroPi). But if you are interested in writing a dedicated script for the EuroPi that requires concurrency and faster performance than MicroPython can offer, then maybe this is the right firmware for you!
