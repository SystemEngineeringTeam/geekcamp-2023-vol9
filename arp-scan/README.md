# ARP SCAN
このプログラムは BSD 3-Clause "New" or "Revised" Licenseに基づいて google/gopacket の「arpscan.go」を改変、利用しています
https://github.com/google/gopacket/blob/master/examples/arpscan/arpscan.go

## 使用方法
1. `ifconfig` を実行し、使用しているネットワークインタフェースを確認
1. `scan-interfase/interfase.go` を実行し、1で確認したインタフェースのindexとmacアドレスを確認
1. .envにインタフェース名、index、macアドレスをそれぞれ記入
1. network.csvに部屋番号、ネットワークアドレス、ネットワークマスクをそれぞれ記入
1. OS側で、`/etc/systemd/system/arpscan.timer`,`/etc/systemd/system/arpscan.service`に以下のように記述する

```arpscan.timer
[Unit]
Description=ArpScan Timer

[Timer]
OnCalendar=Mon..Fri 9..16:0/30
Unit=arpscan.service

[Install]
WantedBy=timers.target
```

```arpscan.service
[Unit]
Description=ArpScan Service

[Service]
User=root
Type=oneshot
WorkingDirectory=<pass>/geekcamp-2023-vol9/arp-scan
ExecStart=go run main.go
```

6. systemdデーモンの再度読み込み
```sudo systemctl deamon-reload```

1. タイマーを有効にして開始
```sudo systemctl enable arpscan.timer```
```sudo systemctl start arpscan.timer```

1. タイマーが実行されていることを確認
```systemctl list-timers```