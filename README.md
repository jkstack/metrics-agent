# metrics-agent

系统状态监控agent

## 兼容性

go语言最低支持linux内核版本`2.6.23`

| 系统   | 版本号 | 内核版本 | 是否支持 |
| :----: | :---: | :-----: | ------- |
| ubuntu | 12.04 | [3.2](https://en.wikipedia.org/wiki/Ubuntu_version_history#Table_of_versions) | ✅ [数据示例](docs/examples/ubuntu12.md) |
| ubuntu | 14.04 | 3.13 | ✅ [数据示例](docs/examples/ubuntu14.md) |
| ubuntu | 16.04 | 4.4  | ✅ [数据示例](docs/examples/ubuntu16.md) |
| ubuntu | 18.04 | 4.15 | ✅ [数据示例](docs/examples/ubuntu18.md) |
| ubuntu | 20.04 | 5.4  | ✅ [数据示例](docs/examples/ubuntu20.md) |
| ubuntu | 22.04 | 5.15 | ✅ [数据示例](docs/examples/ubuntu22.md) |
| suse   | 10(SP4) | [2.6.16.60-0.132.1](https://www.suse.com/support/kb/doc/?id=000019587) | ❌ |
| suse   | 11(SP4) | 3.0.101-108.135.1      | ✅ [数据示例](docs/examples/suse11.md) |
| suse   | 12(SP5) | 4.12.14-122.130.1      | ✅ |
| suse   | 15(SP4) | 5.14.21-150400.24.18.1 | ✅ |
| redhat | 5.11 | [2.6.18-398](https://access.redhat.com/articles/3078) | ❌ |
| redhat | 6.1  | 2.6.32-754     | ✅ |
| redhat | 7.9  | 3.10.0-1160    | ✅ |
| redhat | 8.6  | 4.18.0-372.9.1 | ✅ |
| centos | 5.11 | [2.6.18-398](https://vault.centos.org/5.11/os/Source/) | ❌ |
| centos | 6.1  | [2.6.32-754](https://vault.centos.org/6.10/os/Source/SPackages/) | ✅ [数据示例](docs/examples/centos6.md) |
| centos | 7.9-2009 | [3.10.0-1160](https://vault.centos.org/7.9.2009/os/Source/SPackages/) | ✅ [数据示例](docs/examples/centos7.md) |
| centos | 8.5.2111 | [4.18.0-348](https://vault.centos.org/8.5.2111/BaseOS/Source/SPackages/) | ✅ [数据示例](docs/examples/centos8.md) |
| windows | 7 Enterprise with Service Pack 1             | | ✅ |
| windows | 7 Professional with Service Pack 1, VL Build | | ✅ |
| windows | 10 Enterprise LTSC 2021                      | | ✅ |
| windows | 2008 R2 Enterprise                           | | ✅ |
| windows | 2008 R2 Datacenter                           | | ✅ |
| windows | 2016 VL                                      | | ✅ |
| windows | 2016(Updated Feb 2018)                       | | ✅ |
| windows | 2016 Essentials                              | | ✅ |
| windows | 2019(Updated July 2020)                      | | ✅ |
| windows | 2022(updated Aug 2022)                       | | ✅ |
