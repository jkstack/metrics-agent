static:

```json
{
  "code": 200,
  "payload": {
    "time": 1663317744,
    "host": {
      "name": "jkstack",
      "uptime": 261382
    },
    "os": {
      "name": "linux",
      "platform_name": "ubuntu",
      "platform_version": "20.04",
      "install": 1663035074,
      "startup": 1663056362
    },
    "kernel": {
      "version": "5.4.0-125-generic",
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
      "physical": 4122755072,
      "swap": 4121948160
    },
    "disks": [
      {
        "model": "Virtual_disk",
        "total": 53687091200,
        "type": "HDD",
        "disks": [
          "/dev/sda1",
          "/dev/sda2",
          "/dev/sda3"
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
        "total": 25180848128,
        "inodes": 1572864
      },
      {
        "mount": "/snap/lxd/22753",
        "fstype": "squashfs",
        "opts": [
          "ro",
          "nodev",
          "relatime"
        ],
        "total": 71172096,
        "inodes": 802
      },
      {
        "mount": "/snap/core20/1611",
        "fstype": "squashfs",
        "opts": [
          "ro",
          "nodev",
          "relatime"
        ],
        "total": 65011712,
        "inodes": 11796
      },
      {
        "mount": "/snap/snapd/16292",
        "fstype": "squashfs",
        "opts": [
          "ro",
          "nodev",
          "relatime"
        ],
        "total": 49283072,
        "inodes": 486
      },
      {
        "mount": "/boot",
        "fstype": "ext4",
        "opts": [
          "rw",
          "relatime"
        ],
        "total": 2040373248,
        "inodes": 131072
      },
      {
        "mount": "/snap/core20/1623",
        "fstype": "squashfs",
        "opts": [
          "ro",
          "nodev",
          "relatime"
        ],
        "total": 66322432,
        "inodes": 11882
      },
      {
        "mount": "/snap/snapd/16778",
        "fstype": "squashfs",
        "opts": [
          "ro",
          "nodev",
          "relatime"
        ],
        "total": 50331648,
        "inodes": 486
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
        "mac": "00:50:56:a3:e5:13",
        "addrs": [
          "192.168.2.58/24",
          "fe80::250:56ff:fea3:e513/64"
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
        "name": "systemd Network Management",
        "id": "100",
        "gid": "102"
      },
      {
        "name": "systemd Resolver",
        "id": "101",
        "gid": "103"
      },
      {
        "name": "systemd Time Synchronization",
        "id": "102",
        "gid": "104"
      },
      {
        "name": "",
        "id": "103",
        "gid": "106"
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
        "name": "TPM software stack",
        "id": "106",
        "gid": "111"
      },
      {
        "name": "",
        "id": "107",
        "gid": "112"
      },
      {
        "name": "",
        "id": "108",
        "gid": "113"
      },
      {
        "name": "",
        "id": "109",
        "gid": "115"
      },
      {
        "name": "",
        "id": "110",
        "gid": "1"
      },
      {
        "name": "usbmux daemon",
        "id": "111",
        "gid": "46"
      },
      {
        "name": "",
        "id": "112",
        "gid": "65534"
      },
      {
        "name": "systemd Core Dumper",
        "id": "999",
        "gid": "999"
      },
      {
        "name": "jkstack",
        "id": "1000",
        "gid": "1000"
      },
      {
        "name": "",
        "id": "998",
        "gid": "100"
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
        "usage": 0.31
      },
      "memory": {
        "used": 220819456,
        "free": 2281967616,
        "available": 3646844928,
        "usage": 5.36,
        "swap_used": 0,
        "swap_free": 4121948160
      },
      "partitions": [
        {
          "name": "/",
          "used": 6836555776,
          "free": 17039237120,
          "usage": 28.63,
          "inode_used": 80990,
          "inode_free": 1491874,
          "inode_usage": 5.15
        },
        {
          "name": "/snap/lxd/22753",
          "used": 71172096,
          "free": 0,
          "usage": 100,
          "inode_used": 802,
          "inode_usage": 100
        },
        {
          "name": "/snap/core20/1611",
          "used": 65011712,
          "free": 0,
          "usage": 100,
          "inode_used": 11796,
          "inode_usage": 100
        },
        {
          "name": "/snap/snapd/16292",
          "used": 49283072,
          "free": 0,
          "usage": 100,
          "inode_used": 486,
          "inode_usage": 100
        },
        {
          "name": "/boot",
          "used": 110899200,
          "free": 1805324288,
          "usage": 5.79,
          "inode_used": 312,
          "inode_free": 130760,
          "inode_usage": 0.24
        },
        {
          "name": "/snap/core20/1623",
          "used": 66322432,
          "free": 0,
          "usage": 100,
          "inode_used": 11882,
          "inode_usage": 100
        },
        {
          "name": "/snap/snapd/16778",
          "used": 50331648,
          "free": 0,
          "usage": 100,
          "inode_used": 486,
          "inode_usage": 100
        }
      ],
      "interface": [
        {
          "name": "ens160",
          "bytes_sent": 6468316,
          "bytes_recv": 129677846,
          "packets_sent": 55113,
          "packets_recv": 443137
        },
        {
          "name": "lo",
          "bytes_sent": 30604,
          "bytes_recv": 30604,
          "packets_sent": 348,
          "packets_recv": 348
        }
      ]
    },
    "process": [
      {
        "id": 1,
        "pid": 0,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 13164544,
        "vms": 106307584,
        "swap": 0,
        "memory_usage": 0.32,
        "cmd": [
          "/sbin/init",
          "maybe-ubiquity"
        ],
        "conns": 19
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
        "cpu_usage": 0.03,
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
        "id": 77,
        "pid": 2,
        "user": "root",
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
        "id": 80,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 81,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 82,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 83,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 84,
        "pid": 2,
        "user": "root",
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
        "id": 88,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 89,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 91,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 92,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 93,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 94,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 95,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 96,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 97,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 98,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 99,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 100,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 101,
        "pid": 2,
        "user": "root",
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
        "id": 104,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 105,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 106,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 107,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 108,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 109,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 110,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 111,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 112,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 113,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 114,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 115,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 116,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 117,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 118,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 119,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 120,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 121,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 122,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 123,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 124,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 125,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 126,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 127,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 128,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 130,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 131,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 141,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 144,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 157,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 202,
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
        "id": 203,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 204,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 205,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 206,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 207,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 208,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 209,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 211,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 212,
        "pid": 2,
        "user": "root",
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
        "id": 221,
        "pid": 2,
        "user": "root",
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
        "id": 237,
        "pid": 2,
        "user": "root",
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
        "id": 257,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 258,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 260,
        "pid": 2,
        "user": "root",
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
        "id": 266,
        "pid": 2,
        "user": "root",
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
        "id": 269,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 270,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 271,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 272,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 273,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 274,
        "pid": 2,
        "user": "root",
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
        "id": 276,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 277,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 278,
        "pid": 2,
        "user": "root",
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
        "id": 280,
        "pid": 2,
        "user": "root",
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
        "id": 282,
        "pid": 2,
        "user": "root",
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
        "id": 285,
        "pid": 2,
        "user": "root",
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
        "id": 288,
        "pid": 2,
        "user": "root",
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
        "id": 381,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 425,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0.06,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 426,
        "pid": 2,
        "user": "root",
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
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.05,
        "rss": 18898944,
        "vms": 53379072,
        "swap": 0,
        "memory_usage": 0.46,
        "cmd": [
          "/lib/systemd/systemd-journald"
        ],
        "conns": 5
      },
      {
        "id": 523,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 532,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 7114752,
        "vms": 24113152,
        "swap": 0,
        "memory_usage": 0.17,
        "cmd": [
          "/lib/systemd/systemd-udevd"
        ],
        "conns": 3
      },
      {
        "id": 700,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 701,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 702,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 703,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 704,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.07,
        "rss": 18661376,
        "vms": 354111488,
        "swap": 0,
        "memory_usage": 0.45,
        "cmd": [
          "/sbin/multipathd",
          "-d",
          "-s"
        ],
        "conns": 2
      },
      {
        "id": 715,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 716,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 717,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 726,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 727,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 746,
        "pid": 1,
        "user": "systemd-timesync",
        "cpu_usage": 0,
        "rss": 6135808,
        "vms": 93052928,
        "swap": 0,
        "memory_usage": 0.15,
        "cmd": [
          "/lib/systemd/systemd-timesyncd"
        ],
        "conns": 2
      },
      {
        "id": 750,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 10760192,
        "vms": 48680960,
        "swap": 0,
        "memory_usage": 0.26,
        "cmd": [
          "/usr/bin/VGAuthService"
        ],
        "conns": 3
      },
      {
        "id": 751,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.15,
        "rss": 8179712,
        "vms": 318197760,
        "swap": 0,
        "memory_usage": 0.2,
        "cmd": [
          "/usr/bin/vmtoolsd"
        ],
        "conns": 1
      },
      {
        "id": 837,
        "pid": 1,
        "user": "systemd-network",
        "cpu_usage": 0,
        "rss": 7827456,
        "vms": 19517440,
        "swap": 0,
        "memory_usage": 0.19,
        "cmd": [
          "/lib/systemd/systemd-networkd"
        ],
        "conns": 2
      },
      {
        "id": 847,
        "pid": 1,
        "user": "systemd-resolve",
        "cpu_usage": 0,
        "rss": 12509184,
        "vms": 25124864,
        "swap": 0,
        "memory_usage": 0.3,
        "cmd": [
          "/lib/systemd/systemd-resolved"
        ],
        "listen": [
          53
        ],
        "conns": 4
      },
      {
        "id": 863,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 9326592,
        "vms": 245018624,
        "swap": 0,
        "memory_usage": 0.23,
        "cmd": [
          "/usr/lib/accountsservice/accounts-daemon"
        ],
        "conns": 1
      },
      {
        "id": 866,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3112960,
        "vms": 6975488,
        "swap": 0,
        "memory_usage": 0.08,
        "cmd": [
          "/usr/sbin/cron",
          "-f"
        ],
        "conns": 2
      },
      {
        "id": 867,
        "pid": 1,
        "user": "messagebus",
        "cpu_usage": 0,
        "rss": 4927488,
        "vms": 7843840,
        "swap": 0,
        "memory_usage": 0.12,
        "cmd": [
          "/usr/bin/dbus-daemon",
          "--system",
          "--address=systemd:",
          "--nofork",
          "--nopidfile",
          "--systemd-activation",
          "--syslog-only"
        ],
        "conns": 3
      },
      {
        "id": 874,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 3710976,
        "vms": 83922944,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "/usr/sbin/irqbalance",
          "--foreground"
        ],
        "conns": 2
      },
      {
        "id": 875,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 18956288,
        "vms": 30117888,
        "swap": 0,
        "memory_usage": 0.46,
        "cmd": [
          "/usr/bin/python3",
          "/usr/bin/networkd-dispatcher",
          "--run-startup-triggers"
        ],
        "conns": 1
      },
      {
        "id": 876,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 9228288,
        "vms": 242094080,
        "swap": 0,
        "memory_usage": 0.22,
        "cmd": [
          "/usr/lib/policykit-1/polkitd",
          "--no-debug"
        ],
        "conns": 1
      },
      {
        "id": 878,
        "pid": 1,
        "user": "syslog",
        "cpu_usage": 0.01,
        "rss": 5545984,
        "vms": 229728256,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          "/usr/sbin/rsyslogd",
          "-n",
          "-iNONE"
        ],
        "conns": 2
      },
      {
        "id": 882,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 8114176,
        "vms": 17924096,
        "swap": 0,
        "memory_usage": 0.2,
        "cmd": [
          "/lib/systemd/systemd-logind"
        ],
        "conns": 2
      },
      {
        "id": 884,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 14032896,
        "vms": 405049344,
        "swap": 0,
        "memory_usage": 0.34,
        "cmd": [
          "/usr/lib/udisks2/udisksd"
        ],
        "conns": 2
      },
      {
        "id": 885,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2355200,
        "vms": 3883008,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "/usr/sbin/atd",
          "-f"
        ],
        "conns": 0
      },
      {
        "id": 915,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1888256,
        "vms": 5967872,
        "swap": 0,
        "memory_usage": 0.05,
        "cmd": [
          "/sbin/agetty",
          "-o",
          "-p -- \\u",
          "--noclear",
          "tty1",
          "linux"
        ],
        "conns": 0
      },
      {
        "id": 935,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 21360640,
        "vms": 110514176,
        "swap": 0,
        "memory_usage": 0.52,
        "cmd": [
          "/usr/bin/python3",
          "/usr/share/unattended-upgrades/unattended-upgrade-shutdown",
          "--wait-for-signal"
        ],
        "conns": 1
      },
      {
        "id": 950,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 13668352,
        "vms": 250966016,
        "swap": 0,
        "memory_usage": 0.33,
        "cmd": [
          "/usr/sbin/ModemManager"
        ],
        "conns": 2
      },
      {
        "id": 18222,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 32035,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 32078,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 48742400,
        "vms": 971329536,
        "swap": 0,
        "memory_usage": 1.18,
        "cmd": [
          "/usr/lib/snapd/snapd"
        ],
        "conns": 3
      },
      {
        "id": 41923,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 7499776,
        "vms": 12464128,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "sshd: /usr/sbin/sshd -D [listener] 0 of 10-100 startups"
        ],
        "listen": [
          22,
          22
        ],
        "conns": 3
      },
      {
        "id": 77767,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.04,
        "rss": 30466048,
        "vms": 1269706752,
        "swap": 0,
        "memory_usage": 0.74,
        "cmd": [
          "/opt/metrics-agent/metrics",
          "-conf",
          "/opt/metrics-agent/agent.conf"
        ],
        "conns": 2
      },
      {
        "id": 100339,
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
        "id": 120269,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0.05,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 122505,
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
        "id": 123003,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 123243,
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
        "id": 123376,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0.01,
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
        "fd": 13,
        "pid": 847,
        "type": "tcp4",
        "local": "127.0.0.53:53",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 3,
        "pid": 41923,
        "type": "tcp4",
        "local": "0.0.0.0:22",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 3,
        "pid": 77767,
        "type": "tcp4",
        "local": "192.168.2.58:34350",
        "remote": "192.168.2.52:13081",
        "status": "ESTABLISHED"
      },
      {
        "fd": 4,
        "pid": 41923,
        "type": "tcp6",
        "local": ":::22",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 12,
        "pid": 847,
        "type": "udp4",
        "local": "127.0.0.53:53",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 40,
        "pid": 1,
        "type": "unix",
        "local": "/run/udev/control",
        "status": "NONE"
      },
      {
        "fd": 43,
        "pid": 1,
        "type": "file",
        "local": "/var/snap/lxd/common/lxd/unix.socket",
        "status": "NONE"
      },
      {
        "fd": 48,
        "pid": 1,
        "type": "file",
        "local": "@/org/kernel/linux/storage/multipathd",
        "status": "NONE"
      },
      {
        "fd": 31,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/notify",
        "status": "NONE"
      },
      {
        "fd": 19,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/private",
        "status": "NONE"
      },
      {
        "fd": 20,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/userdb/io.systemd.DynamicUser",
        "status": "NONE"
      },
      {
        "fd": 50,
        "pid": 1,
        "type": "file",
        "local": "/run/lvm/lvmpolld.socket",
        "status": "NONE"
      },
      {
        "fd": 79,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/syslog",
        "status": "NONE"
      },
      {
        "fd": 81,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/fsck.progress",
        "status": "NONE"
      },
      {
        "fd": 55,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/dev-log",
        "status": "NONE"
      },
      {
        "fd": 41,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/stdout",
        "status": "NONE"
      },
      {
        "fd": 42,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/socket",
        "status": "NONE"
      },
      {
        "fd": 10,
        "pid": 499,
        "type": "file",
        "local": "/run/systemd/journal/io.systemd.journal",
        "status": "NONE"
      },
      {
        "fd": 7,
        "pid": 750,
        "type": "file",
        "local": "/var/run/vmware/guestServicePipe",
        "status": "NONE"
      },
      {
        "fd": 35,
        "pid": 1,
        "type": "file",
        "local": "/run/snapd.socket",
        "status": "NONE"
      },
      {
        "fd": 44,
        "pid": 1,
        "type": "file",
        "local": "/run/dbus/system_bus_socket",
        "status": "NONE"
      },
      {
        "fd": 36,
        "pid": 1,
        "type": "file",
        "local": "/run/snapd-snap.socket",
        "status": "NONE"
      },
      {
        "fd": 76,
        "pid": 1,
        "type": "file",
        "local": "/run/uuidd/request",
        "status": "NONE"
      },
      {
        "fd": 5,
        "pid": 874,
        "type": "file",
        "local": "/run/irqbalance//irqbalance874.sock",
        "status": "NONE"
      },
      {
        "fd": 83,
        "pid": 1,
        "type": "file",
        "local": "@ISCSIADM_ABSTRACT_NAMESPACE",
        "status": "NONE"
      },
      {
        "fd": 5,
        "pid": 863,
        "type": "file",
        "status": "NONE"
      },
      {
        "fd": 3,
        "pid": 750,
        "type": "file",
        "status": "NONE"
      }
    ]
  }
}
```