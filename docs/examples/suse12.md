static:

```json
{
  "code": 200,
  "payload": {
    "time": 1663568459,
    "host": {
      "name": "linux-be4e",
      "uptime": 421717
    },
    "os": {
      "name": "linux",
      "platform_name": "suse",
      "platform_version": "12.5",
      "install": 1663116275,
      "startup": 1663146743
    },
    "kernel": {
      "version": "4.12.14-120-default",
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
      "physical": 3877720064,
      "swap": 2153771008
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
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/.snapshots",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/boot/grub2/i386-pc",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/usr/local",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/var/tmp",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/var/log",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/var/spool",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/tmp",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/var/opt",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/var/crash",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/srv",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/opt",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/var/lib/mailman",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/var/lib/pgsql",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/var/lib/libvirt/images",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/boot/grub2/x86_64-efi",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/var/lib/mysql",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/var/lib/named",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/var/lib/machines",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/var/lib/mariadb",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/var/cache",
        "fstype": "btrfs",
        "opts": [
          "rw",
          "relatime",
          "bind"
        ],
        "total": 21015560192
      },
      {
        "mount": "/home",
        "fstype": "xfs",
        "opts": [
          "rw",
          "relatime"
        ],
        "total": 30501810176,
        "inodes": 14900736
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
        "mac": "00:50:56:a3:2c:a7",
        "addrs": [
          "192.168.2.63/24",
          "fe80::250:56ff:fea3:2ca7/64"
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
        "name": "games",
        "id": "12",
        "gid": "100"
      },
      {
        "name": "man",
        "id": "13",
        "gid": "62"
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
        "id": "499",
        "gid": "499"
      },
      {
        "name": "sshd",
        "id": "498",
        "gid": "498"
      },
      {
        "name": "polkitd",
        "id": "497",
        "gid": "495"
      },
      {
        "name": "nscd",
        "id": "496",
        "gid": "494"
      },
      {
        "name": "rpc",
        "id": "495",
        "gid": "65534"
      },
      {
        "name": "openslp",
        "id": "494",
        "gid": "2"
      },
      {
        "name": "systemd-timesync",
        "id": "492",
        "gid": "492"
      },
      {
        "name": "postfix",
        "id": "51",
        "gid": "51"
      },
      {
        "name": "ftpsecure",
        "id": "491",
        "gid": "65534"
      },
      {
        "name": "ntp",
        "id": "74",
        "gid": "491"
      },
      {
        "name": "at",
        "id": "25",
        "gid": "25"
      },
      {
        "name": "statd",
        "id": "490",
        "gid": "65534"
      },
      {
        "name": "scard",
        "id": "489",
        "gid": "490"
      },
      {
        "name": "rtkit",
        "id": "488",
        "gid": "489"
      },
      {
        "name": "pulse",
        "id": "487",
        "gid": "487"
      },
      {
        "name": "srvGeoClue",
        "id": "486",
        "gid": "65534"
      },
      {
        "name": "vnc",
        "id": "485",
        "gid": "485"
      },
      {
        "name": "gdm",
        "id": "484",
        "gid": "484"
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
        "usage": 0.26
      },
      "memory": {
        "used": 368697344,
        "free": 1860255744,
        "available": 3238801408,
        "usage": 9.51,
        "swap_used": 0,
        "swap_free": 2153771008
      },
      "partitions": [
        {
          "name": "/",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/.snapshots",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/boot/grub2/i386-pc",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/usr/local",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/var/tmp",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/var/log",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/var/spool",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/tmp",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/var/opt",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/var/crash",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/srv",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/opt",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/var/lib/mailman",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/var/lib/pgsql",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/var/lib/libvirt/images",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/boot/grub2/x86_64-efi",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/var/lib/mysql",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/var/lib/named",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/var/lib/machines",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/var/lib/mariadb",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/var/cache",
          "used": 3926769664,
          "free": 15231041536,
          "usage": 20.5
        },
        {
          "name": "/home",
          "used": 33783808,
          "free": 30468026368,
          "usage": 0.11,
          "inode_used": 3,
          "inode_free": 14900733
        }
      ],
      "interface": [
        {
          "name": "lo",
          "bytes_sent": 25480,
          "bytes_recv": 25480,
          "packets_sent": 424,
          "packets_recv": 424
        },
        {
          "name": "eth0",
          "bytes_sent": 13935035,
          "bytes_recv": 55093742,
          "packets_sent": 106787,
          "packets_recv": 649346
        }
      ]
    },
    "process": [
      {
        "id": 1,
        "pid": 0,
        "user": "root",
        "cpu_usage": 0,
        "rss": 6246400,
        "vms": 194232320,
        "swap": 0,
        "memory_usage": 0.16,
        "cmd": [
          "/usr/lib/systemd/systemd",
          "--switched-root",
          "--system",
          "--deserialize",
          "23"
        ],
        "conns": 14
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
        "cpu_usage": 0.01,
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
        "id": 341,
        "pid": 2,
        "user": "root",
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
        "id": 364,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 365,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 366,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 367,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 368,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 369,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 370,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 371,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 372,
        "pid": 2,
        "user": "root",
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
        "id": 375,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 376,
        "pid": 2,
        "user": "root",
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
        "id": 379,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 380,
        "pid": 2,
        "user": "root",
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
        "id": 382,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 383,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 384,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 385,
        "pid": 2,
        "user": "root",
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
        "id": 420,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 421,
        "pid": 2,
        "user": "root",
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
        "cpu_usage": 0,
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
        "id": 430,
        "pid": 2,
        "user": "root",
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
        "pid": 2,
        "user": "root",
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
        "id": 442,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 454,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 455,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 456,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 457,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 458,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 459,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 460,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 461,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 462,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 463,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 464,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 465,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 466,
        "pid": 2,
        "user": "root",
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
        "id": 469,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 470,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 471,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 472,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 473,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 474,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 475,
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
        "id": 548,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 8634368,
        "vms": 36196352,
        "swap": 0,
        "memory_usage": 0.22,
        "cmd": [
          "/usr/lib/systemd/systemd-journald"
        ],
        "conns": 4
      },
      {
        "id": 550,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 4751360,
        "vms": 12320768,
        "swap": 0,
        "memory_usage": 0.12,
        "cmd": [
          "/usr/sbin/haveged",
          "-w",
          "1024",
          "-v",
          "0",
          "-F"
        ],
        "conns": 0
      },
      {
        "id": 567,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5537792,
        "vms": 50614272,
        "swap": 0,
        "memory_usage": 0.14,
        "cmd": [
          "/usr/lib/systemd/systemd-udevd"
        ],
        "conns": 3
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
        "id": 789,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 790,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 792,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 793,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 794,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 795,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 796,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 798,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 922,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 2842624,
        "vms": 19865600,
        "swap": 0,
        "memory_usage": 0.07,
        "cmd": [
          "/usr/sbin/irqbalance",
          "--foreground"
        ],
        "conns": 1
      },
      {
        "id": 926,
        "pid": 1,
        "user": "messagebus",
        "cpu_usage": 0,
        "rss": 4751360,
        "vms": 39301120,
        "swap": 0,
        "memory_usage": 0.12,
        "cmd": [
          "/bin/dbus-daemon",
          "--system",
          "--address=systemd:",
          "--nofork",
          "--nopidfile",
          "--systemd-activation"
        ],
        "conns": 3
      },
      {
        "id": 931,
        "pid": 1,
        "user": "nscd",
        "cpu_usage": 0,
        "rss": 2813952,
        "vms": 823980032,
        "swap": 0,
        "memory_usage": 0.07,
        "cmd": [
          "/usr/sbin/nscd"
        ],
        "conns": 2
      },
      {
        "id": 944,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 8261632,
        "vms": 43180032,
        "swap": 0,
        "memory_usage": 0.21,
        "cmd": [
          "/usr/bin/VGAuthService",
          "-s"
        ],
        "conns": 2
      },
      {
        "id": 950,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3047424,
        "vms": 23056384,
        "swap": 0,
        "memory_usage": 0.08,
        "cmd": [
          "/usr/lib/systemd/systemd-logind"
        ],
        "conns": 2
      },
      {
        "id": 972,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 4530176,
        "vms": 341770240,
        "swap": 0,
        "memory_usage": 0.12,
        "cmd": [
          "/usr/sbin/rsyslogd",
          "-n"
        ],
        "conns": 2
      },
      {
        "id": 984,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.11,
        "rss": 10481664,
        "vms": 182874112,
        "swap": 0,
        "memory_usage": 0.27,
        "cmd": [
          "/usr/bin/vmtoolsd"
        ],
        "conns": 0
      },
      {
        "id": 1143,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5672960,
        "vms": 44695552,
        "swap": 0,
        "memory_usage": 0.15,
        "cmd": [
          "/usr/lib/wicked/bin/wickedd-dhcp6",
          "--systemd",
          "--foreground"
        ],
        "conns": 2
      },
      {
        "id": 1144,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5877760,
        "vms": 44691456,
        "swap": 0,
        "memory_usage": 0.15,
        "cmd": [
          "/usr/lib/wicked/bin/wickedd-auto4",
          "--systemd",
          "--foreground"
        ],
        "conns": 2
      },
      {
        "id": 1146,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5804032,
        "vms": 44695552,
        "swap": 0,
        "memory_usage": 0.15,
        "cmd": [
          "/usr/lib/wicked/bin/wickedd-dhcp4",
          "--systemd",
          "--foreground"
        ],
        "conns": 2
      },
      {
        "id": 1147,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 6197248,
        "vms": 44810240,
        "swap": 0,
        "memory_usage": 0.16,
        "cmd": [
          "/usr/sbin/wickedd",
          "--systemd",
          "--foreground"
        ],
        "conns": 2
      },
      {
        "id": 1150,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 6377472,
        "vms": 44720128,
        "swap": 0,
        "memory_usage": 0.16,
        "cmd": [
          "/usr/sbin/wickedd-nanny",
          "--systemd",
          "--foreground"
        ],
        "conns": 2
      },
      {
        "id": 1429,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 10059776,
        "vms": 50778112,
        "swap": 0,
        "memory_usage": 0.26,
        "cmd": [
          "/sbin/iscsid",
          "-f"
        ],
        "conns": 2
      },
      {
        "id": 1626,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 1889,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1708032,
        "vms": 4866048,
        "swap": 0,
        "memory_usage": 0.04,
        "cmd": [
          "/sbin/agetty",
          "--noclear",
          "tty1",
          "linux"
        ],
        "conns": 0
      },
      {
        "id": 1894,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 6823936,
        "vms": 66387968,
        "swap": 0,
        "memory_usage": 0.18,
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
        "id": 2001,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 6914048,
        "vms": 209018880,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "/usr/sbin/gdm"
        ],
        "conns": 1
      },
      {
        "id": 2016,
        "pid": 2001,
        "user": "root",
        "cpu_usage": 0,
        "rss": 7995392,
        "vms": 292192256,
        "swap": 0,
        "memory_usage": 0.21,
        "cmd": [
          "/usr/lib/gdm/gdm-simple-slave",
          "--display-id",
          "/org/gnome/DisplayManager/Displays/_0"
        ],
        "conns": 3
      },
      {
        "id": 2019,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3342336,
        "vms": 35000320,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "/usr/lib/postfix/bin//master",
          "-w"
        ],
        "listen": [
          25,
          25
        ],
        "conns": 24
      },
      {
        "id": 2021,
        "pid": 2019,
        "user": "postfix",
        "cpu_usage": 0,
        "rss": 5791744,
        "vms": 37289984,
        "swap": 0,
        "memory_usage": 0.15,
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
        "id": 2034,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2768896,
        "vms": 19361792,
        "swap": 0,
        "memory_usage": 0.07,
        "cmd": [
          "/usr/sbin/cron",
          "-n"
        ],
        "conns": 2
      },
      {
        "id": 2044,
        "pid": 2016,
        "user": "root",
        "cpu_usage": 0,
        "rss": 58531840,
        "vms": 353091584,
        "swap": 0,
        "memory_usage": 1.51,
        "cmd": [
          "/usr/bin/Xorg",
          ":0",
          "-background",
          "none",
          "-verbose",
          "-auth",
          "/run/gdm/auth-for-gdm-VwObz4/database",
          "-seat",
          "seat0",
          "vt7"
        ],
        "conns": 3
      },
      {
        "id": 2045,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 6868992,
        "vms": 287350784,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "/usr/lib/accounts-daemon"
        ],
        "conns": 1
      },
      {
        "id": 2049,
        "pid": 1,
        "user": "polkitd",
        "cpu_usage": 0,
        "rss": 17563648,
        "vms": 535347200,
        "swap": 0,
        "memory_usage": 0.45,
        "cmd": [
          "/usr/lib/polkit-1/polkitd",
          "--no-debug"
        ],
        "conns": 2
      },
      {
        "id": 2114,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 8257536,
        "vms": 294461440,
        "swap": 0,
        "memory_usage": 0.21,
        "cmd": [
          "/usr/lib/upower/upowerd"
        ],
        "conns": 1
      },
      {
        "id": 2124,
        "pid": 1,
        "user": "rtkit",
        "cpu_usage": 0,
        "rss": 3690496,
        "vms": 184647680,
        "swap": 0,
        "memory_usage": 0.1,
        "cmd": [
          "/usr/lib/rtkit/rtkit-daemon"
        ],
        "conns": 2
      },
      {
        "id": 5584,
        "pid": 2016,
        "user": "root",
        "cpu_usage": 0,
        "rss": 8605696,
        "vms": 259031040,
        "swap": 0,
        "memory_usage": 0.22,
        "cmd": [
          "gdm-session-worker [pam/gdm-password]"
        ],
        "conns": 2
      },
      {
        "id": 5588,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5005312,
        "vms": 42266624,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          "/usr/lib/systemd/systemd",
          "--user"
        ],
        "conns": 4
      },
      {
        "id": 5589,
        "pid": 5588,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1892352,
        "vms": 241569792,
        "swap": 0,
        "memory_usage": 0.05,
        "cmd": [
          "(sd-pam)"
        ],
        "conns": 2
      },
      {
        "id": 5594,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 8896512,
        "vms": 206950400,
        "swap": 0,
        "memory_usage": 0.23,
        "cmd": [
          "/usr/bin/gnome-keyring-daemon",
          "--daemonize",
          "--login"
        ],
        "conns": 4
      },
      {
        "id": 5600,
        "pid": 5584,
        "user": "root",
        "cpu_usage": 0,
        "rss": 15241216,
        "vms": 570195968,
        "swap": 0,
        "memory_usage": 0.39,
        "cmd": [
          "/usr/lib/gnome-session-binary",
          "--session",
          "gnome-classic"
        ],
        "conns": 3
      },
      {
        "id": 5653,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1945600,
        "vms": 14516224,
        "swap": 0,
        "memory_usage": 0.05,
        "cmd": [
          "/usr/bin/dbus-launch",
          "--sh-syntax",
          "--exit-with-session",
          "/usr/bin/gpg-agent",
          "--sh",
          "--daemon",
          "--keep-display",
          "--write-env-file",
          "/root/.gnupg/agent.info-linux-be4e:0",
          "/etc/X11/xinit/xinitrc",
          "--session",
          "gnome-classic"
        ],
        "conns": 1
      },
      {
        "id": 5654,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2871296,
        "vms": 38805504,
        "swap": 0,
        "memory_usage": 0.07,
        "cmd": [
          "/bin/dbus-daemon",
          "--fork",
          "--print-pid",
          "5",
          "--print-address",
          "7",
          "--session"
        ],
        "conns": 2
      },
      {
        "id": 5655,
        "pid": 5600,
        "user": "root",
        "cpu_usage": 0,
        "rss": 229376,
        "vms": 16678912,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "/usr/bin/gpg-agent",
          "--sh",
          "--daemon",
          "--keep-display",
          "--write-env-file",
          "/root/.gnupg/agent.info-linux-be4e:0",
          "/etc/X11/xinit/xinitrc",
          "--session",
          "gnome-classic"
        ],
        "conns": 1
      },
      {
        "id": 5661,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5554176,
        "vms": 343547904,
        "swap": 0,
        "memory_usage": 0.14,
        "cmd": [
          "/usr/lib/at-spi2/at-spi-bus-launcher"
        ],
        "conns": 1
      },
      {
        "id": 5666,
        "pid": 5661,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3915776,
        "vms": 38481920,
        "swap": 0,
        "memory_usage": 0.1,
        "cmd": [
          "/bin/dbus-daemon",
          "--config-file=/usr/share/defaults/at-spi2/accessibility.conf",
          "--nofork",
          "--print-address",
          "3"
        ],
        "conns": 2
      },
      {
        "id": 5668,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 6172672,
        "vms": 199483392,
        "swap": 0,
        "memory_usage": 0.16,
        "cmd": [
          "/usr/lib/at-spi2/at-spi2-registryd",
          "--use-gnome-session"
        ],
        "conns": 1
      },
      {
        "id": 5684,
        "pid": 5600,
        "user": "root",
        "cpu_usage": 0.07,
        "rss": 181686272,
        "vms": 1461669888,
        "swap": 0,
        "memory_usage": 4.69,
        "cmd": [
          "/usr/bin/gnome-shell"
        ],
        "conns": 2
      },
      {
        "id": 5690,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 8511488,
        "vms": 278130688,
        "swap": 0,
        "memory_usage": 0.22,
        "cmd": [
          "/usr/lib/gvfs/gvfsd"
        ],
        "conns": 1
      },
      {
        "id": 5695,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 6074368,
        "vms": 416419840,
        "swap": 0,
        "memory_usage": 0.16,
        "cmd": [
          "/usr/lib/gvfs/gvfsd-fuse",
          "/run/user/0/gvfs",
          "-f",
          "-o",
          "big_writes"
        ],
        "conns": 1
      },
      {
        "id": 5718,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 8019968,
        "vms": 303538176,
        "swap": 0,
        "memory_usage": 0.21,
        "cmd": [
          "/usr/lib/gvfs/gvfs-udisks2-volume-monitor"
        ],
        "conns": 1
      },
      {
        "id": 5721,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 9826304,
        "vms": 376135680,
        "swap": 0,
        "memory_usage": 0.25,
        "cmd": [
          "/usr/lib/udisks2/udisksd",
          "--no-debug"
        ],
        "conns": 2
      },
      {
        "id": 5729,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 4567040,
        "vms": 270196736,
        "swap": 0,
        "memory_usage": 0.12,
        "cmd": [
          "/usr/lib/gvfs/gvfs-mtp-volume-monitor"
        ],
        "conns": 1
      },
      {
        "id": 5734,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5099520,
        "vms": 279482368,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          "/usr/lib/gvfs/gvfs-gphoto2-volume-monitor"
        ],
        "conns": 1
      },
      {
        "id": 5739,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5066752,
        "vms": 260820992,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          "/usr/lib/gvfs/gvfs-goa-volume-monitor"
        ],
        "conns": 1
      },
      {
        "id": 5744,
        "pid": 5600,
        "user": "root",
        "cpu_usage": 0.02,
        "rss": 37601280,
        "vms": 1051344896,
        "swap": 0,
        "memory_usage": 0.97,
        "cmd": [
          "/usr/lib/gnome-settings-daemon-3.0/gnome-settings-daemon"
        ],
        "conns": 1
      },
      {
        "id": 5760,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 8232960,
        "vms": 300924928,
        "swap": 0,
        "memory_usage": 0.21,
        "cmd": [
          "/usr/bin/pulseaudio",
          "--start"
        ],
        "conns": 3
      },
      {
        "id": 5768,
        "pid": 5600,
        "user": "root",
        "cpu_usage": 0,
        "rss": 43077632,
        "vms": 713773056,
        "swap": 0,
        "memory_usage": 1.11,
        "cmd": [
          "nautilus",
          "--no-default-window",
          "--force-desktop"
        ],
        "conns": 1
      },
      {
        "id": 5770,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 11763712,
        "vms": 493944832,
        "swap": 0,
        "memory_usage": 0.3,
        "cmd": [
          "/usr/lib/gnome-settings-daemon-3.0/gsd-printer"
        ],
        "conns": 1
      },
      {
        "id": 5775,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 4620288,
        "vms": 182607872,
        "swap": 0,
        "memory_usage": 0.12,
        "cmd": [
          "/usr/lib/dconf-service"
        ],
        "conns": 1
      },
      {
        "id": 5797,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.27,
        "rss": 34578432,
        "vms": 399773696,
        "swap": 0,
        "memory_usage": 0.89,
        "cmd": [
          "/usr/bin/vmtoolsd",
          "-n",
          "vmusr"
        ],
        "conns": 1
      },
      {
        "id": 5799,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 6692864,
        "vms": 356081664,
        "swap": 0,
        "memory_usage": 0.17,
        "cmd": [
          "/usr/lib/gvfs/gvfsd-trash",
          "--spawner",
          ":1.7",
          "/org/gtk/gvfs/exec_spaw/0"
        ],
        "conns": 3
      },
      {
        "id": 5810,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 6389760,
        "vms": 353632256,
        "swap": 0,
        "memory_usage": 0.16,
        "cmd": [
          "/usr/lib/gvfs/gvfsd-burn",
          "--spawner",
          ":1.7",
          "/org/gtk/gvfs/exec_spaw/1"
        ],
        "conns": 2
      },
      {
        "id": 5822,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 32862208,
        "vms": 591962112,
        "swap": 0,
        "memory_usage": 0.85,
        "cmd": [
          "/usr/lib/gnome-terminal-server"
        ],
        "conns": 1
      },
      {
        "id": 5828,
        "pid": 5822,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5128192,
        "vms": 14417920,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          "bash"
        ],
        "conns": 0
      },
      {
        "id": 11113,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 11206,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.04,
        "rss": 38645760,
        "vms": 1139433472,
        "swap": 0,
        "memory_usage": 1,
        "cmd": [
          "/opt/metrics-agent/metrics",
          "-conf",
          "/opt/metrics-agent/agent.conf"
        ],
        "conns": 2
      },
      {
        "id": 21554,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 21591,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 22509,
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
        "id": 22586,
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
        "id": 22663,
        "pid": 2019,
        "user": "postfix",
        "cpu_usage": 0,
        "rss": 4341760,
        "vms": 36896768,
        "swap": 0,
        "memory_usage": 0.11,
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
        "id": 22664,
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
        "id": 22743,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0.03,
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
        "pid": 1894,
        "type": "tcp4",
        "local": "0.0.0.0:22",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 13,
        "pid": 2019,
        "type": "tcp4",
        "local": "127.0.0.1:25",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 3,
        "pid": 11206,
        "type": "tcp4",
        "local": "192.168.2.63:53612",
        "remote": "192.168.2.52:13081",
        "status": "ESTABLISHED"
      },
      {
        "fd": 4,
        "pid": 1894,
        "type": "tcp6",
        "local": ":::22",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 14,
        "pid": 2019,
        "type": "tcp6",
        "local": "::1:25",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 10,
        "pid": 2016,
        "type": "file",
        "local": "@/tmp/dbus-a0caeDAX",
        "status": "NONE"
      },
      {
        "fd": 15,
        "pid": 5594,
        "type": "file",
        "local": "/run/user/0/keyring/pkcs11",
        "status": "NONE"
      },
      {
        "fd": 15,
        "pid": 5760,
        "type": "file",
        "local": "/run/user/0/pulse/native",
        "status": "NONE"
      },
      {
        "fd": 10,
        "pid": 5600,
        "type": "file",
        "local": "@/tmp/.ICE-unix/5600",
        "status": "NONE"
      },
      {
        "fd": 5,
        "pid": 2044,
        "type": "file",
        "local": "@/tmp/.X11-unix/X0",
        "status": "NONE"
      },
      {
        "fd": 13,
        "pid": 5588,
        "type": "file",
        "local": "/run/user/0/systemd/notify",
        "status": "NONE"
      },
      {
        "fd": 14,
        "pid": 5588,
        "type": "file",
        "local": "/run/user/0/systemd/private",
        "status": "NONE"
      },
      {
        "fd": 21,
        "pid": 2019,
        "type": "file",
        "local": "public/cleanup",
        "status": "NONE"
      },
      {
        "fd": 27,
        "pid": 2019,
        "type": "file",
        "local": "private/rewrite",
        "status": "NONE"
      },
      {
        "fd": 30,
        "pid": 2019,
        "type": "file",
        "local": "private/bounce",
        "status": "NONE"
      },
      {
        "fd": 33,
        "pid": 2019,
        "type": "file",
        "local": "private/defer",
        "status": "NONE"
      },
      {
        "fd": 9,
        "pid": 5594,
        "type": "file",
        "local": "/run/user/0/keyring/control",
        "status": "NONE"
      },
      {
        "fd": 36,
        "pid": 2019,
        "type": "file",
        "local": "private/trace",
        "status": "NONE"
      },
      {
        "fd": 39,
        "pid": 2019,
        "type": "file",
        "local": "private/verify",
        "status": "NONE"
      },
      {
        "fd": 42,
        "pid": 2019,
        "type": "file",
        "local": "public/flush",
        "status": "NONE"
      },
      {
        "fd": 45,
        "pid": 2019,
        "type": "file",
        "local": "private/proxymap",
        "status": "NONE"
      },
      {
        "fd": 48,
        "pid": 2019,
        "type": "file",
        "local": "private/proxywrite",
        "status": "NONE"
      },
      {
        "fd": 51,
        "pid": 2019,
        "type": "file",
        "local": "private/smtp",
        "status": "NONE"
      },
      {
        "fd": 54,
        "pid": 2019,
        "type": "file",
        "local": "private/relay",
        "status": "NONE"
      },
      {
        "fd": 57,
        "pid": 2019,
        "type": "file",
        "local": "public/showq",
        "status": "NONE"
      },
      {
        "fd": 60,
        "pid": 2019,
        "type": "file",
        "local": "private/error",
        "status": "NONE"
      },
      {
        "fd": 63,
        "pid": 2019,
        "type": "file",
        "local": "private/retry",
        "status": "NONE"
      },
      {
        "fd": 66,
        "pid": 2019,
        "type": "file",
        "local": "private/discard",
        "status": "NONE"
      },
      {
        "fd": 69,
        "pid": 2019,
        "type": "file",
        "local": "private/local",
        "status": "NONE"
      },
      {
        "fd": 72,
        "pid": 2019,
        "type": "file",
        "local": "private/virtual",
        "status": "NONE"
      },
      {
        "fd": 75,
        "pid": 2019,
        "type": "file",
        "local": "private/lmtp",
        "status": "NONE"
      },
      {
        "fd": 78,
        "pid": 2019,
        "type": "file",
        "local": "private/anvil",
        "status": "NONE"
      },
      {
        "fd": 81,
        "pid": 2019,
        "type": "file",
        "local": "private/scache",
        "status": "NONE"
      },
      {
        "fd": 127,
        "pid": 1,
        "type": "file",
        "local": "/run/dbus/system_bus_socket",
        "status": "NONE"
      },
      {
        "fd": 129,
        "pid": 1,
        "type": "file",
        "local": "/var/run/pcscd/pcscd.comm",
        "status": "NONE"
      },
      {
        "fd": 5,
        "pid": 5655,
        "type": "file",
        "local": "/tmp/gpg-5W0Bvp/S.gpg-agent",
        "status": "NONE"
      },
      {
        "fd": 11,
        "pid": 5600,
        "type": "file",
        "local": "/tmp/.ICE-unix/5600",
        "status": "NONE"
      },
      {
        "fd": 14,
        "pid": 5760,
        "type": "file",
        "local": "/tmp/.esd-0/socket",
        "status": "NONE"
      },
      {
        "fd": 6,
        "pid": 2044,
        "type": "file",
        "local": "/tmp/.X11-unix/X0",
        "status": "NONE"
      },
      {
        "fd": 11,
        "pid": 2016,
        "type": "file",
        "local": "@/tmp/dbus-9YSrrWjy",
        "status": "NONE"
      },
      {
        "fd": 4,
        "pid": 5654,
        "type": "file",
        "local": "@/tmp/dbus-6Ji52bIABj",
        "status": "NONE"
      },
      {
        "fd": 11,
        "pid": 5594,
        "type": "file",
        "local": "/run/user/0/keyring/ssh",
        "status": "NONE"
      },
      {
        "fd": 76,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/notify",
        "status": "NONE"
      },
      {
        "fd": 77,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/cgroups-agent",
        "status": "NONE"
      },
      {
        "fd": 128,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/dev-log",
        "status": "NONE"
      },
      {
        "fd": 123,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/stdout",
        "status": "NONE"
      },
      {
        "fd": 124,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/socket",
        "status": "NONE"
      },
      {
        "fd": 13,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/private",
        "status": "NONE"
      },
      {
        "fd": 130,
        "pid": 1,
        "type": "file",
        "local": "@ISCSIADM_ABSTRACT_NAMESPACE",
        "status": "NONE"
      },
      {
        "fd": 115,
        "pid": 1,
        "type": "file",
        "local": "/run/lvm/lvmetad.socket",
        "status": "NONE"
      },
      {
        "fd": 119,
        "pid": 1,
        "type": "unix",
        "local": "/run/udev/control",
        "status": "NONE"
      },
      {
        "fd": 5,
        "pid": 5666,
        "type": "file",
        "local": "@/tmp/dbus-m00XglYYGN",
        "status": "NONE"
      },
      {
        "fd": 121,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/syslog",
        "status": "NONE"
      },
      {
        "fd": 7,
        "pid": 944,
        "type": "file",
        "local": "/var/run/vmware/guestServicePipe",
        "status": "NONE"
      },
      {
        "fd": 16,
        "pid": 931,
        "type": "file",
        "local": "/var/run/nscd/socket",
        "status": "NONE"
      },
      {
        "fd": 15,
        "pid": 1,
        "type": "file",
        "status": "NONE"
      },
      {
        "fd": 21,
        "pid": 5760,
        "type": "file",
        "status": "NONE"
      },
      {
        "fd": 12,
        "pid": 5799,
        "type": "file",
        "local": "@/dbus-vfs-daemon/socket-MhWJ6XZE",
        "status": "NONE"
      },
      {
        "fd": 10,
        "pid": 5799,
        "type": "file",
        "local": "@/dbus-vfs-daemon/socket-gSnNcL1P",
        "status": "NONE"
      },
      {
        "fd": 8,
        "pid": 5810,
        "type": "file",
        "local": "@/dbus-vfs-daemon/socket-OZ4qK3JE",
        "status": "NONE"
      }
    ]
  }
}
```