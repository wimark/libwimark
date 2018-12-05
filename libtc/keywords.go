package libtc

const (
	TC_EXECUTABLE = "tc"
	IP_EXECUTABLE = "ip"

	QDISC_OBJECT  = "qdisc"
	CLASS_OBJECT  = "class"
	FILTER_OBJECT = "filter"
	LINK_OBJECT   = "link"

	ADD_COMMAND  = "add"
	DEL_COMMAND  = "del"
	SHOW_COMMAND = "show"

	DEV_PARAM     = "dev"
	PARENT_PARAM  = "parent"
	HANDLE_PARAM  = "handle"
	PRIO_PARAM    = "pref"
	PROTO_PARAM   = "protocol"
	CLASSID_PARAM = "classid"
	ACTION_PARAM  = "action"
	LEAF_PARAM    = "leaf"
	DEFAULT_PARAM = "default"

	U32_HANDLE_PARAM = "fh"
	U32_TABLE_PARAM  = "ht"
	U32_BUCKET_PARAM = "bkt"
	U32_DIV_PARAM    = "divisor"
	U32_FLOW_PARAM   = "flowid"
	U32_LINK_PARAM   = "link"
	U32_MATCH_PARAM  = "match"
	U32_HASH_PARAM   = "hash"
	U32_SAMPLE_PARAM = "sample"

	FW_HANDLE_PARAM = "handle"
	FW_FLOW_PARAM   = "classid"

	HTB_RATE_PARAM = "rate"
	HTB_CEIL_PARAM = "ceil"

	FILTER_U32 = "u32"
	FILTER_FW  = "fw"

	QDISC_HTB = "htb"

	ROOT_VALUE    = "root"
	INGRESS_VALUE = "ingress"

	SENT_PARAM  = "Sent"
	BYTES_PARAM = "bytes"
	PKT_PARAM   = "pkt"
)
