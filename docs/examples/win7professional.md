static:

```json
{
  "code": 200,
  "payload": {
    "time": 1663812792,
    "host": {
      "name": "jkstack-PC",
      "uptime": 156911
    },
    "os": {
      "name": "windows",
      "platform_name": "Microsoft Windows 7 Professional Service Pack 1",
      "platform_version": "6.1.7601 Build 7601",
      "install": 1247538897,
      "startup": 1663655881
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
        "mac": "00:50:56:a3:37:6c",
        "addrs": [
          "192.168.2.71/24"
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
        "name": "isatap.{6104ED38-1193-4091-9501-189BF6ABBA5E}",
        "mtu": 1280,
        "flags": [
          "pointtopoint",
          "multicast"
        ],
        "mac": "00:00:00:00:00:00:00:e0",
        "addrs": [
          "fe80::5efe:c0a8:247/128"
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
        "usage": 0.1
      },
      "memory": {
        "used": 806174720,
        "free": 3488325632,
        "available": 3488325632,
        "usage": 18,
        "swap_used": 0,
        "swap_free": 0
      },
      "partitions": [
        {
          "name": "C:",
          "used": 16373157888,
          "free": 37206974464,
          "usage": 30.56
        }
      ],
      "interface": [
        {
          "name": "本地连接",
          "bytes_sent": 30025171,
          "bytes_recv": 19843168,
          "packets_sent": 84053,
          "packets_recv": 173627
        },
        {
          "name": "Loopback Pseudo-Interface 1",
          "bytes_sent": 0,
          "bytes_recv": 0,
          "packets_sent": 0,
          "packets_recv": 0
        },
        {
          "name": "isatap.{6104ED38-1193-4091-9501-189BF6ABBA5E}",
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
        "rss": 6467584,
        "vms": 245760,
        "swap": 0,
        "memory_usage": 0.15,
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
        "id": 196,
        "pid": 836,
        "user": "jkstack-PC\\Administrator",
        "cpu_usage": 0,
        "rss": 5337088,
        "vms": 1691648,
        "swap": 0,
        "memory_usage": 0.12,
        "cmd": [
          "\"C:\\Windows\\system32\\Dwm.exe\""
        ],
        "conns": 0
      },
      {
        "id": 260,
        "pid": 4,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 1114112,
        "vms": 434176,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "\\SystemRoot\\System32\\smss.exe"
        ],
        "conns": 0
      },
      {
        "id": 336,
        "pid": 324,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 5419008,
        "vms": 2150400,
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
        "id": 356,
        "pid": 476,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 15400960,
        "vms": 10706944,
        "swap": 0,
        "memory_usage": 0.36,
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
        "id": 384,
        "pid": 324,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 5083136,
        "vms": 1564672,
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
        "id": 408,
        "pid": 376,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 4616192,
        "vms": 1802240,
        "swap": 0,
        "memory_usage": 0.11,
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
        "id": 440,
        "pid": 376,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 7127040,
        "vms": 2576384,
        "swap": 0,
        "memory_usage": 0.17,
        "cmd": [
          "winlogon.exe"
        ],
        "conns": 0
      },
      {
        "id": 476,
        "pid": 384,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 10866688,
        "vms": 4452352,
        "swap": 0,
        "memory_usage": 0.25,
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
        "id": 488,
        "pid": 384,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 10883072,
        "vms": 3923968,
        "swap": 0,
        "memory_usage": 0.25,
        "cmd": [
          "C:\\Windows\\system32\\lsass.exe"
        ],
        "listen": [
          49157,
          49157
        ],
        "conns": 2
      },
      {
        "id": 496,
        "pid": 384,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 6103040,
        "vms": 2981888,
        "swap": 0,
        "memory_usage": 0.14,
        "cmd": [
          "C:\\Windows\\system32\\lsm.exe"
        ],
        "conns": 0
      },
      {
        "id": 612,
        "pid": 476,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 9699328,
        "vms": 4100096,
        "swap": 0,
        "memory_usage": 0.23,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "DcomLaunch"
        ],
        "conns": 0
      },
      {
        "id": 688,
        "pid": 476,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 7458816,
        "vms": 3665920,
        "swap": 0,
        "memory_usage": 0.17,
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
        "id": 772,
        "pid": 440,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 19677184,
        "vms": 8773632,
        "swap": 0,
        "memory_usage": 0.46,
        "cmd": [
          "\"LogonUI.exe\"",
          "/flags:0x0"
        ],
        "conns": 0
      },
      {
        "id": 780,
        "pid": 476,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 17317888,
        "vms": 16343040,
        "swap": 0,
        "memory_usage": 0.4,
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
        "pid": 476,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.08,
        "rss": 59285504,
        "vms": 52588544,
        "swap": 0,
        "memory_usage": 1.38,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalSystemNetworkRestricted"
        ],
        "conns": 0
      },
      {
        "id": 860,
        "pid": 476,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 29528064,
        "vms": 15499264,
        "swap": 0,
        "memory_usage": 0.69,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "netsvcs"
        ],
        "listen": [
          49154,
          49154
        ],
        "conns": 2
      },
      {
        "id": 992,
        "pid": 476,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 12689408,
        "vms": 5828608,
        "swap": 0,
        "memory_usage": 0.3,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalService"
        ],
        "conns": 0
      },
      {
        "id": 1036,
        "pid": 476,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 11698176,
        "vms": 6180864,
        "swap": 0,
        "memory_usage": 0.27,
        "cmd": [
          "C:\\Windows\\System32\\spoolsv.exe"
        ],
        "conns": 0
      },
      {
        "id": 1072,
        "pid": 476,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 13291520,
        "vms": 10858496,
        "swap": 0,
        "memory_usage": 0.31,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalServiceNoNetwork"
        ],
        "conns": 0
      },
      {
        "id": 1220,
        "pid": 476,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 19693568,
        "vms": 24047616,
        "swap": 0,
        "memory_usage": 0.46,
        "cmd": [
          "C:\\Windows\\system32\\SearchIndexer.exe",
          "/Embedding"
        ],
        "conns": 0
      },
      {
        "id": 1460,
        "pid": 356,
        "user": "jkstack-PC\\Administrator",
        "cpu_usage": 0,
        "rss": 7569408,
        "vms": 2080768,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "rdpclip"
        ],
        "conns": 0
      },
      {
        "id": 1496,
        "pid": 832,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 12075008,
        "vms": 2265088,
        "swap": 0,
        "memory_usage": 0.28,
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
        "id": 1744,
        "pid": 476,
        "user": "jkstack-PC\\Administrator",
        "cpu_usage": 0,
        "rss": 8339456,
        "vms": 3043328,
        "swap": 0,
        "memory_usage": 0.19,
        "cmd": [
          "\"taskhost.exe\""
        ],
        "conns": 0
      },
      {
        "id": 1752,
        "pid": 832,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 7540736,
        "vms": 2965504,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "winlogon.exe"
        ],
        "conns": 0
      },
      {
        "id": 1920,
        "pid": 476,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 6746112,
        "vms": 2531328,
        "swap": 0,
        "memory_usage": 0.16,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalServiceAndNoImpersonation"
        ],
        "conns": 6
      },
      {
        "id": 1952,
        "pid": 476,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0.01,
        "rss": 12263424,
        "vms": 6209536,
        "swap": 0,
        "memory_usage": 0.29,
        "cmd": [
          "C:\\Windows\\system32\\sppsvc.exe"
        ],
        "conns": 0
      },
      {
        "id": 1988,
        "pid": 476,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.03,
        "rss": 31141888,
        "vms": 65892352,
        "swap": 0,
        "memory_usage": 0.73,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "secsvcs"
        ],
        "conns": 0
      },
      {
        "id": 2112,
        "pid": 832,
        "user": "jkstack-PC\\Administrator",
        "cpu_usage": 0.02,
        "rss": 63660032,
        "vms": 40775680,
        "swap": 0,
        "memory_usage": 1.48,
        "cmd": [
          "C:\\Windows\\Explorer.EXE"
        ],
        "conns": 0
      },
      {
        "id": 2200,
        "pid": 612,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0.08,
        "rss": 7753728,
        "vms": 3604480,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "C:\\Windows\\system32\\wbem\\wmiprvse.exe"
        ],
        "conns": 0
      },
      {
        "id": 2236,
        "pid": 2112,
        "user": "jkstack-PC\\Administrator",
        "cpu_usage": 0.01,
        "rss": 13484032,
        "vms": 3874816,
        "swap": 0,
        "memory_usage": 0.31,
        "cmd": [
          "\"C:\\Windows\\system32\\taskmgr.exe\"",
          "/4"
        ],
        "conns": 0
      },
      {
        "id": 2320,
        "pid": 476,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.01,
        "rss": 36286464,
        "vms": 44646400,
        "swap": 0,
        "memory_usage": 0.84,
        "cmd": [
          "C:\\metrics-agent\\metrics.exe",
          "-conf",
          "C:\\metrics-agent\\agent.conf"
        ],
        "conns": 1
      },
      {
        "id": 2512,
        "pid": 476,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 4018176,
        "vms": 4419584,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "\"C:\\Program",
          "Files\\Windows",
          "Media",
          "Player\\wmpnetwk.exe\""
        ],
        "conns": 0
      },
      {
        "id": 2664,
        "pid": 2840,
        "user": "jkstack-PC\\Administrator",
        "cpu_usage": 0,
        "rss": 12181504,
        "vms": 5382144,
        "swap": 0,
        "memory_usage": 0.28,
        "cmd": [
          "wmic"
        ],
        "conns": 0
      },
      {
        "id": 2840,
        "pid": 2112,
        "user": "jkstack-PC\\Administrator",
        "cpu_usage": 0,
        "rss": 7065600,
        "vms": 3969024,
        "swap": 0,
        "memory_usage": 0.16,
        "cmd": [
          "\"C:\\Windows\\System32\\cmd.exe\"",
          ""
        ],
        "conns": 0
      },
      {
        "id": 2848,
        "pid": 1496,
        "user": "jkstack-PC\\Administrator",
        "cpu_usage": 0,
        "rss": 10518528,
        "vms": 13049856,
        "swap": 0,
        "memory_usage": 0.24,
        "cmd": [
          "\\??\\C:\\Windows\\system32\\conhost.exe"
        ],
        "conns": 0
      }
    ],
    "connections": [
      {
        "fd": 0,
        "pid": 688,
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
        "pid": 356,
        "type": "tcp4",
        "local": "0.0.0.0:3389",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 384,
        "type": "tcp4",
        "local": "0.0.0.0:49152",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 780,
        "type": "tcp4",
        "local": "0.0.0.0:49153",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 860,
        "type": "tcp4",
        "local": "0.0.0.0:49154",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 476,
        "type": "tcp4",
        "local": "0.0.0.0:49155",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 488,
        "type": "tcp4",
        "local": "0.0.0.0:49157",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "tcp4",
        "local": "192.168.2.71:139",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 2320,
        "type": "tcp4",
        "local": "192.168.2.71:49158",
        "remote": "192.168.2.52:13081",
        "status": "ESTABLISHED"
      },
      {
        "fd": 0,
        "pid": 688,
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
        "pid": 356,
        "type": "tcp6",
        "local": ":::3389",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 384,
        "type": "tcp6",
        "local": ":::49152",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 780,
        "type": "tcp6",
        "local": ":::49153",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 860,
        "type": "tcp6",
        "local": ":::49154",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 476,
        "type": "tcp6",
        "local": ":::49155",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 488,
        "type": "tcp6",
        "local": ":::49157",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 356,
        "type": "udp4",
        "local": "0.0.0.0:5355",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1920,
        "type": "udp4",
        "local": "127.0.0.1:1900",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1920,
        "type": "udp4",
        "local": "127.0.0.1:54946",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "udp4",
        "local": "192.168.2.71:137",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "udp4",
        "local": "192.168.2.71:138",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1920,
        "type": "udp4",
        "local": "192.168.2.71:1900",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1920,
        "type": "udp4",
        "local": "192.168.2.71:54945",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1920,
        "type": "udp6",
        "local": "::1:1900",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1920,
        "type": "udp6",
        "local": "::1:54944",
        "status": ""
      }
    ]
  }
}
```