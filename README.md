***Warning: This library should not be used just yet***

# glesys

Go library for the [GleSYS API](https://github.com/GleSYS/API)

[![GoDoc](https://godoc.org/github.com/peterhellberg/glesys?status.svg)](https://godoc.org/github.com/peterhellberg/glesys)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/glesys#license-mit)

## Installation

```bash
go get -u github.com/peterhellberg/glesys
```

## API Modules

I’ll (slowly) work my way through implementing the code for handling each module/endpoint in the GleSYS API.

### server

 - [x] list
 - [ ] details
 - [ ] status
 - [ ] backup
 - [ ] reboot
 - [ ] stop
 - [ ] start
 - [ ] create
 - [ ] destroy
 - [ ] edit
 - [ ] clone
 - [ ] limits
 - [ ] resetlimit
 - [ ] console
 - [ ] resetpassword
 - [ ] templates
 - [ ] allowedarguments
 - [ ] resourceusage
 - [ ] costs
 - [ ] listiso
 - [ ] mountiso
 - [ ] addiso

### ip

 - [x] listfree
 - [ ] listown
 - [ ] details
 - [ ] take
 - [ ] release
 - [ ] add
 - [ ] remove
 - [ ] setptr
 - [ ] resetptr

### domain

 - [ ] list
 - [ ] add
 - [ ] register
 - [ ] transfer
 - [ ] renew
 - [ ] setautorenew
 - [ ] details
 - [ ] available
 - [ ] pricelist
 - [ ] edit
 - [ ] delete
 - [ ] updaterecord
 - [ ] listrecords
 - [ ] addrecord
 - [ ] deleterecord
 - [ ] allowedarguments
 - [ ] changenameservers
 - [ ] legacywebhosting

### archive

 - [ ] details
 - [ ] delete
 - [ ] list
 - [ ] create
 - [ ] changepassword
 - [ ] resize
 - [ ] changedescription
 - [ ] allowedarguments

### email

 - [ ] overview
 - [ ] globalquota
 - [ ] list
 - [ ] editaccount
 - [ ] delete
 - [ ] createaccount
 - [ ] quota
 - [ ] createalias
 - [ ] editalias
 - [ ] costs

### invoice

 - [ ] allowedarguments
 - [ ] billingperiod
 - [ ] settings
 - [ ] list
 - [ ] next
 - [ ] paymenthistory
 - [ ] paybycard
 - [ ] paybypaypal

### changelog

 - [x] api
 - [x] controlpanel

### customer

 - [x] contactinfo

### livechat

 - [ ] newsession
 - [ ] getmessages
 - [ ] closesession
 - [ ] postmessage
 - [ ] status
 - [ ] sessioninfo

### account

 - [ ] changepassword
 - [ ] info

### paymentcard

 - [ ] list
 - [ ] delete
 - [ ] edit
 - [ ] add

### contactperson

 - [ ] list
 - [ ] add
 - [ ] edit
 - [ ] delete

### transaction

 - [ ] list
 - [ ] cancel
 - [ ] start
 - [ ] acknowledge

### vpn

 - [ ] listusers
 - [ ] createuser
 - [ ] deleteuser
 - [ ] edituser

### loadbalancer

 - [ ] list
 - [ ] details
 - [ ] create
 - [ ] edit
 - [ ] addfrontend
 - [ ] editfrontend
 - [ ] addbackend
 - [ ] editbackend
 - [ ] addtarget
 - [ ] removetarget
 - [ ] enabletarget
 - [ ] disabletarget
 - [ ] edittarget
 - [ ] removebackend
 - [ ] removefrontend
 - [ ] destroy
 - [ ] addcertificate
 - [ ] removecertificate
 - [ ] listcertificate

### user

 - [ ] details
 - [ ] edit
 - [ ] login
 - [ ] logout
 - [ ] enabletwofactor
 - [ ] disabletwofactor
 - [ ] changepassword

### api

 - [ ] maintenance
 - [ ] serviceinfo
 - [ ] listfunctions

### sshkey

 - [ ] add
 - [ ] list
 - [ ] remove

## License (MIT)

Copyright (c) 2015 [Peter Hellberg](http://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
