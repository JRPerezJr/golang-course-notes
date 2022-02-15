<h1 align="center">Golang Course Notes</h1>

<br>

<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />
  <a href="https://github.com/JRPerezJr/golang-course-notes/blob/main/LICENSE" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

<br>

<br>

> Golang coursework and lab notes.

## 🛠 Installing Golang

[Download and install](https://go.dev/doc/install)

MacOS

1. Open the package file you downloaded and follow the prompts to install Go.

The package installs the Go distribution to /usr/local/go. The package should put the /usr/local/go/bin directory in your PATH environment variable. You may need to restart any open Terminal sessions for the change to take effect.

2.Verify that you've installed Go by opening a command prompt and typing the following command:

`$ go version`

3. Confirm that the command prints the installed version of Go.


Linux

1. Extract the archive you downloaded into /usr/local, creating a Go tree in /usr/local/go.

    Important: This step will remove a previous installation at /usr/local/go, if any, prior to extracting. Please back up any data before proceeding.

    For example, run the following as root or through sudo:

    `rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.7.linux-amd64.tar.gz`

2. Add `/usr/local/go/bin` to the PATH environment variable.

    You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):

    `export PATH=$PATH:/usr/local/go/bin`

    Note: Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.
 
3. Verify that you've installed Go by opening a command prompt and typing the following command:

    `$ go version`

4. Confirm that the command prints the installed version of Go.



Windows


1. Open the MSI file you downloaded and follow the prompts to install Go.

    By default, the installer will install Go to Program Files or Program Files (x86). You can change the location as needed. After installing, you will need to close and reopen any open command prompts so that changes to the environment made by the installer are reflected at the command prompt.
    
2. Verify that you've installed Go.
   1. In Windows, click the Start menu.
   2. In the menu's search box, type cmd, then press the Enter key.
   3. In the Command Prompt window that appears, type the following command:

    `$ go version`

3. Confirm that the command prints the installed version of Go.



## 👩‍💻 👨‍💻 Basic Usage 
Print current Golang version

```shell
go version
```

Compile and run Go program

```shell
go run <fileName>
```

Build the project and emit an executable binary

```shell
go build
```

Check for concurrency problems

```shell
go build -race
```

Manage modules & dependencies

```shell
go mod
```

Update dependencies

```shell
go mod tidy
```

Run project's test suite

```shell
go test
```

Format all source files

```shell
go fmt
```



## 📓 Author

![Logo](https://user-images.githubusercontent.com/19915910/120965966-81203b00-c7a0-11eb-8ef4-a42c0642db4c.png)

- Github: [@JRPerezJr](https://github.com/JRPerezJr)
- LinkedIn: [@devjperez](https://linkedin.com/in/devjperez)