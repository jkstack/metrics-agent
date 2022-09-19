static:

```json
{
  "code": 200,
  "payload": {
    "time": 1663289101,
    "host": {
      "name": "suse11",
      "uptime": 172146
    },
    "os": {
      "name": "linux",
      "platform_name": "suse",
      "platform_version": "11.4",
      "install": 1663108724,
      "startup": 1663116955
    },
    "kernel": {
      "version": "3.0.101-63-default",
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
      "physical": 4020203520,
      "swap": 2153771008
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
        "fstype": "ext3",
        "opts": [
          "rw",
          "relatime"
        ],
        "total": 50723602432,
        "inodes": 3145728
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
          "127.0.0.2/8",
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
        "mac": "00:50:56:a3:c9:ad",
        "addrs": [
          "192.168.2.62/24",
          "fe80::250:56ff:fea3:c9ad/64"
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
        "name": "lp",
        "id": "4",
        "gid": "7"
      },
      {
        "name": "mail",
        "id": "8",
        "gid": "12"
      },
      {
        "name": "games",
        "id": "12",
        "gid": "100"
      },
      {
        "name": "wwwrun",
        "id": "30",
        "gid": "8"
      },
      {
        "name": "ftp",
        "id": "40",
        "gid": "49"
      },
      {
        "name": "nobody",
        "id": "65534",
        "gid": "65533"
      },
      {
        "name": "messagebus",
        "id": "100",
        "gid": "101"
      },
      {
        "name": "haldaemon",
        "id": "101",
        "gid": "102"
      },
      {
        "name": "sshd",
        "id": "102",
        "gid": "103"
      },
      {
        "name": "man",
        "id": "13",
        "gid": "62"
      },
      {
        "name": "news",
        "id": "9",
        "gid": "13"
      },
      {
        "name": "uucp",
        "id": "10",
        "gid": "14"
      },
      {
        "name": "uuidd",
        "id": "103",
        "gid": "105"
      },
      {
        "name": "puppet",
        "id": "104",
        "gid": "107"
      },
      {
        "name": "at",
        "id": "25",
        "gid": "25"
      },
      {
        "name": "postfix",
        "id": "51",
        "gid": "51"
      },
      {
        "name": "polkituser",
        "id": "105",
        "gid": "108"
      },
      {
        "name": "ntp",
        "id": "74",
        "gid": "109"
      },
      {
        "name": "pulse",
        "id": "106",
        "gid": "110"
      },
      {
        "name": "suse-ncc",
        "id": "107",
        "gid": "112"
      },
      {
        "name": "gdm",
        "id": "108",
        "gid": "113"
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
        "usage": 0.47
      },
      "memory": {
        "used": 267706368,
        "free": 3380326400,
        "available": 3558180864,
        "usage": 6.66,
        "swap_used": 0,
        "swap_free": 2153771008
      },
      "partitions": [
        {
          "name": "/",
          "used": 3521478656,
          "free": 46119948288,
          "usage": 7.09,
          "inode_used": 150972,
          "inode_free": 2994756,
          "inode_usage": 4.8
        }
      ],
      "interface": [
        {
          "name": "lo",
          "bytes_sent": 820,
          "bytes_recv": 820,
          "packets_sent": 14,
          "packets_recv": 14
        },
        {
          "name": "eth0",
          "bytes_sent": 3506992,
          "bytes_recv": 92176603,
          "packets_sent": 39251,
          "packets_recv": 299181
        }
      ]
    },
    "process": [
      {
        "id": 1,
        "pid": 0,
        "user": "root",
        "cpu_usage": 0,
        "rss": 901120,
        "vms": 10813440,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "init [5]"
        ],
        "conns": 0
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
        "id": 90,
        "pid": 2,
        "user": "root",
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
        "id": 147,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 160,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 161,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 184,
        "pid": 2,
        "user": "root",
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
        "id": 601,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1449984,
        "vms": 17981440,
        "swap": 0,
        "memory_usage": 0.04,
        "cmd": [
          "/sbin/udevd",
          "--daemon"
        ],
        "conns": 2
      },
      {
        "id": 875,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 1723,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 663552,
        "vms": 4120576,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/acpid"
        ],
        "conns": 2
      },
      {
        "id": 1738,
        "pid": 1,
        "user": "messagebus",
        "cpu_usage": 0,
        "rss": 1273856,
        "vms": 22151168,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "/bin/dbus-daemon",
          "--system"
        ],
        "conns": 2
      },
      {
        "id": 1756,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 954368,
        "vms": 19791872,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/syslog-ng"
        ],
        "conns": 1
      },
      {
        "id": 1759,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 675840,
        "vms": 4395008,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/klogd",
          "-c",
          "1",
          "-x"
        ],
        "conns": 1
      },
      {
        "id": 1847,
        "pid": 1,
        "user": "haldaemon",
        "cpu_usage": 0,
        "rss": 4534272,
        "vms": 39993344,
        "swap": 0,
        "memory_usage": 0.11,
        "cmd": [
          "/usr/sbin/hald",
          "--daemon=yes"
        ],
        "conns": 4
      },
      {
        "id": 1850,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2842624,
        "vms": 111697920,
        "swap": 0,
        "memory_usage": 0.07,
        "cmd": [
          "/usr/sbin/console-kit-daemon"
        ],
        "conns": 1
      },
      {
        "id": 1851,
        "pid": 1847,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1433600,
        "vms": 22839296,
        "swap": 0,
        "memory_usage": 0.04,
        "cmd": [
          "hald-runner"
        ],
        "conns": 1
      },
      {
        "id": 1946,
        "pid": 1851,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2191360,
        "vms": 39976960,
        "swap": 0,
        "memory_usage": 0.05,
        "cmd": [
          "hald-addon-input: Listening on /dev/input/event0 /dev/input/event2"
        ],
        "conns": 1
      },
      {
        "id": 1968,
        "pid": 1851,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2199552,
        "vms": 39976960,
        "swap": 0,
        "memory_usage": 0.05,
        "cmd": [
          "hald-addon-storage: polling /dev/sr0 (every 16 sec)"
        ],
        "conns": 1
      },
      {
        "id": 1971,
        "pid": 1851,
        "user": "haldaemon",
        "cpu_usage": 0,
        "rss": 2052096,
        "vms": 41775104,
        "swap": 0,
        "memory_usage": 0.05,
        "cmd": [
          "hald-addon-acpi: listening on acpid socket /var/run/acpid.socket"
        ],
        "conns": 1
      },
      {
        "id": 2827,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 638976,
        "vms": 30371840,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/auditd",
          "-s",
          "disable"
        ],
        "conns": 2
      },
      {
        "id": 2829,
        "pid": 2827,
        "user": "root",
        "cpu_usage": 0,
        "rss": 794624,
        "vms": 14794752,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/audispd"
        ],
        "conns": 3
      },
      {
        "id": 2830,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 2851,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 4661248,
        "vms": 8585216,
        "swap": 0,
        "memory_usage": 0.12,
        "cmd": [
          "/sbin/haveged",
          "-w",
          "1024",
          "-v",
          "1"
        ],
        "conns": 1
      },
      {
        "id": 2865,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 655360,
        "vms": 19378176,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/rpcbind"
        ],
        "listen": [
          111,
          111
        ],
        "conns": 7
      },
      {
        "id": 2924,
        "pid": 601,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1458176,
        "vms": 17977344,
        "swap": 0,
        "memory_usage": 0.04,
        "cmd": [
          "/sbin/udevd",
          "--daemon"
        ],
        "conns": 1
      },
      {
        "id": 2925,
        "pid": 601,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1454080,
        "vms": 17977344,
        "swap": 0,
        "memory_usage": 0.04,
        "cmd": [
          "/sbin/udevd",
          "--daemon"
        ],
        "conns": 1
      },
      {
        "id": 2991,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.02,
        "rss": 610304,
        "vms": 8867840,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/usr/sbin/irqbalance"
        ],
        "conns": 0
      },
      {
        "id": 3007,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 843776,
        "vms": 27795456,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/usr/sbin/mcelog",
          "--daemon",
          "--config-file",
          "/etc/mcelog/mcelog.conf"
        ],
        "conns": 2
      },
      {
        "id": 3311,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1286144,
        "vms": 55054336,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "/usr/sbin/sshd",
          "-o",
          "PidFile=/var/run/sshd.init.pid"
        ],
        "listen": [
          22,
          22
        ],
        "conns": 2
      },
      {
        "id": 3325,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2985984,
        "vms": 73547776,
        "swap": 0,
        "memory_usage": 0.07,
        "cmd": [
          "/usr/sbin/cupsd"
        ],
        "listen": [
          631,
          631
        ],
        "conns": 4
      },
      {
        "id": 3340,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1032192,
        "vms": 136904704,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "/usr/sbin/nscd"
        ],
        "conns": 1
      },
      {
        "id": 3430,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2555904,
        "vms": 41226240,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "/usr/lib/postfix/master"
        ],
        "listen": [
          25,
          25
        ],
        "conns": 30
      },
      {
        "id": 3461,
        "pid": 3430,
        "user": "postfix",
        "cpu_usage": 0,
        "rss": 2392064,
        "vms": 43175936,
        "swap": 0,
        "memory_usage": 0.06,
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
        "id": 3497,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3100672,
        "vms": 59195392,
        "swap": 0,
        "memory_usage": 0.08,
        "cmd": [
          "/usr/sbin/gdm"
        ],
        "conns": 1
      },
      {
        "id": 3523,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 737280,
        "vms": 15167488,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/usr/sbin/cron"
        ],
        "conns": 1
      },
      {
        "id": 3594,
        "pid": 3497,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3371008,
        "vms": 59326464,
        "swap": 0,
        "memory_usage": 0.08,
        "cmd": [
          "/usr/lib/gdm/gdm-simple-slave",
          "--display-id",
          "/org/gnome/DisplayManager/Display1"
        ],
        "conns": 3
      },
      {
        "id": 3625,
        "pid": 3594,
        "user": "root",
        "cpu_usage": 0.03,
        "rss": 29839360,
        "vms": 151875584,
        "swap": 0,
        "memory_usage": 0.74,
        "cmd": [
          "/usr/bin/X",
          ":0",
          "-br",
          "-verbose",
          "-auth",
          "/var/run/gdm/auth-for-gdm-orcGpR/database",
          "-nolisten",
          "tcp",
          "vt7"
        ],
        "conns": 3
      },
      {
        "id": 3639,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 802816,
        "vms": 4624384,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/mingetty",
          "--noclear",
          "tty1"
        ],
        "conns": 0
      },
      {
        "id": 3644,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 798720,
        "vms": 4624384,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/mingetty",
          "tty2"
        ],
        "conns": 0
      },
      {
        "id": 3645,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 794624,
        "vms": 4624384,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/mingetty",
          "tty3"
        ],
        "conns": 0
      },
      {
        "id": 3646,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 794624,
        "vms": 4624384,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/mingetty",
          "tty4"
        ],
        "conns": 0
      },
      {
        "id": 3648,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 794624,
        "vms": 4624384,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/mingetty",
          "tty5"
        ],
        "conns": 0
      },
      {
        "id": 3649,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 794624,
        "vms": 4624384,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/sbin/mingetty",
          "tty6"
        ],
        "conns": 0
      },
      {
        "id": 3780,
        "pid": 1,
        "user": "gdm",
        "cpu_usage": 0,
        "rss": 610304,
        "vms": 22757376,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "/usr/bin/dbus-launch",
          "--exit-with-session"
        ],
        "conns": 1
      },
      {
        "id": 5090,
        "pid": 3594,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3735552,
        "vms": 91332608,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "/usr/lib/gdm/gdm-session-worker"
        ],
        "conns": 2
      },
      {
        "id": 7418,
        "pid": 5090,
        "user": "root",
        "cpu_usage": 0,
        "rss": 7749632,
        "vms": 162357248,
        "swap": 0,
        "memory_usage": 0.19,
        "cmd": [
          "/usr/bin/gnome-session"
        ],
        "conns": 4
      },
      {
        "id": 7492,
        "pid": 7418,
        "user": "root",
        "cpu_usage": 0,
        "rss": 446464,
        "vms": 17154048,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "/usr/bin/gpg-agent",
          "--sh",
          "--daemon",
          "--write-env-file",
          "/root/.gnupg/agent.info",
          "/bin/bash",
          "/etc/X11/xinit/xinitrc"
        ],
        "conns": 1
      },
      {
        "id": 7498,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 610304,
        "vms": 22757376,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "dbus-launch",
          "--exit-with-session",
          "/usr/bin/gnome-session"
        ],
        "conns": 1
      },
      {
        "id": 7499,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1048576,
        "vms": 14073856,
        "swap": 0,
        "memory_usage": 0.03,
        "cmd": [
          "/bin/dbus-daemon",
          "--fork",
          "--print-pid",
          "5",
          "--print-address",
          "9",
          "--session"
        ],
        "conns": 2
      },
      {
        "id": 7502,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5283840,
        "vms": 40538112,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          "/usr/lib/GConf/2/gconfd-2"
        ],
        "conns": 2
      },
      {
        "id": 7507,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 13918208,
        "vms": 372121600,
        "swap": 0,
        "memory_usage": 0.35,
        "cmd": [
          "/usr/lib/gnome-settings-daemon/gnome-settings-daemon"
        ],
        "conns": 2
      },
      {
        "id": 7510,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3497984,
        "vms": 97595392,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "gnome-keyring-daemon",
          "--start"
        ],
        "conns": 5
      },
      {
        "id": 7514,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2428928,
        "vms": 44707840,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "/usr/lib64/gvfs/gvfsd"
        ],
        "conns": 1
      },
      {
        "id": 7518,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2449408,
        "vms": 69021696,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "/usr/lib64/gvfs//gvfs-fuse-daemon",
          "/root/.gvfs"
        ],
        "conns": 1
      },
      {
        "id": 7524,
        "pid": 7418,
        "user": "root",
        "cpu_usage": 0,
        "rss": 11427840,
        "vms": 156872704,
        "swap": 0,
        "memory_usage": 0.28,
        "cmd": [
          "/usr/bin/metacity"
        ],
        "conns": 2
      },
      {
        "id": 7529,
        "pid": 7418,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 20746240,
        "vms": 380809216,
        "swap": 0,
        "memory_usage": 0.52,
        "cmd": [
          "gnome-panel"
        ],
        "conns": 2
      },
      {
        "id": 7530,
        "pid": 7418,
        "user": "root",
        "cpu_usage": 0.08,
        "rss": 18735104,
        "vms": 342257664,
        "swap": 0,
        "memory_usage": 0.47,
        "cmd": [
          "nautilus"
        ],
        "conns": 3
      },
      {
        "id": 7532,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3706880,
        "vms": 149995520,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "/usr/lib/bonobo/bonobo-activation-server",
          "--ac-activate",
          "--ior-output-fd=18"
        ],
        "conns": 2
      },
      {
        "id": 7541,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.02,
        "rss": 20332544,
        "vms": 333864960,
        "swap": 0,
        "memory_usage": 0.51,
        "cmd": [
          "/usr/lib64/gnome-main-menu/main-menu",
          "--oaf-activate-iid=OAFIID:GNOME_MainMenu_Factory",
          "--oaf-ior-fd=18"
        ],
        "conns": 2
      },
      {
        "id": 7545,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2936832,
        "vms": 49520640,
        "swap": 0,
        "memory_usage": 0.07,
        "cmd": [
          "/usr/lib64/gvfs/gvfsd-trash",
          "--spawner",
          ":1.7",
          "/org/gtk/gvfs/exec_spaw/0"
        ],
        "conns": 5
      },
      {
        "id": 7547,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 4091904,
        "vms": 55779328,
        "swap": 0,
        "memory_usage": 0.1,
        "cmd": [
          "/usr/lib64/gvfs/gvfs-hal-volume-monitor"
        ],
        "conns": 1
      },
      {
        "id": 7551,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3686400,
        "vms": 64647168,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "/usr/lib64/gvfs/gvfs-gphoto2-volume-monitor"
        ],
        "conns": 1
      },
      {
        "id": 7555,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2486272,
        "vms": 44929024,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "/usr/lib64/gvfs/gvfsd-burn",
          "--spawner",
          ":1.7",
          "/org/gtk/gvfs/exec_spaw/1"
        ],
        "conns": 3
      },
      {
        "id": 7556,
        "pid": 7418,
        "user": "root",
        "cpu_usage": 0,
        "rss": 12693504,
        "vms": 224567296,
        "swap": 0,
        "memory_usage": 0.32,
        "cmd": [
          "gpk-update-icon"
        ],
        "conns": 2
      },
      {
        "id": 7559,
        "pid": 7418,
        "user": "root",
        "cpu_usage": 0,
        "rss": 20791296,
        "vms": 233598976,
        "swap": 0,
        "memory_usage": 0.52,
        "cmd": [
          "python",
          "/usr/lib64/python2.6/site-packages/system-config-printer/applet.py"
        ],
        "conns": 1
      },
      {
        "id": 7573,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3743744,
        "vms": 197406720,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "/usr/bin/pulseaudio",
          "--start"
        ],
        "conns": 4
      },
      {
        "id": 7574,
        "pid": 7418,
        "user": "root",
        "cpu_usage": 0,
        "rss": 12455936,
        "vms": 309284864,
        "swap": 0,
        "memory_usage": 0.31,
        "cmd": [
          "gnome-volume-control-applet"
        ],
        "conns": 2
      },
      {
        "id": 7575,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 10575872,
        "vms": 235995136,
        "swap": 0,
        "memory_usage": 0.26,
        "cmd": [
          "gnome-power-manager"
        ],
        "conns": 2
      },
      {
        "id": 7581,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5672960,
        "vms": 155947008,
        "swap": 0,
        "memory_usage": 0.14,
        "cmd": [
          "gnome-screensaver"
        ],
        "conns": 2
      },
      {
        "id": 7589,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 14417920,
        "vms": 229502976,
        "swap": 0,
        "memory_usage": 0.36,
        "cmd": [
          "/usr/bin/gnome-terminal",
          "-x",
          "/bin/sh",
          "-c",
          "cd '/root/Desktop' && exec $SHELL -l"
        ],
        "conns": 2
      },
      {
        "id": 7590,
        "pid": 7589,
        "user": "root",
        "cpu_usage": 0,
        "rss": 331776,
        "vms": 6651904,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "gnome-pty-helper"
        ],
        "conns": 1
      },
      {
        "id": 7591,
        "pid": 7589,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2985984,
        "vms": 14254080,
        "swap": 0,
        "memory_usage": 0.07,
        "cmd": [
          "/bin/bash",
          "-l"
        ],
        "conns": 0
      },
      {
        "id": 10242,
        "pid": 3311,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3678208,
        "vms": 81707008,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "sshd: root@pts/1"
        ],
        "conns": 2
      },
      {
        "id": 10246,
        "pid": 10242,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3043328,
        "vms": 14581760,
        "swap": 0,
        "memory_usage": 0.08,
        "cmd": [
          "-bash"
        ],
        "conns": 0
      },
      {
        "id": 10392,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 10899,
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
        "id": 10904,
        "pid": 3430,
        "user": "postfix",
        "cpu_usage": 0,
        "rss": 2387968,
        "vms": 43126784,
        "swap": 0,
        "memory_usage": 0.06,
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
        "id": 11314,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.34,
        "rss": 57720832,
        "vms": 803962880,
        "swap": 0,
        "memory_usage": 1.44,
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
        "fd": 8,
        "pid": 2865,
        "type": "tcp4",
        "local": "0.0.0.0:111",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 3,
        "pid": 3311,
        "type": "tcp4",
        "local": "0.0.0.0:22",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 3,
        "pid": 3325,
        "type": "tcp4",
        "local": "127.0.0.1:631",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 12,
        "pid": 3430,
        "type": "tcp4",
        "local": "127.0.0.1:25",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 5,
        "pid": 10242,
        "type": "tcp4",
        "local": "192.168.2.62:22",
        "remote": "10.202.0.9:58798",
        "status": "ESTABLISHED"
      },
      {
        "fd": 7,
        "pid": 11314,
        "type": "tcp4",
        "local": "192.168.2.62:44481",
        "remote": "192.168.2.52:13081",
        "status": "ESTABLISHED"
      },
      {
        "fd": 11,
        "pid": 2865,
        "type": "tcp6",
        "local": ":::111",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 4,
        "pid": 3311,
        "type": "tcp6",
        "local": ":::22",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 1,
        "pid": 3325,
        "type": "tcp6",
        "local": "::1:631",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 13,
        "pid": 3430,
        "type": "tcp6",
        "local": "::1:25",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 5,
        "pid": 3325,
        "type": "udp4",
        "local": "0.0.0.0:631",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 7,
        "pid": 2865,
        "type": "udp4",
        "local": "0.0.0.0:920",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 6,
        "pid": 2865,
        "type": "udp4",
        "local": "0.0.0.0:111",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 10,
        "pid": 2865,
        "type": "udp6",
        "local": ":::920",
        "remote": "::",
        "status": "NONE"
      },
      {
        "fd": 9,
        "pid": 2865,
        "type": "udp6",
        "local": ":::111",
        "remote": "::",
        "status": "NONE"
      },
      {
        "fd": 3,
        "pid": 1756,
        "type": "file",
        "local": "/dev/log",
        "status": "NONE"
      },
      {
        "fd": 5,
        "pid": 2829,
        "type": "file",
        "local": "/var/run/audispd_events",
        "status": "NONE"
      },
      {
        "fd": 3,
        "pid": 3625,
        "type": "file",
        "local": "/tmp/.X11-unix/X0",
        "status": "NONE"
      },
      {
        "fd": 5,
        "pid": 2865,
        "type": "file",
        "local": "/var/run/rpcbind.sock",
        "status": "NONE"
      },
      {
        "fd": 5,
        "pid": 7492,
        "type": "file",
        "local": "/tmp/gpg-AP0r8Y/S.gpg-agent",
        "status": "NONE"
      },
      {
        "fd": 7,
        "pid": 3007,
        "type": "file",
        "local": "/var/run/mcelog-client",
        "status": "NONE"
      },
      {
        "fd": 10,
        "pid": 7418,
        "type": "file",
        "local": "/tmp/.ICE-unix/7418",
        "status": "NONE"
      },
      {
        "fd": 14,
        "pid": 7502,
        "type": "file",
        "local": "/tmp/orbit-root/linc-1d4e-0-1a583d9779404",
        "status": "NONE"
      },
      {
        "fd": 18,
        "pid": 7418,
        "type": "file",
        "local": "/tmp/orbit-root/linc-1cfa-0-5074a91a2fcf",
        "status": "NONE"
      },
      {
        "fd": 15,
        "pid": 7507,
        "type": "file",
        "local": "/tmp/orbit-root/linc-1d53-0-7f8924db81c21",
        "status": "NONE"
      },
      {
        "fd": 10,
        "pid": 3340,
        "type": "file",
        "local": "/var/run/nscd/socket",
        "status": "NONE"
      },
      {
        "fd": 6,
        "pid": 7510,
        "type": "file",
        "local": "/tmp/keyring-xnU0KR/socket",
        "status": "NONE"
      },
      {
        "fd": 13,
        "pid": 7510,
        "type": "file",
        "local": "/tmp/orbit-root/linc-1d4f-0-be90b4793a29",
        "status": "NONE"
      },
      {
        "fd": 15,
        "pid": 7510,
        "type": "file",
        "local": "/tmp/keyring-xnU0KR/socket.ssh",
        "status": "NONE"
      },
      {
        "fd": 16,
        "pid": 7510,
        "type": "file",
        "local": "/tmp/keyring-xnU0KR/socket.pkcs11",
        "status": "NONE"
      },
      {
        "fd": 17,
        "pid": 7530,
        "type": "file",
        "local": "/tmp/unique/org.gnome.Nautilus.:0.0.2445",
        "status": "NONE"
      },
      {
        "fd": 14,
        "pid": 7524,
        "type": "file",
        "local": "/tmp/orbit-root/linc-1d64-0-2ba998ee19d98",
        "status": "NONE"
      },
      {
        "fd": 12,
        "pid": 7529,
        "type": "file",
        "local": "/tmp/orbit-root/linc-1d69-0-5c67c60a52801",
        "status": "NONE"
      },
      {
        "fd": 14,
        "pid": 7532,
        "type": "file",
        "local": "/tmp/orbit-root/linc-1d6c-0-6364b853b7368",
        "status": "NONE"
      },
      {
        "fd": 3,
        "pid": 7499,
        "type": "file",
        "local": "@/tmp/dbus-Y3goIvV0OG",
        "status": "NONE"
      },
      {
        "fd": 5,
        "pid": 7574,
        "type": "file",
        "local": "/tmp/unique/org.gnome.VolumeControlApplet.:0.0.2495",
        "status": "NONE"
      },
      {
        "fd": 14,
        "pid": 7573,
        "type": "file",
        "local": "/tmp/.esd-0/socket",
        "status": "NONE"
      },
      {
        "fd": 15,
        "pid": 7573,
        "type": "file",
        "local": "/root/.pulse/e0095e13b3adca835bd8427c632106f7-runtime/native",
        "status": "NONE"
      },
      {
        "fd": 14,
        "pid": 7530,
        "type": "file",
        "local": "/tmp/orbit-root/linc-1d6a-0-782fc3762c876",
        "status": "NONE"
      },
      {
        "fd": 10,
        "pid": 7541,
        "type": "file",
        "local": "/tmp/orbit-root/linc-1d75-0-7dbf73039a2d2",
        "status": "NONE"
      },
      {
        "fd": 13,
        "pid": 7556,
        "type": "file",
        "local": "/tmp/orbit-root/linc-1d84-0-759b402a1fdb9",
        "status": "NONE"
      },
      {
        "fd": 12,
        "pid": 7581,
        "type": "file",
        "local": "/tmp/orbit-root/linc-1d92-0-3d004ff896b14",
        "status": "NONE"
      },
      {
        "fd": 13,
        "pid": 7575,
        "type": "file",
        "local": "/tmp/orbit-root/linc-1d86-0-55b2a497be926",
        "status": "NONE"
      },
      {
        "fd": 13,
        "pid": 7589,
        "type": "file",
        "local": "/tmp/orbit-root/linc-1da5-0-5ec8985cc909",
        "status": "NONE"
      },
      {
        "fd": 3,
        "pid": 1738,
        "type": "file",
        "local": "/var/run/dbus/system_bus_socket",
        "status": "NONE"
      },
      {
        "fd": 9,
        "pid": 1847,
        "type": "file",
        "local": "@/var/run/hald/dbus-F87hY9KqOB",
        "status": "NONE"
      },
      {
        "fd": 7,
        "pid": 3594,
        "type": "file",
        "local": "@/tmp/gdm-greeter-ExxtIQaP",
        "status": "NONE"
      },
      {
        "fd": 1,
        "pid": 3625,
        "type": "file",
        "local": "@/tmp/.X11-unix/X0",
        "status": "NONE"
      },
      {
        "fd": 7,
        "pid": 1847,
        "type": "file",
        "local": "@/var/run/hald/dbus-mz0fTOy5vF",
        "status": "NONE"
      },
      {
        "fd": 4,
        "pid": 601,
        "type": "file",
        "local": "@/org/kernel/udev/udevd",
        "status": "NONE"
      },
      {
        "fd": 20,
        "pid": 3430,
        "type": "file",
        "local": "public/cleanup",
        "status": "NONE"
      },
      {
        "fd": 26,
        "pid": 3430,
        "type": "file",
        "local": "private/rewrite",
        "status": "NONE"
      },
      {
        "fd": 29,
        "pid": 3430,
        "type": "file",
        "local": "private/bounce",
        "status": "NONE"
      },
      {
        "fd": 32,
        "pid": 3430,
        "type": "file",
        "local": "private/defer",
        "status": "NONE"
      },
      {
        "fd": 35,
        "pid": 3430,
        "type": "file",
        "local": "private/trace",
        "status": "NONE"
      },
      {
        "fd": 38,
        "pid": 3430,
        "type": "file",
        "local": "private/verify",
        "status": "NONE"
      },
      {
        "fd": 41,
        "pid": 3430,
        "type": "file",
        "local": "public/flush",
        "status": "NONE"
      },
      {
        "fd": 44,
        "pid": 3430,
        "type": "file",
        "local": "private/proxymap",
        "status": "NONE"
      },
      {
        "fd": 47,
        "pid": 3430,
        "type": "file",
        "local": "private/smtp",
        "status": "NONE"
      },
      {
        "fd": 50,
        "pid": 3430,
        "type": "file",
        "local": "private/relay",
        "status": "NONE"
      },
      {
        "fd": 53,
        "pid": 3430,
        "type": "file",
        "local": "public/showq",
        "status": "NONE"
      },
      {
        "fd": 56,
        "pid": 3430,
        "type": "file",
        "local": "private/error",
        "status": "NONE"
      },
      {
        "fd": 59,
        "pid": 3430,
        "type": "file",
        "local": "private/discard",
        "status": "NONE"
      },
      {
        "fd": 62,
        "pid": 3430,
        "type": "file",
        "local": "private/local",
        "status": "NONE"
      },
      {
        "fd": 65,
        "pid": 3430,
        "type": "file",
        "local": "private/virtual",
        "status": "NONE"
      },
      {
        "fd": 68,
        "pid": 3430,
        "type": "file",
        "local": "private/lmtp",
        "status": "NONE"
      },
      {
        "fd": 71,
        "pid": 3430,
        "type": "file",
        "local": "private/anvil",
        "status": "NONE"
      },
      {
        "fd": 74,
        "pid": 3430,
        "type": "file",
        "local": "private/scache",
        "status": "NONE"
      },
      {
        "fd": 77,
        "pid": 3430,
        "type": "file",
        "local": "private/maildrop",
        "status": "NONE"
      },
      {
        "fd": 80,
        "pid": 3430,
        "type": "file",
        "local": "private/cyrus",
        "status": "NONE"
      },
      {
        "fd": 83,
        "pid": 3430,
        "type": "file",
        "local": "private/uucp",
        "status": "NONE"
      },
      {
        "fd": 86,
        "pid": 3430,
        "type": "file",
        "local": "private/ifmail",
        "status": "NONE"
      },
      {
        "fd": 89,
        "pid": 3430,
        "type": "file",
        "local": "private/bsmtp",
        "status": "NONE"
      },
      {
        "fd": 92,
        "pid": 3430,
        "type": "file",
        "local": "private/procmail",
        "status": "NONE"
      },
      {
        "fd": 95,
        "pid": 3430,
        "type": "file",
        "local": "private/retry",
        "status": "NONE"
      },
      {
        "fd": 98,
        "pid": 3430,
        "type": "file",
        "local": "private/proxywrite",
        "status": "NONE"
      },
      {
        "fd": 14,
        "pid": 1847,
        "type": "file",
        "local": "@/org/freedesktop/hal/udev_event",
        "status": "NONE"
      },
      {
        "fd": 4,
        "pid": 3325,
        "type": "file",
        "local": "/var/run/cups/cups.sock",
        "status": "NONE"
      },
      {
        "fd": 9,
        "pid": 3594,
        "type": "file",
        "local": "@/tmp/gdm-session-CCrErvDI",
        "status": "NONE"
      },
      {
        "fd": 4,
        "pid": 1723,
        "type": "file",
        "local": "/var/run/acpid.socket",
        "status": "NONE"
      },
      {
        "fd": 9,
        "pid": 7418,
        "type": "file",
        "local": "@/tmp/.ICE-unix/7418",
        "status": "NONE"
      },
      {
        "fd": 7,
        "pid": 10904,
        "type": "file",
        "status": "NONE"
      },
      {
        "fd": 29,
        "pid": 7418,
        "type": "file",
        "status": "NONE"
      },
      {
        "fd": 8,
        "pid": 7555,
        "type": "file",
        "local": "@/dbus-vfs-daemon/socket-DU5l20Af",
        "status": "NONE"
      },
      {
        "fd": 6,
        "pid": 7555,
        "type": "file",
        "local": "@/dbus-vfs-daemon/socket-ss7UglUY",
        "status": "NONE"
      },
      {
        "fd": 15,
        "pid": 7545,
        "type": "file",
        "local": "@/dbus-vfs-daemon/socket-AsDBeSXW",
        "status": "NONE"
      },
      {
        "fd": 10,
        "pid": 7545,
        "type": "file",
        "local": "@/dbus-vfs-daemon/socket-hWh7F9CX",
        "status": "NONE"
      },
      {
        "fd": 11,
        "pid": 7545,
        "type": "file",
        "local": "@/dbus-vfs-daemon/socket-mxeQ8Xy1",
        "status": "NONE"
      },
      {
        "fd": 9,
        "pid": 7545,
        "type": "file",
        "local": "@/dbus-vfs-daemon/socket-9XIlbzJY",
        "status": "NONE"
      }
    ]
  }
}
```