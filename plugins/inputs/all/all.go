package all

import (
	_ "github.com/influxdata/telegraf/plugins/inputs/exec"
	_ "github.com/influxdata/telegraf/plugins/inputs/filestat"
	_ "github.com/influxdata/telegraf/plugins/inputs/http_listener"
	_ "github.com/influxdata/telegraf/plugins/inputs/http_response"
	_ "github.com/influxdata/telegraf/plugins/inputs/httpjson"
	_ "github.com/influxdata/telegraf/plugins/inputs/mysql"
	_ "github.com/influxdata/telegraf/plugins/inputs/net_response"
	_ "github.com/influxdata/telegraf/plugins/inputs/nginx"
	_ "github.com/influxdata/telegraf/plugins/inputs/ping"
	_ "github.com/influxdata/telegraf/plugins/inputs/procstat"
	_ "github.com/influxdata/telegraf/plugins/inputs/socket_listener"
	_ "github.com/influxdata/telegraf/plugins/inputs/system"
	_ "github.com/influxdata/telegraf/plugins/inputs/tcp_listener"
	_ "github.com/influxdata/telegraf/plugins/inputs/webhooks"
)
