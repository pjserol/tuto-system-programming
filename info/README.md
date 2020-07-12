
# Notes

**show current list of running processes for the active user**

    ps -f

**launch process in the background**

    cat file.txt &

**&& executes next instruction only if the exit code of the first one is 0**
- EX_OK 0
- echo blabla >> tst.txt
- stat tst.txt && echo someText >> tst.txt

**list of available signals**

    kill -l

**Anonymous pipes**

    ls -l | grep “test”

**redirect output to a file**

    ls -l > file.txt

**use file as an input for another process**

    cat < file.txt

**>> operator, which will start writing from the end of the file.**

    echo blabla >> tst.txt

**Go stack & heap**

    go tool compile -m -h main.go

**-m print optimisation decision**

    go build -gcflags=“-m”
    go build -gcflags=“-m -l”

**more outputs**

    go build -gcflags="-m=2"

**display go settings**

    go env

**build a binary in the current directory**

    go build .

**get the limit of open files**

    ulimit -n

**check how many files are open by a process**

    lsof -p PID

**count number of line**

    wc -l test.txt

**start app with PID**

    $ go build -o "signal" pjserol/app.go
    $ ./signal &

**convert json to go struct**
https://mholt.github.io/json-to-go/

  
**convert XML to go struct**
https://www.onlinetool.io/xmltogo/