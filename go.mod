module golang.zx2c4.com/wireguard/windows

require (
	github.com/lxn/walk v0.0.0-20191128110447-55ccb3a9f5c1
	github.com/lxn/win v0.0.0-20191128105842-2da648fda5b4
	golang.org/x/crypto v0.0.0-20200221231518-2aa609cf4a9d
	golang.org/x/net v0.0.0-20200222125558-5a598a2470a0
	golang.org/x/sys v0.0.0-20200219091948-cb0a6d8edb6c
	golang.org/x/text v0.3.2
	golang.zx2c4.com/wireguard v0.0.20200122-0.20200214175355-9cbcff10dd3e
)

replace (
	github.com/lxn/walk => golang.zx2c4.com/wireguard/windows v0.0.0-20191128151049-87f28cc339ec
	github.com/lxn/win => golang.zx2c4.com/wireguard/windows v0.0.0-20191128151145-b4e4933852d5
)

go 1.13
