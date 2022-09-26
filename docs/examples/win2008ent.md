static:

```json
{
  "code": 200,
  "payload": {
    "time": 1664186085,
    "host": {
      "name": "WIN-BTLTI307FJ1",
      "uptime": 343807
    },
    "os": {
      "name": "windows",
      "platform_name": "Microsoft Windows Server 2008 R2 Enterprise Service Pack 1",
      "platform_version": "6.1.7601 Build 7601",
      "install": 1663813578,
      "startup": 1663813493
    },
    "kernel": {
      "version": "6.1.7601 Build 7601",
      "arch": "x86_64"
    },
    "cpu": {
      "physical": 4,
      "logical": 4,
      "cores": [
        {
          "processor": 0,
          "model": "Intel(R) Xeon(R) CPU E5-2630 0 @ 2.30GHz",
          "core": 0,
          "cores": 1,
          "physical": 0,
          "mhz": 2300
        },
        {
          "processor": 1,
          "model": "Intel(R) Xeon(R) CPU E5-2630 0 @ 2.30GHz",
          "core": 0,
          "cores": 1,
          "physical": 0,
          "mhz": 2300
        },
        {
          "processor": 2,
          "model": "Intel(R) Xeon(R) CPU E5-2630 0 @ 2.30GHz",
          "core": 0,
          "cores": 1,
          "physical": 0,
          "mhz": 2300
        },
        {
          "processor": 3,
          "model": "Intel(R) Xeon(R) CPU E5-2630 0 @ 2.30GHz",
          "core": 0,
          "cores": 1,
          "physical": 0,
          "mhz": 2300
        }
      ]
    },
    "memory": {
      "physical": 8589467648,
      "swap": 17177047040
    },
    "disks": [
      {
        "model": "VMware Virtual disk SCSI Disk Device",
        "total": 53686402560,
        "type": "SSD"
      }
    ],
    "partitions": [
      {
        "mount": "C:",
        "fstype": "NTFS",
        "opts": [
          "rw",
          "compress"
        ],
        "total": 53580132352
      },
      {
        "mount": "D:",
        "fstype": "UDF",
        "opts": [
          "ro"
        ],
        "total": 3368962048
      }
    ],
    "gateway": "192.168.2.1",
    "interface": [
      {
        "index": 11,
        "name": "本地连接",
        "mtu": 1500,
        "flags": [
          "up",
          "broadcast",
          "multicast"
        ],
        "mac": "00:50:56:a3:21:d1",
        "addrs": [
          "192.168.2.79/24"
        ]
      },
      {
        "index": 1,
        "name": "Loopback Pseudo-Interface 1",
        "mtu": -1,
        "flags": [
          "up",
          "loopback",
          "multicast"
        ],
        "mac": "",
        "addrs": [
          "::1/128",
          "127.0.0.1/8"
        ]
      },
      {
        "index": 12,
        "name": "isatap.{F16C7F31-8800-42E3-A58F-B6B52B1B9900}",
        "mtu": 1280,
        "flags": [
          "pointtopoint",
          "multicast"
        ],
        "mac": "00:00:00:00:00:00:00:e0",
        "addrs": [
          "fe80::5efe:c0a8:24f/128"
        ]
      }
    ]
  }
}
```

dynamic:

