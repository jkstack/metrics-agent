static:

```json
{
  "code": 200,
  "payload": {
    "time": 1663223978,
    "host": {
      "name": "localhost.localdomain",
      "uptime": 533151
    },
    "os": {
      "name": "linux",
      "platform_name": "centos",
      "platform_version": "6.1",
      "install": 1662687772,
      "startup": 1662690828
    },
    "kernel": {
      "version": "2.6.32-131.0.15.el6.x86_64",
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
          "mhz": 2297.34
        },
        {
          "processor": 1,
          "model": "Intel(R) Xeon(R) CPU E5-2670 v3 @ 2.30GHz",
          "core": 0,
          "cores": 1,
          "physical": 0,
          "mhz": 2297.34
        }
      ]
    },
    "memory": {
      "physical": 4155539456,
      "swap": 6308225024
    },
    "disks": [
      {
        "model": "unknown",
        "total": 53687091200,
        "type": "HDD",
        "disks": [
          "/dev/sda1",
          "/dev/sda2"
        ]
      }
    ],
    "partitions": [
      {
        "mount": "/",
        "fstype": "ext4",
        "opts": [
          "rw",
          "relatime"
        ],
        "total": 46114725888,
        "inodes": 2861600
      },
      {
        "mount": "/boot",
        "fstype": "ext4",
        "opts": [
          "rw",
          "relatime"
        ],
        "total": 507744256,
        "inodes": 128016
      }
    ],
    "gateway": "192.168.2.1",
    "interface": [
      {
        "index": 1,
        "name": "lo",
        "mtu": 16436,
        "flags": [
          "up",
          "loopback"
        ],
        "mac": "",
        "addrs": [
          "127.0.0.1/8",
          "::1/128"
        ]
      },
      {
        "index": 2,
        "name": "eth0",
        "mtu": 1500,
        "flags": [
          "up",
          "broadcast",
          "multicast"
        ],
        "mac": "00:50:56:a3:15:7f",
        "addrs": [
          "192.168.2.51/24",
          "fe80::250:56ff:fea3:157f/64"
        ]
      }
    ],
    "user": [
      {
        "name": "root",
        "id": "0",
        "gid": "0"
      },
      {
        "name": "bin",
        "id": "1",
        "gid": "1"
      },
      {
        "name": "daemon",
        "id": "2",
        "gid": "2"
      },
      {
        "name": "adm",
        "id": "3",
        "gid": "4"
      },
      {
        "name": "lp",
        "id": "4",
        "gid": "7"
      },
      {
        "name": "sync",
        "id": "5",
        "gid": "0"
      },
      {
        "name": "shutdown",
        "id": "6",
        "gid": "0"
      },
      {
        "name": "halt",
        "id": "7",
        "gid": "0"
      },
      {
        "name": "mail",
        "id": "8",
        "gid": "12"
      },
      {
        "name": "uucp",
        "id": "10",
        "gid": "14"
      },
      {
        "name": "operator",
        "id": "11",
        "gid": "0"
      },
      {
        "name": "games",
        "id": "12",
        "gid": "100"
      },
      {
        "name": "gopher",
        "id": "13",
        "gid": "30"
      },
      {
        "name": "FTP User",
        "id": "14",
        "gid": "50"
      },
      {
        "name": "Nobody",
        "id": "99",
        "gid": "99"
      },
      {
        "name": "System message bus",
        "id": "81",
        "gid": "81"
      },
      {
        "name": "virtual console memory owner",
        "id": "69",
        "gid": "69"
      },
      {
        "name": "HAL daemon",
        "id": "68",
        "gid": "68"
      },
      {
        "name": "\"Saslauthd user\"",
        "id": "499",
        "gid": "499"
      },
      {
        "name": "",
        "id": "89",
        "gid": "89"
      },
      {
        "name": "Avahi mDNS/DNS-SD Stack",
        "id": "70",
        "gid": "70"
      },
      {
        "name": "Privilege-separated SSH",
        "id": "74",
        "gid": "74"
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
        "usage": 0.03
      },
      "memory": {
        "used": 227143680,
        "free": 3573792768,
        "available": 3742199808,
        "usage": 5.47,
        "swap_used": 0,
        "swap_free": 6308225024
      },
      "partitions": [
        {
          "name": "/",
          "used": 1215700992,
          "free": 42556506112,
          "usage": 2.78,
          "inode_used": 33323,
          "inode_free": 2828277,
          "inode_usage": 1.16
        },
        {
          "name": "/boot",
          "used": 31692800,
          "free": 449837056,
          "usage": 6.58,
          "inode_used": 38,
          "inode_free": 127978,
          "inode_usage": 0.03
        }
      ],
      "interface": [
        {
          "name": "lo",
          "bytes_sent": 30156120,
          "bytes_recv": 30156120,
          "packets_sent": 297147,
          "packets_recv": 297147
        },
        {
          "name": "eth0",
          "bytes_sent": 10808559,
          "bytes_recv": 256141815,
          "packets_sent": 126898,
          "packets_recv": 932694
        }
      ]
    },
    "process": [
      {
        "id": 1,
        "pid": 0,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1613824,
        "vms": 19779584,
        "swap": 0,
        "memory_usage": 0.04,
        "cmd": [
          "/sbin/init"
        ],
        "conns": 3
      },
      {
        "id": 2,
        "pid": 0,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 3,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 4,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 5,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 6,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 7,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 8,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 9,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 10,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 11,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 12,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 13,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 14,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 15,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 16,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 17,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 18,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 19,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 20,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 21,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 22,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 23,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 24,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 25,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 26,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 27,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 28,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 29,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 30,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 31,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 32,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 33,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 34,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 35,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 36,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 37,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 38,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 39,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 40,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 41,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 42,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 43,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 44,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 49,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 50,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 51,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 53,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 54,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 85,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 262,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 263,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 264,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 267,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 268,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 334,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 338,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 353,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 354,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 355,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 432,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1556480,
        "vms": 11751424,
        "swap": 0,
        "memory_usage": 0.04,
        "cmd": [
          "/sbin/udevd",
          "-d"
        ],
        "conns": 2
      },
      {
        "id": 659,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 786,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 787,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 788,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 826,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 852,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 1072,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 937984,
        "vms": 28303360,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "auditd"
        ],
        "conns": 0
      },
      {
        "id": 1088,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1536000,
        "vms": 254648320,
        "swap": 0,
        "memory_usage": 0.04,
        "cmd": [
          "/sbin/rsyslogd",
          "-c",
          "4"
        ],
        "conns": 1
      },
      {
        "id": 1100,
        "pid": 1,
        "user": "dbus",
        "cpu_usage": 0,
        "rss": 1167360,
        "vms": 32501760,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "dbus-daemon",
          "--system"
        ],
        "conns": 2
      },
      {
        "id": 1112,
        "pid": 1,
        "user": "avahi",
        "cpu_usage": 0,
        "rss": 1646592,
        "vms": 28307456,
        "swap": 0,
        "memory_usage": 0.04,
        "cmd": [
          "avahi-daemon: running [linux.local]"
        ],
        "conns": 5
      },
      {
        "id": 1113,
        "pid": 1112,
        "user": "avahi",
        "cpu_usage": 0,
        "rss": 524288,
        "vms": 28307456,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "avahi-daemon: chroot helper"
        ],
        "conns": 2
      },
      {
        "id": 1136,
        "pid": 1,
        "user": "haldaemon",
        "cpu_usage": 0,
        "rss": 3624960,
        "vms": 25657344,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "hald"
        ],
        "conns": 4
      },
      {
        "id": 1137,
        "pid": 1136,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1331200,
        "vms": 18530304,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "hald-runner"
        ],
        "conns": 1
      },
      {
        "id": 1165,
        "pid": 1137,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1265664,
        "vms": 20697088,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "hald-addon-input: Listening on /dev/input/event2 /dev/input/event0"
        ],
        "conns": 1
      },
      {
        "id": 1189,
        "pid": 1137,
        "user": "haldaemon",
        "cpu_usage": 0,
        "rss": 1155072,
        "vms": 18219008,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "hald-addon-acpi: listening on acpi kernel interface /proc/acpi/event"
        ],
        "conns": 1
      },
      {
        "id": 1268,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3264512,
        "vms": 80429056,
        "swap": 0,
        "memory_usage": 0.08,
        "cmd": [
          "/usr/libexec/postfix/master"
        ],
        "listen": [
          25,
          25
        ],
        "conns": 25
      },
      {
        "id": 1275,
        "pid": 1268,
        "user": "postfix",
        "cpu_usage": 0,
        "rss": 3276800,
        "vms": 80580608,
        "swap": 0,
        "memory_usage": 0.08,
        "cmd": [
          "qmgr",
          "-l",
          "-t",
          "fifo",
          "-u"
        ],
        "conns": 2
      },
      {
        "id": 1278,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1376256,
        "vms": 119984128,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "crond"
        ],
        "conns": 1
      },
      {
        "id": 1293,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 589824,
        "vms": 4149248,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "/sbin/mingetty",
          "/dev/tty2"
        ],
        "conns": 0
      },
      {
        "id": 1295,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 585728,
        "vms": 4149248,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "/sbin/mingetty",
          "/dev/tty3"
        ],
        "conns": 0
      },
      {
        "id": 1297,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 589824,
        "vms": 4149248,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "/sbin/mingetty",
          "/dev/tty4"
        ],
        "conns": 0
      },
      {
        "id": 1299,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 581632,
        "vms": 4149248,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "/sbin/mingetty",
          "/dev/tty5"
        ],
        "conns": 0
      },
      {
        "id": 1306,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 589824,
        "vms": 4149248,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "/sbin/mingetty",
          "/dev/tty6"
        ],
        "conns": 0
      },
      {
        "id": 1309,
        "pid": 432,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2588672,
        "vms": 12558336,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "/sbin/udevd",
          "-d"
        ],
        "conns": 1
      },
      {
        "id": 1310,
        "pid": 432,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2588672,
        "vms": 12558336,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "/sbin/udevd",
          "-d"
        ],
        "conns": 1
      },
      {
        "id": 1568,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1134592,
        "vms": 65859584,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "/usr/sbin/sshd"
        ],
        "listen": [
          22,
          22
        ],
        "conns": 2
      },
      {
        "id": 17645,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1458176,
        "vms": 49557504,
        "swap": 0,
        "memory_usage": 0.04,
        "cmd": [
          "su",
          "root",
          "-c",
          "/opt/metrics-agent/metrics \"-conf\" \"/opt/metrics-agent/agent.conf\""
        ],
        "conns": 1
      },
      {
        "id": 17646,
        "pid": 17645,
        "user": "root",
        "cpu_usage": 0.03,
        "rss": 46673920,
        "vms": 1088790528,
        "swap": 0,
        "memory_usage": 1.12,
        "cmd": [
          "/opt/metrics-agent/metrics",
          "-conf",
          "/opt/metrics-agent/agent.conf"
        ],
        "conns": 1
      },
      {
        "id": 17698,
        "pid": 1268,
        "user": "postfix",
        "cpu_usage": 0,
        "rss": 3342336,
        "vms": 80908288,
        "swap": 0,
        "memory_usage": 0.08,
        "cmd": [
          "pickup",
          "-l",
          "-t",
          "fifo",
          "-u"
        ],
        "conns": 2
      },
      {
        "id": 17714,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 28186,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1454080,
        "vms": 49557504,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "su",
          "root",
          "-c",
          "/opt/agent-server/server \"-conf\" \"/opt/agent-server/server.conf\""
        ],
        "conns": 1
      },
      {
        "id": 28188,
        "pid": 28186,
        "user": "root",
        "cpu_usage": 0.02,
        "rss": 48398336,
        "vms": 902987776,
        "swap": 0,
        "memory_usage": 1.16,
        "cmd": [
          "/opt/agent-server/server",
          "-conf",
          "/opt/agent-server/server.conf"
        ],
        "listen": [
          13081
        ],
        "conns": 3
      }
    ],
    "connections": [
      {
        "fd": 3,
        "pid": 1568,
        "type": "tcp4",
        "local": "0.0.0.0:22",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 12,
        "pid": 1268,
        "type": "tcp4",
        "local": "127.0.0.1:25",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 7,
        "pid": 17646,
        "type": "tcp4",
        "local": "127.0.0.1:55227",
        "remote": "127.0.0.1:13081",
        "status": "ESTABLISHED"
      },
      {
        "fd": 4,
        "pid": 1568,
        "type": "tcp6",
        "local": ":::22",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 7,
        "pid": 28188,
        "type": "tcp6",
        "local": ":::13081",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 13,
        "pid": 1268,
        "type": "tcp6",
        "local": "::1:25",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 9,
        "pid": 28188,
        "type": "tcp6",
        "local": "127.0.0.1:13081",
        "remote": "127.0.0.1:55227",
        "status": "ESTABLISHED"
      },
      {
        "fd": 8,
        "pid": 28188,
        "type": "tcp6",
        "local": "192.168.2.51:13081",
        "remote": "10.202.0.227:54146",
        "status": "ESTABLISHED"
      },
      {
        "fd": 13,
        "pid": 1112,
        "type": "udp4",
        "local": "0.0.0.0:5353",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 14,
        "pid": 1112,
        "type": "udp4",
        "local": "0.0.0.0:36275",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 7,
        "pid": 1,
        "type": "file",
        "local": "@/com/ubuntu/upstart",
        "status": "NONE"
      },
      {
        "fd": 0,
        "pid": 1088,
        "type": "file",
        "local": "/dev/log",
        "status": "NONE"
      },
      {
        "fd": 7,
        "pid": 1136,
        "type": "file",
        "local": "@/var/run/hald/dbus-QxcNT5genA",
        "status": "NONE"
      },
      {
        "fd": 9,
        "pid": 1136,
        "type": "file",
        "local": "@/var/run/hald/dbus-VC0hJG5fi5",
        "status": "NONE"
      },
      {
        "fd": 4,
        "pid": 432,
        "type": "file",
        "local": "@/org/kernel/udev/udevd",
        "status": "NONE"
      },
      {
        "fd": 13,
        "pid": 1136,
        "type": "file",
        "local": "@/org/freedesktop/hal/udev_event",
        "status": "NONE"
      },
      {
        "fd": 3,
        "pid": 1100,
        "type": "file",
        "local": "/var/run/dbus/system_bus_socket",
        "status": "NONE"
      },
      {
        "fd": 10,
        "pid": 1112,
        "type": "file",
        "local": "/var/run/avahi-daemon/socket",
        "status": "NONE"
      },
      {
        "fd": 20,
        "pid": 1268,
        "type": "file",
        "local": "public/cleanup",
        "status": "NONE"
      },
      {
        "fd": 26,
        "pid": 1268,
        "type": "file",
        "local": "private/tlsmgr",
        "status": "NONE"
      },
      {
        "fd": 29,
        "pid": 1268,
        "type": "file",
        "local": "private/rewrite",
        "status": "NONE"
      },
      {
        "fd": 32,
        "pid": 1268,
        "type": "file",
        "local": "private/bounce",
        "status": "NONE"
      },
      {
        "fd": 35,
        "pid": 1268,
        "type": "file",
        "local": "private/defer",
        "status": "NONE"
      },
      {
        "fd": 38,
        "pid": 1268,
        "type": "file",
        "local": "private/trace",
        "status": "NONE"
      },
      {
        "fd": 41,
        "pid": 1268,
        "type": "file",
        "local": "private/verify",
        "status": "NONE"
      },
      {
        "fd": 44,
        "pid": 1268,
        "type": "file",
        "local": "public/flush",
        "status": "NONE"
      },
      {
        "fd": 47,
        "pid": 1268,
        "type": "file",
        "local": "private/proxymap",
        "status": "NONE"
      },
      {
        "fd": 50,
        "pid": 1268,
        "type": "file",
        "local": "private/proxywrite",
        "status": "NONE"
      },
      {
        "fd": 53,
        "pid": 1268,
        "type": "file",
        "local": "private/smtp",
        "status": "NONE"
      },
      {
        "fd": 56,
        "pid": 1268,
        "type": "file",
        "local": "private/relay",
        "status": "NONE"
      },
      {
        "fd": 59,
        "pid": 1268,
        "type": "file",
        "local": "public/showq",
        "status": "NONE"
      },
      {
        "fd": 62,
        "pid": 1268,
        "type": "file",
        "local": "private/error",
        "status": "NONE"
      },
      {
        "fd": 65,
        "pid": 1268,
        "type": "file",
        "local": "private/retry",
        "status": "NONE"
      },
      {
        "fd": 68,
        "pid": 1268,
        "type": "file",
        "local": "private/discard",
        "status": "NONE"
      },
      {
        "fd": 71,
        "pid": 1268,
        "type": "file",
        "local": "private/local",
        "status": "NONE"
      },
      {
        "fd": 74,
        "pid": 1268,
        "type": "file",
        "local": "private/virtual",
        "status": "NONE"
      },
      {
        "fd": 77,
        "pid": 1268,
        "type": "file",
        "local": "private/lmtp",
        "status": "NONE"
      },
      {
        "fd": 80,
        "pid": 1268,
        "type": "file",
        "local": "private/anvil",
        "status": "NONE"
      },
      {
        "fd": 83,
        "pid": 1268,
        "type": "file",
        "local": "private/scache",
        "status": "NONE"
      },
      {
        "fd": 7,
        "pid": 17698,
        "type": "file",
        "status": "NONE"
      },
      {
        "fd": 85,
        "pid": 1268,
        "type": "file",
        "status": "NONE"
      }
    ]
  }
}
```