# Installing the Go Lang on the Linux OS
## 1.First thing as to log in as root
    sudo su - root
## 2. change the directory to **temp**.
    cd /tmp
## 3. Download and install Go (version 1.15.6)
   ### Depending on your linux you have to choose download option. we are going with **go1.15.6.linux-amd64.tar.gz**
   >Note: Maku sure you are in **tmp** directory

    wget https://golang.org/dl/go1.15.6.linux-amd64.tar.gz

    sudo tar -C /usr/local -xzf go1.15.6.linux-amd64.tar.gz
## 4. if don not have *git* install on your machine, You can download it now.
    sudo yum install -y git
## 5. check the go directory
    cd /usr/local/go
# Now we have to set up th **GOPath** nad **path** to the Go binaries.

## 1. go back to main root user
    exit
## 2. Create a directory for Go and some subdirectories
     mkdir -p go/{src,bin}
## 3. make modification on *bashrc*
    vim .bashrc
## 4. add the Go path to the bottam of the file.
    export GOPATH=$HOME/go
    export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin   
> for insert i

> for save and quit :wq

![alt](https://github.com/sada498/Go_Lang/blob/main/Installing_Setting_Up_Go_On_Linux/img/bashrc.JPG)
## 5. Reload the shell
    exec $shell
## 6. check the version of the Go
    go version
    
![alt](https://github.com/sada498/Go_Lang/blob/main/Installing_Setting_Up_Go_On_Linux/img/go%20v.JPG)
## 7. Test sample programming
    go run ~/ hello.go
   ![alt](https://github.com/sada498/Go_Lang/blob/main/Installing_Setting_Up_Go_On_Linux/img/hello.JPG)

Done! : )
