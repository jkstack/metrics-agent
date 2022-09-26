static:

```json
{
  "code": 200,
  "payload": {
    "time": 1664185987,
    "host": {
      "name": "WIN-5UGRJ67RCBK",
      "uptime": 522203
    },
    "os": {
      "name": "windows",
      "platform_name": "Microsoft Windows Server 2016 Datacenter",
      "platform_version": "10.0.14393 Build 14393",
      "install": 1663634041,
      "startup": 1663663787
    },
    "kernel": {
      "version": "10.0.14393 Build 14393",
      "arch": "x86_64"
    },
    "cpu": {
      "physical": 4,
      "logical": 4,
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
        },
        {
          "processor": 2,
          "model": "Intel(R) Xeon(R) CPU E5-2670 v3 @ 2.30GHz",
          "core": 0,
          "cores": 1,
          "physical": 0,
          "mhz": 2297
        },
        {
          "processor": 3,
          "model": "Intel(R) Xeon(R) CPU E5-2670 v3 @ 2.30GHz",
          "core": 0,
          "cores": 1,
          "physical": 0,
          "mhz": 2297
        }
      ]
    },
    "memory": {
      "physical": 8589463552,
      "swap": 10602729472
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
        "total": 53160701952
      }
    ],
    "gateway": "192.168.2.1",
    "interface": [
      {
        "index": 3,
        "name": "Ethernet0",
        "mtu": 1500,
        "flags": [
          "up",
          "broadcast",
          "multicast"
        ],
        "mac": "00:50:56:a3:05:0b",
        "addrs": [
          "192.168.2.75/24"
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
        "index": 4,
        "name": "isatap.{56CDED85-8080-41A6-B2D0-80671CA898FD}",
        "mtu": 1280,
        "flags": [
          "pointtopoint",
          "multicast"
        ],
        "mac": "00:00:00:00:00:00:00:e0",
        "addrs": [
          "fe80::5efe:c0a8:24b/128"
        ]
      },
      {
        "index": 2,
        "name": "Teredo Tunneling Pseudo-Interface",
        "mtu": 1280,
        "flags": [
          "up",
          "pointtopoint",
          "multicast"
        ],
        "mac": "00:00:00:00:00:00:00:e0",
        "addrs": [
          "2001:0:348b:fb58:3895:c6d9:4b56:1aed/64",
          "fe80::3895:c6d9:4b56:1aed/64"
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
        "usage": 0.02
      },
      "memory": {
        "used": 1336004608,
        "free": 7253458944,
        "available": 7253458944,
        "usage": 15,
        "swap_used": 0,
        "swap_free": 0
      },
      "partitions": [
        {
          "name": "C:",
          "used": 13579014144,
          "free": 39581687808,
          "usage": 25.54
        }
      ],
      "interface": [
        {
          "name": "Ethernet0",
          "bytes_sent": 39463824,
          "bytes_recv": 1809069666,
          "packets_sent": 371516,
          "packets_recv": 1659672
        },
        {
          "name": "Loopback Pseudo-Interface 1",
          "bytes_sent": 0,
          "bytes_recv": 0,
          "packets_sent": 0,
          "packets_recv": 0
        },
        {
          "name": "isatap.{56CDED85-8080-41A6-B2D0-80671CA898FD}",
          "bytes_sent": 0,
          "bytes_recv": 0,
          "packets_sent": 0,
          "packets_recv": 0
        },
        {
          "name": "Teredo Tunneling Pseudo-Interface",
          "bytes_sent": 2174650,
          "bytes_recv": 1829528,
          "packets_sent": 18434,
          "packets_recv": 19008
        }
      ]
    },
    "process": [
      {
        "id": 4,
        "pid": 0,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.01,
        "rss": 143360,
        "vms": 131072,
        "swap": 0,
        "memory_usage": 0,
        "cmd": [
          ""
        ],
        "listen": [
          445,
          5985,
          47001,
          139,
          445,
          5985,
          47001
        ],
        "conns": 9
      },
      {
        "id": 32,
        "pid": 956,
        "user": "WIN-5UGRJ67RCBK\\Administrator",
        "cpu_usage": 0,
        "rss": 12873728,
        "vms": 3035136,
        "swap": 0,
        "memory_usage": 0.15,
        "cmd": [
          "\"C:\\Windows\\system32\\cmd.exe\"",
          ""
        ],
        "conns": 0
      },
      {
        "id": 156,
        "pid": 2396,
        "user": "Window Manager\\DWM-2",
        "cpu_usage": 0,
        "rss": 55611392,
        "vms": 20992000,
        "swap": 0,
        "memory_usage": 0.65,
        "cmd": [
          "\"dwm.exe\""
        ],
        "conns": 0
      },
      {
        "id": 292,
        "pid": 4,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 1220608,
        "vms": 393216,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          ""
        ],
        "conns": 0
      },
      {
        "id": 316,
        "pid": 580,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 6991872,
        "vms": 1548288,
        "swap": 0,
        "memory_usage": 0.08,
        "cmd": [
          "C:\\Windows\\system32\\vm3dservice.exe"
        ],
        "conns": 0
      },
      {
        "id": 368,
        "pid": 360,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 4841472,
        "vms": 1982464,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          ""
        ],
        "conns": 0
      },
      {
        "id": 444,
        "pid": 432,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 4849664,
        "vms": 1945600,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          ""
        ],
        "conns": 0
      },
      {
        "id": 468,
        "pid": 360,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 5652480,
        "vms": 1073152,
        "swap": 0,
        "memory_usage": 0.07,
        "cmd": [
          ""
        ],
        "listen": [
          49664,
          49664
        ],
        "conns": 2
      },
      {
        "id": 504,
        "pid": 432,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 9052160,
        "vms": 1896448,
        "swap": 0,
        "memory_usage": 0.11,
        "cmd": [
          "winlogon.exe"
        ],
        "conns": 0
      },
      {
        "id": 532,
        "pid": 360,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 16392192,
        "vms": 2277376,
        "swap": 0,
        "memory_usage": 0.19,
        "cmd": [
          ""
        ],
        "conns": 0
      },
      {
        "id": 580,
        "pid": 468,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 7647232,
        "vms": 3198976,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          ""
        ],
        "listen": [
          49668,
          49668
        ],
        "conns": 2
      },
      {
        "id": 588,
        "pid": 468,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 16601088,
        "vms": 7061504,
        "swap": 0,
        "memory_usage": 0.19,
        "cmd": [
          "C:\\Windows\\system32\\lsass.exe"
        ],
        "listen": [
          49671,
          49671
        ],
        "conns": 2
      },
      {
        "id": 700,
        "pid": 580,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 20647936,
        "vms": 6172672,
        "swap": 0,
        "memory_usage": 0.24,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "DcomLaunch"
        ],
        "conns": 0
      },
      {
        "id": 764,
        "pid": 580,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 10194944,
        "vms": 4616192,
        "swap": 0,
        "memory_usage": 0.12,
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
        "id": 856,
        "pid": 580,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.04,
        "rss": 66244608,
        "vms": 38928384,
        "swap": 0,
        "memory_usage": 0.77,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "netsvcs"
        ],
        "listen": [
          49666,
          49666
        ],
        "conns": 9
      },
      {
        "id": 868,
        "pid": 580,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 21688320,
        "vms": 12025856,
        "swap": 0,
        "memory_usage": 0.25,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "termsvcs"
        ],
        "listen": [
          3389,
          3389
        ],
        "conns": 4
      },
      {
        "id": 928,
        "pid": 580,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 22163456,
        "vms": 12673024,
        "swap": 0,
        "memory_usage": 0.26,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalSystemNetworkRestricted"
        ],
        "conns": 0
      },
      {
        "id": 936,
        "pid": 856,
        "user": "WIN-5UGRJ67RCBK\\Administrator",
        "cpu_usage": 0,
        "rss": 20824064,
        "vms": 4390912,
        "swap": 0,
        "memory_usage": 0.24,
        "cmd": [
          "sihost.exe"
        ],
        "conns": 0
      },
      {
        "id": 956,
        "pid": 700,
        "user": "WIN-5UGRJ67RCBK\\Administrator",
        "cpu_usage": 0,
        "rss": 29360128,
        "vms": 11247616,
        "swap": 0,
        "memory_usage": 0.34,
        "cmd": [
          "C:\\Windows\\System32\\RuntimeBroker.exe",
          "-Embedding"
        ],
        "conns": 0
      },
      {
        "id": 968,
        "pid": 580,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 19337216,
        "vms": 8769536,
        "swap": 0,
        "memory_usage": 0.23,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalService"
        ],
        "conns": 3
      },
      {
        "id": 972,
        "pid": 504,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 45383680,
        "vms": 11726848,
        "swap": 0,
        "memory_usage": 0.53,
        "cmd": [
          "\"LogonUI.exe\"",
          "/flags:0x0",
          "/state0:0xa3b64855",
          "/state1:0x41c64e6d"
        ],
        "conns": 0
      },
      {
        "id": 980,
        "pid": 504,
        "user": "Window Manager\\DWM-1",
        "cpu_usage": 0,
        "rss": 31428608,
        "vms": 14024704,
        "swap": 0,
        "memory_usage": 0.37,
        "cmd": [
          "\"dwm.exe\""
        ],
        "conns": 0
      },
      {
        "id": 1004,
        "pid": 580,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 21577728,
        "vms": 12615680,
        "swap": 0,
        "memory_usage": 0.25,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "LocalServiceNetworkRestricted"
        ],
        "listen": [
          49665,
          49665
        ],
        "conns": 2
      },
      {
        "id": 1088,
        "pid": 580,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 24903680,
        "vms": 9994240,
        "swap": 0,
        "memory_usage": 0.29,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "NetworkService"
        ],
        "conns": 2
      },
      {
        "id": 1176,
        "pid": 580,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 7348224,
        "vms": 1835008,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalServiceNetworkRestricted"
        ],
        "conns": 0
      },
      {
        "id": 1256,
        "pid": 700,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0.08,
        "rss": 11431936,
        "vms": 3747840,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          "C:\\Windows\\system32\\wbem\\wmiprvse.exe"
        ],
        "conns": 0
      },
      {
        "id": 1272,
        "pid": 580,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 17072128,
        "vms": 10653696,
        "swap": 0,
        "memory_usage": 0.2,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalServiceNoNetwork"
        ],
        "conns": 0
      },
      {
        "id": 1288,
        "pid": 856,
        "user": "WIN-5UGRJ67RCBK\\Administrator",
        "cpu_usage": 0,
        "rss": 20467712,
        "vms": 6230016,
        "swap": 0,
        "memory_usage": 0.24,
        "cmd": [
          "taskhostw.exe",
          "{222A245B-E637-4AE9-A93F-A59CA119A75E}"
        ],
        "conns": 0
      },
      {
        "id": 1300,
        "pid": 580,
        "user": "WIN-5UGRJ67RCBK\\Administrator",
        "cpu_usage": 0,
        "rss": 20402176,
        "vms": 4464640,
        "swap": 0,
        "memory_usage": 0.24,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "UnistackSvcGroup"
        ],
        "conns": 0
      },
      {
        "id": 1360,
        "pid": 580,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 7335936,
        "vms": 1822720,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalServiceAndNoImpersonation"
        ],
        "conns": 6
      },
      {
        "id": 1448,
        "pid": 580,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 7065600,
        "vms": 1634304,
        "swap": 0,
        "memory_usage": 0.08,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "NetworkServiceNetworkRestricted"
        ],
        "listen": [
          49669,
          49669
        ],
        "conns": 2
      },
      {
        "id": 1500,
        "pid": 580,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 9928704,
        "vms": 2527232,
        "swap": 0,
        "memory_usage": 0.12,
        "cmd": [
          "C:\\Windows\\System32\\msdtc.exe"
        ],
        "conns": 0
      },
      {
        "id": 1560,
        "pid": 580,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 16699392,
        "vms": 5791744,
        "swap": 0,
        "memory_usage": 0.19,
        "cmd": [
          "C:\\Windows\\System32\\spoolsv.exe"
        ],
        "listen": [
          49667,
          49667
        ],
        "conns": 2
      },
      {
        "id": 1700,
        "pid": 580,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 27164672,
        "vms": 10113024,
        "swap": 0,
        "memory_usage": 0.32,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "utcsvc"
        ],
        "conns": 0
      },
      {
        "id": 1784,
        "pid": 580,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 17891328,
        "vms": 6275072,
        "swap": 0,
        "memory_usage": 0.21,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "appmodel"
        ],
        "conns": 0
      },
      {
        "id": 1832,
        "pid": 580,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.02,
        "rss": 103481344,
        "vms": 226627584,
        "swap": 0,
        "memory_usage": 1.2,
        "cmd": [
          ""
        ],
        "conns": 0
      },
      {
        "id": 1912,
        "pid": 580,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 8454144,
        "vms": 2150400,
        "swap": 0,
        "memory_usage": 0.1,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "smbsvcs"
        ],
        "conns": 0
      },
      {
        "id": 2396,
        "pid": 360,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 8171520,
        "vms": 1884160,
        "swap": 0,
        "memory_usage": 0.1,
        "cmd": [
          "winlogon.exe"
        ],
        "conns": 0
      },
      {
        "id": 2568,
        "pid": 316,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 6971392,
        "vms": 1503232,
        "swap": 0,
        "memory_usage": 0.08,
        "cmd": [
          "vm3dservice.exe",
          "-n"
        ],
        "conns": 0
      },
      {
        "id": 2576,
        "pid": 32,
        "user": "WIN-5UGRJ67RCBK\\Administrator",
        "cpu_usage": 0,
        "rss": 15876096,
        "vms": 4136960,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "\\??\\C:\\Windows\\system32\\conhost.exe",
          "0x4"
        ],
        "conns": 0
      },
      {
        "id": 2672,
        "pid": 868,
        "user": "WIN-5UGRJ67RCBK\\Administrator",
        "cpu_usage": 0,
        "rss": 13283328,
        "vms": 2478080,
        "swap": 0,
        "memory_usage": 0.15,
        "cmd": [
          "rdpclip"
        ],
        "conns": 0
      },
      {
        "id": 3116,
        "pid": 700,
        "user": "WIN-5UGRJ67RCBK\\Administrator",
        "cpu_usage": 0,
        "rss": 20987904,
        "vms": 5988352,
        "swap": 0,
        "memory_usage": 0.24,
        "cmd": [
          "C:\\Windows\\System32\\InputMethod\\CHS\\ChsIME.exe",
          "-Embedding"
        ],
        "conns": 0
      },
      {
        "id": 3192,
        "pid": 3168,
        "user": "WIN-5UGRJ67RCBK\\Administrator",
        "cpu_usage": 0.01,
        "rss": 94474240,
        "vms": 28360704,
        "swap": 0,
        "memory_usage": 1.1,
        "cmd": [
          "C:\\Windows\\Explorer.EXE"
        ],
        "conns": 1
      },
      {
        "id": 3656,
        "pid": 3192,
        "user": "WIN-5UGRJ67RCBK\\Administrator",
        "cpu_usage": 0,
        "rss": 34537472,
        "vms": 11530240,
        "swap": 0,
        "memory_usage": 0.4,
        "cmd": [
          "\"C:\\Windows\\system32\\taskmgr.exe\"",
          "/4"
        ],
        "conns": 0
      },
      {
        "id": 3764,
        "pid": 700,
        "user": "WIN-5UGRJ67RCBK\\Administrator",
        "cpu_usage": 0,
        "rss": 114384896,
        "vms": 53821440,
        "swap": 0,
        "memory_usage": 1.33,
        "cmd": [
          "\"C:\\Windows\\SystemApps\\Microsoft.Windows.Cortana_cw5n1h2txyewy\\SearchUI.exe\"",
          "-ServerName:CortanaUI.AppXa50dqqa5gqv4a428c9y1jjw7m3btvepj.mca"
        ],
        "conns": 0
      },
      {
        "id": 4040,
        "pid": 700,
        "user": "WIN-5UGRJ67RCBK\\Administrator",
        "cpu_usage": 0,
        "rss": 68177920,
        "vms": 24653824,
        "swap": 0,
        "memory_usage": 0.79,
        "cmd": [
          "\"C:\\Windows\\SystemApps\\ShellExperienceHost_cw5n1h2txyewy\\ShellExperienceHost.exe\"",
          "-ServerName:App.AppXtk181tbxbce2qsex02s8tw7hfxa9xb3t.mca"
        ],
        "conns": 0
      },
      {
        "id": 4288,
        "pid": 2284,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 11206656,
        "vms": 2912256,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          "\"C:\\Program",
          "Files\\Windows",
          "Defender\\\\MpCmdRun.exe\"",
          "SpyNetService",
          "-RestrictPrivileges",
          "-AccessKey",
          "CFA9E3EA-908C-BECB-872E-03D4739E7942",
          "-Reinvoke"
        ],
        "conns": 0
      },
      {
        "id": 4932,
        "pid": 580,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.03,
        "rss": 40673280,
        "vms": 45113344,
        "swap": 0,
        "memory_usage": 0.47,
        "cmd": [
          "C:\\metrics-agent\\metrics.exe",
          "-conf",
          "C:\\metrics-agent\\agent.conf"
        ],
        "conns": 1
      },
      {
        "id": 4992,
        "pid": 32,
        "user": "WIN-5UGRJ67RCBK\\Administrator",
        "cpu_usage": 0,
        "rss": 14675968,
        "vms": 4681728,
        "swap": 0,
        "memory_usage": 0.17,
        "cmd": [
          "wmic"
        ],
        "conns": 0
      }
    ],
    "connections": [
      {
        "fd": 0,
        "pid": 764,
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
        "pid": 868,
        "type": "tcp4",
        "local": "0.0.0.0:3389",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "tcp4",
        "local": "0.0.0.0:5985",
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
        "pid": 468,
        "type": "tcp4",
        "local": "0.0.0.0:49664",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 1004,
        "type": "tcp4",
        "local": "0.0.0.0:49665",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 856,
        "type": "tcp4",
        "local": "0.0.0.0:49666",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 1560,
        "type": "tcp4",
        "local": "0.0.0.0:49667",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 580,
        "type": "tcp4",
        "local": "0.0.0.0:49668",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 1448,
        "type": "tcp4",
        "local": "0.0.0.0:49669",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 588,
        "type": "tcp4",
        "local": "0.0.0.0:49671",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "tcp4",
        "local": "192.168.2.75:139",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 856,
        "type": "tcp4",
        "local": "192.168.2.75:49696",
        "remote": "20.198.162.76:443",
        "status": "ESTABLISHED"
      },
      {
        "fd": 0,
        "pid": 4932,
        "type": "tcp4",
        "local": "192.168.2.75:50043",
        "remote": "192.168.2.52:13081",
        "status": "ESTABLISHED"
      },
      {
        "fd": 0,
        "pid": 3192,
        "type": "tcp4",
        "local": "192.168.2.75:50354",
        "remote": "20.198.162.78:443",
        "status": "ESTABLISHED"
      },
      {
        "fd": 0,
        "pid": 764,
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
        "pid": 868,
        "type": "tcp6",
        "local": ":::3389",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "tcp6",
        "local": ":::5985",
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
        "pid": 468,
        "type": "tcp6",
        "local": ":::49664",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 1004,
        "type": "tcp6",
        "local": ":::49665",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 856,
        "type": "tcp6",
        "local": ":::49666",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 1560,
        "type": "tcp6",
        "local": ":::49667",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 580,
        "type": "tcp6",
        "local": ":::49668",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 1448,
        "type": "tcp6",
        "local": ":::49669",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 588,
        "type": "tcp6",
        "local": ":::49671",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 968,
        "type": "udp4",
        "local": "0.0.0.0:123",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 856,
        "type": "udp4",
        "local": "0.0.0.0:500",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 868,
        "type": "udp4",
        "local": "0.0.0.0:3389",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 856,
        "type": "udp4",
        "local": "0.0.0.0:3544",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 856,
        "type": "udp4",
        "local": "0.0.0.0:4500",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 968,
        "type": "udp4",
        "local": "0.0.0.0:5050",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1088,
        "type": "udp4",
        "local": "0.0.0.0:5353",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1088,
        "type": "udp4",
        "local": "0.0.0.0:5355",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1360,
        "type": "udp4",
        "local": "127.0.0.1:1900",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1360,
        "type": "udp4",
        "local": "127.0.0.1:49576",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "udp4",
        "local": "192.168.2.75:137",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "udp4",
        "local": "192.168.2.75:138",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1360,
        "type": "udp4",
        "local": "192.168.2.75:1900",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1360,
        "type": "udp4",
        "local": "192.168.2.75:49575",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 856,
        "type": "udp4",
        "local": "192.168.2.75:51482",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 968,
        "type": "udp6",
        "local": ":::123",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 856,
        "type": "udp6",
        "local": ":::500",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 868,
        "type": "udp6",
        "local": ":::3389",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 856,
        "type": "udp6",
        "local": ":::4500",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1360,
        "type": "udp6",
        "local": "::1:1900",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 1360,
        "type": "udp6",
        "local": "::1:49574",
        "status": ""
      }
    ]
  }
}
```