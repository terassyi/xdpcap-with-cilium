#include "hook.h"
#include "bpf_helpers.h"

BPF_MAP_DEF(xdpcap_hook) = {
	.map_type = BPF_MAP_TYPE_PROG_ARRAY,
	.key_size = sizeof(int),
	.value_size = sizeof(int),
	.max_entries = 5,
};
BPF_MAP_ADD(xdpcap_hook);


SEC("xdp")
int prog(struct xdp_md *ctx) {
	return xdpcap_exit(ctx, &xdpcap_hook, XDP_PASS);
}
