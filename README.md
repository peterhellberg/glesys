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

### [account](https://github.com/GleSYS/API/wiki/functions_account)

 - [ ] changepassword
 - [x] info

### [api](https://github.com/GleSYS/API/wiki/functions_api)

 - [x] maintenance
 - [x] serviceinfo
 - [x] listfunctions

### [archive](https://github.com/GleSYS/API/wiki/functions_archive)

 - [ ] details
 - [ ] delete
 - [ ] list
 - [ ] create
 - [ ] changepassword
 - [ ] resize
 - [ ] changedescription
 - [ ] allowedarguments

### [changelog](https://github.com/GleSYS/API/wiki/functions_changelog)

 - [x] api
 - [x] controlpanel

### [contactperson](https://github.com/GleSYS/API/wiki/functions_contactperson)

 - [ ] list
 - [ ] add
 - [ ] edit
 - [ ] delete

### [customer](https://github.com/GleSYS/API/wiki/functions_customer)

 - [x] contactinfo

### [domain](https://github.com/GleSYS/API/wiki/functions_domain)

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

### [email](https://github.com/GleSYS/API/wiki/functions_email)

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

### [invoice](https://github.com/GleSYS/API/wiki/functions_invoice)

 - [ ] allowedarguments
 - [ ] billingperiod
 - [ ] settings
 - [ ] list
 - [ ] next
 - [ ] paymenthistory
 - [ ] paybycard
 - [ ] paybypaypal

### [ip](https://github.com/GleSYS/API/wiki/functions_ip)

 - [x] listfree
 - [ ] listown
 - [ ] details
 - [ ] take
 - [ ] release
 - [ ] add
 - [ ] remove
 - [ ] setptr
 - [ ] resetptr

### [livechat](https://github.com/GleSYS/API/wiki/functions_livechat)

 - [ ] newsession
 - [ ] getmessages
 - [ ] closesession
 - [ ] postmessage
 - [ ] status
 - [ ] sessioninfo

### [loadbalancer](https://github.com/GleSYS/API/wiki/functions_loadbalancer)

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

### [paymentcard](https://github.com/GleSYS/API/wiki/functions_paymentcard)

 - [ ] list
 - [ ] delete
 - [ ] edit
 - [ ] add

### [server](https://github.com/GleSYS/API/wiki/functions_server)

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

### [sshkey](https://github.com/GleSYS/API/wiki/functions_sshkey)

 - [ ] add
 - [ ] list
 - [ ] remove

### [transaction](https://github.com/GleSYS/API/wiki/functions_transaction)

 - [ ] list
 - [ ] cancel
 - [ ] start
 - [ ] acknowledge

### [user](https://github.com/GleSYS/API/wiki/functions_user)

 - [ ] details
 - [ ] edit
 - [ ] login
 - [ ] logout
 - [ ] enabletwofactor
 - [ ] disabletwofactor
 - [ ] changepassword

### [vpn](https://github.com/GleSYS/API/wiki/functions_vpn)

 - [ ] listusers
 - [ ] createuser
 - [ ] deleteuser
 - [ ] edituser

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
