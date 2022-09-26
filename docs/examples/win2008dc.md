static:

```json
{
  "code": 200,
  "payload": {
    "time": 1664186176,
    "host": {
      "name": "WIN-HLLQ0EB9O72",
      "uptime": 343732
    },
    "os": {
      "name": "windows",
      "platform_name": "Microsoft Windows Server 2008 R2 Datacenter Service Pack 1",
      "platform_version": "6.1.7601 Build 7601",
      "install": 1663813721,
      "startup": 1663842444
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
        "mac": "00:50:56:a3:11:64",
        "addrs": [
          "192.168.2.80/24"
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
        "name": "isatap.{4ACE2CC0-D660-423B-AA2A-467EE4C7D6A2}",
        "mtu": 1280,
        "flags": [
          "pointtopoint",
          "multicast"
        ],
        "mac": "00:00:00:00:00:00:00:e0",
        "addrs": [
          "fe80::5efe:c0a8:250/128"
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
        "usage": 0.13
      },
      "memory": {
        "used": 975552512,
        "free": 7613915136,
        "available": 7613915136,
        "usage": 11,
        "swap_used": 0,
        "swap_free": 0
      },
      "partitions": [
        {
          "name": "C:",
          "used": 15868796928,
          "free": 37711335424,
          "usage": 29.62
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
          "bytes_sent": 18001489,
          "bytes_recv": 31806692,
          "packets_sent": 145998,
          "packets_recv": 306554
        },
        {
          "name": "Loopback Pseudo-Interface 1",
          "bytes_sent": 0,
          "bytes_recv": 0,
          "packets_sent": 0,
          "packets_recv": 0
        },
        {
          "name": "isatap.{4ACE2CC0-D660-423B-AA2A-467EE4C7D6A2}",
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
        "cpu_usage": 0.03,
        "rss": 372736,
        "vms": 110592,
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
        "id": 252,
        "pid": 4,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 1224704,
        "vms": 585728,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "\\SystemRoot\\System32\\smss.exe"
        ],
        "conns": 0
      },
      {
        "id": 276,
        "pid": 1412,
        "user": "WIN-HLLQ0EB9O72\\Administrator",
        "cpu_usage": 0.06,
        "rss": 15388672,
        "vms": 67846144,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "\"C:\\Windows\\system32\\oobe.exe\"",
          ""
        ],
        "conns": 0
      },
      {
        "id": 312,
        "pid": 468,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.02,
        "rss": 14520320,
        "vms": 7127040,
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
        "rss": 5976064,
        "vms": 2322432,
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
        "pid": 324,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 5185536,
        "vms": 1769472,
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
        "id": 380,
        "pid": 364,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 13807616,
        "vms": 4227072,
        "swap": 0,
        "memory_usage": 0.16,
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
        "id": 408,
        "pid": 364,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 7671808,
        "vms": 3239936,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "winlogon.exe"
        ],
        "conns": 0
      },
      {
        "id": 468,
        "pid": 372,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.01,
        "rss": 8372224,
        "vms": 4337664,
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
        "pid": 372,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.01,
        "rss": 13799424,
        "vms": 6361088,
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
        "pid": 372,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 6483968,
        "vms": 3313664,
        "swap": 0,
        "memory_usage": 0.08,
        "cmd": [
          "C:\\Windows\\system32\\lsm.exe"
        ],
        "conns": 0
      },
      {
        "id": 596,
        "pid": 468,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.01,
        "rss": 10502144,
        "vms": 5070848,
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
        "id": 668,
        "pid": 468,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 8265728,
        "vms": 4509696,
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
        "id": 680,
        "pid": 468,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 2.85,
        "rss": 7471104,
        "vms": 2420736,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "C:\\Windows\\servicing\\TrustedInstaller.exe"
        ],
        "conns": 0
      },
      {
        "id": 780,
        "pid": 468,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0.08,
        "rss": 14036992,
        "vms": 7020544,
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
        "user": "WIN-HLLQ0EB9O72\\Administrator",
        "cpu_usage": 0,
        "rss": 7503872,
        "vms": 2887680,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "\"taskhost.exe\""
        ],
        "conns": 0
      },
      {
        "id": 824,
        "pid": 468,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 14430208,
        "vms": 7159808,
        "swap": 0,
        "memory_usage": 0.17,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "LocalService"
        ],
        "conns": 2
      },
      {
        "id": 848,
        "pid": 468,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0.02,
        "rss": 42418176,
        "vms": 70832128,
        "swap": 0,
        "memory_usage": 0.49,
        "cmd": [
          "C:\\Windows\\System32\\svchost.exe",
          "-k",
          "NetworkService"
        ],
        "conns": 1
      },
      {
        "id": 900,
        "pid": 468,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0.01,
        "rss": 11894784,
        "vms": 9371648,
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
        "id": 936,
        "pid": 468,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 13090816,
        "vms": 9871360,
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
        "cpu_usage": 0.06,
        "rss": 38940672,
        "vms": 24371200,
        "swap": 0,
        "memory_usage": 0.45,
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
        "id": 996,
        "pid": 1832,
        "user": "WIN-HLLQ0EB9O72\\Administrator",
        "cpu_usage": 0.05,
        "rss": 11853824,
        "vms": 3690496,
        "swap": 0,
        "memory_usage": 0.14,
        "cmd": [
          "\"C:\\Windows\\system32\\taskmgr.exe\"",
          "/4"
        ],
        "conns": 0
      },
      {
        "id": 1004,
        "pid": 468,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 8855552,
        "vms": 3424256,
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
        "id": 1120,
        "pid": 468,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 15192064,
        "vms": 8908800,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "C:\\Windows\\System32\\spoolsv.exe"
        ],
        "conns": 0
      },
      {
        "id": 1164,
        "pid": 468,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 3039232,
        "vms": 1150976,
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
        "id": 1244,
        "pid": 468,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 5578752,
        "vms": 2015232,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "C:\\Windows\\system32\\svchost.exe",
          "-k",
          "NetworkServiceNetworkRestricted"
        ],
        "listen": [
          49158,
          49158
        ],
        "conns": 2
      },
      {
        "id": 1604,
        "pid": 596,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0.24,
        "rss": 8687616,
        "vms": 4349952,
        "swap": 0,
        "memory_usage": 0.1,
        "cmd": [
          "C:\\Windows\\system32\\wbem\\wmiprvse.exe"
        ],
        "conns": 0
      },
      {
        "id": 1832,
        "pid": 2012,
        "user": "WIN-HLLQ0EB9O72\\Administrator",
        "cpu_usage": 0.11,
        "rss": 56795136,
        "vms": 31756288,
        "swap": 0,
        "memory_usage": 0.66,
        "cmd": [
          "C:\\Windows\\Explorer.EXE"
        ],
        "conns": 0
      },
      {
        "id": 2000,
        "pid": 468,
        "user": "NT AUTHORITY\\LOCAL SERVICE",
        "cpu_usage": 0,
        "rss": 4857856,
        "vms": 1765376,
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
        "id": 2024,
        "pid": 312,
        "user": "WIN-HLLQ0EB9O72\\Administrator",
        "cpu_usage": 0,
        "rss": 5173248,
        "vms": 1892352,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "\"C:\\Windows\\system32\\Dwm.exe\""
        ],
        "conns": 0
      },
      {
        "id": 2044,
        "pid": 468,
        "user": "NT AUTHORITY\\NETWORK SERVICE",
        "cpu_usage": 0,
        "rss": 7839744,
        "vms": 3584000,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "C:\\Windows\\System32\\msdtc.exe"
        ],
        "conns": 0
      },
      {
        "id": 2152,
        "pid": 2116,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 6483968,
        "vms": 2084864,
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
      },
      {
        "id": 2180,
        "pid": 2116,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 4849664,
        "vms": 1867776,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "winlogon.exe"
        ],
        "conns": 0
      },
      {
        "id": 2244,
        "pid": 2180,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0,
        "rss": 16031744,
        "vms": 7389184,
        "swap": 0,
        "memory_usage": 0.19,
        "cmd": [
          "\"LogonUI.exe\"",
          "/flags:0x0"
        ],
        "conns": 0
      },
      {
        "id": 2328,
        "pid": 1004,
        "user": "WIN-HLLQ0EB9O72\\Administrator",
        "cpu_usage": 0,
        "rss": 7213056,
        "vms": 2068480,
        "swap": 0,
        "memory_usage": 0.08,
        "cmd": [
          "rdpclip"
        ],
        "conns": 0
      },
      {
        "id": 2372,
        "pid": 2952,
        "user": "WIN-HLLQ0EB9O72\\Administrator",
        "cpu_usage": 0,
        "rss": 10227712,
        "vms": 4313088,
        "swap": 0,
        "memory_usage": 0.12,
        "cmd": [
          "wmic"
        ],
        "conns": 0
      },
      {
        "id": 2800,
        "pid": 468,
        "user": "NT AUTHORITY\\SYSTEM",
        "cpu_usage": 0.02,
        "rss": 36634624,
        "vms": 45543424,
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
        "id": 2912,
        "pid": 380,
        "user": "WIN-HLLQ0EB9O72\\Administrator",
        "cpu_usage": 0,
        "rss": 9588736,
        "vms": 12857344,
        "swap": 0,
        "memory_usage": 0.11,
        "cmd": [
          "\\??\\C:\\Windows\\system32\\conhost.exe"
        ],
        "conns": 0
      },
      {
        "id": 2952,
        "pid": 1832,
        "user": "WIN-HLLQ0EB9O72\\Administrator",
        "cpu_usage": 0,
        "rss": 6656000,
        "vms": 3985408,
        "swap": 0,
        "memory_usage": 0.08,
        "cmd": [
          "\"C:\\Windows\\system32\\cmd.exe\"",
          ""
        ],
        "conns": 0
      }
    ],
    "connections": [
      {
        "fd": 0,
        "pid": 668,
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
        "pid": 1004,
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
        "pid": 372,
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
        "pid": 936,
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
        "pid": 1244,
        "type": "tcp4",
        "local": "0.0.0.0:49158",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "tcp4",
        "local": "192.168.2.80:139",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 2800,
        "type": "tcp4",
        "local": "192.168.2.80:49159",
        "remote": "192.168.2.52:13081",
        "status": "ESTABLISHED"
      },
      {
        "fd": 0,
        "pid": 668,
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
        "pid": 1004,
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
        "pid": 372,
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
        "pid": 936,
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
        "pid": 1244,
        "type": "tcp6",
        "local": ":::49158",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 0,
        "pid": 824,
        "type": "udp4",
        "local": "0.0.0.0:123",
        "status": ""
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
        "pid": 848,
        "type": "udp4",
        "local": "0.0.0.0:5355",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "udp4",
        "local": "192.168.2.80:137",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 4,
        "type": "udp4",
        "local": "192.168.2.80:138",
        "status": ""
      },
      {
        "fd": 0,
        "pid": 824,
        "type": "udp6",
        "local": ":::123",
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