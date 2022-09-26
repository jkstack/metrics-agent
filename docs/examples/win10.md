static:

```json
{
  "code": 200,
  "payload": {
    "time": 1663898964,
    "host": {
      "name": "DESKTOP-VNFDIVT",
      "uptime": 35913
    },
    "os": {
      "name": "windows",
      "platform_name": "Microsoft Windows 10 Enterprise LTSC 2021",
      "platform_version": "10.0.19044 Build 19044",
      "install": 1663797298,
      "startup": 1663863052
    },
    "kernel": {
      "version": "10.0.19044 Build 19044",
      "arch": "x86_64"
    },
    "cpu": {
      "physical": 2,
      "logical": 2,
      "cores": [
        {
          "processor": 0,
          "model": "Intel(R) Xeon(R) Silver 4214R CPU @ 2.40GHz",
          "core": 0,
          "cores": 1,
          "physical": 0,
          "mhz": 2394
        },
        {
          "processor": 1,
          "model": "Intel(R) Xeon(R) Silver 4214R CPU @ 2.40GHz",
          "core": 0,
          "cores": 1,
          "physical": 0,
          "mhz": 2394
        }
      ]
    },
    "memory": {
      "physical": 4294496256,
      "swap": 5770891264
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
        "total": 53032665088
      },
      {
        "mount": "D:",
        "fstype": "UDF",
        "opts": [
          "ro"
        ],
        "total": 5044211712
      }
    ],
    "gateway": "192.168.2.1",
    "interface": [
      {
        "index": 2,
        "name": "Ethernet0",
        "mtu": 1500,
        "flags": [
          "up",
          "broadcast",
          "multicast"
        ],
        "mac": "00:50:56:a3:a0:52",
        "addrs": [
          "fe80::7c0c:6417:eac:e1b4/64",
          "192.168.2.77/24"
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
        "usage": 1.84
      },
      "memory": {
        "used": 2121076736,
        "free": 2173419520,
        "available": 2173419520,
        "usage": 49,
        "swap_used": 0,
        "swap_free": 0
      },
      "partitions": [
        {
          "name": "C:",
          "used": 25344008192,
          "free": 27688656896,
          "usage": 47.79
        },
        {
          "name": "D:",
          "used": 5044211712,
          "free": 0,
          "usage": 100
        }
      ],
      "interface": [
        {
          "name": "Ethernet0",
          "bytes_sent": 3149286,
          "bytes_recv": 9490782,
          "packets_sent": 20603,
          "packets_recv": 38459
        },
        {
          "name": "Loopback Pseudo-Interface 1",
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
        "cpu_usage": 0.11,
        "rss": 155648,
        "vms": 200704,
        "swap": 0,
        "memory_usage": 0,
        "cmd": [
          ""
        ],
        "listen": [
          445,
          5357,
          139,
          445,
          5357
        ],
        "conns": 7
      },
      {
        "id": 92,
        "pid": 4,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 71397376,
        "vms": 1695744,
        "swap": 0,
        "memory_usage": 1.66,
        "cmd": [
          ""
        ],
        "conns": 0
      },
      {
        "id": 320,
        "pid": 4,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 1183744,
        "vms": 1081344,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          ""
        ],
        "conns": 0
      },
      {
        "id": 368,
        "pid": 556,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 58634240,
        "vms": 15425536,
        "swap": 0,
        "memory_usage": 1.37,
        "cmd": [
          "\"LogonUI.exe\"",
          "/flags:0x2",
          "/state0:0xa3b1a055",
          "/state1:0x41c64e6d"
        ],
        "conns": 0
      },
      {
        "id": 376,
        "pid": 556,
        "user": "Window Manager\\DWM-1",
        "cpu_usage": 0,
        "rss": 48762880,
        "vms": 28839936,
        "swap": 0,
        "memory_usage": 1.14,
        "cmd": [
          "\"dwm.exe\""
        ],
        "conns": 0
      },
      {
        "id": 408,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 10203136,
        "vms": 2088960,
        "swap": 0,
        "memory_usage": 0.24,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalSystemNetworkRestricted",
          "-p",
          "-s",
          "NcbService"
        ],
        "conns": 0
      },
      {
        "id": 420,
        "pid": 408,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 5488640,
        "vms": 1789952,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          ""
        ],
        "conns": 0
      },
      {
        "id": 424,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 5709824,
        "vms": 1282048,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalServiceNetworkRestricted",
          "-p",
          "-s",
          "lmhosts"
        ],
        "conns": 0
      },
      {
        "id": 492,
        "pid": 408,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 7389184,
        "vms": 1413120,
        "swap": 0,
        "memory_usage": 0.17,
        "cmd": [
          ""
        ],
        "listen": [
          49665,
          49665
        ],
        "conns": 2
      },
      {
        "id": 500,
        "pid": 484,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 5369856,
        "vms": 1773568,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          ""
        ],
        "conns": 0
      },
      {
        "id": 512,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 16637952,
        "vms": 3993600,
        "swap": 0,
        "memory_usage": 0.39,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalService",
          "-p",
          "-s",
          "CDPSvc"
        ],
        "listen": [
          5040
        ],
        "conns": 2
      },
      {
        "id": 556,
        "pid": 484,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 12087296,
        "vms": 2973696,
        "swap": 0,
        "memory_usage": 0.28,
        "cmd": [
          "winlogon.exe"
        ],
        "conns": 0
      },
      {
        "id": 624,
        "pid": 492,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.01,
        "rss": 10633216,
        "vms": 5259264,
        "swap": 0,
        "memory_usage": 0.25,
        "cmd": [
          ""
        ],
        "listen": [
          49671,
          49671
        ],
        "conns": 2
      },
      {
        "id": 636,
        "pid": 492,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.16,
        "rss": 19591168,
        "vms": 6623232,
        "swap": 0,
        "memory_usage": 0.46,
        "cmd": [
          "C:\\Windows\\system32\\lsass.exe"
        ],
        "listen": [
          49664,
          49664
        ],
        "conns": 2
      },
      {
        "id": 752,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 26009600,
        "vms": 8527872,
        "swap": 0,
        "memory_usage": 0.61,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "DcomLaunch",
          "-p"
        ],
        "conns": 0
      },
      {
        "id": 776,
        "pid": 556,
        "user": "Font Driver Host\\UMFD-1",
        "cpu_usage": 0,
        "rss": 8957952,
        "vms": 3211264,
        "swap": 0,
        "memory_usage": 0.21,
        "cmd": [
          "\"fontdrvhost.exe\""
        ],
        "conns": 0
      },
      {
        "id": 812,
        "pid": 492,
        "user": "Font Driver Host\\UMFD-0",
        "cpu_usage": 0,
        "rss": 4505600,
        "vms": 1683456,
        "swap": 0,
        "memory_usage": 0.1,
        "cmd": [
          "\"fontdrvhost.exe\""
        ],
        "conns": 0
      },
      {
        "id": 876,
        "pid": 624,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0.01,
        "rss": 13606912,
        "vms": 6488064,
        "swap": 0,
        "memory_usage": 0.32,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "RPCSS",
          "-p"
        ],
        "listen": [
          135,
          135
        ],
        "conns": 2
      },
      {
        "id": 892,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 8478720,
        "vms": 1728512,
        "swap": 0,
        "memory_usage": 0.2,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalServiceNetworkRestricted",
          "-p",
          "-s",
          "TimeBrokerSvc"
        ],
        "conns": 0
      },
      {
        "id": 896,
        "pid": 4300,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 28385280,
        "vms": 9400320,
        "swap": 0,
        "memory_usage": 0.66,
        "cmd": [
          "\"ctfmon.exe\""
        ],
        "conns": 0
      },
      {
        "id": 936,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 9248768,
        "vms": 2560000,
        "swap": 0,
        "memory_usage": 0.22,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "DcomLaunch",
          "-p",
          "-s",
          "LSM"
        ],
        "conns": 0
      },
      {
        "id": 948,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 15163392,
        "vms": 2949120,
        "swap": 0,
        "memory_usage": 0.35,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "netsvcs",
          "-p",
          "-s",
          "TokenBroker"
        ],
        "conns": 0
      },
      {
        "id": 1012,
        "pid": 624,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 13185024,
        "vms": 4378624,
        "swap": 0,
        "memory_usage": 0.31,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "NetworkService",
          "-s",
          "TermService"
        ],
        "listen": [
          3389,
          3389
        ],
        "conns": 4
      },
      {
        "id": 1112,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 4554752,
        "vms": 966656,
        "swap": 0,
        "memory_usage": 0.11,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "DcomLaunch",
          "-p",
          "-s",
          "DeviceInstall"
        ],
        "conns": 0
      },
      {
        "id": 1124,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 17571840,
        "vms": 14077952,
        "swap": 0,
        "memory_usage": 0.41,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalServiceNetworkRestricted",
          "-p",
          "-s",
          "EventLog"
        ],
        "listen": [
          49666,
          49666
        ],
        "conns": 2
      },
      {
        "id": 1168,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 8212480,
        "vms": 3637248,
        "swap": 0,
        "memory_usage": 0.19,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalService",
          "-p",
          "-s",
          "nsi"
        ],
        "conns": 0
      },
      {
        "id": 1200,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 7700480,
        "vms": 2068480,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalServiceNetworkRestricted",
          "-p",
          "-s",
          "Dhcp"
        ],
        "conns": 0
      },
      {
        "id": 1208,
        "pid": 944,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0.01,
        "rss": 102195200,
        "vms": 26185728,
        "swap": 0,
        "memory_usage": 2.38,
        "cmd": [
          "C:\\Windows\\Explorer.EXE"
        ],
        "conns": 0
      },
      {
        "id": 1212,
        "pid": 2288,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 25538560,
        "vms": 5103616,
        "swap": 0,
        "memory_usage": 0.59,
        "cmd": [
          "sihost.exe"
        ],
        "conns": 0
      },
      {
        "id": 1260,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 14381056,
        "vms": 3194880,
        "swap": 0,
        "memory_usage": 0.33,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "netsvcs",
          "-p",
          "-s",
          "ProfSvc"
        ],
        "conns": 0
      },
      {
        "id": 1268,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.21,
        "rss": 62623744,
        "vms": 51306496,
        "swap": 0,
        "memory_usage": 1.46,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalSystemNetworkRestricted",
          "-p",
          "-s",
          "SysMain"
        ],
        "conns": 0
      },
      {
        "id": 1276,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 8077312,
        "vms": 1884160,
        "swap": 0,
        "memory_usage": 0.19,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalService",
          "-p",
          "-s",
          "EventSystem"
        ],
        "conns": 0
      },
      {
        "id": 1288,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 6139904,
        "vms": 1261568,
        "swap": 0,
        "memory_usage": 0.14,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "netsvcs",
          "-p",
          "-s",
          "Themes"
        ],
        "conns": 0
      },
      {
        "id": 1392,
        "pid": 624,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 11681792,
        "vms": 4018176,
        "swap": 0,
        "memory_usage": 0.27,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "NetworkService",
          "-p",
          "-s",
          "NlaSvc"
        ],
        "conns": 0
      },
      {
        "id": 1408,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 8765440,
        "vms": 1941504,
        "swap": 0,
        "memory_usage": 0.2,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalServiceNetworkRestricted",
          "-s",
          "RmSvc"
        ],
        "conns": 0
      },
      {
        "id": 1444,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 8523776,
        "vms": 1839104,
        "swap": 0,
        "memory_usage": 0.2,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "netsvcs",
          "-p",
          "-s",
          "SENS"
        ],
        "conns": 0
      },
      {
        "id": 1488,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 7667712,
        "vms": 1519616,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalSystemNetworkRestricted",
          "-p",
          "-s",
          "AudioEndpointBuilder"
        ],
        "conns": 0
      },
      {
        "id": 1496,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 7700480,
        "vms": 1839104,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalService",
          "-p",
          "-s",
          "FontCache"
        ],
        "conns": 0
      },
      {
        "id": 1540,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.01,
        "rss": 18038784,
        "vms": 6254592,
        "swap": 0,
        "memory_usage": 0.42,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "netsvcs",
          "-p",
          "-s",
          "Schedule"
        ],
        "listen": [
          49667,
          49667
        ],
        "conns": 2
      },
      {
        "id": 1548,
        "pid": 1540,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 12296192,
        "vms": 2662400,
        "swap": 0,
        "memory_usage": 0.29,
        "cmd": [
          "taskhostw.exe"
        ],
        "conns": 0
      },
      {
        "id": 1576,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 9256960,
        "vms": 2015232,
        "swap": 0,
        "memory_usage": 0.22,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalServiceNetworkRestricted",
          "-p"
        ],
        "conns": 0
      },
      {
        "id": 1632,
        "pid": 4,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.01,
        "rss": 15196160,
        "vms": 229376,
        "swap": 0,
        "memory_usage": 0.35,
        "cmd": [
          ""
        ],
        "conns": 0
      },
      {
        "id": 1684,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 9465856,
        "vms": 2576384,
        "swap": 0,
        "memory_usage": 0.22,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalService",
          "-p",
          "-s",
          "netprofm"
        ],
        "conns": 0
      },
      {
        "id": 1700,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 7987200,
        "vms": 1617920,
        "swap": 0,
        "memory_usage": 0.19,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalSystemNetworkRestricted",
          "-p",
          "-s",
          "UmRdpService"
        ],
        "conns": 0
      },
      {
        "id": 1720,
        "pid": 624,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 8204288,
        "vms": 2748416,
        "swap": 0,
        "memory_usage": 0.19,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "NetworkService",
          "-p",
          "-s",
          "Dnscache"
        ],
        "conns": 4
      },
      {
        "id": 1728,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 6717440,
        "vms": 1540096,
        "swap": 0,
        "memory_usage": 0.16,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalServiceNetworkRestricted",
          "-p"
        ],
        "conns": 0
      },
      {
        "id": 1756,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 9973760,
        "vms": 2228224,
        "swap": 0,
        "memory_usage": 0.23,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalServiceNetworkRestricted",
          "-p"
        ],
        "conns": 0
      },
      {
        "id": 1800,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 9457664,
        "vms": 3854336,
        "swap": 0,
        "memory_usage": 0.22,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalSystemNetworkRestricted",
          "-p",
          "-s",
          "PcaSvc"
        ],
        "conns": 0
      },
      {
        "id": 1812,
        "pid": 752,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 8167424,
        "vms": 1462272,
        "swap": 0,
        "memory_usage": 0.19,
        "cmd": [
          "C:\\Windows\\System32\\InputMethod\\CHS\\ChsIME.exe",
          "-Embedding"
        ],
        "conns": 0
      },
      {
        "id": 1832,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 13025280,
        "vms": 2097152,
        "swap": 0,
        "memory_usage": 0.3,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "netsvcs",
          "-p",
          "-s",
          "ShellHWDetection"
        ],
        "conns": 0
      },
      {
        "id": 2012,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.02,
        "rss": 26411008,
        "vms": 8048640,
        "swap": 0,
        "memory_usage": 0.61,
        "cmd": [
          "C:\\Windows\\System32\\spoolsv.exe"
        ],
        "listen": [
          49668,
          49668
        ],
        "conns": 2
      },
      {
        "id": 2020,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 7716864,
        "vms": 1716224,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalServiceNetworkRestricted",
          "-p",
          "-s",
          "WinHttpAutoProxySvc"
        ],
        "conns": 0
      },
      {
        "id": 2052,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 10354688,
        "vms": 2224128,
        "swap": 0,
        "memory_usage": 0.24,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalServiceNoNetwork",
          "-p",
          "-s",
          "NcdAutoSetup"
        ],
        "conns": 0
      },
      {
        "id": 2064,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 16855040,
        "vms": 7340032,
        "swap": 0,
        "memory_usage": 0.39,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalServiceNoNetworkFirewall",
          "-p"
        ],
        "conns": 0
      },
      {
        "id": 2076,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 7557120,
        "vms": 1552384,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalService",
          "-p",
          "-s",
          "fdPHost"
        ],
        "conns": 0
      },
      {
        "id": 2172,
        "pid": 624,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 8486912,
        "vms": 2002944,
        "swap": 0,
        "memory_usage": 0.2,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "NetworkService",
          "-p",
          "-s",
          "LanmanWorkstation"
        ],
        "conns": 0
      },
      {
        "id": 2288,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 9883648,
        "vms": 2379776,
        "swap": 0,
        "memory_usage": 0.23,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "netsvcs",
          "-p",
          "-s",
          "UserManager"
        ],
        "conns": 0
      },
      {
        "id": 2364,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 12075008,
        "vms": 2756608,
        "swap": 0,
        "memory_usage": 0.28,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "WbioSvcGroup",
          "-s",
          "WbioSrvc"
        ],
        "conns": 0
      },
      {
        "id": 2368,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 7241728,
        "vms": 1597440,
        "swap": 0,
        "memory_usage": 0.17,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "netsvcs",
          "-s",
          "CertPropSvc"
        ],
        "conns": 0
      },
      {
        "id": 2440,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 9179136,
        "vms": 2117632,
        "swap": 0,
        "memory_usage": 0.21,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "netsvcs",
          "-p",
          "-s",
          "SessionEnv"
        ],
        "listen": [
          49669,
          49669
        ],
        "conns": 2
      },
      {
        "id": 2456,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.02,
        "rss": 21626880,
        "vms": 10829824,
        "swap": 0,
        "memory_usage": 0.5,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "netsvcs",
          "-p",
          "-s",
          "Winmgmt"
        ],
        "conns": 0
      },
      {
        "id": 2548,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 10072064,
        "vms": 2744320,
        "swap": 0,
        "memory_usage": 0.23,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalServiceAndNoImpersonation",
          "-p",
          "-s",
          "FDResPub"
        ],
        "conns": 6
      },
      {
        "id": 2592,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 8327168,
        "vms": 2600960,
        "swap": 0,
        "memory_usage": 0.19,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "netsvcs",
          "-p",
          "-s",
          "IKEEXT"
        ],
        "conns": 4
      },
      {
        "id": 2600,
        "pid": 624,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 7589888,
        "vms": 1683456,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "NetworkServiceNetworkRestricted",
          "-p",
          "-s",
          "PolicyAgent"
        ],
        "listen": [
          49670,
          49670
        ],
        "conns": 2
      },
      {
        "id": 2644,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 6406144,
        "vms": 1421312,
        "swap": 0,
        "memory_usage": 0.15,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalServiceNoNetwork",
          "-p"
        ],
        "conns": 0
      },
      {
        "id": 2652,
        "pid": 624,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0.11,
        "rss": 15097856,
        "vms": 4108288,
        "swap": 0,
        "memory_usage": 0.35,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "NetworkService",
          "-p",
          "-s",
          "CryptSvc"
        ],
        "conns": 0
      },
      {
        "id": 2660,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.02,
        "rss": 37150720,
        "vms": 20316160,
        "swap": 0,
        "memory_usage": 0.87,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "utcsvc",
          "-p"
        ],
        "conns": 0
      },
      {
        "id": 2680,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 11223040,
        "vms": 2740224,
        "swap": 0,
        "memory_usage": 0.26,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "NetSvcs",
          "-p",
          "-s",
          "iphlpsvc"
        ],
        "conns": 1
      },
      {
        "id": 2688,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 31297536,
        "vms": 26263552,
        "swap": 0,
        "memory_usage": 0.73,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalServiceNoNetwork",
          "-p",
          "-s",
          "DPS"
        ],
        "conns": 0
      },
      {
        "id": 2756,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 7618560,
        "vms": 1409024,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalService",
          "-p",
          "-s",
          "DispBrokerDesktopSvc"
        ],
        "conns": 0
      },
      {
        "id": 2772,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 7098368,
        "vms": 1593344,
        "swap": 0,
        "memory_usage": 0.17,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalService",
          "-p",
          "-s",
          "SstpSvc"
        ],
        "conns": 0
      },
      {
        "id": 2796,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 5984256,
        "vms": 1294336,
        "swap": 0,
        "memory_usage": 0.14,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalSystemNetworkRestricted",
          "-p",
          "-s",
          "TrkWks"
        ],
        "conns": 0
      },
      {
        "id": 2812,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 20570112,
        "vms": 4472832,
        "swap": 0,
        "memory_usage": 0.48,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "netsvcs",
          "-p",
          "-s",
          "WpnService"
        ],
        "conns": 1
      },
      {
        "id": 2824,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 9539584,
        "vms": 2420736,
        "swap": 0,
        "memory_usage": 0.22,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "netsvcs",
          "-p",
          "-s",
          "LanmanServer"
        ],
        "conns": 0
      },
      {
        "id": 2880,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.03,
        "rss": 36294656,
        "vms": 42979328,
        "swap": 0,
        "memory_usage": 0.85,
        "cmd": [
          "C:\\metrics-agent\\metrics.exe",
          "-conf",
          "C:\\metrics-agent\\agent.conf"
        ],
        "conns": 1
      },
      {
        "id": 2888,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 1.9,
        "rss": 261435392,
        "vms": 294465536,
        "swap": 0,
        "memory_usage": 6.09,
        "cmd": [
          ""
        ],
        "conns": 0
      },
      {
        "id": 3028,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 5775360,
        "vms": 1273856,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalService",
          "-p",
          "-s",
          "WdiServiceHost"
        ],
        "conns": 0
      },
      {
        "id": 3132,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 19484672,
        "vms": 3846144,
        "swap": 0,
        "memory_usage": 0.45,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalService",
          "-p",
          "-s",
          "LicenseManager"
        ],
        "conns": 0
      },
      {
        "id": 3180,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 13262848,
        "vms": 3534848,
        "swap": 0,
        "memory_usage": 0.31,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "netsvcs"
        ],
        "conns": 0
      },
      {
        "id": 3340,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 6225920,
        "vms": 1486848,
        "swap": 0,
        "memory_usage": 0.14,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalSystemNetworkRestricted",
          "-p",
          "-s",
          "WdiSystemHost"
        ],
        "conns": 0
      },
      {
        "id": 3476,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 14757888,
        "vms": 2457600,
        "swap": 0,
        "memory_usage": 0.34,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "netsvcs",
          "-p"
        ],
        "conns": 0
      },
      {
        "id": 3512,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 7348224,
        "vms": 4255744,
        "swap": 0,
        "memory_usage": 0.17,
        "cmd": [
          ""
        ],
        "conns": 0
      },
      {
        "id": 3636,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 6656000,
        "vms": 1314816,
        "swap": 0,
        "memory_usage": 0.15,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalSystemNetworkRestricted",
          "-p",
          "-s",
          "DeviceAssociationService"
        ],
        "conns": 0
      },
      {
        "id": 3672,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.09,
        "rss": 16781312,
        "vms": 8351744,
        "swap": 0,
        "memory_usage": 0.39,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "netsvcs",
          "-p",
          "-s",
          "wuauserv"
        ],
        "conns": 0
      },
      {
        "id": 3676,
        "pid": 3636,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 12681216,
        "vms": 3260416,
        "swap": 0,
        "memory_usage": 0.3,
        "cmd": [
          "dashost.exe",
          "{aeeeaa4e-9fb0-421e-a8cb3f16e6034ea0}"
        ],
        "conns": 6
      },
      {
        "id": 3880,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 7892992,
        "vms": 2019328,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalServiceAndNoImpersonation",
          "-p",
          "-s",
          "SSDPSRV"
        ],
        "conns": 8
      },
      {
        "id": 4300,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 8318976,
        "vms": 1777664,
        "swap": 0,
        "memory_usage": 0.19,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalSystemNetworkRestricted",
          "-p",
          "-s",
          "TabletInputService"
        ],
        "conns": 0
      },
      {
        "id": 4520,
        "pid": 1540,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 13574144,
        "vms": 4378624,
        "swap": 0,
        "memory_usage": 0.32,
        "cmd": [
          "taskhostw.exe",
          "{222A245B-E637-4AE9-A93F-A59CA119A75E}"
        ],
        "conns": 0
      },
      {
        "id": 4808,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 11751424,
        "vms": 3846144,
        "swap": 0,
        "memory_usage": 0.27,
        "cmd": [
          ""
        ],
        "conns": 0
      },
      {
        "id": 4856,
        "pid": 624,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 32616448,
        "vms": 7139328,
        "swap": 0,
        "memory_usage": 0.76,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "UnistackSvcGroup",
          "-s",
          "WpnUserService"
        ],
        "conns": 0
      },
      {
        "id": 4872,
        "pid": 752,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 1540096,
        "vms": 18157568,
        "swap": 0,
        "memory_usage": 0.04,
        "cmd": [
          "\"C:\\Windows\\ImmersiveControlPanel\\SystemSettings.exe\"",
          "-ServerName:microsoft.windows.immersivecontrolpanel"
        ],
        "conns": 0
      },
      {
        "id": 4948,
        "pid": 752,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 44429312,
        "vms": 10379264,
        "swap": 0,
        "memory_usage": 1.03,
        "cmd": [
          "\"C:\\Windows\\SystemApps\\ShellExperienceHost_cw5n1h2txyewy\\ShellExperienceHost.exe\"",
          "-ServerName:App.AppXtk181tbxbce2qsex02s8tw7hfxa9xb3t.mca"
        ],
        "conns": 0
      },
      {
        "id": 5000,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.01,
        "rss": 14385152,
        "vms": 5636096,
        "swap": 0,
        "memory_usage": 0.33,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "appmodel",
          "-p",
          "-s",
          "StateRepository"
        ],
        "conns": 0
      },
      {
        "id": 5016,
        "pid": 624,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 17301504,
        "vms": 3743744,
        "swap": 0,
        "memory_usage": 0.4,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "UnistackSvcGroup",
          "-s",
          "CDPUserSvc"
        ],
        "conns": 0
      },
      {
        "id": 5240,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.01,
        "rss": 26546176,
        "vms": 18210816,
        "swap": 0,
        "memory_usage": 0.62,
        "cmd": [
          "C:\\Windows\\system32\\SearchIndexer.exe",
          "/Embedding"
        ],
        "conns": 0
      },
      {
        "id": 5260,
        "pid": 1208,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 44843008,
        "vms": 21258240,
        "swap": 0,
        "memory_usage": 1.04,
        "cmd": [
          "\"C:\\Windows\\System32\\Taskmgr.exe\"",
          ""
        ],
        "conns": 0
      },
      {
        "id": 5392,
        "pid": 624,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 17559552,
        "vms": 3047424,
        "swap": 0,
        "memory_usage": 0.41,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "ClipboardSvcGroup",
          "-p",
          "-s",
          "cbdhsvc"
        ],
        "conns": 0
      },
      {
        "id": 5492,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 6434816,
        "vms": 1286144,
        "swap": 0,
        "memory_usage": 0.15,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "netsvcs",
          "-p",
          "-s",
          "Appinfo"
        ],
        "conns": 0
      },
      {
        "id": 5604,
        "pid": 752,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 27054080,
        "vms": 7778304,
        "swap": 0,
        "memory_usage": 0.63,
        "cmd": [
          "C:\\Windows\\system32\\ApplicationFrameHost.exe",
          "-Embedding"
        ],
        "conns": 0
      },
      {
        "id": 5660,
        "pid": 752,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 60489728,
        "vms": 16601088,
        "swap": 0,
        "memory_usage": 1.41,
        "cmd": [
          "\"C:\\Windows\\SystemApps\\Microsoft.Windows.StartMenuExperienceHost_cw5n1h2txyewy\\StartMenuExperienceHost.exe\"",
          "-ServerName:App.AppXywbrabmsek0gm3tkwpr5kwzbs55tkqay.mca"
        ],
        "conns": 0
      },
      {
        "id": 5732,
        "pid": 752,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 22192128,
        "vms": 3547136,
        "swap": 0,
        "memory_usage": 0.52,
        "cmd": [
          "C:\\Windows\\System32\\RuntimeBroker.exe",
          "-Embedding"
        ],
        "conns": 0
      },
      {
        "id": 5764,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 10203136,
        "vms": 2478080,
        "swap": 0,
        "memory_usage": 0.24,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "netsvcs",
          "-p",
          "-s",
          "UsoSvc"
        ],
        "conns": 0
      },
      {
        "id": 5852,
        "pid": 624,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 12513280,
        "vms": 2797568,
        "swap": 0,
        "memory_usage": 0.29,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "UnistackSvcGroup"
        ],
        "conns": 0
      },
      {
        "id": 5860,
        "pid": 752,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.03,
        "rss": 30171136,
        "vms": 15822848,
        "swap": 0,
        "memory_usage": 0.7,
        "cmd": [
          "C:\\Windows\\System32\\mousocoreworker.exe",
          "-Embedding"
        ],
        "conns": 0
      },
      {
        "id": 5976,
        "pid": 752,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0.01,
        "rss": 135290880,
        "vms": 71757824,
        "swap": 0,
        "memory_usage": 3.15,
        "cmd": [
          "\"C:\\Windows\\SystemApps\\Microsoft.Windows.Search_cw5n1h2txyewy\\SearchApp.exe\"",
          "-ServerName:CortanaUI.AppX8z9r6jm96hw4bsbneegw0kyxx296wr9t.mca"
        ],
        "conns": 0
      },
      {
        "id": 6120,
        "pid": 752,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0.01,
        "rss": 29495296,
        "vms": 7499776,
        "swap": 0,
        "memory_usage": 0.69,
        "cmd": [
          "C:\\Windows\\System32\\RuntimeBroker.exe",
          "-Embedding"
        ],
        "conns": 0
      },
      {
        "id": 6532,
        "pid": 752,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 45899776,
        "vms": 10223616,
        "swap": 0,
        "memory_usage": 1.07,
        "cmd": [
          "\"C:\\Windows\\SystemApps\\Microsoft.LockApp_cw5n1h2txyewy\\LockApp.exe\"",
          "-ServerName:WindowsDefaultLockScreen.AppX7y4nbzq37zn4ks9k7amqjywdat7d3j2z.mca"
        ],
        "conns": 0
      },
      {
        "id": 6572,
        "pid": 752,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 27238400,
        "vms": 5382144,
        "swap": 0,
        "memory_usage": 0.63,
        "cmd": [
          "C:\\Windows\\System32\\RuntimeBroker.exe",
          "-Embedding"
        ],
        "conns": 0
      },
      {
        "id": 6712,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 11390976,
        "vms": 1572864,
        "swap": 0,
        "memory_usage": 0.27,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalService",
          "-p",
          "-s",
          "BthAvctpSvc"
        ],
        "conns": 0
      },
      {
        "id": 6764,
        "pid": 624,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 10502144,
        "vms": 2600960,
        "swap": 0,
        "memory_usage": 0.24,
        "cmd": [
          ""
        ],
        "conns": 0
      },
      {
        "id": 6784,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 12115968,
        "vms": 2912256,
        "swap": 0,
        "memory_usage": 0.28,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalSystemNetworkRestricted",
          "-p",
          "-s",
          "StorSvc"
        ],
        "conns": 0
      },
      {
        "id": 6788,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 10629120,
        "vms": 6361088,
        "swap": 0,
        "memory_usage": 0.25,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalSystemNetworkRestricted",
          "-p",
          "-s",
          "DsSvc"
        ],
        "conns": 0
      },
      {
        "id": 6844,
        "pid": 752,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0.1,
        "rss": 11091968,
        "vms": 3133440,
        "swap": 0,
        "memory_usage": 0.26,
        "cmd": [
          "C:\\Windows\\system32\\wbem\\wmiprvse.exe"
        ],
        "conns": 0
      },
      {
        "id": 6916,
        "pid": 752,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 13852672,
        "vms": 2387968,
        "swap": 0,
        "memory_usage": 0.32,
        "cmd": [
          "C:\\Windows\\System32\\RuntimeBroker.exe",
          "-Embedding"
        ],
        "conns": 0
      },
      {
        "id": 6960,
        "pid": 1208,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 9740288,
        "vms": 1724416,
        "swap": 0,
        "memory_usage": 0.23,
        "cmd": [
          "\"C:\\Windows\\System32\\SecurityHealthSystray.exe\"",
          ""
        ],
        "conns": 0
      },
      {
        "id": 6972,
        "pid": 752,
        "user": "DESKTOP-VNFDIVT\\jkstack",
        "cpu_usage": 0,
        "rss": 16687104,
        "vms": 2506752,
        "swap": 0,
        "memory_usage": 0.39,
        "cmd": [
          "C:\\Windows\\System32\\RuntimeBroker.exe",
          "-Embedding"
        ],
        "conns": 0
      },
      {
        "id": 6988,
        "pid": 624,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 16252928,
        "vms": 4325376,
        "swap": 0,
        "memory_usage": 0.38,
        "cmd": [
          ""
        ],
        "conns": 0
      },
      {
        "id": 7096,
        "pid": 624,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0.01,
        "rss": 18505728,
        "vms": 6213632,
        "swap": 0,
        "memory_usage": 0.43,
        "cmd": [
          ""
        ],
        "listen": [
          7680,
          7680
        ],
        "conns": 3
      }
    ],
    "connections": [
      {
        "fd": 0,
        "pid": 876,
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
        "pid": 1012,
        "type": "tcp4",
        "local": "0.0.0.0:3389",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 512,
        "type": "tcp4",
        "local": "0.0.0.0:5040",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "tcp4",
        "local": "0.0.0.0:5357",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 7096,
        "type": "tcp4",
        "local": "0.0.0.0:7680",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 636,
        "type": "tcp4",
        "local": "0.0.0.0:49664",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 492,
        "type": "tcp4",
        "local": "0.0.0.0:49665",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 1124,
        "type": "tcp4",
        "local": "0.0.0.0:49666",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 1540,
        "type": "tcp4",
        "local": "0.0.0.0:49667",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 2012,
        "type": "tcp4",
        "local": "0.0.0.0:49668",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 2440,
        "type": "tcp4",
        "local": "0.0.0.0:49669",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 2600,
        "type": "tcp4",
        "local": "0.0.0.0:49670",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 624,
        "type": "tcp4",
        "local": "0.0.0.0:49671",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "tcp4",
        "local": "192.168.2.77:139",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 2880,
        "type": "tcp4",
        "local": "192.168.2.77:49672",
        "remote": "192.168.2.52:13081",
        "status": "ESTABLISHED"
      },
      {
        "fd": 0,
        "pid": 2812,
        "type": "tcp4",
        "local": "192.168.2.77:49698",
        "remote": "20.197.71.89:443",
        "status": "ESTABLISHED"
      },
      {
        "fd": 0,
        "pid": 7096,
        "type": "tcp4",
        "local": "192.168.2.77:49928",
        "remote": "20.54.24.246:443",
        "status": "ESTABLISHED"
      },
      {
        "fd": 0,
        "pid": 876,
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
        "pid": 1012,
        "type": "tcp6",
        "local": ":::3389",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "tcp6",
        "local": ":::5357",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 7096,
        "type": "tcp6",
        "local": ":::7680",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 636,
        "type": "tcp6",
        "local": ":::49664",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 492,
        "type": "tcp6",
        "local": ":::49665",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 1124,
        "type": "tcp6",
        "local": ":::49666",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 1540,
        "type": "tcp6",
        "local": ":::49667",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 2012,
        "type": "tcp6",
        "local": ":::49668",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 2440,
        "type": "tcp6",
        "local": ":::49669",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 2600,
        "type": "tcp6",
        "local": ":::49670",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 624,
        "type": "tcp6",
        "local": ":::49671",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 2592,
        "type": "udp4",
        "local": "0.0.0.0:500",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1012,
        "type": "udp4",
        "local": "0.0.0.0:3389",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 2548,
        "type": "udp4",
        "local": "0.0.0.0:3702",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 2548,
        "type": "udp4",
        "local": "0.0.0.0:3702",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 3676,
        "type": "udp4",
        "local": "0.0.0.0:3702",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 3676,
        "type": "udp4",
        "local": "0.0.0.0:3702",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 2592,
        "type": "udp4",
        "local": "0.0.0.0:4500",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 512,
        "type": "udp4",
        "local": "0.0.0.0:5050",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1720,
        "type": "udp4",
        "local": "0.0.0.0:5353",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1720,
        "type": "udp4",
        "local": "0.0.0.0:5355",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 2548,
        "type": "udp4",
        "local": "0.0.0.0:56151",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 3676,
        "type": "udp4",
        "local": "0.0.0.0:56157",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 3880,
        "type": "udp4",
        "local": "127.0.0.1:1900",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 2680,
        "type": "udp4",
        "local": "127.0.0.1:56150",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 3880,
        "type": "udp4",
        "local": "127.0.0.1:56156",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "udp4",
        "local": "192.168.2.77:137",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "udp4",
        "local": "192.168.2.77:138",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 3880,
        "type": "udp4",
        "local": "192.168.2.77:1900",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 3880,
        "type": "udp4",
        "local": "192.168.2.77:56155",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 2592,
        "type": "udp6",
        "local": ":::500",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1012,
        "type": "udp6",
        "local": ":::3389",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 2548,
        "type": "udp6",
        "local": ":::3702",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 3676,
        "type": "udp6",
        "local": ":::3702",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 3676,
        "type": "udp6",
        "local": ":::3702",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 2548,
        "type": "udp6",
        "local": ":::3702",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 2592,
        "type": "udp6",
        "local": ":::4500",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1720,
        "type": "udp6",
        "local": ":::5353",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1720,
        "type": "udp6",
        "local": ":::5355",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 2548,
        "type": "udp6",
        "local": ":::56152",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 3676,
        "type": "udp6",
        "local": ":::56158",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 3880,
        "type": "udp6",
        "local": "::1:1900",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 3880,
        "type": "udp6",
        "local": "::1:56154",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 3880,
        "type": "udp6",
        "local": "fe80::7c0c:6417:eac:e1b4:1900",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 3880,
        "type": "udp6",
        "local": "fe80::7c0c:6417:eac:e1b4:56153",
        "status": ""
      }
    ]
  }
}
```