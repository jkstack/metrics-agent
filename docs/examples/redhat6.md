static:

```json
{
  "code": 200,
  "payload": {
    "time": 1663568666,
    "host": {
      "name": "rhel6.1",
      "uptime": 356731
    },
    "os": {
      "name": "linux",
      "platform_name": "redhat",
      "platform_version": "6.1",
      "install": 1663181888,
      "startup": 1663211936
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
      "physical": 4020273152,
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
        "mac": "00:50:56:a3:55:7b",
        "addrs": [
          "192.168.2.65/24",
          "fe80::250:56ff:fea3:557b/64"
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
        "name": "ftp",
        "id": "14",
        "gid": "50"
      },
      {
        "name": "nobody",
        "id": "99",
        "gid": "99"
      },
      {
        "name": "dbus",
        "id": "81",
        "gid": "81"
      },
      {
        "name": "vcsa",
        "id": "69",
        "gid": "69"
      },
      {
        "name": "rpc",
        "id": "32",
        "gid": "32"
      },
      {
        "name": "abrt",
        "id": "173",
        "gid": "173"
      },
      {
        "name": "haldaemon",
        "id": "68",
        "gid": "68"
      },
      {
        "name": "ntp",
        "id": "38",
        "gid": "38"
      },
      {
        "name": "saslauth",
        "id": "499",
        "gid": "499"
      },
      {
        "name": "postfix",
        "id": "89",
        "gid": "89"
      },
      {
        "name": "avahi",
        "id": "70",
        "gid": "70"
      },
      {
        "name": "rpcuser",
        "id": "29",
        "gid": "29"
      },
      {
        "name": "nfsnobody",
        "id": "65534",
        "gid": "65534"
      },
      {
        "name": "sshd",
        "id": "74",
        "gid": "74"
      },
      {
        "name": "tcpdump",
        "id": "72",
        "gid": "72"
      },
      {
        "name": "oprofile",
        "id": "16",
        "gid": "16"
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
        "usage": 0.11
      },
      "memory": {
        "used": 223649792,
        "free": 3326595072,
        "available": 3597355008,
        "usage": 5.56,
        "swap_used": 0,
        "swap_free": 6308225024
      },
      "partitions": [
        {
          "name": "/",
          "used": 2478202880,
          "free": 41294004224,
          "usage": 5.66,
          "inode_used": 70571,
          "inode_free": 2791029,
          "inode_usage": 2.47
        },
        {
          "name": "/boot",
          "used": 36518912,
          "free": 445010944,
          "usage": 7.58,
          "inode_used": 39,
          "inode_free": 127977,
          "inode_usage": 0.03
        }
      ],
      "interface": [
        {
          "name": "lo",
          "bytes_sent": 0,
          "bytes_recv": 0,
          "packets_sent": 0,
          "packets_recv": 0
        },
        {
          "name": "eth0",
          "bytes_sent": 9404705,
          "bytes_recv": 48097309,
          "packets_sent": 72883,
          "packets_recv": 541692
        }
      ]
    },
    "process": [
      {
        "id": 1,
        "pid": 0,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1634304,
        "vms": 19800064,
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
        "id": 226,
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
        "id": 229,
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
        "id": 242,
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
        "id": 243,
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
        "id": 293,
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
        "id": 294,
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
        "id": 295,
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
        "id": 296,
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
        "id": 297,
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
        "id": 298,
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
        "id": 299,
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
        "id": 300,
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
        "id": 301,
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
        "id": 302,
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
        "id": 303,
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
        "id": 304,
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
        "id": 305,
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
        "id": 306,
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
        "id": 307,
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
        "id": 308,
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
        "id": 309,
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
        "id": 310,
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
        "id": 311,
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
        "id": 312,
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
        "id": 313,
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
        "id": 314,
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
        "id": 315,
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
        "id": 316,
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
        "id": 317,
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
        "id": 318,
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
        "id": 319,
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
        "id": 320,
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
        "id": 321,
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
        "id": 322,
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
        "id": 398,
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
        "id": 402,
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
        "id": 417,
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
        "id": 418,
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
        "id": 419,
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
        "id": 506,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1601536,
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
        "id": 699,
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
        "id": 880,
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
        "id": 881,
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
        "id": 882,
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
        "id": 930,
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
        "id": 1347,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.08,
        "rss": 7741440,
        "vms": 176898048,
        "swap": 0,
        "memory_usage": 0.19,
        "cmd": [
          "/usr/sbin/vmtoolsd"
        ],
        "conns": 0
      },
      {
        "id": 1412,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 913408,
        "vms": 28278784,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "auditd"
        ],
        "conns": 0
      },
      {
        "id": 1441,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1540096,
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
        "id": 1474,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 9367552,
        "vms": 48865280,
        "swap": 0,
        "memory_usage": 0.23,
        "cmd": [
          "/usr/lib/vmware-vgauth/VGAuthService",
          "-s"
        ],
        "conns": 1
      },
      {
        "id": 1490,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 561152,
        "vms": 9359360,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "irqbalance"
        ],
        "conns": 0
      },
      {
        "id": 1515,
        "pid": 1,
        "user": "rpc",
        "cpu_usage": 0,
        "rss": 917504,
        "vms": 19406848,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "rpcbind"
        ],
        "listen": [
          111,
          111
        ],
        "conns": 7
      },
      {
        "id": 1554,
        "pid": 1,
        "user": "rpcuser",
        "cpu_usage": 0,
        "rss": 1228800,
        "vms": 23678976,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "rpc.statd"
        ],
        "listen": [
          47389,
          38738
        ],
        "conns": 6
      },
      {
        "id": 1618,
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
        "id": 1619,
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
        "id": 1662,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.05,
        "rss": 5480448,
        "vms": 213180416,
        "swap": 0,
        "memory_usage": 0.14,
        "cmd": [
          "/usr/lib/vmware-caf/pme/bin/ManagementAgentHost"
        ],
        "conns": 0
      },
      {
        "id": 1667,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 561152,
        "vms": 30146560,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "rpc.idmapd"
        ],
        "conns": 1
      },
      {
        "id": 1745,
        "pid": 1,
        "user": "dbus",
        "cpu_usage": 0,
        "rss": 1171456,
        "vms": 32505856,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "dbus-daemon",
          "--system"
        ],
        "conns": 2
      },
      {
        "id": 1757,
        "pid": 1,
        "user": "avahi",
        "cpu_usage": 0,
        "rss": 1654784,
        "vms": 28307456,
        "swap": 0,
        "memory_usage": 0.04,
        "cmd": [
          "avahi-daemon: running [rhel6.local]"
        ],
        "conns": 5
      },
      {
        "id": 1758,
        "pid": 1757,
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
        "id": 1768,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3268608,
        "vms": 193552384,
        "swap": 0,
        "memory_usage": 0.08,
        "cmd": [
          "cupsd",
          "-C",
          "/etc/cups/cupsd.conf"
        ],
        "listen": [
          631,
          631
        ],
        "conns": 4
      },
      {
        "id": 1793,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 667648,
        "vms": 4165632,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/usr/sbin/acpid"
        ],
        "conns": 2
      },
      {
        "id": 1802,
        "pid": 1,
        "user": "haldaemon",
        "cpu_usage": 0,
        "rss": 3874816,
        "vms": 25919488,
        "swap": 0,
        "memory_usage": 0.1,
        "cmd": [
          "hald"
        ],
        "conns": 4
      },
      {
        "id": 1803,
        "pid": 1802,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1318912,
        "vms": 18530304,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "hald-runner"
        ],
        "conns": 1
      },
      {
        "id": 1831,
        "pid": 1803,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1261568,
        "vms": 20697088,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "hald-addon-input: Listening on /dev/input/event2 /dev/input/event0"
        ],
        "conns": 1
      },
      {
        "id": 1847,
        "pid": 1803,
        "user": "haldaemon",
        "cpu_usage": 0,
        "rss": 1155072,
        "vms": 18219008,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "hald-addon-acpi: listening on acpid socket /var/run/acpid.socket"
        ],
        "conns": 1
      },
      {
        "id": 1860,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1818624,
        "vms": 394932224,
        "swap": 0,
        "memory_usage": 0.05,
        "cmd": [
          "automount",
          "--pid-file",
          "/var/run/autofs.pid"
        ],
        "conns": 1
      },
      {
        "id": 1879,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1138688,
        "vms": 65478656,
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
        "id": 1898,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5709824,
        "vms": 270888960,
        "swap": 0,
        "memory_usage": 0.14,
        "cmd": [
          "/usr/sbin/abrtd"
        ],
        "conns": 3
      },
      {
        "id": 1906,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1392640,
        "vms": 119980032,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "crond"
        ],
        "conns": 1
      },
      {
        "id": 1917,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 471040,
        "vms": 21925888,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "/usr/sbin/atd"
        ],
        "conns": 0
      },
      {
        "id": 1932,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 471040,
        "vms": 4145152,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "/usr/bin/rhsmcertd",
          "240"
        ],
        "conns": 0
      },
      {
        "id": 1948,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 585728,
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
        "id": 1950,
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
        "id": 1952,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 585728,
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
        "id": 1954,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 585728,
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
        "id": 1956,
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
        "id": 1964,
        "pid": 506,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2605056,
        "vms": 12533760,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "/sbin/udevd",
          "-d"
        ],
        "conns": 1
      },
      {
        "id": 1965,
        "pid": 506,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2605056,
        "vms": 12533760,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "/sbin/udevd",
          "-d"
        ],
        "conns": 1
      },
      {
        "id": 6740,
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
        "id": 6757,
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
        "id": 31491,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1810432,
        "vms": 66039808,
        "swap": 0,
        "memory_usage": 0.05,
        "cmd": [
          "su",
          "root",
          "-c",
          "/opt/metrics-agent/metrics \"-conf\" \"/opt/metrics-agent/agent.conf\""
        ],
        "conns": 1
      },
      {
        "id": 31492,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 540672,
        "vms": 4149248,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "/sbin/mingetty",
          "/dev/tty1"
        ],
        "conns": 0
      },
      {
        "id": 31493,
        "pid": 31491,
        "user": "root",
        "cpu_usage": 0.04,
        "rss": 51777536,
        "vms": 1155899392,
        "swap": 0,
        "memory_usage": 1.29,
        "cmd": [
          "/opt/metrics-agent/metrics",
          "-conf",
          "/opt/metrics-agent/agent.conf"
        ],
        "conns": 1
      }
    ],
    "connections": [
      {
        "fd": 3,
        "pid": 1879,
        "type": "tcp4",
        "local": "0.0.0.0:22",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 7,
        "pid": 1768,
        "type": "tcp4",
        "local": "127.0.0.1:631",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 9,
        "pid": 1554,
        "type": "tcp4",
        "local": "0.0.0.0:47389",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 8,
        "pid": 1515,
        "type": "tcp4",
        "local": "0.0.0.0:111",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 7,
        "pid": 31493,
        "type": "tcp4",
        "local": "192.168.2.65:49227",
        "remote": "192.168.2.52:13081",
        "status": "ESTABLISHED"
      },
      {
        "fd": 4,
        "pid": 1879,
        "type": "tcp6",
        "local": ":::22",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 6,
        "pid": 1768,
        "type": "tcp6",
        "local": "::1:631",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 11,
        "pid": 1515,
        "type": "tcp6",
        "local": ":::111",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 11,
        "pid": 1554,
        "type": "tcp6",
        "local": ":::38738",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 8,
        "pid": 1554,
        "type": "udp4",
        "local": "0.0.0.0:56416",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 13,
        "pid": 1757,
        "type": "udp4",
        "local": "0.0.0.0:5353",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 14,
        "pid": 1757,
        "type": "udp4",
        "local": "0.0.0.0:54762",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 6,
        "pid": 1515,
        "type": "udp4",
        "local": "0.0.0.0:111",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 5,
        "pid": 1554,
        "type": "udp4",
        "local": "0.0.0.0:882",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 9,
        "pid": 1768,
        "type": "udp4",
        "local": "0.0.0.0:631",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 7,
        "pid": 1515,
        "type": "udp4",
        "local": "0.0.0.0:842",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 10,
        "pid": 1554,
        "type": "udp6",
        "local": ":::58846",
        "remote": "::",
        "status": "NONE"
      },
      {
        "fd": 9,
        "pid": 1515,
        "type": "udp6",
        "local": ":::111",
        "remote": "::",
        "status": "NONE"
      },
      {
        "fd": 10,
        "pid": 1515,
        "type": "udp6",
        "local": ":::842",
        "remote": "::",
        "status": "NONE"
      },
      {
        "fd": 0,
        "pid": 1441,
        "type": "file",
        "local": "/dev/log",
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
        "fd": 4,
        "pid": 506,
        "type": "file",
        "local": "@/org/kernel/udev/udevd",
        "status": "NONE"
      },
      {
        "fd": 13,
        "pid": 1802,
        "type": "file",
        "local": "@/org/freedesktop/hal/udev_event",
        "status": "NONE"
      },
      {
        "fd": 8,
        "pid": 1768,
        "type": "file",
        "local": "/var/run/cups/cups.sock",
        "status": "NONE"
      },
      {
        "fd": 5,
        "pid": 1515,
        "type": "file",
        "local": "/var/run/rpcbind.sock",
        "status": "NONE"
      },
      {
        "fd": 7,
        "pid": 1474,
        "type": "file",
        "local": "/var/run/vmware/guestServicePipe",
        "status": "NONE"
      },
      {
        "fd": 3,
        "pid": 1745,
        "type": "file",
        "local": "/var/run/dbus/system_bus_socket",
        "status": "NONE"
      },
      {
        "fd": 7,
        "pid": 1802,
        "type": "file",
        "local": "@/var/run/hald/dbus-SAYA6iyJqy",
        "status": "NONE"
      },
      {
        "fd": 10,
        "pid": 1757,
        "type": "file",
        "local": "/var/run/avahi-daemon/socket",
        "status": "NONE"
      },
      {
        "fd": 4,
        "pid": 1793,
        "type": "file",
        "local": "/var/run/acpid.socket",
        "status": "NONE"
      },
      {
        "fd": 8,
        "pid": 1898,
        "type": "file",
        "local": "/var/run/abrt/abrt.socket",
        "status": "NONE"
      },
      {
        "fd": 9,
        "pid": 1802,
        "type": "file",
        "local": "@/var/run/hald/dbus-38Y2Js9SMW",
        "status": "NONE"
      },
      {
        "fd": 3,
        "pid": 31491,
        "type": "file",
        "status": "NONE"
      },
      {
        "fd": 9,
        "pid": 1898,
        "type": "file",
        "status": "NONE"
      }
    ]
  }
}
```