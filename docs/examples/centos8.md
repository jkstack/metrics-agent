static:

```json
{
  "code": 200,
  "payload": {
    "time": 1663224204,
    "host": {
      "name": "localhost.localdomain",
      "uptime": 532035
    },
    "os": {
      "name": "linux",
      "platform_name": "centos",
      "platform_version": "8.5.2111",
      "install": 1662690987,
      "startup": 1662692169
    },
    "kernel": {
      "version": "4.18.0-348.el8.x86_64",
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
      "physical": 3917819904,
      "swap": 4257214464
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
        "total": 48328327168,
        "inodes": 23609344
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
        "mac": "00:50:56:a3:95:0e",
        "addrs": [
          "192.168.2.53/24",
          "fe80::250:56ff:fea3:950e/64"
        ]
      },
      {
        "index": 3,
        "name": "virbr0",
        "mtu": 1500,
        "flags": [
          "up",
          "broadcast",
          "multicast"
        ],
        "mac": "52:54:00:09:08:7a",
        "addrs": [
          "192.168.122.1/24"
        ]
      },
      {
        "index": 4,
        "name": "virbr0-nic",
        "mtu": 1500,
        "flags": [
          "broadcast",
          "multicast"
        ],
        "mac": "52:54:00:09:08:7a"
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
        "name": "Kernel Overflow User",
        "id": "65534",
        "gid": "65534"
      },
      {
        "name": "System message bus",
        "id": "81",
        "gid": "81"
      },
      {
        "name": "systemd Core Dumper",
        "id": "999",
        "gid": "997"
      },
      {
        "name": "systemd Resolver",
        "id": "193",
        "gid": "193"
      },
      {
        "name": "Account used for TPM access",
        "id": "59",
        "gid": "59"
      },
      {
        "name": "User for polkitd",
        "id": "998",
        "gid": "996"
      },
      {
        "name": "User for geoclue",
        "id": "997",
        "gid": "995"
      },
      {
        "name": "RealtimeKit",
        "id": "172",
        "gid": "172"
      },
      {
        "name": "PipeWire System Daemon",
        "id": "996",
        "gid": "992"
      },
      {
        "name": "PulseAudio System Daemon",
        "id": "171",
        "gid": "171"
      },
      {
        "name": "qemu user",
        "id": "107",
        "gid": "107"
      },
      {
        "name": "usbmuxd user",
        "id": "113",
        "gid": "113"
      },
      {
        "name": "Unbound DNS resolver",
        "id": "995",
        "gid": "989"
      },
      {
        "name": "GlusterFS daemons",
        "id": "994",
        "gid": "988"
      },
      {
        "name": "Rpcbind Daemon",
        "id": "32",
        "gid": "32"
      },
      {
        "name": "Avahi mDNS/DNS-SD Stack",
        "id": "70",
        "gid": "70"
      },
      {
        "name": "",
        "id": "993",
        "gid": "987"
      },
      {
        "name": "Saslauthd user",
        "id": "992",
        "gid": "76"
      },
      {
        "name": "daemon account for libstoragemgmt",
        "id": "991",
        "gid": "985"
      },
      {
        "name": "Dnsmasq DHCP and DNS server",
        "id": "983",
        "gid": "983"
      },
      {
        "name": "radvd user",
        "id": "75",
        "gid": "75"
      },
      {
        "name": "User for sssd",
        "id": "982",
        "gid": "982"
      },
      {
        "name": "User for cockpit web service",
        "id": "981",
        "gid": "981"
      },
      {
        "name": "User for cockpit-ws instances",
        "id": "980",
        "gid": "980"
      },
      {
        "name": "User for colord",
        "id": "979",
        "gid": "979"
      },
      {
        "name": "RPC Service User",
        "id": "29",
        "gid": "29"
      },
      {
        "name": "",
        "id": "978",
        "gid": "978"
      },
      {
        "name": "User for flatpak system helper",
        "id": "977",
        "gid": "977"
      },
      {
        "name": "",
        "id": "42",
        "gid": "42"
      },
      {
        "name": "Clevis Decryption Framework unprivileged user",
        "id": "976",
        "gid": "976"
      },
      {
        "name": "",
        "id": "975",
        "gid": "975"
      },
      {
        "name": "",
        "id": "72",
        "gid": "72"
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
        "usage": 1.82
      },
      "memory": {
        "used": 780533760,
        "free": 2264530944,
        "available": 2871099392,
        "usage": 19.92,
        "swap_used": 0,
        "swap_free": 4257214464
      },
      "partitions": [
        {
          "name": "/",
          "used": 5092208640,
          "free": 43236118528,
          "usage": 10.54,
          "inode_used": 122559,
          "inode_free": 23486785,
          "inode_usage": 0.52
        },
        {
          "name": "/boot",
          "used": 271273984,
          "free": 791982080,
          "usage": 25.51,
          "inode_used": 310,
          "inode_free": 523978,
          "inode_usage": 0.06
        }
      ],
      "interface": [
        {
          "name": "lo",
          "bytes_sent": 10850460,
          "bytes_recv": 10850460,
          "packets_sent": 111649,
          "packets_recv": 111649
        },
        {
          "name": "ens160",
          "bytes_sent": 3551278,
          "bytes_recv": 98270868,
          "packets_sent": 37997,
          "packets_recv": 792580
        },
        {
          "name": "virbr0",
          "bytes_sent": 0,
          "bytes_recv": 0,
          "packets_sent": 0,
          "packets_recv": 0
        },
        {
          "name": "virbr0-nic",
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
        "cpu_usage": 0.03,
        "rss": 14417920,
        "vms": 247095296,
        "swap": 0,
        "memory_usage": 0.37,
        "cmd": [
          "/usr/lib/systemd/systemd",
          "--switched-root",
          "--system",
          "--deserialize",
          "18"
        ],
        "listen": [
          111,
          111
        ],
        "conns": 28
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
        "id": 75,
        "pid": 2,
        "user": "root",
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
        "id": 499,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.62,
        "rss": 8593408,
        "vms": 58314752,
        "swap": 0,
        "memory_usage": 0.22,
        "cmd": [
          "@usr/sbin/plymouthd",
          "--mode=boot",
          "--pid-file=/var/run/plymouth/pid",
          "--attach-to-session"
        ],
        "conns": 1
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
        "id": 509,
        "pid": 2,
        "user": "root",
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
        "id": 515,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 516,
        "pid": 2,
        "user": "root",
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
        "id": 522,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 525,
        "pid": 2,
        "user": "root",
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
        "id": 529,
        "pid": 2,
        "user": "root",
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
        "id": 531,
        "pid": 2,
        "user": "root",
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
        "id": 605,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 612,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 638,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 639,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 640,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 641,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 642,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 643,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 644,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 645,
        "pid": 2,
        "user": "root",
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
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 743,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 10907648,
        "vms": 91992064,
        "swap": 0,
        "memory_usage": 0.28,
        "cmd": [
          "/usr/lib/systemd/systemd-journald"
        ],
        "conns": 4
      },
      {
        "id": 783,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 13594624,
        "vms": 122015744,
        "swap": 0,
        "memory_usage": 0.35,
        "cmd": [
          "/usr/lib/systemd/systemd-udevd"
        ],
        "conns": 3
      },
      {
        "id": 845,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 846,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 847,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 848,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 849,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 850,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 851,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 876,
        "pid": 1,
        "user": "rpc",
        "cpu_usage": 0,
        "rss": 5591040,
        "vms": 68812800,
        "swap": 0,
        "memory_usage": 0.14,
        "cmd": [
          "/usr/bin/rpcbind",
          "-w",
          "-f"
        ],
        "listen": [
          111,
          111
        ],
        "conns": 6
      },
      {
        "id": 879,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2682880,
        "vms": 154361856,
        "swap": 0,
        "memory_usage": 0.07,
        "cmd": [
          "/sbin/auditd"
        ],
        "conns": 2
      },
      {
        "id": 881,
        "pid": 879,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2326528,
        "vms": 49725440,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "/usr/sbin/sedispatch"
        ],
        "conns": 1
      },
      {
        "id": 887,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 888,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 889,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 908,
        "pid": 1,
        "user": "libstoragemgmt",
        "cpu_usage": 0,
        "rss": 2039808,
        "vms": 20213760,
        "swap": 0,
        "memory_usage": 0.05,
        "cmd": [
          "/usr/bin/lsmd",
          "-d"
        ],
        "conns": 3
      },
      {
        "id": 909,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 14839808,
        "vms": 224714752,
        "swap": 0,
        "memory_usage": 0.38,
        "cmd": [
          "/usr/sbin/sssd",
          "-i",
          "--logger=files"
        ],
        "conns": 3
      },
      {
        "id": 910,
        "pid": 1,
        "user": "dbus",
        "cpu_usage": 0,
        "rss": 7274496,
        "vms": 74555392,
        "swap": 0,
        "memory_usage": 0.19,
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
        "id": 911,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 7036928,
        "vms": 81014784,
        "swap": 0,
        "memory_usage": 0.18,
        "cmd": [
          "/usr/lib/systemd/systemd-machined"
        ],
        "conns": 2
      },
      {
        "id": 912,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2269184,
        "vms": 18210816,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "/usr/sbin/mcelog",
          "--ignorenodev",
          "--daemon",
          "--foreground"
        ],
        "conns": 2
      },
      {
        "id": 913,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 11907072,
        "vms": 88743936,
        "swap": 0,
        "memory_usage": 0.3,
        "cmd": [
          "/usr/bin/VGAuthService",
          "-s"
        ],
        "conns": 3
      },
      {
        "id": 914,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.12,
        "rss": 12800000,
        "vms": 375193600,
        "swap": 0,
        "memory_usage": 0.33,
        "cmd": [
          "/usr/bin/vmtoolsd"
        ],
        "conns": 1
      },
      {
        "id": 915,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 15884288,
        "vms": 569221120,
        "swap": 0,
        "memory_usage": 0.41,
        "cmd": [
          "/usr/libexec/udisks2/udisksd"
        ],
        "conns": 2
      },
      {
        "id": 916,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 5480448,
        "vms": 51470336,
        "swap": 0,
        "memory_usage": 0.14,
        "cmd": [
          "/usr/sbin/smartd",
          "-n",
          "-q",
          "never"
        ],
        "conns": 0
      },
      {
        "id": 917,
        "pid": 1,
        "user": "avahi",
        "cpu_usage": 0,
        "rss": 5001216,
        "vms": 87388160,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          "avahi-daemon: running [linux-2.local]"
        ],
        "conns": 7
      },
      {
        "id": 920,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 12935168,
        "vms": 474140672,
        "swap": 0,
        "memory_usage": 0.33,
        "cmd": [
          "/usr/sbin/ModemManager"
        ],
        "conns": 2
      },
      {
        "id": 921,
        "pid": 1,
        "user": "polkitd",
        "cpu_usage": 0,
        "rss": 29827072,
        "vms": 1679015936,
        "swap": 0,
        "memory_usage": 0.76,
        "cmd": [
          "/usr/lib/polkit-1/polkitd",
          "--no-debug"
        ],
        "conns": 2
      },
      {
        "id": 923,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 5443584,
        "vms": 127987712,
        "swap": 0,
        "memory_usage": 0.14,
        "cmd": [
          "/usr/sbin/irqbalance",
          "--foreground"
        ],
        "conns": 2
      },
      {
        "id": 924,
        "pid": 1,
        "user": "rtkit",
        "cpu_usage": 0,
        "rss": 3633152,
        "vms": 207720448,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "/usr/libexec/rtkit-daemon"
        ],
        "conns": 2
      },
      {
        "id": 925,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.02,
        "rss": 21303296,
        "vms": 418537472,
        "swap": 0,
        "memory_usage": 0.54,
        "cmd": [
          "/usr/sbin/NetworkManager",
          "--no-daemon"
        ],
        "conns": 2
      },
      {
        "id": 929,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2523136,
        "vms": 27222016,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "/bin/bash",
          "/usr/sbin/ksmtuned"
        ],
        "conns": 1
      },
      {
        "id": 992,
        "pid": 917,
        "user": "avahi",
        "cpu_usage": 0,
        "rss": 446464,
        "vms": 87252992,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "avahi-daemon: chroot helper"
        ],
        "conns": 3
      },
      {
        "id": 993,
        "pid": 909,
        "user": "root",
        "cpu_usage": 0,
        "rss": 15863808,
        "vms": 233779200,
        "swap": 0,
        "memory_usage": 0.4,
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
        "id": 995,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 8028160,
        "vms": 96739328,
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
        "id": 996,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.03,
        "rss": 33509376,
        "vms": 508551168,
        "swap": 0,
        "memory_usage": 0.86,
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
        "id": 1005,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3833856,
        "vms": 104435712,
        "swap": 0,
        "memory_usage": 0.1,
        "cmd": [
          "/usr/sbin/gssproxy",
          "-D"
        ],
        "conns": 4
      },
      {
        "id": 1006,
        "pid": 909,
        "user": "root",
        "cpu_usage": 0,
        "rss": 42336256,
        "vms": 235122688,
        "swap": 0,
        "memory_usage": 1.08,
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
        "id": 1016,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 9822208,
        "vms": 337620992,
        "swap": 0,
        "memory_usage": 0.25,
        "cmd": [
          "/usr/libexec/accounts-daemon"
        ],
        "conns": 1
      },
      {
        "id": 1017,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 10125312,
        "vms": 96747520,
        "swap": 0,
        "memory_usage": 0.26,
        "cmd": [
          "/usr/lib/systemd/systemd-logind"
        ],
        "conns": 2
      },
      {
        "id": 1018,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 11448320,
        "vms": 142635008,
        "swap": 0,
        "memory_usage": 0.29,
        "cmd": [
          "/usr/sbin/cupsd",
          "-l"
        ],
        "listen": [
          631,
          631
        ],
        "conns": 5
      },
      {
        "id": 1036,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 12107776,
        "vms": 358211584,
        "swap": 0,
        "memory_usage": 0.31,
        "cmd": [
          "/usr/lib/realmd/realmd"
        ],
        "conns": 2
      },
      {
        "id": 1049,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.01,
        "rss": 7647232,
        "vms": 214515712,
        "swap": 0,
        "memory_usage": 0.2,
        "cmd": [
          "/usr/sbin/rsyslogd",
          "-n"
        ],
        "conns": 1
      },
      {
        "id": 1056,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 11452416,
        "vms": 344629248,
        "swap": 0,
        "memory_usage": 0.29,
        "cmd": [
          "/usr/sbin/gdm"
        ],
        "conns": 5
      },
      {
        "id": 1064,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 3694592,
        "vms": 38060032,
        "swap": 0,
        "memory_usage": 0.09,
        "cmd": [
          "/usr/sbin/crond",
          "-n"
        ],
        "conns": 2
      },
      {
        "id": 1072,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 2478080,
        "vms": 45060096,
        "swap": 0,
        "memory_usage": 0.06,
        "cmd": [
          "/usr/sbin/atd",
          "-f"
        ],
        "conns": 1
      },
      {
        "id": 1073,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 1445888,
        "vms": 16883712,
        "swap": 0,
        "memory_usage": 0.04,
        "cmd": [
          "/usr/bin/plymouth",
          "--wait"
        ],
        "conns": 2
      },
      {
        "id": 1141,
        "pid": 1056,
        "user": "root",
        "cpu_usage": 0,
        "rss": 13795328,
        "vms": 333062144,
        "swap": 0,
        "memory_usage": 0.35,
        "cmd": [
          "gdm-session-worker [pam/gdm-launch-environment]"
        ],
        "conns": 2
      },
      {
        "id": 1324,
        "pid": 1,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 9719808,
        "vms": 91631616,
        "swap": 0,
        "memory_usage": 0.25,
        "cmd": [
          "/usr/lib/systemd/systemd",
          "--user"
        ],
        "conns": 7
      },
      {
        "id": 1330,
        "pid": 1324,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 5124096,
        "vms": 175325184,
        "swap": 0,
        "memory_usage": 0.13,
        "cmd": [
          "(sd-pam)"
        ],
        "conns": 2
      },
      {
        "id": 1576,
        "pid": 1324,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 9867264,
        "vms": 842670080,
        "swap": 0,
        "memory_usage": 0.25,
        "cmd": [
          "/usr/bin/pulseaudio",
          "--daemonize=no",
          "--log-target=journal"
        ],
        "conns": 4
      },
      {
        "id": 1602,
        "pid": 1141,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 8118272,
        "vms": 223383552,
        "swap": 0,
        "memory_usage": 0.21,
        "cmd": [
          "/usr/libexec/gdm-wayland-session",
          "--register-session",
          "dbus-run-session -- gnome-session --autostart /usr/share/gdm/greeter/autostart  --session gnome-initial-setup"
        ],
        "conns": 1
      },
      {
        "id": 1614,
        "pid": 1324,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 5685248,
        "vms": 86691840,
        "swap": 0,
        "memory_usage": 0.15,
        "cmd": [
          "/usr/bin/dbus-daemon",
          "--session",
          "--address=systemd:",
          "--nofork",
          "--nopidfile",
          "--systemd-activation",
          "--syslog-only"
        ],
        "conns": 2
      },
      {
        "id": 1677,
        "pid": 1602,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 2117632,
        "vms": 43089920,
        "swap": 0,
        "memory_usage": 0.05,
        "cmd": [
          "dbus-run-session",
          "--",
          "gnome-session",
          "--autostart",
          "/usr/share/gdm/greeter/autostart",
          "--session",
          "gnome-initial-setup"
        ],
        "conns": 1
      },
      {
        "id": 1678,
        "pid": 1677,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 6455296,
        "vms": 87322624,
        "swap": 0,
        "memory_usage": 0.16,
        "cmd": [
          "dbus-daemon",
          "--nofork",
          "--print-address",
          "4",
          "--session"
        ],
        "conns": 2
      },
      {
        "id": 1679,
        "pid": 1677,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 19595264,
        "vms": 884559872,
        "swap": 0,
        "memory_usage": 0.5,
        "cmd": [
          "/usr/libexec/gnome-session-binary",
          "--autostart",
          "/usr/share/gdm/greeter/autostart",
          "--session",
          "gnome-initial-setup"
        ],
        "conns": 4
      },
      {
        "id": 1776,
        "pid": 1,
        "user": "dnsmasq",
        "cpu_usage": 0,
        "rss": 2666496,
        "vms": 75165696,
        "swap": 0,
        "memory_usage": 0.07,
        "cmd": [
          "/usr/sbin/dnsmasq",
          "--conf-file=/var/lib/libvirt/dnsmasq/default.conf",
          "--leasefile-ro",
          "--dhcp-script=/usr/libexec/libvirt_leaseshelper"
        ],
        "listen": [
          53
        ],
        "conns": 5
      },
      {
        "id": 1777,
        "pid": 1776,
        "user": "root",
        "cpu_usage": 0,
        "rss": 446464,
        "vms": 75059200,
        "swap": 0,
        "memory_usage": 0.01,
        "cmd": [
          "/usr/sbin/dnsmasq",
          "--conf-file=/var/lib/libvirt/dnsmasq/default.conf",
          "--leasefile-ro",
          "--dhcp-script=/usr/libexec/libvirt_leaseshelper"
        ],
        "conns": 0
      },
      {
        "id": 1823,
        "pid": 1679,
        "user": "gnome-initial-setup",
        "cpu_usage": 1.08,
        "rss": 284139520,
        "vms": 3307843584,
        "swap": 0,
        "memory_usage": 7.25,
        "cmd": [
          "gnome-shell",
          "--mode=initial-setup"
        ],
        "conns": 4
      },
      {
        "id": 1840,
        "pid": 1,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 11665408,
        "vms": 320905216,
        "swap": 0,
        "memory_usage": 0.3,
        "cmd": [
          "/usr/libexec/gvfsd"
        ],
        "conns": 1
      },
      {
        "id": 1848,
        "pid": 1823,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 53051392,
        "vms": 882982912,
        "swap": 0,
        "memory_usage": 1.35,
        "cmd": [
          "/usr/bin/Xwayland",
          ":1024",
          "-rootless",
          "-terminate",
          "-accessx",
          "-core",
          "-auth",
          "/run/user/975/.mutter-Xwaylandauth.GQ21R1",
          "-listen",
          "4",
          "-listen",
          "5",
          "-displayfd",
          "6"
        ],
        "conns": 3
      },
      {
        "id": 1865,
        "pid": 1,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 11243520,
        "vms": 381157376,
        "swap": 0,
        "memory_usage": 0.29,
        "cmd": [
          "/usr/libexec/at-spi-bus-launcher"
        ],
        "conns": 1
      },
      {
        "id": 1870,
        "pid": 1865,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 6012928,
        "vms": 86712320,
        "swap": 0,
        "memory_usage": 0.15,
        "cmd": [
          "/usr/bin/dbus-daemon",
          "--config-file=/usr/share/defaults/at-spi2/accessibility.conf",
          "--nofork",
          "--print-address",
          "3"
        ],
        "conns": 2
      },
      {
        "id": 1873,
        "pid": 1,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 8724480,
        "vms": 254672896,
        "swap": 0,
        "memory_usage": 0.22,
        "cmd": [
          "/usr/libexec/at-spi2-registryd",
          "--use-gnome-session"
        ],
        "conns": 1
      },
      {
        "id": 1881,
        "pid": 1823,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 12136448,
        "vms": 399355904,
        "swap": 0,
        "memory_usage": 0.31,
        "cmd": [
          "ibus-daemon",
          "--xim",
          "--panel",
          "disable"
        ],
        "conns": 2
      },
      {
        "id": 1885,
        "pid": 1881,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 11784192,
        "vms": 310644736,
        "swap": 0,
        "memory_usage": 0.3,
        "cmd": [
          "/usr/libexec/ibus-dconf"
        ],
        "conns": 1
      },
      {
        "id": 1886,
        "pid": 1881,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 29413376,
        "vms": 449048576,
        "swap": 0,
        "memory_usage": 0.75,
        "cmd": [
          "/usr/libexec/ibus-extension-gtk3"
        ],
        "conns": 1
      },
      {
        "id": 1890,
        "pid": 1,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 67747840,
        "vms": 1168949248,
        "swap": 0,
        "memory_usage": 1.73,
        "cmd": [
          "/usr/libexec/ibus-x11",
          "--kill-daemon"
        ],
        "conns": 1
      },
      {
        "id": 1896,
        "pid": 1,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 8941568,
        "vms": 310513664,
        "swap": 0,
        "memory_usage": 0.23,
        "cmd": [
          "/usr/libexec/ibus-portal"
        ],
        "conns": 1
      },
      {
        "id": 1897,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 9822208,
        "vms": 320688128,
        "swap": 0,
        "memory_usage": 0.25,
        "cmd": [
          "/usr/libexec/upowerd"
        ],
        "conns": 1
      },
      {
        "id": 1920,
        "pid": 1,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 8712192,
        "vms": 316952576,
        "swap": 0,
        "memory_usage": 0.22,
        "cmd": [
          "/usr/libexec/xdg-permission-store"
        ],
        "conns": 1
      },
      {
        "id": 1921,
        "pid": 1,
        "user": "geoclue",
        "cpu_usage": 0,
        "rss": 21864448,
        "vms": 374489088,
        "swap": 0,
        "memory_usage": 0.56,
        "cmd": [
          "/usr/libexec/geoclue"
        ],
        "conns": 1
      },
      {
        "id": 1930,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 27848704,
        "vms": 483110912,
        "swap": 0,
        "memory_usage": 0.71,
        "cmd": [
          "/usr/libexec/packagekitd"
        ],
        "conns": 2
      },
      {
        "id": 1936,
        "pid": 1679,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 24985600,
        "vms": 523026432,
        "swap": 0,
        "memory_usage": 0.64,
        "cmd": [
          "/usr/libexec/gsd-wacom"
        ],
        "conns": 1
      },
      {
        "id": 1937,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0,
        "rss": 6688768,
        "vms": 62509056,
        "swap": 0,
        "memory_usage": 0.17,
        "cmd": [
          "/usr/sbin/wpa_supplicant",
          "-c",
          "/etc/wpa_supplicant/wpa_supplicant.conf",
          "-u",
          "-s"
        ],
        "conns": 2
      },
      {
        "id": 1938,
        "pid": 1679,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 26791936,
        "vms": 537427968,
        "swap": 0,
        "memory_usage": 0.68,
        "cmd": [
          "/usr/libexec/gsd-xsettings"
        ],
        "conns": 1
      },
      {
        "id": 1940,
        "pid": 1679,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 9220096,
        "vms": 309362688,
        "swap": 0,
        "memory_usage": 0.24,
        "cmd": [
          "/usr/libexec/gsd-a11y-settings"
        ],
        "conns": 1
      },
      {
        "id": 1941,
        "pid": 1679,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 24985600,
        "vms": 382971904,
        "swap": 0,
        "memory_usage": 0.64,
        "cmd": [
          "/usr/libexec/gsd-clipboard"
        ],
        "conns": 1
      },
      {
        "id": 1943,
        "pid": 1679,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 26341376,
        "vms": 703995904,
        "swap": 0,
        "memory_usage": 0.67,
        "cmd": [
          "/usr/libexec/gsd-color"
        ],
        "conns": 2
      },
      {
        "id": 1946,
        "pid": 1679,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 17223680,
        "vms": 471695360,
        "swap": 0,
        "memory_usage": 0.44,
        "cmd": [
          "/usr/libexec/gsd-datetime"
        ],
        "conns": 1
      },
      {
        "id": 1951,
        "pid": 1679,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 10203136,
        "vms": 398532608,
        "swap": 0,
        "memory_usage": 0.26,
        "cmd": [
          "/usr/libexec/gsd-housekeeping"
        ],
        "conns": 2
      },
      {
        "id": 1953,
        "pid": 1679,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 27648000,
        "vms": 536256512,
        "swap": 0,
        "memory_usage": 0.71,
        "cmd": [
          "/usr/libexec/gsd-keyboard"
        ],
        "conns": 1
      },
      {
        "id": 1961,
        "pid": 1679,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 28422144,
        "vms": 1148481536,
        "swap": 0,
        "memory_usage": 0.73,
        "cmd": [
          "/usr/libexec/gsd-media-keys"
        ],
        "conns": 1
      },
      {
        "id": 1962,
        "pid": 1679,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 9150464,
        "vms": 309501952,
        "swap": 0,
        "memory_usage": 0.23,
        "cmd": [
          "/usr/libexec/gsd-mouse"
        ],
        "conns": 1
      },
      {
        "id": 1971,
        "pid": 1679,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 26955776,
        "vms": 637255680,
        "swap": 0,
        "memory_usage": 0.69,
        "cmd": [
          "/usr/libexec/gsd-power"
        ],
        "conns": 1
      },
      {
        "id": 1974,
        "pid": 1679,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 15048704,
        "vms": 371712000,
        "swap": 0,
        "memory_usage": 0.38,
        "cmd": [
          "/usr/libexec/gsd-print-notifications"
        ],
        "conns": 1
      },
      {
        "id": 1978,
        "pid": 1679,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 9080832,
        "vms": 462794752,
        "swap": 0,
        "memory_usage": 0.23,
        "cmd": [
          "/usr/libexec/gsd-rfkill"
        ],
        "conns": 1
      },
      {
        "id": 1983,
        "pid": 1679,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 9883648,
        "vms": 306921472,
        "swap": 0,
        "memory_usage": 0.25,
        "cmd": [
          "/usr/libexec/gsd-screensaver-proxy"
        ],
        "conns": 1
      },
      {
        "id": 1986,
        "pid": 1679,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 13488128,
        "vms": 466960384,
        "swap": 0,
        "memory_usage": 0.34,
        "cmd": [
          "/usr/libexec/gsd-sharing"
        ],
        "conns": 2
      },
      {
        "id": 1991,
        "pid": 1679,
        "user": "gnome-initial-setup",
        "cpu_usage": 0.02,
        "rss": 15405056,
        "vms": 486428672,
        "swap": 0,
        "memory_usage": 0.39,
        "cmd": [
          "/usr/libexec/gsd-smartcard"
        ],
        "conns": 1
      },
      {
        "id": 1993,
        "pid": 1679,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 11874304,
        "vms": 376627200,
        "swap": 0,
        "memory_usage": 0.3,
        "cmd": [
          "/usr/libexec/gsd-sound"
        ],
        "conns": 1
      },
      {
        "id": 2017,
        "pid": 1,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 11046912,
        "vms": 421199872,
        "swap": 0,
        "memory_usage": 0.28,
        "cmd": [
          "/usr/libexec/gvfs-udisks2-volume-monitor"
        ],
        "conns": 1
      },
      {
        "id": 2018,
        "pid": 1,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 7806976,
        "vms": 220090368,
        "swap": 0,
        "memory_usage": 0.2,
        "cmd": [
          "/usr/libexec/dconf-service"
        ],
        "conns": 1
      },
      {
        "id": 2034,
        "pid": 1,
        "user": "colord",
        "cpu_usage": 0,
        "rss": 14131200,
        "vms": 357150720,
        "swap": 0,
        "memory_usage": 0.36,
        "cmd": [
          "/usr/libexec/colord"
        ],
        "conns": 2
      },
      {
        "id": 2041,
        "pid": 1,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 8536064,
        "vms": 308273152,
        "swap": 0,
        "memory_usage": 0.22,
        "cmd": [
          "/usr/libexec/gvfs-mtp-volume-monitor"
        ],
        "conns": 1
      },
      {
        "id": 2053,
        "pid": 1,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 8896512,
        "vms": 321282048,
        "swap": 0,
        "memory_usage": 0.23,
        "cmd": [
          "/usr/libexec/gvfs-gphoto2-volume-monitor"
        ],
        "conns": 1
      },
      {
        "id": 2061,
        "pid": 1,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 15855616,
        "vms": 530563072,
        "swap": 0,
        "memory_usage": 0.4,
        "cmd": [
          "/usr/libexec/gsd-printer"
        ],
        "conns": 1
      },
      {
        "id": 2066,
        "pid": 1,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 10432512,
        "vms": 390979584,
        "swap": 0,
        "memory_usage": 0.27,
        "cmd": [
          "/usr/libexec/gvfs-afc-volume-monitor"
        ],
        "conns": 1
      },
      {
        "id": 2074,
        "pid": 1,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 9076736,
        "vms": 306003968,
        "swap": 0,
        "memory_usage": 0.23,
        "cmd": [
          "/usr/libexec/gvfs-goa-volume-monitor"
        ],
        "conns": 1
      },
      {
        "id": 2080,
        "pid": 1,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 40452096,
        "vms": 844468224,
        "swap": 0,
        "memory_usage": 1.03,
        "cmd": [
          "/usr/libexec/goa-daemon"
        ],
        "conns": 2
      },
      {
        "id": 2087,
        "pid": 1,
        "user": "gnome-initial-setup",
        "cpu_usage": 0.01,
        "rss": 16064512,
        "vms": 453521408,
        "swap": 0,
        "memory_usage": 0.41,
        "cmd": [
          "/usr/libexec/goa-identity-service"
        ],
        "conns": 2
      },
      {
        "id": 2097,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.13,
        "rss": 44302336,
        "vms": 225390592,
        "swap": 0,
        "memory_usage": 1.13,
        "cmd": [
          "/usr/libexec/sssd/sssd_kcm",
          "--uid",
          "0",
          "--gid",
          "0",
          "--logger=files"
        ],
        "conns": 3
      },
      {
        "id": 2103,
        "pid": 1881,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 11726848,
        "vms": 235073536,
        "swap": 0,
        "memory_usage": 0.3,
        "cmd": [
          "/usr/libexec/ibus-engine-simple"
        ],
        "conns": 1
      },
      {
        "id": 2113,
        "pid": 1679,
        "user": "gnome-initial-setup",
        "cpu_usage": 1.2,
        "rss": 159297536,
        "vms": 1814056960,
        "swap": 0,
        "memory_usage": 4.07,
        "cmd": [
          "/usr/libexec/gnome-initial-setup"
        ],
        "conns": 2
      },
      {
        "id": 2129,
        "pid": 1,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 7458816,
        "vms": 315072512,
        "swap": 0,
        "memory_usage": 0.19,
        "cmd": [
          "gnome-keyring-daemon",
          "--unlock"
        ],
        "conns": 4
      },
      {
        "id": 2161,
        "pid": 1881,
        "user": "gnome-initial-setup",
        "cpu_usage": 0,
        "rss": 15024128,
        "vms": 335499264,
        "swap": 0,
        "memory_usage": 0.38,
        "cmd": [
          "/usr/libexec/ibus-engine-libpinyin",
          "--ibus"
        ],
        "conns": 1
      },
      {
        "id": 51782,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.04,
        "rss": 51400704,
        "vms": 1291304960,
        "swap": 0,
        "memory_usage": 1.31,
        "cmd": [
          "/opt/agent-server/server",
          "-conf",
          "/opt/agent-server/server.conf"
        ],
        "listen": [
          13081
        ],
        "conns": 5
      },
      {
        "id": 51892,
        "pid": 1,
        "user": "root",
        "cpu_usage": 0.04,
        "rss": 48631808,
        "vms": 1287446528,
        "swap": 0,
        "memory_usage": 1.24,
        "cmd": [
          "/opt/metrics-agent/metrics",
          "-conf",
          "/opt/metrics-agent/agent.conf"
        ],
        "conns": 2
      },
      {
        "id": 70962,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 76438,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 76599,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 78041,
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
        "id": 78069,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 78503,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 78553,
        "pid": 2,
        "user": "root",
        "cpu_usage": 0,
        "rss": 0,
        "vms": 0,
        "swap": 0,
        "memory_usage": 0,
        "cmd": null,
        "conns": 0
      },
      {
        "id": 78587,
        "pid": 929,
        "user": "root",
        "cpu_usage": 0,
        "rss": 868352,
        "vms": 7700480,
        "swap": 0,
        "memory_usage": 0.02,
        "cmd": [
          "sleep",
          "60"
        ],
        "conns": 1
      }
    ],
    "connections": [
      {
        "fd": 6,
        "pid": 1776,
        "type": "tcp4",
        "local": "192.168.122.1:53",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 4,
        "pid": 995,
        "type": "tcp4",
        "local": "0.0.0.0:22",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 10,
        "pid": 1018,
        "type": "tcp4",
        "local": "127.0.0.1:631",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 147,
        "pid": 1,
        "type": "tcp4",
        "local": "0.0.0.0:111",
        "remote": "0.0.0.0",
        "status": "LISTEN"
      },
      {
        "fd": 39,
        "pid": 1823,
        "type": "tcp4",
        "local": "192.168.2.53:58968",
        "remote": "8.43.85.4:443",
        "status": "CLOSE_WAIT"
      },
      {
        "fd": 3,
        "pid": 51892,
        "type": "tcp4",
        "local": "127.0.0.1:41386",
        "remote": "127.0.0.1:13081",
        "status": "ESTABLISHED"
      },
      {
        "fd": 6,
        "pid": 995,
        "type": "tcp6",
        "local": ":::22",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 9,
        "pid": 1018,
        "type": "tcp6",
        "local": "::1:631",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 7,
        "pid": 51782,
        "type": "tcp6",
        "local": ":::13081",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 150,
        "pid": 1,
        "type": "tcp6",
        "local": ":::111",
        "remote": "::",
        "status": "LISTEN"
      },
      {
        "fd": 8,
        "pid": 51782,
        "type": "tcp6",
        "local": "127.0.0.1:13081",
        "remote": "127.0.0.1:41386",
        "status": "ESTABLISHED"
      },
      {
        "fd": 10,
        "pid": 51782,
        "type": "tcp6",
        "local": "192.168.2.53:13081",
        "remote": "10.202.0.227:54190",
        "status": "ESTABLISHED"
      },
      {
        "fd": 9,
        "pid": 51782,
        "type": "tcp6",
        "local": "192.168.2.53:13081",
        "remote": "10.202.0.9:64736",
        "status": "ESTABLISHED"
      },
      {
        "fd": 15,
        "pid": 917,
        "type": "udp4",
        "local": "0.0.0.0:5353",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 5,
        "pid": 1776,
        "type": "udp4",
        "local": "192.168.122.1:53",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 3,
        "pid": 1776,
        "type": "udp4",
        "local": "0.0.0.0:67",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 148,
        "pid": 1,
        "type": "udp4",
        "local": "0.0.0.0:111",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 17,
        "pid": 917,
        "type": "udp4",
        "local": "0.0.0.0:49405",
        "remote": "0.0.0.0",
        "status": "NONE"
      },
      {
        "fd": 18,
        "pid": 917,
        "type": "udp6",
        "local": ":::54294",
        "remote": "::",
        "status": "NONE"
      },
      {
        "fd": 16,
        "pid": 917,
        "type": "udp6",
        "local": ":::5353",
        "remote": "::",
        "status": "NONE"
      },
      {
        "fd": 151,
        "pid": 1,
        "type": "udp6",
        "local": ":::111",
        "remote": "::",
        "status": "NONE"
      },
      {
        "fd": 52,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/stdout",
        "status": "NONE"
      },
      {
        "fd": 11,
        "pid": 1679,
        "type": "file",
        "local": "@/tmp/.ICE-unix/1679",
        "status": "NONE"
      },
      {
        "fd": 15,
        "pid": 1324,
        "type": "file",
        "local": "/run/user/975/systemd/notify",
        "status": "NONE"
      },
      {
        "fd": 15,
        "pid": 993,
        "type": "file",
        "local": "/var/lib/sss/pipes/private/sbus-dp_implicit_files.993",
        "status": "NONE"
      },
      {
        "fd": 53,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/socket",
        "status": "NONE"
      },
      {
        "fd": 18,
        "pid": 1324,
        "type": "file",
        "local": "/run/user/975/systemd/private",
        "status": "NONE"
      },
      {
        "fd": 9,
        "pid": 1881,
        "type": "file",
        "local": "@/tmp/dbus-yGLSsCHh",
        "status": "NONE"
      },
      {
        "fd": 16,
        "pid": 1006,
        "type": "file",
        "local": "/var/lib/sss/pipes/nss",
        "status": "NONE"
      },
      {
        "fd": 8,
        "pid": 1005,
        "type": "file",
        "local": "/var/lib/gssproxy/default.sock",
        "status": "NONE"
      },
      {
        "fd": 55,
        "pid": 1,
        "type": "file",
        "local": "@/org/kernel/linux/storage/multipathd",
        "status": "NONE"
      },
      {
        "fd": 5,
        "pid": 923,
        "type": "file",
        "local": "@irqbalance923.sock",
        "status": "NONE"
      },
      {
        "fd": 16,
        "pid": 1576,
        "type": "file",
        "local": "/tmp/.esd-975/socket",
        "status": "NONE"
      },
      {
        "fd": 51,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/journal/dev-log",
        "status": "NONE"
      },
      {
        "fd": 7,
        "pid": 1678,
        "type": "file",
        "local": "@/tmp/dbus-VxAWu2Xyy5",
        "status": "NONE"
      },
      {
        "fd": 8,
        "pid": 1870,
        "type": "file",
        "local": "@/tmp/dbus-0uACcoGM8v",
        "status": "NONE"
      },
      {
        "fd": 141,
        "pid": 1,
        "type": "file",
        "local": "/run/lvm/lvmpolld.socket",
        "status": "NONE"
      },
      {
        "fd": 11,
        "pid": 1056,
        "type": "file",
        "local": "@/tmp/dbus-fL2IpYrP",
        "status": "NONE"
      },
      {
        "fd": 11,
        "pid": 1324,
        "type": "file",
        "local": "/run/user/975/bus",
        "status": "NONE"
      },
      {
        "fd": 25,
        "pid": 1324,
        "type": "file",
        "local": "/run/user/975/pulse/native",
        "status": "NONE"
      },
      {
        "fd": 26,
        "pid": 1324,
        "type": "file",
        "local": "/run/user/975/pipewire-0",
        "status": "NONE"
      },
      {
        "fd": 6,
        "pid": 499,
        "type": "file",
        "local": "@/org/freedesktop/plymouthd",
        "status": "NONE"
      },
      {
        "fd": 5,
        "pid": 912,
        "type": "file",
        "local": "/var/run/mcelog-client",
        "status": "NONE"
      },
      {
        "fd": 50,
        "pid": 1,
        "type": "file",
        "local": "@ISCSID_UIP_ABSTRACT_NAMESPACE",
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
        "fd": 32,
        "pid": 1,
        "type": "unix",
        "local": "/run/udev/control",
        "status": "NONE"
      },
      {
        "fd": 6,
        "pid": 2129,
        "type": "file",
        "local": "/run/user/975/keyring/control",
        "status": "NONE"
      },
      {
        "fd": 146,
        "pid": 1,
        "type": "file",
        "local": "/run/rpcbind.sock",
        "status": "NONE"
      },
      {
        "fd": 44,
        "pid": 1,
        "type": "unix",
        "local": "/run/systemd/coredump",
        "status": "NONE"
      },
      {
        "fd": 7,
        "pid": 2129,
        "type": "file",
        "local": "/run/user/975/keyring/ssh",
        "status": "NONE"
      },
      {
        "fd": 11,
        "pid": 2129,
        "type": "file",
        "local": "/run/user/975/keyring/pkcs11",
        "status": "NONE"
      },
      {
        "fd": 5,
        "pid": 908,
        "type": "file",
        "local": "/var/run/lsm/ipc/simc",
        "status": "NONE"
      },
      {
        "fd": 6,
        "pid": 908,
        "type": "file",
        "local": "/var/run/lsm/ipc/sim",
        "status": "NONE"
      },
      {
        "fd": 9,
        "pid": 913,
        "type": "file",
        "local": "/var/run/vmware/guestServicePipe",
        "status": "NONE"
      },
      {
        "fd": 47,
        "pid": 1,
        "type": "file",
        "local": "/run/avahi-daemon/socket",
        "status": "NONE"
      },
      {
        "fd": 35,
        "pid": 1,
        "type": "file",
        "local": "/run/dbus/system_bus_socket",
        "status": "NONE"
      },
      {
        "fd": 153,
        "pid": 1,
        "type": "file",
        "local": "/var/run/.heim_org.h5l.kcm-socket",
        "status": "NONE"
      },
      {
        "fd": 9,
        "pid": 1005,
        "type": "file",
        "local": "/run/gssproxy.sock",
        "status": "NONE"
      },
      {
        "fd": 142,
        "pid": 1,
        "type": "file",
        "local": "/run/libvirt/virtlockd-sock",
        "status": "NONE"
      },
      {
        "fd": 14,
        "pid": 1056,
        "type": "file",
        "local": "@/tmp/dbus-Eyr2TPFi",
        "status": "NONE"
      },
      {
        "fd": 5,
        "pid": 1848,
        "type": "file",
        "local": "/tmp/.X11-unix/X1024",
        "status": "NONE"
      },
      {
        "fd": 26,
        "pid": 1823,
        "type": "file",
        "local": "/run/user/975/wayland-0",
        "status": "NONE"
      },
      {
        "fd": 33,
        "pid": 1,
        "type": "file",
        "local": "@ISCSIADM_ABSTRACT_NAMESPACE",
        "status": "NONE"
      },
      {
        "fd": 4,
        "pid": 1848,
        "type": "file",
        "local": "@/tmp/.X11-unix/X1024",
        "status": "NONE"
      },
      {
        "fd": 45,
        "pid": 1,
        "type": "file",
        "local": "/run/libvirt/virtlogd-sock",
        "status": "NONE"
      },
      {
        "fd": 38,
        "pid": 1,
        "type": "file",
        "local": "/var/run/cups/cups.sock",
        "status": "NONE"
      },
      {
        "fd": 56,
        "pid": 1,
        "type": "file",
        "local": "/run/libvirt/libvirt-sock",
        "status": "NONE"
      },
      {
        "fd": 13,
        "pid": 909,
        "type": "file",
        "local": "/var/lib/sss/pipes/private/sbus-monitor",
        "status": "NONE"
      },
      {
        "fd": 36,
        "pid": 1,
        "type": "file",
        "local": "/run/libvirt/libvirt-sock-ro",
        "status": "NONE"
      },
      {
        "fd": 143,
        "pid": 1,
        "type": "file",
        "local": "/run/libvirt/libvirt-admin-sock",
        "status": "NONE"
      },
      {
        "fd": 12,
        "pid": 1056,
        "type": "file",
        "local": "@/tmp/dbus-CI9Y4MpI",
        "status": "NONE"
      },
      {
        "fd": 12,
        "pid": 1679,
        "type": "file",
        "local": "/tmp/.ICE-unix/1679",
        "status": "NONE"
      },
      {
        "fd": 26,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/notify",
        "status": "NONE"
      },
      {
        "fd": 27,
        "pid": 1,
        "type": "file",
        "local": "/run/systemd/cgroups-agent",
        "status": "NONE"
      },
      {
        "fd": 6,
        "pid": 1073,
        "type": "file",
        "local": "@00009",
        "status": "NONE"
      },
      {
        "fd": 13,
        "pid": 1056,
        "type": "file",
        "local": "@/tmp/dbus-822bBRfw",
        "status": "NONE"
      },
      {
        "fd": 1,
        "pid": 51892,
        "type": "file",
        "status": "NONE"
      },
      {
        "fd": 3,
        "pid": 1930,
        "type": "file",
        "status": "NONE"
      }
    ]
  }
}
```