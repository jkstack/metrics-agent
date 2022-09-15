static:

```json
{
  "code": 200,
  "payload": {
    "time": 1663223328,
    "host": {
      "name": "localhost.localdomain",
      "uptime": 532496
    },
    "os": {
      "name": "linux",
      "platform_name": "centos",
      "platform_version": "7.9.2009",
      "install": 1662689273,
      "startup": 1662690833
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
        "name": "ens160",
        "mtu": 1500,
        "flags": [
          "up",
          "broadcast",
          "multicast"
        ],
        "mac": "00:50:56:a3:92:d3",
        "addrs": [
          "192.168.2.52/24",
          "fe80::8ac3:eaa:8e90:ed38/64"
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
        "name": "systemd Network Management",
        "id": "192",
        "gid": "192"
      },
      {
        "name": "System message bus",
        "id": "81",
        "gid": "81"
      },
      {
        "name": "User for polkitd",
        "id": "999",
        "gid": "998"
      },
      {
        "name": "Privilege-separated SSH",
        "id": "74",
        "gid": "74"
      },
      {
        "name": "",
        "id": "89",
        "gid": "89"
      },
      {
        "name": "",
        "id": "998",
        "gid": "996"
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
        "usage": 0.78
      },
      "memory": {
        "used": 305504256,
        "free": 3213684736,
        "available": 3433676800,
        "usage": 7.69,
        "swap_used": 0,
        "swap_free": 4160745472
      },
      "partitions": [
        {
          "name": "/",
          "used": 1539723264,
          "free": 46880833536,
          "usage": 3.18,
          "inode_used": 27767,
          "inode_free": 23626633,
          "inode_usage": 0.12
        },
        {
          "name": "/boot",
          "used": 157421568,
          "free": 905834496,
          "usage": 14.81,
          "inode_used": 327,
          "inode_free": 523961,
          "inode_usage": 0.06
        }
      ],
      "interface": [
        {
          "name": "lo",
          "bytes_sent": 10803730,
          "bytes_recv": 10803730,
          "packets_sent": 111600,
          "packets_recv": 111600
        },
        {
          "name": "ens160",
          "bytes_sent": 7157769,
          "bytes_recv": 126875150,
          "packets_sent": 77224,
          "packets_recv": 843048
        }
      ]
    },
    "process": [
      {
        "id": 1,
        "pid": 0,
        "user": "root",
        "cpu_usage": 0,
        "rss": 7004160,
        "vms": 131223552,
        "swap": 0,
        "memory_usage": 0.18,
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
        "id": 47,
        "pid": 2,
        "user": "root",
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
        "id": 102,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 286,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 287,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 290,
        "pid": 2,
        "user": "root",
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
        "id": 377,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 378,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 387,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 388,
        "pid": 2,
        "user": "root",
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
        "id": 403,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 404,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 405,
        "pid": 2,
        "user": "root",
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
        "id": 407,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 408,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 409,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 410,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 411,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 412,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0.07,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 413,
        "pid": 2,
        "user": "root",
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
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 4386816,
        "vms": 39993344,
        "swap": 0,
        "memory_usage": 0.11,
        "cmd": [
          "/usr/lib/systemd/systemd-journald"
        ],
        "conns": 4
      },
      {
        "id": 523,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 4222976,
        "vms": 205930496,
        "swap": 0,
        "memory_usage": 0.11,
        "cmd": [
          "/usr/sbin/lvmetad",
          "-f"
        ],
        "conns": 2
      },
      {
        "id": 529,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5775360,
        "vms": 50339840,
        "swap": 0,
        "memory_usage": 0.15,
        "cmd": [
          "/usr/lib/systemd/systemd-udevd"
        ],
        "conns": 3
      },
      {
        "id": 615,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 616,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 617,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 618,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 619,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 620,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 621,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 622,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 646,
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
        "id": 669,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 1347584,
        "vms": 22204416,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "/usr/sbin/irqbalance",
          "--foreground"
        ],
        "conns": 1
      },
      {
        "id": 670,
        "pid": 1,
        "user": "polkitd",
        "cpu_usage": 0,
        "rss": 11235328,
        "vms": 627716096,
        "swap": 0,
        "memory_usage": 0.28,
        "cmd": [
          "/usr/lib/polkit-1/polkitd",
          "--no-debug"
        ],
        "conns": 2
      },
      {
        "id": 672,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5283840,
        "vms": 172343296,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          "/usr/bin/VGAuthService",
          "-s"
        ],
        "conns": 3
      },
      {
        "id": 673,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.16,
        "rss": 5115904,
        "vms": 357363712,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          "/usr/bin/vmtoolsd"
        ],
        "conns": 1
      },
      {
        "id": 679,
        "pid": 1,
        "user": "dbus",
        "cpu_usage": 0,
        "rss": 2670592,
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
        "id": 682,
        "pid": 1,
        "user": "chrony",
        "cpu_usage": 0,
        "rss": 1867776,
        "vms": 120635392,
        "swap": 0,
        "memory_usage": 0.05,
        "cmd": [
          "/usr/sbin/chronyd"
        ],
        "conns": 4
      },
      {
        "id": 692,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1843200,
        "vms": 27058176,
        "swap": 0,
        "memory_usage": 0.05,
        "cmd": [
          "/usr/lib/systemd/systemd-logind"
        ],
        "conns": 2
      },
      {
        "id": 699,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1720320,
        "vms": 129421312,
        "swap": 0,
        "memory_usage": 0.04,
        "cmd": [
          "/usr/sbin/crond",
          "-n"
        ],
        "conns": 2
      },
      {
        "id": 746,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.03,
        "rss": 11370496,
        "vms": 565989376,
        "swap": 0,
        "memory_usage": 0.29,
        "cmd": [
          "/usr/sbin/NetworkManager",
          "--no-daemon"
        ],
        "conns": 2
      },
      {
        "id": 1055,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.02,
        "rss": 19959808,
        "vms": 588062720,
        "swap": 0,
        "memory_usage": 0.5,
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
        "id": 1056,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 4452352,
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
        "id": 1059,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 9101312,
        "vms": 221593600,
        "swap": 0,
        "memory_usage": 0.23,
        "cmd": [
          "/usr/sbin/rsyslogd",
          "-n"
        ],
        "conns": 1
      },
      {
        "id": 1242,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2273280,
        "vms": 91860992,
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
        "id": 1248,
        "pid": 1242,
        "user": "postfix",
        "cpu_usage": 0,
        "rss": 4173824,
        "vms": 92037120,
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
        "id": 2037,
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
        "id": 13042,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.06,
        "rss": 45187072,
        "vms": 955719680,
        "swap": 0,
        "memory_usage": 1.14,
        "cmd": [
          "/opt/agent-server/server",
          "-conf",
          "/opt/agent-server/server.conf"
        ],
        "listen": [
          13081
        ],
        "conns": 6
      },
      {
        "id": 13144,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.04,
        "rss": 39297024,
        "vms": 1141788672,
        "swap": 0,
        "memory_usage": 0.99,
        "cmd": [
          "/opt/metrics-agent/metrics",
          "-conf",
          "/opt/metrics-agent/agent.conf"
        ],
        "conns": 2
      },
      {
        "id": 13208,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0.02,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 13247,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 13335,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 13386,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0.02,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 14552,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 14883,
        "pid": 1242,
        "user": "postfix",
        "cpu_usage": 0,
        "rss": 4149248,
        "vms": 91967488,
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
        "id": 14923,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 14924,
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
        "pid": 1056,
        "type": "tcp4",
        "local": "0.0.0.0:22",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 13,
        "pid": 1242,
        "type": "tcp4",
        "local": "127.0.0.1:25",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 3,
        "pid": 13144,
        "type": "tcp4",
        "local": "127.0.0.1:55374",
        "remote": "127.0.0.1:13081",
        "status": "ESTABLISHED"
      },
      {
        "fd": 4,
        "pid": 1056,
        "type": "tcp6",
        "local": ":::22",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 7,
        "pid": 13042,
        "type": "tcp6",
        "local": ":::13081",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 14,
        "pid": 1242,
        "type": "tcp6",
        "local": "::1:25",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 11,
        "pid": 13042,
        "type": "tcp6",
        "local": "192.168.2.52:13081",
        "remote": "10.202.0.227:54162",
        "status": "ESTABLISHED"
      },
      {
        "fd": 8,
        "pid": 13042,
        "type": "tcp6",
        "local": "127.0.0.1:13081",
        "remote": "127.0.0.1:55374",
        "status": "ESTABLISHED"
      },
      {
        "fd": 10,
        "pid": 13042,
        "type": "tcp6",
        "local": "192.168.2.52:13081",
        "remote": "192.168.2.54:40840",
        "status": "ESTABLISHED"
      },
      {
        "fd": 9,
        "pid": 13042,
        "type": "tcp6",
        "local": "192.168.2.52:13081",
        "remote": "10.202.0.9:63556",
        "status": "ESTABLISHED"
      },
      {
        "fd": 5,
        "pid": 682,
        "type": "udp4",
        "local": "127.0.0.1:323",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 6,
        "pid": 682,
        "type": "udp6",
        "local": "::1:323",
        "remote": "::",
        "status": "NONE"
      },
      {
        "fd": 31,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/stdout",
        "status": "NONE"
      },
      {
        "fd": 32,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/socket",
        "status": "NONE"
      },
      {
        "fd": 33,
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
        "fd": 75,
        "pid": 1,
        "type": "file",
        "local": "/run/dbus/system_bus_socket",
        "status": "NONE"
      },
      {
        "fd": 72,
        "pid": 1,
        "type": "file",
        "local": "/run/lvm/lvmpolld.socket",
        "status": "NONE"
      },
      {
        "fd": 77,
        "pid": 1,
        "type": "unix",
        "local": "/run/udev/control",
        "status": "NONE"
      },
      {
        "fd": 8,
        "pid": 682,
        "type": "file",
        "local": "/var/run/chrony/chronyd.sock",
        "status": "NONE"
      },
      {
        "fd": 74,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/shutdownd",
        "status": "NONE"
      },
      {
        "fd": 73,
        "pid": 1,
        "type": "file",
        "local": "/run/lvm/lvmetad.socket",
        "status": "NONE"
      },
      {
        "fd": 8,
        "pid": 672,
        "type": "file",
        "local": "/var/run/vmware/guestServicePipe",
        "status": "NONE"
      },
      {
        "fd": 44,
        "pid": 1242,
        "type": "file",
        "local": "private/verify",
        "status": "NONE"
      },
      {
        "fd": 50,
        "pid": 1242,
        "type": "file",
        "local": "private/proxymap",
        "status": "NONE"
      },
      {
        "fd": 53,
        "pid": 1242,
        "type": "file",
        "local": "private/proxywrite",
        "status": "NONE"
      },
      {
        "fd": 56,
        "pid": 1242,
        "type": "file",
        "local": "private/smtp",
        "status": "NONE"
      },
      {
        "fd": 59,
        "pid": 1242,
        "type": "file",
        "local": "private/relay",
        "status": "NONE"
      },
      {
        "fd": 65,
        "pid": 1242,
        "type": "file",
        "local": "private/error",
        "status": "NONE"
      },
      {
        "fd": 68,
        "pid": 1242,
        "type": "file",
        "local": "private/retry",
        "status": "NONE"
      },
      {
        "fd": 18,
        "pid": 1242,
        "type": "file",
        "local": "public/pickup",
        "status": "NONE"
      },
      {
        "fd": 71,
        "pid": 1242,
        "type": "file",
        "local": "private/discard",
        "status": "NONE"
      },
      {
        "fd": 22,
        "pid": 1242,
        "type": "file",
        "local": "public/cleanup",
        "status": "NONE"
      },
      {
        "fd": 74,
        "pid": 1242,
        "type": "file",
        "local": "private/local",
        "status": "NONE"
      },
      {
        "fd": 25,
        "pid": 1242,
        "type": "file",
        "local": "public/qmgr",
        "status": "NONE"
      },
      {
        "fd": 77,
        "pid": 1242,
        "type": "file",
        "local": "private/virtual",
        "status": "NONE"
      },
      {
        "fd": 47,
        "pid": 1242,
        "type": "file",
        "local": "public/flush",
        "status": "NONE"
      },
      {
        "fd": 80,
        "pid": 1242,
        "type": "file",
        "local": "private/lmtp",
        "status": "NONE"
      },
      {
        "fd": 62,
        "pid": 1242,
        "type": "file",
        "local": "public/showq",
        "status": "NONE"
      },
      {
        "fd": 83,
        "pid": 1242,
        "type": "file",
        "local": "private/anvil",
        "status": "NONE"
      },
      {
        "fd": 86,
        "pid": 1242,
        "type": "file",
        "local": "private/scache",
        "status": "NONE"
      },
      {
        "fd": 29,
        "pid": 1242,
        "type": "file",
        "local": "private/tlsmgr",
        "status": "NONE"
      },
      {
        "fd": 32,
        "pid": 1242,
        "type": "file",
        "local": "private/rewrite",
        "status": "NONE"
      },
      {
        "fd": 35,
        "pid": 1242,
        "type": "file",
        "local": "private/bounce",
        "status": "NONE"
      },
      {
        "fd": 38,
        "pid": 1242,
        "type": "file",
        "local": "private/defer",
        "status": "NONE"
      },
      {
        "fd": 41,
        "pid": 1242,
        "type": "file",
        "local": "private/trace",
        "status": "NONE"
      },
      {
        "fd": 29,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/notify",
        "status": "NONE"
      },
      {
        "fd": 30,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/cgroups-agent",
        "status": "NONE"
      },
      {
        "fd": 58,
        "pid": 1242,
        "type": "file",
        "status": "NONE"
      },
      {
        "fd": 5,
        "pid": 746,
        "type": "file",
        "status": "NONE"
      }
    ]
  }
}
```