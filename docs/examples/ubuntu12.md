static:

```json
{
  "code": 200,
  "payload": {
    "time": 1663223734,
    "host": {
      "name": "ubuntu12",
      "uptime": 522313
    },
    "os": {
      "name": "linux",
      "platform_name": "ubuntu",
      "platform_version": "12.04",
      "install": 1662695452,
      "startup": 1662701422
    },
    "kernel": {
      "version": "3.11.0-15-generic",
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
      "physical": 4144664576,
      "swap": 4290768896
    },
    "disks": [
      {
        "model": "Virtual_disk",
        "total": 53687091200,
        "type": "HDD",
        "disks": [
          "/dev/sda1",
          "/dev/sda2",
          "/dev/sda5"
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
        "total": 48197730304,
        "inodes": 2998272
      },
      {
        "mount": "/boot",
        "fstype": "ext2",
        "opts": [
          "rw",
          "relatime"
        ],
        "total": 246755328,
        "inodes": 62248
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
        "name": "eth0",
        "mtu": 1500,
        "flags": [
          "up",
          "broadcast",
          "multicast"
        ],
        "mac": "00:50:56:a3:54:34",
        "addrs": [
          "192.168.2.54/24",
          "fe80::250:56ff:fea3:5434/64"
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
        "name": "daemon",
        "id": "1",
        "gid": "1"
      },
      {
        "name": "bin",
        "id": "2",
        "gid": "2"
      },
      {
        "name": "sys",
        "id": "3",
        "gid": "3"
      },
      {
        "name": "sync",
        "id": "4",
        "gid": "65534"
      },
      {
        "name": "games",
        "id": "5",
        "gid": "60"
      },
      {
        "name": "man",
        "id": "6",
        "gid": "12"
      },
      {
        "name": "lp",
        "id": "7",
        "gid": "7"
      },
      {
        "name": "mail",
        "id": "8",
        "gid": "8"
      },
      {
        "name": "news",
        "id": "9",
        "gid": "9"
      },
      {
        "name": "uucp",
        "id": "10",
        "gid": "10"
      },
      {
        "name": "proxy",
        "id": "13",
        "gid": "13"
      },
      {
        "name": "www-data",
        "id": "33",
        "gid": "33"
      },
      {
        "name": "backup",
        "id": "34",
        "gid": "34"
      },
      {
        "name": "Mailing List Manager",
        "id": "38",
        "gid": "38"
      },
      {
        "name": "ircd",
        "id": "39",
        "gid": "39"
      },
      {
        "name": "Gnats Bug-Reporting System (admin)",
        "id": "41",
        "gid": "41"
      },
      {
        "name": "nobody",
        "id": "65534",
        "gid": "65534"
      },
      {
        "name": "",
        "id": "100",
        "gid": "101"
      },
      {
        "name": "",
        "id": "101",
        "gid": "103"
      },
      {
        "name": "",
        "id": "102",
        "gid": "105"
      },
      {
        "name": "",
        "id": "103",
        "gid": "108"
      },
      {
        "name": "",
        "id": "104",
        "gid": "110"
      },
      {
        "name": "",
        "id": "105",
        "gid": "65534"
      },
      {
        "name": "jkstack",
        "id": "1000",
        "gid": "1000"
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
        "usage": 50.12
      },
      "memory": {
        "used": 86929408,
        "free": 3912425472,
        "available": 3898466304,
        "usage": 2.1,
        "swap_used": 0,
        "swap_free": 4290768896
      },
      "partitions": [
        {
          "name": "/",
          "used": 828039168,
          "free": 44897779712,
          "usage": 1.81,
          "inode_used": 56497,
          "inode_free": 2941775,
          "inode_usage": 1.88
        },
        {
          "name": "/boot",
          "used": 32727040,
          "free": 201288704,
          "usage": 13.98,
          "inode_used": 227,
          "inode_free": 62021,
          "inode_usage": 0.36
        }
      ],
      "interface": [
        {
          "name": "eth0",
          "bytes_sent": 6613306,
          "bytes_recv": 92954544,
          "packets_sent": 37707,
          "packets_recv": 775094
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
        "rss": 2281472,
        "vms": 24928256,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "/sbin/init"
        ],
        "conns": 2
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
        "id": 52,
        "pid": 2,
        "user": "root",
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
        "id": 55,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 57,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 58,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 78,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 79,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 216,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 217,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 222,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 223,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 224,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 225,
        "pid": 2,
        "user": "root",
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
        "id": 227,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 228,
        "pid": 2,
        "user": "root",
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
        "id": 230,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 231,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 233,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 235,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 236,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 239,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 240,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 241,
        "pid": 2,
        "user": "root",
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
        "id": 244,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 245,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 246,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 247,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 248,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 249,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 250,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 251,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 252,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 253,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 254,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 255,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 284,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 373,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 374,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 386,
        "pid": 2,
        "user": "root",
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
        "id": 395,
        "pid": 2,
        "user": "root",
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
        "id": 505,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 651264,
        "vms": 17653760,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "upstart-udev-bridge",
          "--daemon"
        ],
        "conns": 1
      },
      {
        "id": 516,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1728512,
        "vms": 22532096,
        "swap": 0,
        "memory_usage": 0.04,
        "cmd": [
          "/sbin/udevd",
          "--daemon"
        ],
        "conns": 2
      },
      {
        "id": 658,
        "pid": 1,
        "user": "messagebus",
        "cpu_usage": 0,
        "rss": 970752,
        "vms": 24395776,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "dbus-daemon",
          "--system",
          "--fork",
          "--activation=upstart"
        ],
        "conns": 2
      },
      {
        "id": 683,
        "pid": 1,
        "user": "syslog",
        "cpu_usage": 0,
        "rss": 1531904,
        "vms": 255463424,
        "swap": 0,
        "memory_usage": 0.04,
        "cmd": [
          "rsyslogd",
          "-c5"
        ],
        "conns": 1
      },
      {
        "id": 734,
        "pid": 516,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1220608,
        "vms": 22528000,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "/sbin/udevd",
          "--daemon"
        ],
        "conns": 1
      },
      {
        "id": 735,
        "pid": 516,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1228800,
        "vms": 22528000,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "/sbin/udevd",
          "--daemon"
        ],
        "conns": 1
      },
      {
        "id": 741,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 770,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 776,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 821,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 401408,
        "vms": 15560704,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "upstart-socket-bridge",
          "--daemon"
        ],
        "conns": 1
      },
      {
        "id": 906,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3002368,
        "vms": 51240960,
        "swap": 0,
        "memory_usage": 0.07,
        "cmd": [
          "/usr/sbin/sshd",
          "-D"
        ],
        "listen": [
          22,
          22
        ],
        "conns": 2
      },
      {
        "id": 1005,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 999424,
        "vms": 15511552,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/getty",
          "-8",
          "38400",
          "tty4"
        ],
        "conns": 0
      },
      {
        "id": 1011,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 995328,
        "vms": 15511552,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/getty",
          "-8",
          "38400",
          "tty5"
        ],
        "conns": 0
      },
      {
        "id": 1023,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 995328,
        "vms": 15511552,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/getty",
          "-8",
          "38400",
          "tty2"
        ],
        "conns": 0
      },
      {
        "id": 1024,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 995328,
        "vms": 15511552,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/getty",
          "-8",
          "38400",
          "tty3"
        ],
        "conns": 0
      },
      {
        "id": 1026,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 995328,
        "vms": 15511552,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/getty",
          "-8",
          "38400",
          "tty6"
        ],
        "conns": 0
      },
      {
        "id": 1031,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 724992,
        "vms": 4440064,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "acpid",
          "-c",
          "/etc/acpi/events",
          "-s",
          "/var/run/acpid.socket"
        ],
        "conns": 2
      },
      {
        "id": 1032,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 962560,
        "vms": 19578880,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "cron"
        ],
        "conns": 0
      },
      {
        "id": 1033,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 385024,
        "vms": 17321984,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "atd"
        ],
        "conns": 0
      },
      {
        "id": 1058,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 708608,
        "vms": 16371712,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/usr/sbin/irqbalance"
        ],
        "conns": 0
      },
      {
        "id": 1070,
        "pid": 1,
        "user": "root",
        "cpu_usage": 3.21,
        "rss": 995328,
        "vms": 15511552,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/getty",
          "-8",
          "38400",
          "tty1"
        ],
        "conns": 0
      },
      {
        "id": 1072,
        "pid": 1,
        "user": "whoopsie",
        "cpu_usage": 0,
        "rss": 4378624,
        "vms": 192172032,
        "swap": 0,
        "memory_usage": 0.11,
        "cmd": [
          "whoopsie"
        ],
        "conns": 1
      },
      {
        "id": 4316,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 7860,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1372160,
        "vms": 38871040,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "su",
          "root",
          "-c",
          "/opt/metrics-agent/metrics \"-conf\" \"/opt/metrics-agent/agent.conf\""
        ],
        "conns": 0
      },
      {
        "id": 7861,
        "pid": 7860,
        "user": "root",
        "cpu_usage": 0.05,
        "rss": 22134784,
        "vms": 1089245184,
        "swap": 0,
        "memory_usage": 0.53,
        "cmd": [
          "/opt/metrics-agent/metrics",
          "-conf",
          "/opt/metrics-agent/agent.conf"
        ],
        "conns": 1
      },
      {
        "id": 7949,
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
        "id": 7950,
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
        "id": 8123,
        "pid": 906,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 3756032,
        "vms": 75206656,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "sshd: root@pts/0"
        ],
        "conns": 2
      },
      {
        "id": 8175,
        "pid": 8123,
        "user": "root",
        "cpu_usage": 0.03,
        "rss": 4055040,
        "vms": 22765568,
        "swap": 0,
        "memory_usage": 0.1,
        "cmd": [
          "-bash"
        ],
        "conns": 0
      }
    ],
    "connections": [
      {
        "fd": 3,
        "pid": 906,
        "type": "tcp4",
        "local": "0.0.0.0:22",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 3,
        "pid": 8123,
        "type": "tcp4",
        "local": "192.168.2.54:22",
        "remote": "10.202.0.9:49561",
        "status": "ESTABLISHED"
      },
      {
        "fd": 7,
        "pid": 7861,
        "type": "tcp4",
        "local": "192.168.2.54:40840",
        "remote": "192.168.2.52:13081",
        "status": "ESTABLISHED"
      },
      {
        "fd": 4,
        "pid": 906,
        "type": "tcp6",
        "local": ":::22",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 3,
        "pid": 516,
        "type": "unix",
        "local": "/run/udev/control",
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
        "fd": 3,
        "pid": 658,
        "type": "file",
        "local": "/var/run/dbus/system_bus_socket",
        "status": "NONE"
      },
      {
        "fd": 5,
        "pid": 1031,
        "type": "file",
        "local": "/var/run/acpid.socket",
        "status": "NONE"
      },
      {
        "fd": 0,
        "pid": 683,
        "type": "file",
        "local": "/dev/log",
        "status": "NONE"
      },
      {
        "fd": 3,
        "pid": 1031,
        "type": "file",
        "status": "NONE"
      },
      {
        "fd": 3,
        "pid": 505,
        "type": "file",
        "status": "NONE"
      }
    ]
  }
}
```