```json
{
  "code": 200,
  "payload": {
    "usage": {
      "cpu": {
        "usage": 0.52
      },
      "memory": {
        "used": 915537920,
        "free": 7673929728,
        "available": 7673929728,
        "usage": 10,
        "swap_used": 0,
        "swap_free": 0
      },
      "partitions": [
        {
          "name": "C:",
          "used": 15875350528,
          "free": 37704781824,
          "usage": 29.63
        },
        {
          "name": "D:",
          "used": 3368962048,
          "free": 0,
          "usage": 100
        }
      ],
      "interface": [
        {
          "name": "本地连接",
          "bytes_sent": 18209415,
          "bytes_recv": 29174916,
          "packets_sent": 147310,
          "packets_recv": 273872
        },
        {
          "name": "Loopback Pseudo-Interface 1",
          "bytes_sent": 0,
          "bytes_recv": 0,
          "packets_sent": 0,
          "packets_recv": 0
        },
        {
          "name": "isatap.{F16C7F31-8800-42E3-A58F-B6B52B1B9900}",
          "bytes_sent": 0,
          "bytes_recv": 0,
          "packets_sent": 0,
          "packets_recv": 0
        }
      ]
    },
    "process": [
      {
        "id": 4,
        "pid": 0,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.02,
        "rss": 372736,
        "vms": 114688,
        "swap": 0,
        "memory_usage": 0,
        "cmd": [
          ""
        ],
        "listen": [
          445,
          47001,
          139,
          445,
          47001
        ],
        "conns": 7
      },
      {
        "id": 152,
        "pid": 468,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 11931648,
        "vms": 9347072,
        "swap": 0,
        "memory_usage": 0.14,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalServiceNoNetwork"
        ],
        "conns": 0
      },
      {
        "id": 252,
        "pid": 4,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 1208320,
        "vms": 565248,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "\\SystemRoot\\System32\\smss.exe"
        ],
        "conns": 0
      },
      {
        "id": 316,
        "pid": 468,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 14209024,
        "vms": 6569984,
        "swap": 0,
        "memory_usage": 0.17,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalSystemNetworkRestricted"
        ],
        "conns": 0
      },
      {
        "id": 332,
        "pid": 324,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 5885952,
        "vms": 2342912,
        "swap": 0,
        "memory_usage": 0.07,
        "cmd": [
          "%SystemRoot%\\system32\\csrss.exe",
          "ObjectDirectory=\\Windows",
          "SharedSection=1024,20480,768",
          "Windows=On",
          "SubSystemType=Windows",
          "ServerDll=basesrv,1",
          "ServerDll=winsrv:UserServerDllInitialization,3",
          "ServerDll=winsrv:ConServerDllInitialization,2",
          "ServerDll=sxssrv,4",
          "ProfileControl=Off",
          "MaxRequestThreads=16"
        ],
        "conns": 0
      },
      {
        "id": 372,
        "pid": 364,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.01,
        "rss": 12963840,
        "vms": 3710976,
        "swap": 0,
        "memory_usage": 0.15,
        "cmd": [
          "%SystemRoot%\\system32\\csrss.exe",
          "ObjectDirectory=\\Windows",
          "SharedSection=1024,20480,768",
          "Windows=On",
          "SubSystemType=Windows",
          "ServerDll=basesrv,1",
          "ServerDll=winsrv:UserServerDllInitialization,3",
          "ServerDll=winsrv:ConServerDllInitialization,2",
          "ServerDll=sxssrv,4",
          "ProfileControl=Off",
          "MaxRequestThreads=16"
        ],
        "conns": 0
      },
      {
        "id": 380,
        "pid": 324,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 5234688,
        "vms": 1835008,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "wininit.exe"
        ],
        "listen": [
          49152,
          49152
        ],
        "conns": 2
      },
      {
        "id": 408,
        "pid": 364,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 7634944,
        "vms": 3231744,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "winlogon.exe"
        ],
        "conns": 0
      },
      {
        "id": 468,
        "pid": 380,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 8761344,
        "vms": 4681728,
        "swap": 0,
        "memory_usage": 0.1,
        "cmd": [
          "C:\\Windows\\system32\\services.exe"
        ],
        "listen": [
          49156,
          49156
        ],
        "conns": 2
      },
      {
        "id": 488,
        "pid": 380,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 13451264,
        "vms": 5935104,
        "swap": 0,
        "memory_usage": 0.16,
        "cmd": [
          "C:\\Windows\\system32\\lsass.exe"
        ],
        "listen": [
          49153,
          49153
        ],
        "conns": 2
      },
      {
        "id": 496,
        "pid": 380,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 6459392,
        "vms": 3117056,
        "swap": 0,
        "memory_usage": 0.08,
        "cmd": [
          "C:\\Windows\\system32\\lsm.exe"
        ],
        "conns": 0
      },
      {
        "id": 600,
        "pid": 468,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 10473472,
        "vms": 4952064,
        "swap": 0,
        "memory_usage": 0.12,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "DcomLaunch"
        ],
        "conns": 0
      },
      {
        "id": 672,
        "pid": 468,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 8769536,
        "vms": 4808704,
        "swap": 0,
        "memory_usage": 0.1,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "RPCSS"
        ],
        "listen": [
          135,
          135
        ],
        "conns": 2
      },
      {
        "id": 748,
        "pid": 1648,
        "user": "WIN-BTLTI307FJ1\\Administrator",
        "cpu_usage": 0.58,
        "rss": 12632064,
        "vms": 3735552,
        "swap": 0,
        "memory_usage": 0.15,
        "cmd": [
          "\"C:\\Windows\\system32\\taskmgr.exe\"",
          "/4"
        ],
        "conns": 0
      },
      {
        "id": 780,
        "pid": 468,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0.01,
        "rss": 13938688,
        "vms": 6995968,
        "swap": 0,
        "memory_usage": 0.16,
        "cmd": [
          "C:\\Windows\\system32\\sppsvc.exe"
        ],
        "conns": 0
      },
      {
        "id": 820,
        "pid": 468,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 14471168,
        "vms": 7417856,
        "swap": 0,
        "memory_usage": 0.17,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalService"
        ],
        "conns": 0
      },
      {
        "id": 844,
        "pid": 468,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0.03,
        "rss": 44097536,
        "vms": 73191424,
        "swap": 0,
        "memory_usage": 0.51,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "NetworkService"
        ],
        "conns": 1
      },
      {
        "id": 932,
        "pid": 468,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0.01,
        "rss": 12767232,
        "vms": 9568256,
        "swap": 0,
        "memory_usage": 0.15,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalServiceNetworkRestricted"
        ],
        "listen": [
          49154,
          49154
        ],
        "conns": 2
      },
      {
        "id": 968,
        "pid": 468,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.04,
        "rss": 39673856,
        "vms": 25624576,
        "swap": 0,
        "memory_usage": 0.46,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "netsvcs"
        ],
        "listen": [
          49155,
          49155
        ],
        "conns": 6
      },
      {
        "id": 976,
        "pid": 468,
        "user": "WIN-BTLTI307FJ1\\Administrator",
        "cpu_usage": 0,
        "rss": 7532544,
        "vms": 2764800,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "\"taskhost.exe\""
        ],
        "conns": 0
      },
      {
        "id": 1128,
        "pid": 468,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 15454208,
        "vms": 9306112,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "C:\\Windows\\System32\\spoolsv.exe"
        ],
        "conns": 0
      },
      {
        "id": 1184,
        "pid": 468,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 3026944,
        "vms": 1138688,
        "swap": 0,
        "memory_usage": 0.04,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "regsvc"
        ],
        "conns": 0
      },
      {
        "id": 1260,
        "pid": 316,
        "user": "WIN-BTLTI307FJ1\\Administrator",
        "cpu_usage": 0,
        "rss": 5726208,
        "vms": 1884160,
        "swap": 0,
        "memory_usage": 0.07,
        "cmd": [
          "\"C:\\Windows\\system32\\Dwm.exe\""
        ],
        "conns": 0
      },
      {
        "id": 1384,
        "pid": 2080,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 15867904,
        "vms": 7467008,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "\"LogonUI.exe\"",
          "/flags:0x0"
        ],
        "conns": 0
      },
      {
        "id": 1412,
        "pid": 468,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 5730304,
        "vms": 2289664,
        "swap": 0,
        "memory_usage": 0.07,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "NetworkServiceNetworkRestricted"
        ],
        "listen": [
          49157,
          49157
        ],
        "conns": 2
      },
      {
        "id": 1648,
        "pid": 1252,
        "user": "WIN-BTLTI307FJ1\\Administrator",
        "cpu_usage": 0.29,
        "rss": 64360448,
        "vms": 34246656,
        "swap": 0,
        "memory_usage": 0.75,
        "cmd": [
          "C:\\Windows\\Explorer.EXE"
        ],
        "conns": 0
      },
      {
        "id": 1748,
        "pid": 2648,
        "user": "WIN-BTLTI307FJ1\\Administrator",
        "cpu_usage": 0,
        "rss": 7282688,
        "vms": 2035712,
        "swap": 0,
        "memory_usage": 0.08,
        "cmd": [
          "rdpclip"
        ],
        "conns": 0
      },
      {
        "id": 1940,
        "pid": 1648,
        "user": "WIN-BTLTI307FJ1\\Administrator",
        "cpu_usage": 0.01,
        "rss": 21270528,
        "vms": 53649408,
        "swap": 0,
        "memory_usage": 0.25,
        "cmd": [
          "\"C:\\Windows\\system32\\mmc.exe\"",
          "\"C:\\Windows\\system32\\wf.msc\"",
          ""
        ],
        "conns": 0
      },
      {
        "id": 1992,
        "pid": 468,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 4857856,
        "vms": 1744896,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalServiceAndNoImpersonation"
        ],
        "conns": 0
      },
      {
        "id": 2028,
        "pid": 468,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 7860224,
        "vms": 3596288,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "C:\\Windows\\System32\\msdtc.exe"
        ],
        "conns": 0
      },
      {
        "id": 2080,
        "pid": 3024,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 4878336,
        "vms": 1851392,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "winlogon.exe"
        ],
        "conns": 0
      },
      {
        "id": 2268,
        "pid": 600,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0.28,
        "rss": 8486912,
        "vms": 4001792,
        "swap": 0,
        "memory_usage": 0.1,
        "cmd": [
          "C:\\Windows\\system32\\wbem\\wmiprvse.exe"
        ],
        "conns": 0
      },
      {
        "id": 2332,
        "pid": 748,
        "user": "WIN-BTLTI307FJ1\\Administrator",
        "cpu_usage": 0.88,
        "rss": 25137152,
        "vms": 13697024,
        "swap": 0,
        "memory_usage": 0.29,
        "cmd": [
          "\"C:\\Windows\\System32\\perfmon.exe\"",
          "/res"
        ],
        "conns": 0
      },
      {
        "id": 2648,
        "pid": 468,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 8851456,
        "vms": 3411968,
        "swap": 0,
        "memory_usage": 0.1,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "termsvcs"
        ],
        "listen": [
          3389,
          3389
        ],
        "conns": 2
      },
      {
        "id": 2900,
        "pid": 468,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.01,
        "rss": 36990976,
        "vms": 45559808,
        "swap": 0,
        "memory_usage": 0.43,
        "cmd": [
          "C:\\metrics-agent\\metrics.exe",
          "-conf",
          "C:\\metrics-agent\\agent.conf"
        ],
        "conns": 1
      },
      {
        "id": 3052,
        "pid": 3024,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 6443008,
        "vms": 2052096,
        "swap": 0,
        "memory_usage": 0.08,
        "cmd": [
          "%SystemRoot%\\system32\\csrss.exe",
          "ObjectDirectory=\\Windows",
          "SharedSection=1024,20480,768",
          "Windows=On",
          "SubSystemType=Windows",
          "ServerDll=basesrv,1",
          "ServerDll=winsrv:UserServerDllInitialization,3",
          "ServerDll=winsrv:ConServerDllInitialization,2",
          "ServerDll=sxssrv,4",
          "ProfileControl=Off",
          "MaxRequestThreads=16"
        ],
        "conns": 0
      }
    ],
    "connections": [
      {
        "fd": 0,
        "pid": 672,
        "type": "tcp4",
        "local": "0.0.0.0:135",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "tcp4",
        "local": "0.0.0.0:445",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 2648,
        "type": "tcp4",
        "local": "0.0.0.0:3389",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "tcp4",
        "local": "0.0.0.0:47001",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 380,
        "type": "tcp4",
        "local": "0.0.0.0:49152",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 488,
        "type": "tcp4",
        "local": "0.0.0.0:49153",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 932,
        "type": "tcp4",
        "local": "0.0.0.0:49154",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 968,
        "type": "tcp4",
        "local": "0.0.0.0:49155",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 468,
        "type": "tcp4",
        "local": "0.0.0.0:49156",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 1412,
        "type": "tcp4",
        "local": "0.0.0.0:49157",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "tcp4",
        "local": "192.168.2.79:139",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 2900,
        "type": "tcp4",
        "local": "192.168.2.79:49159",
        "remote": "192.168.2.52:13081",
        "status": "ESTABLISHED"
      },
      {
        "fd": 0,
        "pid": 672,
        "type": "tcp6",
        "local": ":::135",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "tcp6",
        "local": ":::445",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 2648,
        "type": "tcp6",
        "local": ":::3389",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "tcp6",
        "local": ":::47001",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 380,
        "type": "tcp6",
        "local": ":::49152",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 488,
        "type": "tcp6",
        "local": ":::49153",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 932,
        "type": "tcp6",
        "local": ":::49154",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 968,
        "type": "tcp6",
        "local": ":::49155",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 468,
        "type": "tcp6",
        "local": ":::49156",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 1412,
        "type": "tcp6",
        "local": ":::49157",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 968,
        "type": "udp4",
        "local": "0.0.0.0:500",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 968,
        "type": "udp4",
        "local": "0.0.0.0:4500",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 844,
        "type": "udp4",
        "local": "0.0.0.0:5355",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "udp4",
        "local": "192.168.2.79:137",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "udp4",
        "local": "192.168.2.79:138",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 968,
        "type": "udp6",
        "local": ":::500",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 968,
        "type": "udp6",
        "local": ":::4500",
        "status": ""
      }
    ]
  }
}
```