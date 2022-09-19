static:

```json
{
  "code": 200,
  "payload": {
    "time": 1663568771,
    "host": {
      "name": "localhost.localdomain",
      "uptime": 356829
    },
    "os": {
      "name": "linux",
      "platform_name": "redhat",
      "platform_version": "7.9",
      "install": 1663210990,
      "startup": 1663211942
    },
    "kernel": {
      "version": "3.10.0-1160.el7.x86_64",
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
          "physical": 2,
          "mhz": 2297.34
        }
      ]
    },
    "memory": {
      "physical": 3973709824,
      "swap": 4160745472
    },
    "disks": [
      {
        "model": "Virtual_disk",
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
        "fstype": "xfs",
        "opts": [
          "rw",
          "relatime"
        ],
        "total": 48420556800,
        "inodes": 23654400
      },
      {
        "mount": "/boot",
        "fstype": "xfs",
        "opts": [
          "rw",
          "relatime"
        ],
        "total": 1063256064,
        "inodes": 524288
      }
    ],
    "gateway": "192.168.2.1",
    "interface": [
      {
        "index": 1,
        "name": "lo",
        "mtu": 65536,
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
        "name": "ens192",
        "mtu": 1500,
        "flags": [
          "up",
          "broadcast",
          "multicast"
        ],
        "mac": "00:50:56:a3:c1:e9",
        "addrs": [
          "192.168.2.66/24",
          "fe80::623f:395a:b4d5:c3c9/64"
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
        "name": "systemd-network",
        "id": "192",
        "gid": "192"
      },
      {
        "name": "dbus",
        "id": "81",
        "gid": "81"
      },
      {
        "name": "polkitd",
        "id": "999",
        "gid": "998"
      },
      {
        "name": "sshd",
        "id": "74",
        "gid": "74"
      },
      {
        "name": "postfix",
        "id": "89",
        "gid": "89"
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
        "usage": 0.12
      },
      "memory": {
        "used": 315367424,
        "free": 3471564800,
        "available": 3437207552,
        "usage": 7.94,
        "swap_used": 0,
        "swap_free": 4160745472
      },
      "partitions": [
        {
          "name": "/",
          "used": 1363775488,
          "free": 47056781312,
          "usage": 2.82,
          "inode_used": 29876,
          "inode_free": 23624524,
          "inode_usage": 0.13
        },
        {
          "name": "/boot",
          "used": 156983296,
          "free": 906272768,
          "usage": 14.76,
          "inode_used": 325,
          "inode_free": 523963,
          "inode_usage": 0.06
        }
      ],
      "interface": [
        {
          "name": "ens192",
          "bytes_sent": 9421118,
          "bytes_recv": 48083331,
          "packets_sent": 72720,
          "packets_recv": 541499
        },
        {
          "name": "lo",
          "bytes_sent": 0,
          "bytes_recv": 0,
          "packets_sent": 0,
          "packets_recv": 0
        }
      ]
    },
    "process": [
      {
        "id": 1,
        "pid": 0,
        "user": "root",
        "cpu_usage": 0,
        "rss": 6901760,
        "vms": 131178496,
        "swap": 0,
        "memory_usage": 0.17,
        "cmd": [
          "/usr/lib/systemd/systemd",
          "--switched-root",
          "--system",
          "--deserialize",
          "22"
        ],
        "conns": 13
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
        "cpu_usage": 0.01,
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
        "id": 46,
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
        "id": 48,
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
        "id": 66,
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
        "id": 103,
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
        "id": 275,
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
        "id": 279,
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
        "id": 281,
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
        "id": 283,
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
        "id": 289,
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
        "id": 291,
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
        "id": 292,
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
        "id": 323,
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
        "id": 324,
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
        "id": 325,
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
        "id": 326,
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
        "id": 327,
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
        "id": 328,
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
        "id": 329,
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
        "id": 330,
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
        "id": 331,
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
        "id": 332,
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
        "id": 333,
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
        "id": 335,
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
        "id": 336,
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
        "id": 337,
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
        "id": 339,
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
        "id": 340,
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
        "id": 342,
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
        "id": 343,
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
        "id": 344,
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
        "id": 345,
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
        "id": 346,
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
        "id": 347,
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
        "id": 348,
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
        "id": 349,
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
        "id": 350,
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
        "id": 351,
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
        "id": 352,
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
        "id": 356,
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
        "id": 357,
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
        "id": 358,
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
        "id": 359,
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
        "id": 360,
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
        "id": 361,
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
        "id": 362,
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
        "id": 363,
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
        "id": 391,
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
        "id": 392,
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
        "id": 393,
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
        "id": 396,
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
        "id": 397,
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
        "id": 406,
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
        "id": 467,
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
        "id": 468,
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
        "id": 477,
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
        "id": 478,
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
        "id": 493,
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
        "id": 494,
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
        "id": 495,
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
        "id": 496,
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
        "id": 497,
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
        "id": 498,
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
        "id": 499,
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
        "id": 500,
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
        "id": 501,
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
        "id": 502,
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
        "id": 503,
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
        "id": 504,
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
        "id": 583,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3960832,
        "vms": 39993344,
        "swap": 0,
        "memory_usage": 0.1,
        "cmd": [
          "/usr/lib/systemd/systemd-journald"
        ],
        "conns": 4
      },
      {
        "id": 610,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5885952,
        "vms": 50421760,
        "swap": 0,
        "memory_usage": 0.15,
        "cmd": [
          "/usr/lib/systemd/systemd-udevd"
        ],
        "conns": 3
      },
      {
        "id": 612,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 4169728,
        "vms": 205877248,
        "swap": 0,
        "memory_usage": 0.1,
        "cmd": [
          "/usr/sbin/lvmetad",
          "-f"
        ],
        "conns": 2
      },
      {
        "id": 705,
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
        "id": 706,
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
        "id": 707,
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
        "id": 708,
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
        "id": 709,
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
        "id": 710,
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
        "id": 711,
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
        "id": 712,
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
        "id": 731,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 872448,
        "vms": 56864768,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/auditd"
        ],
        "conns": 2
      },
      {
        "id": 756,
        "pid": 1,
        "user": "polkitd",
        "cpu_usage": 0,
        "rss": 15388672,
        "vms": 627716096,
        "swap": 0,
        "memory_usage": 0.39,
        "cmd": [
          "/usr/lib/polkit-1/polkitd",
          "--no-debug"
        ],
        "conns": 2
      },
      {
        "id": 758,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5140480,
        "vms": 172187648,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          "/usr/bin/VGAuthService",
          "-s"
        ],
        "conns": 3
      },
      {
        "id": 759,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.17,
        "rss": 5066752,
        "vms": 357171200,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          "/usr/bin/vmtoolsd"
        ],
        "conns": 1
      },
      {
        "id": 761,
        "pid": 1,
        "user": "dbus",
        "cpu_usage": 0,
        "rss": 2637824,
        "vms": 68042752,
        "swap": 0,
        "memory_usage": 0.07,
        "cmd": [
          "/usr/bin/dbus-daemon",
          "--system",
          "--address=systemd:",
          "--nofork",
          "--nopidfile",
          "--systemd-activation"
        ],
        "conns": 3
      },
      {
        "id": 769,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1794048,
        "vms": 27017216,
        "swap": 0,
        "memory_usage": 0.05,
        "cmd": [
          "/usr/lib/systemd/systemd-logind"
        ],
        "conns": 2
      },
      {
        "id": 770,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 1335296,
        "vms": 22200320,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "/usr/sbin/irqbalance",
          "--foreground"
        ],
        "conns": 1
      },
      {
        "id": 774,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1699840,
        "vms": 129417216,
        "swap": 0,
        "memory_usage": 0.04,
        "cmd": [
          "/usr/sbin/crond",
          "-n"
        ],
        "conns": 2
      },
      {
        "id": 777,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 868352,
        "vms": 112848896,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/agetty",
          "--noclear",
          "tty1",
          "linux"
        ],
        "conns": 0
      },
      {
        "id": 798,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 30236672,
        "vms": 367353856,
        "swap": 0,
        "memory_usage": 0.76,
        "cmd": [
          "/usr/bin/python2",
          "-Es",
          "/usr/sbin/firewalld",
          "--nofork",
          "--nopid"
        ],
        "conns": 2
      },
      {
        "id": 823,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.03,
        "rss": 11165696,
        "vms": 640913408,
        "swap": 0,
        "memory_usage": 0.28,
        "cmd": [
          "/usr/sbin/NetworkManager",
          "--no-daemon"
        ],
        "conns": 2
      },
      {
        "id": 1121,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.02,
        "rss": 20574208,
        "vms": 600498176,
        "swap": 0,
        "memory_usage": 0.52,
        "cmd": [
          "/usr/bin/python2",
          "-Es",
          "/usr/sbin/tuned",
          "-l",
          "-P"
        ],
        "conns": 1
      },
      {
        "id": 1124,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 4407296,
        "vms": 115609600,
        "swap": 0,
        "memory_usage": 0.11,
        "cmd": [
          "/usr/sbin/sshd",
          "-D"
        ],
        "listen": [
          22,
          22
        ],
        "conns": 3
      },
      {
        "id": 1125,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 6107136,
        "vms": 221593600,
        "swap": 0,
        "memory_usage": 0.15,
        "cmd": [
          "/usr/sbin/rsyslogd",
          "-n"
        ],
        "conns": 1
      },
      {
        "id": 1128,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 716800,
        "vms": 118587392,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/usr/bin/rhsmcertd"
        ],
        "conns": 0
      },
      {
        "id": 1154,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 634880,
        "vms": 110661632,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "rhnsd"
        ],
        "conns": 1
      },
      {
        "id": 1318,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2273280,
        "vms": 91856896,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "/usr/libexec/postfix/master",
          "-w"
        ],
        "listen": [
          25,
          25
        ],
        "conns": 27
      },
      {
        "id": 1366,
        "pid": 1318,
        "user": "postfix",
        "cpu_usage": 0,
        "rss": 4173824,
        "vms": 92033024,
        "swap": 0,
        "memory_usage": 0.11,
        "cmd": [
          "qmgr",
          "-l",
          "-t",
          "unix",
          "-u"
        ],
        "conns": 3
      },
      {
        "id": 11063,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.04,
        "rss": 36753408,
        "vms": 1083092992,
        "swap": 0,
        "memory_usage": 0.92,
        "cmd": [
          "/opt/metrics-agent/metrics",
          "-conf",
          "/opt/metrics-agent/agent.conf"
        ],
        "conns": 2
      },
      {
        "id": 12142,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 12242,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 12266,
        "pid": 1318,
        "user": "postfix",
        "cpu_usage": 0,
        "rss": 4149248,
        "vms": 91963392,
        "swap": 0,
        "memory_usage": 0.1,
        "cmd": [
          "pickup",
          "-l",
          "-t",
          "unix",
          "-u"
        ],
        "conns": 3
      },
      {
        "id": 12293,
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
        "id": 12296,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 12300,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      }
    ],
    "connections": [
      {
        "fd": 3,
        "pid": 1124,
        "type": "tcp4",
        "local": "0.0.0.0:22",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 13,
        "pid": 1318,
        "type": "tcp4",
        "local": "127.0.0.1:25",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 3,
        "pid": 11063,
        "type": "tcp4",
        "local": "192.168.2.66:36872",
        "remote": "192.168.2.52:13081",
        "status": "ESTABLISHED"
      },
      {
        "fd": 4,
        "pid": 1124,
        "type": "tcp6",
        "local": ":::22",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 14,
        "pid": 1318,
        "type": "tcp6",
        "local": "::1:25",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 8,
        "pid": 758,
        "type": "file",
        "local": "/var/run/vmware/guestServicePipe",
        "status": "NONE"
      },
      {
        "fd": 50,
        "pid": 1318,
        "type": "file",
        "local": "private/proxymap",
        "status": "NONE"
      },
      {
        "fd": 53,
        "pid": 1318,
        "type": "file",
        "local": "private/proxywrite",
        "status": "NONE"
      },
      {
        "fd": 56,
        "pid": 1318,
        "type": "file",
        "local": "private/smtp",
        "status": "NONE"
      },
      {
        "fd": 59,
        "pid": 1318,
        "type": "file",
        "local": "private/relay",
        "status": "NONE"
      },
      {
        "fd": 65,
        "pid": 1318,
        "type": "file",
        "local": "private/error",
        "status": "NONE"
      },
      {
        "fd": 68,
        "pid": 1318,
        "type": "file",
        "local": "private/retry",
        "status": "NONE"
      },
      {
        "fd": 22,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/notify",
        "status": "NONE"
      },
      {
        "fd": 71,
        "pid": 1318,
        "type": "file",
        "local": "private/discard",
        "status": "NONE"
      },
      {
        "fd": 74,
        "pid": 1318,
        "type": "file",
        "local": "private/local",
        "status": "NONE"
      },
      {
        "fd": 23,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/cgroups-agent",
        "status": "NONE"
      },
      {
        "fd": 77,
        "pid": 1318,
        "type": "file",
        "local": "private/virtual",
        "status": "NONE"
      },
      {
        "fd": 80,
        "pid": 1318,
        "type": "file",
        "local": "private/lmtp",
        "status": "NONE"
      },
      {
        "fd": 83,
        "pid": 1318,
        "type": "file",
        "local": "private/anvil",
        "status": "NONE"
      },
      {
        "fd": 86,
        "pid": 1318,
        "type": "file",
        "local": "private/scache",
        "status": "NONE"
      },
      {
        "fd": 24,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/stdout",
        "status": "NONE"
      },
      {
        "fd": 25,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/socket",
        "status": "NONE"
      },
      {
        "fd": 26,
        "pid": 1,
        "type": "file",
        "local": "/dev/log",
        "status": "NONE"
      },
      {
        "fd": 12,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/private",
        "status": "NONE"
      },
      {
        "fd": 27,
        "pid": 1,
        "type": "file",
        "local": "/run/dbus/system_bus_socket",
        "status": "NONE"
      },
      {
        "fd": 18,
        "pid": 1318,
        "type": "file",
        "local": "public/pickup",
        "status": "NONE"
      },
      {
        "fd": 22,
        "pid": 1318,
        "type": "file",
        "local": "public/cleanup",
        "status": "NONE"
      },
      {
        "fd": 25,
        "pid": 1318,
        "type": "file",
        "local": "public/qmgr",
        "status": "NONE"
      },
      {
        "fd": 47,
        "pid": 1318,
        "type": "file",
        "local": "public/flush",
        "status": "NONE"
      },
      {
        "fd": 62,
        "pid": 1318,
        "type": "file",
        "local": "public/showq",
        "status": "NONE"
      },
      {
        "fd": 30,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/shutdownd",
        "status": "NONE"
      },
      {
        "fd": 28,
        "pid": 1,
        "type": "file",
        "local": "/run/lvm/lvmpolld.socket",
        "status": "NONE"
      },
      {
        "fd": 33,
        "pid": 1,
        "type": "unix",
        "local": "/run/udev/control",
        "status": "NONE"
      },
      {
        "fd": 31,
        "pid": 1,
        "type": "file",
        "local": "/run/lvm/lvmetad.socket",
        "status": "NONE"
      },
      {
        "fd": 29,
        "pid": 1318,
        "type": "file",
        "local": "private/tlsmgr",
        "status": "NONE"
      },
      {
        "fd": 32,
        "pid": 1318,
        "type": "file",
        "local": "private/rewrite",
        "status": "NONE"
      },
      {
        "fd": 35,
        "pid": 1318,
        "type": "file",
        "local": "private/bounce",
        "status": "NONE"
      },
      {
        "fd": 38,
        "pid": 1318,
        "type": "file",
        "local": "private/defer",
        "status": "NONE"
      },
      {
        "fd": 41,
        "pid": 1318,
        "type": "file",
        "local": "private/trace",
        "status": "NONE"
      },
      {
        "fd": 44,
        "pid": 1318,
        "type": "file",
        "local": "private/verify",
        "status": "NONE"
      },
      {
        "fd": 12,
        "pid": 769,
        "type": "file",
        "status": "NONE"
      },
      {
        "fd": 8,
        "pid": 731,
        "type": "file",
        "status": "NONE"
      }
    ]
  }
}
```