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
| suse   | 12(SP5) | 4.12.14-122.130.1      | ✅ [数据示例](docs/examples/suse12.md) |
| suse   | 15(SP4) | 5.14.21-150400.24.18.1 | ✅ [数据示例](docs/examples/suse15.md) |
| redhat | 5.11 | [2.6.18-398](https://access.redhat.com/articles/3078) | ❌ |
| redhat | 6.1  | 2.6.32-754     | ✅ [数据示例](docs/examples/redhat6.md) |
| redhat | 7.9  | 3.10.0-1160    | ✅ [数据示例](docs/examples/redhat7.md) |
| redhat | 8.6  | 4.18.0-372.9.1 | ✅ [数据示例](docs/examples/redhat8.md) |
| centos | 5.11 | [2.6.18-398](https://vault.centos.org/5.11/os/Source/) | ❌ |
| centos | 6.1  | [2.6.32-754](https://vault.centos.org/6.10/os/Source/SPackages/) | ✅ [数据示例](docs/examples/centos6.md) |
| centos | 7.9-2009 | [3.10.0-1160](https://vault.centos.org/7.9.2009/os/Source/SPackages/) | ✅ [数据示例](docs/examples/centos7.md) |
| centos | 8.5.2111 | [4.18.0-348](https://vault.centos.org/8.5.2111/BaseOS/Source/SPackages/) | ✅ [数据示例](docs/examples/centos8.md) |
| windows | XP | | ❌ |
| windows | 7 Enterprise with Service Pack 1             | | ✅ [数据示例](docs/examples/win7enterprise.md) |
| windows | 7 Professional with Service Pack 1, VL Build | | ✅ [数据示例](docs/examples/win7professional.md) |
| windows | 10 Enterprise LTSC 2021                      | | ✅ [数据示例](docs/examples/win10.md) |
| windows | 2008 R2 Enterprise                           | | ✅ [数据示例](docs/examples/win2008ent.md) |
| windows | 2008 R2 Datacenter                           | | ✅ [数据示例](docs/examples/win2008dc.md) |
| windows | 2016 VL                                      | | ✅ [数据示例](docs/examples/win2016vl.md) |
| windows | 2016(Updated Feb 2018)                       | | ✅ [数据示例](docs/examples/win2016upd2018.md) |
| windows | 2016 Essentials                              | | ✅ [数据示例](docs/examples/win2016ess.md) |
| windows | 2019(Updated July 2020)                      | | ✅ [数据示例](docs/examples/win2019.md) |
| windows | 2022(updated Aug 2022)                       | | ✅ [数据示例](docs/examples/win2022.md) |

## linux系统部署

1. 根据当前操作系统下载`deb`或`rpm`安装包，[下载地址](https://github.com/jkstack/metrics-agent/releases/latest)
2. 使用`rpm`或`dpkg`命令安装该软件包，程序将被安装到`/opt/metrics-agent`目录下
3. 按需修改配置文件，配置文件将被安装在`/opt/metrics-agent/conf/agent.conf`目录下，建议修改以下配置内容
   - basic.id: 客户端ID，在该集群下不可重复
   - basic.server: 服务器端地址
4. 使用以下命令启动客户端程序

       /opt/metrics-agent/bin/metrics-agent -action start
5. 检查当前客户端是否连接成功

       curl http://<服务端IP>:<端口号(默认13081)>/api/agents/<客户端ID>

## windows系统部署

1. 根据当前操作系统下载`exe`或`msi`安装包，[下载地址](https://github.com/jkstack/metrics-agent/releases/latest)
2. 安装该安装包，程序默认会被安装到`C:\Program Files (x86)\metrics-agent`目录下
3. 按需修改配置文件，配置文件将默认被安装在`C:\Program Files (x86)\metrics-agent\conf\agent.conf`目录下，建议修改以下配置内容
   - basic.id: 客户端ID，在该集群下不可重复
   - basic.server: 服务器端地址
4. 使用以下命令打开系统服务管理面板，找到`metrics-agent`服务并启动

       services.msc
5. 检查当前客户端是否连接成功

       curl http://<服务端IP>:<端口号(默认13081)>/api/agents/<客户端ID>