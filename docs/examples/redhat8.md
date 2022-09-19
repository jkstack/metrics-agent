static:

```json
{
  "code": 200,
  "payload": {
    "time": 1663568841,
    "host": {
      "name": "localhost.localdomain",
      "uptime": 271325
    },
    "os": {
      "name": "linux",
      "platform_name": "redhat",
      "platform_version": "8.6",
      "install": 1663211950,
      "startup": 1663297516
    },
    "kernel": {
      "version": "4.18.0-372.9.1.el8.x86_64",
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
      "physical": 3913089024,
      "swap": 4248825856
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
        "total": 48336711680,
        "inodes": 23613440
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
        "mac": "00:50:56:a3:aa:f5",
        "addrs": [
          "192.168.2.67/24",
          "fe80::250:56ff:fea3:aaf5/64"
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
        "id": "65534",
        "gid": "65534"
      },
      {
        "name": "dbus",
        "id": "81",
        "gid": "81"
      },
      {
        "name": "systemd-coredump",
        "id": "999",
        "gid": "997"
      },
      {
        "name": "systemd-resolve",
        "id": "193",
        "gid": "193"
      },
      {
        "name": "tss",
        "id": "59",
        "gid": "59"
      },
      {
        "name": "polkitd",
        "id": "998",
        "gid": "996"
      },
      {
        "name": "unbound",
        "id": "997",
        "gid": "994"
      },
      {
        "name": "sssd",
        "id": "996",
        "gid": "993"
      },
      {
        "name": "sshd",
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
        "usage": 0.43
      },
      "memory": {
        "used": 242163712,
        "free": 3397689344,
        "available": 3435401216,
        "usage": 6.19,
        "swap_used": 0,
        "swap_free": 4248825856
      },
      "partitions": [
        {
          "name": "/",
          "used": 2079227904,
          "free": 46257483776,
          "usage": 4.3,
          "inode_used": 34326,
          "inode_free": 23579114,
          "inode_usage": 0.15
        },
        {
          "name": "/boot",
          "used": 219901952,
          "free": 843354112,
          "usage": 20.68,
          "inode_used": 310,
          "inode_free": 523978,
          "inode_usage": 0.06
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
          "name": "ens192",
          "bytes_sent": 9089660,
          "bytes_recv": 40637374,
          "packets_sent": 70464,
          "packets_recv": 421373
        }
      ]
    },
    "process": [
      {
        "id": 1,
        "pid": 0,
        "user": "root",
        "cpu_usage": 0,
        "rss": 13553664,
        "vms": 179126272,
        "swap": 0,
        "memory_usage": 0.35,
        "cmd": [
          "/usr/lib/systemd/systemd",
          "--switched-root",
          "--system",
          "--deserialize",
          "17"
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
        "cpu_usage": 0.02,
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
        "id": 76,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 178,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 179,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 180,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 181,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 182,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 183,
        "pid": 2,
        "user": "root",
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
        "id": 185,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 186,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 187,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 188,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 189,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 190,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 191,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 192,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 193,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 194,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 195,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 196,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 197,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 198,
        "pid": 2,
        "user": "root",
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
        "id": 200,
        "pid": 2,
        "user": "root",
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
        "id": 505,
        "pid": 2,
        "user": "root",
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
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 507,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 508,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 511,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 512,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 513,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 514,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 518,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 521,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
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
        "id": 524,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 526,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 527,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 528,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 530,
        "pid": 2,
        "user": "root",
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
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 533,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 534,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 535,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 536,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 537,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 538,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 539,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 540,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 541,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 542,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
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
        "id": 544,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 545,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 546,
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
        "id": 547,
        "pid": 2,
        "user": "root",
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
        "id": 550,
        "pid": 2,
        "user": "root",
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
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 554,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 555,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 556,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 557,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 558,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 559,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 560,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 561,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 562,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 563,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 564,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 565,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 566,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 567,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 568,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 569,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 570,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 571,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 572,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 573,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 574,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 575,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 576,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 577,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 578,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 579,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 580,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 581,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 582,
        "pid": 2,
        "user": "root",
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
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 584,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 585,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 586,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 587,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 588,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 589,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 591,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 592,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 594,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 597,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 689,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 698,
        "pid": 2,
        "user": "root",
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
        "id": 728,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 729,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 730,
        "pid": 2,
        "user": "root",
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
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 732,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 734,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 735,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 828,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 9691136,
        "vms": 91639808,
        "swap": 0,
        "memory_usage": 0.25,
        "cmd": [
          "/usr/lib/systemd/systemd-journald"
        ],
        "conns": 4
      },
      {
        "id": 867,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 11739136,
        "vms": 119787520,
        "swap": 0,
        "memory_usage": 0.3,
        "cmd": [
          "/usr/lib/systemd/systemd-udevd"
        ],
        "conns": 3
      },
      {
        "id": 869,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 870,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 871,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 872,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 873,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 874,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
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
        "id": 956,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2330624,
        "vms": 76800000,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "/sbin/auditd"
        ],
        "conns": 2
      },
      {
        "id": 978,
        "pid": 1,
        "user": "polkitd",
        "cpu_usage": 0,
        "rss": 25579520,
        "vms": 1670340608,
        "swap": 0,
        "memory_usage": 0.65,
        "cmd": [
          "/usr/lib/polkit-1/polkitd",
          "--no-debug"
        ],
        "conns": 2
      },
      {
        "id": 979,
        "pid": 1,
        "user": "dbus",
        "cpu_usage": 0,
        "rss": 6139904,
        "vms": 87035904,
        "swap": 0,
        "memory_usage": 0.16,
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
        "id": 980,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 11694080,
        "vms": 88281088,
        "swap": 0,
        "memory_usage": 0.3,
        "cmd": [
          "/usr/bin/VGAuthService",
          "-s"
        ],
        "conns": 3
      },
      {
        "id": 981,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.13,
        "rss": 12083200,
        "vms": 456945664,
        "swap": 0,
        "memory_usage": 0.31,
        "cmd": [
          "/usr/bin/vmtoolsd"
        ],
        "conns": 1
      },
      {
        "id": 982,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 5525504,
        "vms": 127991808,
        "swap": 0,
        "memory_usage": 0.14,
        "cmd": [
          "/usr/sbin/irqbalance",
          "--foreground"
        ],
        "conns": 2
      },
      {
        "id": 983,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 11874304,
        "vms": 197103616,
        "swap": 0,
        "memory_usage": 0.3,
        "cmd": [
          "/usr/sbin/sssd",
          "-i",
          "--logger=files"
        ],
        "conns": 3
      },
      {
        "id": 1044,
        "pid": 983,
        "user": "root",
        "cpu_usage": 0,
        "rss": 12824576,
        "vms": 205787136,
        "swap": 0,
        "memory_usage": 0.33,
        "cmd": [
          "/usr/libexec/sssd/sssd_be",
          "--domain",
          "implicit_files",
          "--uid",
          "0",
          "--gid",
          "0",
          "--logger=files"
        ],
        "conns": 3
      },
      {
        "id": 1049,
        "pid": 983,
        "user": "root",
        "cpu_usage": 0,
        "rss": 38948864,
        "vms": 195969024,
        "swap": 0,
        "memory_usage": 1,
        "cmd": [
          "/usr/libexec/sssd/sssd_nss",
          "--uid",
          "0",
          "--gid",
          "0",
          "--logger=files"
        ],
        "conns": 3
      },
      {
        "id": 1052,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 43552768,
        "vms": 314425344,
        "swap": 0,
        "memory_usage": 1.11,
        "cmd": [
          "/usr/libexec/platform-python",
          "-s",
          "/usr/sbin/firewalld",
          "--nofork",
          "--nopid"
        ],
        "conns": 2
      },
      {
        "id": 1053,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 9691136,
        "vms": 96526336,
        "swap": 0,
        "memory_usage": 0.25,
        "cmd": [
          "/usr/lib/systemd/systemd-logind"
        ],
        "conns": 2
      },
      {
        "id": 1065,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.02,
        "rss": 21127168,
        "vms": 401260544,
        "swap": 0,
        "memory_usage": 0.54,
        "cmd": [
          "/usr/sbin/NetworkManager",
          "--no-daemon"
        ],
        "conns": 2
      },
      {
        "id": 1074,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.3,
        "rss": 33226752,
        "vms": 508354560,
        "swap": 0,
        "memory_usage": 0.85,
        "cmd": [
          "/usr/libexec/platform-python",
          "-Es",
          "/usr/sbin/tuned",
          "-l",
          "-P"
        ],
        "conns": 1
      },
      {
        "id": 1075,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 7860224,
        "vms": 94625792,
        "swap": 0,
        "memory_usage": 0.2,
        "cmd": [
          "/usr/sbin/sshd",
          "-D",
          "-oCiphers=aes256-gcm@openssh.com,chacha20-poly1305@openssh.com,aes256-ctr,aes256-cbc,aes128-gcm@openssh.com,aes128-ctr,aes128-cbc",
          "-oMACs=hmac-sha2-256-etm@openssh.com,hmac-sha1-etm@openssh.com,umac-128-etm@openssh.com,hmac-sha2-512-etm@openssh.com,hmac-sha2-256,hmac-sha1,umac-128@openssh.com,hmac-sha2-512",
          "-oGSSAPIKexAlgorithms=gss-curve25519-sha256-,gss-nistp256-sha256-,gss-group14-sha256-,gss-group16-sha512-,gss-gex-sha1-,gss-group14-sha1-",
          "-oKexAlgorithms=curve25519-sha256,curve25519-sha256@libssh.org,ecdh-sha2-nistp256,ecdh-sha2-nistp384,ecdh-sha2-nistp521,diffie-hellman-group-exchange-sha256,diffie-hellman-group14-sha256,diffie-hellman-group16-sha512,diffie-hellman-group18-sha512,diffie-hellman-group-exchange-sha1,diffie-hellman-group14-sha1",
          "-oHostKeyAlgorithms=ecdsa-sha2-nistp256,ecdsa-sha2-nistp256-cert-v01@openssh.com,ecdsa-sha2-nistp384,ecdsa-sha2-nistp384-cert-v01@openssh.com,ecdsa-sha2-nistp521,ecdsa-sha2-nistp521-cert-v01@openssh.com,ssh-ed25519,ssh-ed25519-cert-v01@openssh.com,rsa-sha2-256,rsa-sha2-256-cert-v01@openssh.com,rsa-sha2-512,rsa-sha2-512-cert-v01@openssh.com,ssh-rsa,ssh-rsa-cert-v01@openssh.com",
          "-oPubkeyAcceptedKeyTypes=ecdsa-sha2-nistp256,ecdsa-sha2-nistp256-cert-v01@openssh.com,ecdsa-sha2-nistp384,ecdsa-sha2-nistp384-cert-v01@openssh.com,ecdsa-sha2-nistp521,ecdsa-sha2-nistp521-cert-v01@openssh.com,ssh-ed25519,ssh-ed25519-cert-v01@openssh.com,rsa-sha2-256,rsa-sha2-256-cert-v01@openssh.com,rsa-sha2-512,rsa-sha2-512-cert-v01@openssh.com,ssh-rsa,ssh-rsa-cert-v01@openssh.com",
          "-oCASignatureAlgorithms=ecdsa-sha2-nistp256,ecdsa-sha2-nistp384,ecdsa-sha2-nistp521,ssh-ed25519,rsa-sha2-256,rsa-sha2-512,ssh-rsa"
        ],
        "listen": [
          22,
          22
        ],
        "conns": 3
      },
      {
        "id": 1080,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2301952,
        "vms": 42790912,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "/usr/bin/rhsmcertd"
        ],
        "conns": 0
      },
      {
        "id": 1086,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3674112,
        "vms": 37875712,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "/usr/sbin/crond",
          "-n"
        ],
        "conns": 2
      },
      {
        "id": 1095,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1744896,
        "vms": 13987840,
        "swap": 0,
        "memory_usage": 0.04,
        "cmd": [
          "/sbin/agetty",
          "-o",
          "-p -- \\u",
          "--noclear",
          "tty1",
          "linux"
        ],
        "conns": 1
      },
      {
        "id": 1279,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 7196672,
        "vms": 214536192,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "/usr/sbin/rsyslogd",
          "-n"
        ],
        "conns": 1
      },
      {
        "id": 2824,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.04,
        "rss": 43577344,
        "vms": 1287208960,
        "swap": 0,
        "memory_usage": 1.11,
        "cmd": [
          "/opt/metrics-agent/metrics",
          "-conf",
          "/opt/metrics-agent/agent.conf"
        ],
        "conns": 2
      },
      {
        "id": 2985,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 4236,
        "pid": 1075,
        "user": "root",
        "cpu_usage": 0,
        "rss": 10686464,
        "vms": 157184000,
        "swap": 0,
        "memory_usage": 0.27,
        "cmd": [
          "sshd: root [priv]"
        ],
        "conns": 3
      },
      {
        "id": 4247,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 9789440,
        "vms": 91664384,
        "swap": 0,
        "memory_usage": 0.25,
        "cmd": [
          "/usr/lib/systemd/systemd",
          "--user"
        ],
        "conns": 5
      },
      {
        "id": 4250,
        "pid": 4247,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5402624,
        "vms": 235749376,
        "swap": 0,
        "memory_usage": 0.14,
        "cmd": [
          "(sd-pam)"
        ],
        "conns": 2
      },
      {
        "id": 4256,
        "pid": 4236,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5738496,
        "vms": 157184000,
        "swap": 0,
        "memory_usage": 0.15,
        "cmd": [
          "sshd: root@pts/0"
        ],
        "conns": 3
      },
      {
        "id": 4257,
        "pid": 4256,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3956736,
        "vms": 26877952,
        "swap": 0,
        "memory_usage": 0.1,
        "cmd": [
          "-bash"
        ],
        "conns": 1
      },
      {
        "id": 4322,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 4388,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 4398,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 4415,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 4417,
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
        "id": 4418,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0.09,
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
        "fd": 5,
        "pid": 1075,
        "type": "tcp4",
        "local": "0.0.0.0:22",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 5,
        "pid": 4236,
        "type": "tcp4",
        "local": "192.168.2.67:22",
        "remote": "10.202.0.9:54790",
        "status": "ESTABLISHED"
      },
      {
        "fd": 3,
        "pid": 2824,
        "type": "tcp4",
        "local": "192.168.2.67:48732",
        "remote": "192.168.2.52:13081",
        "status": "ESTABLISHED"
      },
      {
        "fd": 7,
        "pid": 1075,
        "type": "tcp6",
        "local": ":::22",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 5,
        "pid": 982,
        "type": "file",
        "local": "@irqbalance982.sock",
        "status": "NONE"
      },
      {
        "fd": 15,
        "pid": 4247,
        "type": "file",
        "local": "/run/user/0/systemd/notify",
        "status": "NONE"
      },
      {
        "fd": 18,
        "pid": 4247,
        "type": "file",
        "local": "/run/user/0/systemd/private",
        "status": "NONE"
      },
      {
        "fd": 11,
        "pid": 4247,
        "type": "file",
        "local": "/run/user/0/bus",
        "status": "NONE"
      },
      {
        "fd": 89,
        "pid": 1,
        "type": "file",
        "local": "/run/dbus/system_bus_socket",
        "status": "NONE"
      },
      {
        "fd": 38,
        "pid": 1,
        "type": "file",
        "local": "/var/run/.heim_org.h5l.kcm-socket",
        "status": "NONE"
      },
      {
        "fd": 9,
        "pid": 980,
        "type": "file",
        "local": "/var/run/vmware/guestServicePipe",
        "status": "NONE"
      },
      {
        "fd": 12,
        "pid": 983,
        "type": "file",
        "local": "/var/lib/sss/pipes/private/sbus-monitor",
        "status": "NONE"
      },
      {
        "fd": 15,
        "pid": 1044,
        "type": "file",
        "local": "/var/lib/sss/pipes/private/sbus-dp_implicit_files.1044",
        "status": "NONE"
      },
      {
        "fd": 15,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/private",
        "status": "NONE"
      },
      {
        "fd": 28,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/notify",
        "status": "NONE"
      },
      {
        "fd": 29,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/cgroups-agent",
        "status": "NONE"
      },
      {
        "fd": 41,
        "pid": 1,
        "type": "file",
        "local": "/run/lvm/lvmpolld.socket",
        "status": "NONE"
      },
      {
        "fd": 87,
        "pid": 1,
        "type": "unix",
        "local": "/run/udev/control",
        "status": "NONE"
      },
      {
        "fd": 39,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/dev-log",
        "status": "NONE"
      },
      {
        "fd": 36,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/stdout",
        "status": "NONE"
      },
      {
        "fd": 16,
        "pid": 1049,
        "type": "file",
        "local": "/var/lib/sss/pipes/nss",
        "status": "NONE"
      },
      {
        "fd": 37,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/socket",
        "status": "NONE"
      },
      {
        "fd": 88,
        "pid": 1,
        "type": "unix",
        "local": "/run/systemd/coredump",
        "status": "NONE"
      },
      {
        "fd": 8,
        "pid": 983,
        "type": "file",
        "status": "NONE"
      },
      {
        "fd": 19,
        "pid": 4247,
        "type": "file",
        "status": "NONE"
      }
    ]
  }
}
```