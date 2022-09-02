# Standup

Standup is a cli tool to figure out standup order.

### install with following command
```
go install github.com/wontaeyang/standup@latest
```

### flags
* -start: pin names at the start of the standup. Order is preserved.
* -names: names of standup participants. Order is randomized.
* -end: pin names at the end of the standup. Order is preserved.

### examples:
command with names pinned at the start and end

```
> standup -start wontae -names charlie,may,jake,val,lisa -end jakob

// will result in following output
Today is a workday:
[0] - wontae
[1] - charlie
[2] - jake
[3] - val
[4] - may
[5] - lisa
[6] - jakob
```

All random

```
> standup -names wontae,charlie,may,jake,val,lisa,jakob

// will result in following output
Today is a workday:
[0] - may
[1] - lisa
[2] - charlie
[3] - val
[4] - wontae
[5] - jake
[6] - jakob
```



