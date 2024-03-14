# Cache Variable
Variable write to ram and read from ram instead of disk.

Sometimes, users define variables in the environment on the disk, which may cause frequent disk reads. Using this command can enable reading/writing directly from RAM, reducing disk read/write times and increasing efficiency. For example, you can use the cr get and set commands to read and write variables to RAM.


This project is written in Go.

### Steps for use

1. Download release file `Cache.Variable-xxx-linux-x64.tar.xz` and extract it.

2. Grant it executable permission and move it to bin dir.    
```
chmod +x cr

sudo cp cr /usr/local/bin/
```

3. Run a cache variable instalce.  // Sometime may need run at startup.
```
cr run    // it will run a instance at default port 6366, you can change it with -p xxxx.
```


4. Write a variable.
```
cr set key value 10   // after 10 sec its expire. 
```

5. Read a variable.
```
cr get key            // it will return value.

cr get key -d 0       // it will return 0 when internal error. 
```
also you can use cr get with other command in shell.


### Command Usage And Options
```
cr [command] [flags]
```

### Available Commands
```
cmd   flags                args                                 

run                                                Run a cache_variable server at local.
      -p, --port                                   optional, specific the port, default port is 6366.
      -s, --size                                   optional, specific the cache size, default is 100, unit is MB.


set                        key value expire        Add a variable to ram, must specific the expire time, unit is second.
      -p, --port                                   optional, specific the port of target server, default port is 6366.
      

get                        key                     Get a variable from ram.
      -p, --port                                   optional, specific the port of target server, default port is 6366.
      -d, --defaultValue                           optional, specific the defaultValue when internal error, if not specific it will show error msg.

del                        key                     Delete a variable from ram.
      -p, --port                                   optional, specific the port of target server, default port is 6366.
      -d, --defaultValue                           optional, specific the defaultValue when internal error, if not specific it will show error msg.

help                                               Help about any command.
```



### Use Case Scenario
- Conky config

Sometime conky config may ask a tmp file, update frequently and read frequently, with this command, can store variable to ram instead of disk, it can reducing disk read/write times.

![result](https://github.com/MarsSwimmer/cache_variable/assets/146618222/81bd8b34-ac37-47e7-a6dc-d4f64ac808af)


