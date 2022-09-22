static:

```json
{
  "code": 200,
  "payload": {
    "time": 1663812632,
    "host": {
      "name": "JKSTACK-PC",
      "uptime": 156749
    },
    "os": {
      "name": "windows",
      "platform_name": "Microsoft Windows 7 Enterprise Service Pack 1",
      "platform_version": "6.1.7601 Build 7601",
      "install": 1247538897,
      "startup": 1663655882
    },
    "kernel": {
      "version": "6.1.7601 Build 7601",
      "arch": "x86_64"
    },
    "cpu": {
      "physical": 2,
      "logical": 2,
      "cores": [
        {
          "processor": 0,
          "model": "Intel(R) Xeon(R) CPU E5-2670 v3 @ 2.30GHz",
          "core": 0,
          "cores": 1,
          "physical": 0,
          "mhz": 2297
        },
        {
          "processor": 1,
          "model": "Intel(R) Xeon(R) CPU E5-2670 v3 @ 2.30GHz",
          "core": 0,
          "cores": 1,
          "physical": 0,
          "mhz": 2297
        }
      ]
    },
    "memory": {
      "physical": 4294500352,
      "swap": 8587112448
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
        "mac": "00:50:56:a3:0d:ad",
        "addrs": [
          "192.168.2.70/24"
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
        "name": "isatap.{FC84BCFA-A305-451D-B651-8B4D2E182E8C}",
        "mtu": 1280,
        "flags": [
          "pointtopoint",
          "multicast"
        ],
        "mac": "00:00:00:00:00:00:00:e0"
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
        "usage": 2.02
      },
      "memory": {
        "used": 1164648448,
        "free": 3129851904,
        "available": 3129851904,
        "usage": 27,
        "swap_used": 0,
        "swap_free": 0
      },
      "partitions": [
        {
          "name": "C:",
          "used": 13739737088,
          "free": 39840395264,
          "usage": 25.64
        }
      ],
      "interface": [
        {
          "name": "本地连接",
          "bytes_sent": 28767,
          "bytes_recv": 18246,
          "packets_sent": 263,
          "packets_recv": 235
        },
        {
          "name": "Loopback Pseudo-Interface 1",
          "bytes_sent": 0,
          "bytes_recv": 0,
          "packets_sent": 0,
          "packets_recv": 0
        },
        {
          "name": "isatap.{FC84BCFA-A305-451D-B651-8B4D2E182E8C}",
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
        "cpu_usage": 0.07,
        "rss": 6168576,
        "vms": 237568,
        "swap": 0,
        "memory_usage": 0.14,
        "cmd": [
          ""
        ],
        "listen": [
          445,
          139,
          445
        ],
        "conns": 5
      },
      {
        "id": 276,
        "pid": 4,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 1163264,
        "vms": 471040,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "\\SystemRoot\\System32\\smss.exe"
        ],
        "conns": 0
      },
      {
        "id": 296,
        "pid": 512,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.01,
        "rss": 37195776,
        "vms": 45535232,
        "swap": 0,
        "memory_usage": 0.87,
        "cmd": [
          "C:\\metrics-agent\\metrics.exe",
          "-conf",
          "C:\\metrics-agent\\agent.conf"
        ],
        "conns": 1
      },
      {
        "id": 352,
        "pid": 344,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 5607424,
        "vms": 2138112,
        "swap": 0,
        "memory_usage": 0.13,
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
        "pid": 512,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0.02,
        "rss": 47869952,
        "vms": 79077376,
        "swap": 0,
        "memory_usage": 1.11,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "NetworkService"
        ],
        "listen": [
          3389,
          3389
        ],
        "conns": 3
      },
      {
        "id": 408,
        "pid": 396,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 7815168,
        "vms": 8212480,
        "swap": 0,
        "memory_usage": 0.18,
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
        "id": 416,
        "pid": 344,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 5046272,
        "vms": 1519616,
        "swap": 0,
        "memory_usage": 0.12,
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
        "id": 452,
        "pid": 396,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 7118848,
        "vms": 2772992,
        "swap": 0,
        "memory_usage": 0.17,
        "cmd": [
          "winlogon.exe"
        ],
        "conns": 0
      },
      {
        "id": 512,
        "pid": 416,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.01,
        "rss": 11677696,
        "vms": 5275648,
        "swap": 0,
        "memory_usage": 0.27,
        "cmd": [
          "C:\\Windows\\system32\\services.exe"
        ],
        "listen": [
          49155,
          49155
        ],
        "conns": 2
      },
      {
        "id": 520,
        "pid": 416,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.04,
        "rss": 12144640,
        "vms": 4657152,
        "swap": 0,
        "memory_usage": 0.28,
        "cmd": [
          "C:\\Windows\\system32\\lsass.exe"
        ],
        "listen": [
          49156,
          49156
        ],
        "conns": 2
      },
      {
        "id": 524,
        "pid": 512,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 13705216,
        "vms": 6582272,
        "swap": 0,
        "memory_usage": 0.32,
        "cmd": [
          "C:\\Windows\\System32\\spoolsv.exe"
        ],
        "conns": 0
      },
      {
        "id": 532,
        "pid": 416,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 6746112,
        "vms": 3203072,
        "swap": 0,
        "memory_usage": 0.16,
        "cmd": [
          "C:\\Windows\\system32\\lsm.exe"
        ],
        "conns": 0
      },
      {
        "id": 556,
        "pid": 1720,
        "user": "JKSTACK-PC\\Administrator",
        "cpu_usage": 0.39,
        "rss": 37855232,
        "vms": 25505792,
        "swap": 0,
        "memory_usage": 0.88,
        "cmd": [
          "\"C:\\Windows\\System32\\perfmon.exe\"",
          "/res"
        ],
        "conns": 0
      },
      {
        "id": 628,
        "pid": 512,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.14,
        "rss": 10694656,
        "vms": 4571136,
        "swap": 0,
        "memory_usage": 0.25,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "DcomLaunch"
        ],
        "conns": 0
      },
      {
        "id": 704,
        "pid": 512,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0.01,
        "rss": 9453568,
        "vms": 4857856,
        "swap": 0,
        "memory_usage": 0.22,
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
        "id": 780,
        "pid": 452,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 20893696,
        "vms": 9961472,
        "swap": 0,
        "memory_usage": 0.49,
        "cmd": [
          "\"LogonUI.exe\"",
          "/flags:0x0"
        ],
        "conns": 0
      },
      {
        "id": 804,
        "pid": 512,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0.04,
        "rss": 19681280,
        "vms": 18128896,
        "swap": 0,
        "memory_usage": 0.46,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalServiceNetworkRestricted"
        ],
        "listen": [
          49153,
          49153
        ],
        "conns": 2
      },
      {
        "id": 836,
        "pid": 512,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.22,
        "rss": 88543232,
        "vms": 81276928,
        "swap": 0,
        "memory_usage": 2.06,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalSystemNetworkRestricted"
        ],
        "conns": 0
      },
      {
        "id": 868,
        "pid": 512,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.15,
        "rss": 36044800,
        "vms": 21184512,
        "swap": 0,
        "memory_usage": 0.84,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "netsvcs"
        ],
        "listen": [
          49154,
          49154
        ],
        "conns": 6
      },
      {
        "id": 980,
        "pid": 512,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0.01,
        "rss": 19050496,
        "vms": 9949184,
        "swap": 0,
        "memory_usage": 0.44,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalService"
        ],
        "conns": 0
      },
      {
        "id": 1036,
        "pid": 512,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 4554752,
        "vms": 2314240,
        "swap": 0,
        "memory_usage": 0.11,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "SDRSVC"
        ],
        "conns": 0
      },
      {
        "id": 1052,
        "pid": 512,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0.03,
        "rss": 14082048,
        "vms": 11722752,
        "swap": 0,
        "memory_usage": 0.33,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalServiceNoNetwork"
        ],
        "conns": 0
      },
      {
        "id": 1272,
        "pid": 512,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 8380416,
        "vms": 3633152,
        "swap": 0,
        "memory_usage": 0.2,
        "cmd": [
          "C:\\Windows\\System32\\msdtc.exe"
        ],
        "conns": 0
      },
      {
        "id": 1280,
        "pid": 836,
        "user": "JKSTACK-PC\\Administrator",
        "cpu_usage": 0,
        "rss": 5701632,
        "vms": 1953792,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          "\"C:\\Windows\\system32\\Dwm.exe\""
        ],
        "conns": 0
      },
      {
        "id": 1288,
        "pid": 512,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 10883072,
        "vms": 6123520,
        "swap": 0,
        "memory_usage": 0.25,
        "cmd": [
          "\"C:\\Program",
          "Files\\VMware\\VMware",
          "Tools\\VMware",
          "VGAuth\\VGAuthService.exe\""
        ],
        "conns": 0
      },
      {
        "id": 1332,
        "pid": 512,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.1,
        "rss": 23310336,
        "vms": 14495744,
        "swap": 0,
        "memory_usage": 0.54,
        "cmd": [
          "\"C:\\Program",
          "Files\\VMware\\VMware",
          "Tools\\vmtoolsd.exe\""
        ],
        "conns": 0
      },
      {
        "id": 1572,
        "pid": 512,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 5812224,
        "vms": 1966080,
        "swap": 0,
        "memory_usage": 0.14,
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
        "id": 1592,
        "pid": 1816,
        "user": "JKSTACK-PC\\Administrator",
        "cpu_usage": 0.52,
        "rss": 37748736,
        "vms": 25387008,
        "swap": 0,
        "memory_usage": 0.88,
        "cmd": [
          "\"C:\\Windows\\System32\\perfmon.exe\"",
          "/res"
        ],
        "conns": 0
      },
      {
        "id": 1816,
        "pid": 1828,
        "user": "JKSTACK-PC\\Administrator",
        "cpu_usage": 0.52,
        "rss": 13463552,
        "vms": 3448832,
        "swap": 0,
        "memory_usage": 0.31,
        "cmd": [
          "\"C:\\Windows\\system32\\taskmgr.exe\"",
          "/4"
        ],
        "conns": 0
      },
      {
        "id": 1824,
        "pid": 512,
        "user": "JKSTACK-PC\\Administrator",
        "cpu_usage": 0,
        "rss": 12124160,
        "vms": 8900608,
        "swap": 0,
        "memory_usage": 0.28,
        "cmd": [
          "\"taskhost.exe\""
        ],
        "conns": 0
      },
      {
        "id": 1828,
        "pid": 2248,
        "user": "JKSTACK-PC\\Administrator",
        "cpu_usage": 0.11,
        "rss": 73379840,
        "vms": 42758144,
        "swap": 0,
        "memory_usage": 1.71,
        "cmd": [
          "C:\\Windows\\Explorer.EXE"
        ],
        "conns": 0
      },
      {
        "id": 1908,
        "pid": 628,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0.46,
        "rss": 17920000,
        "vms": 10457088,
        "swap": 0,
        "memory_usage": 0.42,
        "cmd": [
          "C:\\Windows\\system32\\wbem\\wmiprvse.exe"
        ],
        "conns": 0
      },
      {
        "id": 1996,
        "pid": 512,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 11636736,
        "vms": 4276224,
        "swap": 0,
        "memory_usage": 0.27,
        "cmd": [
          "C:\\Windows\\system32\\dllhost.exe",
          "/Processid:{02D4B3F1-FD88-11D1-960D-00805FC79235}"
        ],
        "conns": 0
      },
      {
        "id": 2228,
        "pid": 372,
        "user": "JKSTACK-PC\\Administrator",
        "cpu_usage": 0,
        "rss": 8855552,
        "vms": 2502656,
        "swap": 0,
        "memory_usage": 0.21,
        "cmd": [
          "rdpclip"
        ],
        "conns": 0
      },
      {
        "id": 2244,
        "pid": 896,
        "user": "JKSTACK-PC\\Administrator",
        "cpu_usage": 0.01,
        "rss": 13434880,
        "vms": 3829760,
        "swap": 0,
        "memory_usage": 0.31,
        "cmd": [
          "msinfo32"
        ],
        "conns": 0
      },
      {
        "id": 2380,
        "pid": 1396,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 7544832,
        "vms": 2928640,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "winlogon.exe"
        ],
        "conns": 0
      },
      {
        "id": 2424,
        "pid": 512,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 9879552,
        "vms": 3555328,
        "swap": 0,
        "memory_usage": 0.23,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalServiceAndNoImpersonation"
        ],
        "conns": 6
      },
      {
        "id": 2472,
        "pid": 512,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0.04,
        "rss": 9842688,
        "vms": 6029312,
        "swap": 0,
        "memory_usage": 0.23,
        "cmd": [
          "C:\\Windows\\system32\\sppsvc.exe"
        ],
        "conns": 0
      },
      {
        "id": 2508,
        "pid": 512,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.38,
        "rss": 41910272,
        "vms": 68579328,
        "swap": 0,
        "memory_usage": 0.98,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "secsvcs"
        ],
        "conns": 0
      },
      {
        "id": 2548,
        "pid": 1080,
        "user": "JKSTACK-PC\\Administrator",
        "cpu_usage": 0.01,
        "rss": 21118976,
        "vms": 13131776,
        "swap": 0,
        "memory_usage": 0.49,
        "cmd": [
          "\"C:\\Windows\\system32\\mmc.exe\"",
          "\"C:\\Windows\\system32\\services.msc\"",
          ""
        ],
        "conns": 0
      },
      {
        "id": 2580,
        "pid": 512,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 25452544,
        "vms": 33333248,
        "swap": 0,
        "memory_usage": 0.59,
        "cmd": [
          "C:\\Windows\\system32\\SearchIndexer.exe",
          "/Embedding"
        ],
        "conns": 0
      },
      {
        "id": 2628,
        "pid": 1828,
        "user": "JKSTACK-PC\\Administrator",
        "cpu_usage": 0.11,
        "rss": 10641408,
        "vms": 5107712,
        "swap": 0,
        "memory_usage": 0.25,
        "cmd": [
          "\"C:\\Program",
          "Files\\VMware\\VMware",
          "Tools\\vmtoolsd.exe\"",
          "-n",
          "vmusr"
        ],
        "conns": 0
      },
      {
        "id": 2796,
        "pid": 1816,
        "user": "JKSTACK-PC\\Administrator",
        "cpu_usage": 0.45,
        "rss": 37588992,
        "vms": 25243648,
        "swap": 0,
        "memory_usage": 0.88,
        "cmd": [
          "\"C:\\Windows\\System32\\perfmon.exe\"",
          "/res"
        ],
        "conns": 0
      },
      {
        "id": 2836,
        "pid": 512,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0.01,
        "rss": 16945152,
        "vms": 14630912,
        "swap": 0,
        "memory_usage": 0.39,
        "cmd": [
          "\"C:\\Program",
          "Files\\Windows",
          "Media",
          "Player\\wmpnetwk.exe\""
        ],
        "conns": 0
      },
      {
        "id": 2884,
        "pid": 1396,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.01,
        "rss": 16506880,
        "vms": 2531328,
        "swap": 0,
        "memory_usage": 0.38,
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
        "id": 3600,
        "pid": 1828,
        "user": "JKSTACK-PC\\Administrator",
        "cpu_usage": 0,
        "rss": 3219456,
        "vms": 2064384,
        "swap": 0,
        "memory_usage": 0.07,
        "cmd": [
          "\"C:\\Windows\\system32\\cmd.exe\"",
          ""
        ],
        "conns": 0
      },
      {
        "id": 3616,
        "pid": 2884,
        "user": "JKSTACK-PC\\Administrator",
        "cpu_usage": 0.01,
        "rss": 7225344,
        "vms": 2949120,
        "swap": 0,
        "memory_usage": 0.17,
        "cmd": [
          "\\??\\C:\\Windows\\system32\\conhost.exe"
        ],
        "conns": 0
      },
      {
        "id": 3824,
        "pid": 1816,
        "user": "JKSTACK-PC\\Administrator",
        "cpu_usage": 0.43,
        "rss": 36675584,
        "vms": 24424448,
        "swap": 0,
        "memory_usage": 0.85,
        "cmd": [
          "\"C:\\Windows\\System32\\perfmon.exe\"",
          "/res"
        ],
        "conns": 0
      },
      {
        "id": 3908,
        "pid": 512,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0.14,
        "rss": 10330112,
        "vms": 4202496,
        "swap": 0,
        "memory_usage": 0.24,
        "cmd": [
          "taskhost.exe",
          "$(Arg0)"
        ],
        "conns": 0
      }
    ],
    "connections": [
      {
        "fd": 0,
        "pid": 704,
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
        "pid": 372,
        "type": "tcp4",
        "local": "0.0.0.0:3389",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 416,
        "type": "tcp4",
        "local": "0.0.0.0:49152",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 804,
        "type": "tcp4",
        "local": "0.0.0.0:49153",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 868,
        "type": "tcp4",
        "local": "0.0.0.0:49154",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 512,
        "type": "tcp4",
        "local": "0.0.0.0:49155",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 520,
        "type": "tcp4",
        "local": "0.0.0.0:49156",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 1572,
        "type": "tcp4",
        "local": "0.0.0.0:49157",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "tcp4",
        "local": "192.168.2.70:139",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 296,
        "type": "tcp4",
        "local": "192.168.2.70:49167",
        "remote": "192.168.2.52:13081",
        "status": "ESTABLISHED"
      },
      {
        "fd": 0,
        "pid": 704,
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
        "pid": 372,
        "type": "tcp6",
        "local": ":::3389",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 416,
        "type": "tcp6",
        "local": ":::49152",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 804,
        "type": "tcp6",
        "local": ":::49153",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 868,
        "type": "tcp6",
        "local": ":::49154",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 512,
        "type": "tcp6",
        "local": ":::49155",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 520,
        "type": "tcp6",
        "local": ":::49156",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 1572,
        "type": "tcp6",
        "local": ":::49157",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 868,
        "type": "udp4",
        "local": "0.0.0.0:500",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 868,
        "type": "udp4",
        "local": "0.0.0.0:4500",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 372,
        "type": "udp4",
        "local": "0.0.0.0:5355",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 2424,
        "type": "udp4",
        "local": "127.0.0.1:1900",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 2424,
        "type": "udp4",
        "local": "127.0.0.1:61156",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "udp4",
        "local": "192.168.2.70:137",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "udp4",
        "local": "192.168.2.70:138",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 2424,
        "type": "udp4",
        "local": "192.168.2.70:1900",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 2424,
        "type": "udp4",
        "local": "192.168.2.70:61155",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 868,
        "type": "udp6",
        "local": ":::500",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 868,
        "type": "udp6",
        "local": ":::4500",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 2424,
        "type": "udp6",
        "local": "::1:1900",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 2424,
        "type": "udp6",
        "local": "::1:61154",
        "status": ""
      }
    ]
  }
}
```