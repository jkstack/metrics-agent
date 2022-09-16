static:

```json
{
  "code": 200,
  "payload": {
    "time": 1663317818,
    "host": {
      "name": "jkstack",
      "uptime": 261240
    },
    "os": {
      "name": "linux",
      "platform_name": "ubuntu",
      "platform_version": "22.04",
      "install": 1663035192,
      "startup": 1663056579
    },
    "kernel": {
      "version": "5.15.0-43-generic",
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
      "physical": 4116434944,
      "swap": 4115656704
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
        "mount": "/snap/core20/1587",
        "fstype": "squashfs",
        "opts": [
          "ro",
          "nodev",
          "relatime"
        ],
        "total": 65011712,
        "inodes": 11793
      },
      {
        "mount": "/snap/lxd/22923",
        "fstype": "squashfs",
        "opts": [
          "ro",
          "nodev",
          "relatime"
        ],
        "total": 83886080,
        "inodes": 816
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
        "mount": "/snap/lxd/23541",
        "fstype": "squashfs",
        "opts": [
          "ro",
          "nodev",
          "relatime"
        ],
        "total": 108003328,
        "inodes": 816
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
        "mac": "00:50:56:a3:35:ec",
        "addrs": [
          "192.168.2.59/24",
          "fe80::250:56ff:fea3:35ec/64"
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
        "gid": "65534"
      },
      {
        "name": "systemd Network Management",
        "id": "101",
        "gid": "102"
      },
      {
        "name": "systemd Resolver",
        "id": "102",
        "gid": "103"
      },
      {
        "name": "",
        "id": "103",
        "gid": "104"
      },
      {
        "name": "systemd Time Synchronization",
        "id": "104",
        "gid": "105"
      },
      {
        "name": "",
        "id": "105",
        "gid": "1"
      },
      {
        "name": "",
        "id": "106",
        "gid": "65534"
      },
      {
        "name": "",
        "id": "107",
        "gid": "113"
      },
      {
        "name": "",
        "id": "108",
        "gid": "114"
      },
      {
        "name": "",
        "id": "109",
        "gid": "115"
      },
      {
        "name": "TPM software stack",
        "id": "110",
        "gid": "116"
      },
      {
        "name": "",
        "id": "111",
        "gid": "117"
      },
      {
        "name": "usbmux daemon",
        "id": "112",
        "gid": "46"
      },
      {
        "name": "jkstack",
        "id": "1000",
        "gid": "1000"
      },
      {
        "name": "",
        "id": "999",
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
        "usage": 0.57
      },
      "memory": {
        "used": 316506112,
        "free": 1660735488,
        "available": 3511365632,
        "usage": 7.69,
        "swap_used": 0,
        "swap_free": 4115656704
      },
      "partitions": [
        {
          "name": "/",
          "used": 7783264256,
          "free": 16092528640,
          "usage": 32.6,
          "inode_used": 116753,
          "inode_free": 1456111,
          "inode_usage": 7.42
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
          "name": "/snap/core20/1587",
          "used": 65011712,
          "free": 0,
          "usage": 100,
          "inode_used": 11793,
          "inode_usage": 100
        },
        {
          "name": "/snap/lxd/22923",
          "used": 83886080,
          "free": 0,
          "usage": 100,
          "inode_used": 816,
          "inode_usage": 100
        },
        {
          "name": "/boot",
          "used": 256839680,
          "free": 1659383808,
          "usage": 13.4,
          "inode_used": 320,
          "inode_free": 130752,
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
          "name": "/snap/lxd/23541",
          "used": 108003328,
          "free": 0,
          "usage": 100,
          "inode_used": 816,
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
          "name": "lo",
          "bytes_sent": 33657,
          "bytes_recv": 33657,
          "packets_sent": 356,
          "packets_recv": 356
        },
        {
          "name": "ens160",
          "bytes_sent": 46048069,
          "bytes_recv": 330258306,
          "packets_sent": 99440,
          "packets_recv": 528322
        }
      ]
    },
    "process": [
      {
        "id": 1,
        "pid": 0,
        "user": "root",
        "cpu_usage": 0,
        "rss": 13426688,
        "vms": 171724800,
        "swap": 0,
        "memory_usage": 0.33,
        "cmd": [
          "/sbin/init"
        ],
        "conns": 21
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
        "cpu_usage": 0.02,
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
        "cpu_usage": 0.01,
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
        "id": 86,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 87,
        "pid": 2,
        "user": "root",
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
        "id": 129,
        "pid": 2,
        "user": "root",
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
        "id": 132,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 134,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 136,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 137,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 146,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 149,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 150,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 155,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 199,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 201,
        "pid": 2,
        "user": "root",
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
        "cpu_usage": 0,
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
        "id": 210,
        "pid": 2,
        "user": "root",
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
        "id": 213,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 214,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 215,
        "pid": 2,
        "user": "root",
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
        "id": 218,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 219,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 220,
        "pid": 2,
        "user": "root",
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
        "id": 232,
        "pid": 2,
        "user": "root",
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
        "id": 234,
        "pid": 2,
        "user": "root",
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
        "id": 238,
        "pid": 2,
        "user": "root",
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
        "id": 256,
        "pid": 2,
        "user": "root",
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
        "id": 259,
        "pid": 2,
        "user": "root",
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
        "id": 261,
        "pid": 2,
        "user": "root",
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
        "cpu_usage": 0.02,
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
        "id": 434,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 435,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 509,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 18767872,
        "vms": 48939008,
        "swap": 0,
        "memory_usage": 0.46,
        "cmd": [
          "/lib/systemd/systemd-journald"
        ],
        "conns": 5
      },
      {
        "id": 543,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 548,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 549,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 551,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 552,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 553,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 27754496,
        "vms": 363405312,
        "swap": 0,
        "memory_usage": 0.67,
        "cmd": [
          "/sbin/multipathd",
          "-d",
          "-s"
        ],
        "conns": 2
      },
      {
        "id": 557,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 7819264,
        "vms": 27783168,
        "swap": 0,
        "memory_usage": 0.19,
        "cmd": [
          "/lib/systemd/systemd-udevd"
        ],
        "conns": 3
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
        "id": 731,
        "pid": 1,
        "user": "systemd-timesync",
        "cpu_usage": 0,
        "rss": 6627328,
        "vms": 91496448,
        "swap": 0,
        "memory_usage": 0.16,
        "cmd": [
          "/lib/systemd/systemd-timesyncd"
        ],
        "conns": 2
      },
      {
        "id": 838,
        "pid": 1,
        "user": "systemd-network",
        "cpu_usage": 0,
        "rss": 8282112,
        "vms": 16502784,
        "swap": 0,
        "memory_usage": 0.2,
        "cmd": [
          "/lib/systemd/systemd-networkd"
        ],
        "conns": 2
      },
      {
        "id": 840,
        "pid": 1,
        "user": "systemd-resolve",
        "cpu_usage": 0,
        "rss": 12865536,
        "vms": 26001408,
        "swap": 0,
        "memory_usage": 0.31,
        "cmd": [
          "/lib/systemd/systemd-resolved"
        ],
        "listen": [
          53
        ],
        "conns": 5
      },
      {
        "id": 851,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2940928,
        "vms": 7057408,
        "swap": 0,
        "memory_usage": 0.07,
        "cmd": [
          "/usr/sbin/cron",
          "-f",
          "-P"
        ],
        "conns": 2
      },
      {
        "id": 852,
        "pid": 1,
        "user": "messagebus",
        "cpu_usage": 0,
        "rss": 5128192,
        "vms": 9080832,
        "swap": 0,
        "memory_usage": 0.12,
        "cmd": [
          "@dbus-daemon",
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
        "id": 859,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 3899392,
        "vms": 84815872,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "/usr/sbin/irqbalance",
          "--foreground"
        ],
        "conns": 2
      },
      {
        "id": 860,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 19447808,
        "vms": 33431552,
        "swap": 0,
        "memory_usage": 0.47,
        "cmd": [
          "/usr/bin/python3",
          "/usr/bin/networkd-dispatcher",
          "--run-startup-triggers"
        ],
        "conns": 1
      },
      {
        "id": 861,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 6889472,
        "vms": 240111616,
        "swap": 0,
        "memory_usage": 0.17,
        "cmd": [
          "/usr/libexec/polkitd",
          "--no-debug"
        ],
        "conns": 1
      },
      {
        "id": 863,
        "pid": 1,
        "user": "syslog",
        "cpu_usage": 0,
        "rss": 6066176,
        "vms": 227737600,
        "swap": 0,
        "memory_usage": 0.15,
        "cmd": [
          "/usr/sbin/rsyslogd",
          "-n",
          "-iNONE"
        ],
        "conns": 2
      },
      {
        "id": 866,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 8171520,
        "vms": 24313856,
        "swap": 0,
        "memory_usage": 0.2,
        "cmd": [
          "/lib/systemd/systemd-logind"
        ],
        "conns": 2
      },
      {
        "id": 869,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 13082624,
        "vms": 401989632,
        "swap": 0,
        "memory_usage": 0.32,
        "cmd": [
          "/usr/libexec/udisks2/udisksd"
        ],
        "conns": 2
      },
      {
        "id": 883,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1126400,
        "vms": 6320128,
        "swap": 0,
        "memory_usage": 0.03,
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
        "id": 924,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 22077440,
        "vms": 112349184,
        "swap": 0,
        "memory_usage": 0.54,
        "cmd": [
          "/usr/bin/python3",
          "/usr/share/unattended-upgrades/unattended-upgrade-shutdown",
          "--wait-for-signal"
        ],
        "conns": 1
      },
      {
        "id": 930,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 12062720,
        "vms": 249110528,
        "swap": 0,
        "memory_usage": 0.29,
        "cmd": [
          "/usr/sbin/ModemManager"
        ],
        "conns": 2
      },
      {
        "id": 2533,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 48029696,
        "vms": 971517952,
        "swap": 0,
        "memory_usage": 1.17,
        "cmd": [
          "/usr/lib/snapd/snapd"
        ],
        "conns": 3
      },
      {
        "id": 16417,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 9572352,
        "vms": 15790080,
        "swap": 0,
        "memory_usage": 0.23,
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
        "id": 18065,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 9670656,
        "vms": 17498112,
        "swap": 0,
        "memory_usage": 0.23,
        "cmd": [
          "/lib/systemd/systemd",
          "--user"
        ],
        "conns": 12
      },
      {
        "id": 18066,
        "pid": 18065,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5201920,
        "vms": 174706688,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          "(sd-pam)"
        ],
        "conns": 2
      },
      {
        "id": 18397,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.06,
        "rss": 32010240,
        "vms": 1270042624,
        "swap": 0,
        "memory_usage": 0.78,
        "cmd": [
          "/opt/metrics-agent/metrics",
          "-conf",
          "/opt/metrics-agent/agent.conf"
        ],
        "conns": 2
      },
      {
        "id": 18561,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 21184512,
        "vms": 302641152,
        "swap": 0,
        "memory_usage": 0.51,
        "cmd": [
          "/usr/libexec/packagekitd"
        ],
        "conns": 2
      },
      {
        "id": 18565,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 12267520,
        "vms": 52359168,
        "swap": 0,
        "memory_usage": 0.3,
        "cmd": [
          "/usr/bin/VGAuthService"
        ],
        "conns": 3
      },
      {
        "id": 18567,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.15,
        "rss": 9043968,
        "vms": 246919168,
        "swap": 0,
        "memory_usage": 0.22,
        "cmd": [
          "/usr/bin/vmtoolsd"
        ],
        "conns": 1
      },
      {
        "id": 19547,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0.04,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 19852,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0.04,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 19853,
        "pid": 16417,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 11624448,
        "vms": 17555456,
        "swap": 0,
        "memory_usage": 0.28,
        "cmd": [
          "sshd: root@pts/0"
        ],
        "conns": 3
      },
      {
        "id": 19910,
        "pid": 19853,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5451776,
        "vms": 8994816,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          "-bash"
        ],
        "conns": 0
      },
      {
        "id": 19926,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 19938,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0.08,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 19945,
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
        "id": 19946,
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
        "fd": 3,
        "pid": 16417,
        "type": "tcp4",
        "local": "0.0.0.0:22",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 14,
        "pid": 840,
        "type": "tcp4",
        "local": "127.0.0.53:53",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 3,
        "pid": 18397,
        "type": "tcp4",
        "local": "192.168.2.59:40372",
        "remote": "192.168.2.52:13081",
        "status": "ESTABLISHED"
      },
      {
        "fd": 4,
        "pid": 19853,
        "type": "tcp4",
        "local": "192.168.2.59:22",
        "remote": "10.202.0.9:59049",
        "status": "ESTABLISHED"
      },
      {
        "fd": 4,
        "pid": 16417,
        "type": "tcp6",
        "local": ":::22",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 13,
        "pid": 840,
        "type": "udp4",
        "local": "127.0.0.53:53",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 128,
        "pid": 1,
        "type": "file",
        "local": "/var/snap/lxd/common/lxd-user/unix.socket",
        "status": "NONE"
      },
      {
        "fd": 10,
        "pid": 509,
        "type": "file",
        "local": "/run/systemd/journal/io.systemd.journal",
        "status": "NONE"
      },
      {
        "fd": 16,
        "pid": 18065,
        "type": "file",
        "local": "/run/user/0/systemd/notify",
        "status": "NONE"
      },
      {
        "fd": 19,
        "pid": 18065,
        "type": "file",
        "local": "/run/user/0/systemd/private",
        "status": "NONE"
      },
      {
        "fd": 12,
        "pid": 18065,
        "type": "file",
        "local": "/run/user/0/bus",
        "status": "NONE"
      },
      {
        "fd": 129,
        "pid": 1,
        "type": "file",
        "local": "@/org/kernel/linux/storage/multipathd",
        "status": "NONE"
      },
      {
        "fd": 26,
        "pid": 18065,
        "type": "file",
        "local": "/run/user/0/gnupg/S.dirmngr",
        "status": "NONE"
      },
      {
        "fd": 27,
        "pid": 18065,
        "type": "file",
        "local": "/run/user/0/gnupg/S.gpg-agent.browser",
        "status": "NONE"
      },
      {
        "fd": 28,
        "pid": 18065,
        "type": "file",
        "local": "/run/user/0/gnupg/S.gpg-agent.extra",
        "status": "NONE"
      },
      {
        "fd": 29,
        "pid": 18065,
        "type": "file",
        "local": "/run/user/0/gnupg/S.gpg-agent.ssh",
        "status": "NONE"
      },
      {
        "fd": 30,
        "pid": 18065,
        "type": "file",
        "local": "/run/user/0/gnupg/S.gpg-agent",
        "status": "NONE"
      },
      {
        "fd": 31,
        "pid": 18065,
        "type": "file",
        "local": "/run/user/0/pk-debconf-socket",
        "status": "NONE"
      },
      {
        "fd": 7,
        "pid": 18565,
        "type": "file",
        "local": "/var/run/vmware/guestServicePipe",
        "status": "NONE"
      },
      {
        "fd": 32,
        "pid": 18065,
        "type": "file",
        "local": "/run/user/0/snapd-session-agent.socket",
        "status": "NONE"
      },
      {
        "fd": 15,
        "pid": 840,
        "type": "file",
        "local": "/run/systemd/resolve/io.systemd.Resolve",
        "status": "NONE"
      },
      {
        "fd": 124,
        "pid": 1,
        "type": "file",
        "local": "/run/dbus/system_bus_socket",
        "status": "NONE"
      },
      {
        "fd": 134,
        "pid": 1,
        "type": "file",
        "local": "/run/snapd.socket",
        "status": "NONE"
      },
      {
        "fd": 135,
        "pid": 1,
        "type": "file",
        "local": "/run/snapd-snap.socket",
        "status": "NONE"
      },
      {
        "fd": 127,
        "pid": 1,
        "type": "file",
        "local": "/run/uuidd/request",
        "status": "NONE"
      },
      {
        "fd": 42,
        "pid": 1,
        "type": "file",
        "local": "@ISCSIADM_ABSTRACT_NAMESPACE",
        "status": "NONE"
      },
      {
        "fd": 5,
        "pid": 859,
        "type": "file",
        "local": "/run/irqbalance/irqbalance859.sock",
        "status": "NONE"
      },
      {
        "fd": 108,
        "pid": 1,
        "type": "file",
        "local": "/var/snap/lxd/common/lxd/unix.socket",
        "status": "NONE"
      },
      {
        "fd": 37,
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
        "fd": 21,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/io.system.ManagedOOM",
        "status": "NONE"
      },
      {
        "fd": 66,
        "pid": 1,
        "type": "file",
        "local": "/run/lvm/lvmpolld.socket",
        "status": "NONE"
      },
      {
        "fd": 53,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/syslog",
        "status": "NONE"
      },
      {
        "fd": 73,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/fsck.progress",
        "status": "NONE"
      },
      {
        "fd": 133,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/dev-log",
        "status": "NONE"
      },
      {
        "fd": 125,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/socket",
        "status": "NONE"
      },
      {
        "fd": 126,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/stdout",
        "status": "NONE"
      },
      {
        "fd": 54,
        "pid": 1,
        "type": "unix",
        "local": "/run/udev/control",
        "status": "NONE"
      },
      {
        "fd": 1,
        "pid": 731,
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