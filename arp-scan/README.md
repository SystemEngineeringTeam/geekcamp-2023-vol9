# ARP SCAN
別のプロジェクトで扱う用の機能

このプログラムは BSD 3-Clause "New" or "Revised" Licenseに基づいて google/gopacket の「arpscan.go」を改変、利用しています
https://github.com/google/gopacket/blob/master/examples/arpscan/arpscan.go

## 使用方法
1. `ifconfig` を実行し、使用しているネットワークインタフェースを確認
1. `scan-interfase/interfase.go` を実行し、1で確認したインタフェースのindexとmacアドレスを確認
1. .envにインタフェース名、index、macアドレスとスキャンしたいネットワークのネットワークアドレスをそれぞれ記入
1. `arpscan.go` を実